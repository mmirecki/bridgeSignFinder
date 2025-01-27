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
	chatGPT.QueryDir(outputPath)

}
