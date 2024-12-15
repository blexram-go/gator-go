package main

import (
	"context"
	"fmt"
	"time"

	"github.com/blexram-go/gator-go/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("%s takes an argument", cmd.Name)
	}

	currentUserName := s.cfg.CurrentUserName
	feedURL := cmd.Args[0]

	currentUser, err := s.db.GetUser(context.Background(), currentUserName)
	if err != nil {
		return fmt.Errorf("unable to get user: %w", err)
	}

	currentFeed, err := s.db.GetFeed(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("unable to get feed: %w", err)
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:    currentFeed.ID,
		UserID:    currentUser.ID,
	})
	if err != nil {
		return fmt.Errorf("unable to create feed follow: %w", err)
	}

	fmt.Printf("Feed Name: %s\n", feedFollow.FeedName)
	fmt.Printf("User: %s\n", feedFollow.UserName)
	return nil
}

func handlerGetFeedFollows(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("%s does not take any arguments", cmd.Name)
	}

	currentUserName := s.cfg.CurrentUserName
	currentUser, err := s.db.GetUser(context.Background(), currentUserName)
	if err != nil {
		return fmt.Errorf("unable to get user: %w", err)
	}

	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), currentUser.ID)
	if err != nil {
		return fmt.Errorf("unable to get followed feeds: %w", err)
	}

	fmt.Printf("Followed Feeds for %s\n", currentUserName)
	for _, feedFollow := range feedFollows {
		fmt.Printf("-     %s\n", feedFollow.FeedName)
	}
	return nil
}
