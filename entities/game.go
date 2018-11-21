package entities

import (
	"fmt"
	"github.com/google/jsonapi"
)

type Game struct {
	ID    int      `jsonapi:"primary,games"`
	Name  string   `jsonapi:"attr,name"`
	Teams []*Teams `jsonapi:"relation,posts"`
}


// JSONAPILinks implements the Linkable interface for a game
func (game Game) JSONAPILinks() *jsonapi.Links {
	return &jsonapi.Links{
		"self": fmt.Sprintf("https://example.com/games/%d", game.ID),
	}
}

// JSONAPIRelationshipLinks implements the RelationshipLinkable interface for a game
func (game Game) JSONAPIRelationshipLinks(relation string) *jsonapi.Links {
	if relation == "teams" {
		return &jsonapi.Links{
			"related": fmt.Sprintf("https://example.com/teams/%d/posts", game.ID),
		}
	}
	return nil
}

// JSONAPIMeta implements the Metable interface for a game
func (game Game) JSONAPIMeta() *jsonapi.Meta {
	return &jsonapi.Meta{
		"detail": "extra details regarding the game",
	}
}

// JSONAPIRelationshipMeta implements the RelationshipMetable interface for a game
func (game Game) JSONAPIRelationshipMeta(relation string) *jsonapi.Meta {
	if relation == "posts" {
		return &jsonapi.Meta{
			"detail": "posts meta information",
		}
	}
	if relation == "current_post" {
		return &jsonapi.Meta{
			"detail": "current post meta information",
		}
	}
	return nil
}