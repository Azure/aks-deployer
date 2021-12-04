package log

import (
	"context"
	"runtime"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.opencensus.io/trace"
)

type testTraceExporter struct {
	FinalSpanData *trace.SpanData
}

func (e *testTraceExporter) ExportSpan(sd *trace.SpanData) {
	e.FinalSpanData = sd
}

var _ = Describe("Testing Trace Package", func() {
	var (
		testExporter *testTraceExporter
	)

	BeforeEach(func() {
		trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
		testExporter = &testTraceExporter{}
		trace.RegisterExporter(testExporter)
	})

	It("Test StartSpan", func() {
		ctx := context.Background()

		_, file, line, _ := runtime.Caller(0)
		ctx, span := StartSpan(ctx, "testspan", AKSTeamUnknown)
		span.End()

		spanData := testExporter.FinalSpanData
		Expect(spanData).ToNot(BeNil())
		Expect(spanData.Annotations).To(HaveLen(1))
		stackAttributes := spanData.Annotations[0].Attributes
		Expect(stackAttributes[fileNameFieldName]).To(Equal(file))
		Expect(stackAttributes[lineNumberFieldName]).To(Equal(int64(line + 1)))
	})

	It("Test StartSpanWithTags", func() {
		ctx := context.Background()

		_, file, line, _ := runtime.Caller(0)
		ctx, span := StartSpanWithTags(ctx, "testspan", AKSTeamUnknown, map[string]string{
			"testField":   "testValue",
			"operationID": "test-op-id",
		})

		span.End()

		spanData := testExporter.FinalSpanData
		Expect(spanData).ToNot(BeNil())
		Expect(spanData.Attributes["testField"]).To(Equal("testValue"))
		Expect(spanData.Attributes["operationID"]).To(Equal("test-op-id"))

		Expect(spanData).ToNot(BeNil())
		Expect(spanData.Annotations).To(HaveLen(1))
		stackAttributes := spanData.Annotations[0].Attributes
		Expect(stackAttributes[fileNameFieldName]).To(Equal(file))
		Expect(stackAttributes[lineNumberFieldName]).To(Equal(int64(line + 1)))
	})

	It("StartSpan should set team to unknown when same as parent is specified but there's no team found in parent", func() {
		ctx := context.Background()

		ctx, span := StartSpan(ctx, "testspan", AKSTeamSameAsParent)
		span.End()

		spanData := testExporter.FinalSpanData
		Expect(spanData).ToNot(BeNil())

		Expect(spanData.Attributes[aksTeamAttributeKey]).To(Equal(string(AKSTeamUnknown)))

		teamFromCtx := ctx.Value(teamContextKey{})
		Expect(teamFromCtx).NotTo(BeNil())
		Expect(teamFromCtx.(AKSTeam)).To(Equal(AKSTeamUnknown))
	})

	It("StartSpan should set team to the same value as parent when that's specified", func() {
		ctx, parentSpan := StartSpan(context.Background(), "parent span", AKSTeamRP)
		ctx, childSpan := StartSpan(ctx, "child span", AKSTeamSameAsParent)
		parentSpan.End()
		childSpan.End()

		spanData := testExporter.FinalSpanData
		Expect(spanData).ToNot(BeNil())

		Expect(spanData.Attributes[aksTeamAttributeKey]).To(Equal(string(AKSTeamRP)))

		teamFromCtx := ctx.Value(teamContextKey{})
		Expect(teamFromCtx).NotTo(BeNil())
		Expect(teamFromCtx.(AKSTeam)).To(Equal(AKSTeamRP))
	})

	It("StartSpan should set team to the given value", func() {
		ctx, parentSpan := StartSpan(context.Background(), "parent span", AKSTeamUnknown)
		ctx, childSpan := StartSpan(ctx, "child span", AKSTeamRP)
		parentSpan.End()
		childSpan.End()

		spanData := testExporter.FinalSpanData
		Expect(spanData).ToNot(BeNil())

		Expect(spanData.Attributes[aksTeamAttributeKey]).To(Equal(string(AKSTeamRP)))

		teamFromCtx := ctx.Value(teamContextKey{})
		Expect(teamFromCtx).NotTo(BeNil())
		Expect(teamFromCtx.(AKSTeam)).To(Equal(AKSTeamRP))
	})

	It("NewContextWithTeam and GetTeamFromContext should save and retrieve team value", func() {
		ctx := context.Background()
		Expect(GetTeamFromContext(ctx)).To(Equal(AKSTeamUnknown))

		ctx = NewContextWithTeam(ctx, AKSTeamRP)
		Expect(GetTeamFromContext(ctx)).To(Equal(AKSTeamRP))
	})
})
