package config

import (
	"github.com/pangum/starter/internal/runtime"
)

// Loader 加载器
type Loader interface {
	// Load 加载配置时调用
	Load(config runtime.Pointer) (err error)
}
