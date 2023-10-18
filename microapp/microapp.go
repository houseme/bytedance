/*
 *  Copyright bytedance Author(https://houseme.github.io/bytedance/). All Rights Reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 *  You can obtain one at https://github.com/houseme/bytedance.
 *
 */

// Package microapp mini program
package microapp

import (
	"context"

	"github.com/houseme/bytedance/credential"
	"github.com/houseme/bytedance/domain"
	"github.com/houseme/bytedance/microapp/authorize"
	"github.com/houseme/bytedance/microapp/config"
	"github.com/houseme/bytedance/utility/base"
)

// MicroApp mini program
type MicroApp struct {
	ctxCfg *config.ContextConfig
}

// New micro app
func New(ctx context.Context, cfg *config.Config) (*MicroApp, error) {
	if cfg == nil {
		return nil, base.ErrConfigNotFound
	}
	if cfg.ClientKey() == "" || cfg.ClientSecret() == "" {
		return nil, base.ErrConfigKeyValueEmpty("clientKey or clientSecret")
	}

	if cfg.RedirectURL() == "" {
		return nil, base.ErrConfigKeyValueEmpty("redirect url")
	}

	if cfg.Scopes() == "" {
		return nil, base.ErrConfigKeyValueEmpty("scopes")
	}

	return &MicroApp{
		ctxCfg: &config.ContextConfig{
			Config: cfg,
			AccessTokenHandle: credential.NewDefaultAccessToken(ctx, &domain.Config{
				ClientKey:      cfg.ClientKey(),
				ClientSecret:   cfg.ClientSecret(),
				Cache:          cfg.Cache(),
				Request:        cfg.Request(),
				Logger:         cfg.Logger(),
				RedirectURL:    cfg.RedirectURL(),
				Scopes:         cfg.Scopes(),
				CacheKeyPrefix: credential.CacheKeyPrefix,
			}),
		},
	}, nil
}

// SetAccessTokenHandle 自定义 access_token 获取方式
func (ma *MicroApp) SetAccessTokenHandle(accessTokenHandle credential.AccessTokenHandle) {
	ma.ctxCfg.AccessTokenHandle = accessTokenHandle
}

// GetContext get Context
func (ma *MicroApp) GetContext() *config.ContextConfig {
	return ma.ctxCfg
}

// GetAccessToken 获取 access_token
func (ma *MicroApp) GetAccessToken(ctx context.Context, openID string) (string, error) {
	return ma.ctxCfg.GetAccessToken(ctx, openID)
}

// GetClientToken 获取 client_token
func (ma *MicroApp) GetClientToken(ctx context.Context) (string, error) {
	clientToken, err := ma.ctxCfg.GetClientToken(ctx)
	if err != nil {
		return "", err
	}
	return clientToken.AccessToken, nil
}

// GetAuthorize oauth2 网页授权
func (ma *MicroApp) GetAuthorize() *authorize.Authorize {
	return authorize.NewAuthorize(ma.ctxCfg)
}
