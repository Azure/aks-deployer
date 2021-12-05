package httphandler

type MetricEmitter interface {
	ExpectedResult()
	UnexpectedResult()
	Heartbeat()
}
