package http

import "git.ecobin.ir/ecomicro/template/app/user/domain"

type CreateUserBody struct {
	Name     string
	Username string
	Password string
	IsAdmin  bool
}
type LoginBody struct {
	Username string
	Password string
}
type UpdateUserBody struct {
	Id       int
	Name     string
	Username string
	Password string
	IsAdmin  bool
}

func (c UpdateUserBody) toDomain() domain.User {
	return domain.User{
		Id:       c.Id,
		Name:     c.Name,
		Username: c.Username,
		Password: c.Password,
		IsAdmin:  c.IsAdmin,
	}
}
func (c LoginBody) toDomain() domain.User {
	return domain.User{
		Username: c.Username,
		Password: c.Password,
	}
}

func (c CreateUserBody) toDomain() domain.User {
	return domain.User{
		Name:     c.Name,
		Username: c.Username,
		Password: c.Password,
		IsAdmin:  c.IsAdmin,
	}
}

type UserIdUri struct {
	Id int `uri:"id"`
}
