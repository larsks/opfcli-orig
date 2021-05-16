package models

import (
	log "github.com/sirupsen/logrus"
)

type Group struct {
	Resource `yaml:",inline"`
	Users    []string
}

func CreateGroup(name string) *Group {
	if len(name) == 0 {
		log.Fatal("a group requires a name")
	}

	rsrc := Group{
		Resource: Resource{
			APIVersion: "user.openshift.io/v1",
			Kind:       "Group",
			Metadata: Metadata{
				Name: name,
			},
		},
		Users: make([]string, 0),
	}
	return &rsrc
}
