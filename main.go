package main

import (
	"fmt"
	"github.com/mmirecki/bridgeSignFinder/chatGPT"
	"github.com/mmirecki/bridgeSignFinder/video2image"
)

func main() {

	// TODO: your own path here
	videoPath := "/Users/marcin.mirecki/go/src/github.com/mmirecki/bridgeSignFinder/samples"
	outputPath := fmt.Sprintf("%s/%s", videoPath, "output")

	err := video2image.CheckSamples(videoPath, outputPath)
	if err != nil {
		fmt.Printf("Error processing video samples: %v", err)
		return
	}

	//sample := "/Users/marcin.mirecki/go/src/github.com/mmirecki/bridgeSignFinder/samples/844424930861269-1737384860003_00018.png"

	//chatGPT.Query_image_2(outputPath)
	results, err := chatGPT.QueryDir(outputPath)
	if err != nil {
		fmt.Printf("Error querying images: %v", err)
		return
	}

	fmt.Printf("\n\n ------------------------------------ \n ------------------------------------ \n\n")
	fmt.Printf("\n\n ------------------------------------ \n ------------------------------------ \n\n")
	
	for video, results := range results {
		fmt.Printf("\n=========== Results for video: %s =====\n", video)
		for imageName, result := range results {
			fmt.Printf("    %s   - %s \n", result, imageName)
		}
	}

}
