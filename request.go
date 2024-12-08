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

func newListMotionsRequest() io.Reader {
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

func newListMotionsWithIDsRequest() io.Reader {
	b := &ListMotionsWithIDsRequest{
		commonRequest: commonRequest{
			Action: "list_motions_with_ids",
		},
	}
	return toReader(b)
}

type ListLayersRequest struct {
	commonRequest
	Parent  *string `json:"parent"`
	Channel *string `json:"channel,omitempty"`
}

func newListLayersRequest(parent *string, channel *string) io.Reader {
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

func newPlayMotionsRequest(motions []string, motionIDs []string, channel *string) io.Reader {
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

func newStopMotionsRequest(motions []string, motionIDs []string, channel *string) io.Reader {
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

func newEjectMotionsRequest(motions []string, motionIDs []string, channel *string) io.Reader {
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

func newFinishMotionsRequest(motions []string, motionIDs []string, channel *string) io.Reader {
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

func newPauseMotionsRequest(motions []string, motionIDs []string, channel *string) io.Reader {
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

func newResumeMotionsRequest(motions []string, motionIDs []string, channel *string) io.Reader {
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

func newRestartMotionsRequest(motions []string, motionIDs []string, channel *string) io.Reader {
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

func newFinishAndRestartMotionsRequest(motions []string, motionIDs []string, channel *string) io.Reader {
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

// set_text
// list_grid_names
// list_grid_cells
// activate_grid_cell
// run_data_source_query
// set_data_source_query_parameter
// select_data_source_rows

// tricaster macros??