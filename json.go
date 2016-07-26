package acogo

import (
	"os"
	"fmt"
	"encoding/json"
)

func WriteJsonFile(obj interface{}, outputDirectory string, outputJsonFileName string) {
	if err := os.MkdirAll(outputDirectory, os.ModePerm); err != nil {
		panic(fmt.Sprintf("Cannot create output directory (%s)", err))
	}

	fp, err := os.Create(outputDirectory + "/" + outputJsonFileName)

	if err != nil {
		panic(fmt.Sprintf("Cannot create JSON file (%s)", err))
	}

	defer fp.Close()

	j, err := json.MarshalIndent(obj, "", "  ")

	if err != nil {
		panic(fmt.Sprintf("Cannot encode object as JSON (%s)", err))
	}

	if _, err := fp.Write(j); err != nil {
		panic(fmt.Sprintf("Cannot write JSON file (%s)", err))
	}
}
