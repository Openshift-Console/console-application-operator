package controller

import (
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	appsv1alpha1 "github.com/openshift-console/console-application-operator/api/v1alpha1"
	routev1 "github.com/openshift/api/route/v1"
)

// TODO: Implement "Progressing" status condition in the future.

// SetDegraded sets the Operator Degraded condition with the provided reason and message.
func SetDegraded(consoleApplication *appsv1alpha1.ConsoleApplication, reason, message string) {
	meta.SetStatusCondition(&consoleApplication.Status.Conditions, metav1.Condition{
		Type:               appsv1alpha1.ConditionOperatorDegraded.String(),
		Status:             metav1.ConditionTrue,
		Reason:             reason,
		LastTransitionTime: metav1.NewTime(time.Now()),
		Message:            message,
	})
}

// SetGitServiceCondition sets the GitService condition with the provided status and reason.
func SetGitServiceCondition(consoleApplication *appsv1alpha1.ConsoleApplication, status metav1.ConditionStatus, reason string) {
	meta.SetStatusCondition(&consoleApplication.Status.Conditions, metav1.Condition{
		Type:               appsv1alpha1.ConditionGitRepoReachable.String(),
		Status:             status,
		Reason:             reason,
		LastTransitionTime: metav1.NewTime(time.Now()),
		Message:            fmt.Sprintf("Git Repository Reachable: %s", string(status)),
	})
}

// SetBuildConfigCondition sets the BuildConfig condition with the provided status and reason.
func SetBuildConfigCondition(consoleApplication *appsv1alpha1.ConsoleApplication, status metav1.ConditionStatus, reason, message string) {
	meta.SetStatusCondition(&consoleApplication.Status.Conditions, metav1.Condition{
		Type:               appsv1alpha1.ConditionBuildReady.String(),
		Status:             status,
		Reason:             reason,
		LastTransitionTime: metav1.NewTime(time.Now()),
		Message:            message,
	})
}

// SetWorkloadCondition sets the Deployment condition with the provided status and reason.
func SetWorkloadCondition(consoleApplication *appsv1alpha1.ConsoleApplication, status metav1.ConditionStatus, reason, message string) {
	meta.SetStatusCondition(&consoleApplication.Status.Conditions, metav1.Condition{
		Type:               appsv1alpha1.ConditionWorkloadReady.String(),
		Status:             status,
		Reason:             reason,
		LastTransitionTime: metav1.NewTime(time.Now()),
		Message:            message,
	})
}

// SetServiceCondition sets the Service condition with the provided status and reason.
func SetServiceCondition(consoleApplication *appsv1alpha1.ConsoleApplication, status metav1.ConditionStatus, reason, message string) {
	meta.SetStatusCondition(&consoleApplication.Status.Conditions, metav1.Condition{
		Type:               appsv1alpha1.ConditionServiceReady.String(),
		Status:             status,
		Reason:             reason,
		LastTransitionTime: metav1.NewTime(time.Now()),
		Message:            message,
	})
}

// SetRouteCondition sets the Route condition with the provided status and reason.
func SetRouteCondition(consoleApplication *appsv1alpha1.ConsoleApplication, status metav1.ConditionStatus, reason, message string) {
	meta.SetStatusCondition(&consoleApplication.Status.Conditions, metav1.Condition{
		Type:               appsv1alpha1.ConditionRouteReady.String(),
		Status:             status,
		Reason:             reason,
		LastTransitionTime: metav1.NewTime(time.Now()),
		Message:            message,
	})
}

// SetRouteURL sets the Application URL in the status.
func SetRouteURL(consoleApplication *appsv1alpha1.ConsoleApplication, route *routev1.Route) {
	scheme := "http"
	if route.Spec.TLS != nil {
		scheme = "https"
	}
	consoleApplication.Status.ApplicationURL = fmt.Sprintf("%s://%s", scheme, route.Spec.Host)
}

// SetStarted sets the Operator Progressing condition to True.
func SetStarted(consoleApplication *appsv1alpha1.ConsoleApplication) {
	meta.SetStatusCondition(&consoleApplication.Status.Conditions, metav1.Condition{
		Type:               appsv1alpha1.ConditionReady.String(),
		Status:             metav1.ConditionUnknown,
		Reason:             appsv1alpha1.ReasonInit.String(),
		LastTransitionTime: metav1.NewTime(time.Now()),
		Message:            "Initializing ConsoleApplication",
	})
}

// SetFailed sets the Operator Progressing and Application Ready conditions to False with the provided reason and message.
func SetFailed(consoleApplication *appsv1alpha1.ConsoleApplication, reason, message string) {
	meta.SetStatusCondition(&consoleApplication.Status.Conditions, metav1.Condition{
		Type:               appsv1alpha1.ConditionReady.String(),
		Status:             metav1.ConditionFalse,
		Reason:             reason,
		LastTransitionTime: metav1.NewTime(time.Now()),
		Message:            message,
	})
}

// SetSucceeded sets the Operator Progressing and Application Ready conditions to True.
func SetSucceeded(consoleApplication *appsv1alpha1.ConsoleApplication) {
	meta.SetStatusCondition(&consoleApplication.Status.Conditions, metav1.Condition{
		Type:               appsv1alpha1.ConditionReady.String(),
		Status:             metav1.ConditionTrue,
		Reason:             appsv1alpha1.ReasonAllResourcesReady.String(),
		LastTransitionTime: metav1.NewTime(time.Now()),
		Message:            "All resources are successfully created and ready",
	})
}

// SetCondition sets a generic condition, overwriting existing one by type if present.
func SetCondition(consoleApplication *appsv1alpha1.ConsoleApplication, typ string, status metav1.ConditionStatus, reason string, message string) {
	meta.SetStatusCondition(&consoleApplication.Status.Conditions, metav1.Condition{
		Type:    typ,
		Status:  status,
		Reason:  reason,
		Message: message,
	})
}
