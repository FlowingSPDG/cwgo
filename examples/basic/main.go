package main

import (
	"fmt"

	cw "github.com/FlowingSPDG/cwgo"
)

func main() {
	// client
	client := cw.NewCharacterWorks("localhost")

	// listMotions
	listMotionsResponse, err := client.ListMotions()
	if err != nil {
		panic(err)
	}
	fmt.Printf("listMotionsResponse: %v\n", listMotionsResponse)

	// listMotionsWithID
	listMotionsWithIDResponse, err := client.ListMotionsWithIDs()
	if err != nil {
		panic(err)
	}
	fmt.Printf("listMotionsWithIDResponse: %v\n", listMotionsWithIDResponse)

	// listLayers
	listLayersResponse, err := client.ListLayers(nil, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("listLayersResponse: %v\n", listLayersResponse)

	// playMotions
	playMotionsResponse, err := client.PlayMotions(nil, nil, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("playMotionsResponse: %v\n", playMotionsResponse)

	// stopMotions
	stopMotionsResponse, err := client.StopMotions(nil, nil, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("stopMotionsResponse: %v\n", stopMotionsResponse)

	// ejectMotions
	ejectMotionsResponse, err := client.EjectMotions(nil, nil, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ejectMotionsResponse: %v\n", ejectMotionsResponse)

	// finishMotions
	finishMotionsResponse, err := client.FinishMotions(nil, nil, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("finishMotionsResponse: %v\n", finishMotionsResponse)

	// pauseMotions
	pauseMotionsResponse, err := client.PauseMotions(nil, nil, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pauseMotionsResponse: %v\n", pauseMotionsResponse)

	// resumeMotions
	resumeMotionsResponse, err := client.ResumeMotions(nil, nil, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("resumeMotionsResponse: %v\n", resumeMotionsResponse)

	// restartMotions
	restartMotionsResponse, err := client.RestartMotions(nil, nil, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("restartMotionsResponse: %v\n", restartMotionsResponse)

	// finishAndRestartMotions
	finishAndRestartMotionsResponse, err := client.FinishAndRestartMotions(nil, nil, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("finishAndRestartMotionsResponse: %v\n", finishAndRestartMotionsResponse)
}
