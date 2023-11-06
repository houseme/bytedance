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

package credential

import (
    "context"
    "encoding/json"
    "fmt"
    "sync"
    "time"
    
    "github.com/houseme/bytedance/utility/base"
    "github.com/houseme/bytedance/utility/cache"
    "github.com/houseme/bytedance/utility/request"
)

// 获取 ticket 的 url
const getTicketURL = "https://open.douyin.com/js/getticket?access_token=%s"

// DefaultJsTicket 默认获取 js ticket 方法
type DefaultJsTicket struct {
    appID          string
    cacheKeyPrefix string
    cache          cache.Cache
    request        request.Request
    // jsAPITicket 读写锁 同一个 AppID 一个
    jsAPITicketLock *sync.Mutex
}

// NewDefaultJsTicket new
func NewDefaultJsTicket(_ context.Context, appID, cacheKeyPrefix string, cache cache.Cache, req request.Request) JsTicketHandle {
    return &DefaultJsTicket{
        appID:           appID,
        cache:           cache,
        cacheKeyPrefix:  cacheKeyPrefix,
        request:         req,
        jsAPITicketLock: new(sync.Mutex),
    }
}

// Ticket 请求 jsapi_ticket 返回结果
type Ticket struct {
    base.CommonError
    Ticket    string `json:"ticket"`
    ExpiresIn int64  `json:"expires_in"`
}

// GetTicket 获取 jsapi_ticket
func (t *DefaultJsTicket) GetTicket(ctx context.Context, accessToken string) (ticketStr string, err error) {
    t.jsAPITicketLock.Lock()
    defer t.jsAPITicketLock.Unlock()
    
    // 先从 cache 中取
    jsAPITicketCacheKey := fmt.Sprintf("%s_jsapi_ticket_%s", t.cacheKeyPrefix, t.appID)
    if val := t.cache.Get(ctx, jsAPITicketCacheKey); val != nil {
        ticketStr = val.(string)
        return
    }
    var ticket Ticket
    if ticket, err = GetTicketFromServer(ctx, accessToken, t.request); err != nil {
        return
    }
    expires := ticket.ExpiresIn - 1500
    err = t.cache.Set(ctx, jsAPITicketCacheKey, ticket.Ticket, time.Duration(expires)*time.Second)
    ticketStr = ticket.Ticket
    return
}

// GetTicketFromServer 从服务器中获取 ticket
func GetTicketFromServer(ctx context.Context, accessToken string, req request.Request) (ticket Ticket, err error) {
    var response []byte
    if response, err = req.Get(ctx, fmt.Sprintf(getTicketURL, accessToken)); err != nil {
        return
    }
    if err = json.Unmarshal(response, &ticket); err != nil {
        return
    }
    if ticket.ErrCode != 0 {
        err = fmt.Errorf("getTicket Error : errcode=%d , errmsg=%s", ticket.ErrCode, ticket.ErrMsg)
        return
    }
    return
}
