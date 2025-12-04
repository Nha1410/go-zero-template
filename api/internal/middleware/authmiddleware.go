package middleware

import (
	"context"
	"net/http"

	"github.com/Nha1410/go-zero-template/api/internal/svc"
	"github.com/Nha1410/go-zero-template/common/auth"
	"github.com/Nha1410/go-zero-template/common/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type AuthMiddleware struct {
	svcCtx *svc.ServiceContext
}

func NewAuthMiddleware(svcCtx *svc.ServiceContext) *AuthMiddleware {
	return &AuthMiddleware{
		svcCtx: svcCtx,
	}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract token from request
		token := auth.ExtractTokenFromRequest(r)
		if token == "" {
			httpx.ErrorCtx(r.Context(), w, errors.ErrUnauthorized)
			return
		}

		// Validate token with Zitadel
		userInfo, err := m.svcCtx.Zitadel.ValidateToken(r.Context(), token)
		if err != nil {
			logx.Errorf("Token validation failed: %v", err)
			httpx.ErrorCtx(r.Context(), w, errors.ErrUnauthorized.WithDetails("Invalid or expired token"))
			return
		}

		// Add user info to context
		ctx := context.WithValue(r.Context(), "user_id", userInfo.Sub)
		ctx = context.WithValue(ctx, "user_email", userInfo.Email)
		ctx = context.WithValue(ctx, "user_info", userInfo)

		// Continue with the request
		next(w, r.WithContext(ctx))
	}
}

// OptionalAuthMiddleware allows requests with or without authentication
func (m *AuthMiddleware) OptionalAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := auth.ExtractTokenFromRequest(r)
		if token != "" {
			userInfo, err := m.svcCtx.Zitadel.ValidateToken(r.Context(), token)
			if err == nil {
				ctx := context.WithValue(r.Context(), "user_id", userInfo.Sub)
				ctx = context.WithValue(ctx, "user_email", userInfo.Email)
				ctx = context.WithValue(ctx, "user_info", userInfo)
				r = r.WithContext(ctx)
			}
		}
		next(w, r)
	}
}
