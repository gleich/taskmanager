package secrets

import (
	"github.com/caarlos0/env/v11"
	"github.com/gleich/lumber/v2"
	"github.com/joho/godotenv"
	"github.com/jomei/notionapi"
)

var SECRETS Secrets

type Secrets struct {
	NotionToken notionapi.Token `env:"NOTION_TOKEN"`

	TasksDB notionapi.DatabaseID `env:"TASKS_DB"`
}

func Load() {
	err := godotenv.Load()
	if err != nil {
		lumber.Fatal(err, "loading .env file failed")
	}
	secrets, err := env.ParseAs[Secrets]()
	if err != nil {
		lumber.Fatal(err, "parsing required env vars failed")
	}
	SECRETS = secrets
	lumber.Success("loaded secrets")
}
