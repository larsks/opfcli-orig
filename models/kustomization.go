package models

import (
    log "github.com/sirupsen/logrus"

    "gopkg.in/yaml.v2"
)

type Kustomization struct {
    Resource                `yaml:",inline"`
    Resources   []string    `yaml:",omitempty"`
    Components  []string    `yaml:",omitempty"`
}

func (rsrc *Kustomization) ToYAML() string {
    s, err := yaml.Marshal(&rsrc)
    if err != nil {
        log.Fatalf("failed converting resource to YAML: %v", err)
    }
    return string(s)
}

func CreateKustomization() *Kustomization {
    rsrc := Kustomization{
        Resource: Resource{
            ApiVersion: "kustomize.config.k8s.io/v1beta1",
            Kind: "Kustomization",
        },
    }
    return &rsrc
}
