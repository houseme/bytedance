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

package request

import (
	"context"
)

// DefaultRequest 默认请求
type DefaultRequest struct {
}

// NewDefaultRequest 实例化
func NewDefaultRequest() *DefaultRequest {
	return &DefaultRequest{}
}

// Get http get request
func (srv *DefaultRequest) Get(ctx context.Context, url string) ([]byte, error) {
	return nil, nil
}

// Post http post request
func (srv *DefaultRequest) Post(ctx context.Context, url string, data []byte) ([]byte, error) {
	return nil, nil
}

// PostJSON http post json request
func (srv *DefaultRequest) PostJSON(ctx context.Context, url string, data []byte) ([]byte, error) {
	return nil, nil
}

// PostFile http post file request
func (srv *DefaultRequest) PostFile(ctx context.Context, url string, data []byte, files []MultipartFormField) ([]byte, error) {
	return nil, nil
}

// PostMultipartForm http post multipart form request
func (srv *DefaultRequest) PostMultipartForm(ctx context.Context, url string, data []byte, files []MultipartFormField) ([]byte, error) {
	return nil, nil
}
