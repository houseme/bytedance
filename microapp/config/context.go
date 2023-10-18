/*
 * Copyright icp-filing Author(https://houseme.github.io/bytedance/). All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * You can obtain one at https://github.com/houseme/bytedance.
 *
 */

package config

import (
	"context"

	"github.com/houseme/bytedance/credential"
	"github.com/houseme/bytedance/domain"
)

// ContextConfig 公共配置
type ContextConfig struct {
	*Config
	credential.AccessTokenHandle
}

// NewContextConfig new context config
func NewContextConfig(ctx context.Context, clientKey, clientSecret, redirectURL, scopes string) *ContextConfig {
	ctxCfg := &ContextConfig{
		Config: NewConfig(ctx, clientKey, clientSecret, redirectURL, scopes),
	}
	ctxCfg.AccessTokenHandle = credential.NewDefaultAccessToken(ctx, &domain.Config{
		ClientKey:      clientKey,
		ClientSecret:   clientSecret,
		Cache:          ctxCfg.Cache(),
		Request:        ctxCfg.Request(),
		Logger:         ctxCfg.Logger(),
		RedirectURL:    redirectURL,
		Scopes:         scopes,
		CacheKeyPrefix: credential.CacheKeyPrefix,
	})
	return ctxCfg
}

// SetAccessTokenHandle 设置 AccessTokenHandle
func (cfg *ContextConfig) SetAccessTokenHandle(accessTokenHandle credential.AccessTokenHandle) *ContextConfig {
	cfg.AccessTokenHandle = accessTokenHandle
	return cfg
}

// NewContextConfigWithConfig new context config with config
func NewContextConfigWithConfig(ctx context.Context, cfg *Config) *ContextConfig {
	ctxCfg := &ContextConfig{
		Config: cfg,
	}
	ctxCfg.AccessTokenHandle = credential.NewDefaultAccessToken(ctx, &domain.Config{
		ClientKey:      cfg.ClientKey(),
		ClientSecret:   cfg.ClientSecret(),
		Cache:          cfg.Cache(),
		Request:        cfg.Request(),
		Logger:         cfg.Logger(),
		RedirectURL:    cfg.RedirectURL(),
		Scopes:         cfg.Scopes(),
		CacheKeyPrefix: credential.CacheKeyPrefix,
	})
	return ctxCfg
}

// NewContextConfigWithAccessTokenHandle new context config with access token handle
func NewContextConfigWithAccessTokenHandle(ctx context.Context, cfg *Config, accessTokenHandle credential.AccessTokenHandle) *ContextConfig {
	return &ContextConfig{
		Config:            cfg,
		AccessTokenHandle: accessTokenHandle,
	}
}
