package cli

import (
	"flag"
	"fmt"

	"github.com/ddvk/rmfakecloud/internal/config"
	"github.com/ddvk/rmfakecloud/internal/model"
	"github.com/ddvk/rmfakecloud/internal/storage/fs"
	log "github.com/sirupsen/logrus"
)

// ListUsers lists
func (cli *Cli) ListUsers(args []string) {
	users, err := cli.storage.GetUsers()
	if err != nil {
		log.Fatal(err)
	}
	for _, u := range users {
		fmt.Print(u.ID)
		if u.IsAdmin {
			fmt.Println("\tadmin")
		} else {
			fmt.Println()

		}
	}
	return
}

// SetUser updates or creates the users if not exists
func (cli *Cli) SetUser(args []string) {
	userParam := flag.NewFlagSet("adduser", flag.ExitOnError)
	username := userParam.String("u", "", "username")
	pass := userParam.String("p", "", "password")
	admin := userParam.Bool("a", false, "isadmmin")

	userParam.Parse(args)
	if *username == "" {
		userParam.PrintDefaults()
		return
	}

	usr, err := cli.storage.GetUser(*username)
	if err != nil {
		if *pass == "" {
			*pass, err = model.GenPassword()
			if err != nil {
				log.Fatal(err)
			}
			log.Info("new password:", *pass)
		}
		usr, err = model.NewUser(*username, *pass)
		if err != nil {
			log.Error(err)
			return
		}
	}
	if *pass != "" {
		usr.SetPassword(*pass)
	}
	usr.IsAdmin = *admin
	err = cli.storage.UpdateUser(usr)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Updated/created the user")
}

// Cli cli interface
type Cli struct {
	storage *fs.Storage
}

// New creates
func New(cfg *config.Config) *Cli {
	storage := &fs.Storage{
		Cfg: cfg,
	}
	return &Cli{
		storage: storage,
	}

}

// Handle handles the args
func (cli *Cli) Handle(args []string) bool {
	if len(args) > 1 {
		cmd := args[1]
		otherarg := args[2:]
		switch cmd {
		case "setuser":
			cli.SetUser(otherarg)
		case "listusers":
			cli.ListUsers(otherarg)
		case "rmuser":
		default:
			log.Warn("unknown command: ", cmd)
		}
		return true
	}

	return false
}
