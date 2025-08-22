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

func Get_projectsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["show_all"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("show_all=%v", val))
		}
		if val, ok := args["sort"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sort=%v", val))
		}
		if val, ok := args["direction"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("direction=%v", val))
		}
		if val, ok := args["contact_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("contact_id=%v", val))
		}
		if val, ok := args["company_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("company_id=%v", val))
		}
		if val, ok := args["project_status_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("project_status_id=%v", val))
		}
		if val, ok := args["project_status_ids"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("project_status_ids=%v", val))
		}
		if val, ok := args["name"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("name=%v", val))
		}
		if val, ok := args["erp_project_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("erp_project_id=%v", val))
		}
		if val, ok := args["erp_task_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("erp_task_id=%v", val))
		}
		if val, ok := args["start_time_gte"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("start_time_gte=%v", val))
		}
		if val, ok := args["start_time_lte"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("start_time_lte=%v", val))
		}
		if val, ok := args["start_time_eq"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("start_time_eq=%v", val))
		}
		if val, ok := args["event_start[][gt]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("event_start[][gt]=%v", val))
		}
		if val, ok := args["event_start[][lt]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("event_start[][lt]=%v", val))
		}
		if val, ok := args["event_start[][eq]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("event_start[][eq]=%v", val))
		}
		if val, ok := args["event_end[][gt]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("event_end[][gt]=%v", val))
		}
		if val, ok := args["event_end[][lt]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("event_end[][lt]=%v", val))
		}
		if val, ok := args["event_end[][eq]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("event_end[][eq]=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/projects%s", cfg.BaseURL, queryString)
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

func CreateGet_projectsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_projects",
		mcp.WithDescription("View list of projects"),
		mcp.WithBoolean("show_all", mcp.Description("Unless this is set to `true` only open projects will be shown")),
		mcp.WithString("sort", mcp.Description("Sort projects by `not_invoiced_amount`")),
		mcp.WithString("direction", mcp.Description("")),
		mcp.WithString("contact_id", mcp.Description("Used to filter on the `contact_id` of the projects")),
		mcp.WithString("company_id", mcp.Description("Used to filter on the `company_id` of the projects")),
		mcp.WithString("project_status_id", mcp.Description("Used to filter on the `project_status_id` of the projects")),
		mcp.WithArray("project_status_ids", mcp.Description("Used to filter on the `project_status_id` of the projects (match any of the provided values)")),
		mcp.WithString("name", mcp.Description("Used to search on the `name` of the projects")),
		mcp.WithString("erp_project_id", mcp.Description("Used to search on the `erp_project_id` of the projects")),
		mcp.WithString("erp_task_id", mcp.Description("Used to search on the `erp_task_id` of the projects")),
		mcp.WithString("start_time_gte", mcp.Description("Show projects with start time after than this value")),
		mcp.WithString("start_time_lte", mcp.Description("Show projects with start time before this value")),
		mcp.WithString("start_time_eq", mcp.Description("Show only projects with start time on specific date")),
		mcp.WithString("event_start[][gt]", mcp.Description("")),
		mcp.WithString("event_start[][lt]", mcp.Description("")),
		mcp.WithString("event_start[][eq]", mcp.Description("")),
		mcp.WithString("event_end[][gt]", mcp.Description("")),
		mcp.WithString("event_end[][lt]", mcp.Description("")),
		mcp.WithString("event_end[][eq]", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_projectsHandler(cfg),
	}
}
