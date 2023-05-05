package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/iamuditg/go-grpc-api-gateway/pkg/auth/pb"
	"net/http"
)

type RegisterRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(ctx *gin.Context, c pb.AuthServiceClient) {
	body := RegisterRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Register(context.Background(), &pb.RegisterRequest{
		Email:    body.Email,
		Password: body.Password,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
