package main

import (
	"context"
	"fmt"

	cw "github.com/FlowingSPDG/cwgo"
)

func main() {
	// CharacterWorksインスタンスを作成
	client := cw.NewCharacterWorks("localhost", 10.0) // 10 requests per second limit

	ctx := context.Background()

	// 1. JSON配列リクエストのテスト
	fmt.Println("=== Testing JSON Array Requests ===")
	testBatchRequests(ctx, client)

	// 2. Data Source機能のテスト
	fmt.Println("\n=== Testing Data Source Features ===")
	testDataSourceFeatures(ctx, client)

	// 3. TriCaster Macro機能のテスト
	fmt.Println("\n=== Testing TriCaster Macro Features ===")
	testTriCasterFeatures(ctx, client)
}

func testBatchRequests(ctx context.Context, client *cw.CharacterWorks) {
	// 複数のリクエストを一括送信
	// 型安全な方法でリクエストを作成
	listMotionsReq := cw.NewListMotionsRequest()
	listGridNamesReq := cw.NewListGridNamesRequest()

	fmt.Println("Batch requests created successfully")
	fmt.Printf("ListMotions Request: %+v\n", listMotionsReq)
	fmt.Printf("ListGridNames Request: %+v\n", listGridNamesReq)

	// 実際のAPI呼び出しをテスト（エラーは無視して構造確認のみ）
	fmt.Println("Testing individual API calls...")

	// ListMotions API呼び出し
	if _, err := client.ListMotions(ctx); err != nil {
		fmt.Printf("ListMotions API call failed (expected): %v\n", err)
	} else {
		fmt.Println("ListMotions API call succeeded")
	}

	// ListGridNames API呼び出し
	if _, err := client.ListGridNames(ctx); err != nil {
		fmt.Printf("ListGridNames API call failed (expected): %v\n", err)
	} else {
		fmt.Println("ListGridNames API call succeeded")
	}
}

func testDataSourceFeatures(ctx context.Context, client *cw.CharacterWorks) {
	// Data Source Query
	params := map[string]string{
		"param1": "value1",
		"param2": "value2",
	}

	queryReq := cw.NewRunDataSourceQueryRequest("test_datasource", "SELECT * FROM table", params)
	fmt.Printf("Data Source Query Request: %+v\n", queryReq)

	// Data Source Parameter
	paramReq := cw.NewSetDataSourceQueryParameterRequest("test_datasource", "param1", "new_value")
	fmt.Printf("Data Source Parameter Request: %+v\n", paramReq)

	// Data Source Rows
	rowsReq := cw.NewSelectDataSourceRowsRequest("test_datasource", []int{1, 2, 3})
	fmt.Printf("Data Source Rows Request: %+v\n", rowsReq)

	// 実際のAPI呼び出しをテスト（エラーは無視して構造確認のみ）
	fmt.Println("Testing Data Source API calls...")

	// RunDataSourceQuery API呼び出し
	if _, err := client.RunDataSourceQuery(ctx, "test_datasource", "SELECT * FROM table", params); err != nil {
		fmt.Printf("RunDataSourceQuery API call failed (expected): %v\n", err)
	} else {
		fmt.Println("RunDataSourceQuery API call succeeded")
	}

	// SetDataSourceQueryParameter API呼び出し
	if _, err := client.SetDataSourceQueryParameter(ctx, "test_datasource", "param1", "new_value"); err != nil {
		fmt.Printf("SetDataSourceQueryParameter API call failed (expected): %v\n", err)
	} else {
		fmt.Println("SetDataSourceQueryParameter API call succeeded")
	}

	// SelectDataSourceRows API呼び出し
	if _, err := client.SelectDataSourceRows(ctx, "test_datasource", []int{1, 2, 3}); err != nil {
		fmt.Printf("SelectDataSourceRows API call failed (expected): %v\n", err)
	} else {
		fmt.Println("SelectDataSourceRows API call succeeded")
	}
}

func testTriCasterFeatures(ctx context.Context, client *cw.CharacterWorks) {
	// TriCaster Macro
	macroParams := map[string]string{
		"input":  "1",
		"output": "2",
	}

	macroReq := cw.NewTriCasterMacroRequest("test_macro", macroParams)
	fmt.Printf("TriCaster Macro Request: %+v\n", macroReq)

	// NewTek Command
	newtekParams := map[string]string{
		"command": "TAKE",
		"input":   "1",
	}

	newtekReq := cw.NewNewTekCommandRequest("TAKE", newtekParams)
	fmt.Printf("NewTek Command Request: %+v\n", newtekReq)

	// CW Specific Command
	cwParams := map[string]string{
		"action": "custom_action",
		"value":  "test_value",
	}

	cwReq := cw.NewCWSpecificCommandRequest("custom_action", cwParams)
	fmt.Printf("CW Specific Command Request: %+v\n", cwReq)

	// 実際のAPI呼び出しをテスト（エラーは無視して構造確認のみ）
	fmt.Println("Testing TriCaster API calls...")

	// ExecuteTriCasterMacro API呼び出し
	if _, err := client.ExecuteTriCasterMacro(ctx, "test_macro", macroParams); err != nil {
		fmt.Printf("ExecuteTriCasterMacro API call failed (expected): %v\n", err)
	} else {
		fmt.Println("ExecuteTriCasterMacro API call succeeded")
	}

	// ExecuteNewTekCommand API呼び出し
	if _, err := client.ExecuteNewTekCommand(ctx, "TAKE", newtekParams); err != nil {
		fmt.Printf("ExecuteNewTekCommand API call failed (expected): %v\n", err)
	} else {
		fmt.Println("ExecuteNewTekCommand API call succeeded")
	}

	// ExecuteCWSpecificCommand API呼び出し
	if _, err := client.ExecuteCWSpecificCommand(ctx, "custom_action", cwParams); err != nil {
		fmt.Printf("ExecuteCWSpecificCommand API call failed (expected): %v\n", err)
	} else {
		fmt.Println("ExecuteCWSpecificCommand API call succeeded")
	}
}
