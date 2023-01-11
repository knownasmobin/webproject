package http

import "git.ecobin.ir/ecomicro/template/domain"

type CreateUserBody struct {
	Roles []string `json:"roles"`
	Allow []string `json:"allow"`
	Deny  []string `json:"deny"`
}
type UpdateUserBody struct {
	Id    uint64   `json:"id"`
	Roles []string `json:"roles"`
	Allow []string `json:"allow"`
	Deny  []string `json:"deny"`
}

func (c UpdateUserBody) toDomain() domain.User {
	return domain.User{
		Id:    c.Id,
		Roles: c.Roles,
		Allow: c.Allow,
		Deny:  c.Deny,
	}
}

func (c CreateUserBody) toDomain() domain.User {
	return domain.User{
		Roles: c.Roles,
		Allow: c.Allow,
		Deny:  c.Deny,
	}
}

type UserIdUri struct {
	Id uint64 `uri:"id"`
}
