package video2image

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const video2ImageCommand = "ffmpeg -i %s -r 1 %s"

func CheckSamples(videoPath, outputPath string) error {

	files, err := os.ReadDir(videoPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		fileName := file.Name()
		if strings.Contains(fileName, ".mp4") {
			err := Video2Images(fileName, videoPath, outputPath)
			if err != nil {
				fmt.Printf("Error processing file: %s", fileName)
			}

			break
		}

	}
	return nil
}

func Video2Images(videoFile string, videoFilePath string, outputPath string) error {

	// photos/output_%s/%05d.png

	inputFile := fmt.Sprintf("%s/%s", videoFilePath, videoFile)

	// TODO: assuming there is an extension on filename
	outputDir := strings.Split(videoFile, ".")[0]
	outputDir = fmt.Sprintf("%s/%s", outputPath, outputDir)
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		os.Mkdir(outputDir, 0777)
	}
	outputFilePattern := fmt.Sprintf("%s/%%05d.png", outputDir)

	c := exec.Command(
		"ffmpeg", "-i", inputFile, "-r", "1", outputFilePattern,
	)
	//c.Stderr = os.Stderr
	err := c.Run()
	if err != nil {
		fmt.Printf("Error running command: %+v", err)
		fmt.Println(err)
	}
	return err
}
