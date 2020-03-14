package user

type UserQuery struct {
}

// UserQueryHandler ...
func UserQueryHandler() *UserQuery {
	return &UserQuery{}
}

// UserQueryInterface ...
type UserQueryInterface interface{}
