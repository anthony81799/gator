package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/anthony81799/gator/internal/config"
	"github.com/anthony81799/gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	var s state
	c, err := config.Read()
	if err != nil {
		log.Fatalln(err)
	}
	s.cfg = &c

	db, err := sql.Open("postgres", s.cfg.DBURL)
	if err != nil {
		log.Fatalln(err)
	}
	s.db = database.New(db)

	var cmds commands
	cmds.handlers = map[string]func(*state, command) error{}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))

	arguments := os.Args
	if len(arguments) < 2 {
		log.Fatal("this program needs at least 2 arguments")
	}

	cmd := command{
		name: arguments[1],
		args: arguments[2:],
	}

	err = cmds.run(&s, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
