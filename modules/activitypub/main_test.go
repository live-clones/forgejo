// Copyright 2022 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package activitypub_test

import (
	"testing"

	"forgejo.org/models/unittest"

	_ "forgejo.org/models/repo" // repository is indirectly referenced by auth, and it has a FK to repository, so we need to load this model into the test to initialize the other models
)

func TestMain(m *testing.M) {
	unittest.MainTest(m)
}
