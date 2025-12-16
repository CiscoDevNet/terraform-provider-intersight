# PoolIdMappingPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "pool.IdMappingPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "pool.IdMappingPolicy"]
**PerTypeCombinedSelector** | Pointer to [**[]ResourcePerTypeCombinedSelector**](ResourcePerTypeCombinedSelector.md) |  | [optional] 
**Selectors** | Pointer to [**[]ResourceSelector**](ResourceSelector.md) |  | [optional] 
**UsageCount** | Pointer to **int64** | The number of ID pools to which this ID mapping policy is attached. | [optional] [readonly] 
**Organization** | Pointer to [**NullableOrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewPoolIdMappingPolicy

`func NewPoolIdMappingPolicy(classId string, objectType string, ) *PoolIdMappingPolicy`

NewPoolIdMappingPolicy instantiates a new PoolIdMappingPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolIdMappingPolicyWithDefaults

`func NewPoolIdMappingPolicyWithDefaults() *PoolIdMappingPolicy`

NewPoolIdMappingPolicyWithDefaults instantiates a new PoolIdMappingPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PoolIdMappingPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PoolIdMappingPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PoolIdMappingPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PoolIdMappingPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PoolIdMappingPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PoolIdMappingPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPerTypeCombinedSelector

`func (o *PoolIdMappingPolicy) GetPerTypeCombinedSelector() []ResourcePerTypeCombinedSelector`

GetPerTypeCombinedSelector returns the PerTypeCombinedSelector field if non-nil, zero value otherwise.

### GetPerTypeCombinedSelectorOk

`func (o *PoolIdMappingPolicy) GetPerTypeCombinedSelectorOk() (*[]ResourcePerTypeCombinedSelector, bool)`

GetPerTypeCombinedSelectorOk returns a tuple with the PerTypeCombinedSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerTypeCombinedSelector

`func (o *PoolIdMappingPolicy) SetPerTypeCombinedSelector(v []ResourcePerTypeCombinedSelector)`

SetPerTypeCombinedSelector sets PerTypeCombinedSelector field to given value.

### HasPerTypeCombinedSelector

`func (o *PoolIdMappingPolicy) HasPerTypeCombinedSelector() bool`

HasPerTypeCombinedSelector returns a boolean if a field has been set.

### SetPerTypeCombinedSelectorNil

`func (o *PoolIdMappingPolicy) SetPerTypeCombinedSelectorNil(b bool)`

 SetPerTypeCombinedSelectorNil sets the value for PerTypeCombinedSelector to be an explicit nil

### UnsetPerTypeCombinedSelector
`func (o *PoolIdMappingPolicy) UnsetPerTypeCombinedSelector()`

UnsetPerTypeCombinedSelector ensures that no value is present for PerTypeCombinedSelector, not even an explicit nil
### GetSelectors

`func (o *PoolIdMappingPolicy) GetSelectors() []ResourceSelector`

GetSelectors returns the Selectors field if non-nil, zero value otherwise.

### GetSelectorsOk

`func (o *PoolIdMappingPolicy) GetSelectorsOk() (*[]ResourceSelector, bool)`

GetSelectorsOk returns a tuple with the Selectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectors

`func (o *PoolIdMappingPolicy) SetSelectors(v []ResourceSelector)`

SetSelectors sets Selectors field to given value.

### HasSelectors

`func (o *PoolIdMappingPolicy) HasSelectors() bool`

HasSelectors returns a boolean if a field has been set.

### SetSelectorsNil

`func (o *PoolIdMappingPolicy) SetSelectorsNil(b bool)`

 SetSelectorsNil sets the value for Selectors to be an explicit nil

### UnsetSelectors
`func (o *PoolIdMappingPolicy) UnsetSelectors()`

UnsetSelectors ensures that no value is present for Selectors, not even an explicit nil
### GetUsageCount

`func (o *PoolIdMappingPolicy) GetUsageCount() int64`

GetUsageCount returns the UsageCount field if non-nil, zero value otherwise.

### GetUsageCountOk

`func (o *PoolIdMappingPolicy) GetUsageCountOk() (*int64, bool)`

GetUsageCountOk returns a tuple with the UsageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCount

`func (o *PoolIdMappingPolicy) SetUsageCount(v int64)`

SetUsageCount sets UsageCount field to given value.

### HasUsageCount

`func (o *PoolIdMappingPolicy) HasUsageCount() bool`

HasUsageCount returns a boolean if a field has been set.

### GetOrganization

`func (o *PoolIdMappingPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *PoolIdMappingPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *PoolIdMappingPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *PoolIdMappingPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *PoolIdMappingPolicy) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *PoolIdMappingPolicy) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


