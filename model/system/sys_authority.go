package system

import (
	"github.com/georgezhangjiacheng/my-shop-common/global"
)

/*
 *@Author georgezhangjc@163.com
 *@Date 2022/1/25 上午11:08
 */

type SysAuthority struct {
	global.BaseModel
	AuthorityId     string         `json:"authorityId" gorm:"not null;unique;primary_key;comment:角色ID;size:90"` // 角色ID
	AuthorityName   string         `json:"authorityName" gorm:"comment:角色名"`                                    // 角色名
	ParentId        string         `json:"parentId" gorm:"comment:父角色ID"`                                       // 父角色ID
	//DataAuthorityId []SysAuthority `json:"dataAuthorityId" gorm:"many2many:sys_data_authority_id"`
	//Children        []SysAuthority `json:"children" gorm:"-"`
	//SysBaseMenus    []SysBaseMenu  `json:"menus" gorm:"many2many:sys_authority_menus;"`
}
