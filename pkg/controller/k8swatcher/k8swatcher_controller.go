package k8swatcher

import (
	"context"
	"fmt"
	"time"

	sbtechv1 "github.com/operator-framework/k8s-watcher/pkg/apis/sbtech/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	//"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	//"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"github.com/ashwanthkumar/slack-go-webhook"
	mycache "github.com/patrickmn/go-cache"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_k8swatcher")
var existingPods = &corev1.PodList{}
var cachePods = mycache.New(5*time.Minute, 10*time.Minute)

// Count the pods that are pending or running as available
var existingPodNames = []string{}

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new K8SWatcher Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileK8SWatcher{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	log.Info("add(mgr manager.Manager, r reconcile.Reconciler)")
	c, err := controller.New("k8swatcher-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource K8SWatcher
	//err = c.Watch(&source.Kind{Type: &sbtechv1alpha1.K8SWatcher{}}, &handler.EnqueueRequestForObject{})
	//if err != nil {
	//	return err
	//}

	//Watch for changes to primary resource K8SWatcher
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner K8SWatcher
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &appsv1.Deployment{},
		//OwnerType:    &sbtechv1alpha1.K8SWatcher{},
	})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileK8SWatcher implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileK8SWatcher{}

// ReconcileK8SWatcher reconciles a K8SWatcher object
type ReconcileK8SWatcher struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a K8SWatcher object and makes changes based on the state read
// and what is in the K8SWatcher.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileK8SWatcher) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.PodName", request.Name)
	reqLogger.Info("Reconciling K8SWatcher")

	// Fetch the K8SWatcher instance
	instance := &corev1.Pod{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	notifierConfig := &sbtechv1.K8SWatcherNotifier{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Namespace: instance.Namespace, Name: "k8swatchernotifier"}, notifierConfig)
	if err != nil && errors.IsNotFound(err) {
		log.Info("" + instance.Namespace + " namespace not watched. Please add k8swatchernotifier if needed")
		//reqLogger.Error(err, "notifierConfig not found\n")

	} else {

		log.Info("len(instance.Status.ContainerStatuses)" + string(len(instance.Status.ContainerStatuses)))
		if len(instance.Status.ContainerStatuses) > 0 && instance.Status.ContainerStatuses[0].RestartCount > 0 {
			restartCountnt, found := cachePods.Get(instance.Name)
			if found {
				log.Info("Already exist in cache. restartCountnt.(int32) > instance.Status.ContainerStatuses[0].RestartCount " + string(restartCountnt.(int32)) + ">" + string(instance.Status.ContainerStatuses[0].RestartCount))
				if restartCountnt.(int32) < instance.Status.ContainerStatuses[0].RestartCount {
					reqLogger.Info("restarts increased = " + fmt.Sprintf("%v", restartCountnt))
					cachePods.Set(instance.Name, instance.Status.ContainerStatuses[0].RestartCount, mycache.DefaultExpiration)
					reqLogger.Info("Send Alert")
					msg := "Pod: " + instance.Name + " in " + instance.Namespace + " was restarted " + fmt.Sprintf("%d", instance.Status.ContainerStatuses[0].RestartCount) + " time"
					sendSlackMessage(notifierConfig.Spec.Slack.SlackUrl, notifierConfig.Spec.Slack.Channel, msg)
				} else {
					log.Info("***** cacheRestarts " + fmt.Sprintf("%d", restartCountnt.(int32)))
					log.Info("***** newRestarts " + fmt.Sprintf("%d", instance.Status.ContainerStatuses[0].RestartCount))
				}
			} else {
				cachePods.Add(instance.Name, instance.Status.ContainerStatuses[0].RestartCount, mycache.DefaultExpiration)
				log.Info("First event added to cache")

			}
		}

	}

	return reconcile.Result{}, nil
}

func sendSlackMessage(slackUrl string, channel string, msg string) {

	webhookUrl := slackUrl
	payload := slack.Payload{
		Text:      msg,
		Username:  "k8sWatcher",
		Channel:   "#" + channel,
		IconEmoji: ":fire_engine:",
	}
	err := slack.Send(webhookUrl, "", payload)
	if len(err) > 0 {
		fmt.Printf("error: %s\n", err)
	}
}
