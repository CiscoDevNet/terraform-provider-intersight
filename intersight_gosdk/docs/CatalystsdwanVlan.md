# CatalystsdwanVlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "catalystsdwan.Vlan"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "catalystsdwan.Vlan"]
**AdminState** | Pointer to **string** | The administrative state of the WAN Edge Device. * &#x60;Unknown&#x60; - Administrative state is unknown. * &#x60;Up&#x60; - Administrative state is up. * &#x60;Down&#x60; - Administrative state is down. | [optional] [readonly] [default to "Unknown"]
**Dn** | Pointer to **string** | The distinguished name of the VLAN. | [optional] [readonly] 
**SystemIp** | Pointer to **string** | The system IP address of the WAN Edge Device. | [optional] [readonly] 
**Uuid** | Pointer to **string** | The UUID of the WAN Edge Device to which this VLAN belongs. | [optional] [readonly] 
**VlanId** | Pointer to **int64** | The VLAN number configured on the WAN Edge Device. | [optional] [readonly] 
**VlanState** | Pointer to **string** | The state of the VLAN on the WAN Edge Device. * &#x60;Unknown&#x60; - VLAN state is unknown on the Catalyst SDWAN device. * &#x60;Active&#x60; - VLAN state is active on the Catalyst SDWAN device. * &#x60;Inactive&#x60; - VLAN state is inactive on the Catalyst SDWAN device. * &#x60;Suspended&#x60; - VLAN state is suspended on the Catalyst SDWAN device. | [optional] [readonly] [default to "Unknown"]
**WanEdgeDevice** | Pointer to [**NullableCatalystsdwanWanEdgeDeviceRelationship**](CatalystsdwanWanEdgeDeviceRelationship.md) |  | [optional] 

## Methods

### NewCatalystsdwanVlan

`func NewCatalystsdwanVlan(classId string, objectType string, ) *CatalystsdwanVlan`

NewCatalystsdwanVlan instantiates a new CatalystsdwanVlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalystsdwanVlanWithDefaults

`func NewCatalystsdwanVlanWithDefaults() *CatalystsdwanVlan`

NewCatalystsdwanVlanWithDefaults instantiates a new CatalystsdwanVlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CatalystsdwanVlan) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CatalystsdwanVlan) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CatalystsdwanVlan) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CatalystsdwanVlan) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CatalystsdwanVlan) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CatalystsdwanVlan) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminState

`func (o *CatalystsdwanVlan) GetAdminState() string`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *CatalystsdwanVlan) GetAdminStateOk() (*string, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *CatalystsdwanVlan) SetAdminState(v string)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *CatalystsdwanVlan) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetDn

`func (o *CatalystsdwanVlan) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *CatalystsdwanVlan) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *CatalystsdwanVlan) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *CatalystsdwanVlan) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetSystemIp

`func (o *CatalystsdwanVlan) GetSystemIp() string`

GetSystemIp returns the SystemIp field if non-nil, zero value otherwise.

### GetSystemIpOk

`func (o *CatalystsdwanVlan) GetSystemIpOk() (*string, bool)`

GetSystemIpOk returns a tuple with the SystemIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemIp

`func (o *CatalystsdwanVlan) SetSystemIp(v string)`

SetSystemIp sets SystemIp field to given value.

### HasSystemIp

`func (o *CatalystsdwanVlan) HasSystemIp() bool`

HasSystemIp returns a boolean if a field has been set.

### GetUuid

`func (o *CatalystsdwanVlan) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *CatalystsdwanVlan) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *CatalystsdwanVlan) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *CatalystsdwanVlan) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVlanId

`func (o *CatalystsdwanVlan) GetVlanId() int64`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *CatalystsdwanVlan) GetVlanIdOk() (*int64, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *CatalystsdwanVlan) SetVlanId(v int64)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *CatalystsdwanVlan) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetVlanState

`func (o *CatalystsdwanVlan) GetVlanState() string`

GetVlanState returns the VlanState field if non-nil, zero value otherwise.

### GetVlanStateOk

`func (o *CatalystsdwanVlan) GetVlanStateOk() (*string, bool)`

GetVlanStateOk returns a tuple with the VlanState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanState

`func (o *CatalystsdwanVlan) SetVlanState(v string)`

SetVlanState sets VlanState field to given value.

### HasVlanState

`func (o *CatalystsdwanVlan) HasVlanState() bool`

HasVlanState returns a boolean if a field has been set.

### GetWanEdgeDevice

`func (o *CatalystsdwanVlan) GetWanEdgeDevice() CatalystsdwanWanEdgeDeviceRelationship`

GetWanEdgeDevice returns the WanEdgeDevice field if non-nil, zero value otherwise.

### GetWanEdgeDeviceOk

`func (o *CatalystsdwanVlan) GetWanEdgeDeviceOk() (*CatalystsdwanWanEdgeDeviceRelationship, bool)`

GetWanEdgeDeviceOk returns a tuple with the WanEdgeDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanEdgeDevice

`func (o *CatalystsdwanVlan) SetWanEdgeDevice(v CatalystsdwanWanEdgeDeviceRelationship)`

SetWanEdgeDevice sets WanEdgeDevice field to given value.

### HasWanEdgeDevice

`func (o *CatalystsdwanVlan) HasWanEdgeDevice() bool`

HasWanEdgeDevice returns a boolean if a field has been set.

### SetWanEdgeDeviceNil

`func (o *CatalystsdwanVlan) SetWanEdgeDeviceNil(b bool)`

 SetWanEdgeDeviceNil sets the value for WanEdgeDevice to be an explicit nil

### UnsetWanEdgeDevice
`func (o *CatalystsdwanVlan) UnsetWanEdgeDevice()`

UnsetWanEdgeDevice ensures that no value is present for WanEdgeDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


