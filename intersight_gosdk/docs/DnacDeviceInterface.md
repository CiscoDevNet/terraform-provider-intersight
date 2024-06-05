# DnacDeviceInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "dnac.DeviceInterface"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "dnac.DeviceInterface"]
**AdminStatus** | Pointer to **string** | Administrator status for network device interface. | [optional] 
**DeviceId** | Pointer to **string** | Identity of network device. | [optional] 
**InterfaceId** | Pointer to **string** | Unique identity of network device interface. | [optional] 
**InterfaceName** | Pointer to **string** | The name for the network device interface. | [optional] 
**InterfaceType** | Pointer to **string** | Interface type for network device interface. | [optional] 
**PortType** | Pointer to **string** | Port type for network device interface. | [optional] 

## Methods

### NewDnacDeviceInterface

`func NewDnacDeviceInterface(classId string, objectType string, ) *DnacDeviceInterface`

NewDnacDeviceInterface instantiates a new DnacDeviceInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnacDeviceInterfaceWithDefaults

`func NewDnacDeviceInterfaceWithDefaults() *DnacDeviceInterface`

NewDnacDeviceInterfaceWithDefaults instantiates a new DnacDeviceInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *DnacDeviceInterface) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *DnacDeviceInterface) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *DnacDeviceInterface) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *DnacDeviceInterface) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *DnacDeviceInterface) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *DnacDeviceInterface) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminStatus

`func (o *DnacDeviceInterface) GetAdminStatus() string`

GetAdminStatus returns the AdminStatus field if non-nil, zero value otherwise.

### GetAdminStatusOk

`func (o *DnacDeviceInterface) GetAdminStatusOk() (*string, bool)`

GetAdminStatusOk returns a tuple with the AdminStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminStatus

`func (o *DnacDeviceInterface) SetAdminStatus(v string)`

SetAdminStatus sets AdminStatus field to given value.

### HasAdminStatus

`func (o *DnacDeviceInterface) HasAdminStatus() bool`

HasAdminStatus returns a boolean if a field has been set.

### GetDeviceId

`func (o *DnacDeviceInterface) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *DnacDeviceInterface) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *DnacDeviceInterface) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *DnacDeviceInterface) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetInterfaceId

`func (o *DnacDeviceInterface) GetInterfaceId() string`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *DnacDeviceInterface) GetInterfaceIdOk() (*string, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *DnacDeviceInterface) SetInterfaceId(v string)`

SetInterfaceId sets InterfaceId field to given value.

### HasInterfaceId

`func (o *DnacDeviceInterface) HasInterfaceId() bool`

HasInterfaceId returns a boolean if a field has been set.

### GetInterfaceName

`func (o *DnacDeviceInterface) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *DnacDeviceInterface) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *DnacDeviceInterface) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.

### HasInterfaceName

`func (o *DnacDeviceInterface) HasInterfaceName() bool`

HasInterfaceName returns a boolean if a field has been set.

### GetInterfaceType

`func (o *DnacDeviceInterface) GetInterfaceType() string`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *DnacDeviceInterface) GetInterfaceTypeOk() (*string, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *DnacDeviceInterface) SetInterfaceType(v string)`

SetInterfaceType sets InterfaceType field to given value.

### HasInterfaceType

`func (o *DnacDeviceInterface) HasInterfaceType() bool`

HasInterfaceType returns a boolean if a field has been set.

### GetPortType

`func (o *DnacDeviceInterface) GetPortType() string`

GetPortType returns the PortType field if non-nil, zero value otherwise.

### GetPortTypeOk

`func (o *DnacDeviceInterface) GetPortTypeOk() (*string, bool)`

GetPortTypeOk returns a tuple with the PortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortType

`func (o *DnacDeviceInterface) SetPortType(v string)`

SetPortType sets PortType field to given value.

### HasPortType

`func (o *DnacDeviceInterface) HasPortType() bool`

HasPortType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


