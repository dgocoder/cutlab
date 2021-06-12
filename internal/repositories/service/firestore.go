package servicerepo

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/google/uuid"
	"github.com/testd/cutlab/internal/core/domain"
	"github.com/testd/cutlab/pkg/fstransform"
)

type ServiceFirestoreRespository struct {
	client *firestore.Client
	ctx    context.Context
}

const collection = "customers"

func NewServiceFirestoreRespository(firestoreClient *firestore.Client, ctx context.Context) *ServiceFirestoreRespository {
	return &ServiceFirestoreRespository{client: firestoreClient, ctx: ctx}
}

func (repo *ServiceFirestoreRespository) Get(id uuid.UUID) (record domain.Service, err error) {
	dsnap, _ := repo.client.Collection(collection).Doc(id.String()).Get(repo.ctx)
	err = dsnap.DataTo(&record)
	if err != nil {
		fmt.Println(err)
		return record, err
	}
	return record, nil
}
func (repo *ServiceFirestoreRespository) Save(record domain.Service) error {
	_, err := repo.client.Collection(collection).Doc(record.ID.String()).Set(repo.ctx, fstransform.ToFirestoreMap(record))
	if err != nil {
		return err
	}
	return nil
}
