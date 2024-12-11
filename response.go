package cw

type commonResult struct {
	Result string `json:"result"`
}

type ListMotionsResponse struct {
	commonResult
	Motions []string `json:"motions"`
}

type ListMotionsWithIDsResponse struct {
	commonResult
	Motions []struct {
		Name string `json:"name"`
		ID   string `json:"id"`
	} `json:"motions"`
}

type ListLayersResponse struct {
	commonResult
	Layers []Layer `json:"layers"`
}

type Layer struct {
	Name     string  `json:"name"`
	Path     string  `json:"path"`
	ID       string  `json:"id"`
	Type     string  `json:"type"`
	Children []Layer `json:"children,omitempty"`
}

type commonWarnings struct {
	Warnings []struct {
		Code        string `json:"code"`
		Description string `json:"description"`
	} `json:"warnings"`
}

type commonErrors struct {
	Errors []struct {
		Code        string `json:"code"`
		Description string `json:"description"`
	} `json:"errors"`
}

type PlayMotionsResponse struct {
	commonResult
	commonWarnings
}

type StopMotionsResponse struct {
	commonResult
	commonWarnings
}

type EjectMotionsResponse struct {
	commonResult
	commonWarnings
}

type FinishMotionsResponse struct {
	commonResult
	commonWarnings
}

type PauseMotionsResponse struct {
	commonResult
	commonWarnings
}

type ResumeMotionsResponse struct {
	commonResult
	commonWarnings
}

type RestartMotionsResponse struct {
	commonResult
	commonWarnings
}

type FinishAndRestartMotionsResponse struct {
	commonResult
	commonWarnings
}

type SetTextResponse struct {
	commonResult
	commonErrors
}

// list_grid_names
// list_grid_cells
// activate_grid_cell
// run_data_source_query
// set_data_source_query_parameter
// select_data_source_rows

// tricaster macros??
