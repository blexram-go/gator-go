package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/blexram-go/gator-go/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("%s expects a single argument, the username", cmd.Name)
	}

	name := cmd.Args[0]

	user, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return err
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <name>", cmd.Name)
	}

	name := cmd.Args[0]

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
	})
	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User successfully created!")
	log.Println(user)
	return nil
}

func handlerGetUsers(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("%s does not take any arguments", cmd.Name)
	}

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	if len(users) == 0 {
		fmt.Println("No users in the database")
		return nil
	}

	currentUser := s.cfg.CurrentUserName

	for _, user := range users {
		if currentUser == user.Name {
			fmt.Printf("* %s (current)\n", currentUser)
			continue
		}
		fmt.Printf("* %s\n", user.Name)
	}
	return nil
}
