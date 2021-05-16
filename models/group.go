package models

import (
	log "github.com/sirupsen/logrus"

	"gopkg.in/yaml.v2"
)

type Group struct {
	Resource `yaml:",inline"`
	Users    []string
}

func (rsrc *Group) ToYAML() string {
	s, err := yaml.Marshal(&rsrc)
	if err != nil {
		log.Fatalf("failed converting resource to YAML: %v", err)
	}
	return string(s)
}

func CreateGroup(name string) *Group {
	if len(name) == 0 {
		log.Fatal("a group requires a name")
	}

	rsrc := Group{
		Resource: Resource{
			ApiVersion: "user.openshift.io/v1",
			Kind:       "Group",
			Metadata: Metadata{
				Name: name,
			},
		},
		Users: make([]string, 0),
	}
	return &rsrc
}
