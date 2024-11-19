// Copyright (c) Abstract Machines
// SPDX-License-Identifier: Apache-2.0

package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/absmach/magistrala/internal/api"
	"github.com/absmach/magistrala/pkg/apiutil"
	"github.com/absmach/magistrala/pkg/errors"
	mggroups "github.com/absmach/magistrala/pkg/groups"
	"github.com/go-chi/chi/v5"
)

func DecodeListGroupsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	pm, err := decodePageMeta(r)
	if err != nil {
		return nil, err
	}

	level, err := apiutil.ReadNumQuery[uint64](r, api.LevelKey, api.DefLevel)
	if err != nil {
		return nil, errors.Wrap(apiutil.ErrValidation, err)
	}

	parentID, err := apiutil.ReadStringQuery(r, api.ParentKey, "")
	if err != nil {
		return nil, errors.Wrap(apiutil.ErrValidation, err)
	}

	tree, err := apiutil.ReadBoolQuery(r, api.TreeKey, false)
	if err != nil {
		return nil, errors.Wrap(apiutil.ErrValidation, err)
	}
	dir, err := apiutil.ReadNumQuery[int64](r, api.DirKey, -1)
	if err != nil {
		return nil, errors.Wrap(apiutil.ErrValidation, err)
	}

	memberKind, err := apiutil.ReadStringQuery(r, api.MemberKindKey, "")
	if err != nil {
		return nil, errors.Wrap(apiutil.ErrValidation, err)
	}

	permission, err := apiutil.ReadStringQuery(r, api.PermissionKey, api.DefPermission)
	if err != nil {
		return nil, errors.Wrap(apiutil.ErrValidation, err)
	}

	listPerms, err := apiutil.ReadBoolQuery(r, api.ListPerms, api.DefListPerms)
	if err != nil {
		return nil, errors.Wrap(apiutil.ErrValidation, err)
	}
	req := listGroupsReq{
		tree:       tree,
		memberKind: memberKind,
		memberID:   chi.URLParam(r, "memberID"),
		Page: mggroups.Page{
			Level:      level,
			ParentID:   parentID,
			Permission: permission,
			PageMeta:   pm,
			Direction:  dir,
			ListPerms:  listPerms,
		},
	}
	return req, nil
}

func DecodeListParentsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	pm, err := decodePageMeta(r)
	if err != nil {
		return nil, err
	}

	level, err := apiutil.ReadNumQuery[uint64](r, api.LevelKey, api.DefLevel)
	if err != nil {
		return nil, errors.Wrap(apiutil.ErrValidation, err)
	}

	tree, err := apiutil.ReadBoolQuery(r, api.TreeKey, false)
	if err != nil {
		return nil, errors.Wrap(apiutil.ErrValidation, err)
	}
	permission, err := apiutil.ReadStringQuery(r, api.PermissionKey, api.DefPermission)
	if err != nil {
		return nil, errors.Wrap(apiutil.ErrValidation, err)
	}

	listPerms, err := apiutil.ReadBoolQuery(r, api.ListPerms, api.DefListPerms)
	if err != nil {
		return nil, errors.Wrap(apiutil.ErrValidation, err)
	}
	req := listGroupsReq{
		tree: tree,
		Page: mggroups.Page{
			Level:      level,
			ParentID:   chi.URLParam(r, "groupID"),
			Permission: permission,
			PageMeta:   pm,
			Direction:  +1,
			ListPerms:  listPerms,
		},
	}
	return req, nil
}

func DecodeListChildrenRequest(_ context.Context, r *http.Request) (interface{}, error) {
	pm, err := decodePageMeta(r)
	if err != nil {
		return nil, err
	}

	level, err := apiutil.ReadNumQuery[uint64](r, api.LevelKey, api.DefLevel)
	if err != nil {
		return nil, errors.Wrap(apiutil.ErrValidation, err)
	}

	tree, err := apiutil.ReadBoolQuery(r, api.TreeKey, false)
	if err != nil {
		return nil, errors.Wrap(apiutil.ErrValidation, err)
	}
	permission, err := apiutil.ReadStringQuery(r, api.PermissionKey, api.DefPermission)
	if err != nil {
		return nil, errors.Wrap(apiutil.ErrValidation, err)
	}

	listPerms, err := apiutil.ReadBoolQuery(r, api.ListPerms, api.DefListPerms)
	if err != nil {
		return nil, errors.Wrap(apiutil.ErrValidation, err)
	}
	req := listGroupsReq{
		tree: tree,
		Page: mggroups.Page{
			Level:      level,
			ParentID:   chi.URLParam(r, "groupID"),
			Permission: permission,
			PageMeta:   pm,
			Direction:  -1,
			ListPerms:  listPerms,
		},
	}
	return req, nil
}

