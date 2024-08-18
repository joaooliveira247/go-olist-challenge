package cmd

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joaooliveira247/go-olist-challenge/src/config"
	"github.com/joaooliveira247/go-olist-challenge/src/routes"
	"github.com/urfave/cli/v2"
)

func runAPI(ctx *cli.Context) error {
	api := gin.Default()
	routes.RoutesRegistry(api)
	port := config.APIPort
	if cliPort := ctx.Int("port"); cliPort > 0 {
		port = cliPort
	}
	if err := api.Run(fmt.Sprintf(":%d", port)); err != nil {
		return err
	}
	return nil
}
