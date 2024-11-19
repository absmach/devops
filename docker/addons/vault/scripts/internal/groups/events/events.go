// Copyright (c) Abstract Machines
// SPDX-License-Identifier: Apache-2.0

package events

import (
	"time"

	"github.com/absmach/magistrala/pkg/events"
	groups "github.com/absmach/magistrala/pkg/groups"
)

var (
	groupPrefix          = "group."
	groupCreate          = groupPrefix + "create"
	groupUpdate          = groupPrefix + "update"
	groupChangeStatus    = groupPrefix + "change_status"
	groupView            = groupPrefix + "view"
	groupViewPerms       = groupPrefix + "view_perms"
	groupList            = groupPrefix + "list"
	groupListMemberships = groupPrefix + "list_by_user"
	groupRemove          = groupPrefix + "remove"
	groupAssign          = groupPrefix + "assign"
	groupUnassign        = groupPrefix + "unassign"
)

var (
	_ events.Event = (*assignEvent)(nil)
	_ events.Event = (*unassignEvent)(nil)
	_ events.Event = (*createGroupEvent)(nil)
	_ events.Event = (*updateGroupEvent)(nil)
	_ events.Event = (*changeStatusGroupEvent)(nil)
	_ events.Event = (*viewGroupEvent)(nil)
	_ events.Event = (*deleteGroupEvent)(nil)
	_ events.Event = (*viewGroupEvent)(nil)
	_ events.Event = (*listGroupEvent)(nil)
	_ events.Event = (*listGroupMembershipEvent)(nil)
)

type assignEvent struct {
	memberIDs  []string
	relation   string
	memberKind string
	groupID    string
}

func (cge assignEvent) Encode() (map[string]interface{}, error) {
	return map[string]interface{}{
		"operation":  groupAssign,
		"member_ids": cge.memberIDs,
		"relation":   cge.relation,
		"memberKind": cge.memberKind,
		"group_id":   cge.groupID,
	}, nil
}

type unassignEvent struct {
	memberIDs  []string
	relation   string
	memberKind string
	groupID    string
}

func (cge unassignEvent) Encode() (map[string]interface{}, error) {
	return map[string]interface{}{
		"operation":  groupUnassign,
		"member_ids": cge.memberIDs,
		"relation":   cge.relation,
		"memberKind": cge.memberKind,
		"group_id":   cge.groupID,
	}, nil
}

type createGroupEvent struct {
	groups.Group
}

func (cge createGroupEvent) Encode() (map[string]interface{}, error) {
	val := map[string]interface{}{
		"operation":  groupCreate,
		"id":         cge.ID,
		"status":     cge.Status.String(),
		"created_at": cge.CreatedAt,
	}

	if cge.Domain != "" {
		val["domain"] = cge.Domain
	}
	if cge.Parent != "" {
		val["parent"] = cge.Parent
	}
	if cge.Name != "" {
		val["name"] = cge.Name
	}
	if cge.Description != "" {
		val["description"] = cge.Description
	}
	if cge.Metadata != nil {
		val["metadata"] = cge.Metadata
	}
	if cge.Status.String() != "" {
		val["status"] = cge.Status.String()
	}

	return val, nil
}

type updateGroupEvent struct {
	groups.Group
}

func (uge updateGroupEvent) Encode() (map[string]interface{}, error) {
	val := map[string]interface{}{
		"operation":  groupUpdate,
		"updated_at": uge.UpdatedAt,
		"updated_by": uge.UpdatedBy,
	}

	if uge.ID != "" {
		val["id"] = uge.ID
	}
	if uge.Domain != "" {
		val["domain"] = uge.Domain
	}
	if uge.Parent != "" {
		val["parent"] = uge.Parent
	}
	if uge.Name != "" {
		val["name"] = uge.Name
	}
	if uge.Description != "" {
		val["description"] = uge.Description
	}
	if uge.Metadata != nil {
		val["metadata"] = uge.Metadata
	}
	if !uge.CreatedAt.IsZero() {
		val["created_at"] = uge.CreatedAt
	}
	if uge.Status.String() != "" {
		val["status"] = uge.Status.String()
	}

	return val, nil
}

type changeStatusGroupEvent struct {
	id        string
	status    string
	updatedAt time.Time
	updatedBy string
}

func (rge changeStatusGroupEvent) Encode() (map[string]interface{}, error) {
	return map[string]interface{}{
		"operation":  groupChangeStatus,
		"id":         rge.id,
		"status":     rge.status,
		"updated_at": rge.updatedAt,
		"updated_by": rge.updatedBy,
	}, nil
}

type viewGroupEvent struct {
	groups.Group
}

func (vge viewGroupEvent) Encode() (map[string]interface{}, error) {
	val := map[string]interface{}{
		"operation": groupView,
		"id":        vge.ID,
	}

	if vge.Domain != "" {
		val["domain"] = vge.Domain
	}
	if vge.Parent != "" {
		val["parent"] = vge.Parent
	}
	if vge.Name != "" {
		val["name"] = vge.Name
	}
	if vge.Description != "" {
		val["description"] = vge.Description
	}
	if vge.Metadata != nil {
		val["metadata"] = vge.Metadata
	}
	if !vge.CreatedAt.IsZero() {
		val["created_at"] = vge.CreatedAt
	}
	if !vge.UpdatedAt.IsZero() {
		val["updated_at"] = vge.UpdatedAt
	}
	if vge.UpdatedBy != "" {
		val["updated_by"] = vge.UpdatedBy
	}
	if vge.Status.String() != "" {
		val["status"] = vge.Status.String()
	}

	return val, nil
}

type viewGroupPermsEvent struct {
	permissions []string
}

func (vgpe viewGroupPermsEvent) Encode() (map[string]interface{}, error) {
	return map[string]interface{}{
		"operation":   groupViewPerms,
		"permissions": vgpe.permissions,
	}, nil
}

type listGroupEvent struct {
	groups.Page
}

func (lge listGroupEvent) Encode() (map[string]interface{}, error) {
	val := map[string]interface{}{
		"operation": groupList,
		"total":     lge.Total,
		"offset":    lge.Offset,
		"limit":     lge.Limit,
	}

	if lge.Name != "" {
		val["name"] = lge.Name
	}
	if lge.DomainID != "" {
		val["domain_id"] = lge.DomainID
	}
	if lge.Tag != "" {
		val["tag"] = lge.Tag
	}
	if lge.Metadata != nil {
		val["metadata"] = lge.Metadata
	}
	if lge.Status.String() != "" {
		val["status"] = lge.Status.String()
	}

	return val, nil
}

type listGroupMembershipEvent struct {
	groupID    string
	permission string
	memberKind string
}

func (lgme listGroupMembershipEvent) Encode() (map[string]interface{}, error) {
	return map[string]interface{}{
		"operation":   groupListMemberships,
		"id":          lgme.groupID,
		"permission":  lgme.permission,
		"member_kind": lgme.memberKind,
	}, nil
}

type deleteGroupEvent struct {
	id string
}

func (rge deleteGroupEvent) Encode() (map[string]interface{}, error) {
	return map[string]interface{}{
		"operation": groupRemove,
		"id":        rge.id,
	}, nil
}
