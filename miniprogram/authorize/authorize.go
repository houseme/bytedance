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

// Package authorize mini program authorize
package authorize

import (
    "context"
    "encoding/json"
    "fmt"
    "net/url"
    
    "github.com/houseme/bytedance/credential"
)

const (
    // see: https://developer.open-douyin.com/docs/resource/zh-CN/dop/develop/openapi/account-permission/douyin-get-permission-code
    redirectOauthURL  = "https://open.douyin.com/platform/oauth/connect?client_key=%s&response_type=code&scope=%s&redirect_uri=%s&state=%s" // 抖音获取授权码
    silenceOauthURL   = "https://open.douyin.com/platform/oauth/authorize/v2?client_key=%s&response_type=code&scope=login_id&redirect_uri=%s&state=%s"
    accessTokenURL    = "https://open.douyin.com/oauth/access_token?client_key=%s&client_secret=%s&code=%s&grant_type=authorization_code" // 获取用户授权调用凭证
    jsCode2sessionURL = "https://developer.toutiao.com/api/apps/v2/jscode2session"
)

// Authorize 保存用户授权信息
type Authorize struct {
    *credential.ContextConfig
}

// NewAuthorize 实例化授权信息
func NewAuthorize(ctxCfg *credential.ContextConfig) *Authorize {
    return &Authorize{
        ContextConfig: ctxCfg,
    }
}

// GetRedirectURL 获取授权码的 URL 地址
func (a *Authorize) GetRedirectURL(_ context.Context, state string) string {
    return fmt.Sprintf(redirectOauthURL, a.ClientKey(), a.Scopes(), url.QueryEscape(a.RedirectURL()), state)
}

// GetSilenceOauthURL 获取静默授权码的 URL 地址
func (a *Authorize) GetSilenceOauthURL(_ context.Context, state string) string {
    return fmt.Sprintf(silenceOauthURL, a.ClientKey(), url.QueryEscape(a.RedirectURL()), state)
}

type accessTokenRes struct {
    Message string                 `json:"message"`
    Data    credential.AccessToken `json:"data"`
}

// GetUserAccessToken 通过网页授权的 code 换取 access_token
func (a *Authorize) GetUserAccessToken(ctx context.Context, code string) (accessToken credential.AccessToken, err error) {
    var response []byte
    if response, err = a.Request().Get(ctx, fmt.Sprintf(accessTokenURL, a.ClientKey(), a.ClientSecret(), code)); err != nil {
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
    
    err = a.SetAccessToken(ctx, &result.Data)
    return
}

// CodeToSession 获取用户的 session_key 和 openid
func (a *Authorize) CodeToSession(ctx context.Context, code, anonymousCode string) (res CodeToSessionData, err error) {
    var (
        response []byte
        req      = CodeToSessionReq{
            Appid:         a.ClientKey(),
            Secret:        a.ClientSecret(),
            AnonymousCode: anonymousCode,
            Code:          code,
        }
    )
    if response, err = a.Request().PostJSON(ctx, jsCode2sessionURL, req); err != nil {
        return
    }
    var result CodeToSessionRes
    if err = json.Unmarshal(response, &result); err != nil {
        return
    }
    res = result.Data
    if result.ErrNo != 0 {
        err = fmt.Errorf("GetUserAccessToken error : errcode=%v , errmsg=%v", result.ErrNo, result.ErrTips)
    }
    return
}

// CodeToSessionReq 获取用户的 session_key 和 openid 的请求参数
type CodeToSessionReq struct {
    Appid         string `json:"appid"`
    Secret        string `json:"secret"`
    AnonymousCode string `json:"anonymous_code"`
    Code          string `json:"code"`
}

// CodeToSessionRes 获取用户的 session_key 和 openid
type CodeToSessionRes struct {
    ErrNo   int               `json:"err_no"`
    ErrTips string            `json:"err_tips"`
    Data    CodeToSessionData `json:"data"`
}

// CodeToSessionData 获取用户的 session_key 和 openid
type CodeToSessionData struct {
    SessionKey      string `json:"session_key"`
    Openid          string `json:"openid"`
    AnonymousOpenid string `json:"anonymous_openid"`
    UnionID         string `json:"unionid"`
}
