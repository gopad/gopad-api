package store

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var (
	// Migrations define all database migrations.
	Migrations = []*gormigrate.Migration{
		{
			ID: "0001_create_users_table",
			Migrate: func(tx *gorm.DB) error {
				type User struct {
					ID        string `gorm:"primaryKey;length:20"`
					Username  string `gorm:"unique;length:255"`
					Hashword  string `gorm:"length:255"`
					Email     string `gorm:"length:255"`
					Fullname  string `gorm:"length:255"`
					Active    bool   `gorm:"default:false"`
					Admin     bool   `gorm:"default:false"`
					CreatedAt time.Time
					UpdatedAt time.Time
				}

				return tx.Migrator().CreateTable(&User{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("users")
			},
		},
		{
			ID: "0002_create_user_auths_table",
			Migrate: func(tx *gorm.DB) error {
				type UserAuth struct {
					ID        string `gorm:"primaryKey;length:20"`
					UserID    string `gorm:"length:20"`
					Provider  string `gorm:"length:255"`
					Ref       string `gorm:"length:255"`
					CreatedAt time.Time
					UpdatedAt time.Time
				}

				return tx.Migrator().CreateTable(&UserAuth{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("user_auths")
			},
		},
		{
			ID: "0003_create_user_auths_users_constraint",
			Migrate: func(tx *gorm.DB) error {
				type UserAuth struct {
					UserID string `gorm:"length:20"`
				}

				type User struct {
					ID    string      `gorm:"primaryKey"`
					Auths []*UserAuth `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
				}

				return tx.Migrator().CreateConstraint(&User{}, "Auths")
			},
			Rollback: func(tx *gorm.DB) error {
				type UserAuth struct {
					UserID string `gorm:"length:20"`
				}

				type User struct {
					ID    string      `gorm:"primaryKey"`
					Auths []*UserAuth `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
				}

				return tx.Migrator().DropConstraint(&User{}, "Auths")
			},
		},
		{
			ID: "0004_create_teams_table",
			Migrate: func(tx *gorm.DB) error {
				type Team struct {
					ID        string `gorm:"primaryKey;length:20"`
					Slug      string `gorm:"unique;length:255"`
					Name      string `gorm:"unique;length:255"`
					CreatedAt time.Time
					UpdatedAt time.Time
				}

				return tx.Migrator().CreateTable(&Team{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("teams")
			},
		},
		{
			ID: "0005_create_user_teams_table",
			Migrate: func(tx *gorm.DB) error {
				type UserTeam struct {
					UserID    string `gorm:"primaryKey;autoIncrement:false;length:20"`
					TeamID    string `gorm:"primaryKey;autoIncrement:false;length:20"`
					Perm      string `gorm:"length:64"`
					CreatedAt time.Time
					UpdatedAt time.Time
				}

				return tx.Migrator().CreateTable(&UserTeam{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("user_teams")
			},
		},
		{
			ID: "0006_create_user_teams_teams_constraint",
			Migrate: func(tx *gorm.DB) error {
				type UserTeam struct {
					UserID string `gorm:"primaryKey;autoIncrement:false;length:20"`
					TeamID string `gorm:"primaryKey;autoIncrement:false;length:20"`
				}

				type Team struct {
					ID    string      `gorm:"primaryKey"`
					Users []*UserTeam `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
				}

				return tx.Migrator().CreateConstraint(&Team{}, "Users")
			},
			Rollback: func(tx *gorm.DB) error {
				type UserTeam struct {
					UserID string `gorm:"primaryKey;autoIncrement:false;length:20"`
					TeamID string `gorm:"primaryKey;autoIncrement:false;length:20"`
				}

				type Team struct {
					ID    string      `gorm:"primaryKey"`
					Users []*UserTeam `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
				}

				return tx.Migrator().DropConstraint(&Team{}, "Users")
			},
		},
		{
			ID: "0007_create_user_teams_users_constraint",
			Migrate: func(tx *gorm.DB) error {
				type UserTeam struct {
					UserID string `gorm:"primaryKey;autoIncrement:false;length:20"`
					TeamID string `gorm:"primaryKey;autoIncrement:false;length:20"`
				}

				type User struct {
					ID    string      `gorm:"primaryKey"`
					Teams []*UserTeam `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
				}

				return tx.Migrator().CreateConstraint(&User{}, "Teams")
			},
			Rollback: func(tx *gorm.DB) error {
				type UserTeam struct {
					UserID string `gorm:"primaryKey;autoIncrement:false;length:20"`
					TeamID string `gorm:"primaryKey;autoIncrement:false;length:20"`
				}

				type User struct {
					ID    string      `gorm:"primaryKey"`
					Teams []*UserTeam `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
				}

				return tx.Migrator().DropConstraint(&User{}, "Teams")
			},
		},
	}
)
