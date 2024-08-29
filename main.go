// Ortelius v11 Textfile Microservice that handles creating and retrieving Textfiles
package main

import (
	"context"

	_ "github.com/ortelius/scec-textfile/docs"

	"github.com/arangodb/go-driver/v2/arangodb"
	"github.com/arangodb/go-driver/v2/arangodb/shared"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/ortelius/scec-commons/database"
	"github.com/ortelius/scec-commons/model"
)

var logger = database.InitLogger()
var dbconn = database.InitializeDatabase()

// NewLicense godoc
// @Summary Create a License
// @Description Create a new License and persist it
// @Tags license
// @Accept application/json
// @Produce json
// @Success 200
// @Router /msapi/license [post]
func NewLicense(c *fiber.Ctx) error {

	var err error                  // for error handling
	var meta arangodb.DocumentMeta // data about the document
	var ctx = context.Background() // use default database context
	license := model.NewLicense()  // define a textfile to be returned
	compid := c.Params("compid")

	if err = c.BodyParser(license); err != nil { // parse the JSON into the textfile object
		return c.Status(503).Send([]byte(err.Error()))
	}

	cid, _ := database.MakeNFT(license) // normalize the object into NFTs and JSON string for db persistence
	license.Key = cid                   // use the cid for the key.  The graph will handle the compver to readme relationship

	var resp arangodb.CollectionDocumentCreateResponse
	// add the textfile to the database.  Ignore if it already exists since it will be identical
	if resp, err = dbconn.Collections["licenses"].CreateDocument(ctx, license); err != nil && !shared.IsConflict(err) {
		logger.Sugar().Errorf("Failed to create document: %v", err)
	}
	meta = resp.DocumentMeta
	logger.Sugar().Infof("Created document in collection '%s' in db '%s' key='%s'\n", dbconn.Collections["licenses"].Name(), dbconn.Database.Name(), meta.Key)

	aql := `
		UPSERT { _from: @compid, _to: @cid }
		INSERT { _from: @compid, _to: @cid }
		UPDATE {} IN comp2licenses
	`

	// Define the parameters
	parameters := map[string]interface{}{
		"compid": "components/" + compid,
		"cid":    "licenses/" + cid,
	}
	// Execute the query
	cursor, err := dbconn.Database.Query(ctx, aql, &arangodb.QueryOptions{BindVars: parameters})
	if err != nil {
		logger.Sugar().Infof("Failed to execute query: %v", err)
	}
	defer cursor.Close()

	// associate the compver to the readme in the graph (many -> one)
	logger.Sugar().Infof("%s -> %s\n", compid, cid)
	var res model.ResponseKey
	res.Key = cid

	return c.JSON(res) // return the cid
}

// NewSwagger godoc
// @Summary Create a Swagger
// @Description Create a new Swagger and persist it
// @Tags swagger
// @Accept application/json
// @Produce json
// @Success 200
// @Router /msapi/swagger [post]
func NewSwagger(c *fiber.Ctx) error {

	var err error                  // for error handling
	var meta arangodb.DocumentMeta // data about the document
	var ctx = context.Background() // use default database context
	swagger := model.NewSwagger()  // define a textfile to be returned
	compid := c.Params("compid")

	if err = c.BodyParser(swagger); err != nil { // parse the JSON into the textfile object
		return c.Status(503).Send([]byte(err.Error()))
	}

	cid, _ := database.MakeNFT(swagger) // normalize the object into NFTs and JSON string for db persistence
	swagger.Key = cid                   // use the cid for the key.  The graph will handle the compver to readme relationship

	var resp arangodb.CollectionDocumentCreateResponse
	// add the textfile to the database.  Ignore if it already exists since it will be identical
	if resp, err = dbconn.Collections["swagger"].CreateDocument(ctx, swagger); err != nil && !shared.IsConflict(err) {
		logger.Sugar().Errorf("Failed to create document: %v", err)
	}
	meta = resp.DocumentMeta
	logger.Sugar().Infof("Created document in collection '%s' in db '%s' key='%s'\n", dbconn.Collections["swagger"].Name(), dbconn.Database.Name(), meta.Key)

	aql := `
		UPSERT { _from: @compid, _to: @cid }
		INSERT { _from: @compid, _to: @cid }
		UPDATE {} IN comp2swagger
	`

	// Define the parameters
	parameters := map[string]interface{}{
		"compid": "components/" + compid,
		"cid":    "swagger/" + cid,
	}
	// Execute the query
	cursor, err := dbconn.Database.Query(ctx, aql, &arangodb.QueryOptions{BindVars: parameters})
	if err != nil {
		logger.Sugar().Infof("Failed to execute query: %v", err)
	}
	defer cursor.Close()

	// associate the compver to the readme in the graph (many -> one)
	logger.Sugar().Infof("%s -> %s\n", compid, cid)
	var res model.ResponseKey
	res.Key = cid

	return c.JSON(res) // return the cid
}

