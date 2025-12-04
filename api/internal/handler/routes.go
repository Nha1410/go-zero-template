package handler

import (
	"github.com/Nha1410/go-zero-template/api/internal/middleware"
	"github.com/Nha1410/go-zero-template/api/internal/svc"
	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	// Create middleware
	authMiddleware := middleware.NewAuthMiddleware(serverCtx)

	// Public routes
	server.AddRoute(
		rest.Route{
			Method:  "GET",
			Path:    "/health",
			Handler: HealthCheckHandler(serverCtx),
		},
	)

	// Protected routes - require authentication
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{authMiddleware.Handle},
			[]rest.Route{
				{
					Method:  "POST",
					Path:    "/api/v1/users",
					Handler: CreateUserHandler(serverCtx),
				},
				{
					Method:  "GET",
					Path:    "/api/v1/users/:id",
					Handler: GetUserHandler(serverCtx),
				},
				{
					Method:  "GET",
					Path:    "/api/v1/users",
					Handler: GetUsersHandler(serverCtx),
				},
				{
					Method:  "PUT",
					Path:    "/api/v1/users/:id",
					Handler: UpdateUserHandler(serverCtx),
				},
				{
					Method:  "DELETE",
					Path:    "/api/v1/users/:id",
					Handler: DeleteUserHandler(serverCtx),
				},
			}...,
		),
	)
}
