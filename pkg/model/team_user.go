package model

import (
	"time"
)

// TeamUser within Gopad.
type TeamUser struct {
	TeamID    string `storm:"id,index" gorm:"index:idx_id,unique;length:36"`
	Team      *Team
	UserID    string `storm:"id,index" gorm:"index:idx_id,unique;length:36"`
	User      *User
	Perm      string `gorm:"length:32"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
