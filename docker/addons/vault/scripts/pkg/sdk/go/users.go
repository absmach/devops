// Copyright (c) Abstract Machines
// SPDX-License-Identifier: Apache-2.0

package sdk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/absmach/magistrala/pkg/apiutil"
	"github.com/absmach/magistrala/pkg/errors"
)

const (
	usersEndpoint         = "users"
	assignEndpoint        = "assign"
	unassignEndpoint      = "unassign"
	enableEndpoint        = "enable"
	disableEndpoint       = "disable"
	issueTokenEndpoint    = "tokens/issue"
	refreshTokenEndpoint  = "tokens/refresh"
	membersEndpoint       = "members"
	PasswordResetEndpoint = "password"
)

// User represents magistrala user its credentials.
type User struct {
	ID             string      `json:"id"`
	FirstName      string      `json:"first_name,omitempty"`
	LastName       string      `json:"last_name,omitempty"`
	Email          string      `json:"email,omitempty"`
	Credentials    Credentials `json:"credentials"`
	Tags           []string    `json:"tags,omitempty"`
	Metadata       Metadata    `json:"metadata,omitempty"`
	CreatedAt      time.Time   `json:"created_at,omitempty"`
	UpdatedAt      time.Time   `json:"updated_at,omitempty"`
	Status         string      `json:"status,omitempty"`
	Role           string      `json:"role,omitempty"`
	ProfilePicture string      `json:"profile_picture,omitempty"`
}

func (sdk mgSDK) CreateUser(user User, token string) (User, errors.SDKError) {
	data, err := json.Marshal(user)
	if err != nil {
		return User{}, errors.NewSDKError(err)
	}

	url := fmt.Sprintf("%s/%s", sdk.usersURL, usersEndpoint)

	_, body, sdkerr := sdk.processRequest(http.MethodPost, url, token, data, nil, http.StatusCreated)
	if sdkerr != nil {
		return User{}, sdkerr
	}

	user = User{}
	if err := json.Unmarshal(body, &user); err != nil {
		return User{}, errors.NewSDKError(err)
	}

	return user, nil
}

func (sdk mgSDK) Users(pm PageMetadata, token string) (UsersPage, errors.SDKError) {
	url, err := sdk.withQueryParams(sdk.usersURL, usersEndpoint, pm)
	if err != nil {
		return UsersPage{}, errors.NewSDKError(err)
	}

	_, body, sdkerr := sdk.processRequest(http.MethodGet, url, token, nil, nil, http.StatusOK)
	if sdkerr != nil {
		return UsersPage{}, sdkerr
	}

	var cp UsersPage
	if err := json.Unmarshal(body, &cp); err != nil {
		return UsersPage{}, errors.NewSDKError(err)
	}

	return cp, nil
}

func (sdk mgSDK) Members(groupID string, meta PageMetadata, token string) (UsersPage, errors.SDKError) {
	url, err := sdk.withQueryParams(sdk.usersURL, fmt.Sprintf("%s/%s/%s/%s", meta.DomainID, groupsEndpoint, groupID, usersEndpoint), meta)
	if err != nil {
		return UsersPage{}, errors.NewSDKError(err)
	}

	_, body, sdkerr := sdk.processRequest(http.MethodGet, url, token, nil, nil, http.StatusOK)
	if sdkerr != nil {
		return UsersPage{}, sdkerr
	}

	var up UsersPage
	if err := json.Unmarshal(body, &up); err != nil {
		return UsersPage{}, errors.NewSDKError(err)
	}

	return up, nil
}

func (sdk mgSDK) User(id, token string) (User, errors.SDKError) {
	if id == "" {
		return User{}, errors.NewSDKError(apiutil.ErrMissingID)
	}
	url := fmt.Sprintf("%s/%s/%s", sdk.usersURL, usersEndpoint, id)

	_, body, sdkerr := sdk.processRequest(http.MethodGet, url, token, nil, nil, http.StatusOK)
	if sdkerr != nil {
		return User{}, sdkerr
	}

	var user User
	if err := json.Unmarshal(body, &user); err != nil {
		return User{}, errors.NewSDKError(err)
	}

	return user, nil
}

func (sdk mgSDK) UserProfile(token string) (User, errors.SDKError) {
	url := fmt.Sprintf("%s/%s/profile", sdk.usersURL, usersEndpoint)

	_, body, sdkerr := sdk.processRequest(http.MethodGet, url, token, nil, nil, http.StatusOK)
	if sdkerr != nil {
		return User{}, sdkerr
	}

	var user User
	if err := json.Unmarshal(body, &user); err != nil {
		return User{}, errors.NewSDKError(err)
	}

	return user, nil
}

func (sdk mgSDK) UpdateUser(user User, token string) (User, errors.SDKError) {
	if user.ID == "" {
		return User{}, errors.NewSDKError(apiutil.ErrMissingID)
	}
	url := fmt.Sprintf("%s/%s/%s", sdk.usersURL, usersEndpoint, user.ID)

	data, err := json.Marshal(user)
	if err != nil {
		return User{}, errors.NewSDKError(err)
	}

	_, body, sdkerr := sdk.processRequest(http.MethodPatch, url, token, data, nil, http.StatusOK)
	if sdkerr != nil {
		return User{}, sdkerr
	}

	user = User{}
	if err := json.Unmarshal(body, &user); err != nil {
		return User{}, errors.NewSDKError(err)
	}

	return user, nil
}

