// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Cluster is the golang structure of table ck8s_cluster for DAO operations like Where/Data.
type Cluster struct {
	g.Meta    `orm:"table:ck8s_cluster, do:true"`
	Id        interface{} // 主键
	Name      interface{} // 集群名称
	ConfigStr interface{} // 配置文件
	CreatedBy *gtime.Time // 创建用户
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 修改时间
}
