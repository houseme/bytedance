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

package trade

// QueryOrderRequest 查询订单
type QueryOrderRequest struct {
    OrderID    string `json:"order_id,omitempty" description:"交易订单号，order_id 与 out_order_no 二选一"`
    OutOrderNo string `json:"out_order_no,omitempty" description:"开发者的单号，order_id 与 out_order_no 二选一"`
}

// QueryOrderResponse 查询订单
type QueryOrderResponse struct {
    Data   *QueryOrderData `json:"data"`
    ErrNo  int             `json:"err_no"`
    ErrMsg string          `json:"err_msg"`
    LogID  string          `json:"log_id"`
}

// ItemOrder 商品订单
type ItemOrder struct {
    ItemOrderID     string `json:"item_order_id"`
    SkuID           string `json:"sku_id"`
    ItemOrderAmount int    `json:"item_order_amount"`
}

// QueryOrderData 查询订单
type QueryOrderData struct {
    OrderID        string       `json:"order_id"`
    OutOrderNo     string       `json:"out_order_no"`
    AppID          string       `json:"app_id"`
    PayStatus      string       `json:"pay_status"`
    PayTime        int64        `json:"pay_time"`
    PayChannel     int          `json:"pay_channel"`
    ChannelPayID   string       `json:"channel_pay_id"`
    TradeTime      int64        `json:"trade_time"`
    TotalAmount    int          `json:"total_amount"`
    DiscountAmount int          `json:"discount_amount"`
    MerchantUID    string       `json:"merchant_uid"`
    CpExtra        string       `json:"cp_extra"`
    ItemOrderList  []*ItemOrder `json:"item_order_list"`
}

// CreateOrderRequest 创建订单
type CreateOrderRequest struct {
    SkuList          []*SkuItem `json:"skuList" description:"下单商品信息，注意：目前只支持传入一项"`
    OutOrderNo       string     `json:"outOrderNo" description:"开发者系统生成的订单号，与抖音开平交易单号 order_id 唯一关联，长度 <= 64byte"`
    TotalAmount      int        `json:"totalAmount" description:"订单总金额，单位分，支付金额为 = total_amount - discount_amount"`
    PayExpireSeconds int        `json:"payExpireSeconds,omitempty" description:"支付超时时间，单位秒，例如 300 表示 300 秒后过期；不传或传 0 会使用默认值 300，不能超过 48 小时。"`
    PayNotifyURL     string     `json:"payNotifyUrl,omitempty" description:"支付结果通知地址，长度 <= 512byte"`
    MerchantUID      string     `json:"merchantUid,omitempty" description:"该笔交易卖家商户号，开发者自定义收款商户号"`
    OrderEntrySchema []*Schema  `json:"orderEntrySchema" description:"订单详情页 schema，用于描述订单详情页的跳转协议"`
    LimitPayWayList  []int      `json:"limitPayWayList,omitempty" description:"限制支付渠道，不传或传空数组表示不限制，目前支持的支付渠道有：1：微信，2：支付宝 10：抖音支付"`
}

// CreateOrderResponse 创建订单
type CreateOrderResponse struct {
    ByteAuthorization string `json:"byteAuthorization"`
    Data              string `json:"data"`
}

// Schema Schema
type Schema struct {
    Path   string         `json:"path"`
    Params map[string]any `json:"params"`
}

// SkuItem sku item info
type SkuItem struct {
    SkuID       string   `json:"skuId" description:"外部商品 id，如：号卡商品 id、会员充值套餐 id、某类服务 id、付费工具 id 等"`
    Price       int      `json:"price" description:"商品价格，单位分"`
    Quantity    int      `json:"quantity" description:"商品数量，购买数量 0 < quantity <= 100 "`
    Title       string   `json:"title" description:"商品标题，如：商品名称、服务名称、付费工具名称等"`
    ImageList   []string `json:"imageList" description:"商品图片列表，如：商品主图、服务主图、付费工具主图等，商品图片链接，长度 <= 512 字节注意：目前只支持传入一项"`
    Type        int      `json:"type" description:"商品类型，目前支持的类型有：101：号卡商品 102：通信定制类商品 (彩铃)，103：话费/宽带充值类商品，201：通用咨询类商品，202:  代写文书，301：虚拟工具类商品 ,401：内容消费类商品"`
    TagGroupID  string   `json:"tagGroupId" description:"商品标签组 id，用于商品标签的分组，长度 <= 64 字节"`
    EntrySchema *Schema  `json:"entrySchema,omitempty" description:"商品详情页 schema，用于描述商品详情页的跳转协议"`
}
