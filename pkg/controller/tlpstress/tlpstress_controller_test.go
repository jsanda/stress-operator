package tlpstress

import (
	"context"
	"github.com/jsanda/tlp-stress-operator/pkg/apis/thelastpickle/v1alpha1"
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

func TestTLPStressControllerDefaultsSet(t *testing.T) {
	var (
		name = "tlpstress-controller"
		namespace = "tlpstress"
	)

	tlpStress := &v1alpha1.TLPStress{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
			Namespace: namespace,
		},
		Spec: v1alpha1.TLPStressSpec{
			CassandraService: "cassandra-service",
		},
	}

	objs := []runtime.Object{ tlpStress }

	s := scheme.Scheme
	s.AddKnownTypes(v1alpha1.SchemeGroupVersion, tlpStress)

	cl := fake.NewFakeClient(objs...)

	r := &ReconcileTLPStress{client: cl, scheme: s}

	req := reconcile.Request{
		NamespacedName: types.NamespacedName{
			Name: name,
			Namespace: namespace,
		},
	}

	res, err := r.Reconcile(req)
	if err != nil {
		t.Fatalf("reconcile: (%v)", err)
	}
	// Check the result of reconciliation to make sure it has the desired state.
	if !res.Requeue {
		t.Error("reconcile did not requeue request as expected")
	}

	// make sure defaults are assigned
	instance := &v1alpha1.TLPStress{}
	err = r.client.Get(context.TODO(), req.NamespacedName, instance)
	if err != nil {
		t.Fatalf("get TLPStress: (%v)", err)
	}

	if instance.Spec.Workload != "KeyValue" {
		t.Errorf("Workload (%s) is not the expected value (%s)", instance.Spec.Workload, "KeyValue")
	}

	if instance.Spec.Image != "jsanda/tlp-stress:demo" {
		t.Errorf("Image (%s) is not the expected value (%s)", instance.Spec.Image, "jsanda/tlp-stress:demo")
	}

	if instance.Spec.ImagePullPolicy != corev1.PullAlways {
		t.Errorf("ImagePullPolicy (%s) is not the expected value (%s)", instance.Spec.ImagePullPolicy, corev1.PullAlways)
	}
}

func TestTLPStressControllerDefaultsNotSet(t *testing.T) {
	var (
		name = "tlpstress-controller"
		namespace = "tlpstress"
	)

	tlpStress := &v1alpha1.TLPStress{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
			Namespace: namespace,
		},
		Spec: v1alpha1.TLPStressSpec{
			CassandraService: "cassandra-service",
			Workload: "BasicTimeSeries",
			Image: "jsanda/tlp-stress:test",
			ImagePullPolicy: corev1.PullIfNotPresent,
		},
	}

	objs := []runtime.Object{ tlpStress }

	s := scheme.Scheme
	s.AddKnownTypes(v1alpha1.SchemeGroupVersion, tlpStress)

	cl := fake.NewFakeClient(objs...)

	r := &ReconcileTLPStress{client: cl, scheme: s}

	req := reconcile.Request{
		NamespacedName: types.NamespacedName{
			Name: name,
			Namespace: namespace,
		},
	}

	res, err := r.Reconcile(req)
	if err != nil {
		t.Fatalf("reconcile: (%v)", err)
	}
	// Check the result of reconciliation to make sure it has the desired state.
	if !res.Requeue {
		t.Error("reconcile did not requeue request as expected")
	}

	// make sure defaults are assigned
	instance := &v1alpha1.TLPStress{}
	err = r.client.Get(context.TODO(), req.NamespacedName, instance)
	if err != nil {
		t.Fatalf("get TLPStress: (%v)", err)
	}

	if instance.Spec.Workload != tlpStress.Spec.Workload {
		t.Errorf("Workload (%s) is not the expected value (%s)", instance.Spec.Workload, tlpStress.Spec.Workload)
	}

	if instance.Spec.Image != tlpStress.Spec.Image {
		t.Errorf("Image (%s) is not the expected value (%s)", instance.Spec.Image, tlpStress.Spec.Image)
	}

	if instance.Spec.ImagePullPolicy != tlpStress.Spec.ImagePullPolicy {
		t.Errorf("ImagePullPolicy (%s) is not the expected value (%s)", instance.Spec.ImagePullPolicy, tlpStress.Spec.ImagePullPolicy)
	}
}

