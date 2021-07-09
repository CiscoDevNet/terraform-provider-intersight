# NiatelemetryNiaInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.NiaInventory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.NiaInventory"]
**Cpu** | Pointer to **float32** | CPU usage of device being inventoried. This determines the percentage of CPU resources used. | [optional] 
**CrashResetLogs** | Pointer to **string** | Last crash reset reason of device being inventoried. This determines the last reason for a device&#39;s restart due to crash of the system. | [optional] 
**CustomerDeviceConnector** | Pointer to **string** | Returns the value of the customerDeviceConnector field. | [optional] 
**DcnmLicenseState** | Pointer to **string** | Returns the License state of the device. | [optional] 
**DeviceDiscovery** | Pointer to **string** | Returns the value of the deviceDiscovery field. | [optional] 
**DeviceHealth** | Pointer to **int64** | Returns the device health. | [optional] 
**DeviceId** | Pointer to **string** | Returns the value of the deviceId field. | [optional] 
**DeviceName** | Pointer to **string** | Name of device being inventoried. The name the user assigns to the device is inventoried here. | [optional] 
**DeviceType** | Pointer to **string** | Type of device being inventoried. This determines whether the device is a controller, leaf or spine. | [optional] 
**DeviceUpTime** | Pointer to **int64** | Returns the device up time. | [optional] 
**Disk** | Pointer to [**NullableNiatelemetryDiskinfo**](niatelemetry.Diskinfo.md) |  | [optional] 
**Dn** | Pointer to **string** | Dn for the inventories present. | [optional] 
**FabricName** | Pointer to **string** | Name of the fabric of the device being inventoried. | [optional] 
**FexCount** | Pointer to **int64** | Number of fabric extendors utilized. | [optional] 
**InfraWiNodeCount** | Pointer to **int64** | Number of appliances as physical device that are wired into the cluster. | [optional] 
**IpAddress** | Pointer to **string** | The IP address of the device being inventoried. | [optional] 
**IsVirtualNode** | Pointer to **string** | Flag to specify if the node is virtual. | [optional] 
**LastRebootTime** | Pointer to **string** | Returns the last reboot Time of the device. | [optional] 
**LastResetReason** | Pointer to **string** | Returns the last reset reason of the device. | [optional] 
**LicenseType** | Pointer to **string** | Returns the License type of the device. | [optional] 
**LogInTime** | Pointer to **string** | Last log in time device being inventoried. This determines the last login time on the device. | [optional] 
**LogOutTime** | Pointer to **string** | Last log out time of device being inventoried. This determines the last logout time on the device. | [optional] 
**MacSecCount** | Pointer to **int64** | Number of Macsec configured interfaces on a TOR. | [optional] 
**MacSecFabCount** | Pointer to **int64** | Number of Macsec configured interfaces on a Spine. | [optional] 
**MacsecTotalCount** | Pointer to **int64** | Number of total Macsec configured interfaces for all nodes. | [optional] 
**Memory** | Pointer to **int64** | Memory usage of device being inventoried. This determines the percentage of memory resources used. | [optional] 
**NodeId** | Pointer to **string** | The ID of the device being inventoried. | [optional] 
**NxosBgpMvpn** | Pointer to [**NullableNiatelemetryNxosBgpMvpn**](niatelemetry.NxosBgpMvpn.md) |  | [optional] 
**NxosBootflashDetails** | Pointer to [**NullableNiatelemetryBootflashDetails**](niatelemetry.BootflashDetails.md) |  | [optional] 
**NxosDciInterfaceStatus** | Pointer to **string** | Returns the status of dci interface configured. | [optional] 
**NxosInterfaceBrief** | Pointer to [**NullableNiatelemetryInterface**](niatelemetry.Interface.md) |  | [optional] 
**NxosNveInterfaceStatus** | Pointer to **string** | Returns the value of the nxosNveInterface field. | [optional] 
**NxosNvePacketCounters** | Pointer to [**NullableNiatelemetryNvePacketCounters**](niatelemetry.NvePacketCounters.md) |  | [optional] 
**NxosNveVni** | Pointer to [**NullableNiatelemetryNveVni**](niatelemetry.NveVni.md) |  | [optional] 
**NxosOspfNeighbors** | Pointer to **int64** | Total number of ospf neighbors per switch in DCNM. | [optional] 
**NxosPimNeighbors** | Pointer to **string** | Total number of pim neighbors per switch in DCNM. | [optional] 
**NxosTelnet** | Pointer to **string** | Returns the value of the nxosTelnet field. | [optional] 
**NxosTotalRoutes** | Pointer to **int64** | Total number of routes configured in the DCNM. | [optional] 
**NxosVtp** | Pointer to [**NullableNiatelemetryNxosVtp**](niatelemetry.NxosVtp.md) |  | [optional] 
**RecordType** | Pointer to **string** | Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. | [optional] 
**RecordVersion** | Pointer to **string** | Version of record being pushed. This determines what was the API version for data available from the device. | [optional] 
**RoutePrefixCount** | Pointer to **int64** | Total nuumber of v4 and v6 routes per node. | [optional] 
**RoutePrefixV4Count** | Pointer to **int64** | Number of v4 routes per node. | [optional] 
**RoutePrefixV6Count** | Pointer to **int64** | Number of v6 routes per node. | [optional] 
**Serial** | Pointer to **string** | Serial number of device being invetoried. The serial number is unique per device and will be used as the key. | [optional] 
**SiteName** | Pointer to **string** | Name of fabric domain of the controller. | [optional] 
**SmartAccountId** | Pointer to **int64** | Returns the value of the smartAccountId/CustomerId field. | [optional] 
**SoftwareDownload** | Pointer to **string** | Last software downloaded of device being inventoried. This determines if software download API was used. | [optional] 
**SystemUpTime** | Pointer to **string** | The amount of time that the device being inventoried been up. | [optional] 
**TotalCriticalFaults** | Pointer to **int64** | Returns the total number of critical faults. | [optional] 
**Version** | Pointer to **string** | Software version of device being inventoried. The various software version values for each device are available on cisco.com. | [optional] 
**LicenseState** | Pointer to [**NiatelemetryNiaLicenseStateRelationship**](niatelemetry.NiaLicenseState.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewNiatelemetryNiaInventory

`func NewNiatelemetryNiaInventory(classId string, objectType string, ) *NiatelemetryNiaInventory`

NewNiatelemetryNiaInventory instantiates a new NiatelemetryNiaInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryNiaInventoryWithDefaults

`func NewNiatelemetryNiaInventoryWithDefaults() *NiatelemetryNiaInventory`

NewNiatelemetryNiaInventoryWithDefaults instantiates a new NiatelemetryNiaInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryNiaInventory) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryNiaInventory) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryNiaInventory) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryNiaInventory) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryNiaInventory) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryNiaInventory) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCpu

