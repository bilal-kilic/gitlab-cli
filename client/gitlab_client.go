package client

import (
	"encoding/json"
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

func (client *GitlabApiClient) GetToken(username string, password string) models.GetTokenResponse {
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
		fmt.Printf("An error occurred trying to get Oauth token", err)
	}

	var responseBody = models.GetTokenResponse{}
	_ = json.Unmarshal(response.Body(), &responseBody)

	return responseBody
}

func (client *GitlabApiClient) GetProjects() []models.ProjectResponse {
	page := 1
	firstPage, hasNextPage := client.getProjectsPage(1)
	result := append(firstPage)
	for hasNextPage {
		page++
		var currentPage []models.ProjectResponse
		currentPage, hasNextPage = client.getProjectsPage(page)
		result = append(result, currentPage...)
	}

	return result
}

func (client *GitlabApiClient) getProjectsPage(page int) ([]models.ProjectResponse, bool) {
	request := fasthttp.AcquireRequest()
	request.Header.SetContentType("application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", client.config.Token))
	request.SetRequestURI(fmt.Sprintf("%s/api/v4/projects?simple=true&page=%d", client.config.Url, page))
	request.Header.SetMethod("GET")

	response := fasthttp.AcquireResponse()
	if err := fasthttp.Do(request, response); err != nil {
		fmt.Printf("An error occurred trying to get Oauth token", err)
	}

	var responseBody []models.ProjectResponse
	_ = json.Unmarshal(response.Body(), &responseBody)

	if response.Header.Peek("X-Next-Page") == nil {
		return responseBody, false
	} else {
		return responseBody, true
	}
}

func (client *GitlabApiClient) GetProjectBranches(projectId int64) []models.GetBranchResponse {
	request := fasthttp.AcquireRequest()
	request.Header.SetContentType("application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", client.config.Token))
	request.SetRequestURI(fmt.Sprintf("%s/api/v4/projects/%d/repository/branches", client.config.Url, projectId))
	request.Header.SetMethod("GET")

	response := fasthttp.AcquireResponse()
	if err := fasthttp.Do(request, response); err != nil {
		fmt.Printf("An error occurred trying to get branches token", err)
	}

	var responseBody []models.GetBranchResponse
	_ = json.Unmarshal(response.Body(), &responseBody)

	return responseBody
}

func (client *GitlabApiClient) CreateMergeRequest(id int64, createMergeRequest models.CreateMergeRequest) models.CreateMergeRequestResponse {
	reqJson, _ := json.Marshal(createMergeRequest)

	request := fasthttp.AcquireRequest()
	request.Header.SetContentType("application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", client.config.Token))
	request.SetRequestURI(fmt.Sprintf("%s/api/v4/projects/%d/merge_requests", client.config.Url, id))
	request.Header.SetMethod("POST")
	request.SetBody(reqJson)

	response := fasthttp.AcquireResponse()
	if err := fasthttp.Do(request, response); err != nil {
		fmt.Printf("An error occurred trying to create merge request. Error: ", err)
	}

	if response.StatusCode() != http.StatusCreated {
		fmt.Printf("An error occurred trying to create merge request. Error: ", response.Body())
	}

	var responseBody = models.CreateMergeRequestResponse{}
	_ = json.Unmarshal(response.Body(), &responseBody)

	return responseBody
}
