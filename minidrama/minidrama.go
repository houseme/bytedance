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

// Package minidrama mini drama
package minidrama

import (
    "context"
    
    "github.com/houseme/bytedance/config"
    "github.com/houseme/bytedance/credential"
    "github.com/houseme/bytedance/minidrama/drama"
    "github.com/houseme/bytedance/minidrama/voc"
    "github.com/houseme/bytedance/utility/base"
)

// MiniDrama mini drama
type MiniDrama struct {
    ctxCfg *credential.ContextConfig
}

// New return new mini drama
func New(ctx context.Context, cfg *config.Config) (*MiniDrama, error) {
    if cfg == nil {
        return nil, base.ErrConfigNotFound
    }
    if cfg.ClientKey() == "" || cfg.ClientSecret() == "" {
        return nil, base.ErrConfigKeyValueEmpty("clientKey or clientSecret")
    }
    if cfg.Salt() == "" {
        return nil, base.ErrConfigKeyValueEmpty("salt")
    }
    
    if cfg.Token() == "" {
        return nil, base.ErrConfigKeyValueEmpty("token")
    }
    
    return &MiniDrama{
        ctxCfg: &credential.ContextConfig{
            Config:            cfg,
            AccessTokenHandle: credential.NewDefaultAccessToken(ctx, cfg),
        },
    }, nil
}

// ContextConfig context config
func (d *MiniDrama) ContextConfig() *credential.ContextConfig {
    return d.ctxCfg
}

// Drama return drama
func (d *MiniDrama) Drama() *drama.Drama {
    return drama.NewDrama(d.ContextConfig())
}

// Voc return voc
func (d *MiniDrama) Voc() *voc.Voc {
    return voc.NewVoc(d.ContextConfig())
}