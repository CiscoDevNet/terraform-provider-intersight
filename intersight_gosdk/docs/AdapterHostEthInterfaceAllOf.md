# AdapterHostEthInterfaceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "adapter.HostEthInterface"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "adapter.HostEthInterface"]
**AdminState** | Pointer to **string** | Admin state of the Host Ethernet Interface. | [optional] [readonly] 
**EpDn** | Pointer to **string** | The Endpoint Config Dn of the Host Ethernet Interface. | [optional] [readonly] 
**HostEthInterfaceId** | Pointer to **int64** | Unique Identifier for an Host Ethernet Interface within the adapter object. | [optional] [readonly] 
**InterfaceType** | Pointer to **string** | Type of External Ethernet Interface. | [optional] [readonly] 
**MacAddress** | Pointer to **string** | Mac address of the Host Ethernet Interface. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of Host Ethernet Interface. | [optional] [readonly] 
**OperReason** | Pointer to **[]string** |  | [optional] 
**Operability** | Pointer to **string** | Operability status of Host Ethernet Channel Interface. | [optional] [readonly] 
**OriginalMacAddress** | Pointer to **string** | The factory default Mac address of the Host Ethernet Interface. | [optional] [readonly] 
**PciAddr** | Pointer to **string** | The PCI address of the Host Ethernet Interface. | [optional] [readonly] 
**PeerDn** | Pointer to **string** | The distinguished name of the peer endpoint connected to the Host Ethernet interface. | [optional] [readonly] 
**VirtualizationPreference** | Pointer to **string** | Virtualization Preference of the Host Ethernet Interface indicating if virtualization is enabled or not. | [optional] [readonly] 
**VnicDn** | Pointer to **string** | The Virtual Ethernet Interface DN connected to the Host Ethernet Interface. | [optional] [readonly] 
**AdapterUnit** | Pointer to [**AdapterUnitRelationship**](AdapterUnitRelationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewAdapterHostEthInterfaceAllOf

`func NewAdapterHostEthInterfaceAllOf(classId string, objectType string, ) *AdapterHostEthInterfaceAllOf`

NewAdapterHostEthInterfaceAllOf instantiates a new AdapterHostEthInterfaceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdapterHostEthInterfaceAllOfWithDefaults

`func NewAdapterHostEthInterfaceAllOfWithDefaults() *AdapterHostEthInterfaceAllOf`

NewAdapterHostEthInterfaceAllOfWithDefaults instantiates a new AdapterHostEthInterfaceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AdapterHostEthInterfaceAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AdapterHostEthInterfaceAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AdapterHostEthInterfaceAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AdapterHostEthInterfaceAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AdapterHostEthInterfaceAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AdapterHostEthInterfaceAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminState

`func (o *AdapterHostEthInterfaceAllOf) GetAdminState() string`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *AdapterHostEthInterfaceAllOf) GetAdminStateOk() (*string, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *AdapterHostEthInterfaceAllOf) SetAdminState(v string)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *AdapterHostEthInterfaceAllOf) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetEpDn

`func (o *AdapterHostEthInterfaceAllOf) GetEpDn() string`

GetEpDn returns the EpDn field if non-nil, zero value otherwise.

### GetEpDnOk

`func (o *AdapterHostEthInterfaceAllOf) GetEpDnOk() (*string, bool)`

GetEpDnOk returns a tuple with the EpDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpDn

`func (o *AdapterHostEthInterfaceAllOf) SetEpDn(v string)`

SetEpDn sets EpDn field to given value.

### HasEpDn

`func (o *AdapterHostEthInterfaceAllOf) HasEpDn() bool`

HasEpDn returns a boolean if a field has been set.

### GetHostEthInterfaceId

`func (o *AdapterHostEthInterfaceAllOf) GetHostEthInterfaceId() int64`

GetHostEthInterfaceId returns the HostEthInterfaceId field if non-nil, zero value otherwise.

### GetHostEthInterfaceIdOk

`func (o *AdapterHostEthInterfaceAllOf) GetHostEthInterfaceIdOk() (*int64, bool)`

