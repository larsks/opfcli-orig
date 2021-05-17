package apiresources

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/operate-first/opfcli/models"
)

type APIResource struct {
	Name       string `json:"name"`
	Namespaced bool   `json:"namespaced"`
	Kind       string `json:"kind"`
	Apiversion string `json:"apiVersion"`
	Apigroup   string `json:"apiGroup"`
}

type APIResourceNotFoundError struct {
	Name string
}

func (err *APIResourceNotFoundError) Error() string {
	return fmt.Sprintf("Resource %s not found", err.Name)
}

//go:embed apiresources.json
var apiresourcesJson []byte
var APIResources []APIResource
var APIResourceMap map[string]APIResource = make(map[string]APIResource)

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

func (resource APIResource) String() string {
	out, err := json.MarshalIndent(resource, "", "  ")
	if err != nil {
		panic(err)
	}

	return string(out)
}

func ResourcePath(resource models.Resource) (string, error) {
	var apigroup string

	if strings.Contains(resource.APIVersion, "/") {
		parts := strings.Split(resource.APIVersion, "/")
		apigroup = parts[0]
	} else {
		apigroup = "core"
	}

	key := fmt.Sprintf("%s/%s", apigroup, resource.Kind)
	if spec, ok := APIResourceMap[key]; ok {
		path := fmt.Sprintf(
			"%s/%s/%s/%s",
			spec.Apigroup,
			spec.Name,
			resource.Metadata.Name,
			fmt.Sprintf("%s.yaml", strings.ToLower(spec.Kind)),
		)

		return path, nil

	}
	return "", &APIResourceNotFoundError{Name: key}
}
