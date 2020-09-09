# NiatelemetryNiaInventoryAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | Pointer to **float32** | CPU usage of device being inventoried. This determines the percentage of CPU resources used. | [optional] 
**CrashResetLogs** | Pointer to **string** | Last crash reset reason of device being inventoried. This determines the last reason for a device&#39;s restart due to crash of the system. | [optional] 
**DeviceName** | Pointer to **string** | Name of device being inventoried. The name the user assigns to the device is inventoried here. | [optional] 
**DeviceType** | Pointer to **string** | Type of device being inventoried. This determines whether the device is a controller, leaf or spine. | [optional] 
**Disk** | Pointer to [**NiatelemetryDiskinfo**](niatelemetry.Diskinfo.md) |  | [optional] 
**FexCount** | Pointer to **int64** | Scale of fabric extendors utilized. | [optional] 
**LogInTime** | Pointer to **string** | Last log in time device being inventoried. This determines the last login time on the device. | [optional] 
**LogOutTime** | Pointer to **string** | Last log out time of device being inventoried. This determines the last logout time on the device. | [optional] 
**MacSecCount** | Pointer to **int64** | Number of Macsec configured interfaces on a TOR. | [optional] 
**MacSecFabCount** | Pointer to **int64** | Number of Macsec configured interfaces on a Spine. | [optional] 
**MacsecTotalCount** | Pointer to **int64** | Number of total Macsec configured interfaces for all nodes. | [optional] 
**Memory** | Pointer to **int64** | Memory usage of device being inventoried. This determines the percentage of memory resources used. | [optional] 
**RecordType** | Pointer to **string** | Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. | [optional] 
**RecordVersion** | Pointer to **string** | Version of record being pushed. This determines what was the API version for data available from the device. | [optional] 
**RoutePrefixCount** | Pointer to **int64** | Total nuumber of v4 and v6 routes per node. | [optional] 
**RoutePrefixV4Count** | Pointer to **int64** | Scale of v4 routes per node. | [optional] 
**RoutePrefixV6Count** | Pointer to **int64** | Scale of v6 routes per node. | [optional] 
**Serial** | Pointer to **string** | Serial number of device being invetoried. The serial number is unique per device and will be used as the key. | [optional] 
**SoftwareDownload** | Pointer to **string** | Last software downloaded of device being inventoried. This determines if software download API was used. | [optional] 
**Version** | Pointer to **string** | Software version of device being inventoried. The various software version values for each device are available on cisco.com. | [optional] 
**LicenseState** | Pointer to [**NiatelemetryNiaLicenseStateRelationship**](niatelemetry.NiaLicenseState.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewNiatelemetryNiaInventoryAllOf

`func NewNiatelemetryNiaInventoryAllOf() *NiatelemetryNiaInventoryAllOf`

NewNiatelemetryNiaInventoryAllOf instantiates a new NiatelemetryNiaInventoryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryNiaInventoryAllOfWithDefaults

`func NewNiatelemetryNiaInventoryAllOfWithDefaults() *NiatelemetryNiaInventoryAllOf`

NewNiatelemetryNiaInventoryAllOfWithDefaults instantiates a new NiatelemetryNiaInventoryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *NiatelemetryNiaInventoryAllOf) GetCpu() float32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *NiatelemetryNiaInventoryAllOf) GetCpuOk() (*float32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *NiatelemetryNiaInventoryAllOf) SetCpu(v float32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *NiatelemetryNiaInventoryAllOf) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetCrashResetLogs

`func (o *NiatelemetryNiaInventoryAllOf) GetCrashResetLogs() string`

GetCrashResetLogs returns the CrashResetLogs field if non-nil, zero value otherwise.

### GetCrashResetLogsOk

`func (o *NiatelemetryNiaInventoryAllOf) GetCrashResetLogsOk() (*string, bool)`

GetCrashResetLogsOk returns a tuple with the CrashResetLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrashResetLogs

`func (o *NiatelemetryNiaInventoryAllOf) SetCrashResetLogs(v string)`

SetCrashResetLogs sets CrashResetLogs field to given value.

### HasCrashResetLogs

`func (o *NiatelemetryNiaInventoryAllOf) HasCrashResetLogs() bool`

HasCrashResetLogs returns a boolean if a field has been set.

### GetDeviceName

`func (o *NiatelemetryNiaInventoryAllOf) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *NiatelemetryNiaInventoryAllOf) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *NiatelemetryNiaInventoryAllOf) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *NiatelemetryNiaInventoryAllOf) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetDeviceType

`func (o *NiatelemetryNiaInventoryAllOf) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *NiatelemetryNiaInventoryAllOf) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *NiatelemetryNiaInventoryAllOf) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *NiatelemetryNiaInventoryAllOf) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetDisk

`func (o *NiatelemetryNiaInventoryAllOf) GetDisk() NiatelemetryDiskinfo`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *NiatelemetryNiaInventoryAllOf) GetDiskOk() (*NiatelemetryDiskinfo, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *NiatelemetryNiaInventoryAllOf) SetDisk(v NiatelemetryDiskinfo)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *NiatelemetryNiaInventoryAllOf) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetFexCount

