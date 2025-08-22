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

func Get_project_custom_field_attributes_project_custom_field_attribute_idHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		project_custom_field_attribute_idVal, ok := args["project_custom_field_attribute_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: project_custom_field_attribute_id"), nil
		}
		project_custom_field_attribute_id, ok := project_custom_field_attribute_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: project_custom_field_attribute_id"), nil
		}
		url := fmt.Sprintf("%s/project_custom_field_attributes/%s", cfg.BaseURL, project_custom_field_attribute_id)
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

func CreateGet_project_custom_field_attributes_project_custom_field_attribute_idTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_project_custom_field_attributes_project_custom_field_attribute_id",
		mcp.WithDescription("Details of 1 project custom field attribute"),
		mcp.WithString("project_custom_field_attribute_id", mcp.Required(), mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_project_custom_field_attributes_project_custom_field_attribute_idHandler(cfg),
	}
}
