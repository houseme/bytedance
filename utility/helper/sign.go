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

package helper

import (
    "context"
    "crypto/md5"
    "crypto/sha1"
    "fmt"
    "reflect"
    "sort"
    "strings"
    "unicode"
)

const (
    // OtherSettleParams 其他分账方参数 (Other settle params)
    OtherSettleParams = "other_settle_params"
    // AppID 小程序 appID (Applets appID)
    AppID = "app_id"
    // ThirdPartyID 代小程序进行该笔交易调用的第三方平台服务商 id (The id of the third-party platform service provider that calls the transaction on behalf of the Applets)
    ThirdPartyID = "thirdparty_id"
    // Sign 签名 (sign)
    Sign = "sign"
    // Timestamp 时间戳 (timestamp)
    Timestamp = "timestamp"
    // Nonce 随机字符串 (nonce)
    Nonce = "nonce"
    // Msg 消息体 (msg)
    Msg = "msg"
)

// CallbackSign 担保支付回调签名算法
// 参数："strArr" 所有字段（验证时注意不包含 sign 签名本身，不包含空字段与 type 常量字段）内容与平台上配置的 token
// 返回：签名字符串
//
// CallbackSign Guaranteed payment callback signature algorithm
// Param: "strArr" The content of all fields (note that the sign signature itself is not included during verification, and does not include empty fields and type constant fields) content and the token configured on the platform
// Return: signature string
func CallbackSign(_ context.Context, token string, data any) string {
    var (
        tt     = reflect.TypeOf(data)
        v      = reflect.ValueOf(data)
        count  = v.NumField()
        strArr = make([]string, 0)
    )
    strArr = append(strArr, token)
    for i := 0; i < count; i++ {
        if v.Field(i).CanInterface() { // 判断是否为可导出字段
            tagKey := LcFirst(tt.Field(i).Tag.Get("json"))
            if tagKey == "" {
                continue
            }
            tagKeyArr := strings.Split(tagKey, ",")
            if len(tagKeyArr) < 1 {
                continue
            }
            k := tagKeyArr[0]
            if !(k == Nonce || k == Timestamp || k == Msg) {
                continue
            }
            value := strings.TrimSpace(fmt.Sprintf("%v", v.Field(i).Interface()))
            if value == "" || value == "null" {
                continue
            }
            strArr = append(strArr, value)
        }
    }
    
    sort.Strings(strArr)
    
    var (
        signStr = strings.Join(strArr, "")
        h       = sha1.New()
    )
    h.Write([]byte(signStr))
    return fmt.Sprintf("%x", h.Sum(nil))
}

// RequestSign Guaranteed Payment Request Signature Algorithm
// Param: "paramsMap" all request parameters
// Return: signature string
func RequestSign(_ context.Context, data interface{}, salt string) string {
    var (
        tt    = reflect.TypeOf(data)
        v     = reflect.ValueOf(data)
        count = v.NumField()
    )
    
    var paramsArr []string
    for i := 0; i < count; i++ {
        if v.Field(i).CanInterface() { // 判断是否为可导出字段
            tagKey := LcFirst(tt.Field(i).Tag.Get("json"))
            if tagKey == "" {
                continue
            }
            tagKeyArr := strings.Split(tagKey, ",")
            if len(tagKeyArr) < 1 {
                continue
            }
            k := tagKeyArr[0]
            if k == OtherSettleParams || k == AppID || k == ThirdPartyID || k == Sign {
                continue
            }
            
            if reflect.ValueOf(v.Field(i).Interface()).Kind() == reflect.Ptr && reflect.ValueOf(v.Field(i).Interface()).IsNil() {
                continue
            }
            
            value := strings.TrimSpace(fmt.Sprintf("%v", v.Field(i).Interface()))
            if strings.HasPrefix(value, "\"") && strings.HasSuffix(value, "\"") && len(value) > 1 {
                value = value[1 : len(value)-1]
            }
            value = strings.TrimSpace(value)
            if value == "" || value == "null" {
                continue
            }
            fmt.Printf("%s=%s \n", k, value)
            paramsArr = append(paramsArr, value)
        }
    }
    
    paramsArr = append(paramsArr, salt)
    sort.Strings(paramsArr)
    return fmt.Sprintf("%x", md5.Sum([]byte(strings.Join(paramsArr, "&"))))
}

// UcFirst 首字母大些
func UcFirst(str string) string {
    for i, v := range str {
        return string(unicode.ToUpper(v)) + str[i+1:]
    }
    return ""
}

// LcFirst 首字母小写
func LcFirst(str string) string {
    for i, v := range str {
        return string(unicode.ToLower(v)) + str[i+1:]
    }
    return ""
}
