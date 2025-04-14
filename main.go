package main
import (
	"fmt"
	//"net/url"
	"os"
)

func main(){
	//var quote = "The sun himself is weak when he first rises, and gathers strength and courage as the day gets on. - Charles Dickens."

	//resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)

	//fmt.Printf("The quote for today is: \n\n%s", quote)
	var key = os.Getenv("GEMINI_API_KEY")
	fmt.Printf("%s", key)

}