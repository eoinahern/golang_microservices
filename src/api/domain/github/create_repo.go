package github

//CreateRepoRequest : object used in request to github api
type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
	Private     bool   `json:"private "`
	HasIssues   bool   `json:"has_issues"`
	HasProjects bool   `json:"has_project"`
	HasWiki     bool   `json:"has_wiki"`
}

//CreateRepoResponse : response
type CreateRepoResponse struct {
	ID          int64           `json:"id"`
	Name        string          `json:"name"`
	Fullname    string          `json:"full_name"`
	Owner       RepoOwner       `json:"owner"`
	Permissions RepoPermissions `json:"permissions"`
}

//RepoOwner : owner of repo
type RepoOwner struct {
	ID      int64  `json:"id"`
	Login   string `json:"login"`
	Url     string `json:"url"`
	HtmlUrl string `json:"html_url`
}

type RepoPermissions struct {
	IsAdmin bool `json:"admin"`
	HasPull bool `json:"pull"`
	HasPush bool `json:"push"`
}

func CreateRepo() {

}
