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

package request

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"encoding/pem"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"golang.org/x/crypto/pkcs12"
)

const (
	headerAccessToken      = "access-token"
	headerContentType      = "Content-Type"
	headerContentTypeValue = "application/json;charset=utf-8"
	headerUserAgent        = "User-Agent"
	headerUserAgentValue   = `Mozilla/5.0 (Bytedance-Go-SDK; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36`
)

// DefaultRequest 默认请求
type DefaultRequest struct {
	AccessTokenKey string
}

// NewDefaultRequest 实例化
func NewDefaultRequest(accessTokenKey string) *DefaultRequest {
	return &DefaultRequest{AccessTokenKey: accessTokenKey}
}

// Get HTTP get request
func (srv *DefaultRequest) Get(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	accessToken := ctxValueToString(ctx, srv.AccessTokenKey)
	if strings.TrimSpace(accessToken) != "" {
		req.Header.Set(headerAccessToken, accessToken)
	}
	req.Header.Set(headerUserAgent, headerUserAgentValue)

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

// Post HTTP post request
func (srv *DefaultRequest) Post(ctx context.Context, url string, data []byte) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	accessToken := ctxValueToString(ctx, srv.AccessTokenKey)
	if strings.TrimSpace(accessToken) != "" {
		req.Header.Set(headerAccessToken, accessToken)
	}
	req.Header.Set(headerUserAgent, headerUserAgentValue)

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

// PostJSON HTTP post JSON request
func (srv *DefaultRequest) PostJSON(ctx context.Context, url string, data any) ([]byte, error) {
	var (
		jsonBuf = new(bytes.Buffer)
		enc     = json.NewEncoder(jsonBuf)
	)
	enc.SetEscapeHTML(false)
	if err := enc.Encode(data); err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, jsonBuf)
	if err != nil {
		return nil, err
	}
	req.Header.Set(headerContentType, headerContentTypeValue)
	req.Header.Set(headerUserAgent, headerUserAgentValue)
	accessToken := ctxValueToString(ctx, srv.AccessTokenKey)
	if strings.TrimSpace(accessToken) != "" {
		req.Header.Set(headerAccessToken, accessToken)
	}
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

// PostJSONWithRespContentType HTTP post JSON request with the response content type
func (srv *DefaultRequest) PostJSONWithRespContentType(ctx context.Context, url string, data any) ([]byte, string, error) {
	var (
		jsonBuf = new(bytes.Buffer)
		enc     = json.NewEncoder(jsonBuf)
	)
	enc.SetEscapeHTML(false)
	if err := enc.Encode(data); err != nil {
		return nil, "", err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, jsonBuf)
	if err != nil {
		return nil, "", err
	}
	req.Header.Set(headerContentType, headerContentTypeValue)
	req.Header.Set(headerUserAgent, headerUserAgentValue)
	accessToken := ctxValueToString(ctx, srv.AccessTokenKey)
	if strings.TrimSpace(accessToken) != "" {
		req.Header.Set(headerAccessToken, accessToken)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, "", err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	if resp.StatusCode != http.StatusOK {
		return nil, "", fmt.Errorf("http post error : uri=%v , statusCode=%v", url, resp.StatusCode)
	}
	res, err := io.ReadAll(resp.Body)
	contentType := resp.Header.Get(headerContentType)
	return res, contentType, err
}

// PostFile HTTP post file request
func (srv *DefaultRequest) PostFile(ctx context.Context, url string, files []MultipartFormField) ([]byte, error) {
	return srv.PostMultipartForm(ctx, url, files)
}

// PostMultipartForm HTTP post multipart form request
func (srv *DefaultRequest) PostMultipartForm(ctx context.Context, url string, files []MultipartFormField) (resp []byte, err error) {
	var (
		bodyBuf    = &bytes.Buffer{}
		bodyWriter = multipart.NewWriter(bodyBuf)
	)
	for _, field := range files {
		if field.IsFile {
			fileWriter, e := bodyWriter.CreateFormFile(field.FieldName, field.FileName)
			if e != nil {
				err = fmt.Errorf("error writing to buffer , err=%w", e)
				return
			}

			fh, e := os.Open(field.FileName)
			if e != nil {
				err = fmt.Errorf("error opening file , err=%w", e)
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
				err = fmt.Errorf("error writing to buffer , err=%w", e)
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
	req.Header.Set(headerContentType, contentType)
	req.Header.Set(headerUserAgent, headerUserAgentValue)
	accessToken := ctxValueToString(ctx, srv.AccessTokenKey)
	if strings.TrimSpace(accessToken) != "" {
		req.Header.Set(headerAccessToken, accessToken)
	}
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

// httpWithTLS CA 证书
func (srv *DefaultRequest) httpWithTLS(rootCa, key string) (*http.Client, error) {
	certData, err := os.ReadFile(rootCa)
	if err != nil {
		return nil, fmt.Errorf("unable to find cert path=%s, error=%v", rootCa, err)
	}
	cert := srv.pkcs12ToPem(certData, key)
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	tr := &http.Transport{
		TLSClientConfig:    config,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	return client, nil
}

// pkcs12ToPem 将 Pkcs12 转成 Pem
func (srv *DefaultRequest) pkcs12ToPem(p12 []byte, password string) tls.Certificate {
	blocks, err := pkcs12.ToPEM(p12, password)
	defer func() {
		if x := recover(); x != nil {
			log.Print(x)
		}
	}()
	if err != nil {
		panic(err)
	}
	var pemData []byte
	for _, b := range blocks {
		pemData = append(pemData, pem.EncodeToMemory(b)...)
	}
	cert, err := tls.X509KeyPair(pemData, pemData)
	if err != nil {
		panic(err)
	}
	return cert
}

// PostXML perform the HTTP/POST request with XML body
func (srv *DefaultRequest) PostXML(ctx context.Context, url string, data any) ([]byte, error) {
	xmlData, err := xml.Marshal(data)
	if err != nil {
		return nil, err
	}

	body := bytes.NewBuffer(xmlData)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set(headerContentType, "application/xml;charset=utf-8")
	accessToken := ctxValueToString(ctx, srv.AccessTokenKey)
	if strings.TrimSpace(accessToken) != "" {
		req.Header.Set(headerAccessToken, accessToken)
	}
	req.Header.Set(headerUserAgent, headerUserAgentValue)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = response.Body.Close()
	}()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http code error : uri=%v , statusCode=%v", url, response.StatusCode)
	}
	return io.ReadAll(response.Body)
}

// PostXMLWithTLS perform the HTTP/POST request with XML body and TLS
func (srv *DefaultRequest) PostXMLWithTLS(ctx context.Context, url string, data any, ca, key string) ([]byte, error) {
	xmlData, err := xml.Marshal(data)
	if err != nil {
		return nil, err
	}
	client, err := srv.httpWithTLS(ca, key)
	if err != nil {
		return nil, err
	}

	body := bytes.NewBuffer(xmlData)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set(headerContentType, "application/xml;charset=utf-8")
	accessToken := ctxValueToString(ctx, srv.AccessTokenKey)
	if strings.TrimSpace(accessToken) != "" {
		req.Header.Set(headerAccessToken, accessToken)
	}
	req.Header.Set(headerUserAgent, headerUserAgentValue)
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = response.Body.Close()
	}()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http code error : uri=%v , statusCode=%v", url, response.StatusCode)
	}
	return io.ReadAll(response.Body)
}

func ctxValueToString(ctx context.Context, key string) string {
	val := ctx.Value(key)
	if val == nil {
		return ""
	}

	strVal, ok := val.(string)
	if !ok {
		strVal = fmt.Sprintf("%v", val)
	}

	return strVal
}
