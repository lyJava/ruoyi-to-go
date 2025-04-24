package domain

// DeptLeader 部门负责人
type DeptLeader struct {
	Id           int64         `json:"id,string,omitempty"`                                             // 主键ID
	DeptId       int64         `json:"deptId,string" validate:"required,max=9999,min=100" label:"部门ID"` // 部门主键ID
	Username     string        `json:"username,omitempty"`                                              // 姓名
	Phone        string        `json:"phone,omitempty" validate:"phone"`                                // 手机
	Email        string        `json:"email,omitempty" validate:"omitempty,email" label:""`             // 邮箱
	Address      string        `json:"address,omitempty"`                                               // 地址
	EnableStatus string        `json:"enableStatus,omitempty" validate:"oneof=0 1"`                     // 是否启用（0:启用；1:禁用）
	CreateBy     string        `json:"createBy,omitempty"`                                              // 创建人
	CreateTime   LocalDateTime `json:"createTime,omitempty"`                                            // 创建时间
	UpdateBy     string        `json:"updateBy,omitempty"`                                              // 修改人
	UpdateTime   LocalDateTime `json:"updateTime,omitempty"`                                            // 更新时间
	Remarks      string        `json:"remarks,omitempty"`                                               // 备注
	DelFlag      string        `json:"delFlag,omitempty" validate:"oneof=0 1"`                          // 是否删除(0:正常；1:删除)
}

// TableName gorm指定表名
func (DeptLeader) TableName() string {
	return "sys_dept_leader"
}

// Validator 实现 gin 的 binding.Validator 引擎
/* type Validator struct {
	validate *validator.Validate
}

// ValidateStruct 使用 validator 验证结构体
func (v *Validator) ValidateStruct(obj interface{}) error {
	if kindOfData := reflect.ValueOf(obj).Kind(); kindOfData == reflect.Struct {
		if err := v.validate.Struct(obj); err != nil {
			return err
		}
	}
	return nil
}

// Engine 返回 validator 实例
func (v *Validator) Engine() interface{} {
	return v.validate
}

func NewValidator() *Validator {
	return &Validator{validate: validator.New()}
}
*/
