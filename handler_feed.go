package main

import (
	"context"
	"fmt"
)

func handlerFetchFeed(s *state, cmd command) error {
	fullURL := "https://www.wagslane.dev/index.xml"

	if len(cmd.Args) != 0 {
		return fmt.Errorf("%s does not take any arguments", cmd.Name)
	}

	rssFeed, err := fetchFeed(context.Background(), fullURL)
	if err != nil {
		return fmt.Errorf("error fetching the feed: %w", err)
	}

	fmt.Printf("Feed: %+v\n", rssFeed)
	return nil
}
