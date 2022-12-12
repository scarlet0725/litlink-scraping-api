package usecase_test

import (
	"fmt"
	"log"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/ory/dockertest/v3"
	"github.com/scarlet0725/prism-api/framework"
	"github.com/scarlet0725/prism-api/infrastructure"
	"github.com/scarlet0725/prism-api/model"
	"github.com/scarlet0725/prism-api/schema"
	"github.com/scarlet0725/prism-api/usecase"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dbName     = "prism_api"
	dbUser     = "prism"
	dbPassword = "password"
)

var (
	db *gorm.DB
)

func newTestUser() *model.User {
	hash, _ := bcrypt.GenerateFromPassword([]byte("test"), bcrypt.DefaultCost)

	return &model.User{
		Email:    "testuser@example.com",
		Password: hash,
		Username: "testuser",
		APIKey:   "0c66825fbe3f76f8e802c30277ef1d950b3eb9b5c7806c41a420c26bc2176260167b259beab2178c139307d8a2b4a03a51efb904dc79f81256577c36eced7f51", //Hashed "TestAPIKey"
	}

}

func TestMain(m *testing.M) {

	// Setup docker

	pool, err := dockertest.NewPool("")

	pool.MaxWait = time.Minute * 2

	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	err = pool.Client.Ping()

	if err != nil {
		log.Fatalf("Could not ping docker: %s", err)
	}

	containerOptions := &dockertest.RunOptions{
		Repository: "mysql",
		Tag:        "8.0.31",
		Env: []string{
			"MYSQL_ROOT_PASSWORD=password",
			"MYSQL_USER=prism",
			"MYSQL_PASSWORD=password",
			"MYSQL_DATABASE=prism_api",
		},
	}

	docker, err := pool.RunWithOptions(containerOptions)

	if err != nil {
		log.Fatalf("Could not start docker: %s", err)
	}

	time.Sleep(time.Second * 15)

	dsn := fmt.Sprintf("%s:%s@tcp(localhost:%s)/%s?charset=utf8mb4&parseTime=true&loc=Asia%%2FTokyo", dbUser, dbPassword, docker.GetPort("3306/tcp"), dbName)

	gorm, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Could not connect to database: %s", err)
		pool.Purge(docker)
	}

	db = gorm

	db.AutoMigrate(
		&schema.User{},
		&schema.Artist{},
		&schema.Event{},
		&schema.Venue{},
		&schema.RyzmEvent{},
		&schema.UnStructuredEventInformation{},
		&schema.Role{},
		&model.GoogleOAuthState{},
		&model.GoogleOAuthToken{},
	)

	fmt.Println("Test Start")
	m.Run()
	fmt.Println("Test End")

	pool.Purge(docker)

}

func TestCreateUser(t *testing.T) {
	//Migration

	gorm := infrastructure.NewGORMClient(db)

	random := framework.NewRamdomIDGenerator()
	usecase := usecase.NewUserUsecase(gorm, random)

	hash, err := bcrypt.GenerateFromPassword([]byte("test"), bcrypt.DefaultCost)

	if err != nil {
		t.Fatal(err)
	}

	//Create

	tt := []struct {
		name    string
		user    *model.User
		wantErr bool
	}{
		{
			name: "ユーザーを作成できる",
			user: &model.User{
				Username: "createusertest",
				Password: hash,
				Email:    "create-user@example.com",
			},
			wantErr: false,
		},
		{
			name: "Field: Usernameが重複している場合はユーザーを作成できない",
			user: &model.User{
				Username: "createusertest",
				Password: hash,
				Email:    "duplicate-username@example.com",
			},
			wantErr: true,
		},
		{
			name: "Field: Emailが重複している場合はユーザーを作成できない",
			user: &model.User{
				Username: "duplicate-email",
				Password: hash,
				Email:    "create-user@example.com",
			},
			wantErr: true,
		},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			_, err := usecase.CreateUser(test.user)

			switch test.wantErr {
			case true:
				if err == nil {
					t.Errorf("Failed %+v", err)
				}
			case false:
				if err != nil {
					t.Errorf("Failed %+v", err)
				}
			}

		})
	}

}

func TestGetUserByID(t *testing.T) {
	gorm := infrastructure.NewGORMClient(db)

	random := framework.NewRamdomIDGenerator()
	usecase := usecase.NewUserUsecase(gorm, random)

	user := newTestUser()

	createdUser, err := usecase.CreateUser(user)

	if err != nil {
		t.Fatal(err)
	}

	getUserTests := []struct {
		name    string
		user    *model.User
		wantErr bool
		expect  *model.User
	}{
		{
			name:    "ユーザーを取得できる",
			user:    user,
			wantErr: false,
			expect:  createdUser,
		},
		{
			name: "存在しないユーザーを取得しようとするとエラーが返る",
			user: &model.User{
				Username: "notfounduser",
			},
			wantErr: true,
		},
	}

	for _, test := range getUserTests {
		t.Run(test.name, func(t *testing.T) {
			_, err := usecase.GetUserByUserID(test.user.UserID)

			switch test.wantErr {
			case true:
				if err == nil {
					t.Errorf("Failed %+v", err)
				}
			case false:
				if err != nil {
					t.Errorf("Failed %+v", err)
				}
			}

		})
	}

}
