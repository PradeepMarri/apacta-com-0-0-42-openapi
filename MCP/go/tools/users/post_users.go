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

func Post_usersHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
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
		url := fmt.Sprintf("%s/users", cfg.BaseURL)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
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

func CreatePost_usersTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_users",
		mcp.WithDescription("Add user to company"),
		mcp.WithBoolean("receive_form_mails", mcp.Description("Input parameter: If `true` the employee will receive an email receipt of every form submitted")),
		mcp.WithString("email", mcp.Description("")),
		mcp.WithString("street_name", mcp.Description("")),
		mcp.WithString("phone", mcp.Description("")),
		mcp.WithObject("roles", mcp.Description("")),
		mcp.WithString("city_id", mcp.Description("")),
		mcp.WithBoolean("hide_phone", mcp.Description("")),
		mcp.WithString("first_name", mcp.Required(), mcp.Description("")),
		mcp.WithString("initials", mcp.Description("")),
		mcp.WithString("password", mcp.Description("")),
		mcp.WithString("sale_price", mcp.Description("Input parameter: The price this employee costs per hour when working")),
		mcp.WithString("expected_billable_hours", mcp.Description("")),
		mcp.WithString("cost_price", mcp.Description("Input parameter: Cost of salaries")),
		mcp.WithBoolean("is_active", mcp.Description("")),
		mcp.WithString("language_id", mcp.Description("")),
		mcp.WithString("mobile", mcp.Description("")),
		mcp.WithString("phone_countrycode", mcp.Description("")),
		mcp.WithBoolean("hide_address", mcp.Description("")),
		mcp.WithString("mobile_countrycode", mcp.Description("")),
		mcp.WithString("last_name", mcp.Description("")),
		mcp.WithString("extra_price", mcp.Description("Input parameter: Additional cost on this employee (pension, vacation etc.)")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Post_usersHandler(cfg),
	}
}
