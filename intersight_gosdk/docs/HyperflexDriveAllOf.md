# HyperflexDriveAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.Drive"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.Drive"]
**Capacity** | Pointer to **int64** | Provisioned capacity of the drive in bytes. | [optional] [readonly] 
**DiskUseState** | Pointer to **string** | Disk inventory state. Should be one of values defined in enum. * &#x60;UNKNOWN&#x60; - The use state of the disk is unknown. * &#x60;USED&#x60; - The use state of the disk is used. * &#x60;NOTUSED&#x60; - The use state of the disk is unused. * &#x60;EMPTY&#x60; - The use state of the disk is empty. | [optional] [readonly] [default to "UNKNOWN"]
**HostName** | Pointer to **string** | Host Name. | [optional] [readonly] 
**HostUuid** | Pointer to **string** | The unique identifier of the drive&#39;s host. | [optional] [readonly] 
**ModelNumber** | Pointer to **string** | The model number of the disk or SSD. | [optional] [readonly] 
**NodeUuid** | Pointer to **string** | The unique identifier of the Hyperflex Node to which the disk is attached. | [optional] [readonly] 
**Path** | Pointer to **string** | Device path for the drive. | [optional] [readonly] 
**Protocol** | Pointer to **string** | Disk Protocol - SATA, NVME, SAS. * &#x60;Unknown&#x60; - Disk protocol is unknown. * &#x60;SAS&#x60; - Serial Attached SCSI protocol (SAS) used in disk. * &#x60;NVMe&#x60; - Non-volatile memory express (NVMe) protocol used in disk. * &#x60;SATA&#x60; - Serial Advanced Technology Attachment (SATA) used in disk. | [optional] [readonly] [default to "Unknown"]
**SerialNumber** | Pointer to **string** | Serial number of the Hyperflex drive. | [optional] [readonly] 
**SlotNumber** | Pointer to **string** | The SCSI slot number of the drive. | [optional] [readonly] 
**Status** | Pointer to **string** | Disk inventory state as determined by storfs inventory module. Should be one of values defined in enum. * &#x60;UNKNOWN&#x60; - The state of the disk is unknown. * &#x60;CLAIMED&#x60; - The state of the disk is claimed by storfs and has a valid storfs format. * &#x60;AVAILABLE&#x60; - The disk is available but not claimed by storfs. * &#x60;IGNORED&#x60; - The disk ash been ignored by storfs. * &#x60;BLACKLISTED&#x60; - The disk has been blacklisted by storfs. * &#x60;SECUREERASED&#x60; - The disk has been secure erased. | [optional] [readonly] [default to "UNKNOWN"]
**Type** | Pointer to **string** | Type of disk - UNKNOWN, HDD, SSD. * &#x60;Unknown&#x60; - Default unknown disk type. * &#x60;SSD&#x60; - Storage disk with Solid state disk. * &#x60;HDD&#x60; - Storage disk with Hard disk drive. * &#x60;NVRAM&#x60; - Storage disk with Non-volatile random-access memory type. * &#x60;SATA&#x60; - Disk drive implementation with Serial Advanced Technology Attachment (SATA). * &#x60;BSAS&#x60; - Bridged SAS-SATA disks with added hardware to enable them to be plugged into a SAS-connected storage shelf. * &#x60;FC&#x60; - Storage disk with Fiber Channel. * &#x60;FSAS&#x60; - Near Line SAS. NL-SAS drives are enterprise SATA drives with a SAS interface, head, media, androtational speed of traditional enterprise-class SATA drives with the fully capable SAS interfacetypical for classic SAS drives. * &#x60;LUN&#x60; - Logical Unit Number refers to a logical disk. * &#x60;MSATA&#x60; - SATA disk in multi-disk carrier storage shelf. * &#x60;SAS&#x60; - Storage disk with serial attached SCSI. * &#x60;VMDISK&#x60; - Virtual machine Data Disk. | [optional] [readonly] [default to "Unknown"]
**Usage** | Pointer to **string** | Specifies what the disk is used for. Should be one of values defined in enum. * &#x60;UNKNOWN&#x60; - The usage of the disk is unknown. * &#x60;PERSISTENCE&#x60; - The usage of the disk is for generic data storage. * &#x60;SYSTEM&#x60; - The usage of the disk is related to system storage. * &#x60;CACHING&#x60; - The usage of the disk is for caching. | [optional] [readonly] [default to "UNKNOWN"]
**UsedCapacity** | Pointer to **int64** | Used Capacity of the drive in Bytes. | [optional] [readonly] 
**Uuid** | Pointer to **string** | The unique identifier of the Hyperflex drive. | [optional] [readonly] 
**Version** | Pointer to **string** | The firmware version of the Hyperflex drive. | [optional] [readonly] 
**LocatorLed** | Pointer to [**EquipmentLocatorLedRelationship**](equipment.LocatorLed.Relationship.md) |  | [optional] 
**Node** | Pointer to [**HyperflexNodeRelationship**](hyperflex.Node.Relationship.md) |  | [optional] 

## Methods

### NewHyperflexDriveAllOf

`func NewHyperflexDriveAllOf(classId string, objectType string, ) *HyperflexDriveAllOf`

NewHyperflexDriveAllOf instantiates a new HyperflexDriveAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexDriveAllOfWithDefaults

`func NewHyperflexDriveAllOfWithDefaults() *HyperflexDriveAllOf`

NewHyperflexDriveAllOfWithDefaults instantiates a new HyperflexDriveAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexDriveAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexDriveAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexDriveAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexDriveAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexDriveAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexDriveAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCapacity

`func (o *HyperflexDriveAllOf) GetCapacity() int64`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *HyperflexDriveAllOf) GetCapacityOk() (*int64, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *HyperflexDriveAllOf) SetCapacity(v int64)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *HyperflexDriveAllOf) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetDiskUseState

`func (o *HyperflexDriveAllOf) GetDiskUseState() string`

GetDiskUseState returns the DiskUseState field if non-nil, zero value otherwise.

### GetDiskUseStateOk

`func (o *HyperflexDriveAllOf) GetDiskUseStateOk() (*string, bool)`

GetDiskUseStateOk returns a tuple with the DiskUseState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskUseState

`func (o *HyperflexDriveAllOf) SetDiskUseState(v string)`

SetDiskUseState sets DiskUseState field to given value.

### HasDiskUseState

`func (o *HyperflexDriveAllOf) HasDiskUseState() bool`

HasDiskUseState returns a boolean if a field has been set.

### GetHostName

`func (o *HyperflexDriveAllOf) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *HyperflexDriveAllOf) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *HyperflexDriveAllOf) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *HyperflexDriveAllOf) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetHostUuid

`func (o *HyperflexDriveAllOf) GetHostUuid() string`

GetHostUuid returns the HostUuid field if non-nil, zero value otherwise.

### GetHostUuidOk

`func (o *HyperflexDriveAllOf) GetHostUuidOk() (*string, bool)`

GetHostUuidOk returns a tuple with the HostUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostUuid

`func (o *HyperflexDriveAllOf) SetHostUuid(v string)`

SetHostUuid sets HostUuid field to given value.

### HasHostUuid

`func (o *HyperflexDriveAllOf) HasHostUuid() bool`

HasHostUuid returns a boolean if a field has been set.

### GetModelNumber

`func (o *HyperflexDriveAllOf) GetModelNumber() string`

GetModelNumber returns the ModelNumber field if non-nil, zero value otherwise.

### GetModelNumberOk

`func (o *HyperflexDriveAllOf) GetModelNumberOk() (*string, bool)`

GetModelNumberOk returns a tuple with the ModelNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelNumber

`func (o *HyperflexDriveAllOf) SetModelNumber(v string)`

SetModelNumber sets ModelNumber field to given value.

### HasModelNumber

`func (o *HyperflexDriveAllOf) HasModelNumber() bool`

HasModelNumber returns a boolean if a field has been set.

### GetNodeUuid

`func (o *HyperflexDriveAllOf) GetNodeUuid() string`

GetNodeUuid returns the NodeUuid field if non-nil, zero value otherwise.

### GetNodeUuidOk

`func (o *HyperflexDriveAllOf) GetNodeUuidOk() (*string, bool)`

GetNodeUuidOk returns a tuple with the NodeUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeUuid

`func (o *HyperflexDriveAllOf) SetNodeUuid(v string)`

SetNodeUuid sets NodeUuid field to given value.

### HasNodeUuid

`func (o *HyperflexDriveAllOf) HasNodeUuid() bool`

HasNodeUuid returns a boolean if a field has been set.

### GetPath

`func (o *HyperflexDriveAllOf) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *HyperflexDriveAllOf) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *HyperflexDriveAllOf) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *HyperflexDriveAllOf) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetProtocol

