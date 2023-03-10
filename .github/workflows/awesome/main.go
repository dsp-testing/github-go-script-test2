package main

import (
	"fmt"

	github_go_script "github.com/dsp-testing/github-go-script"
)

func main() {
	github_go_script.Call(func(options *github_go_script.Options) (map[string]string, error) {
		// TODO: Replace this with your code
		repo, _, err := options.Client.Repositories.Get(options.Ctx, options.Context.Repository.OwnerName, options.Context.Repository.Name)
		if err != nil {
			return nil, err
		}
		fmt.Printf("Repository name: %s\n", repo.GetName())
		fmt.Printf("Triggering event name: %s\n", options.Context.EventName)
		return map[string]string{
			"repository_name": repo.GetName(),
			"event_name":      options.Context.EventName,
		}, nil
	})
}
