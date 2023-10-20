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

// Package schema Schema link
package schema

import (
	"context"
	"encoding/json"

	"github.com/houseme/bytedance/credential"
	"github.com/houseme/bytedance/utility/base"
)

const (
	generateSchemaURL = "https://open.douyin.com/api/apps/v1/url/generate_schema"
	querySchemaURL    = "https://open.douyin.com/api/apps/v1/url/query_schema"
	querySchemaQuota  = "https://open.douyin.com/api/apps/v1/url/query_schema_quota"
)

// GenerateSchemaRequest generate schema request
type GenerateSchemaRequest struct {
	AppID      string `json:"app_id"`
	Query      string `json:"query,omitempty"`
	Path       string `json:"path,omitempty"`
	NoExpire   bool   `json:"no_expire"`
	ExpireTime int    `json:"expire_time,omitempty"`
}

// GenerateSchemaResponse generate schema response
type GenerateSchemaResponse struct {
	base.CommonResponse
	Data GenerateSchemaData `json:"data"`
}

// GenerateSchemaData generate schema data
type GenerateSchemaData struct {
	Schema string `json:"schema"`
}

// QuerySchemaRequest query schema request
type QuerySchemaRequest struct {
	Schema string `json:"schema"`
	AppID  string `json:"app_id"`
}

// QuerySchemaResponse query schema response
type QuerySchemaResponse struct {
	base.CommonResponse
	Data QuerySchemaData `json:"data"`
}

// QuerySchemaData query schema data
type QuerySchemaData struct {
	AppID      string `json:"app_id"`
	Path       string `json:"path"`
	Query      string `json:"query"`
	CreateTime int    `json:"create_time"`
	ExpireTime int    `json:"expire_time"`
}

// QuerySchemaQuotaRequest query schema quota request
type QuerySchemaQuotaRequest struct {
	AppID string `json:"app_id"`
}

// QuerySchemaQuotaResponse query schema quota response
type QuerySchemaQuotaResponse struct {
	base.CommonResponse
	Data QuerySchemaQuotaData `json:"data"`
}

// TermSchemaQuota term schema quota
type TermSchemaQuota struct {
	SchemaLimit int `json:"schema_limit"`
	SchemaUsed  int `json:"schema_used"`
}

// QuerySchemaQuotaData query schema quota data
type QuerySchemaQuotaData struct {
	LongTermSchemaQuota  TermSchemaQuota `json:"long_term_schema_quota"`
	ShortTermSchemaQuota TermSchemaQuota `json:"short_term_schema_quota"`
}

// Schema create schema
type Schema struct {
	ctxCfg *credential.ContextConfig
}

// New create schema
func New(cfg *credential.ContextConfig) *Schema {
	return &Schema{
		ctxCfg: cfg,
	}
}

// Generate generate schema
func (s *Schema) Generate(ctx context.Context, request *GenerateSchemaRequest) (response *GenerateSchemaResponse, err error) {
	var resp []byte
	if resp, err = s.ctxCfg.Request().PostJSON(ctx, generateSchemaURL, request); err != nil {
		return
	}
	response = new(GenerateSchemaResponse)
	err = json.Unmarshal(resp, &response)
	return
}

// Query query schema
func (s *Schema) Query(ctx context.Context, request *QuerySchemaRequest) (response *QuerySchemaResponse, err error) {
	var resp []byte
	if resp, err = s.ctxCfg.Request().PostJSON(ctx, querySchemaURL, request); err != nil {
		return
	}
	response = new(QuerySchemaResponse)
	err = json.Unmarshal(resp, &response)
	return
}

// QueryQuota query schema quota
func (s *Schema) QueryQuota(ctx context.Context, request *QuerySchemaQuotaRequest) (response *QuerySchemaQuotaResponse, err error) {
	var resp []byte
	if resp, err = s.ctxCfg.Request().PostJSON(ctx, querySchemaQuota, request); err != nil {
		return
	}
	response = new(QuerySchemaQuotaResponse)
	err = json.Unmarshal(resp, &response)
	return
}
