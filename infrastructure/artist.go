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

func (a *artistRepository) CreateArtist(ctx context.Context, artist *model.Artist) (*model.Artist, error) {
	result, err := a.db.Artist.Create().SetArtistID(artist.ArtistID).SetName(artist.Name).SetURL(artist.URL).Save(ctx)
	if err != nil {
		return nil, err
	}

	artist.ID = uint(result.ID)

	return artist, nil

}

func (a *artistRepository) GetArtistByName(ctx context.Context, name string) (*model.Artist, error) {
	result, err := a.db.Artist.Query().Where(artist.Name(name)).First(ctx)
	if err != nil {
		return nil, err
	}

	return translator.ArtistFromEnt(result), nil
}

func (a *artistRepository) GetArtistByID(ctx context.Context, id string) (*model.Artist, error) {
	result, err := a.db.Artist.Query().Where(artist.ArtistID(id)).First(ctx)
	if err != nil {
		return nil, err
	}

	return translator.ArtistFromEnt(result), nil
}

func (a *artistRepository) GetArtistsByIDs(ctx context.Context, ids []string) ([]*model.Artist, error) {
	result, err := a.db.Artist.Query().Where(artist.ArtistIDIn(ids...)).All(ctx)

	if err != nil {
		return nil, err
	}

	var artists []*model.Artist
	for _, v := range result {
		artists = append(artists, translator.ArtistFromEnt(v))
	}

	return artists, nil
}
