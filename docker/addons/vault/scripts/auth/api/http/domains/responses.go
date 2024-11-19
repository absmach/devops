// Copyright (c) Abstract Machines
// SPDX-License-Identifier: Apache-2.0

package domains

import (
	"net/http"

	"github.com/absmach/magistrala"
	"github.com/absmach/magistrala/auth"
)

var (
	_ magistrala.Response = (*createDomainRes)(nil)
	_ magistrala.Response = (*retrieveDomainRes)(nil)
	_ magistrala.Response = (*assignUsersRes)(nil)
	_ magistrala.Response = (*unassignUsersRes)(nil)
	_ magistrala.Response = (*listDomainsRes)(nil)
)

type createDomainRes struct {
	auth.Domain
}

func (res createDomainRes) Code() int {
	return http.StatusCreated
}

func (res createDomainRes) Headers() map[string]string {
	return map[string]string{}
}

func (res createDomainRes) Empty() bool {
	return false
}

type retrieveDomainRes struct {
	auth.Domain
}

func (res retrieveDomainRes) Code() int {
	return http.StatusOK
}

func (res retrieveDomainRes) Headers() map[string]string {
	return map[string]string{}
}

func (res retrieveDomainRes) Empty() bool {
	return false
}

type retrieveDomainPermissionsRes struct {
	Permissions []string `json:"permissions"`
}

func (res retrieveDomainPermissionsRes) Code() int {
	return http.StatusOK
}

func (res retrieveDomainPermissionsRes) Headers() map[string]string {
	return map[string]string{}
}

func (res retrieveDomainPermissionsRes) Empty() bool {
	return false
}

type updateDomainRes struct {
	auth.Domain
}

func (res updateDomainRes) Code() int {
	return http.StatusOK
}

func (res updateDomainRes) Headers() map[string]string {
	return map[string]string{}
}

func (res updateDomainRes) Empty() bool {
	return false
}

type listDomainsRes struct {
	auth.DomainsPage
}

func (res listDomainsRes) Code() int {
	return http.StatusOK
}

func (res listDomainsRes) Headers() map[string]string {
	return map[string]string{}
}

func (res listDomainsRes) Empty() bool {
	return false
}

type enableDomainRes struct{}

func (res enableDomainRes) Code() int {
	return http.StatusOK
}

func (res enableDomainRes) Headers() map[string]string {
	return map[string]string{}
}

func (res enableDomainRes) Empty() bool {
	return true
}

type disableDomainRes struct{}

func (res disableDomainRes) Code() int {
	return http.StatusOK
}

func (res disableDomainRes) Headers() map[string]string {
	return map[string]string{}
}

func (res disableDomainRes) Empty() bool {
	return true
}

type freezeDomainRes struct{}

func (res freezeDomainRes) Code() int {
	return http.StatusOK
}

func (res freezeDomainRes) Headers() map[string]string {
	return map[string]string{}
}

func (res freezeDomainRes) Empty() bool {
	return true
}

type assignUsersRes struct{}

func (res assignUsersRes) Code() int {
	return http.StatusCreated
}

func (res assignUsersRes) Headers() map[string]string {
	return map[string]string{}
}

func (res assignUsersRes) Empty() bool {
	return true
}

type unassignUsersRes struct{}

func (res unassignUsersRes) Code() int {
	return http.StatusNoContent
}

func (res unassignUsersRes) Headers() map[string]string {
	return map[string]string{}
}

func (res unassignUsersRes) Empty() bool {
	return true
}

type listUserDomainsRes struct {
	auth.DomainsPage
}

func (res listUserDomainsRes) Code() int {
	return http.StatusOK
}

func (res listUserDomainsRes) Headers() map[string]string {
	return map[string]string{}
}

func (res listUserDomainsRes) Empty() bool {
	return false
}
