package eventrepo

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/google/uuid"
	"github.com/testd/cutlab/internal/core/domain"
	"github.com/testd/cutlab/pkg/fstransform"
)

type EventFirestoreRespository struct {
	client *firestore.Client
	ctx    context.Context
}

const collection = "events"

func NewEventFirestoreRespository(firestoreClient *firestore.Client, ctx context.Context) *EventFirestoreRespository {
	return &EventFirestoreRespository{client: firestoreClient, ctx: ctx}
}

func (repo *EventFirestoreRespository) Get(id uuid.UUID) (record domain.Event, err error) {
	dsnap, _ := repo.client.Collection(collection).Doc(id.String()).Get(repo.ctx)
	err = dsnap.DataTo(&record)
	if err != nil {
		fmt.Println(err)
		return record, err
	}
	return record, nil
}
func (repo *EventFirestoreRespository) Save(record domain.Event) error {
	_, err := repo.client.Collection(collection).Doc(record.ID.String()).Set(repo.ctx, fstransform.ToFirestoreMap(record))
	if err != nil {
		return err
	}
	return nil
}