func (sdk mgSDK) UpdateUserTags(user User, token string) (User, errors.SDKError) {
	data, err := json.Marshal(user)
	if err != nil {
		return User{}, errors.NewSDKError(err)
	}

	url := fmt.Sprintf("%s/%s/%s/tags", sdk.usersURL, usersEndpoint, user.ID)

	_, body, sdkerr := sdk.processRequest(http.MethodPatch, url, token, data, nil, http.StatusOK)
	if sdkerr != nil {
		return User{}, sdkerr
	}

	user = User{}
	if err := json.Unmarshal(body, &user); err != nil {
		return User{}, errors.NewSDKError(err)
	}

	return user, nil
}

func (sdk mgSDK) UpdateUserEmail(user User, token string) (User, errors.SDKError) {
	ucir := updateUserEmailReq{token: token, id: user.ID, Email: user.Email}

	data, err := json.Marshal(ucir)
	if err != nil {
		return User{}, errors.NewSDKError(err)
	}

	url := fmt.Sprintf("%s/%s/%s/email", sdk.usersURL, usersEndpoint, user.ID)

	_, body, sdkerr := sdk.processRequest(http.MethodPatch, url, token, data, nil, http.StatusOK)
	if sdkerr != nil {
		return User{}, sdkerr
	}

	user = User{}
	if err := json.Unmarshal(body, &user); err != nil {
		return User{}, errors.NewSDKError(err)
	}

	return user, nil
}

func (sdk mgSDK) ResetPasswordRequest(email string) errors.SDKError {
	rpr := resetPasswordRequestreq{Email: email}

	data, err := json.Marshal(rpr)
	if err != nil {
		return errors.NewSDKError(err)
	}
	url := fmt.Sprintf("%s/%s/reset-request", sdk.usersURL, PasswordResetEndpoint)

	header := make(map[string]string)
	header["Referer"] = sdk.HostURL

	_, _, sdkerr := sdk.processRequest(http.MethodPost, url, "", data, header, http.StatusCreated)

	return sdkerr
}

func (sdk mgSDK) ResetPassword(password, confPass, token string) errors.SDKError {
	rpr := resetPasswordReq{Token: token, Password: password, ConfPass: confPass}

	data, err := json.Marshal(rpr)
	if err != nil {
		return errors.NewSDKError(err)
	}
	url := fmt.Sprintf("%s/%s/reset", sdk.usersURL, PasswordResetEndpoint)

	_, _, sdkerr := sdk.processRequest(http.MethodPut, url, token, data, nil, http.StatusCreated)

	return sdkerr
}

func (sdk mgSDK) UpdatePassword(oldPass, newPass, token string) (User, errors.SDKError) {
	ucsr := updateUserSecretReq{OldSecret: oldPass, NewSecret: newPass}

	data, err := json.Marshal(ucsr)
	if err != nil {
		return User{}, errors.NewSDKError(err)
	}

	url := fmt.Sprintf("%s/%s/secret", sdk.usersURL, usersEndpoint)

	_, body, sdkerr := sdk.processRequest(http.MethodPatch, url, token, data, nil, http.StatusOK)
	if sdkerr != nil {
		return User{}, sdkerr
	}

	var user User
	if err = json.Unmarshal(body, &user); err != nil {
		return User{}, errors.NewSDKError(err)
	}

	return user, nil
}

func (sdk mgSDK) UpdateUserRole(user User, token string) (User, errors.SDKError) {
	data, err := json.Marshal(user)
	if err != nil {
		return User{}, errors.NewSDKError(err)
	}

	url := fmt.Sprintf("%s/%s/%s/role", sdk.usersURL, usersEndpoint, user.ID)

	_, body, sdkerr := sdk.processRequest(http.MethodPatch, url, token, data, nil, http.StatusOK)
	if sdkerr != nil {
		return User{}, sdkerr
	}

	user = User{}
	if err = json.Unmarshal(body, &user); err != nil {
		return User{}, errors.NewSDKError(err)
	}

	return user, nil
}

func (sdk mgSDK) UpdateUsername(user User, token string) (User, errors.SDKError) {
	uur := UpdateUsernameReq{id: user.ID, Username: user.Credentials.Username}
	data, err := json.Marshal(uur)
	if err != nil {
		return User{}, errors.NewSDKError(err)
	}

	url := fmt.Sprintf("%s/%s/%s/username", sdk.usersURL, usersEndpoint, user.ID)

	_, body, sdkerr := sdk.processRequest(http.MethodPatch, url, token, data, nil, http.StatusOK)
	if sdkerr != nil {
		return User{}, sdkerr
	}

	user = User{}
	if err = json.Unmarshal(body, &user); err != nil {
		return User{}, errors.NewSDKError(err)
	}

	return user, nil
}

