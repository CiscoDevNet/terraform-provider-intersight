# BulkMoCloner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "bulk.MoCloner"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "bulk.MoCloner"]
**ExcludeProperties** | Pointer to **[]string** |  | [optional] 
**Responses** | Pointer to [**[]BulkRestResult**](BulkRestResult.md) |  | [optional] 
**Sources** | Pointer to [**[]MoBaseMo**](MoBaseMo.md) |  | [optional] 
**Targets** | Pointer to [**[]MoBaseMo**](MoBaseMo.md) |  | [optional] 
**WorkflowNameSuffix** | Pointer to **string** | A user-friendly short name to identify the workflow. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ), forward slash (/), comma or an underscore (_). | [optional] 
**AsyncResult** | Pointer to [**BulkResultRelationship**](BulkResultRelationship.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewBulkMoCloner

`func NewBulkMoCloner(classId string, objectType string, ) *BulkMoCloner`

NewBulkMoCloner instantiates a new BulkMoCloner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkMoClonerWithDefaults

`func NewBulkMoClonerWithDefaults() *BulkMoCloner`

NewBulkMoClonerWithDefaults instantiates a new BulkMoCloner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BulkMoCloner) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BulkMoCloner) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BulkMoCloner) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BulkMoCloner) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BulkMoCloner) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BulkMoCloner) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetExcludeProperties

`func (o *BulkMoCloner) GetExcludeProperties() []string`

GetExcludeProperties returns the ExcludeProperties field if non-nil, zero value otherwise.

### GetExcludePropertiesOk

`func (o *BulkMoCloner) GetExcludePropertiesOk() (*[]string, bool)`

GetExcludePropertiesOk returns a tuple with the ExcludeProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeProperties

`func (o *BulkMoCloner) SetExcludeProperties(v []string)`

SetExcludeProperties sets ExcludeProperties field to given value.

### HasExcludeProperties

`func (o *BulkMoCloner) HasExcludeProperties() bool`

HasExcludeProperties returns a boolean if a field has been set.

### SetExcludePropertiesNil

`func (o *BulkMoCloner) SetExcludePropertiesNil(b bool)`

 SetExcludePropertiesNil sets the value for ExcludeProperties to be an explicit nil

### UnsetExcludeProperties
`func (o *BulkMoCloner) UnsetExcludeProperties()`

UnsetExcludeProperties ensures that no value is present for ExcludeProperties, not even an explicit nil
### GetResponses

`func (o *BulkMoCloner) GetResponses() []BulkRestResult`

GetResponses returns the Responses field if non-nil, zero value otherwise.

### GetResponsesOk

`func (o *BulkMoCloner) GetResponsesOk() (*[]BulkRestResult, bool)`

GetResponsesOk returns a tuple with the Responses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponses

`func (o *BulkMoCloner) SetResponses(v []BulkRestResult)`

SetResponses sets Responses field to given value.

### HasResponses

`func (o *BulkMoCloner) HasResponses() bool`

HasResponses returns a boolean if a field has been set.

### SetResponsesNil

`func (o *BulkMoCloner) SetResponsesNil(b bool)`

 SetResponsesNil sets the value for Responses to be an explicit nil

### UnsetResponses
`func (o *BulkMoCloner) UnsetResponses()`

UnsetResponses ensures that no value is present for Responses, not even an explicit nil
### GetSources

`func (o *BulkMoCloner) GetSources() []MoBaseMo`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *BulkMoCloner) GetSourcesOk() (*[]MoBaseMo, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *BulkMoCloner) SetSources(v []MoBaseMo)`

SetSources sets Sources field to given value.

### HasSources

`func (o *BulkMoCloner) HasSources() bool`

HasSources returns a boolean if a field has been set.

### SetSourcesNil

`func (o *BulkMoCloner) SetSourcesNil(b bool)`

 SetSourcesNil sets the value for Sources to be an explicit nil

### UnsetSources
`func (o *BulkMoCloner) UnsetSources()`

UnsetSources ensures that no value is present for Sources, not even an explicit nil
### GetTargets

`func (o *BulkMoCloner) GetTargets() []MoBaseMo`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *BulkMoCloner) GetTargetsOk() (*[]MoBaseMo, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *BulkMoCloner) SetTargets(v []MoBaseMo)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *BulkMoCloner) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### SetTargetsNil

`func (o *BulkMoCloner) SetTargetsNil(b bool)`

 SetTargetsNil sets the value for Targets to be an explicit nil

### UnsetTargets
`func (o *BulkMoCloner) UnsetTargets()`

UnsetTargets ensures that no value is present for Targets, not even an explicit nil
### GetWorkflowNameSuffix

`func (o *BulkMoCloner) GetWorkflowNameSuffix() string`

GetWorkflowNameSuffix returns the WorkflowNameSuffix field if non-nil, zero value otherwise.

### GetWorkflowNameSuffixOk

`func (o *BulkMoCloner) GetWorkflowNameSuffixOk() (*string, bool)`

GetWorkflowNameSuffixOk returns a tuple with the WorkflowNameSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowNameSuffix

`func (o *BulkMoCloner) SetWorkflowNameSuffix(v string)`

SetWorkflowNameSuffix sets WorkflowNameSuffix field to given value.

### HasWorkflowNameSuffix

`func (o *BulkMoCloner) HasWorkflowNameSuffix() bool`

HasWorkflowNameSuffix returns a boolean if a field has been set.

### GetAsyncResult

`func (o *BulkMoCloner) GetAsyncResult() BulkResultRelationship`

GetAsyncResult returns the AsyncResult field if non-nil, zero value otherwise.

### GetAsyncResultOk

`func (o *BulkMoCloner) GetAsyncResultOk() (*BulkResultRelationship, bool)`

GetAsyncResultOk returns a tuple with the AsyncResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsyncResult

`func (o *BulkMoCloner) SetAsyncResult(v BulkResultRelationship)`

SetAsyncResult sets AsyncResult field to given value.

### HasAsyncResult

`func (o *BulkMoCloner) HasAsyncResult() bool`

HasAsyncResult returns a boolean if a field has been set.

### GetOrganization

`func (o *BulkMoCloner) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *BulkMoCloner) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *BulkMoCloner) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *BulkMoCloner) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


