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

package credential

import (
    "context"
    "encoding/json"
    "fmt"
    "sync"
    "time"
    
    "github.com/houseme/bytedance/config"
    "github.com/houseme/bytedance/utility/base"
    "github.com/houseme/bytedance/utility/cache"
    "github.com/houseme/bytedance/utility/logger"
    "github.com/houseme/bytedance/utility/request"
)

const (
    refreshTokenURL      = "https://open.douyin.com/oauth/refresh_token?client_key=%s&grant_type=refresh_token&refresh_token=%s"
    renewRefreshTokenURL = "https://open.douyin.com/oauth/renew_refresh_token?client_key=%s&refresh_token=%s"
    clientTokenURL       = "https://open.douyin.com/oauth/client_token?client_key=%s&client_secret=%s&grant_type=client_credential"
    serverAccessTokenURL = "https://developer.toutiao.com/api/apps/v2/token"
)

// DefaultAccessToken 默认 AccessToken 获取
type DefaultAccessToken struct {
    ClientKey       string
    ClientSecret    string
    cacheKeyPrefix  string
    cache           cache.Cache
    request         request.Request
    logger          logger.ILogger
    accessTokenLock *sync.Mutex
}

// NewDefaultAccessToken new DefaultAccessToken
func NewDefaultAccessToken(_ context.Context, cfg *config.Config) AccessTokenHandle {
    return &DefaultAccessToken{
        ClientKey:       cfg.ClientKey(),
        ClientSecret:    cfg.ClientSecret(),
        cache:           cfg.Cache(),
        request:         cfg.Request(),
        logger:          cfg.Logger(),
        cacheKeyPrefix:  cfg.CacheKeyPrefix(),
        accessTokenLock: new(sync.Mutex),
    }
}

// AccessToken struct
type AccessToken struct {
    base.CommonError
    AccessToken    string `json:"access_token"`
    ExpiresIn      int64  `json:"expires_in"`
    RefreshToken   string `json:"refresh_token"`
    RefreshTokenIn int64  `json:"refresh_expires_in"`
    OpenID         string `json:"openid"`
    Scope          string `json:"scope"`
}

// GetAccessToken 获取 access_token，先从 cache 中获取，没有则从服务端获取
func (t *DefaultAccessToken) GetAccessToken(ctx context.Context, openID string) (accessToken string, err error) {
    accessTokenCacheKey := fmt.Sprintf("%s_access_token_%s", t.cacheKeyPrefix, openID)
    if val := t.cache.Get(ctx, accessTokenCacheKey); val != nil {
        if accessToken = val.(string); accessToken != "" {
            return
        }
    }
    
    // 加上 lock，是为了防止在并发获取 token 时，cache 刚好失效，导致从抖音服务器上获取到不同 token
    t.accessTokenLock.Lock()
    defer t.accessTokenLock.Unlock()
    
    // 双检，防止重复从微信服务器获取
    if val := t.cache.Get(ctx, accessTokenCacheKey); val != nil {
        if accessToken = val.(string); accessToken != "" {
            return
        }
    }
    
    // 刷新 AccessToken
    refreshToken := t.cache.Get(ctx, fmt.Sprintf("%s_refresh_token_%s", t.cacheKeyPrefix, openID))
    if refreshToken == nil {
        err = fmt.Errorf("user need auth")
        return
    }
    
    var resAccessToken *AccessToken
    if resAccessToken, err = t.RefreshAccessToken(ctx, refreshToken.(string)); err != nil {
        return
    }
    
    // 缓存 AccessToken
    if err = t.SetAccessToken(ctx, resAccessToken); err != nil {
        return
    }
    
    accessToken = resAccessToken.AccessToken
    return
}

// SetAccessToken 设置 access_token
func (t *DefaultAccessToken) SetAccessToken(ctx context.Context, accessToken *AccessToken) (err error) {
    // access token cache
    if err = t.cache.Set(ctx, t.accessTokenKey(accessToken.OpenID), accessToken.AccessToken, time.Duration(accessToken.ExpiresIn-1500)*time.Second); err != nil {
        return
    }
    
    // refresh access token cache
    if err = t.cache.Set(ctx, t.refreshAccessTokenKey(accessToken.OpenID), accessToken.RefreshToken, time.Duration(accessToken.RefreshTokenIn-1500)*time.Second); err != nil {
        return
    }
    
    return
}

