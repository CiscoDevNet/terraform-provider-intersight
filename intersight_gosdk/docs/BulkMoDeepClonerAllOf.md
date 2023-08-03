# BulkMoDeepClonerAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "bulk.MoDeepCloner"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "bulk.MoDeepCloner"]
**ExcludeProperties** | Pointer to **[]string** |  | [optional] 
**Source** | Pointer to [**MoMoRef**](MoMoRef.md) |  | [optional] 
**Targets** | Pointer to [**[]MoBaseMo**](MoBaseMo.md) |  | [optional] 
**WorkflowNameSuffix** | Pointer to **string** | A user friendly short name to identify the workflow, optionally. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ), forward slash (/), comma or an underscore (_). | [optional] 
**AsyncResult** | Pointer to [**BulkResultRelationship**](BulkResultRelationship.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewBulkMoDeepClonerAllOf

`func NewBulkMoDeepClonerAllOf(classId string, objectType string, ) *BulkMoDeepClonerAllOf`

NewBulkMoDeepClonerAllOf instantiates a new BulkMoDeepClonerAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkMoDeepClonerAllOfWithDefaults

`func NewBulkMoDeepClonerAllOfWithDefaults() *BulkMoDeepClonerAllOf`

NewBulkMoDeepClonerAllOfWithDefaults instantiates a new BulkMoDeepClonerAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BulkMoDeepClonerAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BulkMoDeepClonerAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BulkMoDeepClonerAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BulkMoDeepClonerAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BulkMoDeepClonerAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BulkMoDeepClonerAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetExcludeProperties

`func (o *BulkMoDeepClonerAllOf) GetExcludeProperties() []string`

GetExcludeProperties returns the ExcludeProperties field if non-nil, zero value otherwise.

### GetExcludePropertiesOk

`func (o *BulkMoDeepClonerAllOf) GetExcludePropertiesOk() (*[]string, bool)`

GetExcludePropertiesOk returns a tuple with the ExcludeProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeProperties

`func (o *BulkMoDeepClonerAllOf) SetExcludeProperties(v []string)`

SetExcludeProperties sets ExcludeProperties field to given value.

### HasExcludeProperties

`func (o *BulkMoDeepClonerAllOf) HasExcludeProperties() bool`

HasExcludeProperties returns a boolean if a field has been set.

### SetExcludePropertiesNil

`func (o *BulkMoDeepClonerAllOf) SetExcludePropertiesNil(b bool)`

 SetExcludePropertiesNil sets the value for ExcludeProperties to be an explicit nil

### UnsetExcludeProperties
`func (o *BulkMoDeepClonerAllOf) UnsetExcludeProperties()`

UnsetExcludeProperties ensures that no value is present for ExcludeProperties, not even an explicit nil
### GetSource

`func (o *BulkMoDeepClonerAllOf) GetSource() MoMoRef`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *BulkMoDeepClonerAllOf) GetSourceOk() (*MoMoRef, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *BulkMoDeepClonerAllOf) SetSource(v MoMoRef)`

SetSource sets Source field to given value.

### HasSource

`func (o *BulkMoDeepClonerAllOf) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTargets

`func (o *BulkMoDeepClonerAllOf) GetTargets() []MoBaseMo`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *BulkMoDeepClonerAllOf) GetTargetsOk() (*[]MoBaseMo, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *BulkMoDeepClonerAllOf) SetTargets(v []MoBaseMo)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *BulkMoDeepClonerAllOf) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### SetTargetsNil

`func (o *BulkMoDeepClonerAllOf) SetTargetsNil(b bool)`

 SetTargetsNil sets the value for Targets to be an explicit nil

### UnsetTargets
`func (o *BulkMoDeepClonerAllOf) UnsetTargets()`

UnsetTargets ensures that no value is present for Targets, not even an explicit nil
### GetWorkflowNameSuffix

`func (o *BulkMoDeepClonerAllOf) GetWorkflowNameSuffix() string`

GetWorkflowNameSuffix returns the WorkflowNameSuffix field if non-nil, zero value otherwise.

### GetWorkflowNameSuffixOk

`func (o *BulkMoDeepClonerAllOf) GetWorkflowNameSuffixOk() (*string, bool)`

GetWorkflowNameSuffixOk returns a tuple with the WorkflowNameSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowNameSuffix

`func (o *BulkMoDeepClonerAllOf) SetWorkflowNameSuffix(v string)`

SetWorkflowNameSuffix sets WorkflowNameSuffix field to given value.

### HasWorkflowNameSuffix

`func (o *BulkMoDeepClonerAllOf) HasWorkflowNameSuffix() bool`

HasWorkflowNameSuffix returns a boolean if a field has been set.

### GetAsyncResult

`func (o *BulkMoDeepClonerAllOf) GetAsyncResult() BulkResultRelationship`

GetAsyncResult returns the AsyncResult field if non-nil, zero value otherwise.

### GetAsyncResultOk

`func (o *BulkMoDeepClonerAllOf) GetAsyncResultOk() (*BulkResultRelationship, bool)`

GetAsyncResultOk returns a tuple with the AsyncResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsyncResult

`func (o *BulkMoDeepClonerAllOf) SetAsyncResult(v BulkResultRelationship)`

SetAsyncResult sets AsyncResult field to given value.

### HasAsyncResult

`func (o *BulkMoDeepClonerAllOf) HasAsyncResult() bool`

HasAsyncResult returns a boolean if a field has been set.

### GetOrganization

`func (o *BulkMoDeepClonerAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *BulkMoDeepClonerAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *BulkMoDeepClonerAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *BulkMoDeepClonerAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


