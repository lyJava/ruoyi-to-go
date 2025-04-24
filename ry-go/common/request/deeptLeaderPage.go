package request

type DeptLeaderPageParam struct {
	BasePageParam
	Username string `json:"username,omitempty"` // 用户名
	Phone    string `json:"phone,omitempty"`    // 手机
	Email    string `json:"email,omitempty"`    // 邮箱
	Address  string `json:"address,omitempty"`  //地址
}
