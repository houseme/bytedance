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

// Package request http request
package request

import (
    "context"
)

// Request http request interface
type Request interface {
    Get(ctx context.Context, url string) ([]byte, error)
    Post(ctx context.Context, url string, data []byte) ([]byte, error)
    PostJSON(ctx context.Context, url string, data any) ([]byte, error)
    PostJSONWithRespContentType(ctx context.Context, url string, data any) ([]byte, string, error)
    PostFile(ctx context.Context, url string, files []MultipartFormField) ([]byte, error)
    PostMultipartForm(ctx context.Context, url string, files []MultipartFormField) ([]byte, error)
    PostXML(ctx context.Context, url string, data any) ([]byte, error)
    PostXMLWithTLS(ctx context.Context, url string, data any, ca, key string) ([]byte, error)
}

// MultipartFormField multipart form field
type MultipartFormField struct {
    IsFile    bool   `json:"isFile"`
    Value     []byte `json:"value"`
    FieldName string `json:"fieldName"`
    FileName  string `json:"fileName"`
}
