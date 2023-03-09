package cmd

import (
	"context"
	"fmt"

	"github.com/scarlet0725/prism-api/ent"
)

func Migrate(ctx context.Context, ent *ent.Client) error {
	if err := ent.Schema.Create(ctx); err != nil {
		return fmt.Errorf("failed creating schema resources: %v", err)
	}

	// Create a new user.
	_, err := ent.User.Create().
		SetUsername("Admin").
		SetUserID("7CAMWFMHF4YXS23P").
		SetEmail("example@example.com").
		SetIsAdminVerified(true).
		SetDeleteProtected(true).
		SetPassword([]byte("$2a$10$h2qZT.YdMDWYRZV36LkUQO8A6sB.coL8mzkMl25VA2eOAGfJTcGZ2")).
		SetAPIKey("776f9e9c12eeffad240a488d12d5c8276c947ec3d67dcce5520be08580755f8edff66e5b502f27e7c400f5b96927e478426f44ee823b484951ba789e4ed1e070").
		Save(ctx)

	if err != nil {
		return fmt.Errorf("failed creating user: %v", err)
	}

	return nil
}
