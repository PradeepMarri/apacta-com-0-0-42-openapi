package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/apacta/mcp-server/config"
	"github.com/apacta/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Get_companies_company_id_price_margins_price_margins_idHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		company_idVal, ok := args["company_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: company_id"), nil
		}
		company_id, ok := company_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: company_id"), nil
		}
		price_margins_idVal, ok := args["price_margins_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: price_margins_id"), nil
		}
		price_margins_id, ok := price_margins_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: price_margins_id"), nil
		}
		url := fmt.Sprintf("%s/companies/%s/price_margins/%s", cfg.BaseURL, company_id, price_margins_id)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Fallback to single auth parameter
		if cfg.APIKey != "" {
			req.Header.Set("X-Auth-Token", cfg.APIKey)
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateGet_companies_company_id_price_margins_price_margins_idTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_companies_company_id_price_margins_price_margins_id",
		mcp.WithDescription("Get a list of company price margins"),
		mcp.WithString("company_id", mcp.Required(), mcp.Description("")),
		mcp.WithString("price_margins_id", mcp.Required(), mcp.Description("Automatically added")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_companies_company_id_price_margins_price_margins_idHandler(cfg),
	}
}
