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

// Package bill query merchant bill
package bill

import (
    "context"
    "encoding/json"
    "strings"
    
    "github.com/houseme/bytedance/credential"
    "github.com/houseme/bytedance/payment/constant"
    "github.com/houseme/bytedance/utility/base"
    "github.com/houseme/bytedance/utility/helper"
)

// Bill merchant accounts bill
type Bill struct {
    ctxCfg *credential.ContextConfig
}

// NewBill init
func NewBill(cfg *credential.ContextConfig) *Bill {
    return &Bill{ctxCfg: cfg}
}

// QueryBill query bill
func (b *Bill) QueryBill(ctx context.Context, req *QueryBillRequest) (resp *QueryBillResponse, err error) {
    if req == nil {
        return nil, base.ErrRequestIsEmpty
    }
    
    if strings.TrimSpace(req.ThirdPartyID) == "" && strings.TrimSpace(req.AppID) == "" {
        req.AppID = b.ctxCfg.Config.ClientKey()
    }
    req.Sign = helper.RequestSign(ctx, *req, b.ctxCfg.Config.Salt())
    var (
        values   = req.ToURLValues()
        response []byte
    )
    
    if response, err = b.ctxCfg.Request().Get(ctx, constant.QueryMerchantBill+values.Encode()); err != nil {
        return nil, err
    }
    resp = &QueryBillResponse{}
    err = json.Unmarshal(response, resp)
    return
}
