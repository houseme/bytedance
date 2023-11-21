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

package withdraw

import (
    "context"
    "encoding/json"
    "strings"
    
    "github.com/houseme/bytedance/config"
    "github.com/houseme/bytedance/credential"
    "github.com/houseme/bytedance/utility/base"
)

// Withdraw merchant accounts withdraw
type Withdraw struct {
    ctxCfg *credential.ContextConfig
}

// NewWithdraw init
func NewWithdraw(cfg *credential.ContextConfig) *Withdraw {
    return &Withdraw{ctxCfg: cfg}
}

// getAccessToken 获取 access_token
func (t *Withdraw) getAccessToken(ctx context.Context) (accessToken string, err error) {
    var clientToken *credential.ClientToken
    if clientToken, err = t.ctxCfg.GetClientToken(ctx); err != nil {
        return "", err
    }
    if clientToken == nil {
        return "", base.ErrClientTokenIsEmpty
    }
    
    if strings.TrimSpace(clientToken.AccessToken) == "" {
        return "", base.ErrClientAccessTokenIsEmpty
    }
    
    return clientToken.AccessToken, nil
}

// setContext 设置上下文
func (t *Withdraw) setContext(ctx context.Context) (context.Context, error) {
    accessToken, err := t.getAccessToken(ctx)
    if err != nil {
        return nil, err
    }
    ctx = context.WithValue(
        ctx,
        config.AccessTokenKey,
        accessToken,
    )
    return ctx, nil
}

// QueryBalance query balance
func (t *Withdraw) QueryBalance(ctx context.Context, req *QueryBalanceRequest) (resp *QueryBalanceResponse, err error) {
    if req == nil {
        return nil, base.ErrRequestIsEmpty
    }
    
    if strings.TrimSpace(req.ThirdPartyID) == "" && strings.TrimSpace(req.AppID) == "" {
        req.AppID = t.ctxCfg.Config.ClientKey()
    }
    if ctx, err = t.setContext(ctx); err != nil {
        return nil, err
    }
    t.ctxCfg.Logger().Debug(ctx, "request content:", req," request url:",queryMerchantBalance)
    var response []byte
    if response, err = t.ctxCfg.Request().PostJSON(ctx, queryMerchantBalance, *req); err != nil {
        return nil, err
    }
    t.ctxCfg.Logger().Debug(ctx, "response content:", string(response))
    resp = &QueryBalanceResponse{}
    err = json.Unmarshal(response, resp)
    return
}

// Apply to apply withdrawal
func (t *Withdraw) Apply(ctx context.Context, req *MerchantWithdrawRequest) (resp *MerchantWithdrawResponse, err error) {
    if req == nil {
        return nil, base.ErrRequestIsEmpty
    }
    
    if strings.TrimSpace(req.ThirdPartyID) == "" && strings.TrimSpace(req.AppID) == "" {
        req.AppID = t.ctxCfg.Config.ClientKey()
    }
    if ctx, err = t.setContext(ctx); err != nil {
        return nil, err
    }
    var response []byte
    if response, err = t.ctxCfg.Request().PostJSON(ctx, applyMerchantWithdraw, *req); err != nil {
        return nil, err
    }
    resp = &MerchantWithdrawResponse{}
    err = json.Unmarshal(response, resp)
    return
}

// QueryWithdraw query withdraws
func (t *Withdraw) QueryWithdraw(ctx context.Context, req *QueryMerchantWithdrawRequest) (resp *QueryMerchantWithdrawResponse, err error) {
    if req == nil {
        return nil, base.ErrRequestIsEmpty
    }
    
    if strings.TrimSpace(req.ThirdPartyID) == "" && strings.TrimSpace(req.AppID) == "" {
        req.AppID = t.ctxCfg.Config.ClientKey()
    }
    if ctx, err = t.setContext(ctx); err != nil {
        return nil, err
    }
    var response []byte
    if response, err = t.ctxCfg.Request().PostJSON(ctx, queryWithdrawOrder, *req); err != nil {
        return nil, err
    }
    resp = &QueryMerchantWithdrawResponse{}
    err = json.Unmarshal(response, resp)
    return
}
