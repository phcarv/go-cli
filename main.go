package main
import (
	"fmt"
	"net/url" 
	"os"
	"github.com/joho/godotenv"
)

type Content struct {
	Contents []struct{
		Parts []struct{
			Text string `json:"text"`
		} `json:"parts"`
	} `json:"contents"`
}


func main(){
	err := godotenv.Load()
	if err != nil {
		panic("Wasn't able to load .env")
	}

	jsonRequestData := []byte(`{
		"contents": [
		  {
			"parts": [
			  {
				"text": "give me a simple motivational quote.\nThe the quote must be in the following format: \"the quote itself\" - the author of the quote.\nInside of the double quotes you will put the quote and after the - you will put the name of the author of the quote"
			  }
			]
		  }
		]
	  }`)

	var content Content
	json.Unmarshal(jsonRequestData, &content)

	//resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)

	//fmt.Printf("The quote for today is: \n\n%s", quote)
	var key = os.Getenv("GEMINI_API_KEY")
	fmt.Printf("%s", key)

}