package cw

import (
	"bytes"
	"encoding/json"
	"io"
)

type commonRequest struct {
	Action string `json:"action"`
}

type ListMotionsRequest struct {
	commonRequest
}

func toReader(v interface{}) io.Reader {
	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(v); err != nil {
		panic("unexpected error encoding request")
	}
	return buf
}

func NewListMotionsRequest() io.Reader {
	b := &ListMotionsRequest{
		commonRequest: commonRequest{
			Action: "list_motions",
		},
	}
	return toReader(b)
}

type ListMotionsWithIDsRequest struct {
	commonRequest
}

func NewListMotionsWithIDsRequest() io.Reader {
	b := &ListMotionsWithIDsRequest{
		commonRequest: commonRequest{
			Action: "list_motions_with_ids",
		},
	}
	return toReader(b)
}

type ListLayersRequest struct {
	commonRequest
	Parent  *string `json:"parent,omitempty"`
	Channel *string `json:"channel,omitempty"`
}

func NewListLayersRequest(parent *string, channel *string) io.Reader {
	b := &ListLayersRequest{
		commonRequest: commonRequest{
			Action: "list_layers",
		},
		Parent:  parent,
		Channel: channel,
	}
	return toReader(b)
}

type PlayMotionsRequest struct {
	commonRequest
	Motions   []string `json:"motions"`
	MotionIDs []string `json:"motion_ids"`
	Channel   *string  `json:"channel,omitempty"`
}

func NewPlayMotionsRequest(motions []string, motionIDs []string, channel *string) io.Reader {
	b := &PlayMotionsRequest{
		commonRequest: commonRequest{
			Action: "play_motions",
		},
		Motions:   motions,
		MotionIDs: motionIDs,
		Channel:   channel,
	}
	return toReader(b)
}

type StopMotionsRequest struct {
	commonRequest
	Motions   []string `json:"motions"`
	MotionIDs []string `json:"motion_ids"`
	Channel   *string  `json:"channel,omitempty"`
}

func NewStopMotionsRequest(motions []string, motionIDs []string, channel *string) io.Reader {
	b := &StopMotionsRequest{
		commonRequest: commonRequest{
			Action: "stop_motions",
		},
		Motions:   motions,
		MotionIDs: motionIDs,
		Channel:   channel,
	}
	return toReader(b)
}

type EjectMotionsRequest struct {
	commonRequest
	Motions   []string `json:"motions"`
	MotionIDs []string `json:"motion_ids"`
	Channel   *string  `json:"channel,omitempty"`
}

func NewEjectMotionsRequest(motions []string, motionIDs []string, channel *string) io.Reader {
	b := &EjectMotionsRequest{
		commonRequest: commonRequest{
			Action: "eject_motions",
		},
		Motions:   motions,
		MotionIDs: motionIDs,
		Channel:   channel,
	}
	return toReader(b)
}

type FinishMotionsRequest struct {
	commonRequest
	Motions   []string `json:"motions"`
	MotionIDs []string `json:"motion_ids"`
	Channel   *string  `json:"channel,omitempty"`
}

func NewFinishMotionsRequest(motions []string, motionIDs []string, channel *string) io.Reader {
	b := &FinishMotionsRequest{
		commonRequest: commonRequest{
			Action: "finish_motions",
		},
		Motions:   motions,
		MotionIDs: motionIDs,
		Channel:   channel,
	}
	return toReader(b)
}

type PauseMotionsRequest struct {
	commonRequest
	Motions   []string `json:"motions"`
	MotionIDs []string `json:"motion_ids"`
	Channel   *string  `json:"channel,omitempty"`
}

func NewPauseMotionsRequest(motions []string, motionIDs []string, channel *string) io.Reader {
	b := &PauseMotionsRequest{
		commonRequest: commonRequest{
			Action: "pause_motions",
		},
		Motions:   motions,
		MotionIDs: motionIDs,
		Channel:   channel,
	}
	return toReader(b)
}

type ResumeMotionsRequest struct {
	commonRequest
	Motions   []string `json:"motions"`
	MotionIDs []string `json:"motion_ids"`
	Channel   *string  `json:"channel,omitempty"`
}

func NewResumeMotionsRequest(motions []string, motionIDs []string, channel *string) io.Reader {
	b := &ResumeMotionsRequest{
		commonRequest: commonRequest{
			Action: "resume_motions",
		},
		Motions:   motions,
		MotionIDs: motionIDs,
		Channel:   channel,
	}
	return toReader(b)
}

type RestartMotionsRequest struct {
	commonRequest
	Motions   []string `json:"motions"`
	MotionIDs []string `json:"motion_ids"`
	Channel   *string  `json:"channel,omitempty"`
}

func NewRestartMotionsRequest(motions []string, motionIDs []string, channel *string) io.Reader {
	b := &RestartMotionsRequest{
		commonRequest: commonRequest{
			Action: "restart_motions",
		},
		Motions:   motions,
		MotionIDs: motionIDs,
		Channel:   channel,
	}
	return toReader(b)
}

type FinishAndRestartMotionsRequest struct {
	commonRequest
	Motions   []string `json:"motions"`
	MotionIDs []string `json:"motion_ids"`
	Channel   *string  `json:"channel,omitempty"`
}

func NewFinishAndRestartMotionsRequest(motions []string, motionIDs []string, channel *string) io.Reader {
	b := &FinishAndRestartMotionsRequest{
		commonRequest: commonRequest{
			Action: "finish_and_restart_motions",
		},
		Motions:   motions,
		MotionIDs: motionIDs,
		Channel:   channel,
	}
	return toReader(b)
}

type SetTextRequest struct {
	commonRequest
	Layer   string   `json:"layer"`
	LayerID []string `json:"layer_id"`
	Value   string   `json:"value"`
	Channel *string  `json:"channel,omitempty"`
}

func NewSetTextRequest(layer string, layerID []string, value string, channel *string) io.Reader {
	b := &SetTextRequest{
		commonRequest: commonRequest{
			Action: "set_text",
		},
		Layer:   layer,
		LayerID: layerID,
		Value:   value,
		Channel: channel,
	}
	return toReader(b)
}

// list_grid_names
// list_grid_cells
// activate_grid_cell
// run_data_source_query
// set_data_source_query_parameter
// select_data_source_rows

// tricaster macros??
