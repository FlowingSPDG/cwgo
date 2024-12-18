package main

import (
	"fmt"

	cw "github.com/FlowingSPDG/cwgo"
)

func main() {
	// client
	client := cw.NewCharacterWorks("192.168.10.121", 50)

	listGridNamesResponse, err := client.ListGridNames()
	if err != nil {
		panic(err)
	}
	fmt.Printf("listGridNamesResponse: %v\n", listGridNamesResponse)

	for _, name := range listGridNamesResponse.Grids {
		listGridsResponse, err := client.ListGridCells(name)
		if err != nil {
			panic(err)
		}
		fmt.Printf("listGridsResponse: %v\n", listGridsResponse)

		for _, grid := range listGridsResponse.Grids {
			fmt.Printf("grid: %v\n", grid)
			for _, cell := range grid.Cells {
				activateGridCellResponse, err := client.ActivateGridCell(name, cell.Position)
				if err != nil {
					panic(err)
				}
				fmt.Printf("activateGridCellResponse: %v\n", activateGridCellResponse)
			}
		}
	}
}
