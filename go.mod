module github.com/operator-framework/k8s-watcher

require (
	github.com/NYTimes/gziphandler v1.0.1 // indirect
	github.com/ashwanthkumar/slack-go-webhook v0.0.0-20181208062437-4a19b1a876b7
	github.com/go-openapi/spec v0.19.0
	github.com/maxchagin/go-memorycache-example v0.0.0-20180524131321-f4bebbb0aa08
	github.com/moul/http2curl v1.0.0 // indirect
	github.com/operator-framework/operator-sdk v0.9.1-0.20190807211748-fe8c81ad98c0
	github.com/parnurzeal/gorequest v0.2.15 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/spf13/pflag v1.0.3
	k8s.io/api v0.0.0-20190612125737-db0771252981
	k8s.io/apimachinery v0.0.0-20190612125636-6a5db36e93ad
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/kube-openapi v0.0.0-20190603182131-db7b694dc208
	sigs.k8s.io/controller-runtime v0.1.12
	sigs.k8s.io/controller-tools v0.1.10
)

// Pinned to kubernetes-1.13.4
replace (
	k8s.io/api => k8s.io/api v0.0.0-20190222213804-5cb15d344471
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20190228180357-d002e88f6236
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190221213512-86fb29eff628
	k8s.io/client-go => k8s.io/client-go v0.0.0-20190228174230-b40b2a5939e4
)

replace (
	github.com/coreos/prometheus-operator => github.com/coreos/prometheus-operator v0.29.0
	k8s.io/kube-state-metrics => k8s.io/kube-state-metrics v1.6.0
	sigs.k8s.io/controller-runtime => sigs.k8s.io/controller-runtime v0.1.12
	sigs.k8s.io/controller-tools => sigs.k8s.io/controller-tools v0.1.11-0.20190411181648-9d55346c2bde
)
