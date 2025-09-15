// cw is a package for controlling Character Works.
package cw

import (
	"bytes"
	"context"
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

func DoPost[TResp any](ctx context.Context, c *CharacterWorks, req io.Reader) (*TResp, error) {
	httpReq, err := http.NewRequestWithContext(ctx, "POST", c.host, req)
	if err != nil {
		return nil, xerrors.Errorf("failed to create request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(httpReq)
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

func (c *CharacterWorks) ListMotions(ctx context.Context) (*ListMotionsResponse, error) {
	req := NewListMotionsRequest()
	r, err := RequestToReader(req)
	if err != nil {
		return nil, xerrors.Errorf("failed to create reader: %w", err)
	}
	return DoPost[ListMotionsResponse](ctx, c, r)
}

func (c *CharacterWorks) ListMotionsWithIDs(ctx context.Context) (*ListMotionsWithIDsResponse, error) {
	req := NewListMotionsWithIDsRequest()
	r, err := RequestToReader(req)
	if err != nil {
		return nil, xerrors.Errorf("failed to create reader: %w", err)
	}
	return DoPost[ListMotionsWithIDsResponse](ctx, c, r)
}

func (c *CharacterWorks) ListLayers(ctx context.Context, parent *string, channel *string) (*ListLayersResponse, error) {
	req := NewListLayersRequest(parent, channel)
	r, err := RequestToReader(req)
	if err != nil {
		return nil, xerrors.Errorf("failed to create reader: %w", err)
	}
	return DoPost[ListLayersResponse](ctx, c, r)
}

func (c *CharacterWorks) PlayMotions(ctx context.Context, motions []string, motionIDs []string, channel *string) (*PlayMotionsResponse, error) {
	req := NewPlayMotionsRequest(motions, motionIDs, channel)
	r, err := RequestToReader(req)
	if err != nil {
		return nil, xerrors.Errorf("failed to create reader: %w", err)
	}
	return DoPost[PlayMotionsResponse](ctx, c, r)
}

func (c *CharacterWorks) StopMotions(ctx context.Context, motions []string, motionIDs []string, channel *string) (*StopMotionsResponse, error) {
	req := NewStopMotionsRequest(motions, motionIDs, channel)
	r, err := RequestToReader(req)
	if err != nil {
		return nil, xerrors.Errorf("failed to create reader: %w", err)
	}
	return DoPost[StopMotionsResponse](ctx, c, r)
}

func (c *CharacterWorks) EjectMotions(ctx context.Context, motions []string, motionIDs []string, channel *string) (*EjectMotionsResponse, error) {
	req := NewEjectMotionsRequest(motions, motionIDs, channel)
	r, err := RequestToReader(req)
	if err != nil {
		return nil, xerrors.Errorf("failed to create reader: %w", err)
	}
	return DoPost[EjectMotionsResponse](ctx, c, r)
}

func (c *CharacterWorks) FinishMotions(ctx context.Context, motions []string, motionIDs []string, channel *string) (*FinishMotionsResponse, error) {
	req := NewFinishMotionsRequest(motions, motionIDs, channel)
	r, err := RequestToReader(req)
	if err != nil {
		return nil, xerrors.Errorf("failed to create reader: %w", err)
	}
	return DoPost[FinishMotionsResponse](ctx, c, r)
}

func (c *CharacterWorks) PauseMotions(ctx context.Context, motions []string, motionIDs []string, channel *string) (*PauseMotionsResponse, error) {
	req := NewPauseMotionsRequest(motions, motionIDs, channel)
	r, err := RequestToReader(req)
	if err != nil {
		return nil, xerrors.Errorf("failed to create reader: %w", err)
	}
	return DoPost[PauseMotionsResponse](ctx, c, r)
}

func (c *CharacterWorks) ResumeMotions(ctx context.Context, motions []string, motionIDs []string, channel *string) (*ResumeMotionsResponse, error) {
	req := NewResumeMotionsRequest(motions, motionIDs, channel)
	r, err := RequestToReader(req)
	if err != nil {
		return nil, xerrors.Errorf("failed to create reader: %w", err)
	}
	return DoPost[ResumeMotionsResponse](ctx, c, r)
}

func (c *CharacterWorks) RestartMotions(ctx context.Context, motions []string, motionIDs []string, channel *string) (*RestartMotionsResponse, error) {
	req := NewRestartMotionsRequest(motions, motionIDs, channel)
	r, err := RequestToReader(req)
	if err != nil {
		return nil, xerrors.Errorf("failed to create reader: %w", err)
	}
	return DoPost[RestartMotionsResponse](ctx, c, r)
}

func (c *CharacterWorks) FinishAndRestartMotions(ctx context.Context, motions []string, motionIDs []string, channel *string) (*FinishAndRestartMotionsResponse, error) {
	req := NewFinishAndRestartMotionsRequest(motions, motionIDs, channel)
	r, err := RequestToReader(req)
	if err != nil {
		return nil, xerrors.Errorf("failed to create reader: %w", err)
	}
	return DoPost[FinishAndRestartMotionsResponse](ctx, c, r)
}

func (c *CharacterWorks) SetText(ctx context.Context, layer string, layerID []string, value string, channel *string) (*SetTextResponse, error) {
	req := NewSetTextRequest(layer, layerID, value, channel)
	r, err := RequestToReader(req)
	if err != nil {
		return nil, xerrors.Errorf("failed to create reader: %w", err)
	}
	return DoPost[SetTextResponse](ctx, c, r)
}

func (c *CharacterWorks) ListGridNames(ctx context.Context) (*ListGridNamesResponse, error) {
	req := NewListGridNamesRequest()
	r, err := RequestToReader(req)
	if err != nil {
		return nil, xerrors.Errorf("failed to create reader: %w", err)
	}
	return DoPost[ListGridNamesResponse](ctx, c, r)
}

func (c *CharacterWorks) ListGridCells(ctx context.Context, grid string) (*ListGridCellsResponse, error) {
	req := NewListGridCellsRequest(grid)
	r, err := RequestToReader(req)
	if err != nil {
		return nil, xerrors.Errorf("failed to create reader: %w", err)
	}
	return DoPost[ListGridCellsResponse](ctx, c, r)
}

func (c *CharacterWorks) ActivateGridCell(ctx context.Context, grid string, cell [2]int) (*ActivateGridCellResponse, error) {
	req := NewActivateGridCellRequest(grid, cell)
	r, err := RequestToReader(req)
	if err != nil {
		return nil, xerrors.Errorf("failed to create reader: %w", err)
	}
	return DoPost[ActivateGridCellResponse](ctx, c, r)
}

// Discovery helpers (read-only)

// ListLayerNames returns a flattened list of layer names for the given channel.
// If channel is nil, the default channel is used.
func (c *CharacterWorks) ListLayerNames(ctx context.Context, channel *string) ([]string, error) {
	resp, err := c.ListLayers(ctx, nil, channel)
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
func (c *CharacterWorks) ListFieldNames(ctx context.Context, layer string, channel *string) (map[string]string, error) {
	parent := layer
	resp, err := c.ListLayers(ctx, &parent, channel)
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
func (c *CharacterWorks) ListMotionNamesAndIDs(ctx context.Context, channel *string) ([]string, []string, error) {
	resp, err := c.ListMotionsWithIDs(ctx)
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
