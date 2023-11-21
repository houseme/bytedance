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

package settle

// ApplySettleRequest 申请分账
type ApplySettleRequest struct {
    OutOrderNo   string `json:"out_order_no"`
    OutSettleNo  string `json:"out_settle_no"`
    ItemOrderID  string `json:"item_order_id,omitempty"`
    SettleDesc   string `json:"settle_desc"`
    SettleParams string `json:"settle_params" desc:"其他分账方（除卖家之外的），长度 <= 512 字节 [{\"merchant_uid\":\"merchant_uid_example\",\"amount\":100}]"`
    Ext          string `json:"ext" desc:"开发者自定义透传字段，不支持二进制，会在查询分账接口 cp_extra 字段原样返回，长度 <= 2048 字节"`
    NotifyURL    string `json:"notify_url"`
}

// OtherSettleParam 其他结算参数
type OtherSettleParam struct {
    MerchantUid string `json:"merchant_uid"`
    Amount      int    `json:"amount"`
}

// ApplySettleResponse 申请分账
type ApplySettleResponse struct {
    ErrMsg string           `json:"err_msg"`
    ErrNo  int              `json:"err_no"`
    LogID  string           `json:"log_id"`
    Data   *ApplySettleData `json:"data"`
}

type ApplySettleData struct {
    WalletSettleID string `json:"wallet_settle_id" desc:"小程序底层分账单号"`
    SettleID       string `json:"settle_id" desc:"小程序侧分账单号"`
}

// QuerySettleRequest 查询分账
// 以上 4 个参数选填一个，查询优先级：settle_id > order_id > out_settle_no > out_order_no。
// 例如：请求填写了 settle_id 和 order_id，服务只会按 settle_id 来查询，忽略 order_id。
// 如果未查询到结果，会返回空数组。
type QuerySettleRequest struct {
    OutOrderNo  string `json:"out_order_no,omitempty"`
    OutSettleNo string `json:"out_settle_no,omitempty"`
    OrderID     string `json:"order_id,omitempty"`
    SettleID    string `json:"settle_id,omitempty"`
    AppID       string `json:"app_id"`
}

// QuerySettleResponse 查询分账
type QuerySettleResponse struct {
    Data   []*QuerySettleData `json:"data"`
    ErrMsg string             `json:"err_msg"`
    ErrNo  int                `json:"err_no"`
    LogID  string             `json:"log_id"`
}

// QuerySettleData 查询分账
type QuerySettleData struct {
    SettleAmount   int    `json:"settle_amount"`
    CpExtra        string `json:"cp_extra"`
    OutSettleID    string `json:"out_settle_id"`
    OrderID        string `json:"order_id"`
    OutOrderID     string `json:"out_order_id"`
    PlatformTicket int    `json:"platform_ticket"`
    Rake           int    `json:"rake"`
    SettleAt       int64  `json:"settle_at"`
    SettleDetail   string `json:"settle_detail"`
    Commission     int    `json:"commission"`
    ItemOrderID    string `json:"item_order_id"`
    SettleID       string `json:"settle_id"`
    SettleStatus   string `json:"settle_status"`
}
