package main

import (
	"fmt"
	"log"
	"os"

	"github.com/blexram-go/gator-go/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	programState := &state{
		cfg: &cfg,
	}

	cmds := commands{
		cmdNames: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		fmt.Println("Usage: cli <command> [args...]")
		os.Exit(1)
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]
	newCmd := command{
		Name: cmdName,
		Args: cmdArgs,
	}
	err = cmds.run(programState, newCmd)
	if err != nil {
		log.Fatal(err)
	}
}
