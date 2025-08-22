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

func Get_users_user_id_integration_settings_integration_settings_user_idHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		user_idVal, ok := args["user_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: user_id"), nil
		}
		user_id, ok := user_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: user_id"), nil
		}
		integration_settings_user_idVal, ok := args["integration_settings_user_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: integration_settings_user_id"), nil
		}
		integration_settings_user_id, ok := integration_settings_user_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: integration_settings_user_id"), nil
		}
		url := fmt.Sprintf("%s/users/%s/integration_settings/%s", cfg.BaseURL, user_id, integration_settings_user_id)
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

func CreateGet_users_user_id_integration_settings_integration_settings_user_idTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_users_user_id_integration_settings_integration_settings_user_id",
		mcp.WithDescription("Get a user integration setting"),
		mcp.WithString("user_id", mcp.Required(), mcp.Description("")),
		mcp.WithString("integration_settings_user_id", mcp.Required(), mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_users_user_id_integration_settings_integration_settings_user_idHandler(cfg),
	}
}
