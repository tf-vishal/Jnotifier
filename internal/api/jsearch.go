package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/tf-vishal/JNotifier/internal/models"
)

func FetchFromJsearch() (keyword string) []models.Job{
	url := fmt.Sprintf("https://jsearch.p.rapidapi.com/search?query=%s&num_pages=1", keyword)

	req., _ := http.NewRequest("GET", url, nil)
	req.Header.Add("X-RapidAPI-Key", os.Getenv("RAPIDAPI_KEY"))
	req.Header.Add("X-RapidAPI-Host", "jsearch.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error fetching data from Jsearch:", err)
		return nil
	}

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	car result struct {
		Data []struct {
			JobTitle    string `json:"job_title"`
			Employer string `json:"employer_name"`
			Location string `json:"location"`
			JobLink string `json:"job_apply_link"`
		} `json:"data"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		log.Println("Error unmarshalling JSON:", err)
		return nil
	}

	var jobs []models.Job
	for _, item := range result.Data {
		jobs = append (job,models.Job{
			Title:    item.JobTitle,
			Company:  item.Employer,
			Link:     item.JobLink,
			Location: item.Location,
			Source:   "Jsearch",
		})
	}

	log.Printf("Found %d jobs for: %s", len(jobs), keyword)
	return jobs
}
