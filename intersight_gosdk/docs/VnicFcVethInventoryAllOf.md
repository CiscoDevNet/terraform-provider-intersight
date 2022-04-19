# VnicFcVethInventoryAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.FcVethInventory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.FcVethInventory"]
**Burst** | Pointer to **int64** | The burst traffic, in bytes, allowed on the vNIC. | [optional] [default to 1024]
**Description** | Pointer to **string** | Description of the virtual FC interface. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the virtual FC interface. | [optional] 
**RateLimit** | Pointer to **int64** | The value in Mbps to use for limiting the data rate on the virtual interface. Setting this to zero will turn rate limiting off. | [optional] [default to 0]
**TargetMo** | Pointer to [**MoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewVnicFcVethInventoryAllOf

`func NewVnicFcVethInventoryAllOf(classId string, objectType string, ) *VnicFcVethInventoryAllOf`

NewVnicFcVethInventoryAllOf instantiates a new VnicFcVethInventoryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicFcVethInventoryAllOfWithDefaults

`func NewVnicFcVethInventoryAllOfWithDefaults() *VnicFcVethInventoryAllOf`

NewVnicFcVethInventoryAllOfWithDefaults instantiates a new VnicFcVethInventoryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicFcVethInventoryAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicFcVethInventoryAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicFcVethInventoryAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicFcVethInventoryAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicFcVethInventoryAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicFcVethInventoryAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBurst

`func (o *VnicFcVethInventoryAllOf) GetBurst() int64`

GetBurst returns the Burst field if non-nil, zero value otherwise.

### GetBurstOk

`func (o *VnicFcVethInventoryAllOf) GetBurstOk() (*int64, bool)`

GetBurstOk returns a tuple with the Burst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurst

`func (o *VnicFcVethInventoryAllOf) SetBurst(v int64)`

SetBurst sets Burst field to given value.

### HasBurst

`func (o *VnicFcVethInventoryAllOf) HasBurst() bool`

HasBurst returns a boolean if a field has been set.

### GetDescription

`func (o *VnicFcVethInventoryAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VnicFcVethInventoryAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VnicFcVethInventoryAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VnicFcVethInventoryAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *VnicFcVethInventoryAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VnicFcVethInventoryAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VnicFcVethInventoryAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VnicFcVethInventoryAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRateLimit

`func (o *VnicFcVethInventoryAllOf) GetRateLimit() int64`

GetRateLimit returns the RateLimit field if non-nil, zero value otherwise.

### GetRateLimitOk

`func (o *VnicFcVethInventoryAllOf) GetRateLimitOk() (*int64, bool)`

GetRateLimitOk returns a tuple with the RateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimit

`func (o *VnicFcVethInventoryAllOf) SetRateLimit(v int64)`

SetRateLimit sets RateLimit field to given value.

### HasRateLimit

`func (o *VnicFcVethInventoryAllOf) HasRateLimit() bool`

HasRateLimit returns a boolean if a field has been set.

### GetTargetMo

`func (o *VnicFcVethInventoryAllOf) GetTargetMo() MoBaseMoRelationship`

GetTargetMo returns the TargetMo field if non-nil, zero value otherwise.

### GetTargetMoOk

`func (o *VnicFcVethInventoryAllOf) GetTargetMoOk() (*MoBaseMoRelationship, bool)`

GetTargetMoOk returns a tuple with the TargetMo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMo

`func (o *VnicFcVethInventoryAllOf) SetTargetMo(v MoBaseMoRelationship)`

SetTargetMo sets TargetMo field to given value.

### HasTargetMo

`func (o *VnicFcVethInventoryAllOf) HasTargetMo() bool`

HasTargetMo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


