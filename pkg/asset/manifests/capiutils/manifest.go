package capiutils

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/openshift/installer/pkg/asset"
)

// GenerateClusterAssetsOutput is the output of GenerateClusterAssets.
type GenerateClusterAssetsOutput struct {
	Manifests         []*asset.RuntimeFile
	InfrastructureRef *corev1.ObjectReference
}
