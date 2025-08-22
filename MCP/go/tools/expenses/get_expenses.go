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

func Get_expensesHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["created_by_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("created_by_id=%v", val))
		}
		if val, ok := args["company_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("company_id=%v", val))
		}
		if val, ok := args["contact_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("contact_id=%v", val))
		}
		if val, ok := args["project_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("project_id=%v", val))
		}
		if val, ok := args["due_date"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("due_date=%v", val))
		}
		if val, ok := args["gt_created"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("gt_created=%v", val))
		}
		if val, ok := args["lt_created"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("lt_created=%v", val))
		}
		if val, ok := args["status"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("status=%v", val))
		}
		if val, ok := args["is_imported"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("is_imported=%v", val))
		}
		if val, ok := args["min_amount"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("min_amount=%v", val))
		}
		if val, ok := args["max_amount"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("max_amount=%v", val))
		}
		if val, ok := args["projects"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("projects=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/expenses%s", cfg.BaseURL, queryString)
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

func CreateGet_expensesTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_expenses",
		mcp.WithDescription("Show list of expenses"),
		mcp.WithString("created_by_id", mcp.Description("")),
		mcp.WithString("company_id", mcp.Description("")),
		mcp.WithString("contact_id", mcp.Description("")),
		mcp.WithString("project_id", mcp.Description("")),
		mcp.WithString("due_date", mcp.Description("Filter by [valid=records in future including today], [exceeded=records in past] or [null=all records]")),
		mcp.WithString("gt_created", mcp.Description("Created after date")),
		mcp.WithString("lt_created", mcp.Description("Created before date")),
		mcp.WithString("status", mcp.Description("Filter by status identifier. [null=all records]")),
		mcp.WithBoolean("is_imported", mcp.Description("")),
		mcp.WithString("min_amount", mcp.Description("Expenses `total_selling_price` > `min_amount`")),
		mcp.WithString("max_amount", mcp.Description("Expenses `total_selling_price` < `max_amount`")),
		mcp.WithString("projects", mcp.Description("You can select `all projects`, `no projects` or select `multiple projects`")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_expensesHandler(cfg),
	}
}
