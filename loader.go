package config

import (
	"context"

	"github.com/pangum/config/internal/runtime"
)

// Loader 加载器
type Loader interface {
	// Load 加载配置时调用
	Load(ctx context.Context, config runtime.Pointer) (err error)
}
