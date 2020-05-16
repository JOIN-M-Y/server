package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/JOIN-M-Y/server/study/entity"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/bson"
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
		interestedFieldDetail []string,
		ownerProfileID string,
		membersProfileID []string,
	) (entity.Study, error)
	FindByID(
		studyID string,
	) (entity.Study, error)
	FindByOwnerProfileID(
		ownerProfileID string,
	) ([]entity.Study, error)
	Find(limit int, cursor, interested string) ([]entity.Study, error)
	Update(
		studyID string,
		title string,
		description string,
		recruitment int,
		recruitEndDate time.Time,
		public bool,
		addressFirstDepthName string,
		addressSecondDepthName string,
		interestedField string,
		interestedFieldDetail []string,
		membersProfileID []string,
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
	interestedFieldDetail []string,
	ownerProfileID string,
	membersProfileID []string,
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
		InterestedFieldDetail:  interestedFieldDetail,
		OwnerProfileID:         ownerProfileID,
		MembersProfileID:       membersProfileID,
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

// FindByID find study by studyID
func (repository *Repository) FindByID(
	studyID string,
) (entity.Study, error) {
	studyEntity := entity.Study{}
	if cache := repository.getCache(studyID); cache != nil {
		return *cache, nil
	}
	repository.mongo.FindOne(
		context.TODO(),
		bson.M{"_id": studyID, "deletedAt": nil},
	).Decode(&studyEntity)
	repository.setCache(studyID, &studyEntity)
	return studyEntity, nil
}

// FindByOwnerProfileID find study by study owner profileID
func (repository *Repository) FindByOwnerProfileID(
	ownerProfileID string,
) ([]entity.Study, error) {
	studyEntityList := []entity.Study{}
	cursor, err := repository.mongo.Find(
		context.TODO(),
		bson.M{
			"ownerProfileId": ownerProfileID,
			"deletedAt":      nil,
		},
	)
	if err != nil {
		panic(err)
	}
	cursor.All(context.TODO(), &studyEntityList)
	return studyEntityList, nil
}

// Find find study list
func (repository *Repository) Find(limit int, cursor, interested string) ([]entity.Study, error) {
	studyEntityList := []entity.Study{}
	cursorEntity := entity.Study{}
	repository.mongo.FindOne(
		context.TODO(),
		bson.M{"_id": cursor, "deletedAt": nil},
	).Decode(&cursorEntity)

	if cursorEntity.ID == "" {
		repository.mongo.FindOne(
			context.TODO(),
			bson.M{"deletedAt": nil},
		).Decode(&cursorEntity)
	}

	mongoCursor, err := repository.mongo.Find(
		context.TODO(),
		bson.M{
			"interestedField": interested,
			"createdAt":       cursorEntity.CreatedAt,
			"deletedAt":       nil,
		},
	)
	if err != nil {
		return nil, err
	}
	mongoCursor.All(context.TODO(), &studyEntityList)
	return studyEntityList, nil
}

// Update update study data
func (repository *Repository) Update(
	studyID string,
	title string,
	description string,
	recruitment int,
	recruitEndDate time.Time,
	public bool,
	addressFirstDepthName string,
	addressSecondDepthName string,
	interestedField string,
	interestedFieldDetail []string,
	membersProfileID []string,
) (entity.Study, error) {
	condition := bson.M{"_id": studyID}
	_, err := repository.mongo.UpdateOne(
		context.TODO(),
		condition,
		bson.M{
			"$set": bson.M{
				"title":                  title,
				"description":            description,
				"recruitment":            recruitment,
				"recruitEndDate":         recruitEndDate,
				"public":                 public,
				"addressFirstDepthName":  addressFirstDepthName,
				"addressSecondDepthName": addressSecondDepthName,
				"interestedField":        interestedField,
				"interestedFieldDetail":  interestedFieldDetail,
				"membersProfileId":       membersProfileID,
			},
		},
	)
	if err != nil {
		panic(err)
	}
	updated := entity.Study{}
	repository.mongo.FindOne(
		context.TODO(),
		bson.M{"_id": studyID},
	).Decode(&updated)
	repository.setCache(studyID, &updated)
	return updated, nil
}
