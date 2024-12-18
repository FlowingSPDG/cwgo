// cw is a package for controlling Character Works.
package cw

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"

	"github.com/graze/go-throttled"
	"golang.org/x/time/rate"
	"golang.org/x/xerrors"
)

type CharacterWorks struct {
	// settings
	host string

	// instance
	client *http.Client
}

func NewCharacterWorks(host string) *CharacterWorks {
	c := &CharacterWorks{
		host:   fmt.Sprintf("http://%s", net.JoinHostPort(host, "5201")),
		client: http.DefaultClient,
	}
	c.client.Transport = throttled.NewTransport(http.DefaultTransport, rate.NewLimiter(rate.Limit(30), 1))

	return c
}

func DoPost[TResp any](c *CharacterWorks, req io.Reader) (*TResp, error) {
	resp, err := c.client.Post(c.host, "application/json", req)
	if err != nil {
		return nil, xerrors.Errorf("post request failed: %w", err)
	}
	defer resp.Body.Close()

	b := new(bytes.Buffer)
	if _, err := io.Copy(b, resp.Body); err != nil {
		return nil, xerrors.Errorf("failed to read response: %w", err)
	}

	res := new(TResp)
	bs := b.Bytes()
	if err := json.Unmarshal(bs, res); err != nil {
		return nil, xerrors.Errorf("failed to decode response: %w raw: %s", err, bs)
	}
	return res, nil
}

func (c *CharacterWorks) ListMotions() (*ListMotionsResponse, error) {
	req := NewListMotionsRequest()
	r, err := RequestToReader(req)
	if err != nil {
		return nil, xerrors.Errorf("failed to create reader: %w", err)
	}
	return DoPost[ListMotionsResponse](c, r)
}

func (c *CharacterWorks) ListMotionsWithIDs() (*ListMotionsWithIDsResponse, error) {
	req := NewListMotionsWithIDsRequest()
	r, err := RequestToReader(req)
	if err != nil {
		return nil, xerrors.Errorf("failed to create reader: %w", err)
	}
	return DoPost[ListMotionsWithIDsResponse](c, r)
}

func (c *CharacterWorks) ListLayers(parent *string, channel *string) (*ListLayersResponse, error) {
	req := NewListLayersRequest(parent, channel)
	r, err := RequestToReader(req)
	if err != nil {
		return nil, xerrors.Errorf("failed to create reader: %w", err)
	}
	return DoPost[ListLayersResponse](c, r)
}

func (c *CharacterWorks) PlayMotions(motions []string, motionIDs []string, channel *string) (*PlayMotionsResponse, error) {
	req := NewPlayMotionsRequest(motions, motionIDs, channel)
	r, err := RequestToReader(req)
	if err != nil {
		return nil, xerrors.Errorf("failed to create reader: %w", err)
	}
	return DoPost[PlayMotionsResponse](c, r)
}

func (c *CharacterWorks) StopMotions(motions []string, motionIDs []string, channel *string) (*StopMotionsResponse, error) {
	req := NewStopMotionsRequest(motions, motionIDs, channel)
	r, err := RequestToReader(req)
	if err != nil {
		return nil, xerrors.Errorf("failed to create reader: %w", err)
	}
	return DoPost[StopMotionsResponse](c, r)
}

func (c *CharacterWorks) EjectMotions(motions []string, motionIDs []string, channel *string) (*EjectMotionsResponse, error) {
	req := NewEjectMotionsRequest(motions, motionIDs, channel)
	r, err := RequestToReader(req)
	if err != nil {
		return nil, xerrors.Errorf("failed to create reader: %w", err)
	}
	return DoPost[EjectMotionsResponse](c, r)
}

func (c *CharacterWorks) FinishMotions(motions []string, motionIDs []string, channel *string) (*FinishMotionsResponse, error) {
	req := NewFinishMotionsRequest(motions, motionIDs, channel)
	r, err := RequestToReader(req)
	if err != nil {
		return nil, xerrors.Errorf("failed to create reader: %w", err)
	}
	return DoPost[FinishMotionsResponse](c, r)
}

func (c *CharacterWorks) PauseMotions(motions []string, motionIDs []string, channel *string) (*PauseMotionsResponse, error) {
	req := NewPauseMotionsRequest(motions, motionIDs, channel)
	r, err := RequestToReader(req)
	if err != nil {
		return nil, xerrors.Errorf("failed to create reader: %w", err)
	}
	return DoPost[PauseMotionsResponse](c, r)
}

func (c *CharacterWorks) ResumeMotions(motions []string, motionIDs []string, channel *string) (*ResumeMotionsResponse, error) {
	req := NewResumeMotionsRequest(motions, motionIDs, channel)
	r, err := RequestToReader(req)
	if err != nil {
		return nil, xerrors.Errorf("failed to create reader: %w", err)
	}
	return DoPost[ResumeMotionsResponse](c, r)
}

func (c *CharacterWorks) RestartMotions(motions []string, motionIDs []string, channel *string) (*RestartMotionsResponse, error) {
	req := NewRestartMotionsRequest(motions, motionIDs, channel)
	r, err := RequestToReader(req)
	if err != nil {
		return nil, xerrors.Errorf("failed to create reader: %w", err)
	}
	return DoPost[RestartMotionsResponse](c, r)
}

func (c *CharacterWorks) FinishAndRestartMotions(motions []string, motionIDs []string, channel *string) (*FinishAndRestartMotionsResponse, error) {
	req := NewFinishAndRestartMotionsRequest(motions, motionIDs, channel)
	r, err := RequestToReader(req)
	if err != nil {
		return nil, xerrors.Errorf("failed to create reader: %w", err)
	}
	return DoPost[FinishAndRestartMotionsResponse](c, r)
}

func (c *CharacterWorks) SetText(layer string, layerID []string, value string, channel *string) (*SetTextResponse, error) {
	req := NewSetTextRequest(layer, layerID, value, channel)
	r, err := RequestToReader(req)
	if err != nil {
		return nil, xerrors.Errorf("failed to create reader: %w", err)
	}
	return DoPost[SetTextResponse](c, r)
}
