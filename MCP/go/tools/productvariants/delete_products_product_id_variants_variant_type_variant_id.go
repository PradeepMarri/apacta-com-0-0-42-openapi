package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/apacta/mcp-server/config"
	"github.com/apacta/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Delete_products_product_id_variants_variant_type_variant_idHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		product_idVal, ok := args["product_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: product_id"), nil
		}
		product_id, ok := product_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: product_id"), nil
		}
		variant_typeVal, ok := args["variant_type"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: variant_type"), nil
		}
		variant_type, ok := variant_typeVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: variant_type"), nil
		}
		variant_idVal, ok := args["variant_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: variant_id"), nil
		}
		variant_id, ok := variant_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: variant_id"), nil
		}
		url := fmt.Sprintf("%s/products/%s/variants/%s/%s", cfg.BaseURL, product_id, variant_type, variant_id)
		req, err := http.NewRequest("DELETE", url, nil)
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
		var result models.EmptySuccessResponse
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

func CreateDelete_products_product_id_variants_variant_type_variant_idTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_products_product_id_variants_variant_type_variant_id",
		mcp.WithDescription("Delete a product variant"),
		mcp.WithString("product_id", mcp.Required(), mcp.Description("")),
		mcp.WithString("variant_type", mcp.Required(), mcp.Description("")),
		mcp.WithString("variant_id", mcp.Required(), mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Delete_products_product_id_variants_variant_type_variant_idHandler(cfg),
	}
}
