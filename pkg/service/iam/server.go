// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package iam

import (
	"google.golang.org/grpc"

	"openpitrix.io/openpitrix/pkg/config"
	"openpitrix.io/openpitrix/pkg/constants"
	"openpitrix.io/openpitrix/pkg/manager"
	"openpitrix.io/openpitrix/pkg/pb"
	"openpitrix.io/openpitrix/pkg/pi"
)

type Server struct {
	config.IAMConfig
}

func Serve(cfg *config.Config) {
	pi.SetGlobal(cfg)
	s := Server{cfg.IAM}

	manager.NewGrpcServer("iam-manager", constants.IAMServicePort).
		ShowErrorCause(cfg.Grpc.ShowErrorCause).
		WithChecker(s.Checker).
		Serve(func(server *grpc.Server) {
			pb.RegisterAccountManagerServer(server, &s)
			pb.RegisterTokenManagerServer(server, &s)
		})
}
