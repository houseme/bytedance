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

package solution

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/houseme/bytedance/config"
	"github.com/houseme/bytedance/credential"
	"github.com/houseme/bytedance/utility/base"
)

// Solution 解决方案
type Solution struct {
	ctxCfg *credential.ContextConfig
}

// NewSolution init
func NewSolution(cfg *credential.ContextConfig) *Solution {
	return &Solution{ctxCfg: cfg}
}

// getAccessToken 获取 access_token
func (s *Solution) getAccessToken(ctx context.Context) (accessToken string, err error) {
	var clientToken *credential.ClientToken
	if clientToken, err = s.ctxCfg.GetClientToken(ctx); err != nil {
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

// CreateSolution 创建解决方案
func (s *Solution) CreateSolution(ctx context.Context, req *CreateSolutionRequest) (resp *CreateSolutionResponse, err error) {
	if req == nil {
		return nil, base.ErrRequestIsEmpty
	}

	var accessToken string
	if accessToken, err = s.getAccessToken(ctx); err != nil {
		return nil, err
	}
	ctx = context.WithValue(ctx, config.AccessTokenKey, accessToken)
	var response []byte
	if response, err = s.ctxCfg.Request().PostJSON(ctx, setImpl+accessToken, *req); err != nil {
		return nil, err
	}

	resp = &CreateSolutionResponse{}
	err = json.Unmarshal(response, &resp)

	return
}

// QuerySolution 查询解决方案
func (s *Solution) QuerySolution(ctx context.Context, req *QuerySolutionRequest) (resp *QuerySolutionResponse, err error) {
	if req == nil {
		return nil, base.ErrRequestIsEmpty
	}

	var accessToken string
	if accessToken, err = s.getAccessToken(ctx); err != nil {
		return nil, err
	}
	ctx = context.WithValue(ctx, config.AccessTokenKey, accessToken)
	var response []byte
	if response, err = s.ctxCfg.Request().PostJSON(ctx, queryImpl+accessToken, *req); err != nil {
		return nil, err
	}

	resp = &QuerySolutionResponse{}
	err = json.Unmarshal(response, &resp)

	return
}
