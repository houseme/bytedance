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

// Package voc voc
package voc

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/houseme/bytedance/credential"
	"github.com/houseme/bytedance/utility/base"
)

// Voc mini drama
type Voc struct {
	ctxCfg *credential.ContextConfig
}

// NewVoc init
func NewVoc(cfg *credential.ContextConfig) *Voc {
	return &Voc{ctxCfg: cfg}
}

// getAccessToken 获取 access_token
func (v *Voc) getAccessToken(ctx context.Context) (accessToken string, err error) {
	var clientToken *credential.ClientToken
	if clientToken, err = v.ctxCfg.GetClientToken(ctx); err != nil {
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

// QueryVideoList 查询视频列表
func (v *Voc) QueryVideoList(ctx context.Context, req *QueryListRequest) (resp *QueryListResponse, err error) {
	if req == nil {
		return nil, base.ErrRequestIsEmpty
	}

	var accessToken string
	if accessToken, err = v.getAccessToken(ctx); err != nil {
		return nil, err
	}

	var response []byte
	if response, err = v.ctxCfg.Request().PostJSON(ctx, getVideoList+accessToken, *req); err != nil {
		return nil, err
	}

	resp = &QueryListResponse{}
	if err = json.Unmarshal(response, &resp); err != nil {
		return nil, err
	}

	if resp != nil && resp.ErrNo != 0 {
		err = fmt.Errorf("query video errCode: %d, error:%s", resp.ErrNo, resp.ErrMsg)
		return nil, err
	}

	return
}

// DeleteVideo 删除视频
func (v *Voc) DeleteVideo(ctx context.Context, req *DeleteVideoRequest) (resp *DeleteVideoResponse, err error) {
	if req == nil {
		return nil, base.ErrRequestIsEmpty
	}

	var accessToken string
	if accessToken, err = v.getAccessToken(ctx); err != nil {
		return nil, err
	}

	var response []byte
	if response, err = v.ctxCfg.Request().PostJSON(ctx, deleteVideo+accessToken, *req); err != nil {
		return nil, err
	}

	resp = &DeleteVideoResponse{}
	err = json.Unmarshal(response, &resp)

	return
}

// QueryVideoURL 获取视频播放地址
func (v *Voc) QueryVideoURL(ctx context.Context, req *QueryVideoURLRequest) (resp *QueryVideoURLResponse, err error) {
	if req == nil {
		return nil, base.ErrRequestIsEmpty
	}

	var accessToken string
	if accessToken, err = v.getAccessToken(ctx); err != nil {
		return nil, err
	}

	var response []byte
	if response, err = v.ctxCfg.Request().PostJSON(ctx, getVideoByVID+accessToken, *req); err != nil {
		return nil, err
	}

	resp = &QueryVideoURLResponse{}
	err = json.Unmarshal(response, &resp)

	return
}

// BatchUploadVideoByURL 批量上传视频
func (v *Voc) BatchUploadVideoByURL(ctx context.Context, req *UploadByURLRequest) (resp *UploadByURLResponse, err error) {
	if req == nil {
		return nil, base.ErrRequestIsEmpty
	}

	var accessToken string
	if accessToken, err = v.getAccessToken(ctx); err != nil {
		return nil, err
	}

	var response []byte
	if response, err = v.ctxCfg.Request().PostJSON(ctx, uploadVideoByURLs+accessToken, *req); err != nil {
		return nil, err
	}

	resp = &UploadByURLResponse{}
	err = json.Unmarshal(response, &resp)

	return
}

// QueryUploadVideoJobInfo 查询视频状态
func (v *Voc) QueryUploadVideoJobInfo(ctx context.Context, req *QueryUploadByURLRequest) (resp *QueryUploadByURLResponse, err error) {
	if req == nil {
		return nil, base.ErrRequestIsEmpty
	}

	var accessToken string
	if accessToken, err = v.getAccessToken(ctx); err != nil {
		return nil, err
	}

	if strings.TrimSpace(req.AccessToken) == "" {
		req.AccessToken = accessToken
	}

	var response []byte
	if response, err = v.ctxCfg.Request().PostJSON(ctx, getVideoUploadJobInfo+accessToken, *req); err != nil {
		return nil, err
	}

	resp = &QueryUploadByURLResponse{}
	err = json.Unmarshal(response, &resp)

	return
}

// StartWorkFlow 发起转码处理
func (v *Voc) StartWorkFlow(ctx context.Context, req *StartWorkFlowRequest) (resp *StartWorkFlowResponse, err error) {
	if req == nil {
		return nil, base.ErrRequestIsEmpty
	}

	var accessToken string
	if accessToken, err = v.getAccessToken(ctx); err != nil {
		return nil, err
	}

	var response []byte
	if response, err = v.ctxCfg.Request().PostJSON(ctx, startWorkFlow+accessToken, *req); err != nil {
		return nil, err
	}

	resp = &StartWorkFlowResponse{}
	err = json.Unmarshal(response, &resp)

	return
}

// QueryWorkFlow 查询转码状态
func (v *Voc) QueryWorkFlow(ctx context.Context, req *QueryWorkFlowRequest) (resp *QueryWorkFlowResponse, err error) {
	if req == nil {
		return nil, base.ErrRequestIsEmpty
	}

	var accessToken string
	if accessToken, err = v.getAccessToken(ctx); err != nil {
		return nil, err
	}

	var response []byte
	if response, err = v.ctxCfg.Request().PostJSON(ctx, getWorkFlowStatus+accessToken, *req); err != nil {
		return nil, err
	}

	resp = &QueryWorkFlowResponse{}
	err = json.Unmarshal(response, &resp)

	return
}
