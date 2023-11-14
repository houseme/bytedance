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
    "crypto"
    "crypto/rand"
    "crypto/rsa"
    "crypto/sha256"
    "crypto/x509"
    "encoding/base64"
    "encoding/json"
    "encoding/pem"
    "fmt"
    "strings"
    
    "github.com/houseme/bytedance/config"
    "github.com/houseme/bytedance/credential"
    "github.com/houseme/bytedance/utility/base"
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

// UploadImage 上传图片
func (d *Drama) UploadImage(ctx context.Context, req *UploadImageRequest) (resp *UploadImageResponse, err error) {
    if req == nil {
        return nil, base.ErrRequestIsEmpty
    }
    req.ResourceType = ResourceTypeImage
    
    var accessToken string
    if accessToken, err = d.getAccessToken(ctx); err != nil {
        return
    }
    
    ctx = context.WithValue(ctx, config.AccessTokenKey, accessToken)
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
    
    var accessToken string
    if accessToken, err = d.getAccessToken(ctx); err != nil {
        return
    }
    ctx = context.WithValue(ctx, config.AccessTokenKey, accessToken)
    
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
    
    var accessToken string
    if accessToken, err = d.getAccessToken(ctx); err != nil {
        return
    }
    ctx = context.WithValue(ctx, config.AccessTokenKey, accessToken)
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
    
    var accessToken string
    if accessToken, err = d.getAccessToken(ctx); err != nil {
        return
    }
    ctx = context.WithValue(ctx, config.AccessTokenKey, accessToken)
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
    
    var accessToken string
    if accessToken, err = d.getAccessToken(ctx); err != nil {
        return
    }
    ctx = context.WithValue(ctx, config.AccessTokenKey, accessToken)
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
    
    var accessToken string
    if accessToken, err = d.getAccessToken(ctx); err != nil {
        return
    }
    ctx = context.WithValue(ctx, config.AccessTokenKey, accessToken)
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
    
    var accessToken string
    if accessToken, err = d.getAccessToken(ctx); err != nil {
        return
    }
    ctx = context.WithValue(ctx, config.AccessTokenKey, accessToken)
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
    
    var accessToken string
    if accessToken, err = d.getAccessToken(ctx); err != nil {
        return
    }
    ctx = context.WithValue(ctx, config.AccessTokenKey, accessToken)
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
    if accessToken, err = d.getAccessToken(ctx); err != nil {
        return
    }
    ctx = context.WithValue(ctx, config.AccessTokenKey, accessToken)
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
    if accessToken, err = d.getAccessToken(ctx); err != nil {
        return
    }
    
    ctx = context.WithValue(ctx, config.AccessTokenKey, accessToken)
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
    var accessToken string
    if accessToken, err = d.getAccessToken(ctx); err != nil {
        return
    }
    
    ctx = context.WithValue(ctx, config.AccessTokenKey, accessToken)
    var response []byte
    if response, err = d.ctxCfg.Request().PostJSON(ctx, playInfo, *req); err != nil {
        return
    }
    
    resp = &PlayInfoResponse{}
    err = json.Unmarshal(response, &resp)
    
    return
}

// GenSign 生成签名
func GenSign(method, url, timestamp, nonce, body string, privateKey *rsa.PrivateKey) (string, error) {
    // method 内容必须大写，如 GET、POST，uri 不包含域名，必须以'/'开头
    targetStr := method + "\n" + url + "\n" + timestamp + "\n" + nonce + "\n" + body + "\n"
    h := sha256.New()
    h.Write([]byte(targetStr))
    digestBytes := h.Sum(nil)
    
    signBytes, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, digestBytes)
    if err != nil {
        return "", err
    }
    sign := base64.StdEncoding.EncodeToString(signBytes)
    
    return sign, nil
}

// CheckSign 验证签名
func CheckSign(timestamp, nonce, body, signature, pubKeyStr string) (bool, error) {
    pubKey, err := PemToRSAPublicKey(pubKeyStr)
    if err != nil {
        return false, err
    }
    
    hashed := sha256.Sum256([]byte(timestamp + "\n" + nonce + "\n" + body + "\n"))
    signBytes, err := base64.StdEncoding.DecodeString(signature)
    if err != nil {
        return false, err
    }
    err = rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, hashed[:], signBytes)
    return err == nil, nil
}

// PemToRSAPublicKey pem to rsa public key
func PemToRSAPublicKey(pemKeyStr string) (*rsa.PublicKey, error) {
    block, _ := pem.Decode([]byte(pemKeyStr))
    if block == nil || len(block.Bytes) == 0 {
        return nil, fmt.Errorf("empty block in pem string")
    }
    key, err := x509.ParsePKIXPublicKey(block.Bytes)
    if err != nil {
        return nil, err
    }
    switch key := key.(type) {
    case *rsa.PublicKey:
        return key, nil
    default:
        return nil, fmt.Errorf("not rsa public key")
    }
}
