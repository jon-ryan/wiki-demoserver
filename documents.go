package main

import "errors"

// Document structures a simple document for the wiki
type Document struct {
	Title    string `json:"title"`
	Abstract string `json:"abstract"`
	Body     string `json:"body"`
}

// document get methods
func (d *Document) getTitle() string {
	return d.Title
}

func (d *Document) getAbstract() string {
	return d.Abstract
}

func (d *Document) getBody() string {
	return d.Body
}

// document set methods
func (d *Document) setTitle(title string) string {
	d.Title = title
	return d.Title
}

func (d *Document) setAbstract(abstract string) string {
	d.Abstract = abstract
	return d.Abstract
}

func (d *Document) setBody(body string) string {
	d.Body = body
	return d.Body
}

// DocumentMap groups Documents into a map
type DocumentMap struct {
	Documents map[string]Document `json:"Documents"`
}

// init docmap
func (d *DocumentMap) init() {
	d.Documents = make(map[string]Document)
}

// documentmap methods
func (d *DocumentMap) addDocument(doc *Document) {
	d.Documents[doc.Title] = *doc
}

func (d *DocumentMap) getDocument(title string) (*Document, error) {
	doc := d.Documents[title]

	if doc.getTitle() == "" {
		return &Document{"", "", ""}, errors.New("404 - Document does not exist")
	}
	return &doc, nil
}

func (d *DocumentMap) getMap() map[string]Document {
	return d.Documents
}

// DocumentList groups Documents into a list
type DocumentList struct {
	Documents []Document `json: "Documents"`
}

func (d *DocumentList) addDocument(doc *Document) {
	d.Documents = append(d.Documents, *doc)
}
