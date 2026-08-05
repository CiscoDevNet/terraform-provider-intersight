# MgmtObjectBackupMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "mgmt.ObjectBackupMeta"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "mgmt.ObjectBackupMeta"]
**ExcludeRelations** | Pointer to **bool** | No relations are to be exported, if specified. | [optional] [readonly] 
**ExcludedPeers** | Pointer to **[]string** |  | [optional] 
**ExportTimeout** | Pointer to **int64** | The time required in milli seconds to export one instance of this MO type with all its relations and values populated. The MO export timeout value is used to determine the timeout for the overall backup operation. | [optional] [readonly] [default to 2000]
**Filter** | Pointer to **string** | Odata filter to find matching instances that need to be backed up. | [optional] [readonly] 
**Name** | Pointer to **string** | The MO type to be exported. | [optional] [readonly] 
**RelatedTypeOptions** | Pointer to [**[]BulkRelatedTypeExportOption**](BulkRelatedTypeExportOption.md) |  | [optional] 
**ReportingService** | Pointer to **string** | The service that have reported this backup meta-data. | [optional] [readonly] 
**RestPath** | Pointer to **string** | The rest path for the MO type. | [optional] [readonly] 
**Category** | Pointer to [**NullableMgmtBackupCategoryRelationship**](MgmtBackupCategoryRelationship.md) |  | [optional] 

## Methods

### NewMgmtObjectBackupMeta

`func NewMgmtObjectBackupMeta(classId string, objectType string, ) *MgmtObjectBackupMeta`

NewMgmtObjectBackupMeta instantiates a new MgmtObjectBackupMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMgmtObjectBackupMetaWithDefaults

`func NewMgmtObjectBackupMetaWithDefaults() *MgmtObjectBackupMeta`

NewMgmtObjectBackupMetaWithDefaults instantiates a new MgmtObjectBackupMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MgmtObjectBackupMeta) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MgmtObjectBackupMeta) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MgmtObjectBackupMeta) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MgmtObjectBackupMeta) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MgmtObjectBackupMeta) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MgmtObjectBackupMeta) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetExcludeRelations

`func (o *MgmtObjectBackupMeta) GetExcludeRelations() bool`

GetExcludeRelations returns the ExcludeRelations field if non-nil, zero value otherwise.

### GetExcludeRelationsOk

`func (o *MgmtObjectBackupMeta) GetExcludeRelationsOk() (*bool, bool)`

GetExcludeRelationsOk returns a tuple with the ExcludeRelations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeRelations

`func (o *MgmtObjectBackupMeta) SetExcludeRelations(v bool)`

SetExcludeRelations sets ExcludeRelations field to given value.

### HasExcludeRelations

`func (o *MgmtObjectBackupMeta) HasExcludeRelations() bool`

HasExcludeRelations returns a boolean if a field has been set.

### GetExcludedPeers

`func (o *MgmtObjectBackupMeta) GetExcludedPeers() []string`

GetExcludedPeers returns the ExcludedPeers field if non-nil, zero value otherwise.

### GetExcludedPeersOk

`func (o *MgmtObjectBackupMeta) GetExcludedPeersOk() (*[]string, bool)`

GetExcludedPeersOk returns a tuple with the ExcludedPeers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedPeers

`func (o *MgmtObjectBackupMeta) SetExcludedPeers(v []string)`

SetExcludedPeers sets ExcludedPeers field to given value.

### HasExcludedPeers

`func (o *MgmtObjectBackupMeta) HasExcludedPeers() bool`

HasExcludedPeers returns a boolean if a field has been set.

### SetExcludedPeersNil

`func (o *MgmtObjectBackupMeta) SetExcludedPeersNil(b bool)`

 SetExcludedPeersNil sets the value for ExcludedPeers to be an explicit nil

### UnsetExcludedPeers
`func (o *MgmtObjectBackupMeta) UnsetExcludedPeers()`

UnsetExcludedPeers ensures that no value is present for ExcludedPeers, not even an explicit nil
### GetExportTimeout

`func (o *MgmtObjectBackupMeta) GetExportTimeout() int64`

