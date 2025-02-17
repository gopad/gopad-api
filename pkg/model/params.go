package model

// ListParams defines optional list attributes.
type ListParams struct {
	Search string
	Sort   string
	Order  string
	Limit  int64
	Offset int64
}

// UserGroupParams defines parameters for user groups.
type UserGroupParams struct {
	ListParams

	UserID  string
	GroupID string
	Perm    string
}
