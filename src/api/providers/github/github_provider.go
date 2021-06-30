package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	rc "github.com/tahkiu/golang-microservices/src/api/clients/rest_client"
	"github.com/tahkiu/golang-microservices/src/api/domain/github"
)

const (
	headerAuthorization       = "Authorization"
	headerAuthorizationFormat = "token %s"
	urlCreateRepo             = "https://api.github.com/user/repos"
)

func getAuthorizationHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthorizationFormat, accessToken)
}

func CreateRepo(accessToken string, req *github.CreateRepoParam) (*github.CreateRepoResp, *github.GithubErrorResp) {
	// Authorization: token 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
	headers := http.Header{}
	headers.Set(headerAuthorization, getAuthorizationHeader(accessToken))
	resp, err := rc.Post(urlCreateRepo, req, headers)
	fmt.Println(resp)
	fmt.Println(err)
	if err != nil {
		log.Printf("error when creating new repo in github: %s\n", err.Error())
		return nil, &github.GithubErrorResp{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, &github.GithubErrorResp{
			StatusCode: http.StatusInternalServerError,
			Message:    "invalid response body",
		}
	}
	if resp.StatusCode > 299 {
		var errResp github.GithubErrorResp
		if err := json.Unmarshal(bytes, &errResp); err != nil {
			// different github error api, wrong struct maybe
			return nil, &github.GithubErrorResp{
				StatusCode: http.StatusInternalServerError,
				Message:    "invalid json response body",
			}
		}
		errResp.StatusCode = int64(resp.StatusCode)
		// correct err response from github
		return nil, &errResp
	}

	var result github.CreateRepoResp
	if err := json.Unmarshal(bytes, &result); err != nil {
		log.Printf("error when trying to unmarshal github create repo response: %s\n", err.Error())
		return nil, &github.GithubErrorResp{
			StatusCode: http.StatusInternalServerError,
			Message:    "error when trying to unmarshal github create repo response",
		}
	}

	return &result, nil
}
