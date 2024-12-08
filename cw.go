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

type CharacterWorks interface {
	ListMotions() (*ListLayersResponse, error)
	ListMotionsWithIDs() (*ListMotionsWithIDsResponse, error)
	ListLayers(parent *string, channel *string) (*ListLayersResponse, error)
	PlayMotions(motions []string, motionIDs []string, channel *string) (*PlayMotionsResponse, error)
	StopMotions(motions []string, motionIDs []string, channel *string) (*StopMotionsResponse, error)
	EjectMotions(motions []string, motionIDs []string, channel *string) (*EjectMotionsResponse, error)
	FinishMotions(motions []string, motionIDs []string, channel *string) (*FinishMotionsResponse, error)
	PauseMotions(motions []string, motionIDs []string, channel *string) (*PauseMotionsResponse, error)
	ResumeMotions(motions []string, motionIDs []string, channel *string) (*ResumeMotionsResponse, error)
	RestartMotions(motions []string, motionIDs []string, channel *string) (*RestartMotionsResponse, error)
	FinishAndRestartMotions(motions []string, motionIDs []string, channel *string) (*FinishAndRestartMotionsResponse, error)
	// SetText
	// ListGridNames
	// ListGridCells
	// ActivateGridCell
	// RunDataSourceQuery
	// SetDataSourceQueryParameter
	// SelectDataSourceRows
}

type characterWorks struct {
	// settings
	host string

	// instance
	client *http.Client
}

func NewCharacterWorks(host string) CharacterWorks {
	return &characterWorks{
		host:   fmt.Sprintf("http://%s", net.JoinHostPort(host, "5201")),
		client: http.DefaultClient, // for now
	}
}

func doPost[TResp any](c *characterWorks, req io.Reader) (*TResp, error) {
	resp, err := c.client.Post(c.host, "application/json", req)
	if err != nil {
		return nil, xerrors.Errorf("failed to list motions with id: %w", err)
	}
	defer resp.Body.Close()

	res := new(TResp)
	if err := json.NewDecoder(resp.Body).Decode(res); err != nil {
		return nil, xerrors.Errorf("failed to decode response: %w", err)
	}
	return res, nil
}

func (c *characterWorks) ListMotions() (*ListLayersResponse, error) {
	req := newListMotionsRequest()
	return doPost[ListLayersResponse](c, req)
}

func (c *characterWorks) ListMotionsWithIDs() (*ListMotionsWithIDsResponse, error) {
	req := newListMotionsWithIDsRequest()
	return doPost[ListMotionsWithIDsResponse](c, req)
}

func (c *characterWorks) ListLayers(parent *string, channel *string) (*ListLayersResponse, error) {
	req := newListLayersRequest(parent, channel)
	return doPost[ListLayersResponse](c, req)
}

func (c *characterWorks) PlayMotions(motions []string, motionIDs []string, channel *string) (*PlayMotionsResponse, error) {
	req := newPlayMotionsRequest(motions, motionIDs, channel)
	return doPost[PlayMotionsResponse](c, req)
}

func (c *characterWorks) StopMotions(motions []string, motionIDs []string, channel *string) (*StopMotionsResponse, error) {
	req := newStopMotionsRequest(motions, motionIDs, channel)
	return doPost[StopMotionsResponse](c, req)
}

func (c *characterWorks) EjectMotions(motions []string, motionIDs []string, channel *string) (*EjectMotionsResponse, error) {
	req := newEjectMotionsRequest(motions, motionIDs, channel)
	return doPost[EjectMotionsResponse](c, req)
}

func (c *characterWorks) FinishMotions(motions []string, motionIDs []string, channel *string) (*FinishMotionsResponse, error) {
	req := newFinishMotionsRequest(motions, motionIDs, channel)
	return doPost[FinishMotionsResponse](c, req)
}

func (c *characterWorks) PauseMotions(motions []string, motionIDs []string, channel *string) (*PauseMotionsResponse, error) {
	req := newPauseMotionsRequest(motions, motionIDs, channel)
	return doPost[PauseMotionsResponse](c, req)
}

func (c *characterWorks) ResumeMotions(motions []string, motionIDs []string, channel *string) (*ResumeMotionsResponse, error) {
	req := newResumeMotionsRequest(motions, motionIDs, channel)
	return doPost[ResumeMotionsResponse](c, req)
}

func (c *characterWorks) RestartMotions(motions []string, motionIDs []string, channel *string) (*RestartMotionsResponse, error) {
	req := newRestartMotionsRequest(motions, motionIDs, channel)
	return doPost[RestartMotionsResponse](c, req)
}

func (c *characterWorks) FinishAndRestartMotions(motions []string, motionIDs []string, channel *string) (*FinishAndRestartMotionsResponse, error) {
	req := newFinishAndRestartMotionsRequest(motions, motionIDs, channel)
	return doPost[FinishAndRestartMotionsResponse](c, req)
}
