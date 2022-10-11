package boltdb

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Machiel/slugify"
	"github.com/asdine/storm/v3"
	"github.com/asdine/storm/v3/q"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
	"github.com/gopad/gopad-api/pkg/model"
	"github.com/gopad/gopad-api/pkg/service/teams"
	"github.com/gopad/gopad-api/pkg/validate"
)

// Teams implements teams.Store interface.
type Teams struct {
	client *botldbStore
}

// List implements List from teams.Store interface.
func (t *Teams) List(ctx context.Context) ([]*model.Team, error) {
	records := make([]*model.Team, 0)

	err := t.client.handle.AllByIndex(
		"Name",
		&records,
	)

	if err == storm.ErrNotFound {
		return records, nil
	}

	for _, record := range records {
		users, err := t.ListUsers(ctx, record.ID)

		if err != nil {
			return records, err
		}

		record.Users = users
	}

	return records, nil
}

// Show implements Show from teams.Store interface.
func (t *Teams) Show(ctx context.Context, name string) (*model.Team, error) {
	record := &model.Team{}

	if err := t.client.handle.Select(
		q.Or(
			q.Eq("ID", name),
			q.Eq("Slug", name),
		),
	).First(record); err != nil {
		if err == storm.ErrNotFound {
			return record, teams.ErrNotFound
		}

		return nil, err
	}

	users, err := t.ListUsers(ctx, record.ID)

	if err != nil {
		return record, err
	}

	record.Users = users
	return record, nil
}

// Create implements Create from teams.Store interface.
func (t *Teams) Create(ctx context.Context, team *model.Team) (*model.Team, error) {
	tx, err := t.client.handle.Begin(true)

	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	if team.Slug == "" {
		for i := 0; true; i++ {
			if i == 0 {
				team.Slug = slugify.Slugify(team.Name)
			} else {
				team.Slug = slugify.Slugify(
					fmt.Sprintf("%s-%d", team.Name, i),
				)
			}

			if err := tx.Select(
				q.Eq("Slug", team.Slug),
			).First(new(model.Team)); err != nil {
				if err == storm.ErrNotFound {
					break
				}

				return nil, err
			}
		}
	}

	team.ID = uuid.New().String()
	team.UpdatedAt = time.Now().UTC()
	team.CreatedAt = time.Now().UTC()

	if err := t.validateCreate(team); err != nil {
		return nil, err
	}

	if err := tx.Save(team); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return team, nil
}

// Update implements Update from teams.Store interface.
func (t *Teams) Update(ctx context.Context, team *model.Team) (*model.Team, error) {
	tx, err := t.client.handle.Begin(true)

	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	if team.Slug == "" {
		for i := 0; true; i++ {
			if i == 0 {
				team.Slug = slugify.Slugify(team.Name)
			} else {
				team.Slug = slugify.Slugify(
					fmt.Sprintf("%s-%d", team.Name, i),
				)
			}

			if err := tx.Select(
				q.And(
					q.Eq("Slug", team.Slug),
					q.Not(
						q.Eq("ID", team.ID),
					),
				),
			).First(new(model.Team)); err != nil {
				if err == storm.ErrNotFound {
					break
				}

				return nil, err
			}
		}
	}

	team.UpdatedAt = time.Now().UTC()

	if err := t.validateUpdate(team); err != nil {
		return nil, err
	}

	if err := tx.Save(team); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return team, nil
}

// Delete implements Delete from teams.Store interface.
func (t *Teams) Delete(ctx context.Context, name string) error {
	tx, err := t.client.handle.Begin(true)

	if err != nil {
		return err
	}

	defer tx.Rollback()

	record := &model.Team{}
	if err := tx.Select(
		q.Or(
			q.Eq("ID", name),
			q.Eq("Slug", name),
		),
	).First(record); err != nil {
		if err == storm.ErrNotFound {
			return teams.ErrNotFound
		}

		return err
	}

	if err := tx.Select(
		q.Eq("TeamID", record.ID),
	).Delete(new(model.TeamUser)); err != nil {
		return err
	}

	if err := tx.DeleteStruct(record); err != nil {
		return err
	}

	return tx.Commit()
}

// ListUsers implements ListUsers from teams.Store interface.
func (t *Teams) ListUsers(ctx context.Context, id string) ([]*model.TeamUser, error) {
	records := make([]*model.TeamUser, 0)

	if err := t.client.handle.Select(
		q.Eq("TeamID", id),
	).Find(&records); err != nil && err != storm.ErrNotFound {
		return records, err
	}

	for _, record := range records {
		team := &model.Team{}

		if err := t.client.handle.Select(
			q.Eq("ID", record.TeamID),
		).First(team); err != nil && err != storm.ErrNotFound {
			return records, err
		}

		record.Team = team

		user := &model.User{}

		if err := t.client.handle.Select(
			q.Eq("ID", record.UserID),
		).First(user); err != nil && err != storm.ErrNotFound {
			return records, err
		}

		record.User = user
	}

	return records, nil
}

// AppendUser implements AppendUser from teams.Store interface.
func (t *Teams) AppendUser(ctx context.Context, teamID, userID, perm string) error {
	tx, err := t.client.handle.Begin(true)

	if err != nil {
		return err
	}

	defer tx.Rollback()

	if err := t.client.handle.Select(
		q.And(
			q.Eq("TeamID", teamID),
			q.Eq("UserID", userID),
		),
	).First(new(model.TeamUser)); err == nil {
		return teams.ErrAlreadyAssigned
	}

	record := &model.TeamUser{
		TeamID:    teamID,
		UserID:    userID,
		Perm:      perm,
		UpdatedAt: time.Now().UTC(),
		CreatedAt: time.Now().UTC(),
	}

	if err := t.validatePerm(record.Perm); err != nil {
		return err
	}

	if err := tx.Save(record); err != nil {
		return err
	}

	return tx.Commit()
}

// PermitUser implements PermitUser from teams.Store interface.
func (t *Teams) PermitUser(ctx context.Context, teamID, userID, perm string) error {
	tx, err := t.client.handle.Begin(true)

	if err != nil {
		return err
	}

	defer tx.Rollback()
	record := &model.TeamUser{}

	if err := t.client.handle.Select(
		q.And(
			q.Eq("TeamID", teamID),
			q.Eq("UserID", userID),
		),
	).First(record); err == storm.ErrNotFound {
		return teams.ErrNotAssigned
	}

	record.Perm = perm
	record.UpdatedAt = time.Now().UTC()

	if err := t.validatePerm(record.Perm); err != nil {
		return err
	}

	if err := tx.Save(record); err != nil {
		return err
	}

	return tx.Commit()
}

// DropUser implements DropUser from teams.Store interface.
func (t *Teams) DropUser(ctx context.Context, teamID, userID string) error {
	tx, err := t.client.handle.Begin(true)

	if err != nil {
		return err
	}

	defer tx.Rollback()
	record := &model.TeamUser{}

	if err := t.client.handle.Select(
		q.And(
			q.Eq("TeamID", teamID),
			q.Eq("UserID", userID),
		),
	).First(record); err == storm.ErrNotFound {
		return teams.ErrNotAssigned
	}

	if err := tx.DeleteStruct(record); err != nil {
		return err
	}

	return tx.Commit()
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

		if err := t.client.handle.Select(
			q.And(
				q.Eq(key, val),
				q.Not(
					q.Eq("ID", id),
				),
			),
		).First(new(model.Team)); err == storm.ErrNotFound {
			return nil
		}

		return errors.New("taken")
	}
}
