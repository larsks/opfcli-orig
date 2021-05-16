package models

import (
	log "github.com/sirupsen/logrus"

	"gopkg.in/yaml.v2"
)

type Subject struct {
	ApiGroup string `yaml:"apiGroup"`
	Kind     string
	Name     string
}

type RoleBinding struct {
	Resource `yaml:",inline"`
	RoleRef  Subject `yaml:"roleRef"`
	Subjects []Subject
}

func (rsrc *RoleBinding) ToYAML() string {
	s, err := yaml.Marshal(&rsrc)
	if err != nil {
		log.Fatalf("failed converting resource to YAML: %v", err)
	}
	return string(s)
}

func CreateRoleBinding(name string, role string) *RoleBinding {
	if len(name) == 0 {
		log.Fatal("a group requires a name")
	}

	rsrc := RoleBinding{
		Resource: Resource{
			ApiVersion: "rbac.authorization.k8s.io/v1",
			Kind:       "RoleBinding",
			Metadata: Metadata{
				Name: name,
			},
		},
		RoleRef: Subject{
			ApiGroup: "rbac.authorization.k8s.io",
			Kind:     "ClusterRole",
			Name:     role,
		},
		Subjects: make([]Subject, 0),
	}
	return &rsrc
}

func CreateGroupSubject(name string) *Subject {
	rsrc := Subject{
		ApiGroup: "rbac.authorization.k8s.io",
		Kind:     "Group",
		Name:     name,
	}

	return &rsrc
}
