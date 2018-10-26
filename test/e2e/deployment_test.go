package e2e

import(
	goctx "context"
	"testing"
	framework "github.com/operator-framework/operator-sdk/pkg/test"
	v1alpha1 "github.com/integr8ly/stuff/pkg/apis/example/stuff"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestDeployment(t *testing.T){
	var err error
	t.Parallel()

	ctx := prepare(t)
	defer ctx.Cleanup()

	err = register()
	if err != nil {
		t.Fatalf("Failed to register crd scheme: %v", err)
	}

	err = createCr(framework.Global, ctx)
	if err != nil {
		t.Fatalf("Failed to create cr: %v", err)
	}

	err = validate(t, framework.Global, ctx)
	if err != nil {
		t.Fatalf("Failed to create cr: %v", err)
	}
}

func createCr(f *framework.Framework, ctx *framework.TestCtx) error {
	ns, err := ctx.GetNamespace()
	if err != nil {
		return err
	}

	cr := &v1alpha1.Stuff{
		TypeMeta: metav1.TypeMeta{
			Kind: "Stuff",
			APIVersion: "example.org/stuff",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "my-stuff",
			Namespace: ns,
		},
		Spec:v1alpha1.StuffSpec{},
	}

	err = f.Client.Create(goctx.TODO(), cr, cleanupOpts(ctx))
	if err != nil {
		return err
	}

	return nil
}

func validate(t *testing.T, f *framework.Framework, ctx *framework.TestCtx) error {
	var err error
	ns, err := ctx.GetNamespace()
	if err != nil {
		return err
	}

	err = waitForPod(t, f.KubeClient, ns, "busy-box", retryInterval, timeout)
	if err != nil {
		return err
	}

	return nil

}
