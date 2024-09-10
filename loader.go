package config

import (
	"github.com/pangum/config/internal/runtime"
)

// Loader 加载器
type Loader interface {
	// Supports 支持的类型
	Supports() []string

	// Load 加载配置时调用
	Load(config runtime.Pointer) (err error)
}
