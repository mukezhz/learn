package main

import (
	"context"
	"errors"
	"log"

	"connectrpc.com/connect"
)

const tokenHeader = "authorization"

func NewAuthInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			if req.Spec().IsClient {
				req.Header().Set(tokenHeader, "password")
			} else if req.Header().Get(tokenHeader) == "" ||
				req.Header().Get(tokenHeader) != "password" {
				log.Println("IS IT CLIENT:: ", req.Spec().IsClient)
				return nil, connect.NewError(
					connect.CodeUnauthenticated,
					errors.New("no token provided"),
				)
			}
			return next(ctx, req)
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
