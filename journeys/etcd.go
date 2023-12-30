package journeys

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/viper"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func EtcdPut(cli clientv3.Client, key string, value string) error {

	ctx, _ := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	_, err := cli.Put(ctx, key, value)

	if err != nil {
		return err
	}
	return nil
}

func EtcdGetFirst(cli clientv3.Client, key string) (string, error) {

	ctx, _ := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	resp, err := cli.Get(ctx, key)

	if err != nil {
		return "", err
	}
	for _, kv := range resp.Kvs {
		// log.Println(string(kv.Key), string(kv.Value), kv.Version, kv.ModRevision)
		return string(kv.Value), nil
	}
	return "", fmt.Errorf("key not found")
}

func EtcdGetKeys(cli clientv3.Client, prefix string) (*[]string, error) {

	ctx, _ := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	resp, err := cli.Get(ctx, prefix, clientv3.WithKeysOnly(), clientv3.WithRange(prefix+"å"))
	result := []string{}
	if err != nil {
		return &result, err
	}
	for _, kv := range resp.Kvs {
		result = append(result, string(kv.Key))

	}
	return &result, nil
}

func ConnectETCD() (*clientv3.Client, error) {
	config := clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		Username:    "root",
		Password:    viper.GetString("ETCD_ROOT_PASSWORD"),
		DialTimeout: 5 * time.Second,
	}
	cli, err := clientv3.New(config)
	if err != nil {
		return nil, err
	}

	return cli, nil
}