`func (o *NiatelemetryNiaInventoryAllOf) GetFexCount() int64`

GetFexCount returns the FexCount field if non-nil, zero value otherwise.

### GetFexCountOk

`func (o *NiatelemetryNiaInventoryAllOf) GetFexCountOk() (*int64, bool)`

GetFexCountOk returns a tuple with the FexCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFexCount

`func (o *NiatelemetryNiaInventoryAllOf) SetFexCount(v int64)`

SetFexCount sets FexCount field to given value.

### HasFexCount

`func (o *NiatelemetryNiaInventoryAllOf) HasFexCount() bool`

HasFexCount returns a boolean if a field has been set.

### GetLogInTime

`func (o *NiatelemetryNiaInventoryAllOf) GetLogInTime() string`

GetLogInTime returns the LogInTime field if non-nil, zero value otherwise.

### GetLogInTimeOk

`func (o *NiatelemetryNiaInventoryAllOf) GetLogInTimeOk() (*string, bool)`

GetLogInTimeOk returns a tuple with the LogInTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogInTime

`func (o *NiatelemetryNiaInventoryAllOf) SetLogInTime(v string)`

SetLogInTime sets LogInTime field to given value.

### HasLogInTime

`func (o *NiatelemetryNiaInventoryAllOf) HasLogInTime() bool`

HasLogInTime returns a boolean if a field has been set.

### GetLogOutTime

`func (o *NiatelemetryNiaInventoryAllOf) GetLogOutTime() string`

GetLogOutTime returns the LogOutTime field if non-nil, zero value otherwise.

### GetLogOutTimeOk

`func (o *NiatelemetryNiaInventoryAllOf) GetLogOutTimeOk() (*string, bool)`

GetLogOutTimeOk returns a tuple with the LogOutTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogOutTime

`func (o *NiatelemetryNiaInventoryAllOf) SetLogOutTime(v string)`

SetLogOutTime sets LogOutTime field to given value.

### HasLogOutTime

`func (o *NiatelemetryNiaInventoryAllOf) HasLogOutTime() bool`

HasLogOutTime returns a boolean if a field has been set.

### GetMacSecCount

`func (o *NiatelemetryNiaInventoryAllOf) GetMacSecCount() int64`

GetMacSecCount returns the MacSecCount field if non-nil, zero value otherwise.

### GetMacSecCountOk

`func (o *NiatelemetryNiaInventoryAllOf) GetMacSecCountOk() (*int64, bool)`

GetMacSecCountOk returns a tuple with the MacSecCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacSecCount

`func (o *NiatelemetryNiaInventoryAllOf) SetMacSecCount(v int64)`

SetMacSecCount sets MacSecCount field to given value.

### HasMacSecCount

`func (o *NiatelemetryNiaInventoryAllOf) HasMacSecCount() bool`

HasMacSecCount returns a boolean if a field has been set.

### GetMacSecFabCount

`func (o *NiatelemetryNiaInventoryAllOf) GetMacSecFabCount() int64`

GetMacSecFabCount returns the MacSecFabCount field if non-nil, zero value otherwise.

### GetMacSecFabCountOk

`func (o *NiatelemetryNiaInventoryAllOf) GetMacSecFabCountOk() (*int64, bool)`

GetMacSecFabCountOk returns a tuple with the MacSecFabCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacSecFabCount

`func (o *NiatelemetryNiaInventoryAllOf) SetMacSecFabCount(v int64)`

SetMacSecFabCount sets MacSecFabCount field to given value.

### HasMacSecFabCount

`func (o *NiatelemetryNiaInventoryAllOf) HasMacSecFabCount() bool`

HasMacSecFabCount returns a boolean if a field has been set.

### GetMacsecTotalCount

`func (o *NiatelemetryNiaInventoryAllOf) GetMacsecTotalCount() int64`

GetMacsecTotalCount returns the MacsecTotalCount field if non-nil, zero value otherwise.

### GetMacsecTotalCountOk

`func (o *NiatelemetryNiaInventoryAllOf) GetMacsecTotalCountOk() (*int64, bool)`

GetMacsecTotalCountOk returns a tuple with the MacsecTotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacsecTotalCount

`func (o *NiatelemetryNiaInventoryAllOf) SetMacsecTotalCount(v int64)`

SetMacsecTotalCount sets MacsecTotalCount field to given value.

### HasMacsecTotalCount

`func (o *NiatelemetryNiaInventoryAllOf) HasMacsecTotalCount() bool`

HasMacsecTotalCount returns a boolean if a field has been set.

### GetMemory

`func (o *NiatelemetryNiaInventoryAllOf) GetMemory() int64`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *NiatelemetryNiaInventoryAllOf) GetMemoryOk() (*int64, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *NiatelemetryNiaInventoryAllOf) SetMemory(v int64)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *NiatelemetryNiaInventoryAllOf) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetryNiaInventoryAllOf) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryNiaInventoryAllOf) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryNiaInventoryAllOf) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryNiaInventoryAllOf) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetryNiaInventoryAllOf) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetryNiaInventoryAllOf) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetryNiaInventoryAllOf) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetryNiaInventoryAllOf) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetRoutePrefixCount

