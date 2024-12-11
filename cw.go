// cw is a package for controlling Character Works.
package cw

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"

	"golang.org/x/xerrors"
)

type CharacterWorks struct {
	// settings
	host string

	// instance
	client *http.Client
}

func NewCharacterWorks(host string) *CharacterWorks {
	return &CharacterWorks{
		host:   fmt.Sprintf("http://%s", net.JoinHostPort(host, "5201")),
		client: http.DefaultClient,
	}
}

func doPost[TResp any](c *CharacterWorks, req io.Reader) (*TResp, error) {
	resp, err := c.client.Post(c.host, "application/json", req)
	if err != nil {
		return nil, xerrors.Errorf("post request failed: %w", err)
	}
	defer resp.Body.Close()

	res := new(TResp)
	if err := json.NewDecoder(resp.Body).Decode(res); err != nil {
		return nil, xerrors.Errorf("failed to decode response: %w", err)
	}
	return res, nil
}

func (c *CharacterWorks) ListMotions() (*ListMotionsResponse, error) {
	req := NewListMotionsRequest()
	return doPost[ListMotionsResponse](c, req)
}

func (c *CharacterWorks) ListMotionsWithIDs() (*ListMotionsWithIDsResponse, error) {
	req := NewListMotionsWithIDsRequest()
	return doPost[ListMotionsWithIDsResponse](c, req)
}

func (c *CharacterWorks) ListLayers(parent *string, channel *string) (*ListLayersResponse, error) {
	req := NewListLayersRequest(parent, channel)
	return doPost[ListLayersResponse](c, req)
}

func (c *CharacterWorks) PlayMotions(motions []string, motionIDs []string, channel *string) (*PlayMotionsResponse, error) {
	req := NewPlayMotionsRequest(motions, motionIDs, channel)
	return doPost[PlayMotionsResponse](c, req)
}

func (c *CharacterWorks) StopMotions(motions []string, motionIDs []string, channel *string) (*StopMotionsResponse, error) {
	req := NewStopMotionsRequest(motions, motionIDs, channel)
	return doPost[StopMotionsResponse](c, req)
}

func (c *CharacterWorks) EjectMotions(motions []string, motionIDs []string, channel *string) (*EjectMotionsResponse, error) {
	req := NewEjectMotionsRequest(motions, motionIDs, channel)
	return doPost[EjectMotionsResponse](c, req)
}

func (c *CharacterWorks) FinishMotions(motions []string, motionIDs []string, channel *string) (*FinishMotionsResponse, error) {
	req := NewFinishMotionsRequest(motions, motionIDs, channel)
	return doPost[FinishMotionsResponse](c, req)
}

func (c *CharacterWorks) PauseMotions(motions []string, motionIDs []string, channel *string) (*PauseMotionsResponse, error) {
	req := NewPauseMotionsRequest(motions, motionIDs, channel)
	return doPost[PauseMotionsResponse](c, req)
}

func (c *CharacterWorks) ResumeMotions(motions []string, motionIDs []string, channel *string) (*ResumeMotionsResponse, error) {
	req := NewResumeMotionsRequest(motions, motionIDs, channel)
	return doPost[ResumeMotionsResponse](c, req)
}

func (c *CharacterWorks) RestartMotions(motions []string, motionIDs []string, channel *string) (*RestartMotionsResponse, error) {
	req := NewRestartMotionsRequest(motions, motionIDs, channel)
	return doPost[RestartMotionsResponse](c, req)
}

func (c *CharacterWorks) FinishAndRestartMotions(motions []string, motionIDs []string, channel *string) (*FinishAndRestartMotionsResponse, error) {
	req := NewFinishAndRestartMotionsRequest(motions, motionIDs, channel)
	return doPost[FinishAndRestartMotionsResponse](c, req)
}

func (c *CharacterWorks) SetText(layer string, layerID []string, value string, channel *string) (*SetTextResponse, error) {
	req := NewSetTextRequest(layer, layerID, value, channel)
	return doPost[SetTextResponse](c, req)
}
