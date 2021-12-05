package controllers

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("Test DeploymentComplete", func() {
	var (
		deployment appsv1.Deployment
	)

	BeforeEach(func() {
		replicas := int32(3)
		deployment = appsv1.Deployment{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-deployment",
				Namespace: "test-namespace",
			},
			Spec: appsv1.DeploymentSpec{
				Template: corev1.PodTemplateSpec{
					ObjectMeta: metav1.ObjectMeta{
						Name: "test-container",
					},
				},
				Replicas: &replicas,
			},
			TypeMeta: metav1.TypeMeta{
				Kind: "Deployment",
			},
		}
	})

	It("Test deployment completed", func() {
		deployment.Status.UpdatedReplicas = 3
		deployment.Status.Replicas = 3
		deployment.Status.AvailableReplicas = 3

		complete := DeploymentComplete(&deployment)
		Expect(complete).To(BeTrue())
	})

	It("Test deployment not all updated", func() {
		deployment.Status.UpdatedReplicas = 2
		deployment.Status.Replicas = 3
		deployment.Status.AvailableReplicas = 3

		complete := DeploymentComplete(&deployment)
		Expect(complete).To(BeFalse())
	})

	It("Test deployment not all updated", func() {
		deployment.Status.UpdatedReplicas = 3
		deployment.Status.Replicas = 3
		deployment.Status.AvailableReplicas = 2

		complete := DeploymentComplete(&deployment)
		Expect(complete).To(BeFalse())
	})
})
