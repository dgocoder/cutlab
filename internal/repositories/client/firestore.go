package clientrepo

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/google/uuid"
	"github.com/testd/cutlab/internal/core/domain"
	"github.com/testd/cutlab/pkg/fstransform"
)

type ClientFirestoreRespository struct {
	client *firestore.Client
	ctx    context.Context
}

const collection = "clients"

func NewClientFirestoreRespository(firestoreClient *firestore.Client, ctx context.Context) *ClientFirestoreRespository {
	return &ClientFirestoreRespository{client: firestoreClient, ctx: ctx}
}

func (repo *ClientFirestoreRespository) Get(id uuid.UUID) (record domain.Client, err error) {
	dsnap, _ := repo.client.Collection(collection).Doc(id.String()).Get(repo.ctx)
	err = dsnap.DataTo(&record)
	if err != nil {
		fmt.Println(err)
		return record, err
	}
	return record, nil
}
func (repo *ClientFirestoreRespository) Save(record domain.Client) error {
	_, err := repo.client.Collection(collection).Doc(record.ID.String()).Set(repo.ctx, fstransform.ToFirestoreMap(record))
	if err != nil {
		return err
	}
	return nil
}
