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

