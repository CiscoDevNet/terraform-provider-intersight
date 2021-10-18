# ResourcepoolLeaseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "resourcepool.Lease"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "resourcepool.Lease"]
**Condition** | Pointer to [**[]ResourceSelector**](ResourceSelector.md) |  | [optional] 
**Feature** | Pointer to **string** | Lease opertion applied for the feature. | [optional] 
**LeaseParameters** | Pointer to [**NullableResourcepoolLeaseParameters**](ResourcepoolLeaseParameters.md) |  | [optional] 
**Resource** | Pointer to [**MoBaseMo**](MoBaseMo.md) |  | [optional] 
**ResourceType** | Pointer to **string** | The type of the resource present in the pool, example &#39;server&#39; its combination of RackUnit and Blade. * &#x60;None&#x60; - The resource cannot consider for Resource Pool. * &#x60;Server&#x60; - Resource Pool holds the server kind of resources, example - RackServer, Blade. | [optional] [default to "None"]
**AssignedToEntity** | Pointer to [**MoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 
**LeasedResource** | Pointer to [**ResourcepoolLeaseResourceRelationship**](ResourcepoolLeaseResourceRelationship.md) |  | [optional] 
**Pool** | Pointer to [**ResourcepoolPoolRelationship**](ResourcepoolPoolRelationship.md) |  | [optional] 
**PoolMember** | Pointer to [**ResourcepoolPoolMemberRelationship**](ResourcepoolPoolMemberRelationship.md) |  | [optional] 
**Universe** | Pointer to [**ResourcepoolUniverseRelationship**](ResourcepoolUniverseRelationship.md) |  | [optional] 

## Methods

### NewResourcepoolLeaseAllOf

`func NewResourcepoolLeaseAllOf(classId string, objectType string, ) *ResourcepoolLeaseAllOf`

NewResourcepoolLeaseAllOf instantiates a new ResourcepoolLeaseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcepoolLeaseAllOfWithDefaults

`func NewResourcepoolLeaseAllOfWithDefaults() *ResourcepoolLeaseAllOf`

NewResourcepoolLeaseAllOfWithDefaults instantiates a new ResourcepoolLeaseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ResourcepoolLeaseAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ResourcepoolLeaseAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ResourcepoolLeaseAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ResourcepoolLeaseAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ResourcepoolLeaseAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ResourcepoolLeaseAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCondition

`func (o *ResourcepoolLeaseAllOf) GetCondition() []ResourceSelector`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *ResourcepoolLeaseAllOf) GetConditionOk() (*[]ResourceSelector, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *ResourcepoolLeaseAllOf) SetCondition(v []ResourceSelector)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *ResourcepoolLeaseAllOf) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *ResourcepoolLeaseAllOf) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *ResourcepoolLeaseAllOf) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil
### GetFeature

`func (o *ResourcepoolLeaseAllOf) GetFeature() string`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *ResourcepoolLeaseAllOf) GetFeatureOk() (*string, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *ResourcepoolLeaseAllOf) SetFeature(v string)`

SetFeature sets Feature field to given value.

### HasFeature

`func (o *ResourcepoolLeaseAllOf) HasFeature() bool`

HasFeature returns a boolean if a field has been set.

### GetLeaseParameters

`func (o *ResourcepoolLeaseAllOf) GetLeaseParameters() ResourcepoolLeaseParameters`

GetLeaseParameters returns the LeaseParameters field if non-nil, zero value otherwise.

### GetLeaseParametersOk

`func (o *ResourcepoolLeaseAllOf) GetLeaseParametersOk() (*ResourcepoolLeaseParameters, bool)`

GetLeaseParametersOk returns a tuple with the LeaseParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseParameters

`func (o *ResourcepoolLeaseAllOf) SetLeaseParameters(v ResourcepoolLeaseParameters)`

SetLeaseParameters sets LeaseParameters field to given value.

### HasLeaseParameters

`func (o *ResourcepoolLeaseAllOf) HasLeaseParameters() bool`

HasLeaseParameters returns a boolean if a field has been set.

### SetLeaseParametersNil

`func (o *ResourcepoolLeaseAllOf) SetLeaseParametersNil(b bool)`

 SetLeaseParametersNil sets the value for LeaseParameters to be an explicit nil

### UnsetLeaseParameters
`func (o *ResourcepoolLeaseAllOf) UnsetLeaseParameters()`

UnsetLeaseParameters ensures that no value is present for LeaseParameters, not even an explicit nil
### GetResource

`func (o *ResourcepoolLeaseAllOf) GetResource() MoBaseMo`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ResourcepoolLeaseAllOf) GetResourceOk() (*MoBaseMo, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ResourcepoolLeaseAllOf) SetResource(v MoBaseMo)`

SetResource sets Resource field to given value.

### HasResource

`func (o *ResourcepoolLeaseAllOf) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetResourceType

`func (o *ResourcepoolLeaseAllOf) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ResourcepoolLeaseAllOf) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ResourcepoolLeaseAllOf) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *ResourcepoolLeaseAllOf) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetAssignedToEntity

`func (o *ResourcepoolLeaseAllOf) GetAssignedToEntity() MoBaseMoRelationship`

GetAssignedToEntity returns the AssignedToEntity field if non-nil, zero value otherwise.

### GetAssignedToEntityOk

`func (o *ResourcepoolLeaseAllOf) GetAssignedToEntityOk() (*MoBaseMoRelationship, bool)`

GetAssignedToEntityOk returns a tuple with the AssignedToEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedToEntity

`func (o *ResourcepoolLeaseAllOf) SetAssignedToEntity(v MoBaseMoRelationship)`

SetAssignedToEntity sets AssignedToEntity field to given value.

### HasAssignedToEntity

`func (o *ResourcepoolLeaseAllOf) HasAssignedToEntity() bool`

HasAssignedToEntity returns a boolean if a field has been set.

### GetLeasedResource

`func (o *ResourcepoolLeaseAllOf) GetLeasedResource() ResourcepoolLeaseResourceRelationship`

GetLeasedResource returns the LeasedResource field if non-nil, zero value otherwise.

### GetLeasedResourceOk

`func (o *ResourcepoolLeaseAllOf) GetLeasedResourceOk() (*ResourcepoolLeaseResourceRelationship, bool)`

GetLeasedResourceOk returns a tuple with the LeasedResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeasedResource

`func (o *ResourcepoolLeaseAllOf) SetLeasedResource(v ResourcepoolLeaseResourceRelationship)`

SetLeasedResource sets LeasedResource field to given value.

### HasLeasedResource

`func (o *ResourcepoolLeaseAllOf) HasLeasedResource() bool`

HasLeasedResource returns a boolean if a field has been set.

### GetPool

`func (o *ResourcepoolLeaseAllOf) GetPool() ResourcepoolPoolRelationship`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *ResourcepoolLeaseAllOf) GetPoolOk() (*ResourcepoolPoolRelationship, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *ResourcepoolLeaseAllOf) SetPool(v ResourcepoolPoolRelationship)`

SetPool sets Pool field to given value.

### HasPool

`func (o *ResourcepoolLeaseAllOf) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetPoolMember

`func (o *ResourcepoolLeaseAllOf) GetPoolMember() ResourcepoolPoolMemberRelationship`

GetPoolMember returns the PoolMember field if non-nil, zero value otherwise.

### GetPoolMemberOk

`func (o *ResourcepoolLeaseAllOf) GetPoolMemberOk() (*ResourcepoolPoolMemberRelationship, bool)`

GetPoolMemberOk returns a tuple with the PoolMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolMember

`func (o *ResourcepoolLeaseAllOf) SetPoolMember(v ResourcepoolPoolMemberRelationship)`

SetPoolMember sets PoolMember field to given value.

### HasPoolMember

`func (o *ResourcepoolLeaseAllOf) HasPoolMember() bool`

HasPoolMember returns a boolean if a field has been set.

### GetUniverse

`func (o *ResourcepoolLeaseAllOf) GetUniverse() ResourcepoolUniverseRelationship`

GetUniverse returns the Universe field if non-nil, zero value otherwise.

### GetUniverseOk

`func (o *ResourcepoolLeaseAllOf) GetUniverseOk() (*ResourcepoolUniverseRelationship, bool)`

GetUniverseOk returns a tuple with the Universe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverse

`func (o *ResourcepoolLeaseAllOf) SetUniverse(v ResourcepoolUniverseRelationship)`

SetUniverse sets Universe field to given value.

### HasUniverse

`func (o *ResourcepoolLeaseAllOf) HasUniverse() bool`

HasUniverse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


