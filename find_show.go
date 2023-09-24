package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

// ShowResponse wraps a TV Maze search response
type ShowResponse struct {
	Score float64
	Show  Show
}

// Show wraps a TV Maze show object
type Show struct {
	ID      int
	Name    string
	Type    string
	Genres  []string
	Status  string
	Runtime int
	Summary string
	Remotes map[string]*json.RawMessage `json:"externals"`
	Image   struct {
		Medium   string
		Original string
	}
}

func main() {
	var n int
	var s []int
	baseURL := "https://api.tvmaze.com/"
	resource := "search/shows"
	params := url.Values{}
	params.Add("q", os.Args[1])

	u, _ := url.ParseRequestURI(baseURL)
	u.Path = resource
	u.RawQuery = params.Encode()
	urlStr := fmt.Sprintf("%v", u) // "https://api.tvmaze.com/search/shows?q="

	// fmt.Println(urlStr)

	response, err := http.Get(urlStr)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(string(responseData))
	// fmt.Println(len(responseData))

	var result []ShowResponse
	if err := json.Unmarshal(responseData, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	//fmt.Println(PrettyPrint(result))
    fmt.Println(len(result))
	// Loop through the show node for the Title and ID
	for i, rec := range result {
		fmt.Println(i, ": ", rec.Show.Name, "(", rec.Show.ID, ") - Status:", rec.Show.Status)
		s = append(s, rec.Show.ID)
	}
	
    if len(result) > 1 {
        fmt.Println("\n More than one result, Please Choose:")
		fmt.Scan(&n)
		fmt.Println(s[n])

}
}
// PrettyPrint to print struct in a readable way
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}



