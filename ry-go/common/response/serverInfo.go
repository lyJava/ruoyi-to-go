package response

// Cpu 处理器信息
type Cpu struct {
	CpuNum int64   `json:"cpuNum"` // 处理器核心
	Total  float64 `json:"total"`  // 处理器总使用率
	Sys    float64 `json:"sys"`    // 系统使用率
	Used   float64 `json:"used"`   // 用户使用率
	Wait   float64 `json:"wait"`   // 等待率
	Free   float64 `json:"free"`   // 空闲率
}

// Memory 内存信息
type Memory struct {
	Total float64 `json:"total"` // 内存总使用率
	Used  float64 `json:"used"`  // 内存使用率
	Free  float64 `json:"free"`  // 内存空闲率
	Usage float64 `json:"usage"` // 使用率
}

// OsInfo 系统信息
type OsInfo struct {
	ComputerName string `json:"computerName"` // 主机名
	Ip           string `json:"computerIp"`   // ip地址
	ProPath      string `json:"userDir"`      // 项目路径
	OsName       string `json:"osName"`       // 系统名称
	OsArch       string `json:"osArch"`       // 系统架构
}

// Gvm go sdk信息
type Gvm struct {
	Total     float64 `json:"total"`     // go总使用率
	Max       float64 `json:"used"`      // golang内存使用率
	Free      float64 `json:"free"`      // 内存空闲率
	Runtime   string  `json:"runTime"`   // 运行时长
	StartTime string  `json:"startTime"` // 启动时间
	Version   string  `json:"version"`   // go sdk 版本
	Home      string  `json:"home"`      // go sdk 路径
	InputArgs string  `json:"inputArgs"` // 运行参数
}

// Disk 磁盘信息
type Disk struct {
	DirName     string  `json:"dirName"`     // 盘符路径
	SysTypeName string  `json:"sysTypeName"` // 盘符类型
	TypeName    string  `json:"typeName"`    // 文件类型
	Total       float64 `json:"total"`       // 总大小
	Free        string  `json:"free"`        // 剩余大小
	Used        float64 `json:"used"`        // 已使用
	Usage       float64 `json:"usage"`       // 已用百分比
}
