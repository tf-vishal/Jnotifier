package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/tf-vishal/JNotifier/internal/models"
)

func SendJobNotification(job models.Job) {
	webhookURL := os.Getenv("DISCORD_WEBHOOK")
	fmt.Println("Webhook URL:", os.Getenv("DISCORD_WEBHOOK"))

	message := map[string]string{
		"content": "📢 **" + job.Title + "** at **" + job.Company + "** (" + job.Location + ")\n" + job.Link,
	}

	body, _ := json.Marshal(message)
	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Println("❌ Failed to send message to Discord:", err)
		return
	}
	defer resp.Body.Close()

	log.Println("✅ Sent message to Discord with status:", resp.Status)

}
