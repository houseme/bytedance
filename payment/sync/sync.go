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

// Package sync order sync to douyin
package sync

import (
    "context"
    "encoding/json"
    
    "github.com/houseme/bytedance/credential"
    "github.com/houseme/bytedance/payment/constant"
    "github.com/houseme/bytedance/payment/domain"
    "github.com/houseme/bytedance/utility/base"
)

// Sync sync
type Sync struct {
    ctxCfg *credential.ContextConfig
}

// NewSync init
func NewSync(cfg *credential.ContextConfig) *Sync {
    return &Sync{ctxCfg: cfg}
}

// PushOrder push order
// see https://developer.toutiao.com/docs/miniapps/miniplatform/payment/order-sync/order-sync
func (s *Sync) PushOrder(ctx context.Context, req *domain.OrderSyncRequest) (resp *domain.OrderSyncResponse, err error) {
    if req == nil {
        return nil, base.ErrRequestIsEmpty
    }
    var clientToken *credential.ClientToken
    if clientToken, err = s.ctxCfg.GetClientToken(ctx); err != nil {
        return nil, err
    }
    if clientToken == nil {
        return nil, base.ErrClientTokenIsEmpty
    }
    req.AccessToken = clientToken.AccessToken
    
    var response []byte
    if response, err = s.ctxCfg.Request().PostJSON(ctx, constant.OrderPush, *req); err != nil {
        return nil, err
    }
    resp = &domain.OrderSyncResponse{}
    err = json.Unmarshal(response, &resp)
    return
}
