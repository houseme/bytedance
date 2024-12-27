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

// Package trade order trading
package trade

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/houseme/bytedance/credential"
	"github.com/houseme/bytedance/payment/constant"
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
func (p *Trade) CreatePay(ctx context.Context, req *CreateOrderRequest) (resp *CreateOrderResponse, err error) {
	p.ctxCfg.Logger().Debug(ctx, "CreatePay req:", req)
	if strings.TrimSpace(req.AppID) == "" {
		req.AppID = p.ctxCfg.Config.ClientKey()
	}
	if strings.TrimSpace(req.Sign) == "" {
		req.Sign = helper.RequestSign(ctx, *req, p.ctxCfg.Config.Salt())
	}
	var response []byte
	if response, err = p.ctxCfg.Request().PostJSON(ctx, constant.CreateOrder, req); err != nil {
		return nil, err
	}
	resp = new(CreateOrderResponse)
	err = json.Unmarshal(response, &resp)
	return
}

// QueryPay 查询支付
func (p *Trade) QueryPay(ctx context.Context, req *QueryOrderRequest) (resp *QueryOrderResponse, err error) {
	p.ctxCfg.Logger().Debug(ctx, "QueryPay req:", req)
	if strings.TrimSpace(req.AppID) == "" {
		req.AppID = p.ctxCfg.Config.ClientKey()
	}
	if strings.TrimSpace(req.Sign) == "" {
		req.Sign = helper.RequestSign(ctx, *req, p.ctxCfg.Config.Salt())
	}

	var response []byte
	if response, err = p.ctxCfg.Request().PostJSON(ctx, constant.QueryOrder, req); err != nil {
		return nil, err
	}
	resp = new(QueryOrderResponse)
	err = json.Unmarshal(response, &resp)
	return
}

// AsyncNotify 异步通知
func (p *Trade) AsyncNotify(ctx context.Context, req *AsyncRequest) (resp *AsyncResponse, err error) {
	p.ctxCfg.Logger().Debug(ctx, " async notify request params:", req)
	var sign = helper.CallbackSign(ctx, p.ctxCfg.Config.Token(), *req)
	resp = &AsyncResponse{
		ErrNo:   constant.Success,
		ErrTips: "SUCCESS",
	}
	p.ctxCfg.Logger().Debug(ctx, " async notify request sign:", sign, " msg sign: ", req.MsgSignature)
	if sign != req.MsgSignature {
		resp.ErrNo = constant.FailedToCheckTheSignature
		resp.ErrTips = "failed"
	}
	return
}