`func (o *NiatelemetryNiaInventory) GetCpu() float32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *NiatelemetryNiaInventory) GetCpuOk() (*float32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *NiatelemetryNiaInventory) SetCpu(v float32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *NiatelemetryNiaInventory) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetCrashResetLogs

`func (o *NiatelemetryNiaInventory) GetCrashResetLogs() string`

GetCrashResetLogs returns the CrashResetLogs field if non-nil, zero value otherwise.

### GetCrashResetLogsOk

`func (o *NiatelemetryNiaInventory) GetCrashResetLogsOk() (*string, bool)`

GetCrashResetLogsOk returns a tuple with the CrashResetLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrashResetLogs

`func (o *NiatelemetryNiaInventory) SetCrashResetLogs(v string)`

SetCrashResetLogs sets CrashResetLogs field to given value.

### HasCrashResetLogs

`func (o *NiatelemetryNiaInventory) HasCrashResetLogs() bool`

HasCrashResetLogs returns a boolean if a field has been set.

### GetCustomerDeviceConnector

`func (o *NiatelemetryNiaInventory) GetCustomerDeviceConnector() string`

GetCustomerDeviceConnector returns the CustomerDeviceConnector field if non-nil, zero value otherwise.

### GetCustomerDeviceConnectorOk

`func (o *NiatelemetryNiaInventory) GetCustomerDeviceConnectorOk() (*string, bool)`

GetCustomerDeviceConnectorOk returns a tuple with the CustomerDeviceConnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerDeviceConnector

`func (o *NiatelemetryNiaInventory) SetCustomerDeviceConnector(v string)`

SetCustomerDeviceConnector sets CustomerDeviceConnector field to given value.

### HasCustomerDeviceConnector

`func (o *NiatelemetryNiaInventory) HasCustomerDeviceConnector() bool`

HasCustomerDeviceConnector returns a boolean if a field has been set.

### GetDcnmLicenseState

`func (o *NiatelemetryNiaInventory) GetDcnmLicenseState() string`

GetDcnmLicenseState returns the DcnmLicenseState field if non-nil, zero value otherwise.

### GetDcnmLicenseStateOk

`func (o *NiatelemetryNiaInventory) GetDcnmLicenseStateOk() (*string, bool)`

GetDcnmLicenseStateOk returns a tuple with the DcnmLicenseState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcnmLicenseState

`func (o *NiatelemetryNiaInventory) SetDcnmLicenseState(v string)`

SetDcnmLicenseState sets DcnmLicenseState field to given value.

### HasDcnmLicenseState

`func (o *NiatelemetryNiaInventory) HasDcnmLicenseState() bool`

HasDcnmLicenseState returns a boolean if a field has been set.

### GetDeviceDiscovery

`func (o *NiatelemetryNiaInventory) GetDeviceDiscovery() string`

GetDeviceDiscovery returns the DeviceDiscovery field if non-nil, zero value otherwise.

### GetDeviceDiscoveryOk

`func (o *NiatelemetryNiaInventory) GetDeviceDiscoveryOk() (*string, bool)`

GetDeviceDiscoveryOk returns a tuple with the DeviceDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceDiscovery

`func (o *NiatelemetryNiaInventory) SetDeviceDiscovery(v string)`

SetDeviceDiscovery sets DeviceDiscovery field to given value.

### HasDeviceDiscovery

`func (o *NiatelemetryNiaInventory) HasDeviceDiscovery() bool`

HasDeviceDiscovery returns a boolean if a field has been set.

### GetDeviceHealth

`func (o *NiatelemetryNiaInventory) GetDeviceHealth() int64`

GetDeviceHealth returns the DeviceHealth field if non-nil, zero value otherwise.

### GetDeviceHealthOk

`func (o *NiatelemetryNiaInventory) GetDeviceHealthOk() (*int64, bool)`

GetDeviceHealthOk returns a tuple with the DeviceHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceHealth

`func (o *NiatelemetryNiaInventory) SetDeviceHealth(v int64)`

SetDeviceHealth sets DeviceHealth field to given value.

### HasDeviceHealth

`func (o *NiatelemetryNiaInventory) HasDeviceHealth() bool`

HasDeviceHealth returns a boolean if a field has been set.

### GetDeviceId

`func (o *NiatelemetryNiaInventory) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *NiatelemetryNiaInventory) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *NiatelemetryNiaInventory) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *NiatelemetryNiaInventory) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetDeviceName

