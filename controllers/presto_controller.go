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
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
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

// +kubebuilder:rbac:groups=skittles.quercy.co,namespace=multicloud-k8s-demo-operator-system,resources=prestoes,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=skittles.quercy.co,namespace=multicloud-k8s-demo-operator-system,resources=prestoes/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=apps,namespace=multicloud-k8s-demo-operator-system,resources=deployments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,namespace=multicloud-k8s-demo-operator-system,resources=services,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,namespace=multicloud-k8s-demo-operator-system,resources=pods,verbs=get;list;watch

func (r *PrestoReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("presto", req.NamespacedName)
	presto := &skittlesv1.Presto{}
	err := r.Get(ctx, req.NamespacedName, presto)
	res := ctrl.Result{}
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

	log.Info("Ensuring the coordinator")
	res, err = r.ensureCoordinator(ctx, presto)
	if err != nil || res.Requeue == true {
		return res, err
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

	return res, nil
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
