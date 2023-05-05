package main

import (
	"context"
	"flookybooky/services/flight/ent"

	"entgo.io/ent/dialect"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func init() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
}

func main() {
	logger, err := zap.NewProduction()
	defer logger.Sync()
	if err != nil {
		panic(err)
	}
	log := logger.Sugar()
	POSTGRES_URI := string(viper.GetString("POSTGRES_URI"))
	client, err := ent.Open(dialect.Postgres, POSTGRES_URI)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// listen, err := net.Listen("tcp", ":2220")
	// if err != nil {
	// 	panic(err)
	// }

	// s := grpc.NewServer()
	// h, err := handler.NewFlightHandler(*client)
	// if err != nil {
	// 	panic(err)
	// }

	// reflection.Register(s)
	// pb.RegisterFlightServiceServer(s, h)
	// s.Serve(listen)
}
