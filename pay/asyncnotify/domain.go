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

package asyncnotify

// AsyncRequest 异步通知
type AsyncRequest struct {
    Content          string `json:"content" description:"回调内容，应答中的报文主体（response body）"`
    Version          string `json:"version" description:"固定值：3.0。回调版本，用于开发者识别回调参数的变更"`
    Msg              string `json:"msg" description:"订单相关信息的 JSON 字符串"`
    Type             string `json:"type" description:"回调类型（支付结果回调为 payment）：payment（支付成功/支付取消）"`
    ByteIdentifyName string `json:"Byte-Identifyname" description:"回调标识，用于开发者识别回调来源" in:"header"`
    ByteLogID        string `json:"Byte-LogId" description:"回调日志 ID" in:"header"`
    ByteNonceStr     string `json:"Byte-Nonce-Str" description:"随机字符串" in:"header"`
    ByteSignature    string `json:"Byte-Signature" description:"签名" in:"header"`
    ByteTimestamp    string `json:"Byte-Timestamp" description:"时间戳" in:"header"`
}

// AsyncResponse 异步响应
type AsyncResponse struct {
    ErrNo       int          `json:"err_no"`
    ErrTips     string       `json:"err_tips"`
    Type        string       `json:"type" description:"回调类型（支付结果回调为 payment）：payment（支付成功/支付取消）"`
    PaymentData *PaymentData `json:"paymentData"`
    SettleData  *SettleData  `json:"settleData"`
}

// PaymentData 异步通知
type PaymentData struct {
    AppID          string `json:"app_id" description:"小程序 app_id"`
    OutOrderNo     string `json:"out_order_no" description:"开发者系统生成的订单号，与抖音开平交易单号 order_id 唯一关联，长度 <= 64byte"`
    OrderID        string `json:"order_id" description:"抖音开平侧订单 id，长度 <= 64byte"`
    Status         string `json:"status" description:"支付结果状态，目前有两种状态：SUCCESS（支付成功），CANCEL（支付取消）"`
    TotalAmount    int    `json:"total_amount" description:"订单总金额，单位分，支付金额为 = total_amount - discount_amount"`
    DiscountAmount int    `json:"discount_amount" description:"订单优惠金额，单位分，接入营销时请关注这个字段"`
    PayChannel     int    `json:"pay_channel" description:"支付渠道枚举（支付成功时才有）：1：微信，2：支付宝 10：抖音支付"`
    ChannelPayID   string `json:"channel_pay_id" description:"渠道支付单号，如微信的支付单号，可能为空，可通过查询订单信息重新获取，长度 <= 64byte"`
    MerchantUID    string `json:"merchant_uid" description:"该笔交易卖家商户号，可能为空，可通过查询订单信息重新获取，实际存储的是 int64 类型的值"`
    Message        string `json:"message" description:"该笔交易取消原因，如：USER_CANCEL：用户取消，TIME_OUT：超时取消"`
    CpExtra        string `json:"cp_extra"`
    EventTime      int64  `json:"event_time" description:"用户支付成功/支付取消时间戳，单位为毫秒"`
}

// SettleData  结算异步信息
type SettleData struct {
    AppID        string `json:"app_id" description:"小程序 id "`
    Status       string `json:"status"`
    OrderID      string `json:"order_id"`
    CpExtra      string `json:"cp_extra"`
    Message      string `json:"message"`
    EventTime    int64  `json:"event_time"`
    SettleID     string `json:"settle_id"`
    OutSettleNo  string `json:"out_settle_no"`
    Rake         int    `json:"rake"`
    Commission   int    `json:"commission" desc:"交易参与 CPS 投放等任务产生的佣金，单位分 "`
    SettleDetail string `json:"settle_detail" desc:"分账细节 "`
    SettleAmount int    `json:"settle_amount"`
    ItemOrderID  string `json:"item_order_id"`
    IsAutoSettle bool   `json:"is_auto_settle" desc:"是否自动分账 "`
}
