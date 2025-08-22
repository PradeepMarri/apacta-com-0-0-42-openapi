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

func Get_contactsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["name"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("name=%v", val))
		}
		if val, ok := args["cvr"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("cvr=%v", val))
		}
		if val, ok := args["ean"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ean=%v", val))
		}
		if val, ok := args["erp_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("erp_id=%v", val))
		}
		if val, ok := args["contact_type"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("contact_type=%v", val))
		}
		if val, ok := args["city"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("city=%v", val))
		}
		if val, ok := args["modified_gte"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("modified_gte=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/contacts%s", cfg.BaseURL, queryString)
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

func CreateGet_contactsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_contacts",
		mcp.WithDescription("Get a list of contacts"),
		mcp.WithString("name", mcp.Description("Used to search for a contact with a specific name")),
		mcp.WithString("cvr", mcp.Description("Search for values in CVR field")),
		mcp.WithString("ean", mcp.Description("Search for values in EAN field")),
		mcp.WithString("erp_id", mcp.Description("Search for values in ERP id field")),
		mcp.WithString("contact_type", mcp.Description("Used to show only contacts with with one specific `ContactType`")),
		mcp.WithString("city", mcp.Description("Used to show only contacts with with one specific `City`")),
		mcp.WithString("modified_gte", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_contactsHandler(cfg),
	}
}
