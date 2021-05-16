package models

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

func ToYAML(doc interface{}) string {
	s, err := yaml.Marshal(&doc)
	if err != nil {
		log.Fatalf("failed converting resource to YAML: %v", err)
	}
	return string(s)
}
