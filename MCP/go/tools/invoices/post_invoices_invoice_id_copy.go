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

func Post_invoices_invoice_id_copyHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		invoice_idVal, ok := args["invoice_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: invoice_id"), nil
		}
		invoice_id, ok := invoice_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: invoice_id"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["project_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("project_id=%v", val))
		}
		if val, ok := args["contact_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("contact_id=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/invoices/%s/copy%s", cfg.BaseURL, invoice_id, queryString)
		req, err := http.NewRequest("POST", url, nil)
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
		var result models.CreateSuccessResponse
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

func CreatePost_invoices_invoice_id_copyTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_invoices_invoice_id_copy",
		mcp.WithDescription("Create a copy of an invoice"),
		mcp.WithString("invoice_id", mcp.Required(), mcp.Description("")),
		mcp.WithString("project_id", mcp.Description("")),
		mcp.WithString("contact_id", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Post_invoices_invoice_id_copyHandler(cfg),
	}
}
