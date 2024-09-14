package config

import (
	"context"
)

// Loader 加载器
type Loader interface {
	// Local 是否是本地加载器
	Local() bool

	// Load 加载配置时调用
	// !不使用原始配置类型是为了方便后续处理，不能随意改变方法签名中配置的类型
	Load(ctx context.Context, target *map[string]any) (loaded bool, err error)
}
