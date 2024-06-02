package model

// ListParams defines optional list attributes.
type ListParams struct {
	Search string
	Sort   string
	Order  string
	Limit  int
	Offset int
}

// UserTeamParams defines parameters for user teams.
type UserTeamParams struct {
	ListParams

	UserID string
	TeamID string
	Perm   string
}
