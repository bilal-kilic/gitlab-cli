package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
	"gitlab-cli/cmd/models"
	"gitlab-cli/configuration"
	"net/http"
	"time"
)

type GitlabApiClient struct {
	config *configuration.GitlabConfiguration
	client *fasthttp.Client
}

func NewGitlabClient(config *configuration.GitlabConfiguration) *GitlabApiClient {
	return &GitlabApiClient{
		config: config,
		client: &fasthttp.Client{
			MaxConnsPerHost:     10,
			MaxIdleConnDuration: time.Millisecond * 5000,
			ReadTimeout:         time.Millisecond * 500,
			WriteTimeout:        time.Millisecond * 500,
		},
	}
}

func (client *GitlabApiClient) GetToken(username string, password string) (models.GetTokenResponse, error) {
	getTokenRequest := models.GetTokenRequest{
		GrantType: "password",
		UserName:  username,
		Password:  password,
	}
	reqJson, _ := json.Marshal(getTokenRequest)

	request := fasthttp.AcquireRequest()
	request.Header.SetContentType("application/json")
	request.SetRequestURI(fmt.Sprintf("%s/oauth/token", client.config.Url))
	request.Header.SetMethod("POST")
	request.SetBody(reqJson)

	response := fasthttp.AcquireResponse()
	if err := fasthttp.Do(request, response); err != nil {
		return models.GetTokenResponse{}, err
	}

	if response.StatusCode() != http.StatusOK {
		return models.GetTokenResponse{}, errors.New(fmt.Sprintf("%s", response.Body()))
	}

	var responseBody = models.GetTokenResponse{}
	if err := json.Unmarshal(response.Body(), &responseBody); err != nil {
		return models.GetTokenResponse{}, err
	}

	return responseBody, nil
}

func (client *GitlabApiClient) GetProjects() ([]models.ProjectResponse, error) {
	page := 1
	firstPage, hasNextPage, err := client.getProjectsPage(1)
	result := append(firstPage)
	for hasNextPage {
		if err != nil {
			return nil, err
		}

		page++
		var currentPage []models.ProjectResponse
		currentPage, hasNextPage, err = client.getProjectsPage(page)
		result = append(result, currentPage...)
	}

	return result, nil
}

func (client *GitlabApiClient) getProjectsPage(page int) ([]models.ProjectResponse, bool, error) {
	request := fasthttp.AcquireRequest()
	request.Header.SetContentType("application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", client.config.Token))
	request.SetRequestURI(fmt.Sprintf("%s/api/v4/projects?simple=true&page=%d", client.config.Url, page))
	request.Header.SetMethod("GET")

	response := fasthttp.AcquireResponse()
	if err := fasthttp.Do(request, response); err != nil {
		return nil, false, err
	}

	var responseBody []models.ProjectResponse
	if err := json.Unmarshal(response.Body(), &responseBody); err != nil {
		return nil, false, err
	}

	if response.Header.Peek("X-Next-Page") == nil {
		return responseBody, false, nil
	} else {
		return responseBody, true, nil
	}
}

func (client *GitlabApiClient) GetProjectBranches(projectId int64) ([]models.GetBranchResponse, error) {
	request := fasthttp.AcquireRequest()
	request.Header.SetContentType("application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", client.config.Token))
	request.SetRequestURI(fmt.Sprintf("%s/api/v4/projects/%d/repository/branches", client.config.Url, projectId))
	request.Header.SetMethod("GET")

	response := fasthttp.AcquireResponse()
	if err := fasthttp.Do(request, response); err != nil {
		return nil, err
	}

	var responseBody []models.GetBranchResponse
	if err := json.Unmarshal(response.Body(), &responseBody); err != nil {
		return nil, err
	}

	return responseBody, nil
}

func (client *GitlabApiClient) CreateMergeRequest(id int64, createMergeRequest models.CreateMergeRequest) (models.CreateMergeRequestResponse, error) {
	reqJson, _ := json.Marshal(createMergeRequest)

	request := fasthttp.AcquireRequest()
	request.Header.SetContentType("application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", client.config.Token))
	request.SetRequestURI(fmt.Sprintf("%s/api/v4/projects/%d/merge_requests", client.config.Url, id))
	request.Header.SetMethod("POST")
	request.SetBody(reqJson)

	response := fasthttp.AcquireResponse()
	if err := fasthttp.Do(request, response); err != nil {
		return models.CreateMergeRequestResponse{}, err
	}

	if response.StatusCode() != http.StatusCreated {
		return models.CreateMergeRequestResponse{}, errors.New(fmt.Sprintf("%q", response.Body()))
	}

	var responseBody = models.CreateMergeRequestResponse{}
	if err := json.Unmarshal(response.Body(), &responseBody); err != nil {
		return models.CreateMergeRequestResponse{}, err
	}

	return responseBody, nil
}
