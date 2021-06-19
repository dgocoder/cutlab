package userrepo

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/google/uuid"
	"github.com/testd/cutlab/internal/core/domain"
	"github.com/testd/cutlab/pkg/fstransform"
)

type UserFirestoreRespository struct {
	client *firestore.Client
	ctx    context.Context
}

const collection = "users"

func NewUserFirestoreRespository(firestoreClient *firestore.Client, ctx context.Context) *UserFirestoreRespository {
	return &UserFirestoreRespository{client: firestoreClient, ctx: ctx}
}

func (repo *UserFirestoreRespository) Get(id uuid.UUID) (record domain.User, err error) {
	dsnap, _ := repo.client.Collection(collection).Doc(id.String()).Get(repo.ctx)
	err = dsnap.DataTo(&record)
	if err != nil {
		fmt.Println(err)
		return record, err
	}
	return record, nil
}
func (repo *UserFirestoreRespository) Save(record domain.User) error {
	_, err := repo.client.Collection(collection).Doc(record.ID.String()).Set(repo.ctx, fstransform.ToFirestoreMap(record))
	if err != nil {
		return err
	}
	return nil
}
