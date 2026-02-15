// Copyright 2026 The Forgejo Authors. All rights reserved.
// SPDX-License-Identifier: GPL-3.0-or-later

package auth

import (
	"context"

	"forgejo.org/models/db"
)

type AccessTokenResourceRepo struct {
	ID      int64 `xorm:"pk autoincr"`
	TokenID int64 `xorm:"NOT NULL REFERENCES(access_token, id)"` // needs to be shortened from "AccessTokenID" for the index to fit MySQL table identifier length restrictions
	RepoID  int64 `xorm:"NOT NULL REFERENCES(repository, id)"`
}

func init() {
	db.RegisterModel(new(AccessTokenResourceRepo))
}

func GetRepositoryResourcesForAccessToken(ctx context.Context, accessTokenID int64) ([]*AccessTokenResourceRepo, error) {
	var tokens []*AccessTokenResourceRepo
	err := db.GetEngine(ctx).
		Table(&AccessTokenResourceRepo{}).
		Where("token_id = ?", accessTokenID).
		Find(&tokens)
	if err != nil {
		return nil, err
	}
	return tokens, nil
}
