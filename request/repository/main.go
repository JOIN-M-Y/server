package repositroy

import (
	"context"
	"time"

	"github.com/JOIN-M-Y/server/request/entity"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Interface repository interface
type Interface interface {
	Create(
		requestID string,
		profileID string,
		studyID string,
	) (entity.Request, error)
	Update(
		requestID string,
		status string,
	) (entity.Request, error)
}

// Repository repository for request data
type Repository struct {
	redis *redis.Client
	mongo *mongo.Collection
}

// New create repository instance
func New(
	redis *redis.Client, mongo *mongo.Collection,
) Interface {
	return &Repository{mongo: mongo, redis: redis}
}

// Create create request
func (repository *Repository) Create(
	requestID string,
	profileID string,
	studyID string,
) (entity.Request, error) {
	requestEntity := entity.Request{
		ID:        requestID,
		ProfileID: profileID,
		StudyID:   studyID,
		Status:    "requested",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	insertResult, err := repository.mongo.InsertOne(
		context.TODO(),
		requestEntity,
	)
	if insertResult == nil || err != nil {
		panic(err)
	}
	return requestEntity, nil
}

// Update update request
func (repository *Repository) Update(
	requestID string,
	status string,
) (entity.Request, error) {
	condition := bson.M{"_id": requestID}
	_, err := repository.mongo.UpdateOne(
		context.TODO(),
		condition,
		bson.M{
			"$set": bson.M{
				"status":    status,
				"updatedAt": time.Now(),
			},
		},
	)
	if err != nil {
		panic(err)
	}
	updated := entity.Request{}
	repository.mongo.FindOne(
		context.TODO(),
		bson.M{"_id": requestID},
	).Decode(&updated)
	return updated, nil
}
