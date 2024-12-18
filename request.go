package cw

import (
	"bytes"
	"encoding/json"
	"io"

	"golang.org/x/xerrors"
)

type commonRequest struct {
	Action string `json:"action"`
}

type ListMotionsRequest struct {
	commonRequest
}

func RequestToReader(v any) (io.Reader, error) {
	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(v); err != nil {
		return nil, xerrors.Errorf("unexpected error encoding request")
	}
	return buf, nil
}

func RequestsToReader(vs []any) (io.Reader, error) {
	buf := &bytes.Buffer{}
	encoder := json.NewEncoder(buf)
	for _, v := range vs {
		if err := encoder.Encode(v); err != nil {
			return nil, xerrors.Errorf("unexpected error encoding request")
		}
	}
	return buf, nil
}

func NewListMotionsRequest() *ListMotionsRequest {
	return &ListMotionsRequest{
		commonRequest: commonRequest{
			Action: "list_motions",
		},
	}
}

type ListMotionsWithIDsRequest struct {
	commonRequest
}

func NewListMotionsWithIDsRequest() *ListMotionsWithIDsRequest {
	return &ListMotionsWithIDsRequest{
		commonRequest: commonRequest{
			Action: "list_motions_with_ids",
		},
	}
}

type ListLayersRequest struct {
	commonRequest
	Parent  *string `json:"parent,omitempty"`
	Channel *string `json:"channel,omitempty"`
}

func NewListLayersRequest(parent *string, channel *string) *ListLayersRequest {
	return &ListLayersRequest{
		commonRequest: commonRequest{
			Action: "list_layers",
		},
		Parent:  parent,
		Channel: channel,
	}
}

type PlayMotionsRequest struct {
	commonRequest
	Motions   []string `json:"motions"`
	MotionIDs []string `json:"motion_ids"`
	Channel   *string  `json:"channel,omitempty"`
}

func NewPlayMotionsRequest(motions []string, motionIDs []string, channel *string) *PlayMotionsRequest {
	return &PlayMotionsRequest{
		commonRequest: commonRequest{
			Action: "play_motions",
		},
		Motions:   motions,
		MotionIDs: motionIDs,
		Channel:   channel,
	}
}

type StopMotionsRequest struct {
	commonRequest
	Motions   []string `json:"motions"`
	MotionIDs []string `json:"motion_ids"`
	Channel   *string  `json:"channel,omitempty"`
}

func NewStopMotionsRequest(motions []string, motionIDs []string, channel *string) *StopMotionsRequest {
	return &StopMotionsRequest{
		commonRequest: commonRequest{
			Action: "stop_motions",
		},
		Motions:   motions,
		MotionIDs: motionIDs,
		Channel:   channel,
	}
}

type EjectMotionsRequest struct {
	commonRequest
	Motions   []string `json:"motions"`
	MotionIDs []string `json:"motion_ids"`
	Channel   *string  `json:"channel,omitempty"`
}

func NewEjectMotionsRequest(motions []string, motionIDs []string, channel *string) *EjectMotionsRequest {
	return &EjectMotionsRequest{
		commonRequest: commonRequest{
			Action: "eject_motions",
		},
		Motions:   motions,
		MotionIDs: motionIDs,
		Channel:   channel,
	}
}

type FinishMotionsRequest struct {
	commonRequest
	Motions   []string `json:"motions"`
	MotionIDs []string `json:"motion_ids"`
	Channel   *string  `json:"channel,omitempty"`
}

func NewFinishMotionsRequest(motions []string, motionIDs []string, channel *string) *FinishMotionsRequest {
	return &FinishMotionsRequest{
		commonRequest: commonRequest{
			Action: "finish_motions",
		},
		Motions:   motions,
		MotionIDs: motionIDs,
		Channel:   channel,
	}
}

type PauseMotionsRequest struct {
	commonRequest
	Motions   []string `json:"motions"`
	MotionIDs []string `json:"motion_ids"`
	Channel   *string  `json:"channel,omitempty"`
}

func NewPauseMotionsRequest(motions []string, motionIDs []string, channel *string) *PauseMotionsRequest {
	return &PauseMotionsRequest{
		commonRequest: commonRequest{
			Action: "pause_motions",
		},
		Motions:   motions,
		MotionIDs: motionIDs,
		Channel:   channel,
	}
}

type ResumeMotionsRequest struct {
	commonRequest
	Motions   []string `json:"motions"`
	MotionIDs []string `json:"motion_ids"`
	Channel   *string  `json:"channel,omitempty"`
}

func NewResumeMotionsRequest(motions []string, motionIDs []string, channel *string) *ResumeMotionsRequest {
	return &ResumeMotionsRequest{
		commonRequest: commonRequest{
			Action: "resume_motions",
		},
		Motions:   motions,
		MotionIDs: motionIDs,
		Channel:   channel,
	}
}

type RestartMotionsRequest struct {
	commonRequest
	Motions   []string `json:"motions"`
	MotionIDs []string `json:"motion_ids"`
	Channel   *string  `json:"channel,omitempty"`
}

func NewRestartMotionsRequest(motions []string, motionIDs []string, channel *string) *RestartMotionsRequest {
	return &RestartMotionsRequest{
		commonRequest: commonRequest{
			Action: "restart_motions",
		},
		Motions:   motions,
		MotionIDs: motionIDs,
		Channel:   channel,
	}
}

type FinishAndRestartMotionsRequest struct {
	commonRequest
	Motions   []string `json:"motions"`
	MotionIDs []string `json:"motion_ids"`
	Channel   *string  `json:"channel,omitempty"`
}

func NewFinishAndRestartMotionsRequest(motions []string, motionIDs []string, channel *string) *FinishAndRestartMotionsRequest {
	return &FinishAndRestartMotionsRequest{
		commonRequest: commonRequest{
			Action: "finish_and_restart_motions",
		},
		Motions:   motions,
		MotionIDs: motionIDs,
		Channel:   channel,
	}
}

type SetTextRequest struct {
	commonRequest
	Layer   string   `json:"layer"`
	LayerID []string `json:"layer_id"`
	Value   string   `json:"value"`
	Channel *string  `json:"channel,omitempty"`
}

func NewSetTextRequest(layer string, layerID []string, value string, channel *string) *SetTextRequest {
	return &SetTextRequest{
		commonRequest: commonRequest{
			Action: "set_text",
		},
		Layer:   layer,
		LayerID: layerID,
		Value:   value,
		Channel: channel,
	}
}

type ListGridNamesRequest struct {
	commonRequest
}

func NewListGridNamesRequest() *ListGridNamesRequest {
	return &ListGridNamesRequest{
		commonRequest: commonRequest{
			Action: "list_grid_names",
		},
	}
}

type ListGridCellsRequest struct {
	commonRequest
	Grid string `json:"grid"`
}

func NewListGridCellsRequest(grid string) *ListGridCellsRequest {
	return &ListGridCellsRequest{
		commonRequest: commonRequest{
			Action: "list_grid_cells",
		},
		Grid: grid,
	}
}

type ActivateGridCellRequest struct {
	commonRequest
	Grid string `json:"grid"`
	Cell [2]int `json:"cell"`
}

func NewActivateGridCellRequest(grid string, cell [2]int) *ActivateGridCellRequest {
	return &ActivateGridCellRequest{
		commonRequest: commonRequest{
			Action: "activate_grid_cell",
		},
		Grid: grid,
		Cell: cell,
	}
}

// run_data_source_query
// set_data_source_query_parameter
// select_data_source_rows

// tricaster macros??
