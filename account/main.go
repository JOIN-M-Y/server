package account

import (
	"context"

	"github.com/JOIN-M-Y/server/account/api"
	"github.com/JOIN-M-Y/server/account/aws"
	"github.com/JOIN-M-Y/server/account/command"
	"github.com/JOIN-M-Y/server/account/controller"
	"github.com/JOIN-M-Y/server/account/email"
	"github.com/JOIN-M-Y/server/account/query"
	"github.com/JOIN-M-Y/server/account/repository"
	"github.com/JOIN-M-Y/server/config"
	"github.com/JOIN-M-Y/server/util"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	).Collection("accounts")

	return collection
}

func getRedisClient(config config.Interface) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.Redis().Address(),
		Password: config.Redis().Password(),
	})
}

// Initialize innitialize account module
func Initialize(
	engine *gin.Engine, config config.Interface, util *util.Util,
) {
	mongoClient := getMongoDBClient(config)
	redisClient := getRedisClient(config)
	repository := repository.New(redisClient, mongoClient)
	email := email.New(config)
	aws := aws.New(config)
	api := api.New(config)
	commandBus := command.New(repository, email, aws, config, api)
	queryBus := query.New(config, repository)
	controller.New(engine, commandBus, queryBus, util, config, api)
}
