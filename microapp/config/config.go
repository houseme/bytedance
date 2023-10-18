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

// Package config Configuration information of the DouYin Open Platform
package config

import (
	"context"

	"github.com/houseme/bytedance/microapp/credential"
	"github.com/houseme/bytedance/microapp/domain"
	"github.com/houseme/bytedance/utility/cache"
	"github.com/houseme/bytedance/utility/logger"
	"github.com/houseme/bytedance/utility/request"
)

// Config 抖音开放平台的配置信息
type Config struct {
	clientKey    string
	clientSecret string
	redirectURL  string
	scopes       string
	cache        cache.Cache
	request      request.Request
	logger       logger.ILogger
}

type options struct {
	ClientKey    string
	ClientSecret string
	RedirectURL  string
	Scopes       string
	Cache        cache.Cache
	Logger       logger.ILogger
	Request      request.Request
}

// Option micro app option
type Option func(*options)

// WithClientKey set clientKey
func WithClientKey(clientKey string) Option {
	return func(o *options) {
		o.ClientKey = clientKey
	}
}

// WithClientSecret set clientSecret
func WithClientSecret(clientSecret string) Option {
	return func(o *options) {
		o.ClientSecret = clientSecret
	}
}

// WithRedirectURL set redirectURL
func WithRedirectURL(redirectURL string) Option {
	return func(o *options) {
		o.RedirectURL = redirectURL
	}
}

// WithScopes set scopes
func WithScopes(scopes string) Option {
	return func(o *options) {
		o.Scopes = scopes
	}
}

// WithLogger set logger
func WithLogger(logger logger.ILogger) Option {
	return func(o *options) {
		o.Logger = logger
	}
}

// WithRequest set request
func WithRequest(request request.Request) Option {
	return func(o *options) {
		o.Request = request
	}
}

// New create config
func New(ctx context.Context, opts ...Option) *Config {
	op := options{
		Logger:  logger.NewDefaultLogger(),
		Request: request.NewDefaultRequest(),
		Cache:   cache.NewRedis(ctx, cache.NewDefaultRedisOpts()),
	}
	for _, option := range opts {
		option(&op)
	}

	return &Config{
		clientKey:    op.ClientKey,
		clientSecret: op.ClientSecret,
		redirectURL:  op.RedirectURL,
		scopes:       op.Scopes,
		request:      op.Request,
		logger:       op.Logger,
		cache:        op.Cache,
	}
}

// SetClientKey 设置 clientKey
func (cfg *Config) SetClientKey(clientKey string) *Config {
	cfg.clientKey = clientKey
	return cfg
}

// SetClientSecret 设置 clientSecret
func (cfg *Config) SetClientSecret(clientSecret string) *Config {
	cfg.clientSecret = clientSecret
	return cfg
}

// SetRedirectURL 设置 redirectURL
func (cfg *Config) SetRedirectURL(redirectURL string) *Config {
	cfg.redirectURL = redirectURL
	return cfg
}

// SetScopes 设置 scopes
func (cfg *Config) SetScopes(scopes string) *Config {
	cfg.scopes = scopes
	return cfg
}

// NewConfig new config
func NewConfig(ctx context.Context, clientKey, clientSecret, redirectURL, scopes string) *Config {
	return &Config{
		clientKey:    clientKey,
		clientSecret: clientSecret,
		redirectURL:  redirectURL,
		scopes:       scopes,
		cache:        cache.NewRedis(ctx, cache.NewDefaultRedisOpts()),
		request:      request.NewDefaultRequest(),
		logger:       logger.NewDefaultLogger(),
	}
}

// SetCache 设置缓存
func (cfg *Config) SetCache(cache cache.Cache) *Config {
	cfg.cache = cache
	return cfg
}

// SetRequest 设置请求
func (cfg *Config) SetRequest(request request.Request) *Config {
	cfg.request = request
	return cfg
}

// SetLogger 设置日志
func (cfg *Config) SetLogger(logger logger.ILogger) *Config {
	cfg.logger = logger
	return cfg
}

// ClientKey 获取 clientKey
func (cfg *Config) ClientKey() string {
	return cfg.clientKey
}

// ClientSecret 获取 clientSecret
func (cfg *Config) ClientSecret() string {
	return cfg.clientSecret
}

// RedirectURL 获取 redirectURL
func (cfg *Config) RedirectURL() string {
	return cfg.redirectURL
}

// Scopes 获取 scopes
func (cfg *Config) Scopes() string {
	return cfg.scopes
}

// Cache 获取 cache
func (cfg *Config) Cache() cache.Cache {
	return cfg.cache
}

// Request 获取 request
func (cfg *Config) Request() request.Request {
	return cfg.request
}

// Logger 获取 logger
func (cfg *Config) Logger() logger.ILogger {
	return cfg.logger
}

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
