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

// Package drama drama
package drama

import (
    "context"
    "encoding/json"
    "strings"
    
    "github.com/houseme/bytedance/config"
    "github.com/houseme/bytedance/credential"
    "github.com/houseme/bytedance/utility/base"
    "github.com/houseme/bytedance/utility/helper"
)

// Drama mini drama
type Drama struct {
    ctxCfg *credential.ContextConfig
}

// NewDrama init
func NewDrama(cfg *credential.ContextConfig) *Drama {
    return &Drama{ctxCfg: cfg}
}

// getAccessToken 获取 access_token
func (d *Drama) getAccessToken(ctx context.Context) (accessToken string, err error) {
    var clientToken *credential.ClientToken
    if clientToken, err = d.ctxCfg.GetClientToken(ctx); err != nil {
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
func (d *Drama) setContext(ctx context.Context) (context.Context, string, error) {
    accessToken, err := d.getAccessToken(ctx)
    if err != nil {
        return nil, "", err
    }
    ctx = context.WithValue(
        ctx,
        config.AccessTokenKey,
        accessToken,
    )
    return ctx, accessToken, nil
}

// UploadImage 上传图片
func (d *Drama) UploadImage(ctx context.Context, req *UploadImageRequest) (resp *UploadImageResponse, err error) {
    if req == nil {
        return nil, base.ErrRequestIsEmpty
    }
    req.ResourceType = ResourceTypeImage
    if req.MaAppID == "" {
        req.MaAppID = d.ctxCfg.Config.ClientKey()
    }
    var accessToken string
    if ctx, accessToken, err = d.setContext(ctx); err != nil {
        return nil, err
    }
    
    var response []byte
    if response, err = d.ctxCfg.Request().PostJSON(ctx, resourceUpload+accessToken, *req); err != nil {
        return
    }
    
    resp = &UploadImageResponse{}
    err = json.Unmarshal(response, &resp)
    
    return
}

// UploadVideo 上传视频
func (d *Drama) UploadVideo(ctx context.Context, req *UploadVideoRequest) (resp *UploadVideoResponse, err error) {
    if req == nil {
        return nil, base.ErrRequestIsEmpty
    }
    req.ResourceType = ResourceTypeVideo
    if req.MaAppID == "" {
        req.MaAppID = d.ctxCfg.Config.ClientKey()
    }
    var accessToken string
    if ctx, accessToken, err = d.setContext(ctx); err != nil {
        return nil, err
    }
    
    var response []byte
    if response, err = d.ctxCfg.Request().PostJSON(ctx, resourceUpload+accessToken, *req); err != nil {
        return
    }
    
    resp = &UploadVideoResponse{}
    err = json.Unmarshal(response, &resp)
    
    return
}

// QueryVideo 查询视频
func (d *Drama) QueryVideo(ctx context.Context, req *QueryVideoRequest) (resp *QueryVideoResponse, err error) {
    if req == nil {
        return nil, base.ErrRequestIsEmpty
    }
    if req.MaAppID == "" {
        req.MaAppID = d.ctxCfg.Config.ClientKey()
    }
    
    var accessToken string
    if ctx, accessToken, err = d.setContext(ctx); err != nil {
        return nil, err
    }
    var response []byte
    if response, err = d.ctxCfg.Request().PostJSON(ctx, queryVideo+accessToken, *req); err != nil {
        return
    }
    
    resp = &QueryVideoResponse{}
    err = json.Unmarshal(response, &resp)
    
    return
}

// CreateVideo 创建视频
func (d *Drama) CreateVideo(ctx context.Context, req *CreateVideoRequest) (resp *CreateVideoResponse, err error) {
    if req == nil {
        return nil, base.ErrRequestIsEmpty
    }
    
    if req.MaAppID == "" {
        req.MaAppID = d.ctxCfg.Config.ClientKey()
    }
    
    var accessToken string
    if ctx, accessToken, err = d.setContext(ctx); err != nil {
        return nil, err
    }
    var response []byte
    if response, err = d.ctxCfg.Request().PostJSON(ctx, createVideo+accessToken, *req); err != nil {
        return
    }
    
    resp = &CreateVideoResponse{}
    err = json.Unmarshal(response, &resp)
    
    return
}

// EditVideo 编辑视频
func (d *Drama) EditVideo(ctx context.Context, req *EditVideoRequest) (resp *EditVideoResponse, err error) {
    if req == nil {
        return nil, base.ErrRequestIsEmpty
    }
    
    if req.MaAppID == "" {
        req.MaAppID = d.ctxCfg.Config.ClientKey()
    }
    
    var accessToken string
    if ctx, accessToken, err = d.setContext(ctx); err != nil {
        return nil, err
    }
    var response []byte
    if response, err = d.ctxCfg.Request().PostJSON(ctx, editVideo+accessToken, *req); err != nil {
        return
    }
    
    resp = &EditVideoResponse{}
    err = json.Unmarshal(response, &resp)
    
    return
}

// QueryVideoAlbum 查询视频专辑
func (d *Drama) QueryVideoAlbum(ctx context.Context, req *QueryVideoAlbumRequest) (resp *QueryVideoAlbumResponse, err error) {
    if req == nil {
        return nil, base.ErrRequestIsEmpty
    }
    
    if req.MaAppID == "" {
        req.MaAppID = d.ctxCfg.Config.ClientKey()
    }
    
    var accessToken string
    if ctx, accessToken, err = d.setContext(ctx); err != nil {
        return nil, err
    }
    var response []byte
    if response, err = d.ctxCfg.Request().PostJSON(ctx, queryVideoAlbum+accessToken, *req); err != nil {
        return
    }
    
    resp = &QueryVideoAlbumResponse{}
    err = json.Unmarshal(response, &resp)
    
    return
}

// ReviewVideo 审核视频 短剧送审
func (d *Drama) ReviewVideo(ctx context.Context, req *ReviewVideoRequest) (resp *ReviewVideoResponse, err error) {
    if req == nil {
        return nil, base.ErrRequestIsEmpty
    }
    
    if req.MaAppID == "" {
        req.MaAppID = d.ctxCfg.Config.ClientKey()
    }
    
    var accessToken string
    if ctx, accessToken, err = d.setContext(ctx); err != nil {
        return nil, err
    }
    var response []byte
    if response, err = d.ctxCfg.Request().PostJSON(ctx, reviewVideo+accessToken, *req); err != nil {
        return
    }
    
    resp = &ReviewVideoResponse{}
    err = json.Unmarshal(response, &resp)
    
    return
}

// AuthorizeVideo 短剧授权
func (d *Drama) AuthorizeVideo(ctx context.Context, req *AuthorizeVideoRequest) (resp *AuthorizeVideoResponse, err error) {
    if req == nil {
        return nil, base.ErrRequestIsEmpty
    }
    
    if req.MaAppID == "" {
        req.MaAppID = d.ctxCfg.Config.ClientKey()
    }
    
    var accessToken string
    if ctx, accessToken, err = d.setContext(ctx); err != nil {
        return nil, err
    }
    var response []byte
    if response, err = d.ctxCfg.Request().PostJSON(ctx, authorizeVideo+accessToken, *req); err != nil {
        return
    }
    
    resp = &AuthorizeVideoResponse{}
    err = json.Unmarshal(response, &resp)
    
    return
}

// OnlineAlbum 上线视频专辑
func (d *Drama) OnlineAlbum(ctx context.Context, req *OnlineAlbumRequest) (resp *OnlineAlbumResponse, err error) {
    if req == nil {
        return nil, base.ErrRequestIsEmpty
    }
    
    var accessToken string
    if ctx, accessToken, err = d.setContext(ctx); err != nil {
        return nil, err
    }
    var response []byte
    if response, err = d.ctxCfg.Request().PostJSON(ctx, onlineAlbum+accessToken, *req); err != nil {
        return
    }
    
    resp = &OnlineAlbumResponse{}
    err = json.Unmarshal(response, &resp)
    
    return
}

// BindAlbum 绑定视频专辑
func (d *Drama) BindAlbum(ctx context.Context, req *BindAlbumRequest) (resp *BindAlbumResponse, err error) {
    if req == nil {
        return nil, base.ErrRequestIsEmpty
    }
    
    var accessToken string
    if ctx, accessToken, err = d.setContext(ctx); err != nil {
        return nil, err
    }
    var response []byte
    if response, err = d.ctxCfg.Request().PostJSON(ctx, bindAlbum+accessToken, *req); err != nil {
        return
    }
    
    resp = &BindAlbumResponse{}
    err = json.Unmarshal(response, &resp)
    
    return
}

// PlayInfo 获取视频播放信息
func (d *Drama) PlayInfo(ctx context.Context, req *PlayInfoRequest) (resp *PlayInfoResponse, err error) {
    if req == nil {
        return nil, base.ErrRequestIsEmpty
    }
    
    if req.MaAppID == "" {
        req.MaAppID = d.ctxCfg.Config.ClientKey()
    }
    
    var accessToken string
    if ctx, accessToken, err = d.setContext(ctx); err != nil {
        return nil, err
    }
    var response []byte
    if response, err = d.ctxCfg.Request().PostJSON(ctx, playInfo+accessToken, *req); err != nil {
        return
    }
    
    resp = &PlayInfoResponse{}
    err = json.Unmarshal(response, &resp)
    
    return
}

// AsyncNotify 异步通知
func (d *Drama) AsyncNotify(ctx context.Context, req *AsyncRequest) (resp *AsyncResponse, err error) {
    if req == nil {
        return nil, base.ErrRequestIsEmpty
    }
    resp = &AsyncResponse{
        ErrNo:   ErrNoSuccess,
        ErrTips: ErrTipsSuccess,
    }
    if req.Version != DefaultAsyncVersion {
        resp.ErrNo = ErrNoVersion
        resp.ErrTips = ErrTipsVersion
        return
    }
    
    var checkResult bool
    if checkResult, err = helper.CheckSign(req.ByteTimestamp, req.ByteNonceStr, req.Content, req.ByteSignature, d.ctxCfg.Config.PublicKey()); err != nil {
        return
    }
    if !checkResult {
        resp.ErrNo = ErrNoFailedToCheckTheSignature
        resp.ErrTips = ErrTipsFailedToCheckTheSignature
        return
    }
    
    if req.Type == AlbumAudit {
        var data = new(AsyncAlbumAudit)
        if err = json.Unmarshal([]byte(req.Msg), data); err != nil {
            resp.ErrNo = ErrNoSystemError
            resp.ErrTips = ErrTipsSystemError + err.Error()
            return
        }
        resp.AlbumAudit = data
    }
    
    if req.Type == EpisodeAudit {
        var data = new(AsyncEpisodeAudit)
        if err = json.Unmarshal([]byte(req.Msg), data); err != nil {
            resp.ErrNo = ErrNoSystemError
            resp.ErrTips = ErrTipsSystemError + err.Error()
            return
        }
        resp.EpisodeAudit = data
    }
    
    if req.Type == UploadVideo {
        var data = new(AsyncUploadVideo)
        if err = json.Unmarshal([]byte(req.Msg), data); err != nil {
            resp.ErrNo = ErrNoSystemError
            resp.ErrTips = ErrTipsSystemError + err.Error()
            return
        }
        resp.UploadVideo = data
    }
    
    return
}
