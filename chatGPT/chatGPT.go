package chatGPT

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	resty "github.com/go-resty/resty/v2"
	"io/ioutil"
	"log"
	"os"
)

// const apiKey = <your key here>

const (
	apiEndpoint = "https://api.openai.com/v1/chat/completions"
)

const streetMapsQuery1 = `
Does this link show a bridge %s
Please answer just yes or no.
`

const streetMapsQuery2 = `
What is the height under the bridge: %s
Please answer just yes or no plus the height of the bridge
`

func QueryDir(outputDir string) error {

	outputDirFile, err := os.ReadDir(outputDir)
	if err != nil {
		return err
	}

	for _, imagesDir := range outputDirFile {
		resultsForVideo := make([]string, 0)
		if !imagesDir.IsDir() {
			continue
		}
		imageDirName := fmt.Sprintf("%s/%s", outputDir, imagesDir.Name())
		imageDirFiles, err := os.ReadDir(imageDirName)
		if err != nil {
			return err
		}

		for _, imageFile := range imageDirFiles {
			if imageFile.IsDir() {
				continue
			}
			fileName := fmt.Sprintf("%s/%s", imageDirName, imageFile.Name())

			result := Query_image_2(fileName)
			resultsForVideo = append(resultsForVideo, result)
		}

		fmt.Printf("=================\n\n")
		for _, result := range resultsForVideo {
			fmt.Println(result)
		}

	}

	return nil

}

func Query_image_2(filePath string) string {
	// Use your API KEY here

	//filePath := "/Users/marcin.mirecki/go/src/github.com/mmirecki/examples/chatgpt/LBS_2.png"

	// Read the file
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	encodedFile := base64.StdEncoding.EncodeToString(fileData)
	if len(encodedFile) == -1 {
		fmt.Println()
	}

	client := resty.New()

	response, err := client.R().
		SetAuthToken(apiKey).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"model": "gpt-4o-mini",
			"messages": []interface{}{map[string]interface{}{
				"role": "user",
				"content": []map[string]interface{}{
					{
						"type": "text",
						"text": "Does the image show a bridge? What is the clearance under the bridge? Answer just YES or NO plus the height",
						//
						//
						//"text": "Does the image show a bridge? What is the clearance under the bridge?",
						//"text": "Does the image show a bridge? What is the clearance under the bridge? Answer just YES or NO plus the clearance height",
					},
					{
						"type": "image_url",
						"image_url": map[string]string{
							"url": fmt.Sprintf("data:image/jpeg;base64,{%s}", encodedFile),
						},
					},
				},
			}},
			"max_tokens": 50,
		}).
		Post(apiEndpoint)

	if err != nil {
		log.Fatalf("Error while sending send the request: %v", err)
	}

	body := response.Body()

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error while decoding JSON response:", err)
		return ""
	}

	//fmt.Printf("DATA: %+v\n", data)
	// Extract the content from the JSON response
	content := data["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
	fmt.Println(content)

	return content

}
