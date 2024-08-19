package cmd

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joaooliveira247/go-olist-challenge/src/config"
	"github.com/joaooliveira247/go-olist-challenge/src/db"
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

func createTables(ctx *cli.Context) error {
	dbConnection, err := db.GetDBConnection()

	if err != nil {
		return err
	}

	if err = db.Create(dbConnection); err != nil {
		return err
	}
	return nil
}

func deleteTables(ctx *cli.Context) error {
	dbConnection, err := db.GetDBConnection()

	if err != nil {
		return err
	}

	if err = db.Delete(dbConnection); err != nil {
		return err
	}
	return nil
}

func Gen() *cli.App {
	app := cli.NewApp()
	app.Name = "Book Store API CLI"
	app.Description = "Book Store CLI"
	app.Usage = "Manegment of API ecosystem"

	app.Commands = []*cli.Command{
		{
			Name:  "run",
			Usage: "Run API",
			Flags: []cli.Flag{
				&cli.Int64Flag{
					Name:    "port",
					Aliases: []string{"p"},
					Usage:   "API port, if not definer, it'll use default port in settings",
				},
			},
			Action: runAPI,
		},
	}

	return app
}
