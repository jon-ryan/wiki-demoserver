package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Document structures a simple document for the wiki
type Document struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// DocumentList groups Documents into a list
type DocumentList struct {
	Documents []Document `json:"Documents"`
}

func (d *DocumentList) addDocument(doc *Document) {
	d.Documents = append(d.Documents, *doc)
}

var documentList DocumentList

//var documentList []Document

// getAllDocuments responds with a json object containing all documents
func getAllDocuments(c *fiber.Ctx) error {
	return c.JSON(documentList)
}

// addDocument receives a json object and adds it to the documents and responds with the document
func addDocument(c *fiber.Ctx) error {
	doc := new(Document)

	if err := c.BodyParser(doc); err != nil {
		return err
	}

	documentList.addDocument(doc)
	return c.JSON(doc)
}

func main() {

	doc := Document{"Hello World", "Basic desc"}
	documentList.addDocument(&doc)
	doc2 := Document{"Hello GO!", "Go is nice!"}
	documentList.addDocument(&doc2)

	app := fiber.New()

	app.Use(logger.New())

	app.Static("/", "./static_files")
	app.Static("/home", "./static_files")

	app.Get("/api/getalldocs", getAllDocuments)
	app.Post("/api/addnewdoc", addDocument)

	log.Fatal(app.Listen(":8080"))
}
