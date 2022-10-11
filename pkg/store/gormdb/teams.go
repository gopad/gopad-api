package gormdb

import (
	"context"
	"errors"
	"fmt"

	"github.com/Machiel/slugify"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
	"github.com/gopad/gopad-api/pkg/model"
	"github.com/gopad/gopad-api/pkg/service/teams"
	"github.com/gopad/gopad-api/pkg/validate"
	"gorm.io/gorm"
)

// Teams implements teams.Store interface.
type Teams struct {
	client *gormdbStore
}

// List implements List from teams.Store interface.
func (t *Teams) List(ctx context.Context) ([]*model.Team, error) {
	records := make([]*model.Team, 0)

	err := t.client.handle.Order(
		"name ASC",
	).Preload(
		"Users",
	).Preload(
		"Users.Team",
	).Preload(
		"Users.User",
	).Find(
		&records,
	).Error

	return records, err
}

// Show implements Show from teams.Store interface.
func (t *Teams) Show(ctx context.Context, name string) (*model.Team, error) {
	record := &model.Team{}

	err := t.client.handle.Where(
		"id = ?",
		name,
	).Or(
		"slug = ?",
		name,
	).Preload(
		"Users",
	).Preload(
		"Users.Team",
	).Preload(
		"Users.User",
	).First(
		record,
	).Error

	if err == gorm.ErrRecordNotFound {
		return record, teams.ErrNotFound
	}

	return record, err
}

// Create implements Create from teams.Store interface.
func (t *Teams) Create(ctx context.Context, team *model.Team) (*model.Team, error) {
	tx := t.client.handle.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if team.Slug == "" {
		for i := 0; true; i++ {
			if i == 0 {
				team.Slug = slugify.Slugify(team.Name)
			} else {
				team.Slug = slugify.Slugify(
					fmt.Sprintf("%s-%d", team.Name, i),
				)
			}

			if res := tx.Where(
				"slug = ?",
				team.Slug,
			).First(
				&model.Team{},
			); errors.Is(res.Error, gorm.ErrRecordNotFound) {
				break
			}
		}
	}

	team.ID = uuid.New().String()

	fmt.Printf("%+v\n", team)

	if err := t.validateCreate(team); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Create(team).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return team, nil
}

// Update implements Update from teams.Store interface.
func (t *Teams) Update(ctx context.Context, team *model.Team) (*model.Team, error) {
	tx := t.client.handle.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if team.Slug == "" {
		for i := 0; true; i++ {
			if i == 0 {
				team.Slug = slugify.Slugify(team.Name)
			} else {
				team.Slug = slugify.Slugify(
					fmt.Sprintf("%s-%d", team.Name, i),
				)
			}

			if res := tx.Where(
				"slug = ?",
				team.Slug,
			).Not(
				"id",
				team.ID,
			).First(
				&model.Team{},
			); errors.Is(res.Error, gorm.ErrRecordNotFound) {
				break
			}
		}
	}

	if err := t.validateUpdate(team); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Save(team).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return team, nil
}

