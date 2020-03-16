package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/JOIN-M-Y/server/group/entity"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Interface group repository interface
type Interface interface {
	Create(
		profileID string,
		groupID string,
		title string,
		description string,
		recruitment int,
		recruitEndDate time.Time,
		publicAbility bool,
		interestedField string,
		recommandation int,
		firstRegionDepthName string,
		secondRegionDepthName string,
	) (entity.Group, error)
	Update(
		groupID string,
		title string,
		description string,
		recruitment int,
		recruitEndDate time.Time,
		publicAbility bool,
		interestedField string,
		recommandation int,
		firstRegionDepthName string,
		secondRegionDepthName string,
	) (entity.Group, error)
	Delete(groupID string) (entity.Group, error)
}

// Repository study group repository
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
	key string, groupEntity *entity.Group,
) {
	marshaledEntity, _ := json.Marshal(&groupEntity)
	repository.redis.Set(
		"group:"+key, string(marshaledEntity), time.Second,
	)
}

func (repository *Repository) getCache(
	key string,
) *entity.Group {
	data, getDataFromRedisError :=
		repository.redis.Get("group:" + key).Result()
	if getDataFromRedisError != nil {
		return nil
	}

	entity := &entity.Group{}
	jsonUnmarshalError := json.Unmarshal([]byte(data), entity)
	if jsonUnmarshalError != nil {
		return nil
	}

	if entity.ID == "" {
		return nil
	}
	return entity
}

// Create create study group data
func (repository *Repository) Create(
	profileID string,
	groupID string,
	title string,
	description string,
	recruitment int,
	recruitEndDate time.Time,
	publicAbility bool,
	interestedField string,
	recommandation int,
	firstRegionDepthName string,
	secondRegionDepthName string,
) (entity.Group, error) {
	groupEntity := entity.Group{
		ID: groupID,
		Owner: struct {
			ProfileID string "json:\"profileId\" bson:\"profileId\""
		}{
			ProfileID: profileID,
		},
		Title:                 title,
		Description:           description,
		Recruitment:           recruitment,
		RecruitEndDate:        recruitEndDate,
		PublicAbility:         publicAbility,
		InterestedField:       interestedField,
		Recommandation:        recommandation,
		FirstRegionDepthName:  firstRegionDepthName,
		SecondRegionDepthName: secondRegionDepthName,
		CreatedAt:             time.Now(),
		UpdatedAt:             time.Now(),
	}
	insertResult, err := repository.mongo.InsertOne(
		context.TODO(),
		groupEntity,
	)
	if err != nil || insertResult == nil {
		panic(err)
	}
	repository.setCache(groupID, &groupEntity)
	return groupEntity, nil
}

// Update update study group data
func (repository *Repository) Update(
	groupID string,
	title string,
	description string,
	recruitment int,
	recruitEndDate time.Time,
	publicAbility bool,
	interestedField string,
	recommandation int,
	firstRegionDepthName string,
	secondRegionDepthName string,
) (entity.Group, error) {
	condition := bson.M{"_id": groupID}
	updateData := bson.M{
		"title":                 title,
		"description":           description,
		"recruitment":           recruitment,
		"recruitEndDate":        recruitEndDate,
		"publicAbility":         publicAbility,
		"interestedField":       interestedField,
		"recommandation":        recommandation,
		"firstRegionDepthName":  firstRegionDepthName,
		"secondRegionDepthName": secondRegionDepthName,
		"updatedAt":             time.Now(),
	}
	_, err := repository.mongo.UpdateOne(
		context.TODO(),
		condition,
		bson.M{
			"$set": updateData,
		},
	)
	if err != nil {
		panic(err)
	}
	updated := entity.Group{}
	repository.mongo.FindOne(
		context.TODO(),
		bson.M{"_id": groupID},
	).Decode(&updated)
	repository.setCache(groupID, &updated)
	return updated, nil
}

// Delete delete study group data
func (repository *Repository) Delete(
	groupID string,
) (entity.Group, error) {
	groupEntity := entity.Group{}
	condition := bson.M{"_di": groupID}
	repository.mongo.UpdateOne(
		context.TODO(),
		condition,
		bson.M{
			"$set": bson.M{
				"deletedAt": time.Now(),
			},
		},
	)
	return groupEntity, nil
}
