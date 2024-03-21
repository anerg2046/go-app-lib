package database

import (
	"context"
	"errors"
	"go-app/lib/logger"
	"net/http"
	"syscall"
	"time"

	elastic6 "github.com/olivere/elastic/v6"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
)

type ElasticOption struct {
	Addresses []string
	User      string
	Pass      string
}

func ElasticConn(option ElasticOption) (client *elastic.Client) {
	var err error
	if option.Pass == "" {
		client, err = elastic.NewClient(
			elastic.SetSniff(false), elastic.SetURL(option.Addresses...),
			elastic.SetRetrier(NewEsRetrier()),
		)
	} else {
		client, err = elastic.NewClient(
			elastic.SetSniff(false),
			elastic.SetURL(option.Addresses...),
			elastic.SetBasicAuth(option.User, option.Pass),
			elastic.SetRetrier(NewEsRetrier()),
		)
	}
	if err != nil {
		logger.Error("[Elastic]", zap.Any("连接失败", err))
	}
	return client
}

type EsRetrier struct {
	backoff elastic.Backoff
}

func NewEsRetrier() *EsRetrier {
	return &EsRetrier{
		backoff: elastic.NewExponentialBackoff(10*time.Millisecond, 15*time.Second),
	}
}

func (r *EsRetrier) Retry(ctx context.Context, retry int, req *http.Request, resp *http.Response, err error) (time.Duration, bool, error) {
	// Fail hard on a specific error
	if err == syscall.ECONNREFUSED {
		return 0, false, errors.New("elasticsearch or network down")
	}

	// Stop after 5 retries
	if retry >= 5 {
		return 0, false, nil
	}

	// Let the backoff strategy decide how long to wait and whether to stop
	wait, stop := r.backoff.Next(retry)
	return wait, stop, nil
}

func Elastic6Conn(option ElasticOption) (client *elastic6.Client) {
	var err error
	if option.Pass == "" {
		client, err = elastic6.NewClient(
			elastic6.SetSniff(false), elastic6.SetURL(option.Addresses...),
			elastic6.SetRetrier(NewEs6Retrier()),
		)
	} else {
		client, err = elastic6.NewClient(
			elastic6.SetSniff(false),
			elastic6.SetURL(option.Addresses...),
			elastic6.SetBasicAuth(option.User, option.Pass),
			elastic6.SetRetrier(NewEs6Retrier()),
		)
	}
	if err != nil {
		logger.Fatal("[Elastic]", zap.Any("连接失败", err))
	}
	return client
}

type Es6Retrier struct {
	backoff elastic.Backoff
}

func NewEs6Retrier() *Es6Retrier {
	return &Es6Retrier{
		backoff: elastic6.NewExponentialBackoff(10*time.Millisecond, 15*time.Second),
	}
}

func (r *Es6Retrier) Retry(ctx context.Context, retry int, req *http.Request, resp *http.Response, err error) (time.Duration, bool, error) {
	// Fail hard on a specific error
	if err == syscall.ECONNREFUSED {
		return 0, false, errors.New("elasticsearch or network down")
	}

	// Stop after 5 retries
	if retry >= 5 {
		return 0, false, nil
	}

	// Let the backoff strategy decide how long to wait and whether to stop
	wait, stop := r.backoff.Next(retry)
	return wait, stop, nil
}
