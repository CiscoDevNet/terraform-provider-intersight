# NetworkInterfaceListAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "network.InterfaceList"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "network.InterfaceList"]
**AdminState** | Pointer to **string** | Admin state of the interface list. | [optional] 
**IpAddress** | Pointer to **string** | IP address of the interface list. | [optional] 
**Mac** | Pointer to **string** | MAC address of the interface list. | [optional] 
**Name** | Pointer to **string** | Name of the interface list. | [optional] 
**OperState** | Pointer to **string** | Operational state of the interface list. | [optional] 
**PortType** | Pointer to **string** | Port type of interface list. | [optional] 
**SlotId** | Pointer to **string** | Slot id of the interface list. | [optional] 
**Vlan** | Pointer to **string** | VLAN of the interface list. | [optional] 
**NetworkElement** | Pointer to [**NetworkElementRelationship**](NetworkElementRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNetworkInterfaceListAllOf

`func NewNetworkInterfaceListAllOf(classId string, objectType string, ) *NetworkInterfaceListAllOf`

NewNetworkInterfaceListAllOf instantiates a new NetworkInterfaceListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkInterfaceListAllOfWithDefaults

`func NewNetworkInterfaceListAllOfWithDefaults() *NetworkInterfaceListAllOf`

NewNetworkInterfaceListAllOfWithDefaults instantiates a new NetworkInterfaceListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NetworkInterfaceListAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NetworkInterfaceListAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NetworkInterfaceListAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NetworkInterfaceListAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NetworkInterfaceListAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NetworkInterfaceListAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminState

`func (o *NetworkInterfaceListAllOf) GetAdminState() string`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *NetworkInterfaceListAllOf) GetAdminStateOk() (*string, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *NetworkInterfaceListAllOf) SetAdminState(v string)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *NetworkInterfaceListAllOf) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetIpAddress

`func (o *NetworkInterfaceListAllOf) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *NetworkInterfaceListAllOf) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *NetworkInterfaceListAllOf) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *NetworkInterfaceListAllOf) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetMac

`func (o *NetworkInterfaceListAllOf) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *NetworkInterfaceListAllOf) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *NetworkInterfaceListAllOf) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *NetworkInterfaceListAllOf) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetName

`func (o *NetworkInterfaceListAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkInterfaceListAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkInterfaceListAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkInterfaceListAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperState

`func (o *NetworkInterfaceListAllOf) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *NetworkInterfaceListAllOf) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *NetworkInterfaceListAllOf) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *NetworkInterfaceListAllOf) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetPortType

`func (o *NetworkInterfaceListAllOf) GetPortType() string`

GetPortType returns the PortType field if non-nil, zero value otherwise.

### GetPortTypeOk

`func (o *NetworkInterfaceListAllOf) GetPortTypeOk() (*string, bool)`

GetPortTypeOk returns a tuple with the PortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortType

`func (o *NetworkInterfaceListAllOf) SetPortType(v string)`

SetPortType sets PortType field to given value.

### HasPortType

`func (o *NetworkInterfaceListAllOf) HasPortType() bool`

HasPortType returns a boolean if a field has been set.

### GetSlotId

`func (o *NetworkInterfaceListAllOf) GetSlotId() string`

GetSlotId returns the SlotId field if non-nil, zero value otherwise.

### GetSlotIdOk

`func (o *NetworkInterfaceListAllOf) GetSlotIdOk() (*string, bool)`

GetSlotIdOk returns a tuple with the SlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotId

`func (o *NetworkInterfaceListAllOf) SetSlotId(v string)`

SetSlotId sets SlotId field to given value.

### HasSlotId

`func (o *NetworkInterfaceListAllOf) HasSlotId() bool`

HasSlotId returns a boolean if a field has been set.

### GetVlan

`func (o *NetworkInterfaceListAllOf) GetVlan() string`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *NetworkInterfaceListAllOf) GetVlanOk() (*string, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *NetworkInterfaceListAllOf) SetVlan(v string)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *NetworkInterfaceListAllOf) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetNetworkElement

`func (o *NetworkInterfaceListAllOf) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *NetworkInterfaceListAllOf) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *NetworkInterfaceListAllOf) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *NetworkInterfaceListAllOf) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NetworkInterfaceListAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NetworkInterfaceListAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NetworkInterfaceListAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NetworkInterfaceListAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


