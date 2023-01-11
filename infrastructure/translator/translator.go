package translator

import (
	"github.com/scarlet0725/prism-api/ent"
	"github.com/scarlet0725/prism-api/model"
)

func ArtistFromEnt(source *ent.Artist) *model.Artist {
	artist := &model.Artist{
		ID:       uint(source.ID),
		ArtistID: source.ArtistID,
		Name:     source.Name,
		URL:      source.URL,
	}

	return artist
}