// accessTokenKey 获取 access_token 的 key
func (t *DefaultAccessToken) accessTokenKey(openID string) string {
    return fmt.Sprintf("%s_access_token_%s", t.cacheKeyPrefix, openID)
}

// refreshAccessTokenKey refresh token
func (t *DefaultAccessToken) refreshAccessTokenKey(openID string) string {
    return fmt.Sprintf("%s_refresh_token_%s", t.cacheKeyPrefix, openID)
}

type accessTokenRes struct {
    Message string                `json:"message"`
    Extra   base.CommonErrorExtra `json:"extra"`
    Data    AccessToken           `json:"data"`
}

// RefreshAccessToken 刷新 AccessToken.
// 当 access_token 过期（过期时间 15 天）后，可以通过该接口使用 refresh_token（过期时间 30 天）进行刷新
func (t *DefaultAccessToken) RefreshAccessToken(ctx context.Context, refreshToken string) (accessToken *AccessToken, err error) {
    var response []byte
    if response, err = t.request.Get(ctx, fmt.Sprintf(refreshTokenURL, t.ClientKey, refreshToken)); err != nil {
        return
    }
    var result accessTokenRes
    if err = json.Unmarshal(response, &result); err != nil {
        return
    }
    
    if result.Data.ErrCode != 0 {
        err = fmt.Errorf("GetUserAccessToken error : errcode=%v , errmsg=%v", result.Data.ErrCode, result.Data.ErrMsg)
        return
    }
    
    if err = t.SetAccessToken(ctx, &result.Data); err != nil {
        return
    }
    accessToken = &result.Data
    return
}

// RefreshToken refresh token
type RefreshToken struct {
    base.CommonError
    ExpiresIn    int64  `json:"expires_in"`
    RefreshToken string `json:"refresh_token"`
}

type refreshTokenRes struct {
    Message string                `json:"message"`
    Extra   base.CommonErrorExtra `json:"extra"`
    Data    RefreshToken          `json:"data"`
}

// RenewRefreshToken 刷新 refresh_token.
// 前提：client_key 需要具备 renew_refresh_token 这个权限
// 接口说明：可以通过旧的 refresh_token 获取新的 refresh_token，调用后旧 refresh_token 会失效，新 refresh_token 有 30 天有效期。最多只能获取 5 次新的 refresh_token，5 次过后需要用户重新授权。
func (t *DefaultAccessToken) RenewRefreshToken(ctx context.Context, refreshToken string) (refreshTokenData *RefreshToken, err error) {
    var response []byte
    if response, err = t.request.Get(ctx, fmt.Sprintf(renewRefreshTokenURL, t.ClientKey, refreshToken)); err != nil {
        return
    }
    var result refreshTokenRes
    if err = json.Unmarshal(response, &result); err != nil {
        return
    }
    
    if result.Data.ErrCode != 0 {
        err = fmt.Errorf("RenewRefreshToken error : errcode=%v , errmsg=%v", result.Data.ErrCode, result.Data.ErrMsg)
        return
    }
    refreshTokenData = &result.Data
    return
}

// ClientToken client token
type ClientToken struct {
    base.CommonError
    AccessToken string `json:"access_token"`
    ExpiresIn   int64  `json:"expires_in"`
}

// ServerAccessToken client token
type ServerAccessToken struct {
    ErrNo       int64  `json:"err_no"`
    ErrTips     string `json:"err_tips"`
    AccessToken string `json:"access_token"`
    ExpiresIn   int64  `json:"expires_in"`
}

type clientTokenRes struct {
    Message string      `json:"message"`
    Data    ClientToken `json:"data"`
}

