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

package logger

import (
	"context"
)

// DefaultLogger 默认日志
type DefaultLogger struct {
}

// NewDefaultLogger 实例化
func NewDefaultLogger() *DefaultLogger {
	return &DefaultLogger{}
}

// Print 打印
func (logger *DefaultLogger) Print(ctx context.Context, v ...interface{}) {

}

// Printf 打印
func (logger *DefaultLogger) Printf(ctx context.Context, format string, v ...interface{}) {

}

// Debug 调试
func (logger *DefaultLogger) Debug(ctx context.Context, v ...interface{}) {

}

// Debugf 调试
func (logger *DefaultLogger) Debugf(ctx context.Context, format string, v ...interface{}) {

}

// Info 信息
func (logger *DefaultLogger) Info(ctx context.Context, v ...interface{}) {

}

// Infof 信息
func (logger *DefaultLogger) Infof(ctx context.Context, format string, v ...interface{}) {

}

// Error 错误
func (logger *DefaultLogger) Error(ctx context.Context, v ...interface{}) {

}

// Errorf 错误
func (logger *DefaultLogger) Errorf(ctx context.Context, format string, v ...interface{}) {

}

// Fatal 致命错误
func (logger *DefaultLogger) Fatal(ctx context.Context, v ...interface{}) {

}

// Fatalf 致命错误
func (logger *DefaultLogger) Fatalf(ctx context.Context, format string, v ...interface{}) {

}