GetExportTimeout returns the ExportTimeout field if non-nil, zero value otherwise.

### GetExportTimeoutOk

`func (o *MgmtObjectBackupMeta) GetExportTimeoutOk() (*int64, bool)`

GetExportTimeoutOk returns a tuple with the ExportTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportTimeout

`func (o *MgmtObjectBackupMeta) SetExportTimeout(v int64)`

SetExportTimeout sets ExportTimeout field to given value.

### HasExportTimeout

`func (o *MgmtObjectBackupMeta) HasExportTimeout() bool`

HasExportTimeout returns a boolean if a field has been set.

### GetFilter

`func (o *MgmtObjectBackupMeta) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *MgmtObjectBackupMeta) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *MgmtObjectBackupMeta) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *MgmtObjectBackupMeta) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetName

`func (o *MgmtObjectBackupMeta) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MgmtObjectBackupMeta) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MgmtObjectBackupMeta) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MgmtObjectBackupMeta) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRelatedTypeOptions

`func (o *MgmtObjectBackupMeta) GetRelatedTypeOptions() []BulkRelatedTypeExportOption`

GetRelatedTypeOptions returns the RelatedTypeOptions field if non-nil, zero value otherwise.

### GetRelatedTypeOptionsOk

`func (o *MgmtObjectBackupMeta) GetRelatedTypeOptionsOk() (*[]BulkRelatedTypeExportOption, bool)`

GetRelatedTypeOptionsOk returns a tuple with the RelatedTypeOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedTypeOptions

`func (o *MgmtObjectBackupMeta) SetRelatedTypeOptions(v []BulkRelatedTypeExportOption)`

SetRelatedTypeOptions sets RelatedTypeOptions field to given value.

### HasRelatedTypeOptions

`func (o *MgmtObjectBackupMeta) HasRelatedTypeOptions() bool`

HasRelatedTypeOptions returns a boolean if a field has been set.

### SetRelatedTypeOptionsNil

`func (o *MgmtObjectBackupMeta) SetRelatedTypeOptionsNil(b bool)`

 SetRelatedTypeOptionsNil sets the value for RelatedTypeOptions to be an explicit nil

### UnsetRelatedTypeOptions
`func (o *MgmtObjectBackupMeta) UnsetRelatedTypeOptions()`

UnsetRelatedTypeOptions ensures that no value is present for RelatedTypeOptions, not even an explicit nil
### GetReportingService

`func (o *MgmtObjectBackupMeta) GetReportingService() string`

GetReportingService returns the ReportingService field if non-nil, zero value otherwise.

### GetReportingServiceOk

`func (o *MgmtObjectBackupMeta) GetReportingServiceOk() (*string, bool)`

GetReportingServiceOk returns a tuple with the ReportingService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingService

`func (o *MgmtObjectBackupMeta) SetReportingService(v string)`

SetReportingService sets ReportingService field to given value.

### HasReportingService

`func (o *MgmtObjectBackupMeta) HasReportingService() bool`

HasReportingService returns a boolean if a field has been set.

### GetRestPath

`func (o *MgmtObjectBackupMeta) GetRestPath() string`

GetRestPath returns the RestPath field if non-nil, zero value otherwise.

### GetRestPathOk

`func (o *MgmtObjectBackupMeta) GetRestPathOk() (*string, bool)`

GetRestPathOk returns a tuple with the RestPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestPath

`func (o *MgmtObjectBackupMeta) SetRestPath(v string)`

SetRestPath sets RestPath field to given value.

### HasRestPath

`func (o *MgmtObjectBackupMeta) HasRestPath() bool`

HasRestPath returns a boolean if a field has been set.

### GetCategory

`func (o *MgmtObjectBackupMeta) GetCategory() MgmtBackupCategoryRelationship`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *MgmtObjectBackupMeta) GetCategoryOk() (*MgmtBackupCategoryRelationship, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *MgmtObjectBackupMeta) SetCategory(v MgmtBackupCategoryRelationship)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *MgmtObjectBackupMeta) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *MgmtObjectBackupMeta) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *MgmtObjectBackupMeta) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


