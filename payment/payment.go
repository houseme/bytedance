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

// Package payment payment
package payment

import (
    "context"
    
    "github.com/houseme/bytedance/config"
    "github.com/houseme/bytedance/credential"
    "github.com/houseme/bytedance/payment/account"
    "github.com/houseme/bytedance/payment/refund"
    "github.com/houseme/bytedance/payment/settle"
    "github.com/houseme/bytedance/payment/sync"
    "github.com/houseme/bytedance/payment/trade"
    "github.com/houseme/bytedance/payment/withdraw"
    "github.com/houseme/bytedance/utility/base"
)

// Pay payment
type Pay struct {
    ctxCfg *credential.ContextConfig
}

// NewPay create payment
func NewPay(ctx context.Context, cfg *config.Config) (*Pay, error) {
    if cfg == nil {
        return nil, base.ErrConfigNotFound
    }
    if cfg.ClientKey() == "" || cfg.ClientSecret() == "" {
        return nil, base.ErrConfigKeyValueEmpty("clientKey or clientSecret")
    }
    if cfg.Salt() == "" {
        return nil, base.ErrConfigKeyValueEmpty("salt")
    }
    
    if cfg.Token() == "" {
        return nil, base.ErrConfigKeyValueEmpty("token")
    }
    
    return &Pay{
        ctxCfg: &credential.ContextConfig{
            Config:            cfg,
            AccessTokenHandle: credential.NewDefaultAccessToken(ctx, cfg),
        },
    }, nil
}

// ContextConfig context config
func (p *Pay) ContextConfig() *credential.ContextConfig {
    return p.ctxCfg
}

// Trade payment trade relation
func (p *Pay) Trade() *trade.Trade {
    return trade.NewTrade(p.ContextConfig())
}

// Withdraw cash
func (p *Pay) Withdraw() *withdraw.Withdraw {
    return withdraw.NewWithdraw(p.ContextConfig())
}

// Settle account cash
func (p *Pay) Settle() *settle.Settle {
    return settle.NewSettle(p.ContextConfig())
}

// Refund order cash
func (p *Pay) Refund() *refund.Refund {
    return refund.NewRefund(p.ContextConfig())
}

// Sync order sync to douyin
func (p *Pay) Sync() *sync.Sync {
    return sync.NewSync(p.ContextConfig())
}

// Account merchant accounts
func (p *Pay) Account() *account.Account {
    return account.NewAccount(p.ContextConfig())
}