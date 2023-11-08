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

package helper

import (
    "testing"
)

func TestRandomStr(t *testing.T) {
    type args struct {
        n int
    }
    tests := []struct {
        name string
        args args
        want string
    }{
        {
            name: "TestRandomStr-1",
            args: args{
                n: 2,
            },
            want: "1",
        },
        {
            name: "TestRandomStr-5",
            args: args{
                n: 5,
            },
            want: "1",
        },
        {
            name: "TestRandomStr-10",
            args: args{
                n: 10,
            },
            want: "1",
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := RandomStr(tt.args.n); got == tt.want {
                t.Errorf("RandomStr() = %v, want %v", got, tt.want)
            }
        })
    }
}

// benchmark
func BenchmarkRandomStr(b *testing.B) {
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        RandomStr(6)
    }
}
