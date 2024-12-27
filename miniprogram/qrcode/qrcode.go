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

package qrcode

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/houseme/bytedance/credential"
	"github.com/houseme/bytedance/utility/base"
)

// QRCode 小程序码
type QRCode struct {
	*credential.ContextConfig
}

const qrcodeURL = `https://developer.toutiao.com/api/apps/qrcode`

// NewQRCode 实例
func NewQRCode(ctxCfg *credential.ContextConfig) *QRCode {
	return &QRCode{
		ContextConfig: ctxCfg,
	}
}

// Color QRCode color
type Color struct {
	R string `json:"r"`
	G string `json:"g"`
	B string `json:"b"`
}

// QRCoder 小程序码参数
type QRCoder struct {
	// 服务端 API 调用标识，获取方法
	AccessToken string `json:"access_token"`
	// 是打开二维码的字节系 app 名称，默认为今日头条，取值如下表所示
	AppName string `json:"appname,omitempty"`
	// 小程序/小游戏启动参数，小程序则格式为 encode({path}?{query})，小游戏则格式为 JSON 字符串，默认为空
	Path string `json:"path,omitempty"`
	// 二维码宽度，单位 px，最小 280px，最大 1280px，默认为 430px
	Width string `json:"width,omitempty"`
	// 二维码线条颜色，默认为黑色
	LineColor Color `json:"line_color,omitempty"`
	// 二维码背景颜色，默认为白色
	Background Color `json:"background,omitempty"`
	// 是否展示小程序/小游戏 icon，默认不展示
	SetIcon bool `json:"set_icon"`
}

const (
	// TOUTIAO 今日头条
	TOUTIAO = "toutiao" // 今日头条
	// TOUTIAOLITE 今日头条极速版
	TOUTIAOLITE = "toutiao_lite" // 今日头条极速版
	// DOUYIN 抖音
	DOUYIN = "douyin" // 抖音
	// DOUYINLITE 抖音极速版
	DOUYINLITE = "douyin_lite" // 抖音极速版
	// PIPIXIA 皮皮虾
	PIPIXIA = "pipixia" // 皮皮虾
	// HUOSHAN 火山小视频
	HUOSHAN = "huoshan" // 火山小视频
	// XIGUA 西瓜视频
	XIGUA = "xigua" // 西瓜视频
)

// FetchCode 获取小程序码
func (qrCode *QRCode) FetchCode(ctx context.Context, data QRCoder) (response []byte, err error) {
	if data.AccessToken == "" {
		var clientToken *credential.ClientToken
		if clientToken, err = qrCode.GetClientToken(ctx); err != nil {
			return
		}

		if clientToken == nil {
			return nil, fmt.Errorf("clientToken is nil")
		}
		data.AccessToken = clientToken.AccessToken
	}
	var contentType string
	if response, err = qrCode.Request().PostJSON(ctx, qrcodeURL, data); err != nil {
		return response, err
	}

	if strings.HasPrefix(contentType, "application/json") {
		// 返回错误信息
		var result base.CommonError
		if err = json.Unmarshal(response, &result); err == nil && result.ErrCode != 0 {
			err = fmt.Errorf("fetchCode error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
			return nil, err
		}
	}
	// 返回文件
	if contentType == "image/jpeg" {
		return response, nil
	}
	return nil, fmt.Errorf("fetchCode error : unknown response content type - %v", contentType)
}
