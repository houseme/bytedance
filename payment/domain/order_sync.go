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

package domain

// OrderSyncRequest order sync request
type OrderSyncRequest struct {
    ClientKey   string `json:"client_key,omitempty" description:"第三方在抖音开放平台申请的 ClientKey 注意：POI 订单必传"`
    AccessToken string `json:"access_token" description:"服务端 API 调用标识，通过 getAccessToken 获取"`
    ExtShopID   string `json:"ext_shop_id,omitempty"`
    AppName     string `json:"app_name" description:"做订单展示的字节系 app 名称，目前为固定值“douyin”"`
    OpenID      string `json:"openId" description:"小程序用户的 open_id，通过 code2Session 获取"`
    OrderDetail any    `json:"order_detail" description:"json string，根据不同订单类型有不同的结构体，请参见 order_detail 字段说明（json string）"`
    OrderStatus int64  `json:"order_status" description:"普通小程序订单订单状态，POI 订单可以忽略 0：待支付 1：已支付 2：已取消（用户主动取消或者超时未支付导致的关单）4：已核销（核销状态是整单核销，即一笔订单买了 3 个券，核销是指 3 个券核销的整单）5：退款中 6：已退款 8：退款失败 注意：普通小程序订单必传，担保支付分账依赖该状态"`
    OrderType   int64  `json:"order_type" description:"订单类型，枚举值:0：普通小程序订单（非 POI 订单）9101：团购券订单（POI 订单）9001：景区门票订单（POI 订单）"`
    UpdateTime  int64  `json:"update_time" description:"订单更新时间，时间戳，单位秒"`
    Extra       string `json:"extra,omitempty" description:"自定义字段，用于关联具体业务场景下的特殊参数，长度 < 2048bytes"`
}

// OrderDetailGroup order detail
type OrderDetailGroup struct {
    ExtOrderID  string `json:"ext_order_id" description:"开发者系统侧业务单号。用作幂等控制。该订单号是和担保支付的支付单号绑定的，即预下单时传入的 out_order_no 字段，长度 <= 64byte"`
    Status      int64  `json:"status" description:"订单状态，枚举值：110：已取消（抖音订单中心可看到，状态为已取消）110：待支付 310：未使用 340：已使用 410：退款中 420：退款成功 430：退款失败"`
    ShopName    string `json:"shop_name" description:"门店名称，长度 <= 256byte"`
    EntryType   int64  `json:"entry_type" description:"入口类型，枚举值：1：H5 2：抖音小程序"`
    EntrySchema string `json:"entry_schema" description:"订单详情页的外链跳转 schema 参数，格式为 JSON 字符串。长度 <= 512byte"`
}

// OrderDetailGeneral order detail
type OrderDetailGeneral struct {
    OrderID    string                    `json:"order_id" description:"开发者侧业务单号。用作幂等控制。该订单号是和担保支付的支付单号绑定的，也就是预下单时传入的 out_order_no 字段，长度 <= 64byte"`
    CreateTime int64                     `json:"create_time" description:"订单创建时间，时间戳，13 位毫秒时间戳"`
    Status     string                    `json:"status" description:"订单状态，订单状态，建议采用以下枚举值：待待支付，已支付，已取消，已超时，已核销，退款中。已退款，退款失败"`
    Amount     int64                     `json:"amount" description:"订单商品总数"`
    TotalPrice int64                     `json:"total_price" description:"订单总价，单位为分"`
    DetailURL  string                    `json:"detail_url" description:"小程序订单详情页 path，长度<=1024 byte (备注：该路径需要保证在小程序内配置过，相对路径即可）"`
    ItemList   []OrderDetailGeneralItem `json:"itemList" description:"订单商品列表"`
}

// OrderDetailGeneralItem order detail general item
type OrderDetailGeneralItem struct {
    ItemCode string `json:"item_code" description:"开发者侧商品 ID，长度 <= 64 byte"`
    Img      string `json:"img" description:"子订单商品图片 URL，长度 <= 512 byte"`
    Title    string `json:"title" description:"子订单商品标题，长度 <= 256 byte"`
    SubTitle string `json:"sub_title,omitempty" description:"子订单商品副标题，长度 <= 256 byte"`
    Amount   int64  `json:"amount" description:"子订单商品数量"`
    Price    int64  `json:"price" description:"子订单商品单价，单位为分"`
}

// OrderSyncResponse order sync response
type OrderSyncResponse struct {
    ErrCode int64  `json:"err_code" description:"错误码，0 代表成功，非 0 代表失败"`
    ErrMsg  string `json:"err_msg" description:"错误信息"`
    Body    string `json:"body" description:"POI 等关联业务推送结果，非 POI 订单为空，JSON 字符串"`
}
