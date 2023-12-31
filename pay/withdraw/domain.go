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

// QueryBalanceRequest 查询余额
type QueryBalanceRequest struct {
    ThirdPartyID   string `json:"thirdparty_id,omitempty"`
    AppID          string `json:"app_id,omitempty" description:"小程序的 app_id。在服务商为自己提现的情况下可不填，其他情况必填"`
    MerchantUID    string `json:"merchant_uid" description:"进件完成返回的商户号"`
    ChannelType    string `json:"channel_type" description:"提现渠道枚举值:alipay: 担保支付普通版支付宝，wx: 担保支付普通版微信，hz: 担保支付普通版抖音支付，yzt: 担保支付企业版聚合账户"`
    MerchantEntity int    `json:"merchant_entity,omitempty" description:"抖音信息和光合信号主体标识：不传或传 0 或 1 查抖音信息主体账户余额，传 2 查光合信号主体账户余额"`
}

// QueryBalanceResponse query merchant account response
type QueryBalanceResponse struct {
    ErrNo  int               `json:"err_no"`
    ErrMsg string            `json:"err_msg"`
    LogID  string            `json:"log_id"`
    Data   *QueryBalanceData `json:"data"`
}

// QueryBalanceData query merchant account response
type QueryBalanceData struct {
    AccountInfo    AccountInfo `json:"account_info" description:"余额信息"`
    SettleInfo     SettleInfo  `json:"settle_info" description:"结算信息"`
    MerchantEntity int         `json:"merchant_entity" description:"抖音信息和光合信号标识：1: 当前余额所属抖音信息，2: 当前余额所属光合信号"`
}

// AccountInfo account info
type AccountInfo struct {
    OnlineBalance       int `json:"online_balance" description:"在途余额；CNY、单位分"`
    WithdrawAbleBalance int `json:"withdrawable_balacne" description:"可提现余额；CNY、单位分"`
    FreezeBalance       int `json:"freeze_balance" description:"冻结余额；CNY、单位分"`
}

// SettleInfo settle info
type SettleInfo struct {
    SettleType    int    `json:"settle_type" description:"结算类型枚举值：1: 银行卡结算，2: 支付宝结算"`
    SettleAccount string `json:"settle_account" description:"结算账户，支付宝结算时，支付宝账号"`
    BankcardNo    string `json:"bankcard_no" description:"银行卡结算时，银行卡号"`
    BankName      string `json:"bank_name" description:"银行卡结算时，银行卡对应银行名称"`
}

// MerchantWithdrawRequest 商户提现
type MerchantWithdrawRequest struct {
    ThirdPartyID   string `json:"thirdparty_id,omitempty"`
    AppID          string `json:"app_id,omitempty"`
    MerchantUID    string `json:"merchant_uid"`
    ChannelType    string `json:"channel_type"`
    WithdrawAmount int    `json:"withdraw_amount"`
    OutOrderID     string `json:"out_order_id"`
}

// MerchantWithdrawResponse 商户提现
type MerchantWithdrawResponse struct {
    ErrNo  int                   `json:"err_no"`
    ErrMsg string                `json:"err_msg"`
    LogID  string                `json:"log_id"`
    Data   *MerchantWithdrawData `json:"data"`
}

// MerchantWithdrawData 商户提现
type MerchantWithdrawData struct {
    OrderID        string `json:"order_id"`
    MerchantEntity int    `json:"merchant_entity"`
}

// QueryMerchantWithdrawRequest 查询商户提现
type QueryMerchantWithdrawRequest struct {
    ThirdPartyID string `json:"thirdparty_id,omitempty"`
    AppID        string `json:"app_id"`
    MerchantUID  string `json:"merchant_uid"`
    ChannelType  string `json:"channel_type"`
    OutOrderID   string `json:"out_order_id"`
}

// QueryMerchantWithdrawResponse 查询商户提现
type QueryMerchantWithdrawResponse struct {
    ErrNo  int                        `json:"err_no"`
    ErrMsg string                     `json:"err_msg"`
    LogId  string                     `json:"log_id"`
    Data   *QueryMerchantWithdrawData `json:"data"`
}

// QueryMerchantWithdrawData 查询商户提现
type QueryMerchantWithdrawData struct {
    Status    string `json:"status"`
    StatusMsg string `json:"statusMsg"`
}
