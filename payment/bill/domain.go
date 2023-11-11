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

package bill

import (
    "net/url"
    "strings"
)

// QueryBillRequest query bill request
type QueryBillRequest struct {
    BillDate     string `json:"bill_date" desc:"账单日期，格式为 yyyyMMdd/yyyyMM"`
    MerchantID   string `json:"merchant_id" desc:"商户号"`
    BillType     string `json:"bill_type" desc:"账单类型，包括 payment:支付账单，settle:分账账单，refund:退款账单，return:退分账账单，withdraw:提现账单，rebate:返佣账单，annual_rebate:年框返佣账单"`
    AppID        string `json:"app_id" desc:"应用 ID"`
    ThirdPartyID string `json:"thirdparty_id" desc:"第三方应用 ID"`
    Sign         string `json:"sign" desc:"签名"`
}

// QueryBillResponse query bill response
type QueryBillResponse struct {
    ErrorNo  int    `json:"error_no" desc:"错误码"`
    ErrorMsg string `json:"error_msg" desc:"错误信息"`
}

func (b QueryBillRequest) ToURLValues() url.Values {
    uv := url.Values{}
    uv.Set("bill_date", b.BillDate)
    uv.Set("merchant_id", b.MerchantID)
    uv.Set("bill_type", b.BillType)
    uv.Set("app_id", b.AppID)
    if strings.TrimSpace(b.ThirdPartyID) != "" {
        uv.Set("thirdparty_id", b.ThirdPartyID)
    }
    uv.Set("sign", b.Sign)
    return uv
}