func DecodeGroupCreate(_ context.Context, r *http.Request) (interface{}, error) {
	if !strings.Contains(r.Header.Get("Content-Type"), api.ContentType) {
		return nil, errors.Wrap(apiutil.ErrValidation, apiutil.ErrUnsupportedContentType)
	}
	var g mggroups.Group
	if err := json.NewDecoder(r.Body).Decode(&g); err != nil {
		return nil, errors.Wrap(apiutil.ErrValidation, errors.Wrap(err, errors.ErrMalformedEntity))
	}
	req := createGroupReq{
		Group: g,
	}

	return req, nil
}

func DecodeGroupUpdate(_ context.Context, r *http.Request) (interface{}, error) {
	if !strings.Contains(r.Header.Get("Content-Type"), api.ContentType) {
		return nil, errors.Wrap(apiutil.ErrValidation, apiutil.ErrUnsupportedContentType)
	}
	req := updateGroupReq{
		id: chi.URLParam(r, "groupID"),
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errors.Wrap(apiutil.ErrValidation, errors.Wrap(err, errors.ErrMalformedEntity))
	}
	return req, nil
}

func DecodeGroupRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := groupReq{
		id: chi.URLParam(r, "groupID"),
	}
	return req, nil
}

func DecodeGroupPermsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := groupPermsReq{
		id: chi.URLParam(r, "groupID"),
	}
	return req, nil
}

func DecodeChangeGroupStatus(_ context.Context, r *http.Request) (interface{}, error) {
	req := changeGroupStatusReq{
		id: chi.URLParam(r, "groupID"),
	}
	return req, nil
}

func DecodeAssignMembersRequest(_ context.Context, r *http.Request) (interface{}, error) {
	if !strings.Contains(r.Header.Get("Content-Type"), api.ContentType) {
		return nil, errors.Wrap(apiutil.ErrValidation, apiutil.ErrUnsupportedContentType)
	}
	req := assignReq{
		groupID: chi.URLParam(r, "groupID"),
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errors.Wrap(apiutil.ErrValidation, errors.Wrap(err, errors.ErrMalformedEntity))
	}
	return req, nil
}

func DecodeUnassignMembersRequest(_ context.Context, r *http.Request) (interface{}, error) {
	if !strings.Contains(r.Header.Get("Content-Type"), api.ContentType) {
		return nil, errors.Wrap(apiutil.ErrValidation, apiutil.ErrUnsupportedContentType)
	}
	req := unassignReq{
		groupID: chi.URLParam(r, "groupID"),
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errors.Wrap(apiutil.ErrValidation, errors.Wrap(err, errors.ErrMalformedEntity))
	}
	return req, nil
}

func DecodeListMembersRequest(_ context.Context, r *http.Request) (interface{}, error) {
	memberKind, err := apiutil.ReadStringQuery(r, api.MemberKindKey, "")
	if err != nil {
		return nil, errors.Wrap(apiutil.ErrValidation, err)
	}
	permission, err := apiutil.ReadStringQuery(r, api.PermissionKey, api.DefPermission)
	if err != nil {
		return nil, errors.Wrap(apiutil.ErrValidation, err)
	}
	req := listMembersReq{
		groupID:    chi.URLParam(r, "groupID"),
		permission: permission,
		memberKind: memberKind,
	}
	return req, nil
}

func decodePageMeta(r *http.Request) (mggroups.PageMeta, error) {
	s, err := apiutil.ReadStringQuery(r, api.StatusKey, api.DefGroupStatus)
	if err != nil {
		return mggroups.PageMeta{}, errors.Wrap(apiutil.ErrValidation, err)
	}
	st, err := mggroups.ToStatus(s)
	if err != nil {
		return mggroups.PageMeta{}, errors.Wrap(apiutil.ErrValidation, err)
	}
	offset, err := apiutil.ReadNumQuery[uint64](r, api.OffsetKey, api.DefOffset)
	if err != nil {
		return mggroups.PageMeta{}, errors.Wrap(apiutil.ErrValidation, err)
	}
	limit, err := apiutil.ReadNumQuery[uint64](r, api.LimitKey, api.DefLimit)
	if err != nil {
		return mggroups.PageMeta{}, errors.Wrap(apiutil.ErrValidation, err)
	}
	name, err := apiutil.ReadStringQuery(r, api.NameKey, "")
	if err != nil {
		return mggroups.PageMeta{}, errors.Wrap(apiutil.ErrValidation, err)
	}
	id, err := apiutil.ReadStringQuery(r, api.IDOrder, "")
	if err != nil {
		return mggroups.PageMeta{}, errors.Wrap(apiutil.ErrValidation, err)
	}
	meta, err := apiutil.ReadMetadataQuery(r, api.MetadataKey, nil)
	if err != nil {
		return mggroups.PageMeta{}, errors.Wrap(apiutil.ErrValidation, err)
	}

	ret := mggroups.PageMeta{
		Offset:   offset,
		Limit:    limit,
		Name:     name,
		ID:       id,
		Metadata: meta,
		Status:   st,
	}
	return ret, nil
}
