package resource

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	opservicev1alpha1 "github.com/cdl4566/operator-service/api/v1alpha1"
)

func NewDeploy(app *opservicev1alpha1.OpServiceApplication) *appsv1.Deployment {
	labels := map[string]string{"app": app.Name}
	selector := &metav1.LabelSelector{MatchLabels: labels}
	replica := int32(app.Spec.Replica)
	return &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Deployment",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      app.Name,
			Namespace: app.Namespace,

			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(app, schema.GroupVersionKind{
					Group:   app.GroupVersionKind().Group,
					Version: app.GroupVersionKind().Version,
					Kind:    app.GroupVersionKind().Kind,
				}),
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replica,
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: newContainers(app),
				},
			},
			Selector: selector,
		},
	}
}

func newContainers(app *opservicev1alpha1.OpServiceApplication) []corev1.Container {
	var containerPorts []corev1.ContainerPort
	cport := corev1.ContainerPort{}
	cport.ContainerPort = 80
	containerPorts = append(containerPorts, cport)
	return []corev1.Container{
		{
			Name:            app.Name,
			Image:           app.Spec.Image,
			Ports:           containerPorts,
			ImagePullPolicy: corev1.PullIfNotPresent,
		},
	}
}
