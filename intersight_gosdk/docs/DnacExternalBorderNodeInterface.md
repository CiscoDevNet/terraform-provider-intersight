# DnacExternalBorderNodeInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "dnac.ExternalBorderNodeInterface"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "dnac.ExternalBorderNodeInterface"]
**AdminStatus** | Pointer to **string** | Administrator status for external border node interface. | [optional] 
**ExternalBorderNodeId** | Pointer to **string** | External border node&#39;s id. | [optional] 
**InterfaceId** | Pointer to **string** | The Moid for the interface in the external border node. | [optional] 
**InterfaceName** | Pointer to **string** | The name for the external border node&#39;s interface. | [optional] 
**InterfaceType** | Pointer to **string** | Interface type for external border node interface. | [optional] 
**PortType** | Pointer to **string** | Port type for external border node interface. | [optional] 

## Methods

### NewDnacExternalBorderNodeInterface

`func NewDnacExternalBorderNodeInterface(classId string, objectType string, ) *DnacExternalBorderNodeInterface`

NewDnacExternalBorderNodeInterface instantiates a new DnacExternalBorderNodeInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnacExternalBorderNodeInterfaceWithDefaults

`func NewDnacExternalBorderNodeInterfaceWithDefaults() *DnacExternalBorderNodeInterface`

NewDnacExternalBorderNodeInterfaceWithDefaults instantiates a new DnacExternalBorderNodeInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *DnacExternalBorderNodeInterface) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *DnacExternalBorderNodeInterface) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *DnacExternalBorderNodeInterface) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *DnacExternalBorderNodeInterface) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *DnacExternalBorderNodeInterface) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *DnacExternalBorderNodeInterface) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminStatus

`func (o *DnacExternalBorderNodeInterface) GetAdminStatus() string`

GetAdminStatus returns the AdminStatus field if non-nil, zero value otherwise.

### GetAdminStatusOk

`func (o *DnacExternalBorderNodeInterface) GetAdminStatusOk() (*string, bool)`

GetAdminStatusOk returns a tuple with the AdminStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminStatus

`func (o *DnacExternalBorderNodeInterface) SetAdminStatus(v string)`

SetAdminStatus sets AdminStatus field to given value.

### HasAdminStatus

`func (o *DnacExternalBorderNodeInterface) HasAdminStatus() bool`

HasAdminStatus returns a boolean if a field has been set.

### GetExternalBorderNodeId

`func (o *DnacExternalBorderNodeInterface) GetExternalBorderNodeId() string`

GetExternalBorderNodeId returns the ExternalBorderNodeId field if non-nil, zero value otherwise.

### GetExternalBorderNodeIdOk

`func (o *DnacExternalBorderNodeInterface) GetExternalBorderNodeIdOk() (*string, bool)`

GetExternalBorderNodeIdOk returns a tuple with the ExternalBorderNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalBorderNodeId

`func (o *DnacExternalBorderNodeInterface) SetExternalBorderNodeId(v string)`

SetExternalBorderNodeId sets ExternalBorderNodeId field to given value.

### HasExternalBorderNodeId

`func (o *DnacExternalBorderNodeInterface) HasExternalBorderNodeId() bool`

HasExternalBorderNodeId returns a boolean if a field has been set.

### GetInterfaceId

`func (o *DnacExternalBorderNodeInterface) GetInterfaceId() string`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *DnacExternalBorderNodeInterface) GetInterfaceIdOk() (*string, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *DnacExternalBorderNodeInterface) SetInterfaceId(v string)`

SetInterfaceId sets InterfaceId field to given value.

### HasInterfaceId

`func (o *DnacExternalBorderNodeInterface) HasInterfaceId() bool`

HasInterfaceId returns a boolean if a field has been set.

### GetInterfaceName

`func (o *DnacExternalBorderNodeInterface) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *DnacExternalBorderNodeInterface) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *DnacExternalBorderNodeInterface) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.

### HasInterfaceName

`func (o *DnacExternalBorderNodeInterface) HasInterfaceName() bool`

HasInterfaceName returns a boolean if a field has been set.

### GetInterfaceType

`func (o *DnacExternalBorderNodeInterface) GetInterfaceType() string`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *DnacExternalBorderNodeInterface) GetInterfaceTypeOk() (*string, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *DnacExternalBorderNodeInterface) SetInterfaceType(v string)`

SetInterfaceType sets InterfaceType field to given value.

### HasInterfaceType

`func (o *DnacExternalBorderNodeInterface) HasInterfaceType() bool`

HasInterfaceType returns a boolean if a field has been set.

### GetPortType

`func (o *DnacExternalBorderNodeInterface) GetPortType() string`

GetPortType returns the PortType field if non-nil, zero value otherwise.

### GetPortTypeOk

`func (o *DnacExternalBorderNodeInterface) GetPortTypeOk() (*string, bool)`

GetPortTypeOk returns a tuple with the PortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortType

`func (o *DnacExternalBorderNodeInterface) SetPortType(v string)`

SetPortType sets PortType field to given value.

### HasPortType

`func (o *DnacExternalBorderNodeInterface) HasPortType() bool`

HasPortType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


