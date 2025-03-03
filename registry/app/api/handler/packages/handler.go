//  Copyright 2023 Harness, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package packages

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	usercontroller "github.com/harness/gitness/app/api/controller/user"
	"github.com/harness/gitness/app/auth/authn"
	"github.com/harness/gitness/app/auth/authz"
	corestore "github.com/harness/gitness/app/store"
	urlprovider "github.com/harness/gitness/app/url"
	"github.com/harness/gitness/registry/app/api/controller/metadata"
	"github.com/harness/gitness/registry/app/api/handler/utils"
	artifact2 "github.com/harness/gitness/registry/app/api/openapi/contracts/artifact"
	"github.com/harness/gitness/registry/app/dist_temp/errcode"
	"github.com/harness/gitness/registry/app/pkg"
	"github.com/harness/gitness/registry/app/pkg/commons"
	"github.com/harness/gitness/registry/app/store"
	"github.com/harness/gitness/types/enum"

	"github.com/rs/zerolog/log"
)

const (
	packageNameRegex = `^[a-zA-Z0-9][a-zA-Z0-9._-]*[a-zA-Z0-9]$`
	versionRegex     = `^[a-z0-9][a-z0-9.-]*[a-z0-9]$`
	filenameRegex    = `^[a-zA-Z0-9][a-zA-Z0-9._~@,/-]*[a-zA-Z0-9]$`
	// Add other route types here.
)

func NewHandler(
	registryDao store.RegistryRepository,
	spaceStore corestore.SpaceStore, tokenStore corestore.TokenStore,
	userCtrl *usercontroller.Controller, authenticator authn.Authenticator,
	urlProvider urlprovider.Provider, authorizer authz.Authorizer,
) Handler {
	return &handler{
		RegistryDao:   registryDao,
		SpaceStore:    spaceStore,
		TokenStore:    tokenStore,
		UserCtrl:      userCtrl,
		Authenticator: authenticator,
		URLProvider:   urlProvider,
		Authorizer:    authorizer,
	}
}

type handler struct {
	RegistryDao   store.RegistryRepository
	SpaceStore    corestore.SpaceStore
	TokenStore    corestore.TokenStore
	UserCtrl      *usercontroller.Controller
	Authenticator authn.Authenticator
	URLProvider   urlprovider.Provider
	Authorizer    authz.Authorizer
}

type Handler interface {
	GetRegistryCheckAccess(
		ctx context.Context,
		r *http.Request,
		reqPermissions ...enum.Permission,
	) error
	GetArtifactInfo(r *http.Request) (pkg.GenericArtifactInfo, errcode.Error)
	GetAuthenticator() authn.Authenticator
}

func (h *handler) GetAuthenticator() authn.Authenticator {
	return h.Authenticator
}

func (h *handler) GetRegistryCheckAccess(
	ctx context.Context,
	r *http.Request,
	reqPermissions ...enum.Permission,
) error {
	info, _ := h.GetArtifactInfo(r)
	return pkg.GetRegistryCheckAccess(ctx, h.RegistryDao, h.Authorizer,
		h.SpaceStore,
		info.RegIdentifier, info.ParentID, reqPermissions...)
}

