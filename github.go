package conveyor

import (
	"fmt"
	"strings"

	"github.com/google/go-github/github"
)

// NewHook returns a new github.Hook instance that represents the appropriate
// configuration for the Conveyor webhook.
func NewHook(url, secret string) *github.Hook {
	return &github.Hook{
		Events: []string{"push"},
		Active: github.Bool(true),
		Name:   github.String("web"),
		Config: map[string]interface{}{
			"url":          url,
			"content_type": "json",
			"secret":       secret,
		},
	}
}

// GitHubAPI represents an interface for performing Git operations.
type GitHubAPI interface {
	ResolveBranch(owner, repo, branch string) (sha string, err error)
	InstallHook(owner, repo string, hook *github.Hook) error
}

func NewGitHub(c *github.Client) *GitHub {
	return &GitHub{
		Git:          c.Git,
		Repositories: c.Repositories,
	}
}

// GitHub is an implementation of the Git interface
// backed by the GitHub API.
type GitHub struct {
	Git *github.GitService

	Repositories interface {
		CreateHook(owner, repo string, hook *github.Hook) (*github.Hook, *github.Response, error)
		ListHooks(owner, repo string, opt *github.ListOptions) ([]github.Hook, *github.Response, error)
		EditHook(owner, repo string, id int, hook *github.Hook) (*github.Hook, *github.Response, error)
	}
}

func (g *GitHub) ResolveBranch(owner, repo, branch string) (string, error) {
	ref, _, err := g.Git.GetRef(owner, repo, fmt.Sprintf("refs/heads/%s", branch))
	if err != nil {
		return "", err
	}
	return *ref.Object.SHA, nil
}

// InstallHook installs the hook on the repo.
func (g *GitHub) InstallHook(owner, repo string, hook *github.Hook) (err error) {
	var existingHook *github.Hook
	existingHook, err = g.existingHook(owner, repo, hook)
	if err != nil {
		return
	}

	if existingHook != nil {
		if _, _, err = g.Repositories.EditHook(owner, repo, *existingHook.ID, hook); err != nil {
			return
		}
	} else {
		if _, _, err = g.Repositories.CreateHook(owner, repo, hook); err != nil {
			return
		}
	}

	return
}

// existingHook returns an existing hook if it exists.
func (g *GitHub) existingHook(owner, repo string, hook *github.Hook) (*github.Hook, error) {
	hooks, _, err := g.Repositories.ListHooks(owner, repo, nil)
	if err != nil {
		return nil, err
	}

	for _, existingHook := range hooks {
		if equalHooks(&existingHook, hook) {
			return &existingHook, nil
		}
	}

	return nil, nil
}

func equalHooks(a, b *github.Hook) bool {
	if *a.Name == *b.Name {
		if *a.Name == "web" {
			return a.Config["url"].(string) == b.Config["url"].(string)
		}
	}

	return false
}

func splitRepo(fullRepo string) (owner, repo string) {
	parts := strings.Split(fullRepo, "/")
	owner, repo = parts[0], parts[1]
	return
}
