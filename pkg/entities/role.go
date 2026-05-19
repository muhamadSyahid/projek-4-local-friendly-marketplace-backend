package entities

type Roles string

const (
	RoleBuyer  Roles = "buyer"
	RoleSeller Roles = "seller"
	RoleAdmin  Roles = "admin"
)

func ParseRoles(value string) Roles {
	switch value {
	case "seller":
		return RoleSeller
	case "admin":
		return RoleAdmin
	default:
		return RoleBuyer
	}
}

func (s Roles) String() string {
	return string(s)
}
