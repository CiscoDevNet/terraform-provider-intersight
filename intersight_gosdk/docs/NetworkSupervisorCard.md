# NetworkSupervisorCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "network.SupervisorCard"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "network.SupervisorCard"]
**Description** | Pointer to **string** | The description of the supervisor card. | [optional] [readonly] 
**HardwareVersion** | Pointer to **string** | The hardware version of the supervisor card. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the supervisor card like Supervisor Card-1. | [optional] [readonly] 
**NumberOfPorts** | Pointer to **int64** | The number of ports on the supervisor card. | [optional] [readonly] 
**OperReason** | Pointer to **string** | The health issue of the supervisor card. | [optional] [readonly] 
**OperState** | Pointer to **string** | The operational status of the supervisor card. | [optional] [readonly] 
**PartNumber** | Pointer to **string** | The part number of the supervisor card. | [optional] [readonly] 
**PowerState** | Pointer to **string** | The power state of the supervisor card. | [optional] [readonly] 
**Status** | Pointer to **string** | The status of the supervisor card. | [optional] [readonly] 
**SupervisorId** | Pointer to **string** | The identifier for the supervisor card. | [optional] [readonly] 
**Type** | Pointer to **string** | The type of the supervisor card. | [optional] [readonly] 
**FcPorts** | Pointer to [**[]FcPhysicalPortRelationship**](FcPhysicalPortRelationship.md) | An array of relationships to fcPhysicalPort resources. | [optional] [readonly] 
**NetworkElement** | Pointer to [**NetworkElementRelationship**](NetworkElementRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNetworkSupervisorCard

`func NewNetworkSupervisorCard(classId string, objectType string, ) *NetworkSupervisorCard`

NewNetworkSupervisorCard instantiates a new NetworkSupervisorCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkSupervisorCardWithDefaults

`func NewNetworkSupervisorCardWithDefaults() *NetworkSupervisorCard`

NewNetworkSupervisorCardWithDefaults instantiates a new NetworkSupervisorCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NetworkSupervisorCard) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NetworkSupervisorCard) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NetworkSupervisorCard) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NetworkSupervisorCard) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NetworkSupervisorCard) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NetworkSupervisorCard) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *NetworkSupervisorCard) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkSupervisorCard) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkSupervisorCard) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworkSupervisorCard) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHardwareVersion

`func (o *NetworkSupervisorCard) GetHardwareVersion() string`

GetHardwareVersion returns the HardwareVersion field if non-nil, zero value otherwise.

### GetHardwareVersionOk

`func (o *NetworkSupervisorCard) GetHardwareVersionOk() (*string, bool)`

GetHardwareVersionOk returns a tuple with the HardwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareVersion

`func (o *NetworkSupervisorCard) SetHardwareVersion(v string)`

SetHardwareVersion sets HardwareVersion field to given value.

### HasHardwareVersion

`func (o *NetworkSupervisorCard) HasHardwareVersion() bool`

HasHardwareVersion returns a boolean if a field has been set.

### GetName

`func (o *NetworkSupervisorCard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkSupervisorCard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkSupervisorCard) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkSupervisorCard) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumberOfPorts

`func (o *NetworkSupervisorCard) GetNumberOfPorts() int64`

GetNumberOfPorts returns the NumberOfPorts field if non-nil, zero value otherwise.

### GetNumberOfPortsOk

`func (o *NetworkSupervisorCard) GetNumberOfPortsOk() (*int64, bool)`

GetNumberOfPortsOk returns a tuple with the NumberOfPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfPorts

`func (o *NetworkSupervisorCard) SetNumberOfPorts(v int64)`

SetNumberOfPorts sets NumberOfPorts field to given value.

### HasNumberOfPorts

`func (o *NetworkSupervisorCard) HasNumberOfPorts() bool`

HasNumberOfPorts returns a boolean if a field has been set.

### GetOperReason

`func (o *NetworkSupervisorCard) GetOperReason() string`

GetOperReason returns the OperReason field if non-nil, zero value otherwise.

### GetOperReasonOk

`func (o *NetworkSupervisorCard) GetOperReasonOk() (*string, bool)`

GetOperReasonOk returns a tuple with the OperReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperReason

`func (o *NetworkSupervisorCard) SetOperReason(v string)`

SetOperReason sets OperReason field to given value.

### HasOperReason

`func (o *NetworkSupervisorCard) HasOperReason() bool`

HasOperReason returns a boolean if a field has been set.

### GetOperState

`func (o *NetworkSupervisorCard) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *NetworkSupervisorCard) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *NetworkSupervisorCard) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *NetworkSupervisorCard) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetPartNumber

