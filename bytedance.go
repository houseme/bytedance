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

// Package bytedance 字节系开放平台
package bytedance

import (
	"context"

	"github.com/houseme/bytedance/config"
	"github.com/houseme/bytedance/miniprogram"
	"github.com/houseme/bytedance/payment"
	"github.com/houseme/bytedance/utility/cache"
	"github.com/houseme/bytedance/utility/logger"
	"github.com/houseme/bytedance/utility/request"
)

const (
	version = "0.0.5"
)

// Bytedance 字节系开放平台
type Bytedance struct {
	cache   cache.Cache
	request request.Request
	logger  logger.ILogger
}

// New 初始化字节系开放平台
func New(ctx context.Context) *Bytedance {
	return &Bytedance{
		cache:   cache.NewRedis(ctx, cache.NewDefaultRedisOpts()),
		request: request.NewDefaultRequest(),
		logger:  logger.NewDefaultLogger(),
	}
}

// Version return version no
func Version() string {
	return version
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

// MiniProgram mini program
func (b *Bytedance) MiniProgram(ctx context.Context, cfg *config.Config) (*miniprogram.MicroApp, error) {
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

	return miniprogram.New(ctx, cfg)
}

// Pay create payment
func (b *Bytedance) Pay(ctx context.Context, cfg *config.Config) (*payment.Pay, error) {
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

	return payment.NewPay(ctx, cfg)
}