`func (o *NiatelemetryNiaInventory) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *NiatelemetryNiaInventory) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *NiatelemetryNiaInventory) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *NiatelemetryNiaInventory) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetDeviceType

`func (o *NiatelemetryNiaInventory) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *NiatelemetryNiaInventory) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *NiatelemetryNiaInventory) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *NiatelemetryNiaInventory) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetDeviceUpTime

`func (o *NiatelemetryNiaInventory) GetDeviceUpTime() int64`

GetDeviceUpTime returns the DeviceUpTime field if non-nil, zero value otherwise.

### GetDeviceUpTimeOk

`func (o *NiatelemetryNiaInventory) GetDeviceUpTimeOk() (*int64, bool)`

GetDeviceUpTimeOk returns a tuple with the DeviceUpTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceUpTime

`func (o *NiatelemetryNiaInventory) SetDeviceUpTime(v int64)`

SetDeviceUpTime sets DeviceUpTime field to given value.

### HasDeviceUpTime

`func (o *NiatelemetryNiaInventory) HasDeviceUpTime() bool`

HasDeviceUpTime returns a boolean if a field has been set.

### GetDisk

`func (o *NiatelemetryNiaInventory) GetDisk() NiatelemetryDiskinfo`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *NiatelemetryNiaInventory) GetDiskOk() (*NiatelemetryDiskinfo, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *NiatelemetryNiaInventory) SetDisk(v NiatelemetryDiskinfo)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *NiatelemetryNiaInventory) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### SetDiskNil

`func (o *NiatelemetryNiaInventory) SetDiskNil(b bool)`

 SetDiskNil sets the value for Disk to be an explicit nil

### UnsetDisk
`func (o *NiatelemetryNiaInventory) UnsetDisk()`

UnsetDisk ensures that no value is present for Disk, not even an explicit nil
### GetDn

`func (o *NiatelemetryNiaInventory) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *NiatelemetryNiaInventory) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *NiatelemetryNiaInventory) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *NiatelemetryNiaInventory) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetFabricName

`func (o *NiatelemetryNiaInventory) GetFabricName() string`

GetFabricName returns the FabricName field if non-nil, zero value otherwise.

### GetFabricNameOk

`func (o *NiatelemetryNiaInventory) GetFabricNameOk() (*string, bool)`

GetFabricNameOk returns a tuple with the FabricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricName

`func (o *NiatelemetryNiaInventory) SetFabricName(v string)`

SetFabricName sets FabricName field to given value.

### HasFabricName

`func (o *NiatelemetryNiaInventory) HasFabricName() bool`

HasFabricName returns a boolean if a field has been set.

### GetFexCount

`func (o *NiatelemetryNiaInventory) GetFexCount() int64`

GetFexCount returns the FexCount field if non-nil, zero value otherwise.

### GetFexCountOk

`func (o *NiatelemetryNiaInventory) GetFexCountOk() (*int64, bool)`

GetFexCountOk returns a tuple with the FexCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFexCount

`func (o *NiatelemetryNiaInventory) SetFexCount(v int64)`

SetFexCount sets FexCount field to given value.

### HasFexCount

`func (o *NiatelemetryNiaInventory) HasFexCount() bool`

HasFexCount returns a boolean if a field has been set.

### GetInfraWiNodeCount

`func (o *NiatelemetryNiaInventory) GetInfraWiNodeCount() int64`

GetInfraWiNodeCount returns the InfraWiNodeCount field if non-nil, zero value otherwise.

### GetInfraWiNodeCountOk

`func (o *NiatelemetryNiaInventory) GetInfraWiNodeCountOk() (*int64, bool)`

GetInfraWiNodeCountOk returns a tuple with the InfraWiNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraWiNodeCount

`func (o *NiatelemetryNiaInventory) SetInfraWiNodeCount(v int64)`

SetInfraWiNodeCount sets InfraWiNodeCount field to given value.

### HasInfraWiNodeCount

`func (o *NiatelemetryNiaInventory) HasInfraWiNodeCount() bool`

HasInfraWiNodeCount returns a boolean if a field has been set.

### GetIpAddress

`func (o *NiatelemetryNiaInventory) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *NiatelemetryNiaInventory) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *NiatelemetryNiaInventory) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *NiatelemetryNiaInventory) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetIsVirtualNode

`func (o *NiatelemetryNiaInventory) GetIsVirtualNode() string`

GetIsVirtualNode returns the IsVirtualNode field if non-nil, zero value otherwise.

### GetIsVirtualNodeOk

`func (o *NiatelemetryNiaInventory) GetIsVirtualNodeOk() (*string, bool)`

GetIsVirtualNodeOk returns a tuple with the IsVirtualNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVirtualNode

`func (o *NiatelemetryNiaInventory) SetIsVirtualNode(v string)`