GetHostEthInterfaceIdOk returns a tuple with the HostEthInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostEthInterfaceId

`func (o *AdapterHostEthInterfaceAllOf) SetHostEthInterfaceId(v int64)`

SetHostEthInterfaceId sets HostEthInterfaceId field to given value.

### HasHostEthInterfaceId

`func (o *AdapterHostEthInterfaceAllOf) HasHostEthInterfaceId() bool`

HasHostEthInterfaceId returns a boolean if a field has been set.

### GetInterfaceType

`func (o *AdapterHostEthInterfaceAllOf) GetInterfaceType() string`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *AdapterHostEthInterfaceAllOf) GetInterfaceTypeOk() (*string, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *AdapterHostEthInterfaceAllOf) SetInterfaceType(v string)`

SetInterfaceType sets InterfaceType field to given value.

### HasInterfaceType

`func (o *AdapterHostEthInterfaceAllOf) HasInterfaceType() bool`

HasInterfaceType returns a boolean if a field has been set.

### GetMacAddress

`func (o *AdapterHostEthInterfaceAllOf) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *AdapterHostEthInterfaceAllOf) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *AdapterHostEthInterfaceAllOf) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *AdapterHostEthInterfaceAllOf) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetName

`func (o *AdapterHostEthInterfaceAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdapterHostEthInterfaceAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdapterHostEthInterfaceAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdapterHostEthInterfaceAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperReason

`func (o *AdapterHostEthInterfaceAllOf) GetOperReason() []string`

GetOperReason returns the OperReason field if non-nil, zero value otherwise.

### GetOperReasonOk

`func (o *AdapterHostEthInterfaceAllOf) GetOperReasonOk() (*[]string, bool)`

GetOperReasonOk returns a tuple with the OperReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperReason

`func (o *AdapterHostEthInterfaceAllOf) SetOperReason(v []string)`

SetOperReason sets OperReason field to given value.

### HasOperReason

`func (o *AdapterHostEthInterfaceAllOf) HasOperReason() bool`

HasOperReason returns a boolean if a field has been set.

### SetOperReasonNil

`func (o *AdapterHostEthInterfaceAllOf) SetOperReasonNil(b bool)`

 SetOperReasonNil sets the value for OperReason to be an explicit nil

### UnsetOperReason
`func (o *AdapterHostEthInterfaceAllOf) UnsetOperReason()`

UnsetOperReason ensures that no value is present for OperReason, not even an explicit nil
### GetOperability

`func (o *AdapterHostEthInterfaceAllOf) GetOperability() string`

GetOperability returns the Operability field if non-nil, zero value otherwise.

### GetOperabilityOk

`func (o *AdapterHostEthInterfaceAllOf) GetOperabilityOk() (*string, bool)`

GetOperabilityOk returns a tuple with the Operability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperability

`func (o *AdapterHostEthInterfaceAllOf) SetOperability(v string)`

SetOperability sets Operability field to given value.

### HasOperability

`func (o *AdapterHostEthInterfaceAllOf) HasOperability() bool`

HasOperability returns a boolean if a field has been set.

### GetOriginalMacAddress

`func (o *AdapterHostEthInterfaceAllOf) GetOriginalMacAddress() string`

GetOriginalMacAddress returns the OriginalMacAddress field if non-nil, zero value otherwise.

### GetOriginalMacAddressOk

`func (o *AdapterHostEthInterfaceAllOf) GetOriginalMacAddressOk() (*string, bool)`

GetOriginalMacAddressOk returns a tuple with the OriginalMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalMacAddress

`func (o *AdapterHostEthInterfaceAllOf) SetOriginalMacAddress(v string)`

SetOriginalMacAddress sets OriginalMacAddress field to given value.

### HasOriginalMacAddress

`func (o *AdapterHostEthInterfaceAllOf) HasOriginalMacAddress() bool`

HasOriginalMacAddress returns a boolean if a field has been set.

### GetPciAddr

`func (o *AdapterHostEthInterfaceAllOf) GetPciAddr() string`

GetPciAddr returns the PciAddr field if non-nil, zero value otherwise.

### GetPciAddrOk

`func (o *AdapterHostEthInterfaceAllOf) GetPciAddrOk() (*string, bool)`

GetPciAddrOk returns a tuple with the PciAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciAddr

`func (o *AdapterHostEthInterfaceAllOf) SetPciAddr(v string)`

SetPciAddr sets PciAddr field to given value.

### HasPciAddr

`func (o *AdapterHostEthInterfaceAllOf) HasPciAddr() bool`

HasPciAddr returns a boolean if a field has been set.

### GetPeerDn

`func (o *AdapterHostEthInterfaceAllOf) GetPeerDn() string`

GetPeerDn returns the PeerDn field if non-nil, zero value otherwise.

### GetPeerDnOk

`func (o *AdapterHostEthInterfaceAllOf) GetPeerDnOk() (*string, bool)`

GetPeerDnOk returns a tuple with the PeerDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerDn

`func (o *AdapterHostEthInterfaceAllOf) SetPeerDn(v string)`

SetPeerDn sets PeerDn field to given value.

### HasPeerDn

`func (o *AdapterHostEthInterfaceAllOf) HasPeerDn() bool`

HasPeerDn returns a boolean if a field has been set.

### GetVirtualizationPreference

`func (o *AdapterHostEthInterfaceAllOf) GetVirtualizationPreference() string`

GetVirtualizationPreference returns the VirtualizationPreference field if non-nil, zero value otherwise.

### GetVirtualizationPreferenceOk

`func (o *AdapterHostEthInterfaceAllOf) GetVirtualizationPreferenceOk() (*string, bool)`

GetVirtualizationPreferenceOk returns a tuple with the VirtualizationPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualizationPreference

`func (o *AdapterHostEthInterfaceAllOf) SetVirtualizationPreference(v string)`

SetVirtualizationPreference sets VirtualizationPreference field to given value.

### HasVirtualizationPreference

`func (o *AdapterHostEthInterfaceAllOf) HasVirtualizationPreference() bool`

HasVirtualizationPreference returns a boolean if a field has been set.

### GetVnicDn

`func (o *AdapterHostEthInterfaceAllOf) GetVnicDn() string`

GetVnicDn returns the VnicDn field if non-nil, zero value otherwise.

### GetVnicDnOk

`func (o *AdapterHostEthInterfaceAllOf) GetVnicDnOk() (*string, bool)`

GetVnicDnOk returns a tuple with the VnicDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnicDn

`func (o *AdapterHostEthInterfaceAllOf) SetVnicDn(v string)`

SetVnicDn sets VnicDn field to given value.

### HasVnicDn

`func (o *AdapterHostEthInterfaceAllOf) HasVnicDn() bool`

HasVnicDn returns a boolean if a field has been set.

### GetAdapterUnit

`func (o *AdapterHostEthInterfaceAllOf) GetAdapterUnit() AdapterUnitRelationship`

GetAdapterUnit returns the AdapterUnit field if non-nil, zero value otherwise.

### GetAdapterUnitOk

`func (o *AdapterHostEthInterfaceAllOf) GetAdapterUnitOk() (*AdapterUnitRelationship, bool)`

GetAdapterUnitOk returns a tuple with the AdapterUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterUnit

`func (o *AdapterHostEthInterfaceAllOf) SetAdapterUnit(v AdapterUnitRelationship)`

SetAdapterUnit sets AdapterUnit field to given value.

### HasAdapterUnit

`func (o *AdapterHostEthInterfaceAllOf) HasAdapterUnit() bool`

HasAdapterUnit returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *AdapterHostEthInterfaceAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *AdapterHostEthInterfaceAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *AdapterHostEthInterfaceAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *AdapterHostEthInterfaceAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *AdapterHostEthInterfaceAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *AdapterHostEthInterfaceAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *AdapterHostEthInterfaceAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *AdapterHostEthInterfaceAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


