package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type APIMethod[RESP, BODY any] func(BODY) (*APIResponse[RESP], error)

type APIResponse[T any] struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Body    string `json:"body"`
}

func (a *APIResponse[T]) UnwrapOr(defaultValue *T, catch func(err error)) *T {
	t, err := a.ParseBody()
	if err != nil {
		catch(err)
		return defaultValue
	}
	return t
}

func (a *APIResponse[T]) Ok() bool {
	return a.Status == http.StatusOK
}

func (a *APIResponse[T]) InternalError() bool {
	return a.Status == http.StatusInternalServerError
}

func (a *APIResponse[T]) ParseBody() (*T, error) {
	var v T
	if err := json.Unmarshal([]byte(a.Body), &v); err != nil {
		return nil, err
	}
	return &v, nil
}

type API interface {
	CallData() CallData
}

type DefaultAPI struct {
	API
}

var itemApi ItemsAPI
var surveyApi SurveyAPI

func GetItemAPI() ItemsAPI {
	return itemApi
}

func GetSurveyAPI() SurveyAPI {
	return surveyApi
}

func Setup() {
	itemApi = &ItemsAPIImpl{}
	surveyApi = &SurveyAPIImpl{}
}

func MustCallAndUnwrap[T, U any](method APIMethod[T, U], in U, then func(*T), catchInternal func(error), catchAPI func(*APIResponse[T])) {
	resp, err := method(in)
	fmt.Println(1)
	if err != nil {
		catchInternal(err)
		return
	}
	fmt.Println(2)
	body, err := resp.ParseBody()
	if err != nil {
		catchInternal(err)
		return
	}
	fmt.Println(3)
	if !resp.Ok() {
		catchAPI(resp)
		return
	}
	then(body)
}

func DefaultHeaders(r *http.Request, d CallData) {
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", d.Token)
}