SetIsVirtualNode sets IsVirtualNode field to given value.

### HasIsVirtualNode

`func (o *NiatelemetryNiaInventory) HasIsVirtualNode() bool`

HasIsVirtualNode returns a boolean if a field has been set.

### GetLastRebootTime

`func (o *NiatelemetryNiaInventory) GetLastRebootTime() string`

GetLastRebootTime returns the LastRebootTime field if non-nil, zero value otherwise.

### GetLastRebootTimeOk

`func (o *NiatelemetryNiaInventory) GetLastRebootTimeOk() (*string, bool)`

GetLastRebootTimeOk returns a tuple with the LastRebootTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRebootTime

`func (o *NiatelemetryNiaInventory) SetLastRebootTime(v string)`

SetLastRebootTime sets LastRebootTime field to given value.

### HasLastRebootTime

`func (o *NiatelemetryNiaInventory) HasLastRebootTime() bool`

HasLastRebootTime returns a boolean if a field has been set.

### GetLastResetReason

`func (o *NiatelemetryNiaInventory) GetLastResetReason() string`

GetLastResetReason returns the LastResetReason field if non-nil, zero value otherwise.

### GetLastResetReasonOk

`func (o *NiatelemetryNiaInventory) GetLastResetReasonOk() (*string, bool)`

GetLastResetReasonOk returns a tuple with the LastResetReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResetReason

`func (o *NiatelemetryNiaInventory) SetLastResetReason(v string)`

SetLastResetReason sets LastResetReason field to given value.

### HasLastResetReason

`func (o *NiatelemetryNiaInventory) HasLastResetReason() bool`

HasLastResetReason returns a boolean if a field has been set.

### GetLicenseType

`func (o *NiatelemetryNiaInventory) GetLicenseType() string`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *NiatelemetryNiaInventory) GetLicenseTypeOk() (*string, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *NiatelemetryNiaInventory) SetLicenseType(v string)`

SetLicenseType sets LicenseType field to given value.

### HasLicenseType

`func (o *NiatelemetryNiaInventory) HasLicenseType() bool`

HasLicenseType returns a boolean if a field has been set.

### GetLogInTime

`func (o *NiatelemetryNiaInventory) GetLogInTime() string`

GetLogInTime returns the LogInTime field if non-nil, zero value otherwise.

### GetLogInTimeOk

`func (o *NiatelemetryNiaInventory) GetLogInTimeOk() (*string, bool)`

GetLogInTimeOk returns a tuple with the LogInTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogInTime

`func (o *NiatelemetryNiaInventory) SetLogInTime(v string)`

SetLogInTime sets LogInTime field to given value.

### HasLogInTime

`func (o *NiatelemetryNiaInventory) HasLogInTime() bool`

HasLogInTime returns a boolean if a field has been set.

### GetLogOutTime

`func (o *NiatelemetryNiaInventory) GetLogOutTime() string`

GetLogOutTime returns the LogOutTime field if non-nil, zero value otherwise.

### GetLogOutTimeOk

`func (o *NiatelemetryNiaInventory) GetLogOutTimeOk() (*string, bool)`

GetLogOutTimeOk returns a tuple with the LogOutTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogOutTime

`func (o *NiatelemetryNiaInventory) SetLogOutTime(v string)`

SetLogOutTime sets LogOutTime field to given value.

### HasLogOutTime

`func (o *NiatelemetryNiaInventory) HasLogOutTime() bool`

HasLogOutTime returns a boolean if a field has been set.

### GetMacSecCount

`func (o *NiatelemetryNiaInventory) GetMacSecCount() int64`

GetMacSecCount returns the MacSecCount field if non-nil, zero value otherwise.

### GetMacSecCountOk

`func (o *NiatelemetryNiaInventory) GetMacSecCountOk() (*int64, bool)`

GetMacSecCountOk returns a tuple with the MacSecCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacSecCount

`func (o *NiatelemetryNiaInventory) SetMacSecCount(v int64)`

SetMacSecCount sets MacSecCount field to given value.

### HasMacSecCount

`func (o *NiatelemetryNiaInventory) HasMacSecCount() bool`

HasMacSecCount returns a boolean if a field has been set.

### GetMacSecFabCount

`func (o *NiatelemetryNiaInventory) GetMacSecFabCount() int64`

GetMacSecFabCount returns the MacSecFabCount field if non-nil, zero value otherwise.

### GetMacSecFabCountOk

`func (o *NiatelemetryNiaInventory) GetMacSecFabCountOk() (*int64, bool)`

GetMacSecFabCountOk returns a tuple with the MacSecFabCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacSecFabCount

`func (o *NiatelemetryNiaInventory) SetMacSecFabCount(v int64)`

SetMacSecFabCount sets MacSecFabCount field to given value.

### HasMacSecFabCount

`func (o *NiatelemetryNiaInventory) HasMacSecFabCount() bool`

HasMacSecFabCount returns a boolean if a field has been set.

### GetMacsecTotalCount

`func (o *NiatelemetryNiaInventory) GetMacsecTotalCount() int64`

GetMacsecTotalCount returns the MacsecTotalCount field if non-nil, zero value otherwise.

### GetMacsecTotalCountOk

`func (o *NiatelemetryNiaInventory) GetMacsecTotalCountOk() (*int64, bool)`

GetMacsecTotalCountOk returns a tuple with the MacsecTotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacsecTotalCount

`func (o *NiatelemetryNiaInventory) SetMacsecTotalCount(v int64)`

SetMacsecTotalCount sets MacsecTotalCount field to given value.

### HasMacsecTotalCount

`func (o *NiatelemetryNiaInventory) HasMacsecTotalCount() bool`

HasMacsecTotalCount returns a boolean if a field has been set.

### GetMemory

`func (o *NiatelemetryNiaInventory) GetMemory() int64`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *NiatelemetryNiaInventory) GetMemoryOk() (*int64, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *NiatelemetryNiaInventory) SetMemory(v int64)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *NiatelemetryNiaInventory) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetNodeId

`func (o *NiatelemetryNiaInventory) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *NiatelemetryNiaInventory) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *NiatelemetryNiaInventory) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *NiatelemetryNiaInventory) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetNxosBgpMvpn

