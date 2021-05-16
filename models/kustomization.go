package models

type Komponent struct {
	Resource  `yaml:",inline"`
	Resources []string `yaml:",omitempty"`
}

type Kustomization struct {
	Komponent  `yaml:",inline"`
	Components []string `yaml:",omitempty"`
}

func CreateKustomization() *Kustomization {
	rsrc := Kustomization{
		Komponent: Komponent{
			Resource: Resource{
				APIVersion: "kustomize.config.k8s.io/v1beta1",
				Kind:       "Kustomization",
			},
		},
	}
	return &rsrc
}

func CreateKomponent() *Komponent {
	rsrc := Komponent{
		Resource: Resource{
			APIVersion: "kustomize.config.k8s.io/v1alpha1",
			Kind:       "Component",
		},
	}
	return &rsrc
}