func TestTLPStressControllerJobCreate(t *testing.T) {
	var (
		name = "tlpstress-controller"
		namespace = "tlpstress"
	)

	tlpStress := &v1alpha1.TLPStress{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
			Namespace: namespace,
		},
		Spec: v1alpha1.TLPStressSpec{
			CassandraService: "cassandra-service",
			Workload: "KeyValue",
			Image: "jsanda/tlp-stress:demo",
			ImagePullPolicy: corev1.PullAlways,
		},
	}

	objs := []runtime.Object{ tlpStress }

	s := scheme.Scheme
	s.AddKnownTypes(v1alpha1.SchemeGroupVersion, tlpStress)

	cl := fake.NewFakeClient(objs...)

	r := &ReconcileTLPStress{client: cl, scheme: s}

	req := reconcile.Request{
		NamespacedName: types.NamespacedName{
			Name: name,
			Namespace: namespace,
		},
	}

	res, err := r.Reconcile(req)
	if err != nil {
		t.Fatalf("reconcile: (%v)", err)
	}
	// Check the result of reconciliation to make sure it has the desired state.
	if !res.Requeue {
		t.Error("reconcile did not requeue request as expected")
	}

	job := &v1batch.Job{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: tlpStress.Name, Namespace: tlpStress.Namespace}, job)
	if err != nil {
		t.Fatalf("get job: (%v)", err)
	}
}

func TestTLPStressControllerSetStatus(t *testing.T) {
	var (
		name = "tlpstress-controller"
		namespace = "tlpstress"
	)

	tlpStress := &v1alpha1.TLPStress{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
			Namespace: namespace,
		},
		Spec: v1alpha1.TLPStressSpec{
			CassandraService: "cassandra-service",
			Workload: "KeyValue",
			Image: "jsanda/tlp-stress:demo",
			ImagePullPolicy: corev1.PullAlways,
		},
	}

	objs := []runtime.Object{ tlpStress }

	s := scheme.Scheme
	s.AddKnownTypes(v1alpha1.SchemeGroupVersion, tlpStress)

	cl := fake.NewFakeClient(objs...)

	r := &ReconcileTLPStress{client: cl, scheme: s}

	req := reconcile.Request{
		NamespacedName: types.NamespacedName{
			Name: name,
			Namespace: namespace,
		},
	}

	job := &v1batch.Job{
		TypeMeta: metav1.TypeMeta{
			Kind: "Job",
			APIVersion: "batch/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: tlpStress.Name,
			Namespace: tlpStress.Namespace,
		},
	}

	err := r.client.Create(context.TODO(), job)
	if err != nil {
		t.Fatalf("create job: (%v)", err)
	}

	now := metav1.Date(2019, time.April, 7, 7, 7, 7, 0, time.Local)
	job.Status = v1batch.JobStatus{
		Active: 1,
		StartTime: &now,
	}

	if err = r.client.Status().Update(context.TODO(), job); err != nil {
		t.Fatalf("update status: (%v)", err)
	}

	res, err := r.Reconcile(req)
	if err != nil {
		t.Fatalf("reconcile: (%v)", err)
	}
	// Check the result of reconciliation to make sure it has the desired state.
	if res.Requeue {
		t.Error("reconcile should not have requeued the request")
	}

	 actual := &v1alpha1.TLPStress{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: tlpStress.Name, Namespace: tlpStress.Namespace}, actual)
	if err != nil {
		t.Fatalf("get actual: (%v)", err)
	}

	if !reflect.DeepEqual(*actual.Status.JobStatus, job.Status) {
		t.Errorf("actual.Status.JobStatus (%v) does not match the expected value (%+v)", actual.Status.JobStatus, job.Status)
	}
}