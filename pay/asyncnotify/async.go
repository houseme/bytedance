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

package asyncnotify

import (
    "context"
    "encoding/json"
    
    "github.com/houseme/bytedance/credential"
    "github.com/houseme/bytedance/utility/helper"
)

// AsyncNotify async notify
type AsyncNotify struct {
    ctxCfg *credential.ContextConfig
}

// NewAsyncNotify init
func NewAsyncNotify(cfg *credential.ContextConfig) *AsyncNotify {
    return &AsyncNotify{ctxCfg: cfg}
}

// AsyncNotify 异步通知
func (a *AsyncNotify) AsyncNotify(ctx context.Context, req *AsyncRequest) (resp *AsyncResponse, err error) {
    a.ctxCfg.Logger().Debug(ctx, " async notify request params:", req)
    resp = &AsyncResponse{
        ErrNo:   ErrNoSuccess,
        ErrTips: ErrTipsSuccess,
        Type:    req.Type,
    }
    
    if req.Version != defaultVersion {
        resp.ErrNo = ErrNoRequestParameterError
        resp.ErrTips = "version check failed"
        return
    }
    
    var checkSignResult bool
    if checkSignResult, err = helper.CheckSign(req.ByteTimestamp, req.ByteNonceStr, req.Msg, req.ByteSignature, a.ctxCfg.Config.PrivateKey()); err != nil {
        resp.ErrNo = ErrNoSystemError
        resp.ErrTips = ErrTipsSystemError
        return
    }
    
    if !checkSignResult {
        resp.ErrNo = ErrNoFailedToCheckTheSignature
        resp.ErrTips = "failed"
    }
    if req.Type == AsyncPay {
        var data = new(PaymentData)
        if err = json.Unmarshal([]byte(req.Msg), data); err != nil {
            resp.ErrNo = ErrNoSystemError
            resp.ErrTips = ErrTipsSystemError + err.Error()
            return
        }
        resp.PaymentData = data
    }
    
    if req.Type == AsyncSettle {
        var data = new(SettleData)
        if err = json.Unmarshal([]byte(req.Msg), data); err != nil {
            resp.ErrNo = ErrNoSystemError
            resp.ErrTips = ErrTipsSystemError + err.Error()
            return
        }
        resp.SettleData = data
    }
    
    return
}
