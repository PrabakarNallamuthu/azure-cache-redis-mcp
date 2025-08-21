package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/redismanagementclient/mcp-server/config"
	"github.com/redismanagementclient/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Firewallrules_createorupdateHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		resourceGroupNameVal, ok := args["resourceGroupName"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: resourceGroupName"), nil
		}
		resourceGroupName, ok := resourceGroupNameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: resourceGroupName"), nil
		}
		cacheNameVal, ok := args["cacheName"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: cacheName"), nil
		}
		cacheName, ok := cacheNameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: cacheName"), nil
		}
		ruleNameVal, ok := args["ruleName"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: ruleName"), nil
		}
		ruleName, ok := ruleNameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: ruleName"), nil
		}
		subscriptionIdVal, ok := args["subscriptionId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: subscriptionId"), nil
		}
		subscriptionId, ok := subscriptionIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: subscriptionId"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["api-version"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("api-version=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Cache/Redis/%s/firewallRules/%s%s", cfg.BaseURL, resourceGroupName, cacheName, ruleName, subscriptionId, queryString)
		req, err := http.NewRequest("PUT", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// No authentication required for this endpoint
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

func CreateFirewallrules_createorupdateTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("put_subscriptions_subscriptionId_resourceGroups_resourceGroupName_providers_Microsoft.Cache_Redis_cacheName_firewallRules_ruleName",
		mcp.WithDescription("Create or update a redis cache firewall rule"),
		mcp.WithString("resourceGroupName", mcp.Required(), mcp.Description("The name of the resource group.")),
		mcp.WithString("cacheName", mcp.Required(), mcp.Description("The name of the Redis cache.")),
		mcp.WithString("ruleName", mcp.Required(), mcp.Description("The name of the firewall rule.")),
		mcp.WithString("parameters", mcp.Required(), mcp.Description("Parameters supplied to the create or update redis firewall rule operation.")),
		mcp.WithString("api-version", mcp.Required(), mcp.Description("Client Api Version.")),
		mcp.WithString("subscriptionId", mcp.Required(), mcp.Description("Gets subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription ID forms part of the URI for every service call.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Firewallrules_createorupdateHandler(cfg),
	}
}
