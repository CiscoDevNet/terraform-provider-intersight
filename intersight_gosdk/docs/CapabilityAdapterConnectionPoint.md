# CapabilityAdapterConnectionPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.AdapterConnectionPoint"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.AdapterConnectionPoint"]
**EndPointIds** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewCapabilityAdapterConnectionPoint

`func NewCapabilityAdapterConnectionPoint(classId string, objectType string, ) *CapabilityAdapterConnectionPoint`

NewCapabilityAdapterConnectionPoint instantiates a new CapabilityAdapterConnectionPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityAdapterConnectionPointWithDefaults

`func NewCapabilityAdapterConnectionPointWithDefaults() *CapabilityAdapterConnectionPoint`

NewCapabilityAdapterConnectionPointWithDefaults instantiates a new CapabilityAdapterConnectionPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityAdapterConnectionPoint) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityAdapterConnectionPoint) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityAdapterConnectionPoint) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityAdapterConnectionPoint) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityAdapterConnectionPoint) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityAdapterConnectionPoint) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEndPointIds

`func (o *CapabilityAdapterConnectionPoint) GetEndPointIds() []int64`

GetEndPointIds returns the EndPointIds field if non-nil, zero value otherwise.

### GetEndPointIdsOk

`func (o *CapabilityAdapterConnectionPoint) GetEndPointIdsOk() (*[]int64, bool)`

GetEndPointIdsOk returns a tuple with the EndPointIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPointIds

`func (o *CapabilityAdapterConnectionPoint) SetEndPointIds(v []int64)`

SetEndPointIds sets EndPointIds field to given value.

### HasEndPointIds

`func (o *CapabilityAdapterConnectionPoint) HasEndPointIds() bool`

HasEndPointIds returns a boolean if a field has been set.

### SetEndPointIdsNil

`func (o *CapabilityAdapterConnectionPoint) SetEndPointIdsNil(b bool)`

 SetEndPointIdsNil sets the value for EndPointIds to be an explicit nil

### UnsetEndPointIds
`func (o *CapabilityAdapterConnectionPoint) UnsetEndPointIds()`

UnsetEndPointIds ensures that no value is present for EndPointIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


