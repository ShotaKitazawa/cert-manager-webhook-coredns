module github.com/ShotaKitazawa/cert-manager-webhook-coredns

go 1.14

require (
	github.com/coreos/etcd v3.3.15+incompatible // indirect
	github.com/google/go-github v17.0.0+incompatible // indirect
	github.com/jetstack/cert-manager v1.0.3
	go.etcd.io/etcd v0.5.0-alpha.5.0.20200819165624-17cef6e3e9d5
	gonum.org/v1/netlib v0.0.0-20190331212654-76723241ea4e // indirect
	k8s.io/apiextensions-apiserver v0.19.0
	k8s.io/client-go v0.19.0
	k8s.io/klog v1.0.0
	sigs.k8s.io/structured-merge-diff v1.0.1-0.20191108220359-b1b620dd3f06 // indirect
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
