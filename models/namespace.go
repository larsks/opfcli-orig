package models

import (
	log "github.com/sirupsen/logrus"
)

type Namespace struct {
	Resource `yaml:",inline"`
}

func CreateNamespace(name, owner, description string) *Namespace {
	if len(name) == 0 {
		log.Fatal("a namespace requires a name")
	}

	if len(owner) == 0 {
		log.Fatal("a namespace requires an owner")
	}

	rsrc := Namespace{
		Resource: Resource{
			ApiVersion: "v1",
			Kind:       "Namespace",
			Metadata: Metadata{
				Name:        name,
				Annotations: make(map[string]string),
			},
		},
	}
	rsrc.Metadata.Annotations["openshift.io/requester"] = owner
	if len(description) > 0 {
		rsrc.Metadata.Annotations["openshift.io/display-name"] = description
	}

	return &rsrc
}
