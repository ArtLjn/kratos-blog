package data

const (
	Admin   = "admin"
	User    = "user"
	Visitor = "visitor"
)

type RoleDesc struct {
	role        string
	description string
}

func (desc *RoleDesc) Role() string {
	return desc.role
}

func (desc *RoleDesc) Description() string {
	return desc.description
}

type RoleItem struct {
	RoleDesc
	permission bool
}

func NewRoleItem(role, desc string, permission bool) *RoleItem {
	return &RoleItem{
		RoleDesc: RoleDesc{
			role:        role,
			description: desc,
		},
		permission: permission,
	}
}
