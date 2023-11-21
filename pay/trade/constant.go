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

const (
    // NumberCardCommodityTagID NumberCardCommodity 号卡商品 tag id tag_group_7272625659887943692
    NumberCardCommodityTagID = "tag_group_7272625659887943692"
    // CustomizedServiceRefundTagID 定制服务 tag id 定制后协商退：tag_group_7272625659887960076，定制后不可退：tag_group_7272625659887976460
    CustomizedServiceRefundTagID = "tag_group_7272625659887960076"
    // CustomizedServiceNoRefundTagID 定制服务 tag id 定制后协商退：tag_group_7272625659887960076，定制后不可退：tag_group_7272625659887976460
    CustomizedServiceNoRefundTagID = "tag_group_7272625659887976460"
    // ContentRechargeTagID 内容充值 tag id tag_group_7272625659888041996
    ContentRechargeTagID = "tag_group_7272625659888041996"
    // VirtualRechargeTagID 虚拟充值 tag id tag_group_7272625659887992844
    VirtualRechargeTagID = "tag_group_7272625659887992844"
    // GhostwritingDocumentTagID 代写文书 tag id tag_group_7297888175123382299
    GhostwritingDocumentTagID = "tag_group_7297888175123382299"
    // VirtualServiceTagID 虚拟服务 tag id tag_group_7272625659888058380
    VirtualServiceTagID = "tag_group_7272625659888058380"
    // GeneralConsultationRefundTagID 普通咨询 tag id 开始服务后协商退：tag_group_7272625659888009228，开始服务后不可退：tag_group_7272625659888025612
    GeneralConsultationRefundTagID = "tag_group_7272625659888009228"
    // GeneralConsultationNoRefundTagID 普通咨询 tag id 开始服务后协商退：tag_group_7272625659888009228，开始服务后不可退：tag_group_7272625659888025612
    GeneralConsultationNoRefundTagID = "tag_group_7272625659888025612"
)

const (
    // QueryTag query tag https://open.douyin.com/api/trade_basic/v1/developer/tag_query/
    // see: https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/trade-system/general/tag/tag_group_query
    QueryTag = "https://open.douyin.com/api/trade_basic/v1/developer/tag_query"
    
    // queryOrder query order https://open.douyin.com/api/trade_basic/v1/developer/order_query/
    // see: https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/trade-system/general/order/query_order
    queryOrder = "https://open.douyin.com/api/trade_basic/v1/developer/order_query"
)

const (
    // 支付超时时间，单位秒 默认值 300
    defaultPayExpireSeconds = 300
)