func (h *handler) GetArtifactInfo(r *http.Request) (pkg.GenericArtifactInfo, errcode.Error) {
	ctx := r.Context()
	path := r.URL.Path
	rootIdentifier, registryIdentifier, artifact, tag, fileName, description, err := ExtractPathVars(r)

	if err != nil {
		return pkg.GenericArtifactInfo{}, errcode.ErrCodeInvalidRequest.WithDetail(err)
	}

	if err := metadata.ValidateIdentifier(registryIdentifier); err != nil {
		return pkg.GenericArtifactInfo{}, errcode.ErrCodeInvalidRequest.WithDetail(err)
	}

	if err := validatePackageVersionAndFileName(artifact, tag, fileName); err != nil {
		return pkg.GenericArtifactInfo{}, errcode.ErrCodeInvalidRequest.WithDetail(err)
	}

	rootSpace, err := h.SpaceStore.FindByRefCaseInsensitive(ctx, rootIdentifier)
	if err != nil {
		log.Ctx(ctx).Error().Msgf("Root space not found: %s", rootIdentifier)
		return pkg.GenericArtifactInfo{}, errcode.ErrCodeRootNotFound.WithDetail(err)
	}

	registry, err := h.RegistryDao.GetByRootParentIDAndName(ctx, rootSpace.ID,
		registryIdentifier)
	if err != nil {
		log.Ctx(ctx).Error().Msgf(
			"registry %s not found for root: %s. Reason: %s", registryIdentifier, rootSpace.Identifier, err,
		)
		return pkg.GenericArtifactInfo{}, errcode.ErrCodeRegNotFound.WithDetail(err)
	}

	if registry.PackageType != artifact2.PackageTypeGENERIC {
		log.Ctx(ctx).Error().Msgf(
			"registry %s is not a generic artifact registry for root: %s", registryIdentifier, rootSpace.Identifier,
		)
		return pkg.GenericArtifactInfo{}, errcode.ErrCodeInvalidRequest.WithDetail(fmt.Errorf("registry %s is"+
			" not a generic artifact registry", registryIdentifier))
	}

	_, err = h.SpaceStore.Find(r.Context(), registry.ParentID)
	if err != nil {
		log.Ctx(ctx).Error().Msgf("Parent space not found: %d", registry.ParentID)
		return pkg.GenericArtifactInfo{}, errcode.ErrCodeParentNotFound.WithDetail(err)
	}

	info := &pkg.GenericArtifactInfo{
		ArtifactInfo: &pkg.ArtifactInfo{
			BaseInfo: &pkg.BaseInfo{
				RootIdentifier: rootIdentifier,
				RootParentID:   rootSpace.ID,
				ParentID:       registry.ParentID,
			},
			RegIdentifier: registryIdentifier,
			Image:         artifact,
		},
		RegistryID:  registry.ID,
		Version:     tag,
		FileName:    fileName,
		Description: description,
	}

	log.Ctx(ctx).Info().Msgf("Dispatch: URI: %s", path)
	if commons.IsEmpty(rootSpace.Identifier) {
		log.Ctx(ctx).Error().Msgf("ParentRef not found in context")
		return pkg.GenericArtifactInfo{}, errcode.ErrCodeParentNotFound.WithDetail(err)
	}

	if commons.IsEmpty(registryIdentifier) {
		log.Ctx(ctx).Warn().Msgf("registry not found in context")
		return pkg.GenericArtifactInfo{}, errcode.ErrCodeRegNotFound.WithDetail(err)
	}

	if !commons.IsEmpty(info.Image) && !commons.IsEmpty(info.Version) && !commons.IsEmpty(info.FileName) {
		flag, err2 := utils.MatchArtifactFilter(registry.AllowedPattern, registry.BlockedPattern,
			info.Image+":"+info.Version+":"+info.FileName)
		if !flag || err2 != nil {
			return pkg.GenericArtifactInfo{}, errcode.ErrCodeInvalidRequest.WithDetail(err2)
		}
	}

	return *info, errcode.Error{}
}

// ExtractPathVars extracts registry,image, reference, digest and tag from the path
// Path format: /generic/:rootSpace/:registry/:image/:tag (for ex:
// /generic/myRootSpace/reg1/alpine/v1).
func ExtractPathVars(r *http.Request) (
	rootIdentifier, registry, artifact,
	tag, fileName string, description string, err error,
) {
	path := r.URL.Path

	// Ensure the path starts with "/generic/"
	if !strings.HasPrefix(path, "/generic/") {
		return "", "", "", "", "", "", fmt.Errorf("invalid path: must start with /generic/")
	}

	trimmedPath := strings.TrimPrefix(path, "/generic/")
	firstSlashIndex := strings.Index(trimmedPath, "/")
	if firstSlashIndex == -1 {
		return "", "", "", "", "", "", fmt.Errorf("invalid path format: missing rootIdentifier or registry")
	}
	rootIdentifier = trimmedPath[:firstSlashIndex]

	remainingPath := trimmedPath[firstSlashIndex+1:]
	secondSlashIndex := strings.Index(remainingPath, "/")
	if secondSlashIndex == -1 {
		return "", "", "", "", "", "", fmt.Errorf("invalid path format: missing registry")
	}
	registry = remainingPath[:secondSlashIndex]

	// Extract the artifact and tag from the remaining path
	artifactPath := remainingPath[secondSlashIndex+1:]

	// Check if the artifactPath contains a ":" for tag and filename
	if strings.Contains(artifactPath, ":") {
		segments := strings.SplitN(artifactPath, ":", 3)
		if len(segments) < 3 {
			return "", "", "", "", "", "", fmt.Errorf("invalid artifact format: %s", artifactPath)
		}
		artifact = segments[0]
		tag = segments[1]
		fileName = segments[2]
	} else {
		segments := strings.SplitN(artifactPath, "/", 2)
		if len(segments) < 2 {
			return "", "", "", "", "", "", fmt.Errorf("invalid artifact format: %s", artifactPath)
		}
		artifact = segments[0]
		tag = segments[1]

		fileName = r.FormValue("filename")
		if fileName == "" {
			return "", "", "", "", "", "", fmt.Errorf("filename not provided in path or form parameter")
		}
	}
	description = r.FormValue("description")

	return rootIdentifier, registry, artifact, tag, fileName, description, nil
}

func validatePackageVersionAndFileName(packageName, version, filename string) error {
	// Compile the regular expressions
	packageNameRe := regexp.MustCompile(packageNameRegex)
	versionRe := regexp.MustCompile(versionRegex)
	filenameRe := regexp.MustCompile(filenameRegex)

	// Validate package name
	if !packageNameRe.MatchString(packageName) {
		return fmt.Errorf("invalid package name: %s", packageName)
	}

	// Validate version
	if !versionRe.MatchString(version) {
		return fmt.Errorf("invalid version: %s", version)
	}

	// Validate filename
	if !filenameRe.MatchString(filename) {
		return fmt.Errorf("invalid filename: %s", filename)
	}

	return nil
}
