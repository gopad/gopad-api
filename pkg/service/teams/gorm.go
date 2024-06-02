package teams

import (
	"context"
	"errors"
	"fmt"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gopad/gopad-api/pkg/config"
	"github.com/gopad/gopad-api/pkg/model"
	"github.com/gopad/gopad-api/pkg/store"
	"github.com/gopad/gopad-api/pkg/validate"
	"gorm.io/gorm"
)

// GormService defines the service to store content within a database based on Gorm.
type GormService struct {
	handle    *gorm.DB
	config    *config.Config
	principal *model.User
}

// NewGormService initializes the service to store content within a database based on Gorm.
func NewGormService(
	handle *gorm.DB,
	cfg *config.Config,
) *GormService {
	return &GormService{
		handle: handle,
		config: cfg,
	}
}

// WithPrincipal implements the Service interface for database persistence.
func (s *GormService) WithPrincipal(principal *model.User) Service {
	s.principal = principal
	return s
}

// List implements the Service interface for database persistence.
func (s *GormService) List(ctx context.Context, params model.ListParams) ([]*model.Team, int64, error) {
	counter := int64(0)
	records := make([]*model.Team, 0)
	q := s.query(ctx)

	if val, ok := s.validSort(params.Sort); ok {
		q = q.Order(strings.Join(
			[]string{
				val,
				sortOrder(params.Order),
			},
			" ",
		))
	}

	// if params.Search != "" {
	// 	opts := queryparser.Options{
	// 		CutFn: searchCut,
	// 		Allowed: []string{
	// 			"slug",
	// 			"name",
	// 		},
	// 	}

	// 	parser := queryparser.New(
	// 		params.Search,
	// 		opts,
	// 	).Parse()

	// 	for _, name := range opts.Allowed {
	// 		if parser.Has(name) {

	// 			q = q.Where(
	// 				fmt.Sprintf(
	// 					"%s LIKE ?",
	// 					name,
	// 				),
	// 				strings.ReplaceAll(
	// 					parser.GetOne(name),
	// 					"*",
	// 					"%",
	// 				),
	// 			)
	// 		}
	// 	}
	// }

	if err := q.Count(
		&counter,
	).Error; err != nil {
		return nil, counter, err
	}

	if params.Limit > 0 {
		q = q.Limit(params.Limit)
	}

	if params.Offset > 0 {
		q = q.Offset(params.Offset)
	}

	if err := q.Find(
		&records,
	).Error; err != nil {
		return nil, counter, err
	}

	return records, counter, nil
}

// Show implements the Service interface for database persistence.
func (s *GormService) Show(ctx context.Context, name string) (*model.Team, error) {
	record := &model.Team{}

	err := s.query(ctx).Where(
		"id = ?",
		name,
	).Or(
		"slug = ?",
		name,
	).First(
		record,
	).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return record, ErrNotFound
	}

	return record, err
}

// Create implements the Service interface for database persistence.
func (s *GormService) Create(ctx context.Context, team *model.Team) error {
	tx := s.handle.WithContext(
		ctx,
	).Begin()
	defer tx.Rollback()

	if team.Slug == "" {
		team.Slug = store.Slugify(
			tx.Model(&model.Team{}),
			team.Name,
			"",
			false,
		)
	}

	if err := s.validate(ctx, team, false); err != nil {
		return err
	}

	if err := tx.Create(team).Error; err != nil {
		return err
	}

	if err := tx.Create(&model.UserTeam{
		TeamID: team.ID,
		UserID: s.principal.ID,
		Perm:   "owner",
	}).Error; err != nil {
		return err
	}

	return tx.Commit().Error
}

// Update implements the Service interface for database persistence.
func (s *GormService) Update(ctx context.Context, team *model.Team) error {
	tx := s.handle.WithContext(
		ctx,
	).Begin()
	defer tx.Rollback()

	if team.Slug == "" {
		team.Slug = store.Slugify(
			tx.Model(&model.Team{}),
			team.Name,
			"",
			false,
		)
	}

	if err := s.validate(ctx, team, true); err != nil {
		return err
	}

	if err := tx.Save(team).Error; err != nil {
		return err
	}

	return tx.Commit().Error
}

// Delete implements the Service interface for database persistence.
func (s *GormService) Delete(ctx context.Context, name string) error {
	tx := s.handle.WithContext(
		ctx,
	).Begin()
	defer tx.Rollback()

	if err := tx.Where(
		"id = ?",
		name,
	).Or(
		"slug = ?",
		name,
	).Delete(
		&model.Team{},
	).Error; err != nil {
		return err
	}

	return tx.Commit().Error
}

// Exists implements the Service interface for database persistence.
func (s *GormService) Exists(ctx context.Context, name string) (bool, error) {
	res := s.query(ctx).Where(
		"id = ?",
		name,
	).Or(
		"slug = ?",
		name,
	).Find(
		&model.Team{},
	)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return false, nil
	}

	if res.Error != nil {
		return false, res.Error
	}

	return res.RowsAffected > 0, nil
}

func (s *GormService) validate(ctx context.Context, record *model.Team, _ bool) error {
	errs := validate.Errors{}

	if err := validation.Validate(
		record.Slug,
		validation.Required,
		validation.Length(3, 255),
		validation.By(s.uniqueValueIsPresent(ctx, "slug", record.ID)),
	); err != nil {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "slug",
			Error: err,
		})
	}

	if err := validation.Validate(
		record.Name,
		validation.Required,
		validation.Length(3, 255),
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

func (s *GormService) uniqueValueIsPresent(ctx context.Context, key, id string) func(value interface{}) error {
	return func(value interface{}) error {
		val, _ := value.(string)

		q := s.handle.WithContext(
			ctx,
		).Where(
			fmt.Sprintf("%s = ?", key),
			val,
		)

		if id != "" {
			q = q.Not(
				"id = ?",
				id,
			)
		}

		if q.Find(
			&model.Team{},
		).RowsAffected != 0 {
			return errors.New("is already taken")
		}

		return nil
	}
}

func (s *GormService) query(ctx context.Context) *gorm.DB {
	return s.handle.WithContext(
		ctx,
	).Model(
		&model.Team{},
	).Preload(
		"Users",
	).Preload(
		"Users.User",
	)
}

func (s *GormService) validSort(val string) (string, bool) {
	if val == "" {
		return "name", true
	}

	val = strings.ToLower(val)

	for _, name := range []string{
		"slug",
		"name",
	} {
		if val == name {
			return val, true
		}
	}

	return "name", true
}
