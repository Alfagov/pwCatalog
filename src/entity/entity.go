package entity

import (
	"github.com/Alfagov/purview-go-sdk/src/client"
	"github.com/Alfagov/purview-go-sdk/src/models"
)

type Entity interface {
	AddClassification(
		classification *models.AtlasClassification,
		entityGuids []string,
	) error
	AddClassifications(
		guid string,
		classifications []*models.AtlasClassification,
	) error
	AddClassificationsByUniqueAttribute(
		classifications []*models.AtlasClassification,
		attributes *models.QueryAttributes,
	) error
	AddLabel(guid string, label []string) error
	AddLabelsByUniqueAttribute(
		labels []string, attributes *models.QueryAttributes,
	) error
	AddOrUpdateBusinessMetadata(
		guid string, businessMetadata map[string]interface{},
		options *models.QueryParams,
	) error
	AddOrUpdateBusinessMetadataAttributes(
		businessMetadataAttributes map[string]interface{},
		guid string, bmName string,
	) error
	CreateOrUpdate(
		data *models.EntityCreateOrUpdateRequest,
		options *models.QueryParams,
		customUnmarshaler func(data []byte) (
			*models.EntityMutationResponse,
			error,
		),
	) (*models.EntityMutationResponse, error)
	CreateOrUpdateEntities(
		data *models.EntityCreateOrUpdateBulkRequest,
		options *models.QueryParams,
		customUnmarshaler func(data []byte) (
			*models.EntityMutationResponse,
			error,
		),
	) (*models.EntityMutationResponse, error)
	GetByGuid(
		guid string,
		customUnmarshaler func(data []byte) (
			*models.AtlasEntityWithExtInfo, error,
		),
	) (*models.AtlasEntityWithExtInfo, error)
	GetByUniqueAttributes(
		uniqueAttributes *models.QueryAttributes,
		params *models.QueryParams,
		customUnmarshaler func(data []byte) (
			*models.AtlasEntityWithExtInfo, error,
		),
	) (*models.AtlasEntityWithExtInfo, error)
	DeleteBusinessMetadata(
		guid string,
		options *models.QueryParams,
		businessMetadata map[string]interface{},
	) error
	DeleteBusinessMetadataAttributes(
		guid string, bmName string,
		businessMetadata map[string]interface{},
	) error
	DeleteByGuid(guid string) error
	DeleteByGuids(guids []string) error
	DeleteByUniqueAttribute(uniqueAttributes *models.QueryAttributes) error
	DeleteClassification(guid string, classificationName string) error
	DeleteClassificationByUniqueAttribute(
		name string, uniqueAttributes *models.QueryAttributes,
	) error
	DeleteLabels(guid string, labels []string) error
	DeleteLabelsByUniqueAttribute(
		uniqueAttributes *models.QueryAttributes,
		labels []string,
	) error
	GetClassification(guid string, name string) (
		*models.AtlasClassification,
		error,
	)
	GetClassifications(guid string) (*models.AtlasClassifications, error)
	GetEntitiesByUniqueAttributes(uniqueAttributes map[string]string) (
		*models.AtlasEntityWithExtInfo, error,
	)
	GetHeader(guid string) (*models.AtlasEntityHeader, error)
	GetSampleBusinessMetadataTemplate(templateName string) (
		interface{}, error,
	)
	ImportBusinessMetadata() (*models.BulkImportResponse, error)
	ListByGuids(guids []string, options *models.QueryParams) (
		*models.
			AtlasEntityWithExtInfo, error,
	)
	PartialUpdateEntityAttributeByGuid(
		guid string, name string, value interface{},
		customUnmarshaler func(data []byte) (
			*models.EntityMutationResponse, error,
		),
	) (*models.EntityMutationResponse, error)
	PartialUpdateEntityByUniqueAttributes(
		uniqueAttributes map[string]string,
		entityData map[string]interface{},
		customUnmarshaler func(data []byte) (
			*models.AtlasEntityWithExtInfo, error,
		),
	) (*models.AtlasEntityWithExtInfo, error)
	SetClassifications(
		guid string, classifications map[string]models.AtlasEntityHeader,
	) ([]string, error)
	SetLabels(guid string, labels []string) error
	SetLabelsByUniqueAttribute(
		uniqueAttributes map[string]string, labels []string,
	) error
	UpdateClassifications(
		guid string, classifications []*models.AtlasClassification,
	) error
	UpdateClassificationsByUniqueAttribute(
		classifications []*models.AtlasClassification,
		uniqueAttributes map[string]string,
	) error
}

type entityImpl struct {
	client.HttpClient
}

func NewEntity(httpClient client.HttpClient) Entity {
	return &entityImpl{httpClient}
}
