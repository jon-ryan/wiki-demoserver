package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var documentMap DocumentMap

// getAllDocuments responds with a json object containing all documents
func getAllDocuments(c *fiber.Ctx) error {
	// convert map to list
	var documents DocumentList
	for _, value := range documentMap.getMap() {
		documents.addDocument(&value)
	}

	return c.JSON(documents)
}

// addDocument receives a json object and adds it to the documents and responds with the document
func addDocument(c *fiber.Ctx) error {
	doc := new(Document)

	if err := c.BodyParser(doc); err != nil {
		return err
	}

	documentMap.addDocument(doc)
	return c.JSON(doc)
}

func main() {
	// init documentList
	documentMap.init()

	doc := Document{"Hello World", "This is the abstract", "Basic body"}
	documentMap.addDocument(&doc)
	doc2 := Document{"Hello GO!", "Opinion about go", "Go is nice!"}
	documentMap.addDocument(&doc2)

	app := fiber.New()

	app.Use(logger.New())

	app.Static("/", "./static_files")
	app.Static("/home", "./static_files")

	app.Get("/api/getalldocs", getAllDocuments)
	app.Post("/api/addnewdoc", addDocument)

	log.Fatal(app.Listen(":8080"))
}
