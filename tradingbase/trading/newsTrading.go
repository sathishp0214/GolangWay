package trading

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/jonreiter/govader"
	"github.com/ledongthuc/pdf"
	"github.com/mmcdole/gofeed"
)

// Config
const (
	startFeedHour     = 9
	startFeedMinute   = 30
	feedURL           = "https://nsearchives.nseindia.com/content/RSS/Online_announcements.xml"
	pollInterval      = 2 * time.Minute  // how often to poll the RSS feed
	httpTimeout       = 30 * time.Second // HTTP client timeout
	maxConcurrentJobs = 4                // how many PDFs to fetch+process concurrently
	maxRetries        = 3                // retry count for transient failures
	retryBaseDelay    = 2 * time.Second  // base backoff
	userAgent         = "nse-rss-sentiment-bot/1.0"
)

var NotifiedFeeds = map[string]string{}

// in-memory seen map to avoid re-processing items (persist to DB for production)
var seenMu sync.Mutex
var seen = map[string]time.Time{} // map[itemLink] => time processed

// semaphore to limit concurrency
var sem = make(chan struct{}, maxConcurrentJobs)

// HTTP client
var client = &http.Client{
	Timeout: httpTimeout,
}

// func main() {
// 	log.Printf("Starting NSE RSS feeds---------------")

// 	location, err := time.LoadLocation("Asia/Kolkata") // IST zone
// 	if err != nil {
// 		panic(err)
// 	}

// 	now := time.Now().In(location)

// 	year, month, day := now.Date()

// 	initialTime := time.Date(year, month, day, startFeedHour, startFeedMinute, 0, 0, location)

// 	for {

// 		lastProcessed := feedReader(initialTime)
// 		fmt.Println("Completed Checking for the feeds---------", time.Now(), "---------", lastProcessed)

// 		time.Sleep(1 * time.Minute)

// 		initialTime = lastProcessed

// 	}

// }

func feedReader(initialTime time.Time) time.Time {
	var err error
	// Layout for parsing the above format
	layout := "02-Jan-2006 15:04:05 -0700 MST"
	fp := gofeed.NewParser()
	var feed *gofeed.Feed

	// feed, err := fp.ParseString(feedURL)

	ctx, cancel := context.WithTimeout(context.Background(), 40*time.Second)
	defer cancel()

	val := func() time.Time {
		//defer with anonymous function for recover
		defer func() {
			r := recover() //This alone recovers from panic() function, Then function which called panic occured function, Will continue to run.
			if r != nil {
				fmt.Println("Panic recoverd from the news feed") //This used for logging, Whether recover is happened or not
			}
		}()

		feed, err = fp.ParseURLWithContext(feedURL, ctx)
		if err != nil {
			panic(fmt.Sprintf("ERROR: parsing feed: %v ---------feeds: %v\n", err, feed))
		}

		var lastProcessed time.Time
		var tmp string
		var flag bool

		for _, stock := range feed.Items {

			stock.Published = stock.Published + " +0530 IST"

			// Parse input timestamp
			parsedUTC, err := time.Parse(layout, stock.Published)
			if err != nil {
				panic(fmt.Sprintf("ERROR: parsing feed: %v ---------stock Time: %v ------- parsed Time: %v\n", err, stock.Published, parsedUTC))
			}

			// fmt.Println("----------------------------------------------------", stockPublishedTime, stock.Title)

			if parsedUTC.After(initialTime) {

				//from the nse feeds loads from latest to old, so has to fetch the one value which is the latest
				if !flag {
					tmp = stock.Published
					flag = true
				}

				// fmt.Println("----------------------------------------------------After---", len(feed.Items), stock.Published, stock.Title)

				if strings.Contains(stock.Description, "orders/contracts") || strings.Contains(stock.Description, "Orders/Contracts") || strings.Contains(stock.Description, "orders/contracts") || strings.Contains(stock.Description, "buyback") {

					if NotifiedFeeds[stock.Title] != "" {
						fmt.Println("feed already notified so ignoring it ------------", stock.Title, stock.Published, stock.Description, stock.Link)
						continue
					}

					NotifiedFeeds[stock.Title] = "Done"

					fmt.Println("Contracts feed  ------------", stock.Title, stock.Published, stock.Description, stock.Link)

					message := fmt.Sprintf("STOCK: %v ----- %v\n CATEGORY: %v\n LINK: %v", stock.Title, stock.Published, stock.Description, stock.Link)

					err := SendTelegramMessage(message)
					if err != nil {
						fmt.Println("Error:", err)
					}

				}

			}

		}

		if tmp == "" { //no latest feeds found on this function call, so again setting this initialTime again for next iteration function call
			return initialTime
		}

		lastProcessed, err = time.Parse(layout, tmp)
		if err != nil {
			panic(fmt.Sprintf("%v ------------- %v", err, tmp))
		}

		return lastProcessed
	}()
	// 	// choose key for dedup: prefer GUID, then Link, then Title+PubDate
	// 	key := item.GUID
	// 	if key == "" {
	// 		key = item.Link
	// 	}
	// 	if key == "" {
	// 		key = item.Title + "|" + item.Published
	// 	}
	// 	if wasSeen(key) {
	// 		continue
	// 	}

	// 	// mark seen immediately (to avoid duplicates on re-run). In production, use DB transaction
	// 	markSeen(key)

	// 	wg.Add(1)
	// 	// concurrency-limited goroutine
	// 	go func(it *gofeed.Item, dedupKey string) {
	// 		defer wg.Done()
	// 		sem <- struct{}{}        // acquire
	// 		defer func() { <-sem }() // release
	// 		if err := processItem(it); err != nil {
	// 			log.Printf("ERROR processing item %s : %v\n", dedupKey, err)
	// 		}
	// 	}(item, key)
	// }
	// wg.Wait()

	return val
}

func testReadAndGetSentiment() {

	s, e := fetchAndExtractContent("https://nsearchives.nseindia.com/corporate/TATACHEMYS_06122025104921_SE_Intimation-CRISIL_05122025_v1_signed.pdf")
	// s, e := fetchAndExtractContent("https://nsearchives.nseindia.com/corporate/NEULANDLAB_06122025110638_SignedSE_Intimation.pdf")

	s1 := strings.Split(s, "\n")
	data := strings.Join(s1, "")
	for _, i := range nseNewsSamples {
		sentiment, score := analyzeSentiment(i)
		fmt.Println(i, "---------", sentiment, score)
	}
	sentiment, score := analyzeSentiment(data)
	fmt.Println(e, "-------------", data, "---------", sentiment, score)

}

func wasSeen(key string) bool {
	seenMu.Lock()
	defer seenMu.Unlock()
	_, ok := seen[key]
	return ok
}

func markSeen(key string) {
	seenMu.Lock()
	defer seenMu.Unlock()
	seen[key] = time.Now()
}

// processItem fetches PDF (if any), extracts text, runs sentiment; logs result
func processItem(item *gofeed.Item) error {
	// Basic extraction
	title := strings.TrimSpace(item.Title)
	desc := strings.TrimSpace(item.Description)
	pub := item.Published
	link := strings.TrimSpace(item.Link)

	log.Printf("Processing item: title=%q published=%q link=%s\n", title, pub, link)

	// if there's no link, skip or try enclosure
	if link == "" && len(item.Enclosures) > 0 {
		link = item.Enclosures[0].URL
	}

	var contentText string
	var err error

	if link != "" {
		contentText, err = fetchAndExtractContent(link)
		if err != nil {
			// If PDF extraction fails, fall back to use description + title
			log.Printf("WARN: failed fetch/extract from link %s: %v. Falling back to title+desc.\n", link, err)
			contentText = title + "\n\n" + desc + "\n\n" + contentText
		}
	} else {
		// No link: use title and description
		contentText = title + "\n\n" + desc
	}

	// Compose final text to analyze
	toAnalyze := title + "\n\n" + desc + "\n\n" + contentText
	sentiment, score := analyzeSentiment(toAnalyze)

	// You can replace this with DB insert or call to other services
	log.Printf("RESULT: title=%q published=%q sentiment=%s score=%.3f link=%s\n",
		title, pub, sentiment, score, link)

	return nil
}

// fetchAndExtractContent downloads the link and attempts to extract text from PDF
// If the link is HTML, it will try to find any PDF link inside the HTML and fetch that.
// Returns extracted text or error.
func fetchAndExtractContent(rawURL string) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	// try direct download with retries
	var resp *http.Response
	for attempt := 1; attempt <= maxRetries; attempt++ {
		req, _ := http.NewRequest("GET", u.String(), nil)
		// req.Header.Set("User-Agent", userAgent)
		// req.Header.Set("Accept", "*/*")

		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
		req.Header.Set("Accept", "*/*")
		req.Header.Set("Referer", "https://www.nseindia.com")
		req.Header.Set("Accept-Language", "en-US,en;q=0.9")
		req.Header.Set("Connection", "keep-alive")
		resp, err = client.Do(req)
		if err == nil && resp != nil && resp.StatusCode >= 200 && resp.StatusCode < 300 {
			break
		}
		// close body if non-nil
		if resp != nil && resp.Body != nil {
			io.Copy(io.Discard, resp.Body)
			resp.Body.Close()
		}
		// backoff
		time.Sleep(retryBaseDelay * time.Duration(attempt))
	}
	if err != nil {
		return "", fmt.Errorf("failed to fetch URL: %w", err)
	}
	defer resp.Body.Close()

	ct := resp.Header.Get("Content-Type")
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// if content-type indicates PDF, extract
	if strings.Contains(strings.ToLower(ct), "pdf") || strings.HasSuffix(strings.ToLower(u.Path), ".pdf") {
		return extractTextFromPDFBytes(bodyBytes)
	}

	// If the response is HTML, try to find a PDF link inside the HTML
	if strings.Contains(strings.ToLower(ct), "html") || strings.HasSuffix(strings.ToLower(u.Path), ".htm") || strings.HasSuffix(strings.ToLower(u.Path), "/") {
		pdfURL := findPDFLinkInHTML(bodyBytes, u)
		if pdfURL != "" {
			// fetch pdf and extract
			return fetchAndExtractContent(pdfURL)
		}
		// fallback: return raw HTML stripped text (we keep HTML as last resort)
		text := stripHTMLToText(bodyBytes)
		return text, nil
	}

	// last resort: treat bytes as text
	if isMostlyText(bodyBytes) {
		return string(bodyBytes), nil
	}

	return "", errors.New("content is neither PDF nor HTML nor readable text")
}

// extractTextFromPDFBytes extracts text from PDF bytes using ledongthuc/pdf
func extractTextFromPDFBytes(b []byte) (string, error) {
	// pdf library expects an io.Reader; use bytes.Reader
	reader := bytes.NewReader(b)
	// pdf.Open requires a file-like reader that implements ReadAt; bytes.Reader has that
	r, err := pdf.NewReader(reader, int64(len(b)))
	if err != nil {
		return "", fmt.Errorf("pdf.NewReader error: %w", err)
	}

	var buf strings.Builder
	numPages := r.NumPage()
	for i := 1; i <= numPages; i++ {
		p := r.Page(i)
		if p.V.IsNull() {
			continue
		}
		txt, err := p.GetPlainText(nil)
		if err != nil {
			// continue on page-level errors
			log.Printf("WARN: error extracting page %d: %v\n", i, err)
			continue
		}
		buf.WriteString(txt)
		buf.WriteString("\n\n")
	}
	out := strings.TrimSpace(buf.String())
	if out == "" {
		return "", errors.New("no text extracted from pdf")
	}
	return out, nil
}

// findPDFLinkInHTML scans HTML bytes and returns absolute PDF URL if found
// This is a simple approach: search for ".pdf" and extract surrounding href.
func findPDFLinkInHTML(htmlBytes []byte, base *url.URL) string {
	s := string(htmlBytes)
	lower := strings.ToLower(s)
	idx := strings.Index(lower, ".pdf")
	if idx == -1 {
		return ""
	}
	// walk backwards to find href="
	start := strings.LastIndex(lower[:idx], "href=\"")
	if start == -1 {
		start = strings.LastIndex(lower[:idx], "href='")
		if start == -1 {
			// try naive extraction around idx
			start = idx - 200
			if start < 0 {
				start = 0
			}
		}
	}
	// find the closing quote after idx
	end := strings.IndexAny(lower[idx:], "\"'")
	if end == -1 {
		end = 50
	}
	// try to extract substring
	fragment := s[start : idx+end]
	// find "http" inside fragment
	httpIdx := strings.Index(fragment, "http")
	if httpIdx != -1 {
		raw := fragment[httpIdx:]
		// crop trailing junk
		raw = strings.SplitN(raw, "\"", 2)[0]
		raw = strings.SplitN(raw, "'", 2)[0]
		parsed, err := url.Parse(strings.TrimSpace(raw))
		if err == nil {
			if parsed.IsAbs() {
				return parsed.String()
			}
			abs := base.ResolveReference(parsed)
			return abs.String()
		}
	}
	// Fallback: find first href= before idx and then between quotes
	hrefStart := strings.LastIndex(lower[:idx], "href=")
	if hrefStart == -1 {
		return ""
	}
	// extract quote char
	q := lower[hrefStart+5]
	rest := s[hrefStart+6:]
	parts := strings.SplitN(rest, string(q), 3)
	if len(parts) >= 2 {
		raw := parts[1]
		parsed, err := url.Parse(strings.TrimSpace(raw))
		if err == nil {
			if parsed.IsAbs() {
				return parsed.String()
			}
			abs := base.ResolveReference(parsed)
			return abs.String()
		}
	}
	return ""
}

// stripHTMLToText: minimal HTML to text removal for fallback
func stripHTMLToText(b []byte) string {
	s := string(b)
	// naive removal of tags: remove <...>
	inTag := false
	var out []rune
	for _, r := range s {
		if r == '<' {
			inTag = true
			continue
		}
		if r == '>' {
			inTag = false
			continue
		}
		if !inTag {
			out = append(out, r)
		}
	}
	// collapse whitespace
	txt := strings.Join(strings.Fields(string(out)), " ")
	return txt
}

// isMostlyText heuristics
func isMostlyText(b []byte) bool {
	if len(b) == 0 {
		return false
	}
	sample := b
	if len(b) > 1024 {
		sample = b[:1024]
	}
	textChars := 0
	for _, c := range sample {
		if (c >= 32 && c <= 126) || c == '\n' || c == '\r' || c == '\t' {
			textChars++
		}
	}
	return float64(textChars)/float64(len(sample)) > 0.8
}

// analyzeSentiment returns label and compound score
func analyzeSentiment(text string) (label string, score float64) {
	// Preprocess: trim and limit length (VADER works better with shorter to medium length)
	trimmed := strings.TrimSpace(text)
	if len(trimmed) == 0 {
		return "neutral", 0.0
	}
	if len(trimmed) > 20000 {
		trimmed = trimmed[:20000] // cap large docs to speed up
	}

	analyzer := govader.NewSentimentIntensityAnalyzer()
	s := analyzer.PolarityScores(trimmed)
	compound := s.Compound

	// thresholds are typical for VADER-style
	if compound >= 0.05 {
		return "positive", compound
	} else if compound <= -0.05 {
		return "negative", compound
	}
	return "neutral", compound
}

var nseNewsSamples = []string{
	// ---------- Positive (10) ----------
	"Reliance Industries surges after reporting a 22% jump in quarterly net profit and record revenue from Jio.",
	"Tata Motors rallies 5% as the company posts strong EV sales growth and improved export numbers.",
	"HDFC Bank announces a special dividend and expects loan book growth to remain strong in FY26.",
	"Infosys signs a $2.1 billion multi-year digital transformation deal with a global bank.",
	"Adani Green secures long-term renewable energy project approvals from the government.",
	"Mahindra & Mahindra stock jumps after tractor sales hit an all-time high in November.",
	"Divi’s Labs climbs as US FDA gives positive clearance with zero observations.",
	"JSW Steel rises as global steel prices rebound and demand outlook improves.",
	"Indian Hotels gains after strong tourism and record festive season occupancy.",
	"HCL Tech soars after raising revenue guidance and announcing share buyback worth ₹12,000 crore.",

	// ---------- Negative (10) ----------
	"Zee Entertainment crashes 18% after merger talks with Sony officially collapse.",
	"Paytm shares plunge as RBI imposes restrictions on Paytm Payments Bank operations.",
	"Yes Bank drops after large block deal sparks concerns of further stake dilution.",
	"Wipro falls as management cuts revenue forecast for the next quarter.",
	"Vedanta declines sharply amid renewed debt restructuring worries.",
	"Biocon tumbles after US FDA issues multiple critical observations during facility inspection.",
	"Bandhan Bank stock hits lower circuit after reporting significant spike in NPAs.",
	"IOC slips as crude oil prices surge, raising margin concerns.",
	"Dixon Technologies sinks after fire incident halts production at its largest unit.",
	"Bharat Forge falls on fears of reduced orders due to geopolitical tensions affecting exports.",

	// ---------- Neutral (10) ----------
	"NSE announces quarterly index reshuffle; several stocks added and removed from NIFTY indices.",
	"SEBI releases new guidelines for margin reporting for equity derivatives from January.",
	"RBI maintains status quo on repo rate; markets react flat ahead of policy commentary.",
	"TCS states that hiring will remain stable and campus intake to be similar to previous year.",
	"Maruti Suzuki schedules its annual investor meeting next week.",
	"Coal India reports coal dispatch and production numbers for the month without major deviation.",
	"UltraTech Cement completes phase-2 expansion as per previously published roadmap.",
	"IRCTC clarifies that no ticket price revision proposal is currently under consideration.",
	"PowerGrid announces board meeting to consider fundraising through bonds.",
	"Larsen & Toubro updates its project pipeline and order book status with no major surprises.",
}
