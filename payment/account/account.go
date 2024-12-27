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

package account

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/houseme/bytedance/credential"
	"github.com/houseme/bytedance/payment/constant"
	"github.com/houseme/bytedance/utility/base"
	"github.com/houseme/bytedance/utility/helper"
)

// Account merchant accounts
type Account struct {
	ctxCfg *credential.ContextConfig
}

// NewAccount init
func NewAccount(cfg *credential.ContextConfig) *Account {
	return &Account{ctxCfg: cfg}
}

// QueryBalance query balance
func (a *Account) QueryBalance(ctx context.Context, req *QueryMerchantAccountRequest) (res *QueryMerchantAccountResponse, err error) {
	if req == nil {
		return nil, base.ErrRequestIsEmpty
	}

	if strings.TrimSpace(req.ThirdPartyID) == "" && strings.TrimSpace(req.AppID) == "" {
		req.AppID = a.ctxCfg.Config.ClientKey()
	}
	req.Sign = helper.RequestSign(ctx, *req, a.ctxCfg.Config.Salt())
	var response []byte
	if response, err = a.ctxCfg.Request().PostJSON(ctx, constant.QueryMerchantBalance, *req); err != nil {
		return nil, err
	}
	res = &QueryMerchantAccountResponse{}
	err = json.Unmarshal(response, res)
	return
}
