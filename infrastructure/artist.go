package infrastructure

import (
	"context"

	"github.com/scarlet0725/prism-api/ent"
	"github.com/scarlet0725/prism-api/ent/artist"
	"github.com/scarlet0725/prism-api/infrastructure/repository"
	"github.com/scarlet0725/prism-api/infrastructure/translator"
	"github.com/scarlet0725/prism-api/model"
)

//TODO: GORMの関係でuintになっているところがあるので後々修正する

type artistRepository struct {
	db *ent.Client
}

func NewArtistRepository(ent *ent.Client) repository.Artist {
	return &artistRepository{
		db: ent,
	}
}

func (a *artistRepository) CreateArtist(artist *model.Artist) (*model.Artist, error) {
	result, err := a.db.Artist.Create().SetArtistID(artist.ArtistID).SetName(artist.Name).SetURL(artist.URL).Save(context.Background())
	if err != nil {
		return nil, err
	}

	artist.ID = uint(result.ID)

	return artist, nil

}

func (a *artistRepository) GetArtistByName(name string) (*model.Artist, error) {
	result, err := a.db.Artist.Query().Where(artist.Name(name)).First(context.Background())
	if err != nil {
		return nil, err
	}

	return translator.ArtistFromEnt(result), nil
}

func (a *artistRepository) GetArtistByID(id string) (*model.Artist, error) {
	result, err := a.db.Artist.Query().Where(artist.ArtistID(id)).First(context.Background())
	if err != nil {
		return nil, err
	}

	return translator.ArtistFromEnt(result), nil
}

func (a *artistRepository) GetArtistsByIDs(ids []string) ([]*model.Artist, error) {
	result, err := a.db.Artist.Query().Where(artist.ArtistIDIn(ids...)).All(context.Background())

	if err != nil {
		return nil, err
	}

	var artists []*model.Artist
	for _, v := range result {
		artists = append(artists, translator.ArtistFromEnt(v))
	}

	return artists, nil
}
