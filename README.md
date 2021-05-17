# opfcli

## Examples

### Create a project

```
opfcli create-project project1 group1 -d "This is project1"
```

This will result in:

```
cluster-scope/
├── base
│   ├── core
│   │   └── namespaces
│   │       └── project1
│   │           ├── kustomization.yaml
│   │           └── namespace.yaml
│   └── user.openshift.io
│       └── groups
│           └── group1
│               ├── group.yaml
│               └── kustomization.yaml
└── components
    └── project-admin-rolebindings
        └── group1
            ├── kustomization.yaml
            └── rbac.yaml
```

### Create a group

```
opfcli create-group group2
```

This will result in:

```
cluster-scope/
└── base
    └── user.openshift.io
        └── groups
            └── group1
                ├── group.yaml
                └── kustomization.yaml
```

### Grant access to a project

```
opfcli grant-access project1 group2 view
```

This will result in:

```
cluster-scope/components/project-view-rolebindings/
└── group2
    ├── kustomization.yaml
    └── rbac.yaml
```

(And will modify
`cluster-scope/base/core/namespaces/project1/kustomization.yaml`)

