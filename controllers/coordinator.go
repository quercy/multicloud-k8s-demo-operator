package controllers

import (
	"context"

	"github.com/prometheus/common/log"
	skittlesv1 "github.com/quercy/multicloud-k8s-demo-operator/v2/api/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	ctrl "sigs.k8s.io/controller-runtime"
)

func (r *PrestoReconciler) ensureCoordinator(ctx context.Context, presto *skittlesv1.Presto) (ctrl.Result, error) {
	// Check if the deployment already exists, if not create a new one
	deployment := &appsv1.Deployment{}
	err := r.Get(ctx, types.NamespacedName{Name: presto.Name, Namespace: presto.Namespace}, deployment)
	if err != nil && errors.IsNotFound(err) {
		// Define a new deployment
		dep := r.createCoordinatorDeployment(presto)
		// Set Presto instance as the owner and controller
		ctrl.SetControllerReference(presto, dep, r.Scheme)
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
		dep := r.createCoordinatorService(presto)
		// Set Presto instance as the owner and controller
		ctrl.SetControllerReference(presto, dep, r.Scheme)
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

	return ctrl.Result{}, nil
}

func (r *PrestoReconciler) createCoordinatorDeployment(p *skittlesv1.Presto) *appsv1.Deployment {
	log.Info("Creating a new Deployment", "Deployment.Namespace", p.Namespace, "Deployment.Name", p.Name)
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
	return dep
}

func (r *PrestoReconciler) createCoordinatorService(p *skittlesv1.Presto) *corev1.Service {
	ls := getPrestoLabels(p.Name)

	dep := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      p.Name,
			Namespace: p.Namespace,
		},
		Spec: corev1.ServiceSpec{
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
	return dep
}
