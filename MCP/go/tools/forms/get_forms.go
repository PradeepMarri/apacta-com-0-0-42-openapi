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

func Get_formsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["extended"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("extended=%v", val))
		}
		if val, ok := args["date_from"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("date_from=%v", val))
		}
		if val, ok := args["date_to"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("date_to=%v", val))
		}
		if val, ok := args["show"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("show=%v", val))
		}
		if val, ok := args["project_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("project_id=%v", val))
		}
		if val, ok := args["created_by_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("created_by_id=%v", val))
		}
		if val, ok := args["form_template_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("form_template_id=%v", val))
		}
		if val, ok := args["form_template_type"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("form_template_type=%v", val))
		}
		if val, ok := args["employee_name"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("employee_name=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/forms%s", cfg.BaseURL, queryString)
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

func CreateGet_formsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_forms",
		mcp.WithDescription("Retrieve array of forms"),
		mcp.WithString("extended", mcp.Description("Used to have extended details from the forms from the `Form`'s `FormFields`")),
		mcp.WithString("date_from", mcp.Description("Used in conjunction with `date_to` to only show forms within these dates - format like `2016-28-05`")),
		mcp.WithString("date_to", mcp.Description("Used in conjunction with `date_from` to only show forms within these dates - format like `2016-28-30`")),
		mcp.WithString("show", mcp.Description("Used to show forms with trashed")),
		mcp.WithString("project_id", mcp.Description("Used to filter on the `project_id` of the forms")),
		mcp.WithString("created_by_id", mcp.Description("Used to filter on the `created_by_id` of the forms")),
		mcp.WithArray("form_template_id", mcp.Description("Used to filter on the `form_template_id` of the forms. Accept single value and array.")),
		mcp.WithString("form_template_type", mcp.Description("Filter by `form_templates.identifier` containing string passed in `form_template_type`. Accept strings like [`qa`, `dagseddel`]")),
		mcp.WithString("employee_name", mcp.Description("Used to filter forms by user's first or last name")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_formsHandler(cfg),
	}
}
