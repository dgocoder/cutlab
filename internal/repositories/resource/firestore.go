package resourcerepo

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/google/uuid"
	"github.com/testd/cutlab/internal/core/domain"
	"github.com/testd/cutlab/pkg/fstransform"
)

type ResourceFirestoreRespository struct {
	client *firestore.Client
	ctx    context.Context
}

const collection = "resources"

func NewResourceFirestoreRespository(firestoreClient *firestore.Client, ctx context.Context) *ResourceFirestoreRespository {
	return &ResourceFirestoreRespository{client: firestoreClient, ctx: ctx}
}

func (repo *ResourceFirestoreRespository) Get(id uuid.UUID) (record domain.Resource, err error) {
	dsnap, _ := repo.client.Collection(collection).Doc(id.String()).Get(repo.ctx)
	err = fstransform.DataTo(dsnap.Data(), &record)
	if err != nil {
		fmt.Println(err)
		return record, err
	}
	return record, nil
}
func (repo *ResourceFirestoreRespository) Save(record domain.Resource) error {
	_, err := repo.client.Collection(collection).Doc(record.ID.String()).Set(repo.ctx, fstransform.ToFirestoreMap(record))
	if err != nil {
		return err
	}
	return nil
}
