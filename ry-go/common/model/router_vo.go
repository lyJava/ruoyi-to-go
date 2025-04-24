package model

// RouterVo 路由信息结构
type RouterVo struct {
	Name       string      `json:"name,omitempty"`       // 路由名称
	Path       string      `json:"path,omitempty"`       // 路由地址
	Hidden     bool        `json:"hidden,omitempty"`     // 是否隐藏路由，当设置 true 的时候该路由不会再侧边栏出现
	Redirect   string      `json:"redirect,omitempty"`   // 重定向地址，当设置 noRedirect 的时候该路由在面包屑导航中不可被点击
	Component  string      `json:"component,omitempty"`  // 组件地址
	AlwaysShow bool        `json:"alwaysShow,omitempty"` // 当你一个路由下面的 children 声明的路由大于1个时，自动会变成嵌套的模式--如组件页面
	Meta       *MetaVo     `json:"meta,omitempty"`
	Children   []*RouterVo `json:"children,omitempty"` // 子路由
}

// MetaVo 路由页面信息结构
type MetaVo struct {
	Title   string `json:"title"`   // 设置该路由在侧边栏和面包屑中展示的名字
	Icon    string `json:"icon"`    // 设置该路由的图标，对应路径src/assets/icons/svg
	NoCache bool   `json:"noCache"` // 设置为true，则不会被 <keep-alive>缓存
}

func NewMetaVo(menuName, menuIcon string, isCache bool) *MetaVo {
	return &MetaVo{
		Title:   menuName,
		Icon:    menuIcon,
		NoCache: isCache,
	}
}
