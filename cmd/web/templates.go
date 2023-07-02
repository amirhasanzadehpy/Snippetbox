package main

import "amirhasanzadehpy/Snippetbox.git/pkg/models"

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
