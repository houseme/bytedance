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

package bytedance

import (
	"context"

	"github.com/houseme/bytedance/microapp"
	"github.com/houseme/bytedance/microapp/config"
	"github.com/houseme/bytedance/utility/cache"
	"github.com/houseme/bytedance/utility/logger"
	"github.com/houseme/bytedance/utility/request"
)

// Bytedance 字节系开放平台
type Bytedance struct {
	cache   cache.Cache
	request request.Request
	logger  logger.ILogger
}

// New 初始化字节系开放平台
func New() *Bytedance {
	return &Bytedance{}
}

// SetCache 设置缓存
func (b *Bytedance) SetCache(cache cache.Cache) {
	b.cache = cache
}

// SetRequest 设置请求
func (b *Bytedance) SetRequest(request request.Request) {
	b.request = request
}

// SetLogger 设置日志
func (b *Bytedance) SetLogger(logger logger.ILogger) {
	b.logger = logger
}

// OpenAPI 字节系开放平台
func (b *Bytedance) OpenAPI(ctx context.Context, cfg *config.Config) (*microapp.MicroApp, error) {
	if cfg == nil {
		cfg = config.New(ctx)
	}

	if cfg.Cache() == nil {
		cfg.SetCache(b.cache)
	}

	if cfg.Request() == nil {
		cfg.SetRequest(b.request)
	}

	if cfg.Logger() == nil {
		cfg.SetLogger(b.logger)
	}

	return microapp.New(ctx, cfg)
}
