package database

import (
	"context"
	"errors"
	"net/http"
	"syscall"
	"time"

	elastic6 "github.com/olivere/elastic/v6"
	"github.com/olivere/elastic/v7"
)

func ElasticConn(address ...string) *elastic.Client {
	esClient, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(address...),
		elastic.SetRetrier(NewEsRetrier()),
	)
	if err != nil {
		panic(err)
	}
	return esClient
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

func Elastic6Conn(address ...string) *elastic6.Client {
	esClient, err := elastic6.NewClient(
		elastic6.SetSniff(false),
		elastic6.SetURL(address...),
		elastic6.SetRetrier(NewEs6Retrier()),
	)
	if err != nil {
		panic(err)
	}
	return esClient
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
