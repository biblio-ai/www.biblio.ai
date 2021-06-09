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

	data := Request( "https://hgrqtmovrbvpvxjqungh.supabase.co/rest/v1/item?select=id,url,item_adult(value,score),item_brand(value,score,x,y,width,height),item_category(value,score),item_celebrity(value,score,position_height,position_left,position_top,position_width),item_color(black_and_white,accent_color,dominant_color_background,dominant_color_foreground,dominant_colors),item_description(value,score),item_face(gender,age,position_height,position_left,position_top,position_width),item_landmark(value,score),item_object(value,score,x,y,width,height),item_racy(value,score),item_tag(value,score),item_text(line,value,box),item_text_entity(value,match_text,text_type,text_sub_type,text_score),item_text_key_phrase(value),item_text_language(value,code,score),item_text_sentiment(score)&and=(id.eq."+uuid+")" )
	// print `data` as a string
        //blah := string(data)
       // value := gjson.Get(string(data), "item_description.0.value")
        var blah string
        blah = "<table>"
        blah = blah + "<tr>" 
        blah = blah + "<td>Item ID:</td><td>"+gjson.Get(string(data), "id").String()+"</td>"
        blah = blah + "</tr>"
        blah = blah + "<tr>"
        blah = blah + "<td>Item URL:</td><td>"+gjson.Get(string(data), "url").String()+"</td>"
        blah = blah + "</tr>"
        blah = blah + "<tr>"
        blah = blah + "<td><b>Item description:</b></td><td></td>"
        blah = blah + "</tr>"
        blah = blah + "<tr>"
        blah = blah + "<td>Item description:</td><td>"+gjson.Get(string(data), "item_description.0.value").String()+"</td>"
        blah = blah + "</tr>"
        blah = blah + "<tr>"
        blah = blah + "<td>Item Score:</td><td>"+gjson.Get(string(data), "item_description.0.score").String()+"</td>"
        blah = blah + "</tr>"
        blah = blah + "<tr>"
        blah = blah + "<td><b>Item tag:</b></td><td></td>"
        blah = blah + "</tr>"
        result := gjson.Get(string(data), "item_tag.#.value")
for _, name := range result.Array() {
                blah = blah + "<tr>"
                blah = blah + "<td>Item tag:</td><td>"+name.String()+"</td>"
                blah = blah + "</tr>"
}
        blah = blah + "<tr>"
        blah = blah + "<td><b>Item object:</b></td><td></td>"
        blah = blah + "</tr>"
        result_obj := gjson.Get(string(data), "item_object.#.value")
for _, name := range result_obj.Array() {
                blah = blah + "<tr>"
                blah = blah + "<td>Item object:</td><td>"+name.String()+"</td>"
                blah = blah + "</tr>"
}
        blah = blah + "<tr>"
        blah = blah + "<td><b>Item text:</b></td><td></td>"
        blah = blah + "</tr>"
        result_text := gjson.Get(string(data), "item_text.#.value")
for _, name := range result_text.Array() {
                blah = blah + "<tr>"
                blah = blah + "<td>Item description:</td><td>"+name.String()+"</td>"
                blah = blah + "</tr>"
}
        blah = blah + "<tr>"
        blah = blah + "<td><b>Item text key phrases:</b></td><td></td>"
        blah = blah + "</tr>"
        result_text_key := gjson.Get(string(data), "item_text_key_phrase.#.value")
for _, name := range result_text_key.Array() {
                blah = blah + "<tr>"
                blah = blah + "<td>Key phrase:</td><td>"+name.String()+"</td>"
                blah = blah + "</tr>"
}
        blah = blah + "</table>"
        // Iterating address objects
        /*
	for key, child := range jsonParsed.Search("employees", "address").ChildrenMap() {
		fmt.Printf("Key=>%v, Value=>%v\n", key, child.Data().(string))
	}
        */

	return &events.APIGatewayProxyResponse{
		StatusCode:        200,
		Headers:           map[string]string{"Content-Type": "text/html"},
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

