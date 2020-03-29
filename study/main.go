package study

import (
	"context"

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

// Initialize initialize study module
func Initialize(
	engine *gin.Engine, config config.Interface, util *util.Util,
) {

}
