package main

import (
	"context"
	"fmt"
	"time"

	"github.com/blexram-go/gator-go/internal/database"
	"github.com/google/uuid"
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

func handlerAddFeed(s *state, cmd command) error {
	currentUser := s.cfg.CurrentUserName
	user, err := s.db.GetUser(context.Background(), currentUser)
	if err != nil {
		return err
	}

	if len(cmd.Args) != 2 {
		return fmt.Errorf("%s takes two arguments", cmd.Name)
	}

	rssFeedName := cmd.Args[0]
	rssFeedURL := cmd.Args[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      rssFeedName,
		Url:       rssFeedURL,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}

	fmt.Println("Feed successfully created and added to user!")
	fmt.Printf("%+v\n", feed)
	return nil
}
