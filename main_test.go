package main

import (
	"testing"
)

func TestDocuments(t *testing.T) {
	var documentList DocumentList
	documentList.init()

	doc1Title := "Test document"
	doc1Abstract := "Abstract of test document"
	doc1Body := "This is the body of the test doc"

	// add a document
	documentList.addDocument(&Document{doc1Title, doc1Abstract, doc1Body})

	// get the document
	doc1, err := documentList.getDocument(doc1Title)

	if err != nil {
		t.Errorf("Document 1 could not be retrieved")
	}
	// validate
	if doc1.getTitle() != doc1Title {
		t.Errorf("Title of Doc1 is not correct")
	}

	if doc1.getAbstract() != doc1Abstract {
		t.Errorf("Abstract of Doc 1 is not correct")
	}

	if doc1.getBody() != doc1Body {
		t.Errorf("Body of Doc 1 is not correct")
	}

	// get not existing document
	doc2, err2 := documentList.getDocument("Not existing")
	if err2 == nil {
		t.Errorf("Doc 2 should not be available")
	}
	if doc2.getTitle() != "" || doc2.getAbstract() != "" || doc2.getBody() != "" {
		t.Errorf("Doc 2 should be all null")
	}
}
