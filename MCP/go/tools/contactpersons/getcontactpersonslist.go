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

func GetcontactpersonslistHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		contact_idVal, ok := args["contact_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: contact_id"), nil
		}
		contact_id, ok := contact_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: contact_id"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["q"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("q=%v", val))
		}
		if val, ok := args["created_gte"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("created_gte=%v", val))
		}
		if val, ok := args["created_lte"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("created_lte=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/contacts/%s/contact_persons%s", cfg.BaseURL, contact_id, queryString)
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
		var result interface{}
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

func CreateGetcontactpersonslistTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_contacts_contact_id_contact_persons",
		mcp.WithDescription("Get a list of contact people"),
		mcp.WithString("contact_id", mcp.Required(), mcp.Description("")),
		mcp.WithString("q", mcp.Description("")),
		mcp.WithString("created_gte", mcp.Description("")),
		mcp.WithString("created_lte", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetcontactpersonslistHandler(cfg),
	}
}
