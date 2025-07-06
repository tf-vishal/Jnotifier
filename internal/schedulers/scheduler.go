package schedulers

import (
	"log"
	"time"

	"github.com/tf-vishal/JNotifier/internal/api"
	"github.com/tf-vishal/JNotifier/internal/discord"
)

func StartJobScheduler() {
	ticker := time.NewTicker(30 * time.Minute)

	go func() {
		for range ticker.C {
			checkJobs()
		}
	}()

	checkJobs()
}

func checkJobs() {
	keywords := []string{"Network Engineer", "System Administrator", "DevOps", "Cloud", "CCNA"}
	for _, keyword := range keywords {
		go func(kw string) {
			jobs := api.FetchFromJsearch(kw)
			for _, job := range jobs {
				discord.SendJobNotification(job)
			}
		}(keyword)
	}
	log.Println("Job check completed for keywords:", keywords)
}
