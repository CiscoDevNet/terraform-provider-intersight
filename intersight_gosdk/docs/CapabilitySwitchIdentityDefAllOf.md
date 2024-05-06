# CapabilitySwitchIdentityDefAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "capability.SwitchEquipmentInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "capability.SwitchEquipmentInfo"]
**Pid** | Pointer to **string** | Product Identifier for a Switch/Fabric-Interconnect. | [optional] [readonly] 
**Sku** | Pointer to **string** | SKU information for Switch/Fabric-Interconnect. | [optional] [readonly] 
**Vid** | Pointer to **string** | VID information for Switch/Fabric-Interconnect. | [optional] [readonly] 

## Methods

### NewCapabilitySwitchIdentityDefAllOf

`func NewCapabilitySwitchIdentityDefAllOf(classId string, objectType string, ) *CapabilitySwitchIdentityDefAllOf`

NewCapabilitySwitchIdentityDefAllOf instantiates a new CapabilitySwitchIdentityDefAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilitySwitchIdentityDefAllOfWithDefaults

`func NewCapabilitySwitchIdentityDefAllOfWithDefaults() *CapabilitySwitchIdentityDefAllOf`

NewCapabilitySwitchIdentityDefAllOfWithDefaults instantiates a new CapabilitySwitchIdentityDefAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilitySwitchIdentityDefAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilitySwitchIdentityDefAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilitySwitchIdentityDefAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilitySwitchIdentityDefAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilitySwitchIdentityDefAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilitySwitchIdentityDefAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPid

`func (o *CapabilitySwitchIdentityDefAllOf) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *CapabilitySwitchIdentityDefAllOf) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *CapabilitySwitchIdentityDefAllOf) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *CapabilitySwitchIdentityDefAllOf) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetSku

`func (o *CapabilitySwitchIdentityDefAllOf) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *CapabilitySwitchIdentityDefAllOf) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *CapabilitySwitchIdentityDefAllOf) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *CapabilitySwitchIdentityDefAllOf) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetVid

`func (o *CapabilitySwitchIdentityDefAllOf) GetVid() string`

GetVid returns the Vid field if non-nil, zero value otherwise.

### GetVidOk

`func (o *CapabilitySwitchIdentityDefAllOf) GetVidOk() (*string, bool)`

GetVidOk returns a tuple with the Vid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVid

`func (o *CapabilitySwitchIdentityDefAllOf) SetVid(v string)`

SetVid sets Vid field to given value.

### HasVid

`func (o *CapabilitySwitchIdentityDefAllOf) HasVid() bool`

HasVid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


