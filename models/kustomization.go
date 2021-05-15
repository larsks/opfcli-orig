package models

import (
    "gopkg.in/yaml.v2"
)

type Komponent struct {
    Resource    `yaml:",inline"`
    Resources   []string
}

type Kustomization struct {
    Komponent
    Components []string
}

func (rsrc *Komponent) ToYAML() (string, error) {
    s, err := yaml.Marshal(&rsrc)
    return string(s), err
}

func (rsrc *Kustomization) ToYAML() (string, error) {
    s, err := yaml.Marshal(&rsrc)
    return string(s), err
}

func CreateKomponent() *Komponent {
    rsrc := Komponent{
        Resource: Resource{
            ApiVersion: "kustomize.config.k8s.io/v1beta1",
            Kind: "Component",
        },
    }
    return &rsrc
}

func CreateKustomization() *Kustomization {
    rsrc := Kustomization{
        Komponent: Komponent{
            Resource: Resource{
                ApiVersion: "kustomize.config.k8s.io/v1beta1",
                Kind: "Kustomization",
            },
        },
    }
    return &rsrc
}
