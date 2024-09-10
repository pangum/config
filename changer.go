package config

import (
	"github.com/pangum/config/internal/runtime"
)

// Changer 配置修改
type Changer interface {
	// Changed 当配置发生修改时的回调方法
	Changed(original runtime.Pointer, changed runtime.Pointer)
}
