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

package helper

import (
	"context"
	"testing"

	"github.com/houseme/bytedance/utility/logger"
)

func TestConcatenateSignSource(t *testing.T) {
	type args struct {
		ctx    context.Context
		data   interface{}
		salt   string
		logger logger.ILogger
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestConcatenateSignSource",
			args: args{
				ctx: context.Background(),
				data: struct {
					AppID       string `json:"app_id"`
					OutTradeNo  string `json:"out_trade_no,omitempty"`
					TotalAmount int    `json:"total_amount,omitempty"`
				}{
					AppID:       "appid12345",
					OutTradeNo:  "out_trade_no",
					TotalAmount: 100,
				},
				salt:   "test",
				logger: logger.NewDefaultLogger(),
			},
			want: "2001ff1c8ffa0f6b4bfb592133a84294",
		},
		{
			name: "TestConcatenateSignSource",
			args: args{
				ctx: context.Background(),
				// {"app_id":"tta2e4d4593dd752cc01","out_order_no":"509428213924556800","total_amount":5990,"subject":"充值：59.9 元","body":"充值：59.9 元","valid_time":7200,"sign":"f0ff2bd2339d4aa076e54a42cf6bf46e","cp_extra":"20231106180550.952","disable_msg":0}
				data: struct {
					AppID        string `json:"app_id" description:"小程序 APPID"`
					OutOrderNo   string `json:"out_order_no" description:"开发者侧的订单号。只能是数字、大小写字母_-*且在同一个 app_id 下唯一"`
					TotalAmount  int    `json:"total_amount" description:"订单总金额，单位为分"`
					Subject      string `json:"subject" description:"商品描述。长度限制不超过 128 字节且不超过 42 字符"`
					Body         string `json:"body" description:"商品详情。长度限制不超过 128 字节且不超过 42 字符"`
					ValidTime    int    `json:"valid_time" description:"订单过期时间 (秒)。最小 5 分钟，最大 2 天，小于 5 分钟会被置为 5 分钟，大于 2 天会被置为 2 天，取值范围：[300,172800]"`
					Sign         string `json:"sign" description:"签名，详见签名 DEMO"`
					CpExtra      string `json:"cp_extra,omitempty" description:"开发者自定义字段，回调原样回传，超过最大长度会被截断"`
					NotifyURL    string `json:"notify_url,omitempty" description:"商户自定义回调地址，必须以 HTTPS 开头，支持 443 端口。指定时，支付成功后抖音会请求该地址通知开发者"`
					ThirdPartyID string `json:"thirdparty_id,omitempty" description:"第三方平台服务商 id，非服务商模式留空"`
					StoreUID     string `json:"store_uid,omitempty" description:"门店 id，非门店模式留空"`
					DisableMsg   int    `json:"disable_msg" description:"是否屏蔽支付完成后推送用户抖音消息，1-屏蔽 0-非屏蔽，默认为 0。特别注意：若接入 POI, 请传 1。"`
					MsgPage      string `json:"msg_page,omitempty" description:"支付完成后推送给用户的抖音消息跳转页面，开发者需要传入在 app.json 中定义的链接，如果不传则跳转首页。"`
					LimitPayWay  string `json:"limit_pay_way,omitempty" description:"屏蔽指定支付方式，屏蔽多个支付方式，请使用逗号，分割，枚举值：屏蔽微信支付：LIMIT_WX，屏蔽支付宝支付：LIMIT_ALI，屏蔽抖音支付：LIMIT_DYZF"`
				}{
					AppID:       "tta2e4d4593dd752cc01",
					OutOrderNo:  "509428213924556800",
					TotalAmount: 5990,
					Subject:     "充值 59.9 元",
					Body:        "充值 59.9 元",
					ValidTime:   7200,
					CpExtra:     "20231106180550.952",
					DisableMsg:  0,
				},
				salt:   "n5NZ1sQb8euPHV7BsPxARQ",
				logger: logger.NewDefaultLogger(),
			},
			want: "ac65148dbe5fe6ca969a76aaa66f1c2e",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RequestSign(tt.args.ctx, tt.args.data, tt.args.salt); got != tt.want {
				t.Errorf("ConcatenateSignSource() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCallbackSign(t *testing.T) {
	type args struct {
		ctx   context.Context
		token string
		data  any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestCallbackSign-1",
			args: args{
				ctx:   context.Background(),
				token: "12345",
				data: struct {
					Timestamp int64  `json:"timestamp"`
					Nonce     string `json:"nonce"`
					Msg       string `json:"msg"`
				}{
					Timestamp: 123456689,
					Nonce:     "121212121",
					Msg:       `{\"appid\":\"tt07e3715e98c9aac0\",\"cp_orderno\":\"out_order_no_1\",\"cp_extra\":\"\",\"way\":\"2\",\"payment_order_no\":\"2021070722001450071438803941\",\"total_amount\":9980,\"status\":\"SUCCESS\",\"seller_uid\":\"69631798443938962290\",\"extra\":\"null\",\"item_id\":\"\"}`,
				},
			},
			want: "f01dea530d78c831ca23447bb445aefb44ea2941",
		},
		{
			name: "TestCallbackSign-2",
			args: args{
				ctx:   context.Background(),
				token: "XTaSqbuSx5Mxxxxx",
				data: struct {
					Timestamp int64  `json:"timestamp"`
					Nonce     string `json:"nonce"`
					Msg       string `json:"msg"`
				}{
					Timestamp: 1699323852,
					Nonce:     "5069",
					Msg:       "{\"appid\":\"tta2e4d4593dd752cc01\",\"cp_orderno\":\"509674359901237248\",\"cp_extra\":\"20231107102356.988\",\"way\":\"2\",\"channel_no\":\"2023110722001478031451493029\",\"channel_gateway_no\":\"\",\"payment_order_no\":\"DPS2311071023579038593541272896\",\"out_channel_order_no\":\"2023110722001478031451493029\",\"total_amount\":1,\"status\":\"SUCCESS\",\"seller_uid\":\"72886065940053875560\",\"extra\":\"\",\"item_id\":\"\",\"paid_at\":1699323852,\"message\":\"\",\"order_id\":\"N7298540284277819658\",\"ec_pay_trade_no\":\"DTPP2311071023579038592558232896\"}",
				},
			},
			want: "8af88d7063a638da5d20d0996e9efc6f899929a3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CallbackSign(tt.args.ctx, tt.args.token, tt.args.data); got != tt.want {
				t.Errorf("CallbackSign() = %v, want %v", got, tt.want)
			}
		})
	}
}
