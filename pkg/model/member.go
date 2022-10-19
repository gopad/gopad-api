package model

import (
	"time"
)

// Member within Gopad.
type Member struct {
	TeamID    string `storm:"id,index" gorm:"index:idx_id,unique;length:36"`
	Team      *Team
	UserID    string `storm:"id,index" gorm:"index:idx_id,unique;length:36"`
	User      *User
	CreatedAt time.Time
	UpdatedAt time.Time
}
