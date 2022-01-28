package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type ParamsToConnect struct {
	User, Host, Port, Password, Database, AuthDB string
}

func NewMongo(ctx context.Context, connect ParamsToConnect) (*mongo.Database, error) {
	var uri string
	var isAuth bool

	if connect.User == "" && connect.Password == "" {
		uri = fmt.Sprintf("mongodb://%s:%s", connect.Host, connect.Port)
	} else {
		isAuth = true
		uri = fmt.Sprintf("mongodb://%s:%s@%s:%s", connect.User, connect.Password, connect.Host, connect.Port)
	}

	clientOptions := options.Client().ApplyURI(uri)
	if isAuth {
		if connect.AuthDB == "" {
			connect.AuthDB = connect.Database
		}
		clientOptions.SetAuth(options.Credential{
			AuthSource: connect.AuthDB,
			Username:   connect.User,
			Password:   connect.Password,
		})
	}

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("cant connect to db")
		return nil, err
	}

	if err = client.Ping(ctx, nil); err != nil {
		log.Fatal("can`t ping db")
		return nil, err
	}

	return client.Database(connect.Database), nil
}
