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

// Request interface for all request types
type Request interface {
	// This interface allows us to group all request types without using any/interface{}
}

type ListMotionsRequest struct {
	commonRequest
}

func RequestToReader(v Request) (io.Reader, error) {
	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(v); err != nil {
		return nil, xerrors.Errorf("unexpected error encoding request")
	}
	return buf, nil
}

func RequestsToReader(vs []Request) (io.Reader, error) {
	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(vs); err != nil {
		return nil, xerrors.Errorf("unexpected error encoding request array: %w", err)
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

type RunDataSourceQueryRequest struct {
	commonRequest
	DataSource string            `json:"data_source"`
	Query      string            `json:"query"`
	Parameters map[string]string `json:"parameters,omitempty"`
}

func NewRunDataSourceQueryRequest(dataSource string, query string, parameters map[string]string) *RunDataSourceQueryRequest {
	return &RunDataSourceQueryRequest{
		commonRequest: commonRequest{
			Action: "run_data_source_query",
		},
		DataSource: dataSource,
		Query:      query,
		Parameters: parameters,
	}
}

type SetDataSourceQueryParameterRequest struct {
	commonRequest
	DataSource string `json:"data_source"`
	Parameter  string `json:"parameter"`
	Value      string `json:"value"`
}

func NewSetDataSourceQueryParameterRequest(dataSource string, parameter string, value string) *SetDataSourceQueryParameterRequest {
	return &SetDataSourceQueryParameterRequest{
		commonRequest: commonRequest{
			Action: "set_data_source_query_parameter",
		},
		DataSource: dataSource,
		Parameter:  parameter,
		Value:      value,
	}
}

type SelectDataSourceRowsRequest struct {
	commonRequest
	DataSource string `json:"data_source"`
	Rows       []int  `json:"rows"`
}

func NewSelectDataSourceRowsRequest(dataSource string, rows []int) *SelectDataSourceRowsRequest {
	return &SelectDataSourceRowsRequest{
		commonRequest: commonRequest{
			Action: "select_data_source_rows",
		},
		DataSource: dataSource,
		Rows:       rows,
	}
}

// TriCaster Macro Control

type TriCasterMacroRequest struct {
	commonRequest
	MacroName  string            `json:"macro_name"`
	Parameters map[string]string `json:"parameters,omitempty"`
}

func NewTriCasterMacroRequest(macroName string, parameters map[string]string) *TriCasterMacroRequest {
	return &TriCasterMacroRequest{
		commonRequest: commonRequest{
			Action: "tricaster_macro",
		},
		MacroName:  macroName,
		Parameters: parameters,
	}
}

// NewTek Standard Commands

type NewTekCommandRequest struct {
	commonRequest
	Command    string            `json:"command"`
	Parameters map[string]string `json:"parameters,omitempty"`
}

func NewNewTekCommandRequest(command string, parameters map[string]string) *NewTekCommandRequest {
	return &NewTekCommandRequest{
		commonRequest: commonRequest{
			Action: "newtek_command",
		},
		Command:    command,
		Parameters: parameters,
	}
}

// CW Specific Commands

type CWSpecificCommandRequest struct {
	commonRequest
	Command    string            `json:"command"`
	Parameters map[string]string `json:"parameters,omitempty"`
}

func NewCWSpecificCommandRequest(command string, parameters map[string]string) *CWSpecificCommandRequest {
	return &CWSpecificCommandRequest{
		commonRequest: commonRequest{
			Action: "cw_specific_command",
		},
		Command:    command,
		Parameters: parameters,
	}
}
