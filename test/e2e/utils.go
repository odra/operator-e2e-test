package e2e

import (
	"github.com/operator-framework/operator-sdk/pkg/test/e2eutil"
	"testing"
	framework "github.com/operator-framework/operator-sdk/pkg/test"
	v1alpha1 "github.com/integr8ly/stuff/pkg/apis/example/stuff"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"time"
	"k8s.io/apimachinery/pkg/util/wait"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/api/core/v1"
)

func prepare(t *testing.T) *framework.TestCtx {
	ctx := framework.NewTestCtx(t)
	opt := &framework.CleanupOptions{
		TestContext:ctx,
		RetryInterval: retryInterval,
		Timeout: timeout,
	}

	err := ctx.InitializeClusterResources(opt)
	if err != nil {
		t.Fatalf("Failed to initialize test context: %v", err)
	}

	ns, err := ctx.GetNamespace()
	if err != nil {
		t.Fatalf("Failed to get context namespace: %v", err)
	}

	globalVars := framework.Global

	err = e2eutil.WaitForDeployment(t, globalVars.KubeClient, ns, "stuff", 1, retryInterval, timeout)
	if err != nil {
		t.Fatalf("Operator deployment failed: %v", err)
	}
	return ctx
}

func register() error {
	stuffList := &v1alpha1.StuffList{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Stuff",
			APIVersion: "example.org/stuff",
		},
	}

	err := framework.AddToFrameworkScheme(v1alpha1.AddToScheme, stuffList)
	if err != nil {
		return err
	}

	return nil
}

func cleanupOpts(ctx *framework.TestCtx) *framework.CleanupOptions {
	return &framework.CleanupOptions{
		TestContext: ctx,
		Timeout: timeout,
		RetryInterval: retryInterval,
	}
}

func waitForPod(t *testing.T, kubeclient kubernetes.Interface, namespace, name string, retryInterval, timeout time.Duration) error {
	err := wait.Poll(retryInterval, timeout, func() (done bool, err error) {
		pod, err := kubeclient.CoreV1().Pods(namespace).Get(name, metav1.GetOptions{})
		if err != nil {
			if apierrors.IsNotFound(err) {
				return false, nil
			}
			return false, err
		}

		if pod.Status.Phase == v1.PodRunning {
			return  true, nil
		}

		return false, nil
	})

	return err
}
