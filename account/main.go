package account

import (
	"github.com/JOIN-M-Y/server/account/api"
	"github.com/JOIN-M-Y/server/account/aws"
	"github.com/JOIN-M-Y/server/account/command"
	"github.com/JOIN-M-Y/server/account/controller"
	"github.com/JOIN-M-Y/server/account/email"
	"github.com/JOIN-M-Y/server/account/entity"
	"github.com/JOIN-M-Y/server/account/query"
	"github.com/JOIN-M-Y/server/account/repository"
	"github.com/JOIN-M-Y/server/config"
	"github.com/JOIN-M-Y/server/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func getDatabaseConnection(config config.Interface) *gorm.DB {
	user := config.Database().User()
	password := config.Database().Password()
	host := config.Database().Host()
	port := config.Database().Port()
	name := config.Database().Name()
	dialect := "mysql"
	args := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + name + "?parseTime=true"

	connection, err := gorm.Open(dialect, args)
	if err != nil {
		panic(err)
	}
	connection.LogMode(true)
	connection.AutoMigrate(&entity.Account{})
	return connection
}

// Initialize initialize account module
func Initialize(
	engine *gin.Engine, config config.Interface, util *util.Util,
) {
	connection := getDatabaseConnection(config)
	repository := repository.New(connection)
	email := email.New(config)
	aws := aws.New(config)
	api := api.New(config)
	commandBus := command.New(repository, email, aws, config, api)
	queryBus := query.New(config, repository)
	controller.New(engine, commandBus, queryBus, util, config, api)
}
