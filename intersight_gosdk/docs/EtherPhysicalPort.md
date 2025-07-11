# EtherPhysicalPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "ether.PhysicalPort"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "ether.PhysicalPort"]
**AdminSpeed** | Pointer to **string** | Administratively configured speed for this port. | [optional] [readonly] 
**AdminState** | Pointer to **string** | Administratively configured state (enabled/disabled) for this port. | [optional] [readonly] 
**AggregatePortId** | Pointer to **int64** | Breakout port member in the Fabric Interconnect. | [optional] [readonly] 
**LicenseGrace** | Pointer to **string** | The number of days this port&#39;s license has been in Grace Period for. | [optional] [readonly] 
**LicenseState** | Pointer to **string** | The state of the port&#39;s licensing. | [optional] [readonly] 
**MacsecOperData** | Pointer to [**NullableEtherMacsecOperData**](EtherMacsecOperData.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the Physical Port. | [optional] [readonly] 
**UserLabel** | Pointer to **string** | The user defined label assigned to the port. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**NullableInventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**PortGroup** | Pointer to [**NullablePortGroupRelationship**](PortGroupRelationship.md) |  | [optional] 
**PortSubGroup** | Pointer to [**NullablePortSubGroupRelationship**](PortSubGroupRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewEtherPhysicalPort

`func NewEtherPhysicalPort(classId string, objectType string, ) *EtherPhysicalPort`

NewEtherPhysicalPort instantiates a new EtherPhysicalPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEtherPhysicalPortWithDefaults

`func NewEtherPhysicalPortWithDefaults() *EtherPhysicalPort`

NewEtherPhysicalPortWithDefaults instantiates a new EtherPhysicalPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EtherPhysicalPort) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EtherPhysicalPort) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EtherPhysicalPort) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EtherPhysicalPort) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EtherPhysicalPort) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EtherPhysicalPort) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminSpeed

`func (o *EtherPhysicalPort) GetAdminSpeed() string`

GetAdminSpeed returns the AdminSpeed field if non-nil, zero value otherwise.

### GetAdminSpeedOk

`func (o *EtherPhysicalPort) GetAdminSpeedOk() (*string, bool)`

GetAdminSpeedOk returns a tuple with the AdminSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminSpeed

`func (o *EtherPhysicalPort) SetAdminSpeed(v string)`

SetAdminSpeed sets AdminSpeed field to given value.

### HasAdminSpeed

`func (o *EtherPhysicalPort) HasAdminSpeed() bool`

HasAdminSpeed returns a boolean if a field has been set.

### GetAdminState

`func (o *EtherPhysicalPort) GetAdminState() string`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *EtherPhysicalPort) GetAdminStateOk() (*string, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *EtherPhysicalPort) SetAdminState(v string)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *EtherPhysicalPort) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetAggregatePortId

`func (o *EtherPhysicalPort) GetAggregatePortId() int64`

GetAggregatePortId returns the AggregatePortId field if non-nil, zero value otherwise.

### GetAggregatePortIdOk

`func (o *EtherPhysicalPort) GetAggregatePortIdOk() (*int64, bool)`

GetAggregatePortIdOk returns a tuple with the AggregatePortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregatePortId

`func (o *EtherPhysicalPort) SetAggregatePortId(v int64)`

SetAggregatePortId sets AggregatePortId field to given value.

### HasAggregatePortId

`func (o *EtherPhysicalPort) HasAggregatePortId() bool`

HasAggregatePortId returns a boolean if a field has been set.

### GetLicenseGrace

`func (o *EtherPhysicalPort) GetLicenseGrace() string`

GetLicenseGrace returns the LicenseGrace field if non-nil, zero value otherwise.

### GetLicenseGraceOk

`func (o *EtherPhysicalPort) GetLicenseGraceOk() (*string, bool)`

GetLicenseGraceOk returns a tuple with the LicenseGrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseGrace

`func (o *EtherPhysicalPort) SetLicenseGrace(v string)`

SetLicenseGrace sets LicenseGrace field to given value.

### HasLicenseGrace

`func (o *EtherPhysicalPort) HasLicenseGrace() bool`

HasLicenseGrace returns a boolean if a field has been set.

### GetLicenseState

`func (o *EtherPhysicalPort) GetLicenseState() string`

GetLicenseState returns the LicenseState field if non-nil, zero value otherwise.

### GetLicenseStateOk

`func (o *EtherPhysicalPort) GetLicenseStateOk() (*string, bool)`

GetLicenseStateOk returns a tuple with the LicenseState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseState

`func (o *EtherPhysicalPort) SetLicenseState(v string)`

SetLicenseState sets LicenseState field to given value.

### HasLicenseState

`func (o *EtherPhysicalPort) HasLicenseState() bool`

HasLicenseState returns a boolean if a field has been set.

### GetMacsecOperData

`func (o *EtherPhysicalPort) GetMacsecOperData() EtherMacsecOperData`

