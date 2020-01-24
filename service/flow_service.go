package service

import (
	"encoding/json"
	"flow/cache"
	"flow/logger"
	"flow/model"
	"flow/model/response_dto"
)

type FlowService struct {
	FlowServiceUtil FlowServiceUtil
}

func (u FlowService) GetFlows(merchantId string, tenantId string, channelId string) response_dto.FlowResponsesDto {
	methodName := "GetFlows"
	logger.SugarLogger.Info(methodName, "Recieved request to get all the flows associated with merchant: ", merchantId, " tenantId: ", tenantId, " channelId: ", channelId)
	redisClient := cache.GetRedisClient()
	var flowsResponse response_dto.FlowResponsesDto
	if redisClient == nil {
		logger.SugarLogger.Info(methodName, "Failed to connect with redis client. ")
		return flowsResponse
	}
	logger.SugarLogger.Info(methodName, "Fetching flows from redis cache for merchant: ", merchantId, " tenantId: ", tenantId, " channelId: ", channelId)
	redisKey := model.RedisKey{MerchantId: merchantId,
		TenantId:  tenantId,
		ChannelId: channelId}
	cachedFlow, err := redisClient.Get(redisKey.ToString()).Result()
	if err != nil {
		logger.SugarLogger.Info(methodName, "Failed to fetch flow from redis cache for merchant: ", merchantId, " tenantId: ", tenantId, " channelId: ", channelId, " with error: ", err)
	}
	if cachedFlow == "" {
		logger.SugarLogger.Info(methodName, "No flows exist in redis cache for merchant: ", merchantId, " tenantId: ", tenantId, " channelId: ", channelId)
		flowContext := model.FlowContext{
			MerchantId: merchantId,
			TenantId:   tenantId,
			ChannelId:  channelId}
		flows := u.FlowServiceUtil.FetchAllFlowsFromDB(flowContext)
		flowsResponse, err := u.FlowServiceUtil.GetParsedFlowsResponse(flows)
		if err != nil {
			logger.SugarLogger.Error(methodName, "Failed to fetch parsed flows associated with merchant: ", merchantId, " tenantId: ", tenantId, " channelId: ", channelId, " with error: ", err)
			return flowsResponse
		}
		//Do not set redis key when there is no entry for given flowContext.
		if len(flowsResponse.FlowResponses) == 0 {
			return flowsResponse
		}

		response, err := json.Marshal(flowsResponse)
		if err != nil{
			logger.SugarLogger.Error(methodName, " couldn't update redis as failed to marshal response with err: ", err)
			return flowsResponse
		}
		
		logger.SugarLogger.Info(methodName, " Adding redis key: ",redisKey.ToString())
		setStatus := redisClient.Set(redisKey.ToString(), response, 0)
		logger.SugarLogger.Info(methodName, " Set redis key status: ", setStatus.Val(), " for key: ",redisKey.ToString())
		return flowsResponse
	}
	logger.SugarLogger.Info(methodName, " UnMarshlling the cached flow response")
	json.Unmarshal([]byte(cachedFlow), &flowsResponse)
	return flowsResponse
}
