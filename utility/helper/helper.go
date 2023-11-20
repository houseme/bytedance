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

// Package helper universal tool package
package helper

import (
    "bytes"
    "crypto"
    "crypto/rand"
    "crypto/rsa"
    "crypto/sha256"
    "crypto/x509"
    "encoding/base64"
    "encoding/pem"
    "errors"
    "fmt"
    "sort"
    "strings"
    "time"
    "unicode"
)

// GenSign 生成签名
func GenSign(method, url, timestamp, nonce, body string, privateKey *rsa.PrivateKey) (sign string, err error) {
    // method 内容必须大写，如 GET、POST，URI 不包含域名，必须以'/'开头
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

// Template 对字符串中的和 map 的 key 相同的字符串进行模板替换 仅支持 形如：{name}
func Template(source string, data map[string]interface{}) string {
    sourceCopy := &source
    for k, val := range data {
        valStr := ""
        switch v := val.(type) {
        case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
            valStr = fmt.Sprintf("%d", v)
        case bool:
            valStr = fmt.Sprintf("%v", v)
        default:
            valStr = fmt.Sprintf("%s", v)
        }
        *sourceCopy = strings.Replace(*sourceCopy, strings.Join([]string{"{", k, "}"}, ""), valStr, 1)
    }
    return *sourceCopy
}

// GetCurrTS return current timestamps
func GetCurrTS() int64 {
    return time.Now().Unix()
}

// SliceChunk 用于将字符串切片分块
func SliceChunk(src []string, chunkSize int) (chunks [][]string) {
    total := len(src)
    chunks = make([][]string, 0)
    if chunkSize < 1 {
        chunkSize = 1
    }
    if total == 0 {
        return
    }
    
    chunkNum := total / chunkSize
    if total%chunkSize != 0 {
        chunkNum++
    }
    
    chunks = make([][]string, chunkNum)
    for i := 0; i < chunkNum; i++ {
        for j := 0; j < chunkSize; j++ {
            offset := i*chunkSize + j
            if offset >= total {
                return
            }
            
            if chunks[i] == nil {
                actualChunkSize := chunkSize
                if i == chunkNum-1 && total%chunkSize != 0 {
                    actualChunkSize = total % chunkSize
                }
                chunks[i] = make([]string, actualChunkSize)
            }
            
            chunks[i][j] = src[offset]
        }
    }
    
    return
}

// Query 将 Map 序列化为 Query 参数
func Query(params map[string]interface{}) string {
    finalString := make([]string, 0)
    for key, value := range params {
        valueString := ""
        switch v := value.(type) {
        case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
            valueString = fmt.Sprintf("%d", v)
        case bool:
            valueString = fmt.Sprintf("%v", v)
        default:
            valueString = fmt.Sprintf("%s", v)
        }
        finalString = append(finalString, strings.Join([]string{key, valueString}, "="))
    }
    return strings.Join(finalString, "&")
}

// OrderParam order params
func OrderParam(p map[string]string, bizKey string) (returnStr string) {
    keys := make([]string, 0, len(p))
    for k := range p {
        if k == "sign" {
            continue
        }
        keys = append(keys, k)
    }
    sort.Strings(keys)
    var buf bytes.Buffer
    for _, k := range keys {
        if p[k] == "" {
            continue
        }
        if buf.Len() > 0 {
            buf.WriteByte('&')
        }
        buf.WriteString(k)
        buf.WriteByte('=')
        buf.WriteString(p[k])
    }
    buf.WriteString(bizKey)
    returnStr = buf.String()
    return
}

// RSADecrypt 数据解密
func RSADecrypt(privateKey string, ciphertext []byte) ([]byte, error) {
    block, _ := pem.Decode([]byte(privateKey))
    if block == nil {
        return nil, errors.New("PrivateKey format error")
    }
    privy, err := x509.ParsePKCS1PrivateKey(block.Bytes)
    if err != nil {
        oldErr := err
        key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
        if err != nil {
            return nil, fmt.Errorf("ParsePKCS1PrivateKey error: %s, ParsePKCS8PrivateKey error: %s", oldErr.Error(), err.Error())
        }
        switch t := key.(type) {
        case *rsa.PrivateKey:
            privy = key.(*rsa.PrivateKey)
        default:
            return nil, fmt.Errorf("ParsePKCS1PrivateKey error: %s, ParsePKCS8PrivateKey error: Not supported privatekey format, should be *rsa.PrivateKey, got %T", oldErr.Error(), t)
        }
    }
    return rsa.DecryptPKCS1v15(rand.Reader, privy, ciphertext)
}

// RSADecryptBase64 Base64 解码后再次进行 RSA 解密
func RSADecryptBase64(privateKey string, cryptoText string) ([]byte, error) {
    encryptedData, err := base64.StdEncoding.DecodeString(cryptoText)
    if err != nil {
        return nil, err
    }
    return RSADecrypt(privateKey, encryptedData)
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
