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

func Delete_companies_company_id_integration_settings_companies_integration_setting_idHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		companies_integration_setting_idVal, ok := args["companies_integration_setting_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: companies_integration_setting_id"), nil
		}
		companies_integration_setting_id, ok := companies_integration_setting_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: companies_integration_setting_id"), nil
		}
		url := fmt.Sprintf("%s/companies/%s/integration_settings/%s", cfg.BaseURL, company_id, companies_integration_setting_id)
		req, err := http.NewRequest("DELETE", url, nil)
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
		var result models.EmptySuccessResponse
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

func CreateDelete_companies_company_id_integration_settings_companies_integration_setting_idTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_companies_company_id_integration_settings_companies_integration_setting_id",
		mcp.WithDescription("Delete a company integration setting"),
		mcp.WithString("company_id", mcp.Required(), mcp.Description("")),
		mcp.WithString("companies_integration_setting_id", mcp.Required(), mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Delete_companies_company_id_integration_settings_companies_integration_setting_idHandler(cfg),
	}
}
