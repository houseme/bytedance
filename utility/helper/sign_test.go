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
            name: "TestCallbackSign",
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
            want: "94af40199bb84a6b9340ecd66a4d120f4371cabd",
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
