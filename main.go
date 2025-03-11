package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/natretsel/blog_aggregator/internal/config"
	"github.com/natretsel/blog_aggregator/internal/database"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Print(err)
	}
	db, err := sql.Open("postgres", cfg.Db_url)
	if err != nil {
		fmt.Println("error opening database: ", err)
	}

	dbQueries := database.New(db)
	st := state{
		db:  dbQueries,
		cfg: &cfg,
	}
	cmds := commands{
		commandMap: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	// Read commands from user

	words := os.Args
	/*
		fmt.Println("Length of supplied args: ", len(words))
		for idx, val := range words {
			fmt.Printf("Index: %v, Value: %v\n", idx, val)
		}
	*/

	if len(words[1:]) == 0 {
		fmt.Println("please provide arguments")
		os.Exit(1)
	}
	commandName := words[1]

	/*
		if commandName != "reset" {
			if len(words[1:]) < 2 {
				fmt.Println("two arguments required. Only one received.")
				os.Exit(1)
			}
		}
	*/

	args := words[2:]

	cmd := command{
		name: commandName,
		args: args,
	}
	err = cmds.run(&st, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)

}
