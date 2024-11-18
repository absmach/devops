// Copyright (c) Abstract Machines
// SPDX-License-Identifier: Apache-2.0

package domains_test

import (
	"context"
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/absmach/magistrala"
	"github.com/absmach/magistrala/auth"
	grpcapi "github.com/absmach/magistrala/auth/api/grpc/domains"
	"github.com/absmach/magistrala/pkg/apiutil"
	"github.com/absmach/magistrala/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port            = 8081
	secret          = "secret"
	email           = "test@example.com"
	id              = "testID"
	thingsType      = "things"
	usersType       = "users"
	description     = "Description"
	groupName       = "mgx"
	adminpermission = "admin"

	authoritiesObj  = "authorities"
	memberRelation  = "member"
	loginDuration   = 30 * time.Minute
	refreshDuration = 24 * time.Hour
	invalidDuration = 7 * 24 * time.Hour
	validToken      = "valid"
	inValidToken    = "invalid"
	validPolicy     = "valid"
)

var authAddr = fmt.Sprintf("localhost:%d", port)

func startGRPCServer(svc auth.Service, port int) *grpc.Server {
	listener, _ := net.Listen("tcp", fmt.Sprintf(":%d", port))
	server := grpc.NewServer()
	magistrala.RegisterDomainsServiceServer(server, grpcapi.NewDomainsServer(svc))
	go func() {
		err := server.Serve(listener)
		assert.Nil(&testing.T{}, err, fmt.Sprintf(`"Unexpected error creating auth server %s"`, err))
	}()

	return server
}

func TestDeleteUserFromDomains(t *testing.T) {
	conn, err := grpc.NewClient(authAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	assert.Nil(t, err, fmt.Sprintf("Unexpected error creating client connection %s", err))
	grpcClient := grpcapi.NewDomainsClient(conn, time.Second)

	cases := []struct {
		desc          string
		token         string
		deleteUserReq *magistrala.DeleteUserReq
		deleteUserRes *magistrala.DeleteUserRes
		err           error
	}{
		{
			desc:  "delete valid req",
			token: validToken,
			deleteUserReq: &magistrala.DeleteUserReq{
				Id: id,
			},
			deleteUserRes: &magistrala.DeleteUserRes{Deleted: true},
			err:           nil,
		},
		{
			desc:          "delete invalid req with invalid token",
			token:         inValidToken,
			deleteUserReq: &magistrala.DeleteUserReq{},
			deleteUserRes: &magistrala.DeleteUserRes{Deleted: false},
			err:           apiutil.ErrMissingID,
		},
		{
			desc:  "delete invalid req with invalid token",
			token: inValidToken,
			deleteUserReq: &magistrala.DeleteUserReq{
				Id: id,
			},
			deleteUserRes: &magistrala.DeleteUserRes{Deleted: false},
			err:           apiutil.ErrMissingPolicyEntityType,
		},
	}
	for _, tc := range cases {
		repoCall := svc.On("DeleteUserFromDomains", mock.Anything, tc.deleteUserReq.Id).Return(tc.err)
		dpr, err := grpcClient.DeleteUserFromDomains(context.Background(), tc.deleteUserReq)
		assert.Equal(t, tc.deleteUserRes.GetDeleted(), dpr.GetDeleted(), fmt.Sprintf("%s: expected %v got %v", tc.desc, tc.deleteUserRes.GetDeleted(), dpr.GetDeleted()))
		assert.True(t, errors.Contains(err, tc.err), fmt.Sprintf("%s: expected %s got %s\n", tc.desc, tc.err, err))
		repoCall.Unset()
	}
}
