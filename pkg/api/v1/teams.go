package v1

import (
	"context"
	"errors"
	"net/http"

	"github.com/gopad/gopad-api/pkg/model"
	"github.com/gopad/gopad-api/pkg/service/members"
	"github.com/gopad/gopad-api/pkg/service/teams"
	"github.com/gopad/gopad-api/pkg/validate"
)

// ListTeams implements the v1.ServerInterface.
func (a *API) ListTeams(ctx context.Context, request ListTeamsRequestObject) (ListTeamsResponseObject, error) {
	records, count, err := a.teams.List(
		ctx,
		toListParams(
			string(FromPtr(request.Params.Sort)),
			string(FromPtr(request.Params.Order)),
			request.Params.Limit,
			request.Params.Offset,
			request.Params.Search,
		),
	)

	if err != nil {
		return ListTeams500JSONResponse{
			Message: ToPtr("Failed to load teams"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	payload := make([]Team, len(records))
	for id, record := range records {
		payload[id] = a.convertTeam(record, false)
	}

	return ListTeams200JSONResponse{
		Total: ToPtr(count),
		Teams: ToPtr(payload),
	}, nil
}

// ShowTeam implements the v1.ServerInterface.
func (a *API) ShowTeam(ctx context.Context, request ShowTeamRequestObject) (ShowTeamResponseObject, error) {
	record, err := a.teams.Show(
		ctx,
		request.TeamId,
	)

	if err != nil {
		if err == teams.ErrNotFound {
			return ShowTeam404JSONResponse{
				Message: ToPtr("Failed to find team"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		return ShowTeam500JSONResponse{
			Message: ToPtr("Failed to load team"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	return ShowTeam200JSONResponse(
		a.convertTeam(record, true),
	), nil
}

// CreateTeam implements the v1.ServerInterface.
func (a *API) CreateTeam(ctx context.Context, request CreateTeamRequestObject) (CreateTeamResponseObject, error) {
	record := &model.Team{}

	if request.Body.Slug != nil {
		record.Slug = FromPtr(request.Body.Slug)
	}

	if request.Body.Name != nil {
		record.Name = FromPtr(request.Body.Name)
	}

	if err := a.teams.Create(
		ctx,
		record,
	); err != nil {
		if v, ok := err.(validate.Errors); ok {
			errors := make([]Validation, 0)

			for _, verr := range v.Errors {
				errors = append(
					errors,
					Validation{
						Field:   ToPtr(verr.Field),
						Message: ToPtr(verr.Error.Error()),
					},
				)
			}

			return CreateTeam422JSONResponse{
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Message: ToPtr("Failed to validate team"),
				Errors:  ToPtr(errors),
			}, nil
		}

		return CreateTeam500JSONResponse{
			Message: ToPtr("Failed to create team"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	return CreateTeam200JSONResponse(
		a.convertTeam(record, false),
	), nil
}

// UpdateTeam implements the v1.ServerInterface.
func (a *API) UpdateTeam(ctx context.Context, request UpdateTeamRequestObject) (UpdateTeamResponseObject, error) {
	record, err := a.teams.Show(
		ctx,
		request.TeamId,
	)

	if err != nil {
		if err == teams.ErrNotFound {
			return UpdateTeam404JSONResponse{
				Message: ToPtr("Failed to find team"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		return UpdateTeam500JSONResponse{
			Message: ToPtr("Failed to load team"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	if request.Body.Slug != nil {
		record.Slug = FromPtr(request.Body.Slug)
	}

	if request.Body.Name != nil {
		record.Name = FromPtr(request.Body.Name)
	}

	if err := a.teams.Update(
		ctx,
		record,
	); err != nil {
		if v, ok := err.(validate.Errors); ok {
			errors := make([]Validation, 0)

			for _, verr := range v.Errors {
				errors = append(
					errors,
					Validation{
						Field:   ToPtr(verr.Field),
						Message: ToPtr(verr.Error.Error()),
					},
				)
			}

			return UpdateTeam422JSONResponse{
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Message: ToPtr("Failed to validate team"),
				Errors:  ToPtr(errors),
			}, nil
		}

		return UpdateTeam500JSONResponse{
			Message: ToPtr("Failed to update team"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	return UpdateTeam200JSONResponse(
		a.convertTeam(record, false),
	), nil
}

// DeleteTeam implements the v1.ServerInterface.
func (a *API) DeleteTeam(ctx context.Context, request DeleteTeamRequestObject) (DeleteTeamResponseObject, error) {
	record, err := a.teams.Show(
		ctx,
		request.TeamId,
	)

	if err != nil {
		if err == teams.ErrNotFound {
			return DeleteTeam404JSONResponse{
				Message: ToPtr("Failed to find team"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		return DeleteTeam500JSONResponse{
			Message: ToPtr("Failed to load team"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	if err := a.teams.Delete(
		ctx,
		record.ID,
	); err != nil {
		return DeleteTeam400JSONResponse{
			Status:  ToPtr(http.StatusBadRequest),
			Message: ToPtr("Failed to delete team"),
		}, nil
	}

	return DeleteTeam200JSONResponse{
		Status:  ToPtr(http.StatusOK),
		Message: ToPtr("Successfully deleted team"),
	}, nil
}

// ListTeamUsers implements the v1.ServerInterface.
func (a *API) ListTeamUsers(ctx context.Context, request ListTeamUsersRequestObject) (ListTeamUsersResponseObject, error) {
	record, err := a.teams.Show(
		ctx,
		request.TeamId,
	)

	if err != nil {
		if err == teams.ErrNotFound {
			return ListTeamUsers404JSONResponse{
				Message: ToPtr("Failed to find team"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		return ListTeamUsers500JSONResponse{
			Message: ToPtr("Failed to load team"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	records, count, err := a.members.List(
		ctx,
		model.MemberParams{
			ListParams: toListParams(
				string(FromPtr(request.Params.Sort)),
				string(FromPtr(request.Params.Order)),
				request.Params.Limit,
				request.Params.Offset,
				request.Params.Search,
			),
			TeamID: record.ID,
		},
	)

	if err != nil {
		return ListTeamUsers500JSONResponse{
			Message: ToPtr("Failed to load members"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	payload := make([]UserTeam, len(records))
	for id, record := range records {
		payload[id] = a.convertTeamUser(record)
	}

	return ListTeamUsers200JSONResponse{
		Total: ToPtr(count),
		Team:  ToPtr(a.convertTeam(record, false)),
		Users: ToPtr(payload),
	}, nil
}

// AttachTeamToUser implements the v1.ServerInterface.
func (a *API) AttachTeamToUser(ctx context.Context, request AttachTeamToUserRequestObject) (AttachTeamToUserResponseObject, error) {
	if err := a.members.Attach(
		ctx,
		model.MemberParams{
			TeamID: request.TeamId,
			UserID: request.Body.User,
			Perm:   string(FromPtr(request.Body.Perm)),
		},
	); err != nil {
		if errors.Is(err, members.ErrNotFound) {
			return AttachTeamToUser404JSONResponse{
				Message: ToPtr("Failed to find team or user"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		if errors.Is(err, members.ErrAlreadyAssigned) {
			return AttachTeamToUser412JSONResponse{
				Message: ToPtr("User is already attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			}, nil
		}

		if v, ok := err.(validate.Errors); ok {
			errors := make([]Validation, 0)

			for _, verr := range v.Errors {
				errors = append(
					errors,
					Validation{
						Field:   ToPtr(verr.Field),
						Message: ToPtr(verr.Error.Error()),
					},
				)
			}

			return AttachTeamToUser422JSONResponse{
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Message: ToPtr("Failed to validate team user"),
				Errors:  ToPtr(errors),
			}, nil
		}

		return AttachTeamToUser500JSONResponse{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to attach team to user"),
		}, nil
	}

	return AttachTeamToUser200JSONResponse{
		Message: ToPtr("Successfully attached team to user"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

// PermitTeamUser implements the v1.ServerInterface.
func (a *API) PermitTeamUser(ctx context.Context, request PermitTeamUserRequestObject) (PermitTeamUserResponseObject, error) {
	if err := a.members.Permit(
		ctx,
		model.MemberParams{
			TeamID: request.TeamId,
			UserID: request.Body.User,
			Perm:   string(FromPtr(request.Body.Perm)),
		},
	); err != nil {
		if errors.Is(err, members.ErrNotFound) {
			return PermitTeamUser404JSONResponse{
				Message: ToPtr("Failed to find team or user"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		if errors.Is(err, members.ErrNotAssigned) {
			return PermitTeamUser412JSONResponse{
				Message: ToPtr("User is not attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			}, nil
		}

		if v, ok := err.(validate.Errors); ok {
			errors := make([]Validation, 0)

			for _, verr := range v.Errors {
				errors = append(
					errors,
					Validation{
						Field:   ToPtr(verr.Field),
						Message: ToPtr(verr.Error.Error()),
					},
				)
			}

			return PermitTeamUser422JSONResponse{
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Message: ToPtr("Failed to validate team user"),
				Errors:  ToPtr(errors),
			}, nil
		}

		return PermitTeamUser500JSONResponse{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to update team user perms"),
		}, nil
	}

	return PermitTeamUser200JSONResponse{
		Message: ToPtr("Successfully updated team user perms"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

// DeleteTeamFromUser implements the v1.ServerInterface.
func (a *API) DeleteTeamFromUser(ctx context.Context, request DeleteTeamFromUserRequestObject) (DeleteTeamFromUserResponseObject, error) {
	if err := a.members.Drop(
		ctx,
		model.MemberParams{
			TeamID: request.TeamId,
			UserID: request.Body.User,
		},
	); err != nil {
		if errors.Is(err, members.ErrNotFound) {
			return DeleteTeamFromUser404JSONResponse{
				Message: ToPtr("Failed to find team or user"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			}, nil
		}

		if errors.Is(err, members.ErrNotAssigned) {
			return DeleteTeamFromUser412JSONResponse{
				Message: ToPtr("User is not attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			}, nil
		}

		return DeleteTeamFromUser500JSONResponse{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to drop team from user"),
		}, nil
	}

	return DeleteTeamFromUser200JSONResponse{
		Message: ToPtr("Successfully dropped team from user"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

func (a *API) convertTeam(record *model.Team, includeRefs bool) Team {
	result := Team{
		Id:        ToPtr(record.ID),
		Slug:      ToPtr(record.Slug),
		Name:      ToPtr(record.Name),
		CreatedAt: ToPtr(record.CreatedAt),
		UpdatedAt: ToPtr(record.UpdatedAt),
	}

	if includeRefs {
		users := make([]UserTeam, 0)

		for _, user := range record.Users {
			users = append(
				users,
				a.convertTeamUser(user),
			)
		}

		result.Users = ToPtr(users)
	}

	return result
}

func (a *API) convertTeamUser(record *model.Member) UserTeam {
	result := UserTeam{
		UserId:    record.UserID,
		User:      ToPtr(a.convertUser(record.User, false)),
		TeamId:    record.TeamID,
		Perm:      ToPtr(UserTeamPerm(record.Perm)),
		CreatedAt: ToPtr(record.CreatedAt),
		UpdatedAt: ToPtr(record.UpdatedAt),
	}

	return result
}
