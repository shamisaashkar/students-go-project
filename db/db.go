package db

import (
	"fmt"

	// "github.com/bytedance/sonic/option"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"golang.org/x/net/context"
)

func New(cfg Config) (*mongo.Database, error) {
	opts := options.Client()
	// opts.Monotor = otelmongo.NewMonitor()
	opts.ApplyURI(cfg.URL)

	//craete mongodb connection
	client, err := mongo.NewClient(opts)
	if err != nil {
		return nil, fmt.Errorf("db New client error %w", err)
	}
	//connect to the mongodb
	{
		ctx, done := context.WithTimeout(context.Background(), cfg.ConnecttionTimeout)
		defer done()

		if err := client.Connect(ctx); err != nil {
			return nil, fmt.Errorf("db connectin error %w", err)

		}
	}
	//ping to the mongodb
	{
		ctx, done := context.WithTimeout(context.Background(), cfg.ConnecttionTimeout)
		defer done()

		if err := client.Ping(ctx, readpref.Primary()); err != nil {
			return nil, fmt.Errorf("db ping error %w", err)

		}

	}

}
