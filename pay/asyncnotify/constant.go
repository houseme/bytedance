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

// defaultVersion 默认版本号
const defaultVersion = "3.0"

const (
    // ErrNoSuccess success 错误码
    ErrNoSuccess = 0
    // ErrNoFailedToCheckTheSignature 验签失败
    ErrNoFailedToCheckTheSignature = 400
    // ErrNoRequestParameterError 请求参数错误
    ErrNoRequestParameterError = 401
    // ErrNoSystemError system error 错误码
    ErrNoSystemError = 10000
    
    // ErrTipsSuccess success 提示
    ErrTipsSuccess = "success"
    // ErrTipsSystemError system error 提示
    ErrTipsSystemError = "system error"
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
