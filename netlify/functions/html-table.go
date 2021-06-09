package main

import (
	"net/http"
	"fmt"
	"log"
	"os"
	"io/ioutil"
)

func main() {
	
	// make a sample HTTP GET request
	

	// read all response body
	data := Request( "https://hgrqtmovrbvpvxjqungh.supabase.co/rest/v1/item?select=id,url,item_adult(value,score),item_brand(value,score),item_category(value,score),item_celebrity(value,score),item_color(black_and_white,accent_color,dominant_color_background,dominant_color_foreground,dominant_colors),item_description(value,score),item_face(gender,age),item_landmark(value,score),item_object(value,score),item_racy(value,score),item_tag(value,score),item_text(line,value),item_text_entity(value,match_text,text_type,text_sub_type,text_score),item_text_key_phrase(value),item_text_language(value,code,score),item_text_sentiment(score)&and=(id.eq.beec7a1a-bdc1-41dc-b7e1-5c15565c3efc)" )
	// print `data` as a string
	fmt.Printf( "%s\n", data )
}


func Request(url string) []byte {
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        log.Fatal(err)
    }
    req.Header.Set("apikey",  os.Getenv("SUPABASE_API_KEY"))
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
    return body
}

