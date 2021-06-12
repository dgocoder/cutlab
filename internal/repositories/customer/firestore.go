package customerrepo

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/google/uuid"
	"github.com/testd/cutlab/internal/core/domain"
	"github.com/testd/cutlab/pkg/fstransform"
)

type CustomerFirestoreRespository struct {
	client *firestore.Client
	ctx    context.Context
}

const collection = "customers"

func NewCustomerFirestoreRespository(firestoreClient *firestore.Client, ctx context.Context) *CustomerFirestoreRespository {
	return &CustomerFirestoreRespository{client: firestoreClient, ctx: ctx}
}

func (repo *CustomerFirestoreRespository) Get(id uuid.UUID) (record domain.Customer, err error) {
	dsnap, _ := repo.client.Collection(collection).Doc(id.String()).Get(repo.ctx)
	err = dsnap.DataTo(&record)
	if err != nil {
		fmt.Println(err)
		return record, err
	}
	return record, nil
}
func (repo *CustomerFirestoreRespository) Save(record domain.Customer) error {
	_, err := repo.client.Collection(collection).Doc(record.ID.String()).Set(repo.ctx, fstransform.ToFirestoreMap(record))
	if err != nil {
		return err
	}
	return nil
}
