package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/apacta/mcp-server/config"
	"github.com/apacta/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Get_time_entriesHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["user_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("user_id=%v", val))
		}
		if val, ok := args["form_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("form_id=%v", val))
		}
		if val, ok := args["project_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("project_id=%v", val))
		}
		if val, ok := args["gt_from_time"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("gt_from_time=%v", val))
		}
		if val, ok := args["lt_from_time"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("lt_from_time=%v", val))
		}
		if val, ok := args["gt_to_time"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("gt_to_time=%v", val))
		}
		if val, ok := args["lt_to_time"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("lt_to_time=%v", val))
		}
		if val, ok := args["lt_sum"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("lt_sum=%v", val))
		}
		if val, ok := args["gt_sum"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("gt_sum=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/time_entries%s", cfg.BaseURL, queryString)
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

func CreateGet_time_entriesTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_time_entries",
		mcp.WithDescription("List time entries"),
		mcp.WithString("user_id", mcp.Description("")),
		mcp.WithString("form_id", mcp.Description("")),
		mcp.WithString("project_id", mcp.Description("")),
		mcp.WithString("gt_from_time", mcp.Description("")),
		mcp.WithString("lt_from_time", mcp.Description("")),
		mcp.WithString("gt_to_time", mcp.Description("")),
		mcp.WithString("lt_to_time", mcp.Description("")),
		mcp.WithString("lt_sum", mcp.Description("")),
		mcp.WithString("gt_sum", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_time_entriesHandler(cfg),
	}
}