`func (o *NetworkSupervisorCard) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *NetworkSupervisorCard) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *NetworkSupervisorCard) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *NetworkSupervisorCard) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetPowerState

`func (o *NetworkSupervisorCard) GetPowerState() string`

GetPowerState returns the PowerState field if non-nil, zero value otherwise.

### GetPowerStateOk

`func (o *NetworkSupervisorCard) GetPowerStateOk() (*string, bool)`

GetPowerStateOk returns a tuple with the PowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerState

`func (o *NetworkSupervisorCard) SetPowerState(v string)`

SetPowerState sets PowerState field to given value.

### HasPowerState

`func (o *NetworkSupervisorCard) HasPowerState() bool`

HasPowerState returns a boolean if a field has been set.

### GetStatus

`func (o *NetworkSupervisorCard) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkSupervisorCard) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkSupervisorCard) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NetworkSupervisorCard) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSupervisorId

`func (o *NetworkSupervisorCard) GetSupervisorId() string`

GetSupervisorId returns the SupervisorId field if non-nil, zero value otherwise.

### GetSupervisorIdOk

`func (o *NetworkSupervisorCard) GetSupervisorIdOk() (*string, bool)`

GetSupervisorIdOk returns a tuple with the SupervisorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupervisorId

`func (o *NetworkSupervisorCard) SetSupervisorId(v string)`

SetSupervisorId sets SupervisorId field to given value.

### HasSupervisorId

`func (o *NetworkSupervisorCard) HasSupervisorId() bool`

HasSupervisorId returns a boolean if a field has been set.

### GetType

`func (o *NetworkSupervisorCard) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkSupervisorCard) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkSupervisorCard) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NetworkSupervisorCard) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFcPorts

`func (o *NetworkSupervisorCard) GetFcPorts() []FcPhysicalPortRelationship`

GetFcPorts returns the FcPorts field if non-nil, zero value otherwise.

### GetFcPortsOk

`func (o *NetworkSupervisorCard) GetFcPortsOk() (*[]FcPhysicalPortRelationship, bool)`

GetFcPortsOk returns a tuple with the FcPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcPorts

`func (o *NetworkSupervisorCard) SetFcPorts(v []FcPhysicalPortRelationship)`

SetFcPorts sets FcPorts field to given value.

### HasFcPorts

`func (o *NetworkSupervisorCard) HasFcPorts() bool`

HasFcPorts returns a boolean if a field has been set.

### SetFcPortsNil

`func (o *NetworkSupervisorCard) SetFcPortsNil(b bool)`

 SetFcPortsNil sets the value for FcPorts to be an explicit nil

### UnsetFcPorts
`func (o *NetworkSupervisorCard) UnsetFcPorts()`

UnsetFcPorts ensures that no value is present for FcPorts, not even an explicit nil
### GetNetworkElement

`func (o *NetworkSupervisorCard) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *NetworkSupervisorCard) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *NetworkSupervisorCard) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *NetworkSupervisorCard) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NetworkSupervisorCard) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NetworkSupervisorCard) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NetworkSupervisorCard) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NetworkSupervisorCard) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


