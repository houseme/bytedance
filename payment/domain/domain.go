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

// Package domain domain
package domain

// AsyncResponse async response
type AsyncResponse struct {
    ErrNo   int    `json:"err_no" description:"返回码，0 代表成功 非 0 代表失败"`
    ErrTips string `json:"err_tips" description:"返回码信息"`
}

// AsyncRequest async request
type AsyncRequest struct {
    Timestamp    string `json:"timestamp"`
    Nonce        string `json:"nonce"`
    Msg          string `json:"msg"`
    MsgSignature string `json:"msg_signature"`
    Type         string `json:"type"`
}

// AsyncPaymentData async payment data
type AsyncPaymentData struct {
    AppID          string `json:"appid"`
    CpOrderNo      string `json:"cp_orderno" description:"开发者侧的订单号。只能是数字、大小写字母_-*且在同一个 app_id 下唯一"`
    CpExtra        string `json:"cp_extra" description:"预下单时开发者传入字段"`
    Way            string `json:"way" description:"支付渠道，1-微信支付，2-支付宝支付，10-抖音支付"`
    ChannelNo      string `json:"channel_no" description:"支付渠道侧单号 (抖音平台请求下游渠道微信或支付宝时传入的单号)"`
    PaymentOrderNo string `json:"payment_order_no" description:"支付渠道侧 PC 单号，支付页面可见 (微信支付宝侧的订单号)"`
    TotalAmount    int    `json:"total_amount" description:"订单总金额，单位为分"`
    Status         string `json:"status" description:"固定 SUCCESS"`
    SellerUID      string `json:"seller_uid" description:"该笔交易卖家商户号"`
    Extra          string `json:"extra"`
    ItemID         string `json:"item_id" description:"订单来源视频对应视频 ID"`
    OrderID        string `json:"order_id" description:"抖音侧订单号"`
    PaidAt         int    `json:"paid_at" description:"支付时间，Unix 时间戳，10 位，整型数"`
}

// AsyncSettleData async settle data
type AsyncSettleData struct {
    AppID           string `json:"app_id"`
    CpSettleNo      string `json:"cp_settle_no"`
    CpExtra         string `json:"cp_extra"`
    Status          string `json:"status"`
    Rake            int    `json:"rake"`
    Commission      int    `json:"commission"`
    SettleDetail    string `json:"settle_detail"`
    SettledAt       int    `json:"settled_at"`
    Message         string `json:"message"`
    OrderID         string `json:"order_id" description:"分账对应原支付单单号"`
    ChannelSettleID string `json:"channel_settle_id" description:"渠道结算单号"`
    SettleAmount    int    `json:"settle_amount" description:"参与分账总金额，单位为分"`
    SettleNo        string `json:"settle_no" description:"分账对应平台单号"`
    OutOrderNo      string `json:"out_order_no" description:"开发者侧的订单号。只能是数字、大小写字母_-*且在同一个 app_id 下唯一"`
    IsAutoSettle    bool   `json:"is_auto_settle" description:"是否自动结算"`
}

// AsyncRefundData async refund data
type AsyncRefundData struct {
    AppID        string `json:"appid"`
    CpRefundNo   string `json:"cp_refundno"`
    CpExtra      string `json:"cp_extra"`
    Status       string `json:"status"`
    RefundAmount int    `json:"refund_amount"`
    IsAllSettled bool   `json:"is_all_settled"`
    RefundedAt   int    `json:"refunded_at"`
    Message      string `json:"message"`
    OrderID      string `json:"order_id"`
    RefundNo     string `json:"refund_no"`
}

// AsyncWithdrawData async withdraw data
type AsyncWithdrawData struct {
    Status     string `json:"status"`
    Extra      string `json:"extra"`
    Message    string `json:"message"`
    WithdrawAt int    `json:"withdraw_at"`
    OrderID    string `json:"order_id"`
    OutOrderID string `json:"out_order_id"`
    ChOrderID  string `json:"ch_order_id"`
}

