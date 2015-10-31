package libs

import (
	"fmt"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type GithubRepo struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Fullname    *string `json:"Fullname,omitempty"`
	Branch      *string `json:"Branch,omitempty"`
	Url         *string `json:"Url,omitempty"`
	Language    *string `json:"Language,omitempty"`
	Forks       *int    `json:"Forks,omitempty"`
	Stars       *int    `json:"Stars,omitempty"`
	Watches     *int    `json:"Watches,omitempty"`
	Badge       *string `json:"Badge,omitempty"`
}

type Github struct {
	client *github.Client
	config *Config
}

func NewGithub(config *Config) *Github {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: config.Github.Token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	return &Github{
		client: github.NewClient(tc),
		config: config,
	}
}

func (g *Github) GetRepos() ([]GithubRepo, error) {
	opt := &github.RepositoryListOptions{
		Type:      "owner",
		Sort:      "pushed",
		Direction: "desc",
	}
	repos, _, err := g.client.Repositories.List(g.config.Github.Username, opt)
	if err != nil {
		return nil, err
	}

	var gr []GithubRepo
	for _, repo := range repos {
		badge := fmt.Sprintf(
			"https://api.travis-ci.org/%s.svg?branch=%s",
			*repo.FullName,
			*repo.DefaultBranch,
		)
		gr = append(gr, GithubRepo{
			Name:        repo.Name,
			Description: repo.Description,
			Fullname:    repo.FullName,
			Branch:      repo.DefaultBranch,
			Url:         repo.HTMLURL,
			Language:    repo.Language,
			Forks:       repo.ForksCount,
			Stars:       repo.StargazersCount,
			Watches:     repo.SubscribersCount,
			Badge:       &badge,
		})
	}

	return gr, nil
}