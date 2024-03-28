package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"

	"github.com/go-chi/chi"
)

const ListeningPort = ":10000"

var InputRegionTrack = "India"

var RegionTopTrack = map[string]Track{}
var AllTracks = []Track{}

type Track struct {
	Id         int               `json:"id"`
	Name       string            `json:"name"`
	Lyrics     string            `json:"lyrics"`
	ArtistInfo ArtistInformation `json:"artistinfo"`
	Tag        string            `json:"tag"`
}

type ArtistInformation struct {
	Name      string `json:"name"`
	ImageLink string `json:"imagelink"`
}

func init() {
	uploadTrackSampleValues()
}

func uploadTrackSampleValues() {
	track1 := Track{1, "Without You", "Without You track's lyrics", ArtistInformation{Name: "Harry Nilsson", ImageLink: "https://lastfm.freetls.fastly.net/i/u/34s/2a96cbd8b46e442fc41c2b86b821562f.png"}, "pop"}
	track2 := Track{2, "You’re So Vain", "You’re So Vain track's lyrics", ArtistInformation{Name: "Carly Simon", ImageLink: "https://lastfm.freetls.fastly.net/i/u/34s/2a96cbd8b46e442fc41c2b86b821562f.png"}, "pop"}
	track3 := Track{3, "Time After Time", "Time After Time track's lyrics", ArtistInformation{Name: "Cyndi Lauper", ImageLink: "https://lastfm.freetls.fastly.net/i/u/34s/2a96cbd8b46e442fc41c2b86b821562f.png"}, "pop"}
	track4 := Track{4, "Where Is My Mind?", "Where Is My Mind? track's lyrics", ArtistInformation{Name: "The Pixies", ImageLink: "https://lastfm.freetls.fastly.net/i/u/34s/2a96cbd8b46e442fc41c2b86b821562f.png"}, "melody"}
	track5 := Track{5, "So What", "So What track's lyrics", ArtistInformation{Name: "Cyndi Lauper", ImageLink: "https://lastfm.freetls.fastly.net/i/u/34s/2a96cbd8b46e442fc41c2b86b821562f.png"}, "melody"}
	track6 := Track{6, "Welcome to the Jungle", "Welcome to the Jungle track's lyrics", ArtistInformation{Name: "Guns N Roses", ImageLink: "https://lastfm.freetls.fastly.net/i/u/34s/2a96cbd8b46e442fc41c2b86b821562f.png"}, "beat"}
	track7 := Track{7, "Old Town Road", "Old Town Road track's lyrics", ArtistInformation{Name: "Lil Nas X", ImageLink: "https://lastfm.freetls.fastly.net/i/u/34s/2a96cbd8b46e442fc41c2b86b821562f.png"}, "beat"}
	track8 := Track{8, "Cannonball", "Cannonball track's lyrics", ArtistInformation{Name: "The Breeders", ImageLink: "https://lastfm.freetls.fastly.net/i/u/34s/2a96cbd8b46e442fc41c2b86b821562f.png"}, "Harry Nilsson"}
	track9 := Track{9, "House of Balloons", "House of Balloons track's lyrics", ArtistInformation{Name: "Miless Davis", ImageLink: "https://lastfm.freetls.fastly.net/i/u/34s/2a96cbd8b46e442fc41c2b86b821562f.png"}, "pop"}

	AllTracks = append(AllTracks, track1, track2, track3, track4, track5, track6, track7, track8, track9)

	RegionTopTrack["india"] = track3
	RegionTopTrack["italy"] = track5
	RegionTopTrack["usa"] = track4
	RegionTopTrack["russia"] = track3

}

func main() {
	router := chi.NewRouter()
	router.Get("/region/{name}", GetRegionTopTrack)

	// router.Use()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go http.ListenAndServe(ListeningPort, router)

	//http client request for accessing the above API
	response, err := http.Get("http://localhost:10000/region/" + InputRegionTrack)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	byteResponse, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("API Response status -- ", response.Status)
	fmt.Println("API Response for the request -- ", string(byteResponse))

	wg.Wait()

	// c := chi.NewRouteContext()
	// f := chi.ChainHandler{}

	// d := http.Transport{}

}

// http request
// http response

func GetRegionTopTrack(response http.ResponseWriter, r *http.Request) {

	fmt.Println("Printing basic request information---------------")
	fmt.Println(r.UserAgent())
	fmt.Println(r.Host)
	fmt.Println(r.URL, r.URL.Path)
	fmt.Println(r.Cookie("exampleCookie"))
	fmt.Println(r.Cookies())
	fmt.Println(r.ContentLength)

	//validates whether we gets expected http request or not
	if r.Method != http.MethodGet {
		http.Error(response, "UnAllowed method", http.StatusMethodNotAllowed) //405 http code for unallowed method.
		return
	}

	response.Header().Set("Content-Type", "application/json")
	regionName := chi.URLParam(r, "name")
	regionName = strings.ToLower(regionName)
	_, found := RegionTopTrack[regionName]

	if !found {
		http.Error(response, "Top track is not found for this Region/Country name - "+string(regionName), http.StatusNotFound)
		return
	}

	result := map[string]interface{}{}
	topTrack := RegionTopTrack[regionName]

	//Adding the region's top track in the response
	result["Top Track of this Region - "+string(regionName)] = topTrack

	//Based on the region's top track, Also adding the suggestion tracks based on the track's tag and artist in the response
	suggestedTracks := []Track{}
	for _, track := range AllTracks {
		//Ensures top track is again not duplicates to the suggested tracks
		if track.Id == topTrack.Id {
			continue
		}
		if track.Tag == topTrack.Tag || track.ArtistInfo.Name == topTrack.ArtistInfo.Name {
			suggestedTracks = append(suggestedTracks, track)
		}
	}

	//setting example cookie
	// cookie := http.Cookie{
	// 	Name:     "exampleCookie",
	// 	Value:    "Hello world!",
	// 	Path:     "/",
	// 	MaxAge:   3600,
	// 	HttpOnly: true,
	// 	Secure:   true,
	// 	SameSite: http.SameSiteLaxMode,
	// }

	// http.SetCookie(response, &cookie)

	result["Suggested tracks"] = suggestedTracks
	json.NewEncoder(response).Encode(result)
}
