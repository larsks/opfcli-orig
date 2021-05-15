package models

import (
    "gopkg.in/yaml.v2"
)

type kvmap map[string]string

type Metadata struct {
    Name string
    Annotations kvmap           `yaml:",omitempty"`
    Labels kvmap                `yaml:",omitempty"`
}

type Resource struct {
    ApiVersion string           `yaml:"apiVersion"`
    Kind string
    Metadata Metadata           `yaml:",omitempty"`
}

type ResourceImpl interface {
    ToYAML() string
}

func (resource *Resource) ToYAML() (string, error) {
    s, err := yaml.Marshal(&resource)
    return string(s), err
}
