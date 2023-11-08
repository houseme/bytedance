/*
 *  Copyright bytedance Author(https://houseme.github.io/bytedance/). All Rights Reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 *  You can obtain one at https://github.com/houseme/bytedance.
 *
 */

// Package base universal tool package
package base

import (
    "fmt"
)

// CommonError 抖音返回的通用错误。
type CommonError struct {
    ErrCode int64  `json:"error_code"`
    ErrMsg  string `json:"description"`
}

// CommonErrorExtra 抖音返回的错误额外信息。
type CommonErrorExtra struct {
    LogID string `json:"logid"`
    Now   int64  `json:"now"`
}

var (
    // ErrConfigNotFound config didn't find
    ErrConfigNotFound = Error{
        ErrCode: 10404,
        ErrMsg:  "config not found",
    }
    // ErrConfigKeyValueEmpty client key not found
    errConfigKeyValueEmpty = Error{
        ErrCode: 10404,
        ErrMsg:  "params key not found",
    }
    
    // ErrRequestIsEmpty request is empty
    ErrRequestIsEmpty = Error{
        ErrCode: 10404,
        ErrMsg:  "request is empty",
    }
    
    ErrClientTokenIsEmpty = Error{
        ErrCode: 10404,
        ErrMsg:  "client token is empty",
    }
)

// ErrConfigKeyValueEmpty params key not found
func ErrConfigKeyValueEmpty(key string) error {
    if key == "" {
        return errConfigKeyValueEmpty
    }
    return Error{
        ErrCode: 10404,
        ErrMsg:  fmt.Sprintf("%s not found", key),
    }
}

// Error base error
type Error struct {
    ErrCode int    `json:"errCode"`
    ErrMsg  string `json:"errMsg"`
    Err     error  `json:"err"`
}

// String return the error string
func (e Error) Error() string {
    return fmt.Sprintf("errCode: %d, errMsg: %s, err: %s", e.ErrCode, e.ErrMsg, e.Err)
}
