// Copyright 2021 Harness Inc. All rights reserved.
// Use of this source code is governed by the Polyform Free Trial License
// that can be found in the LICENSE.md file for this repository.

package authn

import (
	"github.com/google/wire"
	"github.com/harness/gitness/internal/store"
)

// WireSet provides a wire set for this package.
var WireSet = wire.NewSet(
	ProvideAuthenticator,
)

func ProvideAuthenticator(userStore store.UserStore, saStore store.ServiceAccountStore,
	tokenStore store.TokenStore) Authenticator {
	return NewTokenAuthenticator(userStore, saStore, tokenStore)
}
