package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	varsJson := os.Getenv("VARS_JSON")
	varsProp := make(map[string]string)
	err := json.Unmarshal([]byte(varsJson), &varsProp)
	if err != nil {
		log.Fatal("Failed to parse JSON:", err)
	}
	for key, value := range varsProp {
		if propName, ok := strings.CutPrefix(key, "BUILD_"); ok {
			propValue := value
			fmt.Printf("%s=%s\n", propName, propValue)
		}
	}
}
