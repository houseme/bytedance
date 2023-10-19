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

// Package link generate short links
package link

import (
	"github.com/houseme/bytedance/credential"
	"github.com/houseme/bytedance/utility/base"
)

const (
	generateLink     = "https://developer.toutiao.com/api/apps/url_link/generate"
	generateLinkV2   = "https://open.douyin.com/api/apps/v1/url_link/generate"
	queryLinkV2      = "https://open.douyin.com/api/apps/v1/url_link/query_info"
	queryLinkQuotaV2 = "https://open.douyin.com/api/apps/v1/url_link/query_quota"
)

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
	*credential.ContextConfig
}

// New create short link
func New(cfg *credential.ContextConfig) *Link {
	return &Link{
		ContextConfig: cfg,
	}
}