`func (o *HyperflexDriveAllOf) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *HyperflexDriveAllOf) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *HyperflexDriveAllOf) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *HyperflexDriveAllOf) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSerialNumber

`func (o *HyperflexDriveAllOf) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *HyperflexDriveAllOf) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *HyperflexDriveAllOf) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *HyperflexDriveAllOf) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSlotNumber

`func (o *HyperflexDriveAllOf) GetSlotNumber() string`

GetSlotNumber returns the SlotNumber field if non-nil, zero value otherwise.

### GetSlotNumberOk

`func (o *HyperflexDriveAllOf) GetSlotNumberOk() (*string, bool)`

GetSlotNumberOk returns a tuple with the SlotNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotNumber

`func (o *HyperflexDriveAllOf) SetSlotNumber(v string)`

SetSlotNumber sets SlotNumber field to given value.

### HasSlotNumber

`func (o *HyperflexDriveAllOf) HasSlotNumber() bool`

HasSlotNumber returns a boolean if a field has been set.

### GetStatus

`func (o *HyperflexDriveAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HyperflexDriveAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HyperflexDriveAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HyperflexDriveAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *HyperflexDriveAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HyperflexDriveAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HyperflexDriveAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HyperflexDriveAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUsage

`func (o *HyperflexDriveAllOf) GetUsage() string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *HyperflexDriveAllOf) GetUsageOk() (*string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *HyperflexDriveAllOf) SetUsage(v string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *HyperflexDriveAllOf) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetUsedCapacity

`func (o *HyperflexDriveAllOf) GetUsedCapacity() int64`

GetUsedCapacity returns the UsedCapacity field if non-nil, zero value otherwise.

### GetUsedCapacityOk

`func (o *HyperflexDriveAllOf) GetUsedCapacityOk() (*int64, bool)`

GetUsedCapacityOk returns a tuple with the UsedCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedCapacity

`func (o *HyperflexDriveAllOf) SetUsedCapacity(v int64)`

SetUsedCapacity sets UsedCapacity field to given value.

### HasUsedCapacity

`func (o *HyperflexDriveAllOf) HasUsedCapacity() bool`

HasUsedCapacity returns a boolean if a field has been set.

### GetUuid

`func (o *HyperflexDriveAllOf) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *HyperflexDriveAllOf) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *HyperflexDriveAllOf) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *HyperflexDriveAllOf) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVersion

`func (o *HyperflexDriveAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HyperflexDriveAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HyperflexDriveAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HyperflexDriveAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetLocatorLed

`func (o *HyperflexDriveAllOf) GetLocatorLed() EquipmentLocatorLedRelationship`

GetLocatorLed returns the LocatorLed field if non-nil, zero value otherwise.

### GetLocatorLedOk

`func (o *HyperflexDriveAllOf) GetLocatorLedOk() (*EquipmentLocatorLedRelationship, bool)`

GetLocatorLedOk returns a tuple with the LocatorLed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocatorLed

`func (o *HyperflexDriveAllOf) SetLocatorLed(v EquipmentLocatorLedRelationship)`

SetLocatorLed sets LocatorLed field to given value.

### HasLocatorLed

`func (o *HyperflexDriveAllOf) HasLocatorLed() bool`

HasLocatorLed returns a boolean if a field has been set.

### GetNode

`func (o *HyperflexDriveAllOf) GetNode() HyperflexNodeRelationship`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *HyperflexDriveAllOf) GetNodeOk() (*HyperflexNodeRelationship, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *HyperflexDriveAllOf) SetNode(v HyperflexNodeRelationship)`

SetNode sets Node field to given value.

### HasNode

`func (o *HyperflexDriveAllOf) HasNode() bool`

HasNode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


