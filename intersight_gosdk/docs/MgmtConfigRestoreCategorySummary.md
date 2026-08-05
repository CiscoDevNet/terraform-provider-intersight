# MgmtConfigRestoreCategorySummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "mgmt.ConfigRestoreCategorySummary"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "mgmt.ConfigRestoreCategorySummary"]
**CategoryName** | Pointer to **string** | The name of the sub category. The objects are categorized based on the type of the object. For organization - it can be pools, policies, profiles, templates. | [optional] 
**Name** | Pointer to **string** | The name of the organization or account. | [optional] 
**Scope** | Pointer to **string** | The scope of the category, the objects belong to an organization or an account. | [optional] 
**SummarySet** | Pointer to [**[]MgmtConfigObjectRestoreDetails**](MgmtConfigObjectRestoreDetails.md) |  | [optional] 
**RestoreOperation** | Pointer to [**NullableMgmtConfigRestoreOperationRelationship**](MgmtConfigRestoreOperationRelationship.md) |  | [optional] 

## Methods

### NewMgmtConfigRestoreCategorySummary

`func NewMgmtConfigRestoreCategorySummary(classId string, objectType string, ) *MgmtConfigRestoreCategorySummary`

NewMgmtConfigRestoreCategorySummary instantiates a new MgmtConfigRestoreCategorySummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMgmtConfigRestoreCategorySummaryWithDefaults

`func NewMgmtConfigRestoreCategorySummaryWithDefaults() *MgmtConfigRestoreCategorySummary`

NewMgmtConfigRestoreCategorySummaryWithDefaults instantiates a new MgmtConfigRestoreCategorySummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MgmtConfigRestoreCategorySummary) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MgmtConfigRestoreCategorySummary) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MgmtConfigRestoreCategorySummary) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MgmtConfigRestoreCategorySummary) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MgmtConfigRestoreCategorySummary) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MgmtConfigRestoreCategorySummary) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCategoryName

`func (o *MgmtConfigRestoreCategorySummary) GetCategoryName() string`

GetCategoryName returns the CategoryName field if non-nil, zero value otherwise.

### GetCategoryNameOk

`func (o *MgmtConfigRestoreCategorySummary) GetCategoryNameOk() (*string, bool)`

GetCategoryNameOk returns a tuple with the CategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryName

`func (o *MgmtConfigRestoreCategorySummary) SetCategoryName(v string)`

SetCategoryName sets CategoryName field to given value.

### HasCategoryName

`func (o *MgmtConfigRestoreCategorySummary) HasCategoryName() bool`

HasCategoryName returns a boolean if a field has been set.

### GetName

`func (o *MgmtConfigRestoreCategorySummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MgmtConfigRestoreCategorySummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MgmtConfigRestoreCategorySummary) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MgmtConfigRestoreCategorySummary) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScope

`func (o *MgmtConfigRestoreCategorySummary) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *MgmtConfigRestoreCategorySummary) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *MgmtConfigRestoreCategorySummary) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *MgmtConfigRestoreCategorySummary) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSummarySet

`func (o *MgmtConfigRestoreCategorySummary) GetSummarySet() []MgmtConfigObjectRestoreDetails`

GetSummarySet returns the SummarySet field if non-nil, zero value otherwise.

### GetSummarySetOk

`func (o *MgmtConfigRestoreCategorySummary) GetSummarySetOk() (*[]MgmtConfigObjectRestoreDetails, bool)`

GetSummarySetOk returns a tuple with the SummarySet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummarySet

`func (o *MgmtConfigRestoreCategorySummary) SetSummarySet(v []MgmtConfigObjectRestoreDetails)`

SetSummarySet sets SummarySet field to given value.

### HasSummarySet

`func (o *MgmtConfigRestoreCategorySummary) HasSummarySet() bool`

HasSummarySet returns a boolean if a field has been set.

### SetSummarySetNil

`func (o *MgmtConfigRestoreCategorySummary) SetSummarySetNil(b bool)`

 SetSummarySetNil sets the value for SummarySet to be an explicit nil

### UnsetSummarySet
`func (o *MgmtConfigRestoreCategorySummary) UnsetSummarySet()`

UnsetSummarySet ensures that no value is present for SummarySet, not even an explicit nil
### GetRestoreOperation

`func (o *MgmtConfigRestoreCategorySummary) GetRestoreOperation() MgmtConfigRestoreOperationRelationship`

GetRestoreOperation returns the RestoreOperation field if non-nil, zero value otherwise.

### GetRestoreOperationOk

`func (o *MgmtConfigRestoreCategorySummary) GetRestoreOperationOk() (*MgmtConfigRestoreOperationRelationship, bool)`

GetRestoreOperationOk returns a tuple with the RestoreOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreOperation

`func (o *MgmtConfigRestoreCategorySummary) SetRestoreOperation(v MgmtConfigRestoreOperationRelationship)`

SetRestoreOperation sets RestoreOperation field to given value.

### HasRestoreOperation

`func (o *MgmtConfigRestoreCategorySummary) HasRestoreOperation() bool`

HasRestoreOperation returns a boolean if a field has been set.

### SetRestoreOperationNil

`func (o *MgmtConfigRestoreCategorySummary) SetRestoreOperationNil(b bool)`

 SetRestoreOperationNil sets the value for RestoreOperation to be an explicit nil

### UnsetRestoreOperation
`func (o *MgmtConfigRestoreCategorySummary) UnsetRestoreOperation()`

UnsetRestoreOperation ensures that no value is present for RestoreOperation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


