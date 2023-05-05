package auth

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/iamuditg/go-grpc-api-gateway/pkg/auth/pb"
	"net/http"
	"strings"
)

type AuthMiddlewareConfig struct {
	svc *ServiceClient
}

func InitAuthMiddleware(svc *ServiceClient) AuthMiddlewareConfig {
	return AuthMiddlewareConfig{
		svc: svc,
	}
}

func (c *AuthMiddlewareConfig) AuthRequired(ctx *gin.Context) {
	authorication := ctx.Request.Header.Get("authorization")
	if authorication == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token := strings.Split(authorication, "Bearer ")

	if len(token) < 0 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	res, err := c.svc.Client.Validate(context.Background(), &pb.ValidateRequest{
		Token: token[1],
	})

	if err != nil || res.Status != http.StatusOK {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Set("userId", res.UserId)
	ctx.Next()
}
