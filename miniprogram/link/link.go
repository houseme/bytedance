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

// Package link generate short links
package link

import (
    "context"
    "encoding/json"
    
    "github.com/houseme/bytedance/credential"
    "github.com/houseme/bytedance/utility/base"
)

const (
    generateLink     = "https://developer.toutiao.com/api/apps/url_link/generate"
    generateLinkV2   = "https://open.douyin.com/api/apps/v1/url_link/generate"
    queryLinkV2      = "https://open.douyin.com/api/apps/v1/url_link/query_info"
    queryLinkQuotaV2 = "https://open.douyin.com/api/apps/v1/url_link/query_quota"
)

// GenerateV1Request generate link request
type GenerateV1Request struct {
    AccessToken string `json:"access_token"`
    MaAppID     string `json:"ma_app_id"`
    AppName     string `json:"app_name"`
    Path        string `json:"path,omitempty"`
    Query       string `json:"query,omitempty"`
    ExpireTime  int    `json:"expire_time"`
}

// GenerateV1Response generate link response
type GenerateV1Response struct {
    ErrNo   int    `json:"err_no"`
    ErrTips string `json:"err_tips"`
    URLLink string `json:"url_link"`
}

// GenerateLinkRequest generate link request
type GenerateLinkRequest struct {
    AppID      string `json:"app_id"`
    AppName    string `json:"app_name"`
    Path       string `json:"path,omitempty"`
    Query      string `json:"query,omitempty"`
    ExpireTime int    `json:"expire_time"`
}

// GenerateLinkResponse generate link response
type GenerateLinkResponse struct {
    base.CommonResponse
    Data GenerateLinkData `json:"data"`
}

// GenerateLinkData generate link data
type GenerateLinkData struct {
    URLLink string `json:"url_link"`
}

// QueryLinkQuotaRequest query link quota request
type QueryLinkQuotaRequest struct {
    AppID string `json:"app_id"`
}

// QueryLinkQuotaResponse query link quota response
type QueryLinkQuotaResponse struct {
    base.CommonResponse
    URLLinkQuota URLLinkQuota `json:"url_link_quota"`
}

// URLLinkQuota url link quota
type URLLinkQuota struct {
    URLLinkUsed  int `json:"url_link_used"`
    URLLinkLimit int `json:"url_link_limit"`
}

// QueryLinkRequest query link request
type QueryLinkRequest struct {
    AppID   string `json:"app_id"`
    URLLink string `json:"url_link"`
}

// QueryLinkResponse query link response
type QueryLinkResponse struct {
    base.CommonResponse
    Data QueryLinkData `json:"data"`
}

// QueryLinkData query link data
type QueryLinkData struct {
    AppName    string `json:"app_name"`
    AppID      string `json:"app_id"`
    Path       string `json:"path"`
    Query      string `json:"query"`
    CreateTime int    `json:"create_time"`
    ExpireTime int    `json:"expire_time"`
}

// Link short link relation
type Link struct {
    ctxCfg *credential.ContextConfig
}

// New create short link
func New(cfg *credential.ContextConfig) *Link {
    return &Link{
        ctxCfg: cfg,
    }
}

// Generate generate short link
func (l *Link) Generate(ctx context.Context, req *GenerateV1Request) (resp *GenerateV1Response, err error) {
    response, err := l.ctxCfg.Request().PostJSON(ctx, generateLink, req)
    if err != nil {
        return
    }
    resp = new(GenerateV1Response)
    err = json.Unmarshal(response, &resp)
    return
}

// GenerateV2 generate short link v2
func (l *Link) GenerateV2(ctx context.Context, req *GenerateLinkRequest) (resp *GenerateLinkResponse, err error) {
    response, err := l.ctxCfg.Request().PostJSON(ctx, generateLinkV2, req)
    if err != nil {
        return
    }
    resp = new(GenerateLinkResponse)
    err = json.Unmarshal(response, &resp)
    return
}

// QueryQuotaV2 query link quota v2
func (l *Link) QueryQuotaV2(ctx context.Context, req *QueryLinkQuotaRequest) (resp *QueryLinkQuotaResponse, err error) {
    response, err := l.ctxCfg.Request().PostJSON(ctx, queryLinkQuotaV2, req)
    if err != nil {
        return
    }
    resp = new(QueryLinkQuotaResponse)
    err = json.Unmarshal(response, &resp)
    return
}

// QueryV2 query link v2
func (l *Link) QueryV2(ctx context.Context, req *QueryLinkRequest) (resp *QueryLinkResponse, err error) {
    response, err := l.ctxCfg.Request().PostJSON(ctx, queryLinkV2, req)
    if err != nil {
        return
    }
    resp = new(QueryLinkResponse)
    err = json.Unmarshal(response, &resp)
    return
}
