package trading

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	marketauxEndpoint = "https://api.marketaux.com/v1/news/all"
	// marketauxEndpoint = "https://api.marketaux.com/v1/entity/stats/aggregation"
	apiToken       = "0tyWPUQbIwjs80ihYeJnfB0qO1iAvHgHhSwkCVen" // replace with your actual token
	country        = "in"
	maxNewsPerCall = 50 // you can adjust
)

// NewsItem represents structure from Marketaux response (partial)
type NewsItem struct {
	UUID        string    `json:"uuid"`
	Title       string    `json:"title"`
	Snippet     string    `json:"snippet"`
	Description string    `json:"description"`
	Url         string    `json:"url"`
	Entities    []Entity  `json:"entities"`
	PublishedAt time.Time `json:"published_at"`
	// add more fields if needed
}

type Entity struct {
	Symbol         string  `json:"symbol"`
	Name           string  `json:"name"`
	MatchScore     float64 `json:"match_score"`
	SentimentScore float64 `json:"sentiment_score"`
}

type MarketauxResponse struct {
	Meta struct {
		Found    int `json:"found"`
		Returned int `json:"returned"`
	} `json:"meta"`
	Data []NewsItem `json:"data"`
}

func fetchIndianMarketNews(since time.Time) ([]NewsItem, error) {
	q := url.Values{}
	q.Set("countries", country)
	q.Set("api_token", apiToken)
	q.Set("language", "en")
	q.Set("sentiment_gte", "0.35")
	// q.Set("min_match_score", "70.0")
	q.Set("symbol", "HDFCBANK.BO, TCS.BO, BEL.BO, POWERGRID.NS, WIPRO.BO, HCL.BO")
	// q.Set("limit", fmt.Sprintf("%d", maxNewsPerCall))
	// filter articles published after 'since'
	// q.Set("published_after", since.Format(time.RFC3339))
	q.Set("published_after", "2025-12-02T01:00")
	q.Set("sort", "published_at")

	reqURL := marketauxEndpoint + "?" + q.Encode()

	resp, err := http.Get(reqURL)
	if err != nil {
		return nil, err
	}
	body1, err := ioutil.ReadAll(resp.Body)
	fmt.Println("API Response:", string(body1), err)

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status NOT OK: %d", resp.StatusCode)
	}

	// body, err := ioutil.ReadAll(resp.Body)
	// fmt.Println("API Response11111111:", string(body))
	// if err != nil {
	// 	return nil, err
	// }

	var mr MarketauxResponse
	if err := json.Unmarshal(body1, &mr); err != nil {
		return nil, err
	}

	return mr.Data, nil
}

// func main() {
// 	lastFetch := time.Now()
// 	news, err := fetchIndianMarketNews(lastFetch)
// 	if err != nil {
// 		fmt.Println("Error fetching news:", err)
// 	} else {
// 		for _, n := range news {
// 			// fmt.Printf("[%s] %s\n%s\nURL: %s\n\n",
// 			// 	n.PublishedAt.Format("2006-01-02 15:04:05"),
// 			// 	n.Title,
// 			// 	n.Description,
// 			// 	n.Url,
// 			// )

// 			fmt.Printf("News--------- %+v", n)
// 			fmt.Println("-----------------------------------------------------------------------------------")
// 			// Here you can add sentiment filtering, store in DB, or trigger algo logic

// 		}
// 	}

// 	// // Example: poll every 5 minutes between 9:15 and 20:00
// 	// lastFetch := time.Now().Add(-10 * time.Minute)

// 	// ticker := time.NewTicker(5 * time.Minute)
// 	// defer ticker.Stop()

// 	// for {
// 	// 	now := time.Now()
// 	// 	if now.Hour() < 20 {
// 	// 		news, err := fetchIndianMarketNews(lastFetch)
// 	// 		if err != nil {
// 	// 			fmt.Println("Error fetching news:", err)
// 	// 		} else {
// 	// 			for _, n := range news {
// 	// 				fmt.Printf("[%s] %s\n%s\nURL: %s\n\n",
// 	// 					n.PublishedAt.Format("2006-01-02 15:04:05"),
// 	// 					n.Title,
// 	// 					n.Description,
// 	// 					n.Url,
// 	// 				)
// 	// 				// Here you can add sentiment filtering, store in DB, or trigger algo logic
// 	// 			}
// 	// 		}
// 	// 		lastFetch = now
// 	// 	}
// 	// 	<-ticker.C
// 	// }
// }
