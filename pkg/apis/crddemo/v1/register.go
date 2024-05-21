package v1

import (
	"crddemo/pkg/apis/crddemo"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var SchemeGroupVersion = schema.GroupVersion{
	Group:   crddemo.GroupName,
	Version: crddemo.Version,
}

var (
	SchemeBuilder = runtime.NewSchemeBuilder()
	AddToScheme   = SchemeBuilder.AddToScheme
)

func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

func addKownType(schema *runtime.Scheme) error {
	schema.AddKnownTypes(SchemeGroupVersion,
		&Mydemo{},
		&MydemoList{},
	)
	metav1.AddToGroupVersion(schema, SchemeGroupVersion)
	return nil
}
