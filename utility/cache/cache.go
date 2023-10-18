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

// Package cache universal tool package
package cache

import (
	"context"
	"time"
)

// Cache interface
type Cache interface {
	Get(ctx context.Context, key string) interface{}
	Set(ctx context.Context, key string, val interface{}, timeout time.Duration) error
	IsExist(ctx context.Context, key string) bool
	Delete(ctx context.Context, key string) error
}
