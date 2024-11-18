// Copyright (c) Abstract Machines
// SPDX-License-Identifier: Apache-2.0

package api

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/absmach/magistrala/internal/api"
	gapi "github.com/absmach/magistrala/internal/groups/api"
	"github.com/absmach/magistrala/pkg/apiutil"
	"github.com/absmach/magistrala/pkg/authn"
	mgauthn "github.com/absmach/magistrala/pkg/authn"
	"github.com/absmach/magistrala/pkg/errors"
	svcerr "github.com/absmach/magistrala/pkg/errors/service"
	"github.com/absmach/magistrala/pkg/groups"
	"github.com/absmach/magistrala/pkg/policies"
	"github.com/go-chi/chi/v5"
	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

// MakeHandler returns a HTTP handler for Groups API endpoints.
func groupsHandler(svc groups.Service, authn mgauthn.Authentication, r *chi.Mux, logger *slog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(apiutil.LoggingErrorEncoder(logger, api.EncodeError)),
	}

	r.Group(func(r chi.Router) {
		r.Use(api.AuthenticateMiddleware(authn, true))

		r.Route("/{domainID}/groups", func(r chi.Router) {
			r.Post("/", otelhttp.NewHandler(kithttp.NewServer(
				gapi.CreateGroupEndpoint(svc, policies.NewGroupKind),
				gapi.DecodeGroupCreate,
				api.EncodeResponse,
				opts...,
			), "create_group").ServeHTTP)

			r.Get("/{groupID}", otelhttp.NewHandler(kithttp.NewServer(
				gapi.ViewGroupEndpoint(svc),
				gapi.DecodeGroupRequest,
				api.EncodeResponse,
				opts...,
			), "view_group").ServeHTTP)

			r.Delete("/{groupID}", otelhttp.NewHandler(kithttp.NewServer(
				gapi.DeleteGroupEndpoint(svc),
				gapi.DecodeGroupRequest,
				api.EncodeResponse,
				opts...,
			), "delete_group").ServeHTTP)

			r.Get("/{groupID}/permissions", otelhttp.NewHandler(kithttp.NewServer(
				gapi.ViewGroupPermsEndpoint(svc),
				gapi.DecodeGroupPermsRequest,
				api.EncodeResponse,
				opts...,
			), "view_group_permissions").ServeHTTP)

			r.Put("/{groupID}", otelhttp.NewHandler(kithttp.NewServer(
				gapi.UpdateGroupEndpoint(svc),
				gapi.DecodeGroupUpdate,
				api.EncodeResponse,
				opts...,
			), "update_group").ServeHTTP)

			r.Get("/", otelhttp.NewHandler(kithttp.NewServer(
				gapi.ListGroupsEndpoint(svc, "groups", "users"),
				gapi.DecodeListGroupsRequest,
				api.EncodeResponse,
				opts...,
			), "list_groups").ServeHTTP)

			r.Get("/{groupID}/children", otelhttp.NewHandler(kithttp.NewServer(
				gapi.ListGroupsEndpoint(svc, "groups", "users"),
				gapi.DecodeListChildrenRequest,
				api.EncodeResponse,
				opts...,
			), "list_children").ServeHTTP)

			r.Get("/{groupID}/parents", otelhttp.NewHandler(kithttp.NewServer(
				gapi.ListGroupsEndpoint(svc, "groups", "users"),
				gapi.DecodeListParentsRequest,
				api.EncodeResponse,
				opts...,
			), "list_parents").ServeHTTP)

			r.Post("/{groupID}/enable", otelhttp.NewHandler(kithttp.NewServer(
				gapi.EnableGroupEndpoint(svc),
				gapi.DecodeChangeGroupStatus,
				api.EncodeResponse,
				opts...,
			), "enable_group").ServeHTTP)

			r.Post("/{groupID}/disable", otelhttp.NewHandler(kithttp.NewServer(
				gapi.DisableGroupEndpoint(svc),
				gapi.DecodeChangeGroupStatus,
				api.EncodeResponse,
				opts...,
			), "disable_group").ServeHTTP)

			r.Post("/{groupID}/users/assign", otelhttp.NewHandler(kithttp.NewServer(
				assignUsersEndpoint(svc),
				decodeAssignUsersRequest,
				api.EncodeResponse,
				opts...,
			), "assign_users").ServeHTTP)

			r.Post("/{groupID}/users/unassign", otelhttp.NewHandler(kithttp.NewServer(
				unassignUsersEndpoint(svc),
				decodeUnassignUsersRequest,
				api.EncodeResponse,
				opts...,
			), "unassign_users").ServeHTTP)

			r.Post("/{groupID}/groups/assign", otelhttp.NewHandler(kithttp.NewServer(
				assignGroupsEndpoint(svc),
				decodeAssignGroupsRequest,
				api.EncodeResponse,
				opts...,
			), "assign_groups").ServeHTTP)

			r.Post("/{groupID}/groups/unassign", otelhttp.NewHandler(kithttp.NewServer(
				unassignGroupsEndpoint(svc),
				decodeUnassignGroupsRequest,
				api.EncodeResponse,
				opts...,
			), "unassign_groups").ServeHTTP)
		})

		// The ideal placeholder name should be {channelID}, but gapi.DecodeListGroupsRequest uses {memberID} as a placeholder for the ID.
		// So here, we are using {memberID} as the placeholder.
		r.Get("/{domainID}/channels/{memberID}/groups", otelhttp.NewHandler(kithttp.NewServer(
			gapi.ListGroupsEndpoint(svc, "groups", "channels"),
			gapi.DecodeListGroupsRequest,
			api.EncodeResponse,
			opts...,
		), "list_groups_by_channel_id").ServeHTTP)

		r.Get("/{domainID}/users/{memberID}/groups", otelhttp.NewHandler(kithttp.NewServer(
			gapi.ListGroupsEndpoint(svc, "groups", "users"),
			gapi.DecodeListGroupsRequest,
			api.EncodeResponse,
			opts...,
		), "list_groups_by_user_id").ServeHTTP)
	})

	return r
}

func decodeAssignUsersRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := assignUsersReq{
		groupID: chi.URLParam(r, "groupID"),
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errors.Wrap(apiutil.ErrValidation, errors.Wrap(err, errors.ErrMalformedEntity))
	}
	return req, nil
}

func decodeUnassignUsersRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := unassignUsersReq{
		groupID: chi.URLParam(r, "groupID"),
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errors.Wrap(apiutil.ErrValidation, errors.Wrap(err, errors.ErrMalformedEntity))
	}
	return req, nil
}

func assignUsersEndpoint(svc groups.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(assignUsersReq)
		if err := req.validate(); err != nil {
			return nil, errors.Wrap(apiutil.ErrValidation, err)
		}

		session, ok := ctx.Value(api.SessionKey).(authn.Session)
		if !ok {
			return nil, svcerr.ErrAuthorization
		}
		if err := svc.Assign(ctx, session, req.groupID, req.Relation, "users", req.UserIDs...); err != nil {
			return nil, err
		}
		return assignUsersRes{}, nil
	}
}

func unassignUsersEndpoint(svc groups.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(unassignUsersReq)
		if err := req.validate(); err != nil {
			return nil, errors.Wrap(apiutil.ErrValidation, err)
		}

		session, ok := ctx.Value(api.SessionKey).(authn.Session)
		if !ok {
			return nil, svcerr.ErrAuthorization
		}

		if err := svc.Unassign(ctx, session, req.groupID, req.Relation, "users", req.UserIDs...); err != nil {
			return nil, err
		}
		return unassignUsersRes{}, nil
	}
}

func decodeAssignGroupsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := assignGroupsReq{
		groupID:  chi.URLParam(r, "groupID"),
		domainID: chi.URLParam(r, "domainID"),
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errors.Wrap(apiutil.ErrValidation, errors.Wrap(err, errors.ErrMalformedEntity))
	}
	return req, nil
}

func decodeUnassignGroupsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := unassignGroupsReq{
		groupID:  chi.URLParam(r, "groupID"),
		domainID: chi.URLParam(r, "domainID"),
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errors.Wrap(apiutil.ErrValidation, errors.Wrap(err, errors.ErrMalformedEntity))
	}
	return req, nil
}

func assignGroupsEndpoint(svc groups.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(assignGroupsReq)
		if err := req.validate(); err != nil {
			return nil, errors.Wrap(apiutil.ErrValidation, err)
		}

		session, ok := ctx.Value(api.SessionKey).(authn.Session)
		if !ok {
			return nil, svcerr.ErrAuthorization
		}
		if err := svc.Assign(ctx, session, req.groupID, policies.ParentGroupRelation, policies.GroupsKind, req.GroupIDs...); err != nil {
			return nil, err
		}
		return assignUsersRes{}, nil
	}
}

func unassignGroupsEndpoint(svc groups.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(unassignGroupsReq)
		if err := req.validate(); err != nil {
			return nil, errors.Wrap(apiutil.ErrValidation, err)
		}

		session, ok := ctx.Value(api.SessionKey).(authn.Session)
		if !ok {
			return nil, svcerr.ErrAuthorization
		}

		if err := svc.Unassign(ctx, session, req.groupID, policies.ParentGroupRelation, policies.GroupsKind, req.GroupIDs...); err != nil {
			return nil, err
		}
		return unassignUsersRes{}, nil
	}
}
