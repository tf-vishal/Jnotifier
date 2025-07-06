package discord

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"

	"github.com/tf-vishal/JNotifier/internal/models"
)

func SendJobNotifcation(job models.Job) error {
	webhoookURL := os.Getenv("DISCORD_WEEKHOOK")

	message := map[string]string{
		"content": "📢 **" + job.Title + "** at **" + job.Company + "** (" + job.Location + ")\n" + job.Link,
	}

	body, _ := json.Marshal(message)
	http.Post(webhoookURL, "application/json", bytes.NewBuffer(body))
}
