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

package solution

// Solution 解决方案

// CreateSolutionRequest 创建解决方案
type CreateSolutionRequest struct {
    Operator          string           `json:"operator" description:"操作者"`
    ReleaseReason     string           `json:"release_reason" description:"发布原因"`
    IndustryImplList  []*IndustryImpl  `json:"industry_impl_list" description:"行业实现列表"`
    AppConfigItemList []*AppConfigItem `json:"app_config_item_list" description:"应用配置项列表"`
}

// IndustryImpl 行业实现
type IndustryImpl struct {
    SolutionID          string             `json:"solution_id,omitempty" description:"解决方案 ID"`
    TemplateID          int64              `json:"template_id,omitempty" description:"模板 ID"`
    OpenAbilityImplList []*OpenAbilityImpl `json:"open_ability_impl_list" description:"开放能力实现列表"`
}

// OpenAbilityImpl 开放能力实现
type OpenAbilityImpl struct {
    AbilityIdentity string `json:"ability_identity" description:"能力标识"`
    ISDelete        bool   `json:"is_delete" description:"是否删除"`
    TestURL         string `json:"test_url,omitempty" description:"测试 URL is_delete 为 false 时，必填"`
    ProdURL         string `json:"prod_url,omitempty" description:"生产 URL is_delete 为 false 时，必填"`
    AbilityType     int64  `json:"ability_type" description:"开放能力类型：1 为扩展点，2 为消息"`
    ImplName        string `json:"impl_name" description:"实现名称"`
}

// AppConfigItem 应用配置项
type AppConfigItem struct {
    Identity string `json:"identity" description:"配置项标识"`
    Value    string `json:"value" description:"配置项值"`
}

// CreateSolutionResponse 创建解决方案响应
type CreateSolutionResponse struct {
    Extra *Extra        `json:"extra"`
    Data  *ResponseData `json:"data"`
}

// QuerySolutionRequest 查询解决方案
type QuerySolutionRequest struct {
    SolutionIDList []string `json:"solution_id_list,omitempty" description:"解决方案 ID 列表"`
    TemplateIDList []int    `json:"template_id_list,omitempty" description:"模板 ID 列表"`
}

// QuerySolutionResponse 查询解决方案响应
type QuerySolutionResponse struct {
    Extra                      *Extra            `json:"extra"`
    SolutionEffAbilityImplList []*EffAbilityImpl `json:"solution_eff_ability_impl_list"`
    Data                       *ResponseData     `json:"data"`
}

// ResponseData 查询解决方案数据
type ResponseData struct {
    ErrorCode   int    `json:"error_code"`
    Description string `json:"description"`
}

// Extra 额外信息
type Extra struct {
    ErrorCode      int    `json:"error_code"`
    Description    string `json:"description"`
    SubErrorCode   int    `json:"sub_error_code"`
    SubDescription string `json:"sub_description"`
    LogID          string `json:"logid"`
    Now            int    `json:"now"`
}

// MessageAbilityImpl 消息能力实现
type MessageAbilityImpl struct {
    ProtocolType    int    `json:"protocol_type"`
    TestURL         string `json:"test_url"`
    AbilityIdentity string `json:"ability_identity"`
    AppID           string `json:"app_id"`
    IndustryCode    string `json:"industry_code"`
    Name            string `json:"name"`
    ProdURL         string `json:"prod_url"`
}

// ExtensionAbilityImpl 扩展能力实现
type ExtensionAbilityImpl struct {
    AbilityIdentity string `json:"ability_identity"`
    AppID           string `json:"app_id"`
    IndustryCode    string `json:"industry_code"`
    Name            string `json:"name"`
    ProdURL         string `json:"prod_url"`
    ProtocolType    int    `json:"protocol_type"`
    TestURL         string `json:"test_url"`
}

// EffAbilityImpl 解决方案有效能力实现
type EffAbilityImpl struct {
    ExtensionAbilityImplList []*ExtensionAbilityImpl `json:"extension_ability_impl_list"`
    MessageAbilityImplList   []*MessageAbilityImpl   `json:"message_ability_impl_list"`
    SolutionID               string                  `json:"solution_id"`
}
