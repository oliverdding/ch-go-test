package chgotest_test

import (
	"context"
	"testing"

	"github.com/ClickHouse/ch-go"
	"go.uber.org/zap"
)

func TestSimplest(t *testing.T) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		t.Fatal(err)
	}
	options := ch.Options{
		Logger:   logger,
		Address:  "127.0.0.1:9000",
		Database: "default",
		User:     "root",
		Password: "root",
	}

	client, err := ch.Dial(context.Background(), options)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	err = client.Do(context.Background(), ch.Query{
		Body: "CREATE DATABASE IF NOT EXISTS otel ON CLUSTER default_cluster;",
	})
	if err != nil {
		t.Fatal(err)
	}
}
