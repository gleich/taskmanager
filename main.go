package main

import (
	"context"
	"time"

	"github.com/gleich/lumber/v2"
	"github.com/gleich/taskmanager/internal/secrets"
	"github.com/jomei/notionapi"
)

func main() {
	secrets.Load()
	client := notionapi.NewClient(secrets.SECRETS.NotionToken)

	threeMonthsAgo := notionapi.Date(time.Now().AddDate(0, -3, 0))

	db, err := client.Database.Query(
		context.Background(),
		secrets.SECRETS.TasksDB,
		&notionapi.DatabaseQueryRequest{
			Filter: notionapi.AndCompoundFilter{
				notionapi.PropertyFilter{
					Property: "When",
					Date: &notionapi.DateFilterCondition{
						OnOrAfter: &threeMonthsAgo,
					},
				},
				notionapi.PropertyFilter{
					Property: "Done",
					Checkbox: &notionapi.CheckboxFilterCondition{Equals: true},
				},
			},
		},
	)
	if err != nil {
		lumber.Fatal(err, "getting database query for tasks failed")
	}

	for _, result := range db.Results {
		lumber.Debug(result.Properties)
	}
}
