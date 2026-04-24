package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type SingBoxSourceFile struct {
	Version uint8        `json:"version"`
	Rules   []DomainRule `json:"rules"`
}

type DomainRule struct {
	Domains []string `json:"domain"`
}

func main() {
	var fileNames = []string{
		"iplist-main.json",
		"iplist-beta.json",
		"iplist-russian.json",
	}
	for _, fileName := range fileNames {
		fileBytes, err := os.ReadFile(fileName)
		if err != nil {
			panic(err)
		}

		// map[domain_group][]domains
		var fileData map[string][]string
		err = json.Unmarshal(fileBytes, &fileData)
		if err != nil {
			panic(err)
		}

		var RuleSetOut SingBoxSourceFile
		RuleSetOut.Version = 4
		RuleSetOut.Rules = make([]DomainRule, 0)
		var fileRules DomainRule
		fileRules.Domains = make([]string, 0)
		for _, domains := range fileData {
			fileRules.Domains = append(fileRules.Domains, domains...)
		}

		RuleSetOut.Rules = append(RuleSetOut.Rules, fileRules)

		outputBytes, err := json.MarshalIndent(RuleSetOut, "", "\t")
		if err != nil {
			panic(err)
		}
		file, err := os.OpenFile(fmt.Sprintf("source-%s", fileName), os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		_, err = file.Write(outputBytes)
		if err != nil {
			panic(err)
		}
	}
}
