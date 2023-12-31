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
)

// AccessTokenHandle AccessToken 接口
type AccessTokenHandle interface {
    GetAccessToken(ctx context.Context, openID string) (accessToken string, err error)
    SetAccessToken(ctx context.Context, accessToken *AccessToken) (err error)
    GetClientToken(ctx context.Context) (clientToken *ClientToken, err error)
    GetServerAccessToken(ctx context.Context) (serverAccessToken *ServerAccessToken, err error)
}
