package model

type Role int

const (
	MEMBER Role = iota
	ADMIN
	SUPERADMIN
)

// this is literture
// MEMBER = 0 , ADMIN = 1, SUPERADMIN = 2

func (r Role) StringRole() string {
	return []string{"MEMBER", "ADMIN", "SUPERADMIN"}[r]
}