// Delete implements Delete from teams.Store interface.
func (t *Teams) Delete(ctx context.Context, name string) error {
	tx := t.client.handle.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	record := &model.Team{}

	if err := tx.Where(
		"id = ?",
		name,
	).Or(
		"slug = ?",
		name,
	).First(
		record,
	).Error; err != nil {
		tx.Rollback()

		if err == gorm.ErrRecordNotFound {
			return teams.ErrNotFound
		}

		return err
	}

	if err := tx.Where(
		"team_id = ?",
		record.ID,
	).Delete(
		&model.TeamUser{},
	).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Delete(
		record,
	).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// ListUsers implements ListUsers from teams.Store interface.
func (t *Teams) ListUsers(ctx context.Context, id string) ([]*model.TeamUser, error) {
	records := make([]*model.TeamUser, 0)

	err := t.client.handle.Where(
		"team_id = ?",
		id,
	).Model(
		&model.TeamUser{},
	).Preload(
		"Team",
	).Preload(
		"User",
	).Find(
		&records,
	).Error

	return records, err
}

// AppendUser implements AppendUser from teams.Store interface.
func (t *Teams) AppendUser(ctx context.Context, teamID, userID, perm string) error {
	if t.isAssignedToUser(teamID, userID) {
		return teams.ErrAlreadyAssigned
	}

	tx := t.client.handle.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	record := &model.TeamUser{
		TeamID: teamID,
		UserID: userID,
		Perm:   perm,
	}

	if err := t.validatePerm(record.Perm); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(record).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// PermitUser implements PermitUser from teams.Store interface.
func (t *Teams) PermitUser(ctx context.Context, teamID, userID, perm string) error {
	if t.isUnassignedFromUser(teamID, userID) {
		return teams.ErrNotAssigned
	}

	tx := t.client.handle.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	record := &model.TeamUser{}
	record.Perm = perm

	if err := t.validatePerm(record.Perm); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where(
		"team_id = ? AND user_id = ?",
		teamID,
		userID,
	).Model(
		&model.TeamUser{},
	).Updates(
		record,
	).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// DropUser implements DropUser from teams.Store interface.
func (t *Teams) DropUser(ctx context.Context, teamID, userID string) error {
	if t.isUnassignedFromUser(teamID, userID) {
		return teams.ErrNotAssigned
	}

	tx := t.client.handle.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Where(
		"team_id = ? AND user_id = ?",
		teamID,
		userID,
	).Delete(
		&model.TeamUser{},
	).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (t *Teams) isAssignedToUser(teamID, userID string) bool {
	res := t.client.handle.Where(
		"team_id = ? AND user_id = ?",
		teamID,
		userID,
	).Find(
		&model.TeamUser{},
	)

	return res.RowsAffected != 0
}

func (t *Teams) isUnassignedFromUser(teamID, userID string) bool {
	res := t.client.handle.Where(
		"team_id = ? AND user_id = ?",
		teamID,
		userID,
	).Find(
		&model.TeamUser{},
	)

	return res.RowsAffected == 0
}

func (t *Teams) validateCreate(record *model.Team) error {
	errs := validate.Errors{}

	if err := validation.Validate(
		record.Slug,
		validation.Length(3, 255),
		validation.By(t.uniqueValueIsPresent("slug", record.ID)),
	); err != nil {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "slug",
			Error: err,
		})
	}

	if err := validation.Validate(
		record.Name,
		validation.Length(3, 255),
		validation.By(t.uniqueValueIsPresent("name", record.ID)),
	); err != nil {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "name",
			Error: err,
		})
	}

	if len(errs.Errors) > 0 {
		return errs
	}

	return nil
}

func (t *Teams) validateUpdate(record *model.Team) error {
	errs := validate.Errors{}

	if err := validation.Validate(
		record.ID,
		validation.Required,
		is.UUIDv4,
		validation.By(t.uniqueValueIsPresent("id", record.ID)),
	); err != nil {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "id",
			Error: err,
		})
	}

	if err := validation.Validate(
		record.Slug,
		validation.Length(3, 255),
		validation.By(t.uniqueValueIsPresent("slug", record.ID)),
	); err != nil {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "slug",
			Error: err,
		})
	}

	if err := validation.Validate(
		record.Name,
		validation.Length(3, 255),
		validation.By(t.uniqueValueIsPresent("name", record.ID)),
	); err != nil {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "name",
			Error: err,
		})
	}

	if len(errs.Errors) > 0 {
		return errs
	}

	return nil
}

func (t *Teams) validatePerm(perm string) error {
	if err := validation.Validate(
		perm,
		validation.In("user", "admin", "owner"),
	); err != nil {
		return validate.Errors{
			Errors: []validate.Error{
				{
					Field: "perm",
					Error: fmt.Errorf("invalid permission value"),
				},
			},
		}
	}

	return nil
}

func (t *Teams) uniqueValueIsPresent(key, id string) func(value interface{}) error {
	return func(value interface{}) error {
		val, _ := value.(string)

		res := t.client.handle.Where(
			fmt.Sprintf("%s = ?", key),
			val,
		).Not(
			"id = ?",
			id,
		).Find(
			&model.Team{},
		)

		if res.RowsAffected != 0 {
			return errors.New("is already taken")
		}

		return nil
	}
}
