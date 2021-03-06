package repository

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
	Update(requestID string) (entity.Request, error)
	FindByStudyID(studyID string) ([]entity.Request, error)
	FindByID(requestID string) entity.Request
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
) (entity.Request, error) {
	condition := bson.M{"_id": requestID}
	_, err := repository.mongo.UpdateOne(
		context.TODO(),
		condition,
		bson.M{
			"$set": bson.M{
				"status":    "accepted",
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

// FindByStudyID find request data using studyId
func (repository *Repository) FindByStudyID(studyID string) ([]entity.Request, error) {
	requestEntityList := []entity.Request{}
	cursor, err := repository.mongo.Find(context.TODO(), bson.M{"studyId": studyID})
	if err != nil {
		panic(err)
	}
	cursor.All(context.TODO(), &requestEntityList)
	return requestEntityList, err
}

// FindByID find request by id
func (repository *Repository) FindByID(requestID string) entity.Request {
	request := entity.Request{}
	repository.mongo.FindOne(context.TODO(), bson.M{"_id": requestID}).Decode(&request)
	return request
}
