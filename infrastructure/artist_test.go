package infrastructure_test

import (
	"context"
	"testing"
	"time"

	"log"

	"github.com/scarlet0725/prism-api/cmd"
	"github.com/scarlet0725/prism-api/infrastructure"
	"github.com/scarlet0725/prism-api/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ory/dockertest/v3"
)

func TestCreateArtist(t *testing.T) {
	t.Parallel()
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	err = pool.Client.Ping()

	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	runOptions := &dockertest.RunOptions{
		Repository: "mysql",
		Tag:        "8.0",
		Env: []string{
			"MYSQL_ROOT_PASSWORD=password",
			"MYSQL_DATABASE=prism_api",
			"MYSQL_USER=prism",
			"MYSQL_PASSWORD=password",
		},
	}

	container, err := pool.RunWithOptions(runOptions)

	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	conf := cmd.DBConfig{
		Host:     "localhost",
		Port:     container.GetPort("3306/tcp"),
		User:     "prism",
		Password: "password",
		Database: "prism_api",
	}

	time.Sleep(15 * time.Second)

	entClient, err := cmd.InitDB(conf)

	if err != nil {
		log.Fatalf("Could not Initialization DB: %s", err)
		pool.Purge(container)
	}

	err = entClient.Schema.Create(context.Background())

	if err != nil {
		log.Fatalf("Could not auto migration: %s", err)
		pool.Purge(container)
	}

	artistRepository := infrastructure.NewArtistRepository(entClient)

	t.Cleanup(func() {
		err := pool.Purge(container)
		if err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}
	})

	t.Run("artistが作成できる", func(t *testing.T) {
		_, err := artistRepository.CreateArtist(&model.Artist{
			ArtistID: "ABCDEFGHIJKLMNPQ",
			Name:     "Test",
			URL:      "",
		})

		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("artistが取得できる", func(t *testing.T) {
		_, err := artistRepository.GetArtistByID("ABCDEFGHIJKLMNPQ")

		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("artistが更新できる", func(t *testing.T) {})

}