func (sdk mgSDK) UpdateProfilePicture(user User, token string) (User, errors.SDKError) {
	data, err := json.Marshal(user)
	if err != nil {
		return User{}, errors.NewSDKError(err)
	}

	url := fmt.Sprintf("%s/%s/%s/picture", sdk.usersURL, usersEndpoint, user.ID)

	_, body, sdkerr := sdk.processRequest(http.MethodPatch, url, token, data, nil, http.StatusOK)
	if sdkerr != nil {
		return User{}, sdkerr
	}

	user = User{}
	if err = json.Unmarshal(body, &user); err != nil {
		return User{}, errors.NewSDKError(err)
	}

	return user, nil
}

func (sdk mgSDK) ListUserChannels(userID string, pm PageMetadata, token string) (ChannelsPage, errors.SDKError) {
	url, err := sdk.withQueryParams(sdk.thingsURL, fmt.Sprintf("%s/%s/%s/%s", pm.DomainID, usersEndpoint, userID, channelsEndpoint), pm)
	if err != nil {
		return ChannelsPage{}, errors.NewSDKError(err)
	}

	_, body, sdkerr := sdk.processRequest(http.MethodGet, url, token, nil, nil, http.StatusOK)
	if sdkerr != nil {
		return ChannelsPage{}, sdkerr
	}
	cp := ChannelsPage{}
	if err := json.Unmarshal(body, &cp); err != nil {
		return ChannelsPage{}, errors.NewSDKError(err)
	}

	return cp, nil
}

func (sdk mgSDK) ListUserGroups(userID string, pm PageMetadata, token string) (GroupsPage, errors.SDKError) {
	url, err := sdk.withQueryParams(sdk.usersURL, fmt.Sprintf("%s/%s/%s/%s", pm.DomainID, usersEndpoint, userID, groupsEndpoint), pm)
	if err != nil {
		return GroupsPage{}, errors.NewSDKError(err)
	}
	_, body, sdkerr := sdk.processRequest(http.MethodGet, url, token, nil, nil, http.StatusOK)
	if sdkerr != nil {
		return GroupsPage{}, sdkerr
	}
	gp := GroupsPage{}
	if err := json.Unmarshal(body, &gp); err != nil {
		return GroupsPage{}, errors.NewSDKError(err)
	}

	return gp, nil
}

func (sdk mgSDK) ListUserThings(userID string, pm PageMetadata, token string) (ThingsPage, errors.SDKError) {
	url, err := sdk.withQueryParams(sdk.thingsURL, fmt.Sprintf("%s/%s/%s/%s", pm.DomainID, usersEndpoint, userID, thingsEndpoint), pm)
	if err != nil {
		return ThingsPage{}, errors.NewSDKError(err)
	}
	_, body, sdkerr := sdk.processRequest(http.MethodGet, url, token, nil, nil, http.StatusOK)
	if sdkerr != nil {
		return ThingsPage{}, sdkerr
	}
	tp := ThingsPage{}
	if err := json.Unmarshal(body, &tp); err != nil {
		return ThingsPage{}, errors.NewSDKError(err)
	}

	return tp, nil
}

func (sdk mgSDK) SearchUsers(pm PageMetadata, token string) (UsersPage, errors.SDKError) {
	url, err := sdk.withQueryParams(sdk.usersURL, fmt.Sprintf("%s/search", usersEndpoint), pm)
	if err != nil {
		return UsersPage{}, errors.NewSDKError(err)
	}

	_, body, sdkerr := sdk.processRequest(http.MethodGet, url, token, nil, nil, http.StatusOK)
	if sdkerr != nil {
		return UsersPage{}, sdkerr
	}

	var cp UsersPage
	if err := json.Unmarshal(body, &cp); err != nil {
		return UsersPage{}, errors.NewSDKError(err)
	}

	return cp, nil
}

func (sdk mgSDK) EnableUser(id, token string) (User, errors.SDKError) {
	return sdk.changeUserStatus(token, id, enableEndpoint)
}

func (sdk mgSDK) DisableUser(id, token string) (User, errors.SDKError) {
	return sdk.changeUserStatus(token, id, disableEndpoint)
}

func (sdk mgSDK) changeUserStatus(token, id, status string) (User, errors.SDKError) {
	url := fmt.Sprintf("%s/%s/%s/%s", sdk.usersURL, usersEndpoint, id, status)

	_, body, sdkerr := sdk.processRequest(http.MethodPost, url, token, nil, nil, http.StatusOK)
	if sdkerr != nil {
		return User{}, sdkerr
	}

	user := User{}
	if err := json.Unmarshal(body, &user); err != nil {
		return User{}, errors.NewSDKError(err)
	}

	return user, nil
}

func (sdk mgSDK) DeleteUser(id, token string) errors.SDKError {
	if id == "" {
		return errors.NewSDKError(apiutil.ErrMissingID)
	}
	url := fmt.Sprintf("%s/%s/%s", sdk.usersURL, usersEndpoint, id)
	_, _, sdkerr := sdk.processRequest(http.MethodDelete, url, token, nil, nil, http.StatusNoContent)
	return sdkerr
}
