package v1

import (
	"context"
	"errors"
	"net/http"

	"github.com/gopad/gopad-api/pkg/middleware/current"
	"github.com/gopad/gopad-api/pkg/model"
	"github.com/gopad/gopad-api/pkg/service/members"
	"github.com/gopad/gopad-api/pkg/service/users"
	"github.com/gopad/gopad-api/pkg/validate"
)

// ListUsers implements the v1.ServerInterface.
func (a *API) ListUsers(ctx context.Context, request ListUsersRequestObject) (ListUsersResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return ListUsers403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	records, count, err := a.users.List(
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
		return ListUsers500JSONResponse{
			Message: ToPtr("Failed to load users"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	payload := make([]User, len(records))
	for id, record := range records {
		payload[id] = a.convertUser(record, true)
	}

	return ListUsers200JSONResponse{
		Total: ToPtr(count),
		Users: ToPtr(payload),
	}, nil
}

// ShowUser implements the v1.ServerInterface.
func (a *API) ShowUser(ctx context.Context, request ShowUserRequestObject) (ShowUserResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return ShowUser403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	record, err := a.users.Show(
		ctx,
		request.UserId,
	)

	if err != nil {
		if err == users.ErrNotFound {
			return ShowUser404JSONResponse{
				Message: ToPtr("Failed to find user"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		return ShowUser500JSONResponse{
			Message: ToPtr("Failed to load user"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	return ShowUser200JSONResponse(
		a.convertUser(record, true),
	), nil
}

// CreateUser implements the v1.ServerInterface.
func (a *API) CreateUser(ctx context.Context, request CreateUserRequestObject) (CreateUserResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return CreateUser403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	record := &model.User{}

	if request.Body.Username != nil {
		record.Username = FromPtr(request.Body.Username)
	}

	if request.Body.Password != nil {
		record.Password = FromPtr(request.Body.Password)
	}

	if request.Body.Email != nil {
		record.Email = FromPtr(request.Body.Email)
	}

	if request.Body.Fullname != nil {
		record.Fullname = FromPtr(request.Body.Fullname)
	}

	if request.Body.Admin != nil {
		record.Admin = FromPtr(request.Body.Admin)
	}

	if request.Body.Active != nil {
		record.Active = FromPtr(request.Body.Active)
	}

	if err := a.users.Create(
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

			return CreateUser422JSONResponse{
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Message: ToPtr("Failed to validate user"),
				Errors:  ToPtr(errors),
			}, nil
		}

		return CreateUser500JSONResponse{
			Message: ToPtr("Failed to create user"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	return CreateUser200JSONResponse(
		a.convertUser(record, false),
	), nil
}

// UpdateUser implements the v1.ServerInterface.
func (a *API) UpdateUser(ctx context.Context, request UpdateUserRequestObject) (UpdateUserResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return UpdateUser403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	record, err := a.users.Show(
		ctx,
		request.UserId,
	)

	if err != nil {
		if err == users.ErrNotFound {
			return UpdateUser404JSONResponse{
				Message: ToPtr("Failed to find user"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		return UpdateUser500JSONResponse{
			Message: ToPtr("Failed to load user"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	if request.Body.Username != nil {
		record.Username = FromPtr(request.Body.Username)
	}

	if request.Body.Password != nil {
		record.Password = FromPtr(request.Body.Password)
	}

	if request.Body.Email != nil {
		record.Email = FromPtr(request.Body.Email)
	}

	if request.Body.Fullname != nil {
		record.Fullname = FromPtr(request.Body.Fullname)
	}

	if request.Body.Admin != nil {
		record.Admin = FromPtr(request.Body.Admin)
	}

	if request.Body.Active != nil {
		record.Active = FromPtr(request.Body.Active)
	}

	if err := a.users.Update(
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

			return UpdateUser422JSONResponse{
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Message: ToPtr("Failed to validate user"),
				Errors:  ToPtr(errors),
			}, nil
		}

		return UpdateUser500JSONResponse{
			Message: ToPtr("Failed to update user"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	return UpdateUser200JSONResponse(
		a.convertUser(record, false),
	), nil
}

// DeleteUser implements the v1.ServerInterface.
func (a *API) DeleteUser(ctx context.Context, request DeleteUserRequestObject) (DeleteUserResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return DeleteUser403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	record, err := a.users.Show(
		ctx,
		request.UserId,
	)

	if err != nil {
		if err == users.ErrNotFound {
			return DeleteUser404JSONResponse{
				Message: ToPtr("Failed to find user"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		return DeleteUser500JSONResponse{
			Message: ToPtr("Failed to load user"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	if err := a.users.Delete(
		ctx,
		record.ID,
	); err != nil {
		return DeleteUser400JSONResponse{
			Status:  ToPtr(http.StatusBadRequest),
			Message: ToPtr("Failed to delete user"),
		}, nil
	}

	return DeleteUser200JSONResponse{
		Status:  ToPtr(http.StatusOK),
		Message: ToPtr("Successfully deleted user"),
	}, nil
}

// ListUserTeams implements the v1.ServerInterface.
func (a *API) ListUserTeams(ctx context.Context, request ListUserTeamsRequestObject) (ListUserTeamsResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return ListUserTeams403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	record, err := a.users.Show(
		ctx,
		request.UserId,
	)

	if err != nil {
		if err == users.ErrNotFound {
			return ListUserTeams404JSONResponse{
				Message: ToPtr("Failed to find user"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		return ListUserTeams500JSONResponse{
			Message: ToPtr("Failed to load user"),
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
			UserID: record.ID,
		},
	)

	if err != nil {
		return ListUserTeams500JSONResponse{
			Message: ToPtr("Failed to load members"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	payload := make([]UserTeam, len(records))
	for id, record := range records {
		payload[id] = a.convertUserTeam(record)
	}

	return ListUserTeams200JSONResponse{
		Total: ToPtr(count),
		User:  ToPtr(a.convertUser(record, false)),
		Teams: ToPtr(payload),
	}, nil
}

// AttachUserToTeam implements the v1.ServerInterface.
func (a *API) AttachUserToTeam(ctx context.Context, request AttachUserToTeamRequestObject) (AttachUserToTeamResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return AttachUserToTeam403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	if err := a.members.Attach(
		ctx,
		model.MemberParams{
			UserID: request.UserId,
			TeamID: request.Body.Team,
			Perm:   string(FromPtr(request.Body.Perm)),
		},
	); err != nil {
		if errors.Is(err, members.ErrNotFound) {
			return AttachUserToTeam404JSONResponse{
				Message: ToPtr("Failed to find user or team"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		if errors.Is(err, members.ErrAlreadyAssigned) {
			return AttachUserToTeam412JSONResponse{
				Message: ToPtr("Team is already attached"),
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

			return AttachUserToTeam422JSONResponse{
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Message: ToPtr("Failed to validate user team"),
				Errors:  ToPtr(errors),
			}, nil
		}

		return AttachUserToTeam500JSONResponse{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to attach user to team"),
		}, nil
	}

	return AttachUserToTeam200JSONResponse{
		Message: ToPtr("Successfully attached user to team"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

// PermitUserTeam implements the v1.ServerInterface.
func (a *API) PermitUserTeam(ctx context.Context, request PermitUserTeamRequestObject) (PermitUserTeamResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return PermitUserTeam403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	if err := a.members.Permit(
		ctx,
		model.MemberParams{
			UserID: request.UserId,
			TeamID: request.Body.Team,
			Perm:   string(FromPtr(request.Body.Perm)),
		},
	); err != nil {
		if errors.Is(err, members.ErrNotFound) {
			return PermitUserTeam404JSONResponse{
				Message: ToPtr("Failed to find user or team"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		if errors.Is(err, members.ErrNotAssigned) {
			return PermitUserTeam412JSONResponse{
				Message: ToPtr("Team is not attached"),
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

			return PermitUserTeam422JSONResponse{
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Message: ToPtr("Failed to validate user team"),
				Errors:  ToPtr(errors),
			}, nil
		}

		return PermitUserTeam500JSONResponse{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to update user team perms"),
		}, nil
	}

	return PermitUserTeam200JSONResponse{
		Message: ToPtr("Successfully updated user team perms"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

// DeleteUserFromTeam implements the v1.ServerInterface.
func (a *API) DeleteUserFromTeam(ctx context.Context, request DeleteUserFromTeamRequestObject) (DeleteUserFromTeamResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return DeleteUserFromTeam403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	if err := a.members.Drop(
		ctx,
		model.MemberParams{
			UserID: request.UserId,
			TeamID: request.Body.Team,
		},
	); err != nil {
		if errors.Is(err, members.ErrNotFound) {
			return DeleteUserFromTeam404JSONResponse{
				Message: ToPtr("Failed to find user or team"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			}, nil
		}

		if errors.Is(err, members.ErrNotAssigned) {
			return DeleteUserFromTeam412JSONResponse{
				Message: ToPtr("Team is not attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			}, nil
		}

		return DeleteUserFromTeam500JSONResponse{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to drop user from team"),
		}, nil
	}

	return DeleteUserFromTeam200JSONResponse{
		Message: ToPtr("Successfully dropped user from team"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

func (a *API) convertUser(record *model.User, includeRefs bool) User {
	result := User{
		Id:        ToPtr(record.ID),
		Username:  ToPtr(record.Username),
		Email:     ToPtr(record.Email),
		Fullname:  ToPtr(record.Fullname),
		Profile:   ToPtr(gravatarFor(record.Email)),
		Active:    ToPtr(record.Active),
		Admin:     ToPtr(record.Admin),
		CreatedAt: ToPtr(record.CreatedAt),
		UpdatedAt: ToPtr(record.UpdatedAt),
	}

	if includeRefs {
		auths := make([]UserAuth, 0)

		for _, auth := range record.Auths {
			auths = append(
				auths,
				a.convertUserAuth(auth),
			)
		}

		result.Auths = ToPtr(auths)

		teams := make([]UserTeam, 0)

		for _, team := range record.Teams {
			teams = append(
				teams,
				a.convertUserTeam(team),
			)
		}

		result.Teams = ToPtr(teams)
	}

	return result
}

func (a *API) convertUserAuth(record *model.UserAuth) UserAuth {
	result := UserAuth{
		Provider:  ToPtr(record.Provider),
		Ref:       ToPtr(record.Ref),
		CreatedAt: ToPtr(record.CreatedAt),
		UpdatedAt: ToPtr(record.UpdatedAt),
	}

	return result
}

func (a *API) convertUserTeam(record *model.Member) UserTeam {
	result := UserTeam{
		TeamId:    record.TeamID,
		Team:      ToPtr(a.convertTeam(record.Team, false)),
		UserId:    record.UserID,
		Perm:      ToPtr(UserTeamPerm(record.Perm)),
		CreatedAt: ToPtr(record.CreatedAt),
		UpdatedAt: ToPtr(record.UpdatedAt),
	}

	return result
}
