package main
import (
	"fmt"
	"bytes"
	"io"
	"net/http" 
	"os"
	"github.com/joho/godotenv"
)

// type Content struct {
// 	Contents []struct{
// 		Parts []struct{
// 			Text string `json:"text"`
// 		} `json:"parts"`
// 	} `json:"contents"`
// }


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
				"text": "give me a random motivational quote.\nThe the quote must be in the following format: \"the quote itself\" - the author of the quote.\nInside of the double quotes you will put the quote and after the - you will put the name of the author of the quote - VERY IMPORTANT: Don't repeat previous quotes"
			  }
			]
		  }
		]
	  }`)


	key := os.Getenv("GEMINI_API_KEY")
	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash:generateContent?key=%s", key)
	
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonRequestData))
	if err != nil {
			message := "Fail to request API: \n " + fmt.Sprintf("%s", err)
			panic(message)
		}

	defer resp.Body.Close()
	
	fmt.Println(resp.StatusCode)

	// Read the entire response body
	bodyBytes, err := io.ReadAll(resp.Body)

	// Convert the byte slice to a string and print it
	bodyString := string(bodyBytes)
	fmt.Println("Response Body:\n", bodyString)
	//fmt.Printf(resp)
	
}