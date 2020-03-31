package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/JOIN-M-Y/server/study/entity"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
)

// Interface repository interface
type Interface interface {
	Create(
		studyID string,
		title string,
		description string,
		recruitment int,
		recruitEndDate time.Time,
		public bool,
		addressFirstDepthName string,
		addressSecondDepthName string,
		interestedField string,
		ownerProfileID string,
	) (entity.Study, error)
}

// Repository repository for study data
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

func (repository *Repository) setCache(
	key string, studyEntity *entity.Study,
) {
	marshaledEntity, _ := json.Marshal(&studyEntity)
	repository.redis.Set(
		"study:"+key, string(marshaledEntity), time.Second,
	)
}

func (repository *Repository) getCache(
	key string,
) *entity.Study {
	data, getDataFromRedisError :=
		repository.redis.Get("study:" + key).Result()
	if getDataFromRedisError != nil {
		return nil
	}

	entity := &entity.Study{}
	jsonUnmarshalError := json.Unmarshal([]byte(data), entity)
	if jsonUnmarshalError != nil {
		return nil
	}

	if entity.ID == "" {
		return nil
	}
	return entity
}

// Create creat e study
func (repository *Repository) Create(
	studyID string,
	title string,
	description string,
	recruitment int,
	recruitEndDate time.Time,
	public bool,
	addressFirstDepthName string,
	addressSecondDepthName string,
	interestedField string,
	ownerProfileID string,
) (entity.Study, error) {
	studyEntity := entity.Study{
		ID:                     studyID,
		Title:                  title,
		Description:            description,
		Recruitment:            recruitment,
		RecruitEndDate:         recruitEndDate,
		Public:                 public,
		AddressFirstDepthName:  addressFirstDepthName,
		AddressSecondDepthName: addressSecondDepthName,
		InterestedField:        interestedField,
		OwnerProfileID:         ownerProfileID,
		CreatedAt:              time.Now(),
		UpdatedAt:              time.Now(),
	}
	insertResult, err := repository.mongo.InsertOne(
		context.TODO(),
		studyEntity,
	)
	if insertResult == nil || err != nil {
		panic(err)
	}
	repository.setCache(studyID, &studyEntity)
	return studyEntity, nil
}
