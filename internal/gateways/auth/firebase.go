package authgtw

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/testd/cutlab/internal/core/domain"
)

type AuthFirebase struct {
	client *auth.Client
	ctx    context.Context
}

func NewAuthFirebase(firebaseAuthClient *firebase.App, ctx context.Context) *AuthFirebase {
	client, _ := firebaseAuthClient.Auth(ctx)
	return &AuthFirebase{client: client, ctx: ctx}
}

func (a *AuthFirebase) CreateUser(email string, password string, name string) (domain.User, error) {
	user := domain.NewUser(name, email, nil)
	fuser := (&auth.UserToCreate{}).DisplayName(name).Email(email).Password(password).UID(user.ID.String())
	_, err := a.client.CreateUser(a.ctx, fuser)
	if err != nil {
		return user, err
	}
	return user, nil
}
