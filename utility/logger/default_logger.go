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
    "fmt"
    "log/slog"
    "os"
)

// DefaultLogger 默认日志
type DefaultLogger struct {
    logger *slog.Logger
}

// NewDefaultLogger 实例化
func NewDefaultLogger() *DefaultLogger {
    return &DefaultLogger{
        logger: slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
            AddSource: true,
            Level:     slog.LevelDebug,
        })),
    }
}

// Debug 调试
func (logger *DefaultLogger) Debug(ctx context.Context, v ...any) {
    logger.logger.DebugContext(ctx, "debug level", v...)
}

// Debugf 调试
func (logger *DefaultLogger) Debugf(ctx context.Context, format string, v ...any) {
    logger.logger.DebugContext(ctx, "debug level", fmt.Sprintf(format, v...), "")
}

// Info 信息
func (logger *DefaultLogger) Info(ctx context.Context, v ...any) {
    logger.logger.InfoContext(ctx, "info level", v...)
}

// Infof 信息
func (logger *DefaultLogger) Infof(ctx context.Context, format string, v ...any) {
    logger.logger.InfoContext(ctx, "info level", fmt.Sprintf(format, v...), nil)
}

// Warning 警告
func (logger *DefaultLogger) Warning(ctx context.Context, v ...any) {
    logger.logger.WarnContext(ctx, "warn level", v...)
}

// Warningf 警告
func (logger *DefaultLogger) Warningf(ctx context.Context, format string, v ...any) {
    logger.logger.WarnContext(ctx, "warn level", fmt.Sprintf(format, v...), nil)
}

// Error 错误
func (logger *DefaultLogger) Error(ctx context.Context, v ...any) {
    logger.logger.ErrorContext(ctx, "error level", v...)
}

// Errorf 错误
func (logger *DefaultLogger) Errorf(ctx context.Context, format string, v ...any) {
    logger.logger.ErrorContext(ctx, "error level", fmt.Sprintf(format, v...), nil)
}

// Fatal 致命错误
func (logger *DefaultLogger) Fatal(ctx context.Context, v ...any) {
    logger.logger.ErrorContext(ctx, "fatal level", v...)
    os.Exit(1)
}

// Fatalf 致命错误
func (logger *DefaultLogger) Fatalf(ctx context.Context, format string, v ...any) {
    logger.logger.ErrorContext(ctx, "fatal level", fmt.Sprintf(format, v...), nil)
    os.Exit(1)
}
