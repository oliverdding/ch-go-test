package chgotest_test

import (
	"context"
	"testing"

	"github.com/ClickHouse/ch-go"
	"github.com/ClickHouse/ch-go/proto"
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

	var (
		host                proto.ColStr
		port                proto.ColUInt16
		status              proto.ColInt64
		error               proto.ColStr
		num_hosts_remaining proto.ColUInt64
		num_hosts_active    proto.ColUInt64
	)
	err = client.Do(context.Background(), ch.Query{
		Body: "CREATE DATABASE IF NOT EXISTS otel ON CLUSTER default_cluster;",
		Result: proto.Results{
			{Name: "host", Data: &host},
			{Name: "port", Data: &port},
			{Name: "status", Data: &status},
			{Name: "error", Data: &error},
			{Name: "num_hosts_remaining", Data: &num_hosts_remaining},
			{Name: "num_hosts_active", Data: &num_hosts_active},
		},
	})
	if err != nil {
		logger.Sugar().Fatal(err)
	}
	for i := 0; i < status.Rows(); i++ {
		if status.Row(i) != 0 {
			logger.Sugar().Errorf("%s:%d failed with error %s", host.Row(i), port.Row(i), error.Row(i))
		} else {
			logger.Sugar().Infof("%s:%d success", host.Row(i), port.Row(i))
		}
	}
}
