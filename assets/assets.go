package assets

import (
	"embed"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
)

var (
	// go:embed manifests/*
	manifests  embed.FS
	appSchema  = runtime.NewScheme()
	appsCodecs = serializer.NewCodecFactory(appSchema)
)

func init() {
	if err := appsv1.AddToScheme(appSchema); err != nil {
		panic(err)
	}
}

func GetDeploymemtFromFile(name string) *appsv1.Deployment {
	deploymentBytes, err := manifests.ReadFile(name)
	if err != nil {
		panic(err)
	}

	deploymentObject, err := runtime.Decode(
		appsCodecs.UniversalDecoder(appsv1.SchemaGroupVersion),
		deploymentBytes,
	)

	if err != nil {
		panic(err)
	}
	return deploymentObject.(*appsv1.Deployment)
}
