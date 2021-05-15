package models

import (
    "log"

    "gopkg.in/yaml.v2"
)

type Group struct {
    Resource    `yaml:",inline"`
    Users []string
}

func (rsrc *Group) ToYAML() (string, error) {
    s, err := yaml.Marshal(&rsrc)
    return string(s), err
}

func CreateGroup(name string) *Group {
    if len(name) == 0 {
        log.Fatal("a group requires a name")
    }

    rsrc := Group{
        Resource: Resource{
            ApiVersion: "user.openshift.io/v1",
            Kind: "Group",
            Metadata: Metadata{
                Name: name,
            },
        },
        Users: make([]string, 0),
    }
    return &rsrc
}
