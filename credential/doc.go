/*
 * Copyright Bytedance Author(https://houseme.github.io/bytedance/). All Rights Reserved.
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

// Package credential mini program credentials
package credential

import (
	"context"

	"github.com/houseme/bytedance/config"
)

// ContextConfig 公共配置
type ContextConfig struct {
	*config.Config
	AccessTokenHandle
}

// SetAccessTokenHandle 设置 AccessTokenHandle
func (cfg *ContextConfig) SetAccessTokenHandle(accessTokenHandle AccessTokenHandle) *ContextConfig {
	cfg.AccessTokenHandle = accessTokenHandle
	return cfg
}

// NewContextConfigWithConfig new context config with config
func NewContextConfigWithConfig(ctx context.Context, cfg *config.Config) *ContextConfig {
	ctxCfg := &ContextConfig{
		Config: cfg,
	}
	ctxCfg.AccessTokenHandle = NewDefaultAccessToken(ctx, cfg)
	return ctxCfg
}

// NewContextConfigWithAccessTokenHandle new context config with access_token handle
func NewContextConfigWithAccessTokenHandle(ctx context.Context, cfg *config.Config, accessTokenHandle AccessTokenHandle) *ContextConfig {
	return &ContextConfig{
		Config:            cfg,
		AccessTokenHandle: accessTokenHandle,
	}
}