// GetClientToken 该接口用于获取接口调用的凭证 client_access_token，主要用于调用不需要用户授权就可以调用的接口。
func (t *DefaultAccessToken) GetClientToken(ctx context.Context) (clientToken *ClientToken, err error) {
    if val := t.cache.Get(ctx, t.clientTokenKey()); val != nil {
        if accessToken := val.(string); accessToken != "" {
            clientToken = &ClientToken{
                AccessToken: accessToken,
            }
            return
        }
    }
    
    // 加上 lock，是为了防止在并发获取 token 时，cache 刚好失效，导致从抖音服务器上获取到不同 token
    t.accessTokenLock.Lock()
    defer t.accessTokenLock.Unlock()
    
    // 双检，防止重复从微信服务器获取
    if val := t.cache.Get(ctx, t.clientTokenKey()); val != nil {
        if accessToken := val.(string); accessToken != "" {
            clientToken = &ClientToken{
                AccessToken: accessToken,
            }
            return
        }
    }
    
    var (
        response []byte
        param    = map[string]string{
            "grant_type": "client_credential",
            "appid":      t.ClientKey,
            "secret":     t.ClientSecret,
        }
        data []byte
    )
    if data, err = json.Marshal(param); err != nil {
        return
    }
    if response, err = t.request.Post(ctx, fmt.Sprintf(clientTokenURL, t.ClientKey, t.ClientSecret), data); err != nil {
        return
    }
    var result clientTokenRes
    if err = json.Unmarshal(response, &result); err != nil {
        return
    }
    
    if result.Data.ErrCode != 0 {
        err = fmt.Errorf("get client token error : errcode=%v , errmsg=%v", result.Data.ErrCode, result.Data.ErrMsg)
        return
    }
    if err = t.SetClientToken(ctx, &result.Data); err != nil {
        return
    }
    clientToken = &result.Data
    return
}

// SetClientToken 设置 client_token
func (t *DefaultAccessToken) SetClientToken(ctx context.Context, clientToken *ClientToken) (err error) {
    // access token cache
    if err = t.cache.Set(ctx, t.clientTokenKey(), clientToken.AccessToken, time.Duration(clientToken.ExpiresIn-1500)*time.Second); err != nil {
        return
    }
    return
}

// clientTokenKey 获取 client_token 的 key
func (t *DefaultAccessToken) clientTokenKey() string {
    return fmt.Sprintf("%s_client_token_%s", t.cacheKeyPrefix, t.ClientKey)
}

// serverAccessTokenKey 获取 server_access_token 的 key
func (t *DefaultAccessToken) serverAccessTokenKey() string {
    return fmt.Sprintf("%s_server_client_token_%s", t.cacheKeyPrefix, t.ClientKey)
}

// SetServerAccessToken 设置 client_token
func (t *DefaultAccessToken) SetServerAccessToken(ctx context.Context, serverAccessToken *ServerAccessToken) (err error) {
    // access token cache
    if err = t.cache.Set(ctx, t.clientTokenKey(), serverAccessToken.AccessToken, time.Duration(serverAccessToken.ExpiresIn-1500)*time.Second); err != nil {
        return
    }
    return
}

// GetServerAccessToken 该接口用于获取接口调用的凭证 client_access_token，主要用于调用不需要用户授权就可以调用的接口。
func (t *DefaultAccessToken) GetServerAccessToken(ctx context.Context) (serverAccessToken *ServerAccessToken, err error) {
    if val := t.cache.Get(ctx, t.serverAccessTokenKey()); val != nil {
        if accessToken := val.(string); accessToken != "" {
            serverAccessToken = &ServerAccessToken{
                AccessToken: accessToken,
            }
            return
        }
    }
    
    // 加上 lock，是为了防止在并发获取 token 时，cache 刚好失效，导致从抖音服务器上获取到不同 token
    t.accessTokenLock.Lock()
    defer t.accessTokenLock.Unlock()
    
    // 双检，防止重复从微信服务器获取
    if val := t.cache.Get(ctx, t.serverAccessTokenKey()); val != nil {
        if accessToken := val.(string); accessToken != "" {
            serverAccessToken = &ServerAccessToken{
                AccessToken: accessToken,
            }
            return
        }
    }
    
    var (
        response []byte
        param    = map[string]string{
            "grant_type": "client_credential",
            "appid":      t.ClientKey,
            "secret":     t.ClientSecret,
        }
        data []byte
    )
    if data, err = json.Marshal(param); err != nil {
        return
    }
    if response, err = t.request.Post(ctx, serverAccessTokenURL, data); err != nil {
        return
    }
    serverAccessToken = &ServerAccessToken{}
    if err = json.Unmarshal(response, serverAccessToken); err != nil {
        return
    }
    
    if serverAccessToken.ErrNo != 0 {
        err = fmt.Errorf("get client token error : errcode=%v , errmsg=%v", serverAccessToken.ErrNo, serverAccessToken.ErrTips)
        return
    }
    if err = t.SetServerAccessToken(ctx, serverAccessToken); err != nil {
        return
    }
    return
}
