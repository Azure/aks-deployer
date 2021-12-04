package consts

// AsyncOperationStatus represents the current state of async operation.
type AsyncOperationStatus string

const (
	// InProgress indicates an ongoing operation
	InProgress AsyncOperationStatus = "InProgress"
	// Succeeded means operation completed successfully
	Succeeded AsyncOperationStatus = "Succeeded"
	// Failed indicates that operation has failed
	Failed AsyncOperationStatus = "Failed"
	// Canceled indicates that operation was Canceled
	Canceled AsyncOperationStatus = "Canceled"
	// FailedNoAutoReconcile indicates that operation is failed but no auto reconciliation should be initialized.
	// This status should only be used in AsyncOperationTracking but never provisioningState.
	FailedNoAutoReconcile AsyncOperationStatus = "FailedNoAutoReconcile"
)
