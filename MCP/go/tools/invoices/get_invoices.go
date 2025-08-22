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

func Get_invoicesHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["contact_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("contact_id=%v", val))
		}
		if val, ok := args["project_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("project_id=%v", val))
		}
		if val, ok := args["invoice_number"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("invoice_number=%v", val))
		}
		if val, ok := args["offer_number"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("offer_number=%v", val))
		}
		if val, ok := args["is_draft"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("is_draft=%v", val))
		}
		if val, ok := args["is_offer"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("is_offer=%v", val))
		}
		if val, ok := args["is_locked"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("is_locked=%v", val))
		}
		if val, ok := args["is_fixed_price"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("is_fixed_price=%v", val))
		}
		if val, ok := args["date_from"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("date_from=%v", val))
		}
		if val, ok := args["date_to"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("date_to=%v", val))
		}
		if val, ok := args["issued_date"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("issued_date=%v", val))
		}
		if val, ok := args["sent_as_draft"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sent_as_draft=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/invoices%s", cfg.BaseURL, queryString)
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

func CreateGet_invoicesTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_invoices",
		mcp.WithDescription("View list of invoices"),
		mcp.WithString("contact_id", mcp.Description("Used to filter on the `contact_id` of the invoices")),
		mcp.WithString("project_id", mcp.Description("Used to filter on the `project_id` of the invoices")),
		mcp.WithString("invoice_number", mcp.Description("Used to filter on the `invoice_number` of the invoices")),
		mcp.WithString("offer_number", mcp.Description("")),
		mcp.WithNumber("is_draft", mcp.Description("")),
		mcp.WithNumber("is_offer", mcp.Description("")),
		mcp.WithNumber("is_locked", mcp.Description("")),
		mcp.WithNumber("is_fixed_price", mcp.Description("")),
		mcp.WithString("date_from", mcp.Description("")),
		mcp.WithString("date_to", mcp.Description("")),
		mcp.WithString("issued_date", mcp.Description("")),
		mcp.WithNumber("sent_as_draft", mcp.Description("Used to filter invoices which are sent as draft to integration")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_invoicesHandler(cfg),
	}
}
