package utils

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

type ResponseData[T any] struct {
	Status string
	Data   T
}

func Post[T any](path string, requestData any) (interface{}, error) {
	client := resty.New()

	resp, err := client.
		SetBaseURL("http://localhost:8000").
		SetRetryCount(2).
		SetRetryWaitTime(5*time.Second).
		R().
		SetHeader("AppSvcName", "appSvc1").
		SetBody(requestData).
		SetResult(&ResponseData[T]{}).
		Post(path)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		log.Println(resp)
		return nil, errors.New(resp.Status())
	}

	return resp.Result(), nil
}
