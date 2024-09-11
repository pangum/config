package config

import (
	"context"

	"github.com/pangum/config/internal/runtime"
)

// Loader 加载器
type Loader interface {
	// Local 是否是本地加载器
	Local() bool

	// Load 加载配置时调用
	Load(ctx context.Context, config runtime.Pointer) (loaded bool, err error)
}