`func (o *NiatelemetryNiaInventory) GetNxosBgpMvpn() NiatelemetryNxosBgpMvpn`

GetNxosBgpMvpn returns the NxosBgpMvpn field if non-nil, zero value otherwise.

### GetNxosBgpMvpnOk

`func (o *NiatelemetryNiaInventory) GetNxosBgpMvpnOk() (*NiatelemetryNxosBgpMvpn, bool)`

GetNxosBgpMvpnOk returns a tuple with the NxosBgpMvpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxosBgpMvpn

`func (o *NiatelemetryNiaInventory) SetNxosBgpMvpn(v NiatelemetryNxosBgpMvpn)`

SetNxosBgpMvpn sets NxosBgpMvpn field to given value.

### HasNxosBgpMvpn

`func (o *NiatelemetryNiaInventory) HasNxosBgpMvpn() bool`

HasNxosBgpMvpn returns a boolean if a field has been set.

### SetNxosBgpMvpnNil

`func (o *NiatelemetryNiaInventory) SetNxosBgpMvpnNil(b bool)`

 SetNxosBgpMvpnNil sets the value for NxosBgpMvpn to be an explicit nil

### UnsetNxosBgpMvpn
`func (o *NiatelemetryNiaInventory) UnsetNxosBgpMvpn()`

UnsetNxosBgpMvpn ensures that no value is present for NxosBgpMvpn, not even an explicit nil
### GetNxosBootflashDetails

`func (o *NiatelemetryNiaInventory) GetNxosBootflashDetails() NiatelemetryBootflashDetails`

GetNxosBootflashDetails returns the NxosBootflashDetails field if non-nil, zero value otherwise.

### GetNxosBootflashDetailsOk

`func (o *NiatelemetryNiaInventory) GetNxosBootflashDetailsOk() (*NiatelemetryBootflashDetails, bool)`

GetNxosBootflashDetailsOk returns a tuple with the NxosBootflashDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxosBootflashDetails

`func (o *NiatelemetryNiaInventory) SetNxosBootflashDetails(v NiatelemetryBootflashDetails)`

SetNxosBootflashDetails sets NxosBootflashDetails field to given value.

### HasNxosBootflashDetails

`func (o *NiatelemetryNiaInventory) HasNxosBootflashDetails() bool`

HasNxosBootflashDetails returns a boolean if a field has been set.

### SetNxosBootflashDetailsNil

`func (o *NiatelemetryNiaInventory) SetNxosBootflashDetailsNil(b bool)`

 SetNxosBootflashDetailsNil sets the value for NxosBootflashDetails to be an explicit nil

### UnsetNxosBootflashDetails
`func (o *NiatelemetryNiaInventory) UnsetNxosBootflashDetails()`

UnsetNxosBootflashDetails ensures that no value is present for NxosBootflashDetails, not even an explicit nil
### GetNxosDciInterfaceStatus

`func (o *NiatelemetryNiaInventory) GetNxosDciInterfaceStatus() string`

GetNxosDciInterfaceStatus returns the NxosDciInterfaceStatus field if non-nil, zero value otherwise.

### GetNxosDciInterfaceStatusOk

`func (o *NiatelemetryNiaInventory) GetNxosDciInterfaceStatusOk() (*string, bool)`

GetNxosDciInterfaceStatusOk returns a tuple with the NxosDciInterfaceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxosDciInterfaceStatus

`func (o *NiatelemetryNiaInventory) SetNxosDciInterfaceStatus(v string)`

SetNxosDciInterfaceStatus sets NxosDciInterfaceStatus field to given value.

### HasNxosDciInterfaceStatus

`func (o *NiatelemetryNiaInventory) HasNxosDciInterfaceStatus() bool`

HasNxosDciInterfaceStatus returns a boolean if a field has been set.

### GetNxosInterfaceBrief

`func (o *NiatelemetryNiaInventory) GetNxosInterfaceBrief() NiatelemetryInterface`

GetNxosInterfaceBrief returns the NxosInterfaceBrief field if non-nil, zero value otherwise.

### GetNxosInterfaceBriefOk

`func (o *NiatelemetryNiaInventory) GetNxosInterfaceBriefOk() (*NiatelemetryInterface, bool)`

GetNxosInterfaceBriefOk returns a tuple with the NxosInterfaceBrief field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxosInterfaceBrief

`func (o *NiatelemetryNiaInventory) SetNxosInterfaceBrief(v NiatelemetryInterface)`

SetNxosInterfaceBrief sets NxosInterfaceBrief field to given value.

### HasNxosInterfaceBrief

`func (o *NiatelemetryNiaInventory) HasNxosInterfaceBrief() bool`

HasNxosInterfaceBrief returns a boolean if a field has been set.

### SetNxosInterfaceBriefNil

`func (o *NiatelemetryNiaInventory) SetNxosInterfaceBriefNil(b bool)`

 SetNxosInterfaceBriefNil sets the value for NxosInterfaceBrief to be an explicit nil

### UnsetNxosInterfaceBrief
`func (o *NiatelemetryNiaInventory) UnsetNxosInterfaceBrief()`

UnsetNxosInterfaceBrief ensures that no value is present for NxosInterfaceBrief, not even an explicit nil
### GetNxosNveInterfaceStatus

`func (o *NiatelemetryNiaInventory) GetNxosNveInterfaceStatus() string`

GetNxosNveInterfaceStatus returns the NxosNveInterfaceStatus field if non-nil, zero value otherwise.

### GetNxosNveInterfaceStatusOk

`func (o *NiatelemetryNiaInventory) GetNxosNveInterfaceStatusOk() (*string, bool)`

GetNxosNveInterfaceStatusOk returns a tuple with the NxosNveInterfaceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxosNveInterfaceStatus

`func (o *NiatelemetryNiaInventory) SetNxosNveInterfaceStatus(v string)`

SetNxosNveInterfaceStatus sets NxosNveInterfaceStatus field to given value.

### HasNxosNveInterfaceStatus

`func (o *NiatelemetryNiaInventory) HasNxosNveInterfaceStatus() bool`

HasNxosNveInterfaceStatus returns a boolean if a field has been set.

### GetNxosNvePacketCounters

`func (o *NiatelemetryNiaInventory) GetNxosNvePacketCounters() NiatelemetryNvePacketCounters`

GetNxosNvePacketCounters returns the NxosNvePacketCounters field if non-nil, zero value otherwise.

### GetNxosNvePacketCountersOk

`func (o *NiatelemetryNiaInventory) GetNxosNvePacketCountersOk() (*NiatelemetryNvePacketCounters, bool)`

GetNxosNvePacketCountersOk returns a tuple with the NxosNvePacketCounters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxosNvePacketCounters

`func (o *NiatelemetryNiaInventory) SetNxosNvePacketCounters(v NiatelemetryNvePacketCounters)`

SetNxosNvePacketCounters sets NxosNvePacketCounters field to given value.

### HasNxosNvePacketCounters

`func (o *NiatelemetryNiaInventory) HasNxosNvePacketCounters() bool`

HasNxosNvePacketCounters returns a boolean if a field has been set.

### SetNxosNvePacketCountersNil

`func (o *NiatelemetryNiaInventory) SetNxosNvePacketCountersNil(b bool)`

 SetNxosNvePacketCountersNil sets the value for NxosNvePacketCounters to be an explicit nil

### UnsetNxosNvePacketCounters
`func (o *NiatelemetryNiaInventory) UnsetNxosNvePacketCounters()`

UnsetNxosNvePacketCounters ensures that no value is present for NxosNvePacketCounters, not even an explicit nil
### GetNxosNveVni

`func (o *NiatelemetryNiaInventory) GetNxosNveVni() NiatelemetryNveVni`

GetNxosNveVni returns the NxosNveVni field if non-nil, zero value otherwise.

### GetNxosNveVniOk

`func (o *NiatelemetryNiaInventory) GetNxosNveVniOk() (*NiatelemetryNveVni, bool)`

GetNxosNveVniOk returns a tuple with the NxosNveVni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxosNveVni

`func (o *NiatelemetryNiaInventory) SetNxosNveVni(v NiatelemetryNveVni)`

SetNxosNveVni sets NxosNveVni field to given value.

### HasNxosNveVni

`func (o *NiatelemetryNiaInventory) HasNxosNveVni() bool`

HasNxosNveVni returns a boolean if a field has been set.

### SetNxosNveVniNil

`func (o *NiatelemetryNiaInventory) SetNxosNveVniNil(b bool)`

 SetNxosNveVniNil sets the value for NxosNveVni to be an explicit nil

### UnsetNxosNveVni
`func (o *NiatelemetryNiaInventory) UnsetNxosNveVni()`

UnsetNxosNveVni ensures that no value is present for NxosNveVni, not even an explicit nil
### GetNxosOspfNeighbors

`func (o *NiatelemetryNiaInventory) GetNxosOspfNeighbors() int64`

GetNxosOspfNeighbors returns the NxosOspfNeighbors field if non-nil, zero value otherwise.

### GetNxosOspfNeighborsOk

`func (o *NiatelemetryNiaInventory) GetNxosOspfNeighborsOk() (*int64, bool)`

GetNxosOspfNeighborsOk returns a tuple with the NxosOspfNeighbors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxosOspfNeighbors

`func (o *NiatelemetryNiaInventory) SetNxosOspfNeighbors(v int64)`

SetNxosOspfNeighbors sets NxosOspfNeighbors field to given value.

### HasNxosOspfNeighbors

`func (o *NiatelemetryNiaInventory) HasNxosOspfNeighbors() bool`

HasNxosOspfNeighbors returns a boolean if a field has been set.

### GetNxosPimNeighbors

