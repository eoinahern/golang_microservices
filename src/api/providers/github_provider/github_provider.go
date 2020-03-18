package github_provider

import (
	"fmt"

	"github.com/eoinahern/golang_microservices/microservice_project/src/api/domain/github"
)

const (
	headerAuthorization = "Authorization"
	headerAuthFormat    = "token %s"
)

func getAuthHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthFormat, accessToken)
}

//CreateRepo call external api to create repo
func CreateRepo(request github.CreateRepoRequest) (*github.CreateRepoResponse, *github.GithubErrorResponse) {
	return &github.CreateRepoResponse{}, nil
}
