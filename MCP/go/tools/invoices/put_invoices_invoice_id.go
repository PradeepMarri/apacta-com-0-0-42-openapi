package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/apacta/mcp-server/config"
	"github.com/apacta/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Put_invoices_invoice_idHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		// Create properly typed request body using the generated schema
		var requestBody map[string]interface{}
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/invoices/%s", cfg.BaseURL, invoice_id)
		req, err := http.NewRequest("PUT", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
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

func CreatePut_invoices_invoice_idTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("put_invoices_invoice_id",
		mcp.WithDescription("Edit invoice"),
		mcp.WithString("invoice_id", mcp.Required(), mcp.Description("")),
		mcp.WithString("issued_date", mcp.Description("")),
		mcp.WithString("date_from", mcp.Description("")),
		mcp.WithString("payment_due_date", mcp.Description("")),
		mcp.WithString("erp_payment_term_id", mcp.Description("")),
		mcp.WithString("contact_id", mcp.Description("")),
		mcp.WithNumber("offer_number", mcp.Description("")),
		mcp.WithString("erp_id", mcp.Description("")),
		mcp.WithString("payment_term_id", mcp.Description("")),
		mcp.WithBoolean("is_offer", mcp.Description("")),
		mcp.WithString("reference", mcp.Description("")),
		mcp.WithNumber("invoice_number", mcp.Description("")),
		mcp.WithNumber("vat_percent", mcp.Description("")),
		mcp.WithString("project_id", mcp.Description("")),
		mcp.WithString("date_to", mcp.Description("")),
		mcp.WithBoolean("is_draft", mcp.Description("")),
		mcp.WithBoolean("is_locked", mcp.Description("")),
		mcp.WithString("message", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Put_invoices_invoice_idHandler(cfg),
	}
}
