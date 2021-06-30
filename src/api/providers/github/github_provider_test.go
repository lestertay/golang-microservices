package github

import (
	"errors"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	rc "github.com/tahkiu/golang-microservices/src/api/clients/rest_client"
	"github.com/tahkiu/golang-microservices/src/api/domain/github"
)

const (
	testUrlCreateRepo = "https://api.github.com/user/repos"
)

func TestMain(m *testing.M) {
	rc.StartMockups()
	defer rc.StopMockups()
	os.Exit(m.Run())
}
func TestGetAuthorizationHeader(t *testing.T) {
	header := getAuthorizationHeader("abc123")
	assert.EqualValues(t, "token abc123", header)
}

func TestCreateRepoErrorRestclient(t *testing.T) {
	rc.FlushMockups()
	rc.AddMockup(rc.Mock{
		Url:        testUrlCreateRepo,
		HttpMethod: http.MethodPost,
		Err:        errors.New("invalid rest_client response"),
	})
	resp, err := CreateRepo("", &github.CreateRepoParam{})
	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, "invalid rest_client response", err.Message)
}

func TestCreateRepoInvalidResponseBody(t *testing.T) {
	rc.FlushMockups()
	invalidCloser, _ := os.Open("-asf3")
	rc.AddMockup(rc.Mock{
		Url:        testUrlCreateRepo,
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       invalidCloser,
		},
	})
	resp, err := CreateRepo("", &github.CreateRepoParam{})
	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "invalid response body", err.Message)
}
