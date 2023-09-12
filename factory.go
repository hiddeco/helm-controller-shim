package shim

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	kuberecorder "k8s.io/client-go/tools/record"
	"sigs.k8s.io/cli-utils/pkg/kstatus/polling"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/fluxcd/helm-controller/internal/controller"
	runtimeClient "github.com/fluxcd/pkg/runtime/client"
	helper "github.com/fluxcd/pkg/runtime/controller"
)

type HelmReleaseReconcilerFactory struct {
	client.Client
	helper.Metrics

	Config                *rest.Config
	Scheme                *runtime.Scheme
	EventRecorder         kuberecorder.EventRecorder
	DefaultServiceAccount string
	NoCrossNamespaceRef   bool
	ClientOpts            runtimeClient.Options
	KubeConfigOpts        runtimeClient.KubeConfigOptions
	StatusPoller          *polling.StatusPoller
	PollingOpts           polling.Options
	ControllerName        string
}

type HelmReleaseReconcilerOptions controller.HelmReleaseReconcilerOptions

func (f *HelmReleaseReconcilerFactory) SetupWithManager(ctx context.Context, mgr ctrl.Manager, opts HelmReleaseReconcilerOptions) error {
	r := &controller.HelmReleaseReconciler{
		Client:                f.Client,
		Metrics:               f.Metrics,
		Config:                f.Config,
		Scheme:                f.Scheme,
		EventRecorder:         f.EventRecorder,
		DefaultServiceAccount: f.DefaultServiceAccount,
		NoCrossNamespaceRef:   f.NoCrossNamespaceRef,
		ClientOpts:            f.ClientOpts,
		KubeConfigOpts:        f.KubeConfigOpts,
		StatusPoller:          f.StatusPoller,
		PollingOpts:           f.PollingOpts,
		ControllerName:        f.ControllerName,
	}
	return r.SetupWithManager(ctx, mgr, controller.HelmReleaseReconcilerOptions(opts))
}
