package locationrepo

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/google/uuid"
	"github.com/testd/cutlab/internal/core/domain"
	"github.com/testd/cutlab/pkg/fstransform"
)

type LocationFirestoreRespository struct {
	client *firestore.Client
	ctx    context.Context
}

const collection = "locations"

func NewLocationFirestoreRespository(firestoreClient *firestore.Client, ctx context.Context) *LocationFirestoreRespository {
	return &LocationFirestoreRespository{client: firestoreClient, ctx: ctx}
}

func (repo *LocationFirestoreRespository) Get(id uuid.UUID) (record domain.Location, err error) {
	dsnap, _ := repo.client.Collection(collection).Doc(id.String()).Get(repo.ctx)
	err = fstransform.DataTo(dsnap.Data(), &record)
	if err != nil {
		fmt.Println(err)
		return record, err
	}
	return record, nil
}
func (repo *LocationFirestoreRespository) Save(record domain.Location) error {
	_, err := repo.client.Collection(collection).Doc(record.ID.String()).Set(repo.ctx, fstransform.ToFirestoreMap(record))
	if err != nil {
		return err
	}
	return nil
}
