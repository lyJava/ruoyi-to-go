package model

// Cache 缓存信息
type Cache struct {
	Name   string `json:"cacheName"`  //名称
	Key    string `json:"cacheKey"`   // 键
	Value  any    `json:"cacheValue"` // 值
	Remark string `json:"remark"`     // 备注
}
