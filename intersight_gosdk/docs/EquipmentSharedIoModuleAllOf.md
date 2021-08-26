# EquipmentSharedIoModuleAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "equipment.SharedIoModule"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "equipment.SharedIoModule"]
**ConfigState** | Pointer to **string** | This field identifies the configuration state for this SIOM Unit. | [optional] [readonly] 
**Discovery** | Pointer to **string** | This field identifies the discovery state of SIOM. | [optional] [readonly] 
**MacOfSharedIomAside** | Pointer to **string** | This field identifies the MAC of IOM-A side. | [optional] [readonly] 
**MacOfSharedIomBside** | Pointer to **string** | This field identifies the MAC of IOM-B side. | [optional] [readonly] 
**OperState** | Pointer to **string** | This field identifies the SIOM operational state. | [optional] [readonly] 
**PartNumber** | Pointer to **string** | This field identifies the Part Number for this SIOM Unit. | [optional] [readonly] 
**Reachability** | Pointer to **string** | This field identifies the reachability to FI-A and B side. | [optional] [readonly] 
**UsrLbl** | Pointer to **string** | User label configured for the SIOM. | [optional] [readonly] 
**Vid** | Pointer to **string** | This field identifies the vendor id for this SIOM Unit. | [optional] [readonly] 
**Controller** | Pointer to [**ManagementControllerRelationship**](ManagementControllerRelationship.md) |  | [optional] 
**EquipmentSystemIoController** | Pointer to [**EquipmentSystemIoControllerRelationship**](EquipmentSystemIoControllerRelationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**PortGroups** | Pointer to [**[]PortGroupRelationship**](PortGroupRelationship.md) | An array of relationships to portGroup resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewEquipmentSharedIoModuleAllOf

`func NewEquipmentSharedIoModuleAllOf(classId string, objectType string, ) *EquipmentSharedIoModuleAllOf`

NewEquipmentSharedIoModuleAllOf instantiates a new EquipmentSharedIoModuleAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentSharedIoModuleAllOfWithDefaults

`func NewEquipmentSharedIoModuleAllOfWithDefaults() *EquipmentSharedIoModuleAllOf`

NewEquipmentSharedIoModuleAllOfWithDefaults instantiates a new EquipmentSharedIoModuleAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EquipmentSharedIoModuleAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EquipmentSharedIoModuleAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EquipmentSharedIoModuleAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EquipmentSharedIoModuleAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EquipmentSharedIoModuleAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EquipmentSharedIoModuleAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConfigState

`func (o *EquipmentSharedIoModuleAllOf) GetConfigState() string`

GetConfigState returns the ConfigState field if non-nil, zero value otherwise.

### GetConfigStateOk

`func (o *EquipmentSharedIoModuleAllOf) GetConfigStateOk() (*string, bool)`

GetConfigStateOk returns a tuple with the ConfigState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigState

`func (o *EquipmentSharedIoModuleAllOf) SetConfigState(v string)`

SetConfigState sets ConfigState field to given value.

### HasConfigState

`func (o *EquipmentSharedIoModuleAllOf) HasConfigState() bool`

HasConfigState returns a boolean if a field has been set.

### GetDiscovery

`func (o *EquipmentSharedIoModuleAllOf) GetDiscovery() string`

GetDiscovery returns the Discovery field if non-nil, zero value otherwise.

### GetDiscoveryOk

`func (o *EquipmentSharedIoModuleAllOf) GetDiscoveryOk() (*string, bool)`

GetDiscoveryOk returns a tuple with the Discovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscovery

`func (o *EquipmentSharedIoModuleAllOf) SetDiscovery(v string)`

SetDiscovery sets Discovery field to given value.

### HasDiscovery

`func (o *EquipmentSharedIoModuleAllOf) HasDiscovery() bool`

HasDiscovery returns a boolean if a field has been set.

### GetMacOfSharedIomAside

`func (o *EquipmentSharedIoModuleAllOf) GetMacOfSharedIomAside() string`

GetMacOfSharedIomAside returns the MacOfSharedIomAside field if non-nil, zero value otherwise.

### GetMacOfSharedIomAsideOk

`func (o *EquipmentSharedIoModuleAllOf) GetMacOfSharedIomAsideOk() (*string, bool)`

GetMacOfSharedIomAsideOk returns a tuple with the MacOfSharedIomAside field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacOfSharedIomAside

`func (o *EquipmentSharedIoModuleAllOf) SetMacOfSharedIomAside(v string)`

SetMacOfSharedIomAside sets MacOfSharedIomAside field to given value.

### HasMacOfSharedIomAside

`func (o *EquipmentSharedIoModuleAllOf) HasMacOfSharedIomAside() bool`

HasMacOfSharedIomAside returns a boolean if a field has been set.

### GetMacOfSharedIomBside

`func (o *EquipmentSharedIoModuleAllOf) GetMacOfSharedIomBside() string`

GetMacOfSharedIomBside returns the MacOfSharedIomBside field if non-nil, zero value otherwise.

### GetMacOfSharedIomBsideOk

`func (o *EquipmentSharedIoModuleAllOf) GetMacOfSharedIomBsideOk() (*string, bool)`

GetMacOfSharedIomBsideOk returns a tuple with the MacOfSharedIomBside field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacOfSharedIomBside

`func (o *EquipmentSharedIoModuleAllOf) SetMacOfSharedIomBside(v string)`

SetMacOfSharedIomBside sets MacOfSharedIomBside field to given value.

### HasMacOfSharedIomBside

`func (o *EquipmentSharedIoModuleAllOf) HasMacOfSharedIomBside() bool`

HasMacOfSharedIomBside returns a boolean if a field has been set.

### GetOperState

`func (o *EquipmentSharedIoModuleAllOf) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *EquipmentSharedIoModuleAllOf) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *EquipmentSharedIoModuleAllOf) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *EquipmentSharedIoModuleAllOf) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetPartNumber

`func (o *EquipmentSharedIoModuleAllOf) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *EquipmentSharedIoModuleAllOf) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *EquipmentSharedIoModuleAllOf) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *EquipmentSharedIoModuleAllOf) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetReachability

`func (o *EquipmentSharedIoModuleAllOf) GetReachability() string`

GetReachability returns the Reachability field if non-nil, zero value otherwise.

### GetReachabilityOk

`func (o *EquipmentSharedIoModuleAllOf) GetReachabilityOk() (*string, bool)`

GetReachabilityOk returns a tuple with the Reachability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachability

`func (o *EquipmentSharedIoModuleAllOf) SetReachability(v string)`

SetReachability sets Reachability field to given value.

### HasReachability

`func (o *EquipmentSharedIoModuleAllOf) HasReachability() bool`

HasReachability returns a boolean if a field has been set.

### GetUsrLbl

`func (o *EquipmentSharedIoModuleAllOf) GetUsrLbl() string`

GetUsrLbl returns the UsrLbl field if non-nil, zero value otherwise.

### GetUsrLblOk

`func (o *EquipmentSharedIoModuleAllOf) GetUsrLblOk() (*string, bool)`

GetUsrLblOk returns a tuple with the UsrLbl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsrLbl

`func (o *EquipmentSharedIoModuleAllOf) SetUsrLbl(v string)`

SetUsrLbl sets UsrLbl field to given value.

### HasUsrLbl

`func (o *EquipmentSharedIoModuleAllOf) HasUsrLbl() bool`

HasUsrLbl returns a boolean if a field has been set.

### GetVid

`func (o *EquipmentSharedIoModuleAllOf) GetVid() string`

GetVid returns the Vid field if non-nil, zero value otherwise.

### GetVidOk

`func (o *EquipmentSharedIoModuleAllOf) GetVidOk() (*string, bool)`

GetVidOk returns a tuple with the Vid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVid

`func (o *EquipmentSharedIoModuleAllOf) SetVid(v string)`

SetVid sets Vid field to given value.

### HasVid

`func (o *EquipmentSharedIoModuleAllOf) HasVid() bool`

HasVid returns a boolean if a field has been set.

### GetController

`func (o *EquipmentSharedIoModuleAllOf) GetController() ManagementControllerRelationship`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *EquipmentSharedIoModuleAllOf) GetControllerOk() (*ManagementControllerRelationship, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *EquipmentSharedIoModuleAllOf) SetController(v ManagementControllerRelationship)`

SetController sets Controller field to given value.

### HasController

`func (o *EquipmentSharedIoModuleAllOf) HasController() bool`

HasController returns a boolean if a field has been set.

### GetEquipmentSystemIoController

`func (o *EquipmentSharedIoModuleAllOf) GetEquipmentSystemIoController() EquipmentSystemIoControllerRelationship`

GetEquipmentSystemIoController returns the EquipmentSystemIoController field if non-nil, zero value otherwise.

### GetEquipmentSystemIoControllerOk

`func (o *EquipmentSharedIoModuleAllOf) GetEquipmentSystemIoControllerOk() (*EquipmentSystemIoControllerRelationship, bool)`

GetEquipmentSystemIoControllerOk returns a tuple with the EquipmentSystemIoController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentSystemIoController

`func (o *EquipmentSharedIoModuleAllOf) SetEquipmentSystemIoController(v EquipmentSystemIoControllerRelationship)`

SetEquipmentSystemIoController sets EquipmentSystemIoController field to given value.

### HasEquipmentSystemIoController

`func (o *EquipmentSharedIoModuleAllOf) HasEquipmentSystemIoController() bool`

HasEquipmentSystemIoController returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *EquipmentSharedIoModuleAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *EquipmentSharedIoModuleAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *EquipmentSharedIoModuleAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *EquipmentSharedIoModuleAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetPortGroups

`func (o *EquipmentSharedIoModuleAllOf) GetPortGroups() []PortGroupRelationship`

GetPortGroups returns the PortGroups field if non-nil, zero value otherwise.

### GetPortGroupsOk

`func (o *EquipmentSharedIoModuleAllOf) GetPortGroupsOk() (*[]PortGroupRelationship, bool)`

GetPortGroupsOk returns a tuple with the PortGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortGroups

`func (o *EquipmentSharedIoModuleAllOf) SetPortGroups(v []PortGroupRelationship)`

SetPortGroups sets PortGroups field to given value.

### HasPortGroups

`func (o *EquipmentSharedIoModuleAllOf) HasPortGroups() bool`

HasPortGroups returns a boolean if a field has been set.

### SetPortGroupsNil

`func (o *EquipmentSharedIoModuleAllOf) SetPortGroupsNil(b bool)`

 SetPortGroupsNil sets the value for PortGroups to be an explicit nil

### UnsetPortGroups
`func (o *EquipmentSharedIoModuleAllOf) UnsetPortGroups()`

UnsetPortGroups ensures that no value is present for PortGroups, not even an explicit nil
### GetRegisteredDevice

`func (o *EquipmentSharedIoModuleAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EquipmentSharedIoModuleAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EquipmentSharedIoModuleAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EquipmentSharedIoModuleAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


