package microapp

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
)

// MicroApp mini program
type MicroApp struct {
	opt    options
	logger glog.ILogger
}

type options struct {
	Logger    glog.ILogger
	KeyID     string
	AccessKey string
}

// Option micro app option
type Option func(*options)

// WithLogger set logger
func WithLogger(logger glog.ILogger) Option {
	return func(o *options) {
		o.Logger = logger
	}
}

// WithKeyID set key id
func WithKeyID(keyID string) Option {
	return func(o *options) {
		o.KeyID = keyID
	}
}

// WithAccessKey set access key
func WithAccessKey(accessKey string) Option {
	return func(o *options) {
		o.AccessKey = accessKey
	}
}

// New micro app
func New(ctx context.Context, opts ...Option) *MicroApp {
	op := options{
		Logger: g.Log(),
	}
	for _, option := range opts {
		option(&op)
	}
	return &MicroApp{
		opt:    op,
		logger: op.Logger,
	}
}
