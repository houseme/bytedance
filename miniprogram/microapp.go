/*
 *  Copyright bytedance Author(https://houseme.github.io/bytedance/). All Rights Reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 *  You can obtain one at https://github.com/houseme/bytedance.
 *
 */

// Package miniprogram mini program
package miniprogram

import (
    "context"
    
    "github.com/houseme/bytedance/config"
    "github.com/houseme/bytedance/credential"
    "github.com/houseme/bytedance/miniprogram/authorize"
    "github.com/houseme/bytedance/miniprogram/link"
    "github.com/houseme/bytedance/miniprogram/payment"
    "github.com/houseme/bytedance/miniprogram/qrcode"
    "github.com/houseme/bytedance/miniprogram/schema"
    "github.com/houseme/bytedance/utility/base"
)

// MicroApp mini program
type MicroApp struct {
    ctxCfg *credential.ContextConfig
}

// New micro app
func New(ctx context.Context, cfg *config.Config) (*MicroApp, error) {
    if cfg == nil {
        return nil, base.ErrConfigNotFound
    }
    if cfg.ClientKey() == "" || cfg.ClientSecret() == "" {
        return nil, base.ErrConfigKeyValueEmpty("clientKey or clientSecret")
    }
    
    return &MicroApp{
        ctxCfg: &credential.ContextConfig{
            Config:            cfg,
            AccessTokenHandle: credential.NewDefaultAccessToken(ctx, cfg),
        },
    }, nil
}

// SetAccessTokenHandle 自定义 access_token 获取方式
func (ma *MicroApp) SetAccessTokenHandle(accessTokenHandle credential.AccessTokenHandle) {
    ma.ctxCfg.AccessTokenHandle = accessTokenHandle
}

// GetContext get Context
func (ma *MicroApp) GetContext() *credential.ContextConfig {
    return ma.ctxCfg
}

// GetAccessToken 获取 access_token
func (ma *MicroApp) GetAccessToken(ctx context.Context, openID string) (string, error) {
    return ma.ctxCfg.GetAccessToken(ctx, openID)
}

// GetClientToken 获取 client_token
func (ma *MicroApp) GetClientToken(ctx context.Context) (string, error) {
    clientToken, err := ma.ctxCfg.GetClientToken(ctx)
    if err != nil {
        return "", err
    }
    return clientToken.AccessToken, nil
}

// GetAuthorize oauth2 网页授权
func (ma *MicroApp) GetAuthorize() *authorize.Authorize {
    return authorize.NewAuthorize(ma.ctxCfg)
}

// GetQrcode 获取小程序码
func (ma *MicroApp) GetQrcode() *qrcode.QRCode {
    return qrcode.NewQRCode(ma.ctxCfg)
}

// GetLink 获取小程序 link
func (ma *MicroApp) GetLink() *link.Link {
    return link.New(ma.ctxCfg)
}

// GetSchema 获取小程序 schema
func (ma *MicroApp) GetSchema() *schema.Schema {
    return schema.New(ma.ctxCfg)
}

// GetPay 获取支付
func (ma *MicroApp) GetPay() (*payment.Pay, error) {
    return payment.NewPay(ma.ctxCfg.Config)
}
