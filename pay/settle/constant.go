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

const (
    applySettle = "https://open.douyin.com/api/trade_basic/v1/developer/settle_create/"
    querySettle = "https://open.douyin.com/api/trade_basic/v1/developer/settle_query/"
)

const (
    // StateInit 分账状态：INIT：初始化，PROCESSING：处理中，SUCCESS：处理成功，FAIL：处理失败
    StateInit = "INIT"
    // StateProcessing 分账状态：INIT：初始化，PROCESSING：处理中，SUCCESS：处理成功，FAIL：处理失败
    StateProcessing = "PROCESSING"
    // StateSuccess 分账状态：INIT：初始化，PROCESSING：处理中，SUCCESS：处理成功，FAIL：处理失败
    StateSuccess = "SUCCESS"
    // StateFail 分账状态：INIT：初始化，PROCESSING：处理中，SUCCESS：处理成功，FAIL：处理失败
    StateFail = "FAIL"
)
