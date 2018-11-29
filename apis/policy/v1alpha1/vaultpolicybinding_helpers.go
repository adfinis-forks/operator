package v1alpha1

import (
	"fmt"

	crdutils "github.com/appscode/kutil/apiextensions/v1beta1"
	meta_util "github.com/appscode/kutil/meta"
	"github.com/kubevault/operator/apis"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
)

func (v VaultPolicyBinding) GetKey() string {
	return ResourceVaultPolicyBinding + "/" + v.Namespace + "/" + v.Name
}

func (v VaultPolicyBinding) PolicyBindingName() string {
	if v.Spec.RoleName != "" {
		return v.Spec.RoleName
	}

	cluster := "-"
	if v.ClusterName != "" {
		cluster = v.ClusterName
	}
	return fmt.Sprintf("k8s.%s.%s.%s", cluster, v.Namespace, v.Name)
}

func (v VaultPolicyBinding) OffshootSelectors() map[string]string {
	return map[string]string{
		"app": "vault",
		"vault_policy_binding": v.Name,
	}
}

func (v VaultPolicyBinding) OffshootLabels() map[string]string {
	return meta_util.FilterKeys("kubevault.com", v.OffshootSelectors(), v.Labels)
}

func (v VaultPolicyBinding) CustomResourceDefinition() *apiextensions.CustomResourceDefinition {
	return crdutils.NewCustomResourceDefinition(crdutils.Config{
		Group:         SchemeGroupVersion.Group,
		Plural:        ResourceVaultPolicyBindings,
		Singular:      ResourceVaultPolicyBinding,
		Kind:          ResourceKindVaultPolicyBinding,
		ShortNames:    []string{"vpb"},
		Categories:    []string{"vault", "appscode", "all"},
		ResourceScope: string(apiextensions.NamespaceScoped),
		Versions: []apiextensions.CustomResourceDefinitionVersion{
			{
				Name:    SchemeGroupVersion.Version,
				Served:  true,
				Storage: true,
			},
		},
		Labels: crdutils.Labels{
			LabelsMap: map[string]string{"app": "vault"},
		},
		SpecDefinitionName:      "github.com/kubevault/operator/apis/policy/v1alpha1.VaultPolicyBinding",
		EnableValidation:        true,
		GetOpenAPIDefinitions:   GetOpenAPIDefinitions,
		EnableStatusSubresource: apis.EnableStatusSubresource,
		AdditionalPrinterColumns: []apiextensions.CustomResourceColumnDefinition{
			{
				Name:     "Status",
				Type:     "string",
				JSONPath: ".status.status",
			},
			{
				Name:     "Age",
				Type:     "date",
				JSONPath: ".metadata.creationTimestamp",
			},
		},
	}, apis.SetNameSchema)
}

func (v VaultPolicyBinding) IsValid() error {
	return nil
}

func (v *VaultPolicyBinding) SetDefaults() {
	if v == nil {
		return
	}

	if v.Spec.RoleName == "" {
		v.Spec.RoleName = v.PolicyBindingName()
	}

	if v.Spec.AuthPath == "" {
		v.Spec.AuthPath = "kubernetes"
	}
}