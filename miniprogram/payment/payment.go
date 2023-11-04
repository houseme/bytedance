/*
 * Copyright icp-filing Author(https://houseme.github.io/bytedance/). All Rights Reserved.
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

// Package payment payment
package payment

import (
    "context"
    "encoding/json"
    
    "github.com/houseme/bytedance/config"
    "github.com/houseme/bytedance/credential"
    "github.com/houseme/bytedance/miniprogram/payment/constant"
    "github.com/houseme/bytedance/miniprogram/payment/domain"
    "github.com/houseme/bytedance/utility/base"
    "github.com/houseme/bytedance/utility/helper"
)

// Pay payment
type Pay struct {
    ctxCfg *credential.ContextConfig
}

// NewPay create payment
func NewPay(cfg *config.Config) (*Pay, error) {
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
    
    return &Pay{
        ctxCfg: &credential.ContextConfig{
            Config:            cfg,
            AccessTokenHandle: credential.NewDefaultAccessToken(context.Background(), cfg),
        },
    }, nil
}

// CreatePay 创建支付
func (p *Pay) CreatePay(ctx context.Context, req *domain.CreateOrderRequest) (resp *domain.CreateOrderResponse, err error) {
    p.ctxCfg.Logger().Debug(ctx, "CreatePay req:", req)
    req.AppID = p.ctxCfg.Config.ClientKey()
    req.Sign = helper.RequestSign(ctx, *req, p.ctxCfg.Config.Salt())
    
    var response []byte
    if response, err = p.ctxCfg.Request().PostJSON(ctx, constant.CreateOrder, req); err != nil {
        return nil, err
    }
    resp = new(domain.CreateOrderResponse)
    err = json.Unmarshal(response, &resp)
    return
}

// QueryPay 查询支付
func (p *Pay) QueryPay(ctx context.Context, req *domain.QueryOrderRequest) (resp *domain.QueryOrderResponse, err error) {
    p.ctxCfg.Logger().Debug(ctx, "QueryPay req:", req)
    req.AppID = p.ctxCfg.Config.ClientKey()
    req.Sign = helper.RequestSign(ctx, *req, p.ctxCfg.Config.Salt())
    
    var response []byte
    if response, err = p.ctxCfg.Request().PostJSON(ctx, constant.QueryOrder, req); err != nil {
        return nil, err
    }
    resp = new(domain.QueryOrderResponse)
    err = json.Unmarshal(response, &resp)
    return
}
