package github

type GithubErrorResponse struct {
	Message          string        `json:"message"`
	DocumentationUrl string        `json:"documentation_url"`
	Errors           []GithubError `json:"errors"`
}

type GithubError struct {
	Resource string `json:"resource"`
	Code     string `json:"code"`
	Field    string `json:"field"`
	Message  string `json:"message"`
}
