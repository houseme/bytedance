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

// Package constant  payment constant
package constant

const (
    // CreateOrder creates order
    CreateOrder = "https://developer.toutiao.com/api/apps/ecpay/v1/create_order"
    
    // QueryOrder query order
    QueryOrder = "https://developer.toutiao.com/api/apps/ecpay/v1/query_order"
    
    // CreateSettle settle trade
    CreateSettle = "https://developer.toutiao.com/api/apps/ecpay/v1/settle"
    
    // QuerySettle query settle
    QuerySettle = "https://developer.toutiao.com/api/apps/ecpay/v1/query_settle"
    
    // UnsettleAmount query unsettle amount
    UnsettleAmount = "https://developer.toutiao.com/api/apps/ecpay/v1/unsettle_amount"
    
    // QueryPlatformOrder query platform order 自动结算结果查询
    QueryPlatformOrder = "https://developer.toutiao.com/api/apps/ecpay/v1/query_platform_order"
    
    // CreateRefund create refund
    CreateRefund = "https://developer.toutiao.com/api/apps/ecpay/v1/create_refund"
    
    // QueryRefund query refund
    QueryRefund = "https://developer.toutiao.com/api/apps/ecpay/v1/query_refund"
    
    // OrderPush order push
    OrderPush = "https://developer.toutiao.com/api/apps/order/v2/push"
    
    // QueryMerchantBalance query merchant balance
    QueryMerchantBalance = "https://developer.toutiao.com/api/apps/ecpay/saas/query_merchant_balance"
    
    // MerchantWithdraw merchant withdraws
    MerchantWithdraw = "https://developer.toutiao.com/api/apps/ecpay/saas/merchant_withdraw"
    
    // QueryWithdrawOrder query withdraws order
    QueryWithdrawOrder = "https://developer.toutiao.com/api/apps/ecpay/saas/query_withdraw_order"
)

const (
    // AsyncPay 异步支付类型
    AsyncPay = "payment"
    
    // AsyncSettle 异步结算类型
    AsyncSettle = "settle"
    
    // AsyncRefund 异步退款类型
    AsyncRefund = "refund"
    
    // AsyncWithdraw 异步提现类型
    AsyncWithdraw = "withdraw"
    
    // AsyncTransfer 异步转账类型
    AsyncTransfer = "transfer"
    
    // AsyncSettleFinish 异步结算完成类型 自动分账
    AsyncSettleFinish = "settle_finish"
)

const (
    // SettleFail 结算失败
    SettleFail = "FAIL"
    // SettleSuccess 结算成功
    SettleSuccess = "SUCCESS"
)

const (
    // LimitWx 屏蔽微信支付
    LimitWx = "LIMIT_WX"
    // LimitAli 屏蔽支付宝支付
    LimitAli = "LIMIT_ALI"
    // LimitDyzf 屏蔽抖音支付
    LimitDyzf = "LIMIT_DYZF"
)

const (
    // DisableMsgOne 是否屏蔽消息  1-屏蔽 0-非屏蔽
    DisableMsgOne = 1
    // DisableMsgZero 是否屏蔽消息  1-屏蔽 0-非屏蔽
    DisableMsgZero = 0
)

const (
    // Success 通知成功 0
    Success = 0
    
    // FailedToCheckTheSignature 验签失败
    FailedToCheckTheSignature = 400
)
