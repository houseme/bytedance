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

// Package helper universal tool package
package helper

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
)

// GenSign 生成签名
func GenSign(method, url, timestamp, nonce, body string, privateKey *rsa.PrivateKey) (sign string, err error) {
	// method 内容必须大写，如 GET、POST，uri 不包含域名，必须以'/'开头
	targetStr := method + "\n" + url + "\n" + timestamp + "\n" + nonce + "\n" + body + "\n"
	h := sha256.New()
	n, err := h.Write([]byte(targetStr))
	if err != nil {
		return "", err
	}
	if n != len(targetStr) {
		return "", fmt.Errorf("sha256 write bytes not match")
	}
	digestBytes := h.Sum(nil)
	signBytes, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, digestBytes)
	if err != nil {
		return "", err
	}
	sign = base64.StdEncoding.EncodeToString(signBytes)
	return
}

// CheckSign 校验签名
func CheckSign(timestamp, nonce, body, signature, pubKeyStr string) (bool, error) {
	pubKey, err := PemToRSAPublicKey(pubKeyStr)
	if err != nil {
		return false, err
	}

	hashed := sha256.Sum256([]byte(timestamp + "\n" + nonce + "\n" + body + "\n"))
	signBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false, err
	}
	err = rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, hashed[:], signBytes)
	return err == nil, nil
}

// PemToRSAPublicKey pem to rsa public key
func PemToRSAPublicKey(pemKeyStr string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pemKeyStr))
	if block == nil || len(block.Bytes) == 0 {
		return nil, fmt.Errorf("empty block in pem string")
	}
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	switch key := key.(type) {
	case *rsa.PublicKey:
		return key, nil
	default:
		return nil, fmt.Errorf("not rsa public key")
	}
}
