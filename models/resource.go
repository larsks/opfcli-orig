package models

type kvmap map[string]string

type Metadata struct {
	Name        string
	Annotations kvmap `yaml:",omitempty"`
	Labels      kvmap `yaml:",omitempty"`
}

type Resource struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string
	Metadata   Metadata `yaml:",omitempty"`
}
