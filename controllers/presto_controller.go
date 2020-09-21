/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"reflect"

	"github.com/go-logr/logr"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	skittlesv1 "github.com/quercy/multicloud-k8s-demo-operator/v2/api/v1"
)

// PrestoReconciler reconciles a Presto object
type PrestoReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// func ensureDeployment(log logr.Logger, presto skittlesv1.Presto) {

// }

// +kubebuilder:rbac:groups=skittles.quercy.co,namespace=multicloud-k8s-demo-operator-system,resources=prestoes,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=skittles.quercy.co,namespace=multicloud-k8s-demo-operator-system,resources=prestoes/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apps,namespace=multicloud-k8s-demo-operator-system,resources=deployments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;
// +kubebuilder:rbac:groups=core,resources=deployments,verbs=get;list;
// +kubebuilder:rbac:groups=core,namespace=multicloud-k8s-demo-operator-system,resources=pods,verbs=get;list;
// +kubebuilder:rbac:groups=core,namespace=multicloud-k8s-demo-operator-system,resources=services,verbs=get;list;watch;create;update;patch;delete

func (r *PrestoReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("presto", req.NamespacedName)
	presto := &skittlesv1.Presto{}
	err := r.Get(ctx, req.NamespacedName, presto)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			log.Info("Presto resource not found. Ignoring since object must be deleted")
			return ctrl.Result{}, nil
		}
		// Error reading the object - requeue the request.
		log.Error(err, "Failed to get Presto")
		return ctrl.Result{}, err
	}

	// Check if the deployment already exists, if not create a new one
	deployment := &appsv1.Deployment{}
	err = r.Get(ctx, types.NamespacedName{Name: presto.Name, Namespace: presto.Namespace}, deployment)
	if err != nil && errors.IsNotFound(err) {
		// Define a new deployment
		dep := r.deployPresto(presto)
		log.Info("Creating a new Deployment", "Deployment.Namespace", dep.Namespace, "Deployment.Name", dep.Name)
		err = r.Create(ctx, dep)
		if err != nil {
			log.Error(err, "Failed to create new Deployment", "Deployment.Namespace", dep.Namespace, "Deployment.Name", dep.Name)
			return ctrl.Result{}, err
		}
		// Deployment created successfully - return and requeue
		return ctrl.Result{Requeue: true}, nil
	} else if err != nil {
		log.Error(err, "Failed to get Deployment")
		return ctrl.Result{}, err
	}

	// Ensure the deployment size is the same as the spec
	size := presto.Spec.Workers
	if *deployment.Spec.Replicas != size {
		deployment.Spec.Replicas = &size
		err = r.Update(ctx, deployment)
		if err != nil {
			log.Error(err, "Failed to update Deployment", "Deployment.Namespace", deployment.Namespace, "Deployment.Name", deployment.Name)
			return ctrl.Result{}, err
		}
		// Spec updated - return and requeue
		return ctrl.Result{Requeue: true}, nil
	}

	// Check if the service already exists, if not create a new one
	service := &corev1.Service{}
	err = r.Get(ctx, types.NamespacedName{Name: presto.Name, Namespace: presto.Namespace}, service)
	if err != nil && errors.IsNotFound(err) {
		// Define a new deployment
		dep := r.deployPrestoService(presto)
		log.Info("Creating a new Service", "Service.Namespace", dep.Namespace, "Service.Name", dep.Name)
		err = r.Create(ctx, dep)
		if err != nil {
			log.Error(err, "Failed to create new Service", "Service.Namespace", dep.Namespace, "Service.Name", dep.Name)
			return ctrl.Result{}, err
		}
		// Service created successfully - return and requeue
		return ctrl.Result{Requeue: true}, nil
	} else if err != nil {
		log.Error(err, "Failed to get Service")
		return ctrl.Result{}, err
	}

	// Update the Presto status with the pod names
	// List the pods for this Presto's deployment
	podList := &corev1.PodList{}
	listOpts := []client.ListOption{
		client.InNamespace(presto.Namespace),
		client.MatchingLabels(getPrestoLabels(presto.Name)),
	}
	if err = r.List(ctx, podList, listOpts...); err != nil {
		log.Error(err, "Failed to list pods", "Presto.Namespace", presto.Namespace, "Presto.Name", presto.Name)
		return ctrl.Result{}, err
	}
	podNames := getPodNames(podList.Items)

	// Update status.Nodes if needed
	if !reflect.DeepEqual(podNames, presto.Status.Nodes) {
		presto.Status.Nodes = podNames
		err := r.Status().Update(ctx, presto)
		if err != nil {
			log.Error(err, "Failed to update Presto status")
			return ctrl.Result{}, err
		}
	}

	return ctrl.Result{}, nil
}

// getPodNames returns the pod names of the array of pods passed in
func getPodNames(pods []corev1.Pod) []string {
	var podNames []string
	for _, pod := range pods {
		podNames = append(podNames, pod.Name)
	}
	return podNames
}
func (r *PrestoReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&skittlesv1.Presto{}).
		Complete(r)
}

func getPrestoLabels(name string) map[string]string {
	return map[string]string{"app": "presto", "presto_cr": name}
}

func (r *PrestoReconciler) deployPresto(p *skittlesv1.Presto) *appsv1.Deployment {
	ls := getPrestoLabels(p.Name)
	replicas := p.Spec.Workers

	dep := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      p.Name,
			Namespace: p.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: ls,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: ls,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Image: p.Spec.Image.Repository + ":" + p.Spec.Image.Tag,
						Name:  "presto",
						Ports: []corev1.ContainerPort{{
							ContainerPort: p.Spec.Config.HTTPPort,
							Name:          "http-coord",
						}},
					}},
				},
			},
		},
	}
	// Set Presto instance as the owner and controller
	ctrl.SetControllerReference(p, dep, r.Scheme)
	return dep
}

func (r *PrestoReconciler) deployPrestoService(p *skittlesv1.Presto) *corev1.Service {
	ls := getPrestoLabels(p.Name)

	dep := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      p.Name,
			Namespace: p.Namespace,
		},
		Spec: corev1.ServiceSpec{
			// Type: "ClusterIp",
			Ports: []corev1.ServicePort{
				{
					TargetPort: intstr.IntOrString{IntVal: p.Spec.Config.HTTPPort},
					Protocol:   corev1.ProtocolTCP,
					Port:       p.Spec.Config.HTTPPort,
				},
			},
			Selector: ls,
		},
	}
	// Set Presto instance as the owner and controller
	ctrl.SetControllerReference(p, dep, r.Scheme)
	return dep
}
