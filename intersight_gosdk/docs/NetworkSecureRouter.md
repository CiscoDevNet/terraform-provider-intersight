# NetworkSecureRouter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "network.SecureRouter"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "network.SecureRouter"]
**AlarmSummary** | Pointer to [**NullableComputeAlarmSummary**](ComputeAlarmSummary.md) |  | [optional] 
**BootProgressLastState** | Pointer to **string** | The last reported boot progress of the secure router. * &#x60;Unknown&#x60; - The boot progress state of the system is unknown. * &#x60;None&#x60; - The system has not started the boot process. * &#x60;OSBootStarted&#x60; - The system has started the OS boot process. * &#x60;OSRunning&#x60; - The system has completed the boot process and the OS is running. | [optional] [readonly] [default to "Unknown"]
**ChassisId** | Pointer to **int64** | The id of the chassis that the blade is discovered in. | [optional] [readonly] 
**FirmwareVersion** | Pointer to **string** | The version of the secure router CIMC firmware. | [optional] [readonly] 
**Lifecycle** | Pointer to **string** | The lifecycle state of the secure router. This will map to the discovery lifecycle as represented in the Identity object. * &#x60;None&#x60; - Default state of an equipment. This should be an initial state when no state is defined for an equipment. * &#x60;Active&#x60; - Default Lifecycle State for a physical entity. * &#x60;Decommissioned&#x60; - Decommission Lifecycle state. * &#x60;SlotReserved&#x60; - Blade is inserted in a slot that is reserved for Secure Router. * &#x60;DiscoveryInProgress&#x60; - DiscoveryInProgress Lifecycle state. * &#x60;DiscoveryFailed&#x60; - DiscoveryFailed Lifecycle state. * &#x60;FirmwareUpgradeInProgress&#x60; - Firmware upgrade is in progress on given physical entity. * &#x60;DiagnosticsInProgress&#x60; - Diagnostics is in progress on given physical entity. * &#x60;SecureEraseInProgress&#x60; - Secure Erase is in progress on given physical entity. * &#x60;ScrubInProgress&#x60; - Scrub is in progress on given physical entity. * &#x60;BladeMigrationInProgress&#x60; - Server slot migration is in progress on given physical entity. * &#x60;SlotMismatch&#x60; - The blade server is detected in a different chassis/slot than it was previously. * &#x60;Removed&#x60; - The blade server has been removed from its discovered slot, and not detected anywhere else. Blade inventory can be cleaned up by performing a software remove operation on the physically removed blade. * &#x60;Moved&#x60; - The blade server has been moved from its discovered location to a new location. Blade inventory can be updated by performing a rediscover operation on the moved blade. * &#x60;Replaced&#x60; - The blade server has been removed from its discovered location and another blade has been inserted in that location. Blade inventory can be cleaned up and updated by doing a software remove operation on the physically removed blade. * &#x60;MovedAndReplaced&#x60; - The blade server has been moved from its discovered location to a new location and another blade has been inserted into the old discovered location. Blade inventory can be updated by performing a rediscover operation on the moved blade. * &#x60;DecommissionAndRemoveInProgress&#x60; - Decommission and remove operation is in progress. * &#x60;DecommissionedForRemove&#x60; - Server is decommissioned and once it is physically pulled from chassis, it will be automatically removed from inventory. | [optional] [readonly] [default to "None"]
**Name** | Pointer to **string** | The name of the secure router. The value is set to the UCSFI domain name along with the chassis and router id. | [optional] [readonly] 
**OperPowerState** | Pointer to **string** | The power state of the secure router. | [optional] [readonly] 
**OperReason** | Pointer to **[]string** |  | [optional] 
**PackageVersion** | Pointer to **string** | Bundle version which the CIMC firmware belongs to. | [optional] [readonly] 
**SlotId** | Pointer to **int64** | The slot number in the chassis that the blade is discovered in. | [optional] [readonly] 
**UserLabel** | Pointer to **string** | The user defined label assigned to the secure router. | [optional] 
**Bmc** | Pointer to [**NullableManagementControllerRelationship**](ManagementControllerRelationship.md) |  | [optional] 
**EquipmentChassis** | Pointer to [**NullableEquipmentChassisRelationship**](EquipmentChassisRelationship.md) |  | [optional] 
**LanPorts** | Pointer to [**[]EtherLanPortRelationship**](EtherLanPortRelationship.md) | An array of relationships to etherLanPort resources. | [optional] [readonly] 
**LocatorLed** | Pointer to [**NullableEquipmentLocatorLedRelationship**](EquipmentLocatorLedRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNetworkSecureRouter

`func NewNetworkSecureRouter(classId string, objectType string, ) *NetworkSecureRouter`

NewNetworkSecureRouter instantiates a new NetworkSecureRouter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkSecureRouterWithDefaults

`func NewNetworkSecureRouterWithDefaults() *NetworkSecureRouter`

NewNetworkSecureRouterWithDefaults instantiates a new NetworkSecureRouter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NetworkSecureRouter) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NetworkSecureRouter) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NetworkSecureRouter) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NetworkSecureRouter) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NetworkSecureRouter) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NetworkSecureRouter) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAlarmSummary

`func (o *NetworkSecureRouter) GetAlarmSummary() ComputeAlarmSummary`

GetAlarmSummary returns the AlarmSummary field if non-nil, zero value otherwise.

### GetAlarmSummaryOk

`func (o *NetworkSecureRouter) GetAlarmSummaryOk() (*ComputeAlarmSummary, bool)`

GetAlarmSummaryOk returns a tuple with the AlarmSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmSummary

`func (o *NetworkSecureRouter) SetAlarmSummary(v ComputeAlarmSummary)`

SetAlarmSummary sets AlarmSummary field to given value.

### HasAlarmSummary

`func (o *NetworkSecureRouter) HasAlarmSummary() bool`

HasAlarmSummary returns a boolean if a field has been set.

### SetAlarmSummaryNil

`func (o *NetworkSecureRouter) SetAlarmSummaryNil(b bool)`

 SetAlarmSummaryNil sets the value for AlarmSummary to be an explicit nil

### UnsetAlarmSummary
`func (o *NetworkSecureRouter) UnsetAlarmSummary()`

UnsetAlarmSummary ensures that no value is present for AlarmSummary, not even an explicit nil
### GetBootProgressLastState

`func (o *NetworkSecureRouter) GetBootProgressLastState() string`

GetBootProgressLastState returns the BootProgressLastState field if non-nil, zero value otherwise.

### GetBootProgressLastStateOk

`func (o *NetworkSecureRouter) GetBootProgressLastStateOk() (*string, bool)`

GetBootProgressLastStateOk returns a tuple with the BootProgressLastState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootProgressLastState

`func (o *NetworkSecureRouter) SetBootProgressLastState(v string)`

SetBootProgressLastState sets BootProgressLastState field to given value.

### HasBootProgressLastState

`func (o *NetworkSecureRouter) HasBootProgressLastState() bool`

HasBootProgressLastState returns a boolean if a field has been set.

### GetChassisId

`func (o *NetworkSecureRouter) GetChassisId() int64`

GetChassisId returns the ChassisId field if non-nil, zero value otherwise.

### GetChassisIdOk

`func (o *NetworkSecureRouter) GetChassisIdOk() (*int64, bool)`

GetChassisIdOk returns a tuple with the ChassisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisId

`func (o *NetworkSecureRouter) SetChassisId(v int64)`

SetChassisId sets ChassisId field to given value.

### HasChassisId

`func (o *NetworkSecureRouter) HasChassisId() bool`

HasChassisId returns a boolean if a field has been set.

### GetFirmwareVersion

`func (o *NetworkSecureRouter) GetFirmwareVersion() string`

GetFirmwareVersion returns the FirmwareVersion field if non-nil, zero value otherwise.

### GetFirmwareVersionOk

`func (o *NetworkSecureRouter) GetFirmwareVersionOk() (*string, bool)`

GetFirmwareVersionOk returns a tuple with the FirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareVersion

`func (o *NetworkSecureRouter) SetFirmwareVersion(v string)`

SetFirmwareVersion sets FirmwareVersion field to given value.

### HasFirmwareVersion

`func (o *NetworkSecureRouter) HasFirmwareVersion() bool`

HasFirmwareVersion returns a boolean if a field has been set.

### GetLifecycle

`func (o *NetworkSecureRouter) GetLifecycle() string`

GetLifecycle returns the Lifecycle field if non-nil, zero value otherwise.

### GetLifecycleOk

`func (o *NetworkSecureRouter) GetLifecycleOk() (*string, bool)`

GetLifecycleOk returns a tuple with the Lifecycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycle

`func (o *NetworkSecureRouter) SetLifecycle(v string)`

SetLifecycle sets Lifecycle field to given value.

### HasLifecycle

`func (o *NetworkSecureRouter) HasLifecycle() bool`

HasLifecycle returns a boolean if a field has been set.

### GetName

`func (o *NetworkSecureRouter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkSecureRouter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkSecureRouter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkSecureRouter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperPowerState

`func (o *NetworkSecureRouter) GetOperPowerState() string`

GetOperPowerState returns the OperPowerState field if non-nil, zero value otherwise.

### GetOperPowerStateOk

`func (o *NetworkSecureRouter) GetOperPowerStateOk() (*string, bool)`

GetOperPowerStateOk returns a tuple with the OperPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperPowerState

`func (o *NetworkSecureRouter) SetOperPowerState(v string)`

SetOperPowerState sets OperPowerState field to given value.

### HasOperPowerState

`func (o *NetworkSecureRouter) HasOperPowerState() bool`

HasOperPowerState returns a boolean if a field has been set.

### GetOperReason

`func (o *NetworkSecureRouter) GetOperReason() []string`

GetOperReason returns the OperReason field if non-nil, zero value otherwise.

### GetOperReasonOk

`func (o *NetworkSecureRouter) GetOperReasonOk() (*[]string, bool)`

GetOperReasonOk returns a tuple with the OperReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperReason

`func (o *NetworkSecureRouter) SetOperReason(v []string)`

SetOperReason sets OperReason field to given value.

### HasOperReason

`func (o *NetworkSecureRouter) HasOperReason() bool`

HasOperReason returns a boolean if a field has been set.

### SetOperReasonNil

`func (o *NetworkSecureRouter) SetOperReasonNil(b bool)`

 SetOperReasonNil sets the value for OperReason to be an explicit nil

### UnsetOperReason
`func (o *NetworkSecureRouter) UnsetOperReason()`

UnsetOperReason ensures that no value is present for OperReason, not even an explicit nil
### GetPackageVersion

`func (o *NetworkSecureRouter) GetPackageVersion() string`

GetPackageVersion returns the PackageVersion field if non-nil, zero value otherwise.

### GetPackageVersionOk

`func (o *NetworkSecureRouter) GetPackageVersionOk() (*string, bool)`

GetPackageVersionOk returns a tuple with the PackageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageVersion

`func (o *NetworkSecureRouter) SetPackageVersion(v string)`

SetPackageVersion sets PackageVersion field to given value.

### HasPackageVersion

`func (o *NetworkSecureRouter) HasPackageVersion() bool`

HasPackageVersion returns a boolean if a field has been set.

### GetSlotId

`func (o *NetworkSecureRouter) GetSlotId() int64`

GetSlotId returns the SlotId field if non-nil, zero value otherwise.

### GetSlotIdOk

`func (o *NetworkSecureRouter) GetSlotIdOk() (*int64, bool)`

GetSlotIdOk returns a tuple with the SlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotId

`func (o *NetworkSecureRouter) SetSlotId(v int64)`

SetSlotId sets SlotId field to given value.

### HasSlotId

`func (o *NetworkSecureRouter) HasSlotId() bool`

HasSlotId returns a boolean if a field has been set.

### GetUserLabel

`func (o *NetworkSecureRouter) GetUserLabel() string`

GetUserLabel returns the UserLabel field if non-nil, zero value otherwise.

### GetUserLabelOk

`func (o *NetworkSecureRouter) GetUserLabelOk() (*string, bool)`

GetUserLabelOk returns a tuple with the UserLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLabel

`func (o *NetworkSecureRouter) SetUserLabel(v string)`

SetUserLabel sets UserLabel field to given value.

### HasUserLabel

`func (o *NetworkSecureRouter) HasUserLabel() bool`

HasUserLabel returns a boolean if a field has been set.

### GetBmc

`func (o *NetworkSecureRouter) GetBmc() ManagementControllerRelationship`

GetBmc returns the Bmc field if non-nil, zero value otherwise.

### GetBmcOk

`func (o *NetworkSecureRouter) GetBmcOk() (*ManagementControllerRelationship, bool)`

GetBmcOk returns a tuple with the Bmc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBmc

`func (o *NetworkSecureRouter) SetBmc(v ManagementControllerRelationship)`

SetBmc sets Bmc field to given value.

### HasBmc

`func (o *NetworkSecureRouter) HasBmc() bool`

HasBmc returns a boolean if a field has been set.

### SetBmcNil

`func (o *NetworkSecureRouter) SetBmcNil(b bool)`

 SetBmcNil sets the value for Bmc to be an explicit nil

### UnsetBmc
`func (o *NetworkSecureRouter) UnsetBmc()`

UnsetBmc ensures that no value is present for Bmc, not even an explicit nil
### GetEquipmentChassis

`func (o *NetworkSecureRouter) GetEquipmentChassis() EquipmentChassisRelationship`

GetEquipmentChassis returns the EquipmentChassis field if non-nil, zero value otherwise.

### GetEquipmentChassisOk

`func (o *NetworkSecureRouter) GetEquipmentChassisOk() (*EquipmentChassisRelationship, bool)`

GetEquipmentChassisOk returns a tuple with the EquipmentChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentChassis

`func (o *NetworkSecureRouter) SetEquipmentChassis(v EquipmentChassisRelationship)`

SetEquipmentChassis sets EquipmentChassis field to given value.

### HasEquipmentChassis

`func (o *NetworkSecureRouter) HasEquipmentChassis() bool`

HasEquipmentChassis returns a boolean if a field has been set.

### SetEquipmentChassisNil

`func (o *NetworkSecureRouter) SetEquipmentChassisNil(b bool)`

 SetEquipmentChassisNil sets the value for EquipmentChassis to be an explicit nil

### UnsetEquipmentChassis
`func (o *NetworkSecureRouter) UnsetEquipmentChassis()`

UnsetEquipmentChassis ensures that no value is present for EquipmentChassis, not even an explicit nil
### GetLanPorts

`func (o *NetworkSecureRouter) GetLanPorts() []EtherLanPortRelationship`

GetLanPorts returns the LanPorts field if non-nil, zero value otherwise.

### GetLanPortsOk

`func (o *NetworkSecureRouter) GetLanPortsOk() (*[]EtherLanPortRelationship, bool)`

GetLanPortsOk returns a tuple with the LanPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanPorts

`func (o *NetworkSecureRouter) SetLanPorts(v []EtherLanPortRelationship)`

SetLanPorts sets LanPorts field to given value.

### HasLanPorts

`func (o *NetworkSecureRouter) HasLanPorts() bool`

HasLanPorts returns a boolean if a field has been set.

### SetLanPortsNil

`func (o *NetworkSecureRouter) SetLanPortsNil(b bool)`

 SetLanPortsNil sets the value for LanPorts to be an explicit nil

### UnsetLanPorts
`func (o *NetworkSecureRouter) UnsetLanPorts()`

UnsetLanPorts ensures that no value is present for LanPorts, not even an explicit nil
### GetLocatorLed

`func (o *NetworkSecureRouter) GetLocatorLed() EquipmentLocatorLedRelationship`

GetLocatorLed returns the LocatorLed field if non-nil, zero value otherwise.

### GetLocatorLedOk

`func (o *NetworkSecureRouter) GetLocatorLedOk() (*EquipmentLocatorLedRelationship, bool)`

GetLocatorLedOk returns a tuple with the LocatorLed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocatorLed

`func (o *NetworkSecureRouter) SetLocatorLed(v EquipmentLocatorLedRelationship)`

SetLocatorLed sets LocatorLed field to given value.

### HasLocatorLed

`func (o *NetworkSecureRouter) HasLocatorLed() bool`

HasLocatorLed returns a boolean if a field has been set.

### SetLocatorLedNil

`func (o *NetworkSecureRouter) SetLocatorLedNil(b bool)`

 SetLocatorLedNil sets the value for LocatorLed to be an explicit nil

### UnsetLocatorLed
`func (o *NetworkSecureRouter) UnsetLocatorLed()`

UnsetLocatorLed ensures that no value is present for LocatorLed, not even an explicit nil
### GetRegisteredDevice

`func (o *NetworkSecureRouter) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NetworkSecureRouter) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NetworkSecureRouter) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NetworkSecureRouter) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *NetworkSecureRouter) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *NetworkSecureRouter) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


