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

// Package config Configuration information of the DouYin Open Platform
package config

import (
	"context"

	"github.com/houseme/bytedance/utility/cache"
	"github.com/houseme/bytedance/utility/logger"
	"github.com/houseme/bytedance/utility/request"
)

const (
	// CacheKeyPrefix 抖音 open cache key 前缀
	CacheKeyPrefix = "bytedance_douyin_lite"
)

const (
	// AccessTokenKey AccessToken Key
	AccessTokenKey = "accessTokenKey"
)

// Secret defines the private key type
type Secret uint

const (
	PKCS1 Secret = iota
	PKCS8
)

// Config 抖音开放平台的配置信息
type Config struct {
	version        string
	cacheKeyPrefix string
	clientKey      string
	clientSecret   string
	redirectURL    string
	scopes         string
	token          string
	salt           string // 支付密钥值
	privateKey     string // 私钥
	publicKey      string // 公钥
	keyVersion     int    // 秘钥版本
	keyType        Secret
	cache          cache.Cache
	request        request.Request
	logger         logger.ILogger
}

type options struct {
	CacheKeyPrefix string
	ClientKey      string
	ClientSecret   string
	RedirectURL    string
	Scopes         string
	Token          string
	Salt           string // 支付密钥值
	PrivateKey     string // 私钥
	PublicKey      string // 公钥
	KeyVersion     int    // 秘钥版本
	KeyType        Secret
	Cache          cache.Cache
	Logger         logger.ILogger
	Request        request.Request
}

// Option micro app option
type Option func(*options)

// WithCacheKeyPrefix set cacheKeyPrefix
func WithCacheKeyPrefix(cacheKeyPrefix string) Option {
	return func(o *options) {
		o.CacheKeyPrefix = cacheKeyPrefix
	}
}

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

// WithToken set token
func WithToken(token string) Option {
	return func(o *options) {
		o.Token = token
	}
}

// WithSalt set salt
func WithSalt(salt string) Option {
	return func(o *options) {
		o.Salt = salt
	}
}

// WithPrivateKey set privateKey
func WithPrivateKey(privateKey string) Option {
	return func(o *options) {
		o.PrivateKey = privateKey
	}
}

// WithPublicKey set publicKey
func WithPublicKey(publicKey string) Option {
	return func(o *options) {
		o.PublicKey = publicKey
	}
}

// WithKeyVersion set keyVersion
func WithKeyVersion(keyVersion int) Option {
	return func(o *options) {
		o.KeyVersion = keyVersion
	}
}

// WithKeyType set keyType
func WithKeyType(keyType Secret) Option {
	return func(o *options) {
		o.KeyType = keyType
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

// WithCache set cache
func WithCache(cache cache.Cache) Option {
	return func(o *options) {
		o.Cache = cache
	}
}

// New create config
func New(ctx context.Context, opts ...Option) *Config {
	op := options{
		Logger:         logger.NewDefaultLogger(),
		Request:        request.NewDefaultRequest(AccessTokenKey),
		Cache:          cache.NewRedis(ctx, cache.NewDefaultRedisOpts()),
		CacheKeyPrefix: CacheKeyPrefix,
	}
	for _, option := range opts {
		option(&op)
	}

	return &Config{
		cacheKeyPrefix: op.CacheKeyPrefix,
		clientKey:      op.ClientKey,
		clientSecret:   op.ClientSecret,
		redirectURL:    op.RedirectURL,
		scopes:         op.Scopes,
		salt:           op.Salt,
		token:          op.Token,
		privateKey:     op.PrivateKey,
		publicKey:      op.PublicKey,
		keyVersion:     op.KeyVersion,
		keyType:        op.KeyType,
		request:        op.Request,
		logger:         op.Logger,
		cache:          op.Cache,
	}
}

// SetVersion 设置 version
func (cfg *Config) SetVersion(version string) *Config {
	cfg.version = version
	return cfg
}

// SetCacheKeyPrefix 设置 cacheKeyPrefix
func (cfg *Config) SetCacheKeyPrefix(cacheKeyPrefix string) *Config {
	cfg.cacheKeyPrefix = cacheKeyPrefix
	return cfg
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

// SetSalt 设置 salt
func (cfg *Config) SetSalt(salt string) *Config {
	cfg.salt = salt
	return cfg
}

// SetToken 设置 token
func (cfg *Config) SetToken(token string) *Config {
	cfg.token = token
	return cfg
}

// SetPrivateKey 设置 privateKey
func (cfg *Config) SetPrivateKey(privateKey string) *Config {
	cfg.privateKey = privateKey
	return cfg
}

// SetPublicKey 设置 publicKey
func (cfg *Config) SetPublicKey(publicKey string) *Config {
	cfg.publicKey = publicKey
	return cfg
}

// SetKeyVersion 设置 keyVersion
func (cfg *Config) SetKeyVersion(keyVersion int) *Config {
	cfg.keyVersion = keyVersion
	return cfg
}

// SetKeyType 设置 keyType
func (cfg *Config) SetKeyType(keyType Secret) *Config {
	cfg.keyType = keyType
	return cfg
}

// NewConfig new config
func NewConfig(ctx context.Context, clientKey, clientSecret, redirectURL, scopes, salt, token string) *Config {
	return &Config{
		clientKey:    clientKey,
		clientSecret: clientSecret,
		redirectURL:  redirectURL,
		scopes:       scopes,
		salt:         salt,
		token:        token,
		cache:        cache.NewRedis(ctx, cache.NewDefaultRedisOpts()),
		request:      request.NewDefaultRequest(AccessTokenKey),
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

// Version 获取 version
func (cfg *Config) Version() string {
	return cfg.version
}

// CacheKeyPrefix 获取 cacheKeyPrefix
func (cfg *Config) CacheKeyPrefix() string {
	return cfg.cacheKeyPrefix
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

// Token 获取 token
func (cfg *Config) Token() string {
	return cfg.token
}

// Salt 获取 salt
func (cfg *Config) Salt() string {
	return cfg.salt
}

// PrivateKey 获取 privateKey
func (cfg *Config) PrivateKey() string {
	return cfg.privateKey
}

// PublicKey 获取 publicKey
func (cfg *Config) PublicKey() string {
	return cfg.publicKey
}

// KeyVersion 获取 keyVersion
func (cfg *Config) KeyVersion() int {
	return cfg.keyVersion
}

// KeyType 获取 keyType
func (cfg *Config) KeyType() Secret {
	return cfg.keyType
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
