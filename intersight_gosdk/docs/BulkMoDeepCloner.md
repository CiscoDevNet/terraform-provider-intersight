# BulkMoDeepCloner

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

### NewBulkMoDeepCloner

`func NewBulkMoDeepCloner(classId string, objectType string, ) *BulkMoDeepCloner`

NewBulkMoDeepCloner instantiates a new BulkMoDeepCloner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkMoDeepClonerWithDefaults

`func NewBulkMoDeepClonerWithDefaults() *BulkMoDeepCloner`

NewBulkMoDeepClonerWithDefaults instantiates a new BulkMoDeepCloner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BulkMoDeepCloner) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BulkMoDeepCloner) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BulkMoDeepCloner) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BulkMoDeepCloner) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BulkMoDeepCloner) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BulkMoDeepCloner) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetExcludeProperties

`func (o *BulkMoDeepCloner) GetExcludeProperties() []string`

GetExcludeProperties returns the ExcludeProperties field if non-nil, zero value otherwise.

### GetExcludePropertiesOk

`func (o *BulkMoDeepCloner) GetExcludePropertiesOk() (*[]string, bool)`

GetExcludePropertiesOk returns a tuple with the ExcludeProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeProperties

`func (o *BulkMoDeepCloner) SetExcludeProperties(v []string)`

SetExcludeProperties sets ExcludeProperties field to given value.

### HasExcludeProperties

`func (o *BulkMoDeepCloner) HasExcludeProperties() bool`

HasExcludeProperties returns a boolean if a field has been set.

### SetExcludePropertiesNil

`func (o *BulkMoDeepCloner) SetExcludePropertiesNil(b bool)`

 SetExcludePropertiesNil sets the value for ExcludeProperties to be an explicit nil

### UnsetExcludeProperties
`func (o *BulkMoDeepCloner) UnsetExcludeProperties()`

UnsetExcludeProperties ensures that no value is present for ExcludeProperties, not even an explicit nil
### GetSource

`func (o *BulkMoDeepCloner) GetSource() MoMoRef`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *BulkMoDeepCloner) GetSourceOk() (*MoMoRef, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *BulkMoDeepCloner) SetSource(v MoMoRef)`

SetSource sets Source field to given value.

### HasSource

`func (o *BulkMoDeepCloner) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTargets

`func (o *BulkMoDeepCloner) GetTargets() []MoBaseMo`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *BulkMoDeepCloner) GetTargetsOk() (*[]MoBaseMo, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *BulkMoDeepCloner) SetTargets(v []MoBaseMo)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *BulkMoDeepCloner) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### SetTargetsNil

`func (o *BulkMoDeepCloner) SetTargetsNil(b bool)`

 SetTargetsNil sets the value for Targets to be an explicit nil

### UnsetTargets
`func (o *BulkMoDeepCloner) UnsetTargets()`

UnsetTargets ensures that no value is present for Targets, not even an explicit nil
### GetWorkflowNameSuffix

`func (o *BulkMoDeepCloner) GetWorkflowNameSuffix() string`

GetWorkflowNameSuffix returns the WorkflowNameSuffix field if non-nil, zero value otherwise.

### GetWorkflowNameSuffixOk

`func (o *BulkMoDeepCloner) GetWorkflowNameSuffixOk() (*string, bool)`

GetWorkflowNameSuffixOk returns a tuple with the WorkflowNameSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowNameSuffix

`func (o *BulkMoDeepCloner) SetWorkflowNameSuffix(v string)`

SetWorkflowNameSuffix sets WorkflowNameSuffix field to given value.

### HasWorkflowNameSuffix

`func (o *BulkMoDeepCloner) HasWorkflowNameSuffix() bool`

HasWorkflowNameSuffix returns a boolean if a field has been set.

### GetAsyncResult

`func (o *BulkMoDeepCloner) GetAsyncResult() BulkResultRelationship`

GetAsyncResult returns the AsyncResult field if non-nil, zero value otherwise.

### GetAsyncResultOk

`func (o *BulkMoDeepCloner) GetAsyncResultOk() (*BulkResultRelationship, bool)`

GetAsyncResultOk returns a tuple with the AsyncResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsyncResult

`func (o *BulkMoDeepCloner) SetAsyncResult(v BulkResultRelationship)`

SetAsyncResult sets AsyncResult field to given value.

### HasAsyncResult

`func (o *BulkMoDeepCloner) HasAsyncResult() bool`

HasAsyncResult returns a boolean if a field has been set.

### GetOrganization

`func (o *BulkMoDeepCloner) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *BulkMoDeepCloner) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *BulkMoDeepCloner) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *BulkMoDeepCloner) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


