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

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/sirupsen/logrus"

	deployerv1 "github.com/Azure/aks-deployer/pkg/api/v1"
	"github.com/Azure/aks-deployer/pkg/configmaps"
)

// ConfigMapReconciler reconciles a ConfigMap object
type ConfigMapReconciler struct {
	client.Client
	Logger    *logrus.Entry
	Scheme    *runtime.Scheme
	Namespace string
}

// +kubebuilder:rbac:groups=core,resources=configmaps,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=configmaps/status,verbs=get;update;patch

// Reconcile reconciles ConfigMaps
func (r *ConfigMapReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Logger.WithField("configmap", req.NamespacedName)

	if req.NamespacedName.String() != configmaps.GetNamespacedClusterConfigMapName(r.Namespace) {
		return ctrl.Result{}, nil
	}

	log.Info("Start reconciling ConfigMaps")
	var err error

	// 1. Retrieve the cluster configuration
	var cm corev1.ConfigMap
	if err = r.Get(ctx, req.NamespacedName, &cm); err != nil {
		if errors.IsNotFound(err) {
			log.Infof("config map no longer exists, %s", err.Error())
			return ctrl.Result{}, nil
		}
		log.Errorf("unable to get cluster config map, %s", err.Error())
		return ctrl.Result{}, err
	}

	log.Info("Parse the configuration data")
	// 2. Parse the configuration data
	objs, err := configmaps.ParseConfig(log, r.Scheme, cm.Data["config"])
	if err != nil {
		log.Errorf("unable to parse cluster configuration, %s", err.Error())
		return ctrl.Result{}, err
	}

	log.Info("Create or update AksApp component CRDs")
	// 3. Create or update AksApp component CRDs
	for _, obj := range objs {
		app := obj.(*deployerv1.AksApp)
		var tmpApp deployerv1.AksApp
		nn := types.NamespacedName{
			Namespace: app.Namespace,
			Name:      app.Name,
		}

		// 3.0 Check if AksApp namespace exists or not
		var tmpNs corev1.Namespace
		if err = r.Get(ctx, types.NamespacedName{Name: app.Namespace}, &tmpNs); err != nil {
			if errors.IsNotFound(err) {
				ns := &corev1.Namespace{
					ObjectMeta: metav1.ObjectMeta{
						Name: app.Namespace,
					},
				}
				if err = r.Create(ctx, ns); err != nil {
					log.Errorf("unable to create aksapp namespace, %s", err.Error())
					continue
				}
			} else {
				log.Errorf("unable to get aksapp namespace, %s", err.Error())
				continue
			}
		}

		// 3.1 Check if CRD exists or not
		if err = r.Get(ctx, nn, &tmpApp); err != nil {
			// 3.1.1 Create CRD if not found
			if errors.IsNotFound(err) {
				if err = r.Create(ctx, app); err != nil {
					log.Errorf("unable to create aksapp component, %s", err.Error())
					continue
				}
				log.Infof("Created aksapp component %s/%s", app.Namespace, app.Name)
			} else {
				log.Errorf("unable to get aksapp component, %s", err.Error())
				continue
			}
		} else {
			// 3.2 Update CRD if version does not match
			if !reflect.DeepEqual(tmpApp.Spec, app.Spec) {
				app.ObjectMeta.ResourceVersion = tmpApp.ObjectMeta.ResourceVersion // metadata.resourceVersion must be specified for an update
				if err = r.Update(ctx, app); err != nil {
					log.Errorf("unable to update aksapp component, %s", err.Error())
					continue
				}
				log.Infof("Updated aksapp component %s/%s", app.Namespace, app.Name)
			}
		}
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up ConfigMap manager
func (r *ConfigMapReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&corev1.ConfigMap{}).
		Owns(&deployerv1.AksApp{}).
		Complete(r)
}
