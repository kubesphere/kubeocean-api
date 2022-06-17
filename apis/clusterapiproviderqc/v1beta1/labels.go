package v1beta1

const (
	// NameQingCloudProviderPrefix is the tag prefix for
	// cluster-api-provider-qingcloud owned components.
	NameQcloudCloudProviderPrefix = "sigs-k8s-io:capqc"
	// APIServerRoleTagValue describes the value for the apiserver role.
	APIServerRoleTagValue = "apiserver"
	// NodeRoleTagValue describes the value for the node role.
	NodeRoleTagValue = "worker"
	//	ControlPlaneQcMachineTemplateLabel indicate QcMachineTemplate used from cluster control plane
	ControlPlaneQcMachineTemplateLabel = "infrastructure.cluster.x-k8s.io/controlplane-qcmachinetemplate"
)
