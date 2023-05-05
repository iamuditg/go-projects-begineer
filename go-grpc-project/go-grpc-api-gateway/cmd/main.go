package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iamuditg/go-grpc-api-gateway/pkg/auth"
	"github.com/iamuditg/go-grpc-api-gateway/pkg/config"
	"github.com/iamuditg/go-grpc-api-gateway/pkg/order"
	"github.com/iamuditg/go-grpc-api-gateway/pkg/product"
	"log"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, &authSvc)
	order.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}
