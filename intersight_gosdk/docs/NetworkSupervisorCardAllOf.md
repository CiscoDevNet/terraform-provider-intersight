# NetworkSupervisorCardAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "network.SupervisorCard"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "network.SupervisorCard"]
**Description** | Pointer to **string** | The description of the supervisor card. | [optional] 
**HardwareVersion** | Pointer to **string** | The hardware version of the supervisor card. | [optional] 
**Name** | Pointer to **string** | The name of the supervisor card like Supervisor Card-1. | [optional] 
**NumberOfPorts** | Pointer to **int64** | The number of ports on the supervisor card. | [optional] 
**OperState** | Pointer to **string** | The operational status of the supervisor card. | [optional] 
**PartNumber** | Pointer to **string** | The part number of the supervisor card. | [optional] 
**PowerState** | Pointer to **string** | The power state of the supervisor card. | [optional] 
**Status** | Pointer to **string** | The status of the supervisor card. | [optional] 
**SupervisorId** | Pointer to **string** | The identifier for the supervisor card. | [optional] 
**Type** | Pointer to **string** | The type of the supervisor card. | [optional] 
**FcPorts** | Pointer to [**[]FcPhysicalPortRelationship**](FcPhysicalPortRelationship.md) | An array of relationships to fcPhysicalPort resources. | [optional] 
**NetworkElement** | Pointer to [**NetworkElementRelationship**](NetworkElementRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNetworkSupervisorCardAllOf

`func NewNetworkSupervisorCardAllOf(classId string, objectType string, ) *NetworkSupervisorCardAllOf`

NewNetworkSupervisorCardAllOf instantiates a new NetworkSupervisorCardAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkSupervisorCardAllOfWithDefaults

`func NewNetworkSupervisorCardAllOfWithDefaults() *NetworkSupervisorCardAllOf`

NewNetworkSupervisorCardAllOfWithDefaults instantiates a new NetworkSupervisorCardAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NetworkSupervisorCardAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NetworkSupervisorCardAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NetworkSupervisorCardAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NetworkSupervisorCardAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NetworkSupervisorCardAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NetworkSupervisorCardAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *NetworkSupervisorCardAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkSupervisorCardAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkSupervisorCardAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworkSupervisorCardAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHardwareVersion

`func (o *NetworkSupervisorCardAllOf) GetHardwareVersion() string`

GetHardwareVersion returns the HardwareVersion field if non-nil, zero value otherwise.

### GetHardwareVersionOk

`func (o *NetworkSupervisorCardAllOf) GetHardwareVersionOk() (*string, bool)`

GetHardwareVersionOk returns a tuple with the HardwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareVersion

`func (o *NetworkSupervisorCardAllOf) SetHardwareVersion(v string)`

SetHardwareVersion sets HardwareVersion field to given value.

### HasHardwareVersion

`func (o *NetworkSupervisorCardAllOf) HasHardwareVersion() bool`

HasHardwareVersion returns a boolean if a field has been set.

### GetName

`func (o *NetworkSupervisorCardAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkSupervisorCardAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkSupervisorCardAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkSupervisorCardAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumberOfPorts

`func (o *NetworkSupervisorCardAllOf) GetNumberOfPorts() int64`

GetNumberOfPorts returns the NumberOfPorts field if non-nil, zero value otherwise.

### GetNumberOfPortsOk

`func (o *NetworkSupervisorCardAllOf) GetNumberOfPortsOk() (*int64, bool)`

GetNumberOfPortsOk returns a tuple with the NumberOfPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfPorts

`func (o *NetworkSupervisorCardAllOf) SetNumberOfPorts(v int64)`

SetNumberOfPorts sets NumberOfPorts field to given value.

### HasNumberOfPorts

`func (o *NetworkSupervisorCardAllOf) HasNumberOfPorts() bool`

HasNumberOfPorts returns a boolean if a field has been set.

### GetOperState

`func (o *NetworkSupervisorCardAllOf) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *NetworkSupervisorCardAllOf) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *NetworkSupervisorCardAllOf) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *NetworkSupervisorCardAllOf) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetPartNumber

`func (o *NetworkSupervisorCardAllOf) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *NetworkSupervisorCardAllOf) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *NetworkSupervisorCardAllOf) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *NetworkSupervisorCardAllOf) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetPowerState

`func (o *NetworkSupervisorCardAllOf) GetPowerState() string`

GetPowerState returns the PowerState field if non-nil, zero value otherwise.

### GetPowerStateOk

`func (o *NetworkSupervisorCardAllOf) GetPowerStateOk() (*string, bool)`

GetPowerStateOk returns a tuple with the PowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerState

`func (o *NetworkSupervisorCardAllOf) SetPowerState(v string)`

SetPowerState sets PowerState field to given value.

### HasPowerState

`func (o *NetworkSupervisorCardAllOf) HasPowerState() bool`

HasPowerState returns a boolean if a field has been set.

### GetStatus

`func (o *NetworkSupervisorCardAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkSupervisorCardAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkSupervisorCardAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NetworkSupervisorCardAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSupervisorId

`func (o *NetworkSupervisorCardAllOf) GetSupervisorId() string`

GetSupervisorId returns the SupervisorId field if non-nil, zero value otherwise.

### GetSupervisorIdOk

`func (o *NetworkSupervisorCardAllOf) GetSupervisorIdOk() (*string, bool)`

GetSupervisorIdOk returns a tuple with the SupervisorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupervisorId

`func (o *NetworkSupervisorCardAllOf) SetSupervisorId(v string)`

SetSupervisorId sets SupervisorId field to given value.

### HasSupervisorId

`func (o *NetworkSupervisorCardAllOf) HasSupervisorId() bool`

HasSupervisorId returns a boolean if a field has been set.

### GetType

`func (o *NetworkSupervisorCardAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkSupervisorCardAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkSupervisorCardAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NetworkSupervisorCardAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFcPorts

`func (o *NetworkSupervisorCardAllOf) GetFcPorts() []FcPhysicalPortRelationship`

GetFcPorts returns the FcPorts field if non-nil, zero value otherwise.

### GetFcPortsOk

`func (o *NetworkSupervisorCardAllOf) GetFcPortsOk() (*[]FcPhysicalPortRelationship, bool)`

GetFcPortsOk returns a tuple with the FcPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcPorts

`func (o *NetworkSupervisorCardAllOf) SetFcPorts(v []FcPhysicalPortRelationship)`

SetFcPorts sets FcPorts field to given value.

### HasFcPorts

`func (o *NetworkSupervisorCardAllOf) HasFcPorts() bool`

HasFcPorts returns a boolean if a field has been set.

### SetFcPortsNil

`func (o *NetworkSupervisorCardAllOf) SetFcPortsNil(b bool)`

 SetFcPortsNil sets the value for FcPorts to be an explicit nil

### UnsetFcPorts
`func (o *NetworkSupervisorCardAllOf) UnsetFcPorts()`

UnsetFcPorts ensures that no value is present for FcPorts, not even an explicit nil
### GetNetworkElement

`func (o *NetworkSupervisorCardAllOf) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *NetworkSupervisorCardAllOf) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *NetworkSupervisorCardAllOf) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *NetworkSupervisorCardAllOf) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NetworkSupervisorCardAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NetworkSupervisorCardAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NetworkSupervisorCardAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NetworkSupervisorCardAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


