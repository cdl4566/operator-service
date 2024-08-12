/*
Copyright 2024.

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

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	opservicev1alpha1 "github.com/cdl4566/operator-service/api/v1alpha1"
)

// OpServiceApplicationReconciler reconciles a OpServiceApplication object
type OpServiceApplicationReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=op.service.cdl4566.com,resources=opserviceapplications,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=op.service.cdl4566.com,resources=opserviceapplications/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=op.service.cdl4566.com,resources=opserviceapplications/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the OpServiceApplication object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.12.1/pkg/reconcile
func (r *OpServiceApplicationReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	// TODO(user): your logic here
	var opServiceApplication opservicev1alpha1.OpServiceApplication
	if err := r.Get(ctx, req.NamespacedName, &opServiceApplication); err != nil {
		// 如果 algo application 是被删除的，应该忽略掉，交给kubernetes的垃圾回收处理
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	logger.Info("get opServiceApplication", "opServiceApplication", opServiceApplication)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *OpServiceApplicationReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&opservicev1alpha1.OpServiceApplication{}).
		Complete(r)
}
