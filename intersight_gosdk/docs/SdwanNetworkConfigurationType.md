# SdwanNetworkConfigurationType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "sdwan.NetworkConfigurationType"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "sdwan.NetworkConfigurationType"]
**NetworkType** | Pointer to **string** | Type of the Port group being added. * &#x60;WAN&#x60; - Port-group being added is used for WAN traffic. * &#x60;LAN&#x60; - Port-group being added is used for LAN traffic. * &#x60;Management&#x60; - Port-group being added is used for Management traffic. | [optional] [default to "WAN"]
**PortGroup** | Pointer to **string** | Name of the Port Group to create. | [optional] 
**Vlan** | Pointer to **int64** | VLAN to be added to the Port Group. | [optional] 

## Methods

### NewSdwanNetworkConfigurationType

`func NewSdwanNetworkConfigurationType(classId string, objectType string, ) *SdwanNetworkConfigurationType`

NewSdwanNetworkConfigurationType instantiates a new SdwanNetworkConfigurationType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdwanNetworkConfigurationTypeWithDefaults

`func NewSdwanNetworkConfigurationTypeWithDefaults() *SdwanNetworkConfigurationType`

NewSdwanNetworkConfigurationTypeWithDefaults instantiates a new SdwanNetworkConfigurationType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SdwanNetworkConfigurationType) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SdwanNetworkConfigurationType) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SdwanNetworkConfigurationType) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SdwanNetworkConfigurationType) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SdwanNetworkConfigurationType) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SdwanNetworkConfigurationType) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetNetworkType

`func (o *SdwanNetworkConfigurationType) GetNetworkType() string`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *SdwanNetworkConfigurationType) GetNetworkTypeOk() (*string, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *SdwanNetworkConfigurationType) SetNetworkType(v string)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *SdwanNetworkConfigurationType) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.

### GetPortGroup

`func (o *SdwanNetworkConfigurationType) GetPortGroup() string`

GetPortGroup returns the PortGroup field if non-nil, zero value otherwise.

### GetPortGroupOk

`func (o *SdwanNetworkConfigurationType) GetPortGroupOk() (*string, bool)`

GetPortGroupOk returns a tuple with the PortGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortGroup

`func (o *SdwanNetworkConfigurationType) SetPortGroup(v string)`

SetPortGroup sets PortGroup field to given value.

### HasPortGroup

`func (o *SdwanNetworkConfigurationType) HasPortGroup() bool`

HasPortGroup returns a boolean if a field has been set.

### GetVlan

`func (o *SdwanNetworkConfigurationType) GetVlan() int64`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *SdwanNetworkConfigurationType) GetVlanOk() (*int64, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *SdwanNetworkConfigurationType) SetVlan(v int64)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *SdwanNetworkConfigurationType) HasVlan() bool`

HasVlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


