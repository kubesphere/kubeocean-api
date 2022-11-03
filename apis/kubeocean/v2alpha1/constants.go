package v2alpha1

const (
	ClusterNameLabel = "kubeocean.kubesphere.io/cluster-name"
	ClusterIDLabel   = "kubeocean.kubesphere.io/ks-cloud-cluster-id"
	UserIDLabel      = "kubeocean.kubesphere.io/ks-cloud-user-id"

	TagSubscriptionPrefix     = "subscription-"
	TagSubscriptionBasic      = TagSubscriptionPrefix + "basic"
	TagSubscriptionStandard   = TagSubscriptionPrefix + "standard"
	TagSubscriptionEnterprise = TagSubscriptionPrefix + "enterprise"
	TagSubscriptionCustomized = TagSubscriptionPrefix + "customized"
)
