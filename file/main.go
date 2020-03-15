package file

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/JOIN-M-Y/server/config"
	"github.com/JOIN-M-Y/server/file/api"
	"github.com/JOIN-M-Y/server/file/aws"
	"github.com/JOIN-M-Y/server/file/command"
	"github.com/JOIN-M-Y/server/file/controller"
	"github.com/JOIN-M-Y/server/file/query"
	"github.com/JOIN-M-Y/server/file/repository"
	"github.com/JOIN-M-Y/server/util"
)

func getMongoDBClient(config config.Interface) *mongo.Collection {
	user := config.Database().User()
	password := config.Database().Password()
	host := config.Database().Host()
	port := config.Database().Port()
	clientOptions := options.Client().ApplyURI(
		"mongodb://" + user + ":" + password + "@" + host + ":" + port,
	)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}
	client.Ping(context.TODO(), nil)
	collection := client.Database(
		config.Database().Name(),
	).Collection("files")

	return collection
}

func getRedisClient(config config.Interface) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.Redis().Address(),
		Password: config.Redis().Password(),
	})
}

// Initialize init file module
func Initialize(
	engine *gin.Engine, config config.Interface, util *util.Util,
) {
	mongoClient := getMongoDBClient(config)
	redisClient := getRedisClient(config)
	repository := repository.New(redisClient, mongoClient)
	api := api.New(config)
	aws := aws.New(config)
	commandBus := command.New(repository, api, aws, config)
	queryBus := query.New(config, repository)
	controller.New(engine, commandBus, queryBus, util, api)
}
