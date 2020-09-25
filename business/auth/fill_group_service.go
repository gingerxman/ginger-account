package auth

import (
	"context"
	"github.com/gingerxman/eel"
	m_auth "github.com/gingerxman/ginger-account/models/corp"
)

type FillGroupService struct {
	eel.ServiceBase
}

func NewFillGroupService(ctx context.Context) *FillGroupService {
	service := new(FillGroupService)
	service.Ctx = ctx
	return service
}

func (this *FillGroupService) FillOne(group *Group, option eel.FillOption) {
	this.Fill([]*Group{group}, option)
}

func (this *FillGroupService) Fill(groups []*Group, option eel.FillOption)  {
	if len(groups) == 0 {
		return
	}
	
	ids := make([]int, 0)
	for _, group := range groups {
		ids = append(ids, group.Id)
	}
	if enableOption, ok := option["with_permission"]; ok && enableOption {
		this.fillPermission(groups, ids)
	}
}

func (this *FillGroupService) fillPermission(groups []*Group, ids []int) {
	// 搜集permission ids
	var models []*m_auth.GroupHasPermission
	eel.GetOrmFromContext(this.Ctx).Model(&m_auth.GroupHasPermission{}).Where("group_id__in", ids).Find(&models)
	
	permissionIds := make([]int, 0)
	for _, model := range models {
		permissionIds = append(permissionIds, model.PermissionId)
	}
	
	// 获取permissions
	permissions := NewPermissionRepository(this.Ctx).GetPermissionsByIds(permissionIds)
	//构建<id, permission>
	id2permission := make(map[int]*Permission)
	for _, permission := range permissions {
		id2permission[permission.Id] = permission
	}
	
	//构造<id, group>
	id2group := make(map[int]*Group)
	for _, group := range groups {
		id2group[group.Id] = group
	}
	
	// 填充group.Permissions
	for _, model := range models {
		if group, ok := id2group[model.GroupId]; ok {
			if permission, ok2 := id2permission[model.PermissionId]; ok2 {
				group.Permissions = append(group.Permissions, permission)
			}
		}
	}
}

func init() {
}
