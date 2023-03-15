package internal

import (
	"context"
	"fmt"

	flag "github.com/spf13/pflag"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/json"
	kube "k8s.io/cli-runtime/pkg/genericclioptions"
)

type ToggleOptions struct {
	Service string
	Ingress string
	Kube    *kube.ConfigFlags
}

type Patch struct {
	Operation string      `json:"op"`
	Path      string      `json:"path"`
	Value     interface{} `json:"value,omitempty"`
}

func (opts *ToggleOptions) Parse() *flag.FlagSet {
	fs := flag.NewFlagSet("options", flag.ExitOnError)
	opts.Kube = &kube.ConfigFlags{
		KubeConfig:  stringptr(""),
		ClusterName: stringptr(""),
		Context:     stringptr(""),
		Namespace:   stringptr(""),
	}
	namespace, _, err := opts.Kube.ToRawKubeConfigLoader().Namespace()
	if err == nil {
		opts.Kube.Namespace = stringptr(namespace)
	}
	opts.Kube.AddFlags(fs)
	return fs
}

func (opts *ToggleOptions) Patch(patch *Patch) error {
	k8s, err := NewClient(opts.Kube)
	if err != nil {
		return err
	}

	body, err := encode(patch)
	if err != nil {
		return err
	}

	options := metav1.PatchOptions{
		FieldManager: "kubectl-patch",
	}

	if opts.Service != "" {
		svc, err := k8s.
			CoreV1().
			Services(*opts.Kube.Namespace).
			Patch(context.Background(), opts.Service, types.JSONPatchType, body, options)
		if err == nil {
			fmt.Printf("Service %s/%s has been patched\n", svc.Namespace, svc.Name)
		}
	} else if opts.Ingress != "" {
		ingress, err := k8s.
			NetworkingV1().
			Ingresses(*opts.Kube.Namespace).
			Patch(context.Background(), opts.Ingress, types.JSONPatchType, body, options)
		if err == nil {
			fmt.Printf("Ingress %s/%s has been patched\n", ingress.Namespace, ingress.Name)
		}
	}

	return err
}

func encode(patch *Patch) ([]byte, error) {
	data := [1]Patch{*patch}
	return json.Marshal(data)
}

func stringptr(s string) *string {
	return &s
}
