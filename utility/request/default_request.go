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
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
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
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error : uri=%v , statusCode=%v", url, resp.StatusCode)
	}
	return io.ReadAll(resp.Body)
}

// Post http post request
func (srv *DefaultRequest) Post(ctx context.Context, url string, data []byte) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http post error : uri=%v , statusCode=%v", url, resp.StatusCode)
	}
	return io.ReadAll(resp.Body)
}

// PostJSON http post json request
func (srv *DefaultRequest) PostJSON(ctx context.Context, url string, data []byte) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	jsonData = bytes.Replace(jsonData, []byte("\\u003c"), []byte("<"), -1)
	jsonData = bytes.Replace(jsonData, []byte("\\u003e"), []byte(">"), -1)
	jsonData = bytes.Replace(jsonData, []byte("\\u0026"), []byte("&"), -1)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http post error : uri=%v , statusCode=%v", url, resp.StatusCode)
	}
	return io.ReadAll(resp.Body)
}

// PostFile http post file request
func (srv *DefaultRequest) PostFile(ctx context.Context, url string, files []MultipartFormField) ([]byte, error) {
	return nil, nil
}

// PostMultipartForm http post multipart form request
func (srv *DefaultRequest) PostMultipartForm(ctx context.Context, url string, files []MultipartFormField) (resp []byte, err error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	for _, field := range files {
		if field.IsFile {
			fileWriter, e := bodyWriter.CreateFormFile(field.FieldName, field.FileName)
			if e != nil {
				err = fmt.Errorf("error writing to buffer , err=%v", e)
				return
			}

			fh, e := os.Open(field.FileName)
			if e != nil {
				err = fmt.Errorf("error opening file , err=%v", e)
				return
			}

			if _, err = io.Copy(fileWriter, fh); err != nil {
				_ = fh.Close()
				return
			}
			_ = fh.Close()
		} else {
			partWriter, e := bodyWriter.CreateFormField(field.FieldName)
			if e != nil {
				err = fmt.Errorf("error writing to buffer , err=%v", e)
				return
			}
			valueReader := bytes.NewReader(field.Value)
			if _, err = io.Copy(partWriter, valueReader); err != nil {
				return
			}
		}
	}

	contentType := bodyWriter.FormDataContentType()
	_ = bodyWriter.Close()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bodyBuf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = response.Body.Close()
	}()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http post error : uri=%v , statusCode=%v", url, response.StatusCode)
	}
	return io.ReadAll(response.Body)
}