GetMacsecOperData returns the MacsecOperData field if non-nil, zero value otherwise.

### GetMacsecOperDataOk

`func (o *EtherPhysicalPort) GetMacsecOperDataOk() (*EtherMacsecOperData, bool)`

GetMacsecOperDataOk returns a tuple with the MacsecOperData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacsecOperData

`func (o *EtherPhysicalPort) SetMacsecOperData(v EtherMacsecOperData)`

SetMacsecOperData sets MacsecOperData field to given value.

### HasMacsecOperData

`func (o *EtherPhysicalPort) HasMacsecOperData() bool`

HasMacsecOperData returns a boolean if a field has been set.

### SetMacsecOperDataNil

`func (o *EtherPhysicalPort) SetMacsecOperDataNil(b bool)`

 SetMacsecOperDataNil sets the value for MacsecOperData to be an explicit nil

### UnsetMacsecOperData
`func (o *EtherPhysicalPort) UnsetMacsecOperData()`

UnsetMacsecOperData ensures that no value is present for MacsecOperData, not even an explicit nil
### GetName

`func (o *EtherPhysicalPort) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EtherPhysicalPort) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EtherPhysicalPort) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EtherPhysicalPort) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserLabel

`func (o *EtherPhysicalPort) GetUserLabel() string`

GetUserLabel returns the UserLabel field if non-nil, zero value otherwise.

### GetUserLabelOk

`func (o *EtherPhysicalPort) GetUserLabelOk() (*string, bool)`

GetUserLabelOk returns a tuple with the UserLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLabel

`func (o *EtherPhysicalPort) SetUserLabel(v string)`

SetUserLabel sets UserLabel field to given value.

### HasUserLabel

`func (o *EtherPhysicalPort) HasUserLabel() bool`

HasUserLabel returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *EtherPhysicalPort) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *EtherPhysicalPort) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *EtherPhysicalPort) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *EtherPhysicalPort) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### SetInventoryDeviceInfoNil

`func (o *EtherPhysicalPort) SetInventoryDeviceInfoNil(b bool)`

 SetInventoryDeviceInfoNil sets the value for InventoryDeviceInfo to be an explicit nil

### UnsetInventoryDeviceInfo
`func (o *EtherPhysicalPort) UnsetInventoryDeviceInfo()`

UnsetInventoryDeviceInfo ensures that no value is present for InventoryDeviceInfo, not even an explicit nil
### GetPortGroup

`func (o *EtherPhysicalPort) GetPortGroup() PortGroupRelationship`

GetPortGroup returns the PortGroup field if non-nil, zero value otherwise.

### GetPortGroupOk

`func (o *EtherPhysicalPort) GetPortGroupOk() (*PortGroupRelationship, bool)`

GetPortGroupOk returns a tuple with the PortGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortGroup

`func (o *EtherPhysicalPort) SetPortGroup(v PortGroupRelationship)`

SetPortGroup sets PortGroup field to given value.

### HasPortGroup

`func (o *EtherPhysicalPort) HasPortGroup() bool`

HasPortGroup returns a boolean if a field has been set.

### SetPortGroupNil

`func (o *EtherPhysicalPort) SetPortGroupNil(b bool)`

 SetPortGroupNil sets the value for PortGroup to be an explicit nil

### UnsetPortGroup
`func (o *EtherPhysicalPort) UnsetPortGroup()`

UnsetPortGroup ensures that no value is present for PortGroup, not even an explicit nil
### GetPortSubGroup

`func (o *EtherPhysicalPort) GetPortSubGroup() PortSubGroupRelationship`

GetPortSubGroup returns the PortSubGroup field if non-nil, zero value otherwise.

### GetPortSubGroupOk

`func (o *EtherPhysicalPort) GetPortSubGroupOk() (*PortSubGroupRelationship, bool)`

GetPortSubGroupOk returns a tuple with the PortSubGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortSubGroup

`func (o *EtherPhysicalPort) SetPortSubGroup(v PortSubGroupRelationship)`

SetPortSubGroup sets PortSubGroup field to given value.

### HasPortSubGroup

`func (o *EtherPhysicalPort) HasPortSubGroup() bool`

HasPortSubGroup returns a boolean if a field has been set.

### SetPortSubGroupNil

`func (o *EtherPhysicalPort) SetPortSubGroupNil(b bool)`

 SetPortSubGroupNil sets the value for PortSubGroup to be an explicit nil

### UnsetPortSubGroup
`func (o *EtherPhysicalPort) UnsetPortSubGroup()`

UnsetPortSubGroup ensures that no value is present for PortSubGroup, not even an explicit nil
### GetRegisteredDevice

`func (o *EtherPhysicalPort) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EtherPhysicalPort) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EtherPhysicalPort) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EtherPhysicalPort) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *EtherPhysicalPort) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *EtherPhysicalPort) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