// CreateOrderRequest create order request
// see https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/ecpay/pay-list/pay
type CreateOrderRequest struct {
    AppID           string           `json:"app_id" description:"小程序 APPID"`
    OutOrderNo      string           `json:"out_order_no" description:"开发者侧的订单号。只能是数字、大小写字母_-*且在同一个 app_id 下唯一"`
    TotalAmount     int              `json:"total_amount" description:"订单总金额，单位为分"`
    Subject         string           `json:"subject" description:"商品描述。长度限制不超过 128 字节且不超过 42 字符"`
    Body            string           `json:"body" description:"商品详情。长度限制不超过 128 字节且不超过 42 字符"`
    ValidTime       int              `json:"valid_time" description:"订单过期时间 (秒)。最小 5 分钟，最大 2 天，小于 5 分钟会被置为 5 分钟，大于 2 天会被置为 2 天，取值范围：[300,172800]"`
    Sign            string           `json:"sign" description:"签名，详见签名 DEMO"`
    CpExtra         string           `json:"cp_extra,omitempty" description:"开发者自定义字段，回调原样回传，超过最大长度会被截断"`
    NotifyURL       string           `json:"notify_url,omitempty" description:"商户自定义回调地址，必须以 HTTPS 开头，支持 443 端口。指定时，支付成功后抖音会请求该地址通知开发者"`
    ThirdPartyID    string           `json:"thirdparty_id,omitempty" description:"第三方平台服务商 id，非服务商模式留空"`
    StoreUID        string           `json:"store_uid,omitempty" description:"门店 id，非门店模式留空"`
    DisableMsg      int              `json:"disable_msg" description:"是否屏蔽支付完成后推送用户抖音消息，1-屏蔽 0-非屏蔽，默认为 0。特别注意：若接入 POI, 请传 1。"`
    MsgPage         string           `json:"msg_page,omitempty" description:"支付完成后推送给用户的抖音消息跳转页面，开发者需要传入在 app.json 中定义的链接，如果不传则跳转首页。"`
    ExpandOrderInfo *ExpandOrderInfo `json:"expand_order_info,omitempty" description:"订单拓展信息，详见下面 expand_order_info 参数说明"`
    LimitPayWay     string           `json:"limit_pay_way,omitempty" description:"屏蔽指定支付方式，屏蔽多个支付方式，请使用逗号，分割，枚举值：屏蔽微信支付：LIMIT_WX，屏蔽支付宝支付：LIMIT_ALI，屏蔽抖音支付：LIMIT_DYZF"`
}

// ExpandOrderInfo expand order info
type ExpandOrderInfo struct {
    OriginalDeliveryFee int `json:"original_delivery_fee" description:"原始配送费，单位为分"`
    ActualDeliveryFee   int `json:"actual_delivery_fee" description:"实际配送费，单位为分"`
}

// CreateOrderResponse create order response
type CreateOrderResponse struct {
    ErrNo   int              `json:"err_no" description:"返回码，0 代表成功 非 0 代表失败"`
    ErrTips string           `json:"err_tips" description:"返回码信息"`
    Data    *CreateOrderData `json:"data,omitempty"`
}

// CreateOrderData create order data
type CreateOrderData struct {
    OrderId    string `json:"order_id"`
    OrderToken string `json:"order_token"`
}

// QueryOrderRequest query order request
// see https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/ecpay/pay-list/query
type QueryOrderRequest struct {
    AppID        string `json:"app_id"`
    OutOrderNo   string `json:"out_order_no" description:"开发者侧的订单号。只能是数字、大小写字母_-*且在同一个 app_id 下唯一"`
    Sign         string `json:"sign"`
    ThirdPartyID string `json:"thirdparty_id,omitempty" description:"第三方平台服务商 id，非服务商模式留空"`
}

// QueryOrderResponse query order response
type QueryOrderResponse struct {
    ErrNo       int          `json:"err_no"`
    ErrTips     string       `json:"err_tips"`
    OutOrderNo  string       `json:"out_order_no"`
    OrderId     string       `json:"order_id"`
    PaymentInfo *PaymentInfo `json:"payment_info"`
    CpsInfo     string       `json:"cps_info"`
}

// PaymentInfo payment info
type PaymentInfo struct {
    TotalFee    int    `json:"total_fee" description:"订单总金额，单位为分"`
    OrderStatus string `json:"order_status" description:"支付状态枚举值：SUCCESS：成功 TIMEOUT：超时未支付 PROCESSING：处理中 FAIL：失败"`
    PayTime     string `json:"pay_time" description:"支付完成时间，order_status 不为 SUCCESS 时会返回默认值空字符串，order_status 为 SUCCESS 时返回非空字符串，格式为'yyyy-MM-dd HH:mm:ss'"`
    Way         int    `json:"way" description:"支付渠道，order_status 不为 SUCCESS 时会返回默认值 0，order_status 为 SUCCESS 时会返回以下枚举：1-微信支付，2-支付宝支付，10-抖音支付"`
    ChannelNo   string `json:"channel_no" description:"支付渠道侧的支付单号"`
    SellerUid   string `json:"seller_uid" description:"该笔交易卖家商户号"`
    ItemId      string `json:"item_id" description:"订单来源视频对应视频 ID"`
    CpExtra     string `json:"cp_extra" description:"开发者自定义字段，回调原样回传，超过最大长度会被截断"`
}

// CpsInfo cps info
type CpsInfo struct {
    ShareAmount string `json:"share_amount" description:"达人分佣金额，单位为分。"`
    DouYinID    string `json:"douyin_id" description:"达人抖音号"`
    Nickname    string `json:"nickname" description:"达人昵称"`
}