`func (o *NiatelemetryNiaInventory) GetNxosPimNeighbors() string`

GetNxosPimNeighbors returns the NxosPimNeighbors field if non-nil, zero value otherwise.

### GetNxosPimNeighborsOk

`func (o *NiatelemetryNiaInventory) GetNxosPimNeighborsOk() (*string, bool)`

GetNxosPimNeighborsOk returns a tuple with the NxosPimNeighbors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxosPimNeighbors

`func (o *NiatelemetryNiaInventory) SetNxosPimNeighbors(v string)`

SetNxosPimNeighbors sets NxosPimNeighbors field to given value.

### HasNxosPimNeighbors

`func (o *NiatelemetryNiaInventory) HasNxosPimNeighbors() bool`

HasNxosPimNeighbors returns a boolean if a field has been set.

### GetNxosTelnet

`func (o *NiatelemetryNiaInventory) GetNxosTelnet() string`

GetNxosTelnet returns the NxosTelnet field if non-nil, zero value otherwise.

### GetNxosTelnetOk

`func (o *NiatelemetryNiaInventory) GetNxosTelnetOk() (*string, bool)`

GetNxosTelnetOk returns a tuple with the NxosTelnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxosTelnet

`func (o *NiatelemetryNiaInventory) SetNxosTelnet(v string)`

SetNxosTelnet sets NxosTelnet field to given value.

### HasNxosTelnet

`func (o *NiatelemetryNiaInventory) HasNxosTelnet() bool`

HasNxosTelnet returns a boolean if a field has been set.

### GetNxosTotalRoutes

`func (o *NiatelemetryNiaInventory) GetNxosTotalRoutes() int64`

GetNxosTotalRoutes returns the NxosTotalRoutes field if non-nil, zero value otherwise.

### GetNxosTotalRoutesOk

`func (o *NiatelemetryNiaInventory) GetNxosTotalRoutesOk() (*int64, bool)`

GetNxosTotalRoutesOk returns a tuple with the NxosTotalRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxosTotalRoutes

`func (o *NiatelemetryNiaInventory) SetNxosTotalRoutes(v int64)`

SetNxosTotalRoutes sets NxosTotalRoutes field to given value.

### HasNxosTotalRoutes

`func (o *NiatelemetryNiaInventory) HasNxosTotalRoutes() bool`

HasNxosTotalRoutes returns a boolean if a field has been set.

### GetNxosVtp

`func (o *NiatelemetryNiaInventory) GetNxosVtp() NiatelemetryNxosVtp`

GetNxosVtp returns the NxosVtp field if non-nil, zero value otherwise.

### GetNxosVtpOk

`func (o *NiatelemetryNiaInventory) GetNxosVtpOk() (*NiatelemetryNxosVtp, bool)`

GetNxosVtpOk returns a tuple with the NxosVtp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxosVtp

`func (o *NiatelemetryNiaInventory) SetNxosVtp(v NiatelemetryNxosVtp)`

SetNxosVtp sets NxosVtp field to given value.

### HasNxosVtp

`func (o *NiatelemetryNiaInventory) HasNxosVtp() bool`

HasNxosVtp returns a boolean if a field has been set.

### SetNxosVtpNil

`func (o *NiatelemetryNiaInventory) SetNxosVtpNil(b bool)`

 SetNxosVtpNil sets the value for NxosVtp to be an explicit nil

### UnsetNxosVtp
`func (o *NiatelemetryNiaInventory) UnsetNxosVtp()`

UnsetNxosVtp ensures that no value is present for NxosVtp, not even an explicit nil
### GetRecordType

`func (o *NiatelemetryNiaInventory) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryNiaInventory) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryNiaInventory) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryNiaInventory) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetryNiaInventory) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetryNiaInventory) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetryNiaInventory) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetryNiaInventory) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetRoutePrefixCount

`func (o *NiatelemetryNiaInventory) GetRoutePrefixCount() int64`

GetRoutePrefixCount returns the RoutePrefixCount field if non-nil, zero value otherwise.

### GetRoutePrefixCountOk

`func (o *NiatelemetryNiaInventory) GetRoutePrefixCountOk() (*int64, bool)`

GetRoutePrefixCountOk returns a tuple with the RoutePrefixCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutePrefixCount

`func (o *NiatelemetryNiaInventory) SetRoutePrefixCount(v int64)`

SetRoutePrefixCount sets RoutePrefixCount field to given value.

### HasRoutePrefixCount

`func (o *NiatelemetryNiaInventory) HasRoutePrefixCount() bool`

HasRoutePrefixCount returns a boolean if a field has been set.

### GetRoutePrefixV4Count

`func (o *NiatelemetryNiaInventory) GetRoutePrefixV4Count() int64`

GetRoutePrefixV4Count returns the RoutePrefixV4Count field if non-nil, zero value otherwise.

### GetRoutePrefixV4CountOk

`func (o *NiatelemetryNiaInventory) GetRoutePrefixV4CountOk() (*int64, bool)`

GetRoutePrefixV4CountOk returns a tuple with the RoutePrefixV4Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutePrefixV4Count

`func (o *NiatelemetryNiaInventory) SetRoutePrefixV4Count(v int64)`

SetRoutePrefixV4Count sets RoutePrefixV4Count field to given value.

### HasRoutePrefixV4Count

`func (o *NiatelemetryNiaInventory) HasRoutePrefixV4Count() bool`

HasRoutePrefixV4Count returns a boolean if a field has been set.

### GetRoutePrefixV6Count

`func (o *NiatelemetryNiaInventory) GetRoutePrefixV6Count() int64`

GetRoutePrefixV6Count returns the RoutePrefixV6Count field if non-nil, zero value otherwise.

### GetRoutePrefixV6CountOk

`func (o *NiatelemetryNiaInventory) GetRoutePrefixV6CountOk() (*int64, bool)`

GetRoutePrefixV6CountOk returns a tuple with the RoutePrefixV6Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutePrefixV6Count

`func (o *NiatelemetryNiaInventory) SetRoutePrefixV6Count(v int64)`

SetRoutePrefixV6Count sets RoutePrefixV6Count field to given value.

### HasRoutePrefixV6Count

`func (o *NiatelemetryNiaInventory) HasRoutePrefixV6Count() bool`

HasRoutePrefixV6Count returns a boolean if a field has been set.

### GetSerial

`func (o *NiatelemetryNiaInventory) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *NiatelemetryNiaInventory) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *NiatelemetryNiaInventory) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *NiatelemetryNiaInventory) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetryNiaInventory) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetryNiaInventory) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetryNiaInventory) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetryNiaInventory) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetSmartAccountId

