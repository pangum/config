package config

// Changer 配置修改
type Changer interface {
	// Changed 当配置发生修改时的回调方法
	Changed(key string, original string, changed string)
}
