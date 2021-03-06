package stress

import (
	"context"
	i8ly "github.com/integr8ly/grafana-operator/v3/pkg/apis/integreatly/v1alpha1"
	"github.com/thelastpickle/stress-operator/pkg/apis/thelastpickle/v1alpha1"
	"github.com/thelastpickle/stress-operator/pkg/casskop"
	"github.com/thelastpickle/stress-operator/pkg/monitoring"
	"github.com/thelastpickle/stress-operator/test"
	v1batch "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	"reflect"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"testing"
	"time"
)

var (
	name          = "tlpstress-controller"
	namespace     = "tlpstress"
	namespaceName = types.NamespacedName{
		Namespace: namespace,
		Name:      name,
	}
)

func setupReconcile(t *testing.T, state ...runtime.Object) (*ReconcileStress, reconcile.Result) {
	cl := fake.NewFakeClient(state...)
	r := &ReconcileStress{client: cl, scheme: scheme.Scheme}
	req := reconcile.Request{
		NamespacedName: types.NamespacedName{
			Name:      name,
			Namespace: namespace,
		},
	}
	res, err := r.Reconcile(req)
	if err != nil {
		t.Fatalf("reconcile: (%v)", err)
	}

	return r, res
}

func setupReconcileWithRequeue(t *testing.T, state ...runtime.Object) *ReconcileStress {
	r, res := setupReconcile(t, state...)

	// Check the result of reconciliation to make sure it has the desired state.
	if !res.Requeue {
		t.Error("reconcile did not requeue request as expected")
	}

	return r
}

func setupReconcileWithoutRequeue(t *testing.T, state ...runtime.Object) *ReconcileStress {
	r, res := setupReconcile(t, state...)

	if res.Requeue {
		t.Error("did not expect reconcile to requeue the request")
	}

	return r
}

func TestReconcile(t *testing.T) {
	fdc := test.NewFakeDiscoveryClient()
	monitoring.Init(fdc)
	casskop.Init(fdc)

	test.InitScheme(t)

	t.Run("DefaultsSet", testTLPStressControllerDefaultsSet)
	t.Run("DefaultsNotSet", testTLPStressControllerDefaultsNotSet)
	t.Run("MetricsServiceCreate", testTLPStressControllerMetricsServiceCreate)
	//t.Run("ServiceMonitorCreate", testTLPStressControllerServiceMonitorCreate)
	t.Run("DashboardCreate", testTLPStressControllerDashboardCreate)
	t.Run("JobCreate", testTLPStressControllerJobCreate)
	t.Run("SetStatus", testTLPStressControllerSetStatus)
}

func testTLPStressControllerDefaultsSet(t *testing.T) {
	tlpStress := &v1alpha1.Stress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: v1alpha1.StressSpec{
			CassandraConfig: v1alpha1.CassandraConfig{
				CassandraService: "cassandra-service",
			},
		},
	}

	objs := []runtime.Object{tlpStress}

	r := setupReconcileWithRequeue(t, objs...)

	// make sure defaults are assigned
	instance := &v1alpha1.Stress{}
	err := r.client.Get(context.TODO(), namespaceName, instance)
	if err != nil {
		t.Fatalf("get Stress: (%v)", err)
	}

	if instance.Spec.StressConfig.Workload != DefaultWorkload {
		t.Errorf("Workload (%s) is not the expected value (%s)", instance.Spec.StressConfig.Workload, DefaultWorkload)
	}

	if instance.Spec.Image != DefaultImage {
		t.Errorf("Image (%s) is not the expected value (%s)", instance.Spec.Image, DefaultImage)
	}

	if instance.Spec.ImagePullPolicy != DefaultImagePullPolicy {
		t.Errorf("ImagePullPolicy (%s) is not the expected value (%s)", instance.Spec.ImagePullPolicy, DefaultImagePullPolicy)
	}
}

