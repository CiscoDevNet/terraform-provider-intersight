# CapabilitySwitchCapabilityDefAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**Pid** | Pointer to **string** | Product Identifier for a Switch/Fabric-Interconnect. * &#x60;UCS-FI-6454&#x60; - The standard 4th generation UCS Fabric Interconnect with 54 ports. * &#x60;UCS-FI-64108&#x60; - The expanded 4th generation UCS Fabric Interconnect with 108 ports. * &#x60;unknown&#x60; - Unknown device type, usage is TBD. | [optional] [default to "UCS-FI-6454"]
**Sku** | Pointer to **string** | SKU information for Switch/Fabric-Interconnect. | [optional] 
**Vid** | Pointer to **string** | VID information for Switch/Fabric-Interconnect. | [optional] 

## Methods

### NewCapabilitySwitchCapabilityDefAllOf

`func NewCapabilitySwitchCapabilityDefAllOf(classId string, objectType string, ) *CapabilitySwitchCapabilityDefAllOf`

NewCapabilitySwitchCapabilityDefAllOf instantiates a new CapabilitySwitchCapabilityDefAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilitySwitchCapabilityDefAllOfWithDefaults

`func NewCapabilitySwitchCapabilityDefAllOfWithDefaults() *CapabilitySwitchCapabilityDefAllOf`

NewCapabilitySwitchCapabilityDefAllOfWithDefaults instantiates a new CapabilitySwitchCapabilityDefAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilitySwitchCapabilityDefAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilitySwitchCapabilityDefAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilitySwitchCapabilityDefAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilitySwitchCapabilityDefAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilitySwitchCapabilityDefAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilitySwitchCapabilityDefAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPid

`func (o *CapabilitySwitchCapabilityDefAllOf) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *CapabilitySwitchCapabilityDefAllOf) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *CapabilitySwitchCapabilityDefAllOf) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *CapabilitySwitchCapabilityDefAllOf) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetSku

`func (o *CapabilitySwitchCapabilityDefAllOf) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *CapabilitySwitchCapabilityDefAllOf) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *CapabilitySwitchCapabilityDefAllOf) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *CapabilitySwitchCapabilityDefAllOf) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetVid

`func (o *CapabilitySwitchCapabilityDefAllOf) GetVid() string`

GetVid returns the Vid field if non-nil, zero value otherwise.

### GetVidOk

`func (o *CapabilitySwitchCapabilityDefAllOf) GetVidOk() (*string, bool)`

GetVidOk returns a tuple with the Vid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVid

`func (o *CapabilitySwitchCapabilityDefAllOf) SetVid(v string)`

SetVid sets Vid field to given value.

### HasVid

`func (o *CapabilitySwitchCapabilityDefAllOf) HasVid() bool`

HasVid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


