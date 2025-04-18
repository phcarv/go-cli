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
				"text": "Consider that you are a wise writter that knows all the verses from the bible, I need you to give me a quote/verse from the bible in the following format:\n \" The bible verse \" - Reference to the verse. \n So inside of the double quotes you will insert a random verse and after the dash(-) you will put the reference like: John 3:16"
			  },
			  {
				"text": "The quote should be for helping me on a sad day"
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