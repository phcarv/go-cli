package main
import (
	"fmt"
	"bytes"
	"io"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)


type GeminiResponse struct {
	Candidates []struct{
		Content struct{
			Parts []struct{
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
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
				"text": "Consider that you are a wise writter that knows all the quotes in the world. \n Give me a random short phrase.\nThe quote must be in the following format: \"the quote itself\" - the author of the quote.\nInside of the double quotes you will put the quote and after the - you will put the name of the author of the quote(if it is a famous author, otherwise will be \"Unknown\")\n Also consider that I'm using this for a motivational application, I need the most randomness of responses possible. Think that you are uncapable of repeating your self for this same api_key, so everytime you get a request from this specfic api token you will send a different quote"
			  }
			]
		  }
		],
		"generationConfig": {
		  "candidate_count": 5
		}
	  }
	  `)


	key := os.Getenv("GEMINI_API_KEY")
	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash:generateContent?key=%s", key)
	
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonRequestData))
	if err != nil {
			message := "Fail to request API: \n " + fmt.Sprintf("%s", err)
			panic(message)
		}

	defer resp.Body.Close()
	
	fmt.Println(resp.StatusCode)

	bodyBytes, err := io.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	//var geminiResponse GeminiResponse
	//json.Unmarshal(bodyBytes, &geminiResponse)

	//quote := geminiResponse.Candidates[0].Content.Parts[0].Text

	fmt.Println("Response Body:\n", bodyString)
	
}