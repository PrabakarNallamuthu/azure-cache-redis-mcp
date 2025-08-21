package main

import (
	"github.com/redismanagementclient/mcp-server/config"
	"github.com/redismanagementclient/mcp-server/models"
	tools_redis "github.com/redismanagementclient/mcp-server/tools/redis"
	tools_operations "github.com/redismanagementclient/mcp-server/tools/operations"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_redis.CreateFirewallrules_deleteTool(cfg),
		tools_redis.CreateFirewallrules_getTool(cfg),
		tools_redis.CreateFirewallrules_createorupdateTool(cfg),
		tools_redis.CreateRedis_importdataTool(cfg),
		tools_redis.CreateLinkedserver_deleteTool(cfg),
		tools_redis.CreateLinkedserver_getTool(cfg),
		tools_redis.CreateLinkedserver_createTool(cfg),
		tools_redis.CreateRedis_exportdataTool(cfg),
		tools_redis.CreateRedis_listkeysTool(cfg),
		tools_redis.CreatePatchschedules_listbyredisresourceTool(cfg),
		tools_operations.CreateOperations_listTool(cfg),
		tools_redis.CreateRedis_checknameavailabilityTool(cfg),
		tools_redis.CreateRedis_regeneratekeyTool(cfg),
		tools_redis.CreateRedis_forcerebootTool(cfg),
		tools_redis.CreateRedis_listTool(cfg),
		tools_redis.CreateFirewallrules_listbyredisresourceTool(cfg),
		tools_redis.CreateRedis_listupgradenotificationsTool(cfg),
		tools_redis.CreatePatchschedules_deleteTool(cfg),
		tools_redis.CreatePatchschedules_getTool(cfg),
		tools_redis.CreatePatchschedules_createorupdateTool(cfg),
		tools_redis.CreateLinkedserver_listTool(cfg),
		tools_redis.CreateRedis_listbyresourcegroupTool(cfg),
		tools_redis.CreateRedis_getTool(cfg),
		tools_redis.CreateRedis_updateTool(cfg),
		tools_redis.CreateRedis_createTool(cfg),
		tools_redis.CreateRedis_deleteTool(cfg),
	}
}
