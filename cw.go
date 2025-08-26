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

func NewCharacterWorks(host string, limit float64) *CharacterWorks {
	c := &CharacterWorks{
		host:   fmt.Sprintf("http://%s", net.JoinHostPort(host, "5201")),
		client: http.DefaultClient,
	}
	c.client.Transport = throttled.NewTransport(http.DefaultTransport, rate.NewLimiter(rate.Limit(limit), 1))

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

func (c *CharacterWorks) ListGridNames() (*ListGridNamesResponse, error) {
	req := NewListGridNamesRequest()
	r, err := RequestToReader(req)
	if err != nil {
		return nil, xerrors.Errorf("failed to create reader: %w", err)
	}
	return DoPost[ListGridNamesResponse](c, r)
}

func (c *CharacterWorks) ListGridCells(grid string) (*ListGridCellsResponse, error) {
	req := NewListGridCellsRequest(grid)
	r, err := RequestToReader(req)
	if err != nil {
		return nil, xerrors.Errorf("failed to create reader: %w", err)
	}
	return DoPost[ListGridCellsResponse](c, r)
}

func (c *CharacterWorks) ActivateGridCell(grid string, cell [2]int) (*ActivateGridCellResponse, error) {
	req := NewActivateGridCellRequest(grid, cell)
	r, err := RequestToReader(req)
	if err != nil {
		return nil, xerrors.Errorf("failed to create reader: %w", err)
	}
	return DoPost[ActivateGridCellResponse](c, r)
}

// Discovery helpers (read-only)

// ListLayerNames returns a flattened list of layer names for the given channel.
// If channel is nil, the default channel is used.
func (c *CharacterWorks) ListLayerNames(channel *string) ([]string, error) {
	resp, err := c.ListLayers(nil, channel)
	if err != nil {
		return nil, err
	}
	var names []string
	var walk func(layers []Layer)
	walk = func(layers []Layer) {
		for _, l := range layers {
			names = append(names, l.Name)
			if len(l.Children) > 0 {
				walk(l.Children)
			}
		}
	}
	walk(resp.Layers)
	return names, nil
}

// ListFieldNames returns a map of field names to optional metadata for a given parent layer name.
// This uses the ListLayers request with Parent set to the provided layer, and returns child layer names.
// If channel is nil, the default channel is used.
func (c *CharacterWorks) ListFieldNames(layer string, channel *string) (map[string]string, error) {
	parent := layer
	resp, err := c.ListLayers(&parent, channel)
	if err != nil {
		return nil, err
	}
	fields := make(map[string]string, len(resp.Layers))
	for _, child := range resp.Layers {
		fields[child.Name] = ""
	}
	return fields, nil
}

// ListMotionNamesAndIDs returns parallel slices of motion names and motion IDs.
// The underlying API does not currently support channel scoping for listing; the channel parameter is accepted for future compatibility but ignored.
func (c *CharacterWorks) ListMotionNamesAndIDs(channel *string) ([]string, []string, error) {
	resp, err := c.ListMotionsWithIDs()
	if err != nil {
		return nil, nil, err
	}
	names := make([]string, 0, len(resp.Motions))
	ids := make([]string, 0, len(resp.Motions))
	for _, m := range resp.Motions {
		names = append(names, m.Name)
		ids = append(ids, m.ID)
	}
	return names, ids, nil
}
