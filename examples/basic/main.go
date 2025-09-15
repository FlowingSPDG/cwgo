package main

import (
	"context"
	"fmt"

	cw "github.com/FlowingSPDG/cwgo"
)

func main() {
	// client
	client := cw.NewCharacterWorks("localhost", 50)
	ctx := context.Background()

	// listMotions
	listMotionsResponse, err := client.ListMotions(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("listMotionsResponse: %v\n", listMotionsResponse)

	// listMotionsWithID
	listMotionsWithIDResponse, err := client.ListMotionsWithIDs(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("listMotionsWithIDResponse: %v\n", listMotionsWithIDResponse)

	// listLayers
	listLayersResponse, err := client.ListLayers(ctx, nil, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("listLayersResponse: %v\n", listLayersResponse)

	// playMotions
	playMotionsResponse, err := client.PlayMotions(ctx, nil, nil, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("playMotionsResponse: %v\n", playMotionsResponse)

	// stopMotions
	stopMotionsResponse, err := client.StopMotions(ctx, nil, nil, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("stopMotionsResponse: %v\n", stopMotionsResponse)

	// ejectMotions
	ejectMotionsResponse, err := client.EjectMotions(ctx, nil, nil, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ejectMotionsResponse: %v\n", ejectMotionsResponse)

	// finishMotions
	finishMotionsResponse, err := client.FinishMotions(ctx, nil, nil, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("finishMotionsResponse: %v\n", finishMotionsResponse)

	// pauseMotions
	pauseMotionsResponse, err := client.PauseMotions(ctx, nil, nil, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pauseMotionsResponse: %v\n", pauseMotionsResponse)

	// resumeMotions
	resumeMotionsResponse, err := client.ResumeMotions(ctx, nil, nil, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("resumeMotionsResponse: %v\n", resumeMotionsResponse)

	// restartMotions
	restartMotionsResponse, err := client.RestartMotions(ctx, nil, nil, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("restartMotionsResponse: %v\n", restartMotionsResponse)

	// finishAndRestartMotions
	finishAndRestartMotionsResponse, err := client.FinishAndRestartMotions(ctx, nil, nil, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("finishAndRestartMotionsResponse: %v\n", finishAndRestartMotionsResponse)
}