`func (o *NiatelemetryNiaInventory) GetSmartAccountId() int64`

GetSmartAccountId returns the SmartAccountId field if non-nil, zero value otherwise.

### GetSmartAccountIdOk

`func (o *NiatelemetryNiaInventory) GetSmartAccountIdOk() (*int64, bool)`

GetSmartAccountIdOk returns a tuple with the SmartAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartAccountId

`func (o *NiatelemetryNiaInventory) SetSmartAccountId(v int64)`

SetSmartAccountId sets SmartAccountId field to given value.

### HasSmartAccountId

`func (o *NiatelemetryNiaInventory) HasSmartAccountId() bool`

HasSmartAccountId returns a boolean if a field has been set.

### GetSoftwareDownload

`func (o *NiatelemetryNiaInventory) GetSoftwareDownload() string`

GetSoftwareDownload returns the SoftwareDownload field if non-nil, zero value otherwise.

### GetSoftwareDownloadOk

`func (o *NiatelemetryNiaInventory) GetSoftwareDownloadOk() (*string, bool)`

GetSoftwareDownloadOk returns a tuple with the SoftwareDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareDownload

`func (o *NiatelemetryNiaInventory) SetSoftwareDownload(v string)`

SetSoftwareDownload sets SoftwareDownload field to given value.

### HasSoftwareDownload

`func (o *NiatelemetryNiaInventory) HasSoftwareDownload() bool`

HasSoftwareDownload returns a boolean if a field has been set.

### GetSystemUpTime

`func (o *NiatelemetryNiaInventory) GetSystemUpTime() string`

GetSystemUpTime returns the SystemUpTime field if non-nil, zero value otherwise.

### GetSystemUpTimeOk

`func (o *NiatelemetryNiaInventory) GetSystemUpTimeOk() (*string, bool)`

GetSystemUpTimeOk returns a tuple with the SystemUpTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemUpTime

`func (o *NiatelemetryNiaInventory) SetSystemUpTime(v string)`

SetSystemUpTime sets SystemUpTime field to given value.

### HasSystemUpTime

`func (o *NiatelemetryNiaInventory) HasSystemUpTime() bool`

HasSystemUpTime returns a boolean if a field has been set.

### GetTotalCriticalFaults

`func (o *NiatelemetryNiaInventory) GetTotalCriticalFaults() int64`

GetTotalCriticalFaults returns the TotalCriticalFaults field if non-nil, zero value otherwise.

### GetTotalCriticalFaultsOk

`func (o *NiatelemetryNiaInventory) GetTotalCriticalFaultsOk() (*int64, bool)`

GetTotalCriticalFaultsOk returns a tuple with the TotalCriticalFaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCriticalFaults

`func (o *NiatelemetryNiaInventory) SetTotalCriticalFaults(v int64)`

SetTotalCriticalFaults sets TotalCriticalFaults field to given value.

### HasTotalCriticalFaults

`func (o *NiatelemetryNiaInventory) HasTotalCriticalFaults() bool`

HasTotalCriticalFaults returns a boolean if a field has been set.

### GetVersion

`func (o *NiatelemetryNiaInventory) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NiatelemetryNiaInventory) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NiatelemetryNiaInventory) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NiatelemetryNiaInventory) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetLicenseState

`func (o *NiatelemetryNiaInventory) GetLicenseState() NiatelemetryNiaLicenseStateRelationship`

GetLicenseState returns the LicenseState field if non-nil, zero value otherwise.

### GetLicenseStateOk

`func (o *NiatelemetryNiaInventory) GetLicenseStateOk() (*NiatelemetryNiaLicenseStateRelationship, bool)`

GetLicenseStateOk returns a tuple with the LicenseState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseState

`func (o *NiatelemetryNiaInventory) SetLicenseState(v NiatelemetryNiaLicenseStateRelationship)`

SetLicenseState sets LicenseState field to given value.

### HasLicenseState

`func (o *NiatelemetryNiaInventory) HasLicenseState() bool`

HasLicenseState returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryNiaInventory) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryNiaInventory) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryNiaInventory) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryNiaInventory) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


