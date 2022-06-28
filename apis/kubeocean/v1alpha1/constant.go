package v1alpha1

const (
	KsInstalling                = "installing"
	KsInstalled                 = "installed"
	KsInstallFailed             = "failed"
	DolphinServiceProtocolHttp  = "http"
	DolphinServiceProtocolHttps = "https"
	//	supported kubesphere version
	Kubesphere3_2_1Version = "kubesphere 3.2.1"
	Kubesphere3_3Version   = "kubesphere 3.3"
)

var (
	SupportedKubesphereVersions = []string{Kubesphere3_2_1Version, Kubesphere3_3Version}
)
