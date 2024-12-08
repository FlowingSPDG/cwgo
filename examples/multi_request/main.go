package main

import (
	"fmt"

	cw "github.com/FlowingSPDG/cwgo"
)

func main() {
	// client
	client := cw.NewCharacterWorks("localhost")

	requestsBuilder := client.RequestsBuilder()
	requestsBuilder.AddListMotionsRequest()
	requestsBuilder.AddListMotionsWithIDsRequest()
	requestsBuilder.AddListLayersRequest(nil, nil)
	requestsBuilder.AddPlayMotionsRequest(nil, nil, nil)
	requestsBuilder.AddStopMotionsRequest(nil, nil, nil)
	requestsBuilder.AddEjectMotionsRequest(nil, nil, nil)
	requestsBuilder.AddFinishMotionsRequest(nil, nil, nil)
	requestsBuilder.AddPauseMotionsRequest(nil, nil, nil)
	requestsBuilder.AddResumeMotionsRequest(nil, nil, nil)
	requestsBuilder.AddRestartMotionsRequest(nil, nil, nil)
	requestsBuilder.AddFinishAndRestartMotionsRequest(nil, nil, nil)

	resp, err := requestsBuilder.Do(client)
	if err != nil {
		panic(err)
	}

	fmt.Printf("response: %v\n", resp)
}
