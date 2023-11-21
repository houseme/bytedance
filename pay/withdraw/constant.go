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

const (
    queryMerchantBalance  = "https://open.douyin.com/api/apps/ecpay/saas/query_merchant_balance"
    applyMerchantWithdraw = "https://open.douyin.com/api/apps/ecpay/saas/merchant_withdraw"
    queryWithdrawOrder    = "https://open.douyin.com/api/apps/ecpay/saas/query_withdraw_order"
)

const (
    // Alipay 提现渠道枚举值:alipay: 担保支付普通版支付宝，wx: 担保支付普通版微信，hz: 担保支付普通版抖音支付，yzt: 担保支付企业版聚合账户
    Alipay = "alipay"
    // Wx 微信
    Wx = "wx"
    // Hz 抖音
    Hz = "hz"
    // Yzt 聚合账户
    Yzt = "yzt"
)

const (
    // MerchantEntityDefault 抖音信息和光合信号主体标识：不传或传 0 或 1 查抖音信息主体账户余额，传 2 查光合信号主体账户余额
    MerchantEntityDefault = 0
    // MerchantEntityDy 抖音信息主体
    MerchantEntityDy = 1
    // MerchantEntityGh 光合信号主体
    MerchantEntityGh = 2
)

const (
    // StateSuccess 提现状态 状态枚举值：成功:SUCCESS，失败：FAIL，处理中：PROCESSING，退票：REEXCHANGE
    StateSuccess = "SUCCESS"
    // StateFail 提现状态 状态枚举值：成功:SUCCESS，失败：FAIL，处理中：PROCESSING，退票：REEXCHANGE
    StateFail = "FAIL"
    // StateProcessing 提现状态 状态枚举值：成功:SUCCESS，失败：FAIL，处理中：PROCESSING，退票：REEXCHANGE
    StateProcessing = "PROCESSING"
    // StateReExchange 提现状态 状态枚举值：成功:SUCCESS，失败：FAIL，处理中：PROCESSING，退票：REEXCHANGE
    StateReExchange = "REEXCHANGE"
)
