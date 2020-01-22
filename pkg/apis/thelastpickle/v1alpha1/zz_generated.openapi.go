// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/jsanda/tlp-stress-operator/pkg/apis/thelastpickle/v1alpha1.TLPStress":              schema_pkg_apis_thelastpickle_v1alpha1_TLPStress(ref),
		"github.com/jsanda/tlp-stress-operator/pkg/apis/thelastpickle/v1alpha1.TLPStressContext":       schema_pkg_apis_thelastpickle_v1alpha1_TLPStressContext(ref),
		"github.com/jsanda/tlp-stress-operator/pkg/apis/thelastpickle/v1alpha1.TLPStressContextSpec":   schema_pkg_apis_thelastpickle_v1alpha1_TLPStressContextSpec(ref),
		"github.com/jsanda/tlp-stress-operator/pkg/apis/thelastpickle/v1alpha1.TLPStressContextStatus": schema_pkg_apis_thelastpickle_v1alpha1_TLPStressContextStatus(ref),
		"github.com/jsanda/tlp-stress-operator/pkg/apis/thelastpickle/v1alpha1.TLPStressSpec":          schema_pkg_apis_thelastpickle_v1alpha1_TLPStressSpec(ref),
		"github.com/jsanda/tlp-stress-operator/pkg/apis/thelastpickle/v1alpha1.TLPStressStatus":        schema_pkg_apis_thelastpickle_v1alpha1_TLPStressStatus(ref),
	}
}

func schema_pkg_apis_thelastpickle_v1alpha1_TLPStress(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TLPStress is the Schema for the tlpstresses API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/jsanda/tlp-stress-operator/pkg/apis/thelastpickle/v1alpha1.TLPStressSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/jsanda/tlp-stress-operator/pkg/apis/thelastpickle/v1alpha1.TLPStressStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/jsanda/tlp-stress-operator/pkg/apis/thelastpickle/v1alpha1.TLPStressSpec", "github.com/jsanda/tlp-stress-operator/pkg/apis/thelastpickle/v1alpha1.TLPStressStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_thelastpickle_v1alpha1_TLPStressContext(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TLPStressContext is the Schema for the tlpstresscontexts API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/jsanda/tlp-stress-operator/pkg/apis/thelastpickle/v1alpha1.TLPStressContextSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/jsanda/tlp-stress-operator/pkg/apis/thelastpickle/v1alpha1.TLPStressContextStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/jsanda/tlp-stress-operator/pkg/apis/thelastpickle/v1alpha1.TLPStressContextSpec", "github.com/jsanda/tlp-stress-operator/pkg/apis/thelastpickle/v1alpha1.TLPStressContextStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_thelastpickle_v1alpha1_TLPStressContextSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TLPStressContextSpec defines the desired state of TLPStressContext",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"installPrometheus": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"installGrafana": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
				},
			},
		},
	}
}

func schema_pkg_apis_thelastpickle_v1alpha1_TLPStressContextStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TLPStressContextStatus defines the observed state of TLPStressContext",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_thelastpickle_v1alpha1_TLPStressSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TLPStressSpec defines the desired state of TLPStress",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"cassandraConfig": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/jsanda/tlp-stress-operator/pkg/apis/thelastpickle/v1alpha1.CassandraConfig"),
						},
					},
					"image": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"imagePullPolicy": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"stressConfig": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/jsanda/tlp-stress-operator/pkg/apis/thelastpickle/v1alpha1.TLPStressConfig"),
						},
					},
					"jobConfig": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/jsanda/tlp-stress-operator/pkg/apis/thelastpickle/v1alpha1.JobConfig"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/jsanda/tlp-stress-operator/pkg/apis/thelastpickle/v1alpha1.CassandraConfig", "github.com/jsanda/tlp-stress-operator/pkg/apis/thelastpickle/v1alpha1.JobConfig", "github.com/jsanda/tlp-stress-operator/pkg/apis/thelastpickle/v1alpha1.TLPStressConfig"},
	}
}

func schema_pkg_apis_thelastpickle_v1alpha1_TLPStressStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TLPStressStatus defines the observed state of TLPStress",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"jobStatus": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/api/batch/v1.JobStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/api/batch/v1.JobStatus"},
	}
}
