// Copyright 2022 Harness Inc. All rights reserved.
// Use of this source code is governed by the Polyform Free Trial License
// that can be found in the LICENSE.md file for this repository.

package pullreq

import (
	"context"
	"time"

	"github.com/harness/gitness/internal/api/usererror"
	"github.com/harness/gitness/internal/auth"
	"github.com/harness/gitness/internal/store/database/dbtx"
	"github.com/harness/gitness/types"
	"github.com/harness/gitness/types/enum"
)

type CreateInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`

	SourceRepoRef string `json:"source_repo_ref"`
	SourceBranch  string `json:"source_branch"`
	TargetBranch  string `json:"target_branch"`
}

// Create creates a new pull request.
func (c *Controller) Create(
	ctx context.Context,
	session *auth.Session,
	repoRef string,
	in *CreateInput,
) (*types.PullReqInfo, error) {
	now := time.Now().UnixMilli()

	targetRepo, err := c.getRepoCheckAccess(ctx, session, repoRef, enum.PermissionRepoEdit)
	if err != nil {
		return nil, err
	}

	sourceRepo := targetRepo
	if in.SourceRepoRef != "" {
		sourceRepo, err = c.getRepoCheckAccess(ctx, session, in.SourceRepoRef, enum.PermissionRepoView)
		if err != nil {
			return nil, err
		}
	}

	if sourceRepo.ID == targetRepo.ID && in.TargetBranch == in.SourceBranch {
		return nil, usererror.BadRequest("target and source branch can't be the same")
	}

	if errBranch := c.verifyBranchExistence(ctx, sourceRepo, in.SourceBranch); errBranch != nil {
		return nil, errBranch
	}
	if errBranch := c.verifyBranchExistence(ctx, targetRepo, in.TargetBranch); errBranch != nil {
		return nil, errBranch
	}

	var pr *types.PullReq

	err = dbtx.New(c.db).WithTx(ctx, func(ctx context.Context) error {
		var lastNumber int64

		lastNumber, err = c.pullreqStore.LastNumber(ctx, targetRepo.ID)
		if err != nil {
			return err
		}

		// create new pull request object
		pr = &types.PullReq{
			ID:            0, // the ID will be populated in the data layer
			CreatedBy:     session.Principal.ID,
			Created:       now,
			Updated:       now,
			Number:        lastNumber + 1,
			State:         enum.PullReqStateOpen,
			Title:         in.Title,
			Description:   in.Description,
			SourceRepoID:  sourceRepo.ID,
			SourceBranch:  in.SourceBranch,
			TargetRepoID:  targetRepo.ID,
			TargetBranch:  in.TargetBranch,
			MergedBy:      nil,
			Merged:        nil,
			MergeStrategy: nil,
		}

		return c.pullreqStore.Create(ctx, pr)
	})
	if err != nil {
		return nil, err
	}

	pri := &types.PullReqInfo{
		PullReq:     *pr,
		AuthorID:    session.Principal.ID,
		AuthorName:  session.Principal.DisplayName,
		AuthorEmail: session.Principal.Email,
	}

	return pri, nil
}