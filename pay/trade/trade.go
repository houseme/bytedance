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

// Package trade 交易
package trade

import (
    "context"
    "encoding/base64"
    "encoding/json"
    "fmt"
    "strconv"
    "strings"
    
    "github.com/houseme/bytedance/config"
    "github.com/houseme/bytedance/credential"
    "github.com/houseme/bytedance/utility/base"
    "github.com/houseme/bytedance/utility/helper"
)

// Trade creates trade relation
type Trade struct {
    ctxCfg *credential.ContextConfig
}

// NewTrade create trade relation
func NewTrade(cfg *credential.ContextConfig) *Trade {
    return &Trade{ctxCfg: cfg}
}

// getAccessToken 获取 access_token
func (t *Trade) getAccessToken(ctx context.Context) (accessToken string, err error) {
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
func (t *Trade) setContext(ctx context.Context) (context.Context, error) {
    accessToken, err := t.getAccessToken(ctx)
    if err != nil {
        return nil, err
    }
    ctx = context.WithValue(ctx, config.AccessTokenKey, accessToken)
    return ctx, nil
}

// QueryTrade query trade relation
func (t *Trade) QueryTrade(ctx context.Context, req *QueryOrderRequest) (resp *QueryOrderResponse, err error) {
    t.ctxCfg.Logger().Debug(ctx, "QueryPay req:", req)
    
    if req.OutOrderNo == "" && req.OrderID == "" {
        return nil, base.ErrParamKeyValueEmpty("OutOrderNo and OrderID")
    }
    
    if ctx, err = t.setContext(ctx); err != nil {
        return nil, err
    }
    
    var response []byte
    if response, err = t.ctxCfg.Request().PostJSON(ctx, queryOrder, req); err != nil {
        return nil, err
    }
    resp = new(QueryOrderResponse)
    err = json.Unmarshal(response, &resp)
    return
}

// CreateTrade create trade relation
func (t *Trade) CreateTrade(ctx context.Context, req *CreateOrderRequest) (resp *CreateOrderResponse, err error) {
    t.ctxCfg.Logger().Debug(ctx, "CreatePay req:", req)
    if req.OutOrderNo == "" {
        return nil, base.ErrParamKeyValueEmpty("OutOrderNo")
    }
    
    if req.TotalAmount < 1 {
        return nil, base.ErrParamKeyValueEmpty("TotalAmount")
    }
    
    if req.SkuList == nil || len(req.SkuList) < 1 {
        return nil, base.ErrParamKeyValueEmpty("SkuList")
    }
    
    if req.OrderEntrySchema == nil {
        return nil, base.ErrParamKeyValueEmpty("OrderEntrySchema")
    }
    
    if req.PayExpireSeconds < 1 {
        req.PayExpireSeconds = defaultPayExpireSeconds
    }
    var reqByte []byte
    if reqByte, err = json.Marshal(req); err != nil {
        return
    }
    resp = &CreateOrderResponse{
        Data:              string(reqByte),
        ByteAuthorization: "",
    }
    if resp.ByteAuthorization, err = t.getByteAuthorization(t.ctxCfg.PrivateKey(), resp.Data, t.ctxCfg.ClientKey(), helper.RandomStr(10), strconv.FormatInt(helper.GetCurrTS(), 10), strconv.Itoa(t.ctxCfg.KeyVersion())); err != nil {
        return
    }
    return
}

func (t *Trade) getByteAuthorization(privateKeyStr, data, appId, nonceStr, timestamp, keyVersion string) (string, error) {
    // 读取私钥
    key, err := base64.StdEncoding.DecodeString(strings.ReplaceAll(privateKeyStr, "\n", ""))
    if err != nil {
        return "", err
    }
    
    // 生成签名
    signature, err := helper.GenSign("POST", "/requestOrder", timestamp, nonceStr, data, key, t.ctxCfg.KeyType())
    if err != nil {
        return "", err
    }
    // 构造 byteAuthorization
    byteAuthorization := fmt.Sprintf("SHA256-RSA2048 appid=%s,nonce_str=%s,timestamp=%s,key_version=%s,signature=%s", appId, nonceStr, timestamp, keyVersion, signature)
    return byteAuthorization, nil
}
