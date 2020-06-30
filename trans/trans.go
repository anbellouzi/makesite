package trans

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/translate"
	"github.com/joho/godotenv"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
)

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func TranslateText(targetLanguage, text string) (string, error) {
	// text := "The Go Gopher is cute"
	ctx := context.Background()

	apiKey, _ := os.LookupEnv("TRANSLATE_KEY")

	lang, err := language.Parse(targetLanguage)
	if err != nil {
		return "", fmt.Errorf("language.Parse: %v", err)
	}

	client, err := translate.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return "", err
	}
	defer client.Close()

	resp, err := client.Translate(ctx, []string{text}, lang, nil)
	if err != nil {
		return "", fmt.Errorf("Translate: %v", err)
	}
	if len(resp) == 0 {
		return "", fmt.Errorf("Translate returned empty response to text: %s", text)
	}
	return resp[0].Text, nil
}
