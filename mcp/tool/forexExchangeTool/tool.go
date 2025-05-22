package tool

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mark3labs/mcp-go/server"
	"io/ioutil"
	"net/http"

	"github.com/mark3labs/mcp-go/mcp"
)

func NewForexExchangeRateTool(s *server.MCPServer) {
	tl := mcp.NewTool("forex_exchange_rate",
		mcp.WithDescription("获取指定货币对的实时汇率数据"),
		mcp.WithString("currencyPair", mcp.Required(), mcp.Description("货币对代码，如 USDCNY 美元兑换人民币")))
	s.AddTool(tl, exchangeRateHandler)
}

func exchangeRateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	currencyPair, ok := request.GetArguments()["currencyPair"].(string)
	if !ok {
		return nil, errors.New("to_currency must be a string")
	}

	url := fmt.Sprintf("https://forex.finance.sina.cn/api/getExchangeRate?symbol=%s", currencyPair)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch exchange rate: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	var rateResp struct {
		Status struct {
			Code int    `json:"code"`
			Msg  string `json:"msg"`
		} `json:"status"`
		Data struct {
			Symbol string  `json:"symbol"`
			Price  float64 `json:"price"`
		} `json:"data"`
	}

	if err := json.Unmarshal(body, &rateResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	result := fmt.Sprintf("Exchange Rate %s/%s: %.4f",
		currencyPair[0:3], currencyPair[3:6],
		rateResp.Data.Price)
	return mcp.NewToolResultText(result), nil
}
