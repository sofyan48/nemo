package query

// UserQuery ...
type UserQuery struct {
}

// UserQueryHandler ...
func UserQueryHandler() *UserQuery {
	return &UserQuery{}
}

// UserQueryInterface ...
type UserQueryInterface interface {
	GetByID(uuid string)
}

// GetByID ...
func (userQuery *UserQuery) GetByID(uuid string) {

}