// NewReadme godoc
// @Summary Create a Textfile
// @Description Create a new Textfile and persist it
// @Tags textfile
// @Accept application/json
// @Produce json
// @Success 200
// @Router /msapi/textfile [post]
func NewReadme(c *fiber.Ctx) error {

	var err error                  // for error handling
	var meta arangodb.DocumentMeta // data about the document
	var ctx = context.Background() // use default database context
	readme := model.NewReadme()    // define a textfile to be returned
	compid := c.Params("compid")

	if err = c.BodyParser(readme); err != nil { // parse the JSON into the textfile object
		return c.Status(503).Send([]byte(err.Error()))
	}

	cid, _ := database.MakeNFT(readme) // normalize the object into NFTs and JSON string for db persistence
	readme.Key = cid                   // use the cid for the key.  The graph will handle the compver to readme relationship

	var resp arangodb.CollectionDocumentCreateResponse
	// add the textfile to the database.  Ignore if it already exists since it will be identical
	if resp, err = dbconn.Collections["readmes"].CreateDocument(ctx, readme); err != nil && !shared.IsConflict(err) {
		logger.Sugar().Errorf("Failed to create document: %v", err)
	}
	meta = resp.DocumentMeta
	logger.Sugar().Infof("Created document in collection '%s' in db '%s' key='%s'\n", dbconn.Collections["readmes"].Name(), dbconn.Database.Name(), meta.Key)

	aql := `
		UPSERT { _from: @compid, _to: @cid }
		INSERT { _from: @compid, _to: @cid }
		UPDATE {} IN comp2readmes
	`

	// Define the parameters
	parameters := map[string]interface{}{
		"compid": "components/" + compid,
		"cid":    "readmes/" + cid,
	}
	// Execute the query
	cursor, err := dbconn.Database.Query(ctx, aql, &arangodb.QueryOptions{BindVars: parameters})
	if err != nil {
		logger.Sugar().Infof("Failed to execute query: %v", err)
	}
	defer cursor.Close()

	// associate the compver to the readme in the graph (many -> one)
	logger.Sugar().Infof("%s -> %s\n", compid, cid)
	var res model.ResponseKey
	res.Key = cid

	return c.JSON(res) // return the cid
}

// setupRoutes defines maps the routes to the functions
func setupRoutes(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault)  // handle displaying the swagger
	app.Post("/msapi/readme/:compid", NewReadme)   // save a single textfile
	app.Post("/msapi/license/:compid", NewLicense) // save a single license
	app.Post("/msapi/swagger/:compid", NewSwagger) // save a single textfile
}

// @title Ortelius v11 Textfile Microservice
// @version 11.0.0
// @description RestAPI for the Readme, License and Swagger objects.  Only for new objects.  Retrieval will be done directly against the db by other microservices.
// @description ![Release](https://img.shields.io/github/v/release/ortelius/scec-textfile?sort=semver)
// @description ![license](https://img.shields.io/github/license/ortelius/.github)
// @description
// @description ![Build](https://img.shields.io/github/actions/workflow/status/ortelius/scec-textfile/build-push-chart.yml)
// @description [![MegaLinter](https://github.com/ortelius/scec-textfile/workflows/MegaLinter/badge.svg?branch=main)](https://github.com/ortelius/scec-textfile/actions?query=workflow%3AMegaLinter+branch%3Amain)
// @description ![CodeQL](https://github.com/ortelius/scec-textfile/workflows/CodeQL/badge.svg)
// @description [![OpenSSF-Scorecard](https://api.securityscorecards.dev/projects/github.com/ortelius/scec-textfile/badge)](https://api.securityscorecards.dev/projects/github.com/ortelius/scec-textfile)
// @description
// @description ![Discord](https://img.shields.io/discord/722468819091849316)

// @termsOfService http://swagger.io/terms/
// @contact.name Ortelius Google Group
// @contact.email ortelius-dev@googlegroups.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /msapi/textfile
func main() {
	port := ":" + database.GetEnvDefault("MS_POST", "8084")
	app := fiber.New()                       // create a new fiber application
	setupRoutes(app)                         // define the routes for this microservice
	if err := app.Listen(port); err != nil { // start listening for incoming connections
		logger.Sugar().Fatalf("Failed get the microservice running: %v", err)
	}
}
