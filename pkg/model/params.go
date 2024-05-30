package model

// ListParams defines optional list attributes.
type ListParams struct {
	Search string
	Sort   string
	Order  string
	Limit  int
	Offset int
}

// MemberParams defines parameters for members.
type MemberParams struct {
	ListParams

	UserID string
	TeamID string
	Perm   string
}
