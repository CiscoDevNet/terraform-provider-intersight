# ResourcepoolChildLeaseCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "resourcepool.ChildLeaseCondition"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "resourcepool.ChildLeaseCondition"]
**AssignedToEntity** | Pointer to [**MoMoRef**](MoMoRef.md) |  | [optional] 
**Condition** | Pointer to [**[]ResourceResourceQualifier**](ResourceResourceQualifier.md) |  | [optional] 
**Feature** | Pointer to **string** | Lease operation applied for the feature. | [optional] [readonly] 
**Lease** | Pointer to [**MoMoRef**](MoMoRef.md) |  | [optional] 

## Methods

### NewResourcepoolChildLeaseCondition

`func NewResourcepoolChildLeaseCondition(classId string, objectType string, ) *ResourcepoolChildLeaseCondition`

NewResourcepoolChildLeaseCondition instantiates a new ResourcepoolChildLeaseCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcepoolChildLeaseConditionWithDefaults

`func NewResourcepoolChildLeaseConditionWithDefaults() *ResourcepoolChildLeaseCondition`

NewResourcepoolChildLeaseConditionWithDefaults instantiates a new ResourcepoolChildLeaseCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ResourcepoolChildLeaseCondition) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ResourcepoolChildLeaseCondition) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ResourcepoolChildLeaseCondition) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ResourcepoolChildLeaseCondition) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ResourcepoolChildLeaseCondition) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ResourcepoolChildLeaseCondition) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAssignedToEntity

`func (o *ResourcepoolChildLeaseCondition) GetAssignedToEntity() MoMoRef`

GetAssignedToEntity returns the AssignedToEntity field if non-nil, zero value otherwise.

### GetAssignedToEntityOk

`func (o *ResourcepoolChildLeaseCondition) GetAssignedToEntityOk() (*MoMoRef, bool)`

GetAssignedToEntityOk returns a tuple with the AssignedToEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedToEntity

`func (o *ResourcepoolChildLeaseCondition) SetAssignedToEntity(v MoMoRef)`

SetAssignedToEntity sets AssignedToEntity field to given value.

### HasAssignedToEntity

`func (o *ResourcepoolChildLeaseCondition) HasAssignedToEntity() bool`

HasAssignedToEntity returns a boolean if a field has been set.

### GetCondition

`func (o *ResourcepoolChildLeaseCondition) GetCondition() []ResourceResourceQualifier`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *ResourcepoolChildLeaseCondition) GetConditionOk() (*[]ResourceResourceQualifier, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *ResourcepoolChildLeaseCondition) SetCondition(v []ResourceResourceQualifier)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *ResourcepoolChildLeaseCondition) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *ResourcepoolChildLeaseCondition) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *ResourcepoolChildLeaseCondition) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil
### GetFeature

`func (o *ResourcepoolChildLeaseCondition) GetFeature() string`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *ResourcepoolChildLeaseCondition) GetFeatureOk() (*string, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *ResourcepoolChildLeaseCondition) SetFeature(v string)`

SetFeature sets Feature field to given value.

### HasFeature

`func (o *ResourcepoolChildLeaseCondition) HasFeature() bool`

HasFeature returns a boolean if a field has been set.

### GetLease

`func (o *ResourcepoolChildLeaseCondition) GetLease() MoMoRef`

GetLease returns the Lease field if non-nil, zero value otherwise.

### GetLeaseOk

`func (o *ResourcepoolChildLeaseCondition) GetLeaseOk() (*MoMoRef, bool)`

GetLeaseOk returns a tuple with the Lease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLease

`func (o *ResourcepoolChildLeaseCondition) SetLease(v MoMoRef)`

SetLease sets Lease field to given value.

### HasLease

`func (o *ResourcepoolChildLeaseCondition) HasLease() bool`

HasLease returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


