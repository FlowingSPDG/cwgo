package cw

import "io"

type RequestsBuilder struct {
	requests  []any
	responses []any
}

func AddRequest[TResp any](b *RequestsBuilder, req any) {
	if b.requests == nil {
		b.requests = make([]any, 0)
	}
	if b.responses == nil {
		b.responses = make([]any, 0)
	}
	b.requests = append(b.requests, req)
	b.responses = append(b.responses, new(TResp))
}

func (b *RequestsBuilder) Build() io.Reader {
	return toReader(b.requests)
}

func (b *RequestsBuilder) Do(cw *CharacterWorks) ([]map[string]any, error) {
	// TODO: reflect response type
	req := b.Build()
	resp, err := doPost[[]map[string]any](cw, req)
	if err != nil {
		return nil, err
	}
	return *resp, nil
}

func (b *RequestsBuilder) AddListMotionsRequest() {
	AddRequest[ListLayersResponse](b, NewListMotionsRequest())
}

func (b *RequestsBuilder) AddListMotionsWithIDsRequest() {
	AddRequest[ListMotionsWithIDsResponse](b, NewListMotionsWithIDsRequest())
}

func (b *RequestsBuilder) AddListLayersRequest(parent *string, channel *string) {
	AddRequest[ListLayersResponse](b, NewListLayersRequest(parent, channel))
}

func (b *RequestsBuilder) AddPlayMotionsRequest(motions []string, motionIDs []string, channel *string) {
	AddRequest[PlayMotionsResponse](b, NewPlayMotionsRequest(motions, motionIDs, channel))
}

func (b *RequestsBuilder) AddStopMotionsRequest(motions []string, motionIDs []string, channel *string) {
	AddRequest[StopMotionsResponse](b, NewStopMotionsRequest(motions, motionIDs, channel))
}

func (b *RequestsBuilder) AddEjectMotionsRequest(motions []string, motionIDs []string, channel *string) {
	AddRequest[EjectMotionsResponse](b, NewEjectMotionsRequest(motions, motionIDs, channel))
}

func (b *RequestsBuilder) AddFinishMotionsRequest(motions []string, motionIDs []string, channel *string) {
	AddRequest[FinishMotionsResponse](b, NewFinishMotionsRequest(motions, motionIDs, channel))
}

func (b *RequestsBuilder) AddPauseMotionsRequest(motions []string, motionIDs []string, channel *string) {
	AddRequest[PauseMotionsResponse](b, NewPauseMotionsRequest(motions, motionIDs, channel))
}

func (b *RequestsBuilder) AddResumeMotionsRequest(motions []string, motionIDs []string, channel *string) {
	AddRequest[ResumeMotionsResponse](b, NewResumeMotionsRequest(motions, motionIDs, channel))
}

func (b *RequestsBuilder) AddRestartMotionsRequest(motions []string, motionIDs []string, channel *string) {
	AddRequest[RestartMotionsResponse](b, NewRestartMotionsRequest(motions, motionIDs, channel))
}

func (b *RequestsBuilder) AddFinishAndRestartMotionsRequest(motions []string, motionIDs []string, channel *string) {
	AddRequest[FinishAndRestartMotionsResponse](b, NewFinishAndRestartMotionsRequest(motions, motionIDs, channel))
}
