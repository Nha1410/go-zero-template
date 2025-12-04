package handler

import (
	"net/http"

	"github.com/Nha1410/go-zero-template/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func HealthCheckHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		httpx.OkJsonCtx(r.Context(), w, map[string]interface{}{
			"code":    200,
			"message": "OK",
			"data": map[string]string{
				"status": "healthy",
			},
		})
	}
}
