package consts

type ScaleDownDisabledReason string

const (
	ScaleDownDisabledAnnotationKey       = "cluster-autoscaler.kubernetes.io/scale-down-disabled"
	ScaleDownDisabledReasonAnnotationKey = AKSPrefix + "azure-cluster-autoscaler-scale-down-disabled-reason"

	ScaleDownDisabledReasonUpgrade ScaleDownDisabledReason = "upgrade"
)
