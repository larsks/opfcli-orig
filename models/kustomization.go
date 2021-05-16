package models

type Kustomization struct {
	Resource   `yaml:",inline"`
	Resources  []string `yaml:",omitempty"`
	Components []string `yaml:",omitempty"`
}

func CreateKustomization() *Kustomization {
	rsrc := Kustomization{
		Resource: Resource{
			APIVersion: "kustomize.config.k8s.io/v1beta1",
			Kind:       "Kustomization",
		},
	}
	return &rsrc
}
