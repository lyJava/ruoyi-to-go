package dao

import (
	"context"
	"ry-go/business/domain"
	"ry-go/common/request"
)

// SysDeptDao 系统部门Dao
type SysDeptDao interface {
	Insert(ctx context.Context, dept *domain.Dept) (*domain.Dept, error)
	Update(ctx context.Context, dept *domain.Dept) (*domain.Dept, error)
	SelectById(ctx context.Context, id int64) (*domain.Dept, error)
	CheckDept(ctx context.Context, deptId int64) (bool, error)
	CheckDeptByName(ctx context.Context, deptName string) (bool, error)
	SelectDeptByName(ctx context.Context, deptName string) (*domain.Dept, error)
	SelectDeptById(ctx context.Context, id int64) (*domain.Dept, error)
	CheckDeptByIdAndName(ctx context.Context, id int64, deptName string) (bool, error)
	BatchInsert(ctx context.Context, list []*domain.Dept) ([]int64, error)
	BatchInsertByList(ctx context.Context, list []*request.DeptInsert) ([]int64, error)
	SelectDeptList(ctx context.Context, param *request.DeptPageParam) ([]*domain.Dept, error)
	// SelectPage 分页查询
	//
	// 参数:
	//     param: 包含分页查询的条件参数，类型为 *param.TestUserPageParam，包含以下字段：
	//     - Page: 当前页码
	//     - Size: 每页条数
	//     - Column: 排序的列（可选）
	//     - Order: 排序方式（可选）
	//     - Username: 用户名（可选）
	//     - Email: 邮箱（可选）
	//     - Phone: 电话号码（可选）
	//     - Address: 地址（可选）
	//
	// 返回:¬
	//     -[]*domain.TestUser: 符合条件的用户列表，包含每个用户的详细信息
	//     -int64: 查询结果的总记录数
	//     -int64: 总页数，基于查询的总记录数和每页条数计算
	//     -error: 如果执行过程中出现任何错误，则返回相应的错误信息。如果没有错误，返回nil
	SelectPage(ctx context.Context, param *request.DeptPageParam) ([]*domain.Dept, int64, int64, error)
	BatchDelete(ctx context.Context, ids []any) (int64, error)
	SelectListByDept(ctx context.Context, dept *domain.Dept) ([]*domain.Dept, error)
	BuildDeptTree(list []*domain.Dept) ([]*domain.Dept, error)
	BuildDeptTreeSelect(list []*domain.Dept) []*domain.TreeSelect
	SelectListByRoleId(ctx context.Context, roleId int64) ([]int64, error)
}
