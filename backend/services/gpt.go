package services

import (
	"context"
	"os"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/openai/openai-go/responses"
)

func CleanJobDescription(jobdescription string) (cleanedJD string, err error) {
	client := openai.NewClient( //a client that will communicate with OpenAI API
		//option.WithAPIKey tells client which API key to use
		option.WithAPIKey(os.Getenv("OPENAI_API_KEY")),
	)
	prompt := CleanJDPrompt + jobdescription
	resp, err := client.Responses.New( //Create a new response from the model.
		context.Background(), //req information,gpt takes 30s but your browser closes after 5s
		responses.ResponseNewParams{
			Model: openai.ChatModelGPT4_1Mini,
			Input: responses.ResponseNewParamsInputUnion{ //union-> one of the several possible types of input(text,image,file)
				OfString: openai.String(prompt), //OfString -> means input is just a string
			},
		},
	)
	cleanedJD = resp.OutputText() //helper method provided by the SDK
	return
}
