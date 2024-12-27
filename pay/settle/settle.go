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

package settle

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/houseme/bytedance/config"
	"github.com/houseme/bytedance/credential"
	"github.com/houseme/bytedance/utility/base"
)

// Settle merchant account settle
type Settle struct {
	ctxCfg *credential.ContextConfig
}

// NewSettle init
func NewSettle(cfg *credential.ContextConfig) *Settle {
	return &Settle{ctxCfg: cfg}
}

// getAccessToken 获取 access_token
func (t *Settle) getAccessToken(ctx context.Context) (accessToken string, err error) {
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
func (t *Settle) setContext(ctx context.Context) (context.Context, error) {
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

// Apply 申请结算
func (t *Settle) Apply(ctx context.Context, req *ApplySettleRequest) (resp *ApplySettleResponse, err error) {
	if req == nil {
		return nil, base.ErrRequestIsEmpty
	}

	if ctx, err = t.setContext(ctx); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = t.ctxCfg.Request().PostJSON(ctx, applySettle, *req); err != nil {
		return nil, err
	}
	resp = &ApplySettleResponse{}
	err = json.Unmarshal(response, resp)
	return
}

// Query 查询结算
func (t *Settle) Query(ctx context.Context, req *QuerySettleRequest) (resp *QuerySettleResponse, err error) {
	if req == nil {
		return nil, base.ErrRequestIsEmpty
	}

	if ctx, err = t.setContext(ctx); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = t.ctxCfg.Request().PostJSON(ctx, querySettle, *req); err != nil {
		return nil, err
	}
	resp = &QuerySettleResponse{}
	err = json.Unmarshal(response, resp)
	return
}
