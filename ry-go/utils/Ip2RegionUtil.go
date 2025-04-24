package utils

import (
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
)

// Ip2RegionSearch Ip2Region地址库解析
func Ip2RegionSearch(ip string) (string, error) {
	xdbPath := "resources/ip2region.xdb"
	// 1、从 dbPath 加载 VectorIndex 缓存，把下述 vIndex 变量全局到内存里面。
	vIndex, err := xdb.LoadVectorIndexFromFile(xdbPath)
	if err != nil {
		return "", err
	}

	// 2、用全局的 vIndex 创建带 VectorIndex 缓存的查询对象。
	searcher, err := xdb.NewWithVectorIndex(xdbPath, vIndex)
	if err != nil {
		return "", err
	}

	defer searcher.Close()

	region, err := searcher.SearchByStr(ip)
	if err != nil {
		return "", err
	}

	return region, nil
}
