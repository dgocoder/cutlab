package companyrepo

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/google/uuid"
	"github.com/testd/cutlab/internal/core/domain"
	"github.com/testd/cutlab/pkg/fstransform"
)

type CompanyFirestoreRespository struct {
	client *firestore.Client
	ctx    context.Context
}

const collection = "companies"

func NewCompanyFirestoreRespository(firestoreClient *firestore.Client, ctx context.Context) *CompanyFirestoreRespository {
	return &CompanyFirestoreRespository{client: firestoreClient, ctx: ctx}
}

func (repo *CompanyFirestoreRespository) Get(id uuid.UUID) (record domain.Company, err error) {
	dsnap, _ := repo.client.Collection(collection).Doc(id.String()).Get(repo.ctx)
	err = dsnap.DataTo(&record)
	if err != nil {
		fmt.Println(err)
		return record, err
	}
	return record, nil
}
func (repo *CompanyFirestoreRespository) Save(record domain.Company) error {
	_, err := repo.client.Collection(collection).Doc(record.ID.String()).Set(repo.ctx, fstransform.ToFirestoreMap(record))
	if err != nil {
		return err
	}
	return nil
}
