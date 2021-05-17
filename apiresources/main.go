package apiresources

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"strings"
)

type Resource struct {
	Name       string `json:"name"`
	Namespaced bool   `json:"namespaced"`
	Kind       string `json:"kind"`
	Apiversion string `json:"apiVersion"`
	Apigroup   string `json:"apiGroup"`
}

//go:embed apiresources.json
var apiresourcesJson []byte
var APIResources []Resource
var APIResourceMap map[string]Resource = make(map[string]Resource)

func init() {
	json.Unmarshal(apiresourcesJson, &APIResources)

	for _, resource := range APIResources {
		if strings.Contains(resource.Name, "/") {
			continue
		}

		key := fmt.Sprintf("%s/%s", resource.Apigroup, resource.Kind)
		APIResourceMap[key] = resource
	}
}

func (resource Resource) String() string {
	out, err := json.MarshalIndent(resource, "", "  ")
	if err != nil {
		panic(err)
	}

	return string(out)
}
