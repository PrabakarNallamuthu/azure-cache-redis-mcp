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

func Linkedserver_createHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		nameVal, ok := args["name"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: name"), nil
		}
		name, ok := nameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: name"), nil
		}
		linkedServerNameVal, ok := args["linkedServerName"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: linkedServerName"), nil
		}
		linkedServerName, ok := linkedServerNameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: linkedServerName"), nil
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
		url := fmt.Sprintf("%s/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Cache/Redis/%s/linkedServers/%s%s", cfg.BaseURL, resourceGroupName, name, linkedServerName, subscriptionId, queryString)
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

func CreateLinkedserver_createTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("put_subscriptions_subscriptionId_resourceGroups_resourceGroupName_providers_Microsoft.Cache_Redis_name_linkedServers_linkedServerName",
		mcp.WithDescription("Adds a linked server to the Redis cache (requires Premium SKU)."),
		mcp.WithString("resourceGroupName", mcp.Required(), mcp.Description("The name of the resource group.")),
		mcp.WithString("name", mcp.Required(), mcp.Description("The name of the Redis cache.")),
		mcp.WithString("linkedServerName", mcp.Required(), mcp.Description("The name of the linked server that is being added to the Redis cache.")),
		mcp.WithString("parameters", mcp.Required(), mcp.Description("Parameters supplied to the Create Linked server operation.")),
		mcp.WithString("api-version", mcp.Required(), mcp.Description("Client Api Version.")),
		mcp.WithString("subscriptionId", mcp.Required(), mcp.Description("Gets subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription ID forms part of the URI for every service call.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Linkedserver_createHandler(cfg),
	}
}
