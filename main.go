package main

import (
	"context"
	"net/http"
	"privilege/domain/catalog"
	"privilege/domain/category"
	"privilege/domain/privilege"
	privilegetype "privilege/domain/privilege_type"
	"privilege/domain/user"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return err
	}
	return nil
}

func main() {
	initConfig()

	e := echo.New()

	ctx, cancel := context.WithTimeout(context.Background(),
		10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:password@127.0.0.1:27017/?authSource=admin"))
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	mongodb := client.Database(viper.GetString("mongo.db"))
	
	dsn := viper.GetString("postgresql.connection")
	postgres, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	e.Validator = &CustomValidator{validator: validator.New()}

	e.GET("/healths", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK!")
	})

	e.POST("/login", user.Login(user.GetUser(postgres)))

	r := e.Group("")
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(user.JwtCustomClaims)
		},
		SigningMethod: jwt.SigningMethodHS256.Alg(),
		SigningKey:    []byte(viper.GetString("jwt.secret")),
	}

	r.Use(echojwt.WithConfig(config))

	r.GET("/catalogs", catalog.GetCatalogHandler(catalog.GetCatalog(mongodb)))
	r.GET("/categories", category.GetCatalogHandler(category.GetCategories(mongodb)))
	r.GET("/privileges/type", privilegetype.GetPrivilegeTypeHandler(privilegetype.GetPrivilegeTypes(mongodb)))

	r.GET("/privileges", privilege.GetPrivilegeHandler(privilege.GetPrivilege(mongodb)))
	r.GET("/privileges/:id", privilege.GetPrivilegeByIdHandler(privilege.GetPrivilegeById(mongodb)))
	r.POST("/privileges", privilege.CreatePrivilegeHandler(privilege.CreatePrivilege(mongodb)))
	r.PUT("/privileges", privilege.UpdatePrivilegeHandler(privilege.UpdatePrivilege(mongodb)))
	r.DELETE("/privileges/:id", privilege.DeletePrivilegeHandler(privilege.DeletePrivilege(mongodb)))

	e.Logger.Fatal(e.Start(":" + viper.GetString("app.port")))
}

func initConfig() {
	viper.SetDefault("app.port", "1323")
	viper.SetDefault("mongo.uri", "mongodb:localhost:27017")
	viper.SetDefault("mongo.db", "sxexpo")
	viper.SetDefault("mongo.user", "root")
	viper.SetDefault("mongo.pass", "password")

	viper.SetDefault("postgresql.connection", "host=localhost user=root password=password dbname=postgres port=5432 sslmode=disable")

	viper.SetDefault("jwt.secret", "secret")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Warn(err)
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
}
