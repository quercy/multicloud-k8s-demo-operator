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

	"github.com/go-logr/logr"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

// +kubebuilder:rbac:groups=skittles.quercy.co,resources=prestoes,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=skittles.quercy.co,resources=prestoes/status,verbs=get;update;patch

func (r *PrestoReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("presto", req.NamespacedName)

	// presto := &skittlesv1.Presto{}
	// err := r.Get(ctx, req.NamespacedName, presto)
	// found := &appsv1.Deployment{}
	// err = r.Get(ctx, types.NamespacedName{Name: presto.Name, Namespace: presto.Namespace}, found)
	// if err != nil && errors.IsNotFound(err) {
	// 	// Define a new deployment
	// 	dep := r.deployPresto(presto)
	// 	log.Info("Creating a new Deployment", "Deployment.Namespace", dep.Namespace, "Deployment.Name", dep.Name)
	// 	err = r.Create(ctx, dep)
	// 	if err != nil {
	// 		log.Error(err, "Failed to create new Deployment", "Deployment.Namespace", dep.Namespace, "Deployment.Name", dep.Name)
	// 		return ctrl.Result{}, err
	// 	}
	// 	// Deployment created successfully - return and requeue
	// 	return ctrl.Result{Requeue: true}, nil
	// } else if err != nil {
	// 	log.Error(err, "Failed to get Deployment")
	// 	return ctrl.Result{}, err
	// }

	return ctrl.Result{}, nil
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
