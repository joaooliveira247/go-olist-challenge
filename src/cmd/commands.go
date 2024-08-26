package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/joaooliveira247/go-olist-challenge/src/config"
	"github.com/joaooliveira247/go-olist-challenge/src/db"
	"github.com/joaooliveira247/go-olist-challenge/src/models"
	"github.com/joaooliveira247/go-olist-challenge/src/repositories"
	"github.com/joaooliveira247/go-olist-challenge/src/routes"
	"github.com/joaooliveira247/go-olist-challenge/src/utils"
	"github.com/urfave/cli/v2"
)

var composePath string = filepath.Join(config.BASE_DIR, "docker-compose.yaml")

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

func dockerRun(ctx *cli.Context) error {
	cmd := exec.Command("docker", "compose", "-f", composePath, "up", "-d")

	if err := cmd.Run(); err != nil {
		return err
	}
	fmt.Println("database container up")
	return nil
}

func dockerStop(ctx *cli.Context) error {

	cmd := exec.Command("docker", "compose", "-f", composePath, "down")

	if err := cmd.Run(); err != nil {
		return err
	}

	fmt.Println("database container down")
	return nil
}

func importAuthorsFromCSV(ctx *cli.Context) error {
	header := ctx.Bool("header")
	path := ctx.Path("path")

	if _, err := os.Stat(path); err != nil {
		return err
	}

	authorsCSV, err := utils.CSVToAuthor(path, header)

	if err != nil {
		return err
	}

	var authors []models.Authors

	for _, authorCSV := range authorsCSV {
		var author models.Authors
		if err := author.ParseValidate(authorCSV); err != nil {
			return err
		}
		authors = append(authors, author)
	}

	db, err := db.GetDBConnection()
	if err != nil {
		return err
	}

	repository := repositories.NewAuthorRepository(db)

	IDs, err := repository.InsertAuthors(authors)

	if err != nil {
		return err
	}

	fmt.Println(IDs)
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
		{
			Name:    "database",
			Aliases: []string{"db"},
			Usage:   "Interact with database",
			Flags:   nil,
			Subcommands: []*cli.Command{
				{
					Name:    "create",
					Usage:   "Create all tables",
					Aliases: []string{"c"},
					Action:  createTables,
				},
				{
					Name:    "delete",
					Usage:   "Delete all tables",
					Aliases: []string{"d"},
					Action:  deleteTables,
				},
				{
					Name:    "start",
					Usage:   "Start container with Postgres image.",
					Aliases: []string{"up"},
					Action:  dockerRun,
				},
				{
					Name:    "stop",
					Usage:   "Stop container with Postgres image",
					Aliases: []string{"down"},
					Action:  dockerStop,
				},
				{
					Name:    "import",
					Usage:   "Import authors from CSV",
					Aliases: []string{"i"},
					Flags: []cli.Flag{
						&cli.PathFlag{
							Name:     "path",
							Aliases:  []string{"p"},
							Usage:    "Path to CSV file",
							Required: true,
						},
						&cli.BoolFlag{
							Name:  "header",
							Usage: "CSV header true or false",
						},
					},
					Action: importAuthorsFromCSV,
				},
			},
		},
	}

	return app
}
