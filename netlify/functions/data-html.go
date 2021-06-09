package main

import (
  	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
	"log"
	"os"
	"io/ioutil"
        "github.com/tidwall/gjson"
)
func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	// read all response body
      	//Get the path parameter that was sent
	//uuid := request.PathParameters["uuid"]
        uuid := request.QueryStringParameters["uuid"]

	data := Request( "https://hgrqtmovrbvpvxjqungh.supabase.co/rest/v1/item?select=id,url,item_adult(value,score),item_brand(value,score),item_category(value,score),item_celebrity(value,score),item_color(black_and_white,accent_color,dominant_color_background,dominant_color_foreground,dominant_colors),item_description(value,score),item_face(gender,age),item_landmark(value,score),item_object(value,score),item_racy(value,score),item_tag(value,score),item_text(line,value),item_text_entity(value,match_text,text_type,text_sub_type,text_score),item_text_key_phrase(value),item_text_language(value,code,score),item_text_sentiment(score)&and=(id.eq."+uuid+")" )
	// print `data` as a string
        //blah := string(data)
       // value := gjson.Get(string(data), "item_description.0.value")
        var blah string
        blah = "<table>"
        blah = blah + "<tr>"
        blah = blah + "<td>Item description:</td><td>"+gjson.Get(string(data), "item_description.0.value").String()+"</td>"
        blah = blah + "</tr>"
        blah = blah + "<tr>"
        blah = blah + "<td>Item Score:</td><td>"+gjson.Get(string(data), "item_description.0.score").String()+"</td>"
        blah = blah + "</tr>"
        blah = blah + "</table>"
        // Iterating address objects
        /*
	for key, child := range jsonParsed.Search("employees", "address").ChildrenMap() {
		fmt.Printf("Key=>%v, Value=>%v\n", key, child.Data().(string))
	}
        */

	return &events.APIGatewayProxyResponse{
		StatusCode:        200,
		Headers:           map[string]string{"Content-Type": "text/plain"},
		MultiValueHeaders: http.Header{"Set-Cookie": {"Ding", "Ping"}},
                Body:              string(blah),
		IsBase64Encoded:   false,
	}, nil
}

func main() {
	
    lambda.Start(handler)
	// make a sample HTTP GET request
	

}


func Request(url string) []byte {
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        log.Fatal(err)
    }
    req.Header.Set("apikey",  os.Getenv("SUPABASE_API_KEY"))
    req.Header.Set("Accept",  "application/vnd.pgrst.object+json")
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

