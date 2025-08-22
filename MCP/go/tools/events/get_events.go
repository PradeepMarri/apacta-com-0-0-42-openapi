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

func Get_eventsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["user_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("user_id=%v", val))
		}
		if val, ok := args["project_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("project_id=%v", val))
		}
		if val, ok := args["start[][gt]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("start[][gt]=%v", val))
		}
		if val, ok := args["start[][lt]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("start[][lt]=%v", val))
		}
		if val, ok := args["start[][eq]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("start[][eq]=%v", val))
		}
		if val, ok := args["end[][gt]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("end[][gt]=%v", val))
		}
		if val, ok := args["end[][lt]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("end[][lt]=%v", val))
		}
		if val, ok := args["end[][eq]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("end[][eq]=%v", val))
		}
		if val, ok := args["tags"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("tags=%v", val))
		}
		if val, ok := args["without_users"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("without_users=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/events%s", cfg.BaseURL, queryString)
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

func CreateGet_eventsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_events",
		mcp.WithDescription("Show list of events"),
		mcp.WithString("user_id", mcp.Description("")),
		mcp.WithString("project_id", mcp.Description("")),
		mcp.WithString("start[][gt]", mcp.Description("")),
		mcp.WithString("start[][lt]", mcp.Description("")),
		mcp.WithString("start[][eq]", mcp.Description("")),
		mcp.WithString("end[][gt]", mcp.Description("")),
		mcp.WithString("end[][lt]", mcp.Description("")),
		mcp.WithString("end[][eq]", mcp.Description("")),
		mcp.WithString("tags", mcp.Description("List events with given tag ids separated by comma. Example tags=0377d6ce-db5e-4b1e-ac3a-8ca39ea3142e,8cec327e-a559-4b40-9ed6-316b9de46743")),
		mcp.WithBoolean("without_users", mcp.Description("List events without attached user")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_eventsHandler(cfg),
	}
}
