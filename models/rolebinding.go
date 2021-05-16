package models

import (
	log "github.com/sirupsen/logrus"
)

type Subject struct {
	APIGroup string `yaml:"apiGroup"`
	Kind     string
	Name     string
}

type RoleBinding struct {
	Resource `yaml:",inline"`
	RoleRef  Subject `yaml:"roleRef"`
	Subjects []Subject
}

func CreateRoleBinding(name string, role string) *RoleBinding {
	if len(name) == 0 {
		log.Fatal("a group requires a name")
	}

	rsrc := RoleBinding{
		Resource: Resource{
			APIVersion: "rbac.authorization.k8s.io/v1",
			Kind:       "RoleBinding",
			Metadata: Metadata{
				Name: name,
			},
		},
		RoleRef: Subject{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "ClusterRole",
			Name:     role,
		},
		Subjects: make([]Subject, 0),
	}
	return &rsrc
}

func CreateGroupSubject(groupName string) *Subject {
	rsrc := Subject{
		APIGroup: "rbac.authorization.k8s.io",
		Kind:     "Group",
		Name:     groupName,
	}

	return &rsrc
}

func (rolebinding *RoleBinding) AddGroup(groupName string) {
	sub := CreateGroupSubject(groupName)
	if len(rolebinding.Subjects) == 0 {
		rolebinding.Subjects = []Subject{*sub}
	} else {
		rolebinding.Subjects = append(rolebinding.Subjects, *sub)
	}
}
