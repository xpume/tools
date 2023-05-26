package main

import (
	"fmt"
	"github.com/ahmetb/go-linq/v3"
	"github.com/golang-module/carbon/v2"
	"github.com/xanzy/go-gitlab"
)

const (
	token   = ""
	baseURL = ""
)

func main() {
	cli, err := gitlab.NewClient(token, gitlab.WithBaseURL(baseURL))
	if err != nil {
		panic(err)
	}

	getMergeRequest(cli, 1)
}

func getMergeRequest(cli *gitlab.Client, page int) {
	startOfWeek := gitlab.ISOTime(carbon.Now().SetWeekStartsAt(carbon.Monday).StartOfWeek().ToStdTime())
	listOptions := gitlab.ListOptions{Page: page, PerPage: 100}
	events, response, err := cli.Events.ListCurrentUserContributionEvents(
		&gitlab.ListContributionEventsOptions{After: &startOfWeek, ListOptions: listOptions})
	if err != nil {
		panic(err)
	}

	printEvents(events)

	if response.NextPage != 0 {
		getMergeRequest(cli, response.NextPage)
	}
}

func printEvents(events []*gitlab.ContributionEvent) {
	var contributionEvents []linq.Group
	linq.From(events).
		Where(func(e interface{}) bool {
			return e.(*gitlab.ContributionEvent).TargetType == string(gitlab.TodoTargetMergeRequest)
		}).
		GroupBy(func(e interface{}) interface{} {
			return e.(*gitlab.ContributionEvent).TargetTitle
		}, func(e interface{}) interface{} {
			return e.(*gitlab.ContributionEvent).TargetTitle
		}).ToSlice(&contributionEvents)

	for _, event := range contributionEvents {
		fmt.Println(event.Key.(string))
	}
}
