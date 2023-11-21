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

package withdraw

// MerchantWithdrawRequest 商户提现
type MerchantWithdrawRequest struct {
    ThirdPartyID   string `json:"thirdparty_id,omitempty"` // 三方用户唯一标识
    AppID          string `json:"appId"`
    MerchantUID    string `json:"merchant_uid" description:"商户号"`
    ChannelType    string `json:"channel_type" description:"渠道类型" desc:"提现渠道枚举值:alipay: 支付宝，wx: 微信，hz: 抖音支付，yeepay: 易宝，yzt: 担保支付企业版聚合账户"`
    WithdrawAmount int    `json:"withdraw_amount" description:"提现金额；单位分"`
    OutOrderID     string `json:"out_order_id" description:"外部单号（开发者侧）；唯一标识一笔提现请求"`
    Callback       string `json:"callback" description:"提现结果回调地址，现结果通知接口（开发者自己的 HTTPS 服务）；如果不传默认用支付设置中的回调地址"`
    CpExtra        string `json:"cp_extra" description:"开发者自定义数据"`
    MerchantEntity int    `json:"merchant_entity" description:"抖音信息和光合信号主体标识：不传或传 0 或 1 查抖音信息主体账户余额，传 2 查光合信号主体账户余额"`
    Sign           string `json:"sign" description:"签名"`
}

// MerchantWithdrawResponse 商户提现
type MerchantWithdrawResponse struct {
    ErrNo          int    `json:"err_no"`
    ErrTips        string `json:"err_tips"`
    OrderID        string `json:"order_id" description:"平台侧的受理单号"`
    MerchantEntity int    `json:"merchant_entity" description:"抖音信息和光合信号主体标识：1 查抖音信息主体账户余额，2 查光合信号主体账户余额"`
}

// QueryWithdrawRequest 查询提现
type QueryWithdrawRequest struct {
    ThirdPartyID string `json:"thirdparty_id,omitempty"` // 三方用户唯一标识
    AppID        string `json:"appId"`
    MerchantUID  string `json:"merchant_uid" description:"商户号"`
    ChannelType  string `json:"channel_type" description:"渠道类型" desc:"提现渠道枚举值:alipay: 支付宝，wx: 微信，hz: 抖音支付，yeepay: 易宝，yzt: 担保支付企业版聚合账户"`
    OutOrderID   string `json:"out_order_id" description:"外部单号（开发者侧）；唯一标识一笔提现请求"`
    Sign         string `json:"sign" description:"签名"`
}

// QueryWithdrawResponse 查询提现
// 注：
// 退票：商户的提现申请请求通过渠道（微信/支付宝/抖音支付）提交给银行处理后，银行返回结果是处理成功，渠道返回给商户提现成功，
// 但间隔一段时间后，银行再次通知渠道处理失败并返还款项给渠道，渠道再将该笔失败款返还至商户在渠道的账户余额中
type QueryWithdrawResponse struct {
    ErrNo     int    `json:"err_no"`
    ErrTips   string `json:"err_tips"`
    Status    string `json:"status" description:"状态枚举值：成功:SUCCESS，失败：FAIL，处理中：PROCESSING，退票：REEXCHANGE"`
    StatusMsg string `json:"statusMsg" description:"状态描述"`
}