`func (o *NiatelemetryNiaInventoryAllOf) GetRoutePrefixCount() int64`

GetRoutePrefixCount returns the RoutePrefixCount field if non-nil, zero value otherwise.

### GetRoutePrefixCountOk

`func (o *NiatelemetryNiaInventoryAllOf) GetRoutePrefixCountOk() (*int64, bool)`

GetRoutePrefixCountOk returns a tuple with the RoutePrefixCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutePrefixCount

`func (o *NiatelemetryNiaInventoryAllOf) SetRoutePrefixCount(v int64)`

SetRoutePrefixCount sets RoutePrefixCount field to given value.

### HasRoutePrefixCount

`func (o *NiatelemetryNiaInventoryAllOf) HasRoutePrefixCount() bool`

HasRoutePrefixCount returns a boolean if a field has been set.

### GetRoutePrefixV4Count

`func (o *NiatelemetryNiaInventoryAllOf) GetRoutePrefixV4Count() int64`

GetRoutePrefixV4Count returns the RoutePrefixV4Count field if non-nil, zero value otherwise.

### GetRoutePrefixV4CountOk

`func (o *NiatelemetryNiaInventoryAllOf) GetRoutePrefixV4CountOk() (*int64, bool)`

GetRoutePrefixV4CountOk returns a tuple with the RoutePrefixV4Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutePrefixV4Count

`func (o *NiatelemetryNiaInventoryAllOf) SetRoutePrefixV4Count(v int64)`

SetRoutePrefixV4Count sets RoutePrefixV4Count field to given value.

### HasRoutePrefixV4Count

`func (o *NiatelemetryNiaInventoryAllOf) HasRoutePrefixV4Count() bool`

HasRoutePrefixV4Count returns a boolean if a field has been set.

### GetRoutePrefixV6Count

`func (o *NiatelemetryNiaInventoryAllOf) GetRoutePrefixV6Count() int64`

GetRoutePrefixV6Count returns the RoutePrefixV6Count field if non-nil, zero value otherwise.

### GetRoutePrefixV6CountOk

`func (o *NiatelemetryNiaInventoryAllOf) GetRoutePrefixV6CountOk() (*int64, bool)`

GetRoutePrefixV6CountOk returns a tuple with the RoutePrefixV6Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutePrefixV6Count

`func (o *NiatelemetryNiaInventoryAllOf) SetRoutePrefixV6Count(v int64)`

SetRoutePrefixV6Count sets RoutePrefixV6Count field to given value.

### HasRoutePrefixV6Count

`func (o *NiatelemetryNiaInventoryAllOf) HasRoutePrefixV6Count() bool`

HasRoutePrefixV6Count returns a boolean if a field has been set.

### GetSerial

`func (o *NiatelemetryNiaInventoryAllOf) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *NiatelemetryNiaInventoryAllOf) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *NiatelemetryNiaInventoryAllOf) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *NiatelemetryNiaInventoryAllOf) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetSoftwareDownload

`func (o *NiatelemetryNiaInventoryAllOf) GetSoftwareDownload() string`

GetSoftwareDownload returns the SoftwareDownload field if non-nil, zero value otherwise.

### GetSoftwareDownloadOk

`func (o *NiatelemetryNiaInventoryAllOf) GetSoftwareDownloadOk() (*string, bool)`

GetSoftwareDownloadOk returns a tuple with the SoftwareDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareDownload

`func (o *NiatelemetryNiaInventoryAllOf) SetSoftwareDownload(v string)`

SetSoftwareDownload sets SoftwareDownload field to given value.

### HasSoftwareDownload

`func (o *NiatelemetryNiaInventoryAllOf) HasSoftwareDownload() bool`

HasSoftwareDownload returns a boolean if a field has been set.

### GetVersion

`func (o *NiatelemetryNiaInventoryAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NiatelemetryNiaInventoryAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NiatelemetryNiaInventoryAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NiatelemetryNiaInventoryAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetLicenseState

`func (o *NiatelemetryNiaInventoryAllOf) GetLicenseState() NiatelemetryNiaLicenseStateRelationship`

GetLicenseState returns the LicenseState field if non-nil, zero value otherwise.

### GetLicenseStateOk

`func (o *NiatelemetryNiaInventoryAllOf) GetLicenseStateOk() (*NiatelemetryNiaLicenseStateRelationship, bool)`

GetLicenseStateOk returns a tuple with the LicenseState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseState

`func (o *NiatelemetryNiaInventoryAllOf) SetLicenseState(v NiatelemetryNiaLicenseStateRelationship)`

SetLicenseState sets LicenseState field to given value.

### HasLicenseState

`func (o *NiatelemetryNiaInventoryAllOf) HasLicenseState() bool`

HasLicenseState returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryNiaInventoryAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryNiaInventoryAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryNiaInventoryAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryNiaInventoryAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