func testTLPStressControllerDefaultsNotSet(t *testing.T) {
	tlpStress := createStress()

	objs := []runtime.Object{tlpStress}

	r := setupReconcileWithRequeue(t, objs...)

	// make sure defaults are assigned
	instance := &v1alpha1.Stress{}
	err := r.client.Get(context.TODO(), namespaceName, instance)
	if err != nil {
		t.Fatalf("get Stress: (%v)", err)
	}

	if instance.Spec.StressConfig.Workload != tlpStress.Spec.StressConfig.Workload {
		t.Errorf("Workload (%s) is not the expected value (%s)", instance.Spec.StressConfig.Workload,
			tlpStress.Spec.StressConfig.Workload)
	}

	if instance.Spec.Image != tlpStress.Spec.Image {
		t.Errorf("Image (%s) is not the expected value (%s)", instance.Spec.Image, tlpStress.Spec.Image)
	}

	if instance.Spec.ImagePullPolicy != tlpStress.Spec.ImagePullPolicy {
		t.Errorf("ImagePullPolicy (%s) is not the expected value (%s)", instance.Spec.ImagePullPolicy, tlpStress.Spec.ImagePullPolicy)
	}
}

func testTLPStressControllerMetricsServiceCreate(t *testing.T) {
	tlpStress := createStress()

	objs := []runtime.Object{tlpStress}

	r := setupReconcileWithRequeue(t, objs...)

	svc := &corev1.Service{}
	if err := r.client.Get(context.TODO(), types.NamespacedName{Namespace: namespace, Name: monitoring.GetMetricsServiceName(tlpStress)}, svc); err != nil {
		t.Fatalf("get metrics service: (%v)", err)
	}
}

func testTLPStressControllerDashboardCreate(t *testing.T) {
	tlpStress := createStress()

	metricsService := createMetricsService(tlpStress)

	objs := []runtime.Object{tlpStress, metricsService}

	r := setupReconcileWithRequeue(t, objs...)

	dashboard := &i8ly.GrafanaDashboard{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Namespace: tlpStress.Namespace, Name: tlpStress.Name}, dashboard)
	if err != nil {
		t.Fatalf("get dashboard: (%v)", err)
	}
}

func testTLPStressControllerJobCreate(t *testing.T) {
	tlpStress := createStress()
	metricsService :=createMetricsService(tlpStress)
	dashboard := createDashboard(tlpStress)

	objs := []runtime.Object{tlpStress, metricsService, dashboard}

	r := setupReconcileWithRequeue(t, objs...)

	job := &v1batch.Job{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: tlpStress.Name, Namespace: tlpStress.Namespace}, job)
	if err != nil {
		t.Fatalf("get job: (%v)", err)
	}
}

func testTLPStressControllerSetStatus(t *testing.T) {
	tlpStress := createStress()
	metricsService := createMetricsService(tlpStress)
	dashboard := createDashboard(tlpStress)

	job := &v1batch.Job{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Job",
			APIVersion: "batch/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      tlpStress.Name,
			Namespace: tlpStress.Namespace,
		},
	}

	now := metav1.Date(2019, time.April, 7, 7, 7, 7, 0, time.Local)
	job.Status = v1batch.JobStatus{
		Active:    1,
		StartTime: &now,
	}

	objs := []runtime.Object{tlpStress, metricsService, dashboard, job}

	r := setupReconcileWithoutRequeue(t, objs...)

	actual := &v1alpha1.Stress{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: tlpStress.Name, Namespace: tlpStress.Namespace}, actual)
	if err != nil {
		t.Fatalf("get actual: (%v)", err)
	}

	if !reflect.DeepEqual(*actual.Status.JobStatus, job.Status) {
		t.Errorf("actual.Status.JobStatus (%v) does not match the expected value (%+v)", actual.Status.JobStatus, job.Status)
	}
}

func createStress() *v1alpha1.Stress {
	return &v1alpha1.Stress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: v1alpha1.StressSpec{
			CassandraConfig: v1alpha1.CassandraConfig{
				CassandraService: "cassandra-service",
			},
			StressConfig: v1alpha1.StressConfig{
				Workload: v1alpha1.KeyValueWorkload,
			},
			Image:           "jsanda/tlp-stress:demo",
			ImagePullPolicy: corev1.PullAlways,
		},
	}
}

func createMetricsService(tlpStress *v1alpha1.Stress)  *corev1.Service {
	return &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      monitoring.GetMetricsServiceName(tlpStress),
		},
	}
}

func createDashboard(tlpStress *v1alpha1.Stress) *i8ly.GrafanaDashboard {
	return &i8ly.GrafanaDashboard{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name: tlpStress.Name,
		},
	}
}
