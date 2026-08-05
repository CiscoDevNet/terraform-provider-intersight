# MgmtConfigCategorySummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "mgmt.ConfigCategorySummary"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "mgmt.ConfigCategorySummary"]
**CategoryName** | Pointer to **string** | The name of the sub category. The objects are categorized based on the type of the object. For organization - it can be pools, policies, profiles, templates. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the organization or account. | [optional] [readonly] 
**Scope** | Pointer to **string** | The scope of the category, the objects belong to an organization or an account. | [optional] [readonly] 
**SummarySet** | Pointer to [**[]MgmtConfigObjectDetails**](MgmtConfigObjectDetails.md) |  | [optional] 
**BackupInstance** | Pointer to [**NullableMgmtConfigBackupInstanceRelationship**](MgmtConfigBackupInstanceRelationship.md) |  | [optional] 

## Methods

### NewMgmtConfigCategorySummary

`func NewMgmtConfigCategorySummary(classId string, objectType string, ) *MgmtConfigCategorySummary`

NewMgmtConfigCategorySummary instantiates a new MgmtConfigCategorySummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMgmtConfigCategorySummaryWithDefaults

`func NewMgmtConfigCategorySummaryWithDefaults() *MgmtConfigCategorySummary`

NewMgmtConfigCategorySummaryWithDefaults instantiates a new MgmtConfigCategorySummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MgmtConfigCategorySummary) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MgmtConfigCategorySummary) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MgmtConfigCategorySummary) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MgmtConfigCategorySummary) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MgmtConfigCategorySummary) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MgmtConfigCategorySummary) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCategoryName

`func (o *MgmtConfigCategorySummary) GetCategoryName() string`

GetCategoryName returns the CategoryName field if non-nil, zero value otherwise.

### GetCategoryNameOk

`func (o *MgmtConfigCategorySummary) GetCategoryNameOk() (*string, bool)`

GetCategoryNameOk returns a tuple with the CategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryName

`func (o *MgmtConfigCategorySummary) SetCategoryName(v string)`

SetCategoryName sets CategoryName field to given value.

### HasCategoryName

`func (o *MgmtConfigCategorySummary) HasCategoryName() bool`

HasCategoryName returns a boolean if a field has been set.

### GetName

`func (o *MgmtConfigCategorySummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MgmtConfigCategorySummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MgmtConfigCategorySummary) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MgmtConfigCategorySummary) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScope

`func (o *MgmtConfigCategorySummary) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *MgmtConfigCategorySummary) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *MgmtConfigCategorySummary) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *MgmtConfigCategorySummary) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSummarySet

`func (o *MgmtConfigCategorySummary) GetSummarySet() []MgmtConfigObjectDetails`

GetSummarySet returns the SummarySet field if non-nil, zero value otherwise.

### GetSummarySetOk

`func (o *MgmtConfigCategorySummary) GetSummarySetOk() (*[]MgmtConfigObjectDetails, bool)`

GetSummarySetOk returns a tuple with the SummarySet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummarySet

`func (o *MgmtConfigCategorySummary) SetSummarySet(v []MgmtConfigObjectDetails)`

SetSummarySet sets SummarySet field to given value.

### HasSummarySet

`func (o *MgmtConfigCategorySummary) HasSummarySet() bool`

HasSummarySet returns a boolean if a field has been set.

### SetSummarySetNil

`func (o *MgmtConfigCategorySummary) SetSummarySetNil(b bool)`

 SetSummarySetNil sets the value for SummarySet to be an explicit nil

### UnsetSummarySet
`func (o *MgmtConfigCategorySummary) UnsetSummarySet()`

UnsetSummarySet ensures that no value is present for SummarySet, not even an explicit nil
### GetBackupInstance

`func (o *MgmtConfigCategorySummary) GetBackupInstance() MgmtConfigBackupInstanceRelationship`

GetBackupInstance returns the BackupInstance field if non-nil, zero value otherwise.

### GetBackupInstanceOk

`func (o *MgmtConfigCategorySummary) GetBackupInstanceOk() (*MgmtConfigBackupInstanceRelationship, bool)`

GetBackupInstanceOk returns a tuple with the BackupInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupInstance

`func (o *MgmtConfigCategorySummary) SetBackupInstance(v MgmtConfigBackupInstanceRelationship)`

SetBackupInstance sets BackupInstance field to given value.

### HasBackupInstance

`func (o *MgmtConfigCategorySummary) HasBackupInstance() bool`

HasBackupInstance returns a boolean if a field has been set.

### SetBackupInstanceNil

`func (o *MgmtConfigCategorySummary) SetBackupInstanceNil(b bool)`

 SetBackupInstanceNil sets the value for BackupInstance to be an explicit nil

### UnsetBackupInstance
`func (o *MgmtConfigCategorySummary) UnsetBackupInstance()`

UnsetBackupInstance ensures that no value is present for BackupInstance, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


