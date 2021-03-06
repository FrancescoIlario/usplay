package mongostore

import (
	"context"

	"github.com/FrancescoIlario/usplay/internal/services/order/storage"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *mongoStore) Read(ctx context.Context, id uuid.UUID) (*storage.Order, error) {
	filter := bson.M{"_id": id.String()}

	sr := s.Collection.FindOne(ctx, filter)
	if err := sr.Err(); err != nil {
		return nil, err
	}

	var res storage.Order
	if err := sr.Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}
