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

package trade

import (
    "context"
    "encoding/json"
    
    "github.com/houseme/bytedance/credential"
    "github.com/houseme/bytedance/payment/constant"
    "github.com/houseme/bytedance/payment/domain"
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

// CreatePay 创建支付
func (p *Trade) CreatePay(ctx context.Context, req *domain.CreateOrderRequest) (resp *domain.CreateOrderResponse, err error) {
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
func (p *Trade) QueryPay(ctx context.Context, req *domain.QueryOrderRequest) (resp *domain.QueryOrderResponse, err error) {
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

// AsyncNotify 异步通知
func (p *Trade) AsyncNotify(ctx context.Context, req *domain.AsyncRequest) (resp *domain.AsyncResponse, err error) {
    p.ctxCfg.Logger().Debug(ctx, " async notify request params:", req)
    var sign = helper.CallbackSign(ctx, p.ctxCfg.Config.Token(), *req)
    resp = &domain.AsyncResponse{
        ErrNo:   constant.FailedToCheckTheSignature,
        ErrTips: "success",
    }
    if sign != req.MsgSignature {
        resp.ErrNo = constant.FailedToCheckTheSignature
    }
    return
}
