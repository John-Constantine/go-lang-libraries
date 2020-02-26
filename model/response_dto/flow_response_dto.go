package response_dto

import (
	uuid "github.com/google/uuid"
)

type JourneyResponseDto struct {
	Name       string                     `json:"name"`
	ExternalId uuid.UUID                  `json:"externalId"`
	Version    string                     `json:"version"`
	Type       string                     `json:"type"`
	Modules    []ModuleVersionResponseDto `json:"modules"`
}

type JourneyResponseDtoList struct {
	Name       string        `json:"name"`
	ExternalId uuid.UUID     `json:"externalId"`
	Version    string        `json:"version"`
	Type       string        `json:"type"`
	Modules    []ModuleVersionResponseWithExecutionStrategyDTO `json:"modules"`
	Sections   []ResponseDTO `json:"sections"`
	Fields     []ResponseDTO `json:"fields"`
}
