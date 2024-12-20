package main

import (
	"context"
	"fmt"

	"github.com/blexram-go/gator-go/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("%s takes an argument", cmd.Name)
	}

	feedURL := cmd.Args[0]
	feedToUnfollow, err := s.db.GetFeed(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("unable to get feed: %w", err)
	}

	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feedToUnfollow.ID,
	})
	if err != nil {
		return fmt.Errorf("unable to unfollow feed: %w", err)
	}

	fmt.Println("Unfollowed feed!")
	return nil
}
