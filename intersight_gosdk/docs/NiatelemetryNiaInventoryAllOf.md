# NiatelemetryNiaInventoryAllOf

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

### NewNiatelemetryNiaInventoryAllOf

`func NewNiatelemetryNiaInventoryAllOf(classId string, objectType string, ) *NiatelemetryNiaInventoryAllOf`

NewNiatelemetryNiaInventoryAllOf instantiates a new NiatelemetryNiaInventoryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryNiaInventoryAllOfWithDefaults

`func NewNiatelemetryNiaInventoryAllOfWithDefaults() *NiatelemetryNiaInventoryAllOf`

NewNiatelemetryNiaInventoryAllOfWithDefaults instantiates a new NiatelemetryNiaInventoryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryNiaInventoryAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryNiaInventoryAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryNiaInventoryAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryNiaInventoryAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryNiaInventoryAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryNiaInventoryAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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

### GetCustomerDeviceConnector

`func (o *NiatelemetryNiaInventoryAllOf) GetCustomerDeviceConnector() string`

GetCustomerDeviceConnector returns the CustomerDeviceConnector field if non-nil, zero value otherwise.

### GetCustomerDeviceConnectorOk

`func (o *NiatelemetryNiaInventoryAllOf) GetCustomerDeviceConnectorOk() (*string, bool)`

GetCustomerDeviceConnectorOk returns a tuple with the CustomerDeviceConnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerDeviceConnector

`func (o *NiatelemetryNiaInventoryAllOf) SetCustomerDeviceConnector(v string)`

SetCustomerDeviceConnector sets CustomerDeviceConnector field to given value.

### HasCustomerDeviceConnector

`func (o *NiatelemetryNiaInventoryAllOf) HasCustomerDeviceConnector() bool`

HasCustomerDeviceConnector returns a boolean if a field has been set.

### GetDcnmLicenseState

`func (o *NiatelemetryNiaInventoryAllOf) GetDcnmLicenseState() string`

GetDcnmLicenseState returns the DcnmLicenseState field if non-nil, zero value otherwise.

### GetDcnmLicenseStateOk

`func (o *NiatelemetryNiaInventoryAllOf) GetDcnmLicenseStateOk() (*string, bool)`

GetDcnmLicenseStateOk returns a tuple with the DcnmLicenseState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcnmLicenseState

`func (o *NiatelemetryNiaInventoryAllOf) SetDcnmLicenseState(v string)`

SetDcnmLicenseState sets DcnmLicenseState field to given value.

### HasDcnmLicenseState

`func (o *NiatelemetryNiaInventoryAllOf) HasDcnmLicenseState() bool`

HasDcnmLicenseState returns a boolean if a field has been set.

### GetDeviceDiscovery

`func (o *NiatelemetryNiaInventoryAllOf) GetDeviceDiscovery() string`

GetDeviceDiscovery returns the DeviceDiscovery field if non-nil, zero value otherwise.

### GetDeviceDiscoveryOk

`func (o *NiatelemetryNiaInventoryAllOf) GetDeviceDiscoveryOk() (*string, bool)`

GetDeviceDiscoveryOk returns a tuple with the DeviceDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceDiscovery

`func (o *NiatelemetryNiaInventoryAllOf) SetDeviceDiscovery(v string)`

SetDeviceDiscovery sets DeviceDiscovery field to given value.

### HasDeviceDiscovery

`func (o *NiatelemetryNiaInventoryAllOf) HasDeviceDiscovery() bool`

HasDeviceDiscovery returns a boolean if a field has been set.

### GetDeviceHealth

`func (o *NiatelemetryNiaInventoryAllOf) GetDeviceHealth() int64`

GetDeviceHealth returns the DeviceHealth field if non-nil, zero value otherwise.

### GetDeviceHealthOk

`func (o *NiatelemetryNiaInventoryAllOf) GetDeviceHealthOk() (*int64, bool)`

GetDeviceHealthOk returns a tuple with the DeviceHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceHealth

`func (o *NiatelemetryNiaInventoryAllOf) SetDeviceHealth(v int64)`

SetDeviceHealth sets DeviceHealth field to given value.

### HasDeviceHealth

`func (o *NiatelemetryNiaInventoryAllOf) HasDeviceHealth() bool`

HasDeviceHealth returns a boolean if a field has been set.

### GetDeviceId

`func (o *NiatelemetryNiaInventoryAllOf) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *NiatelemetryNiaInventoryAllOf) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *NiatelemetryNiaInventoryAllOf) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *NiatelemetryNiaInventoryAllOf) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

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

### GetDeviceUpTime

`func (o *NiatelemetryNiaInventoryAllOf) GetDeviceUpTime() int64`

GetDeviceUpTime returns the DeviceUpTime field if non-nil, zero value otherwise.

### GetDeviceUpTimeOk

`func (o *NiatelemetryNiaInventoryAllOf) GetDeviceUpTimeOk() (*int64, bool)`

GetDeviceUpTimeOk returns a tuple with the DeviceUpTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceUpTime

`func (o *NiatelemetryNiaInventoryAllOf) SetDeviceUpTime(v int64)`

SetDeviceUpTime sets DeviceUpTime field to given value.

### HasDeviceUpTime

`func (o *NiatelemetryNiaInventoryAllOf) HasDeviceUpTime() bool`

HasDeviceUpTime returns a boolean if a field has been set.

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

### SetDiskNil

`func (o *NiatelemetryNiaInventoryAllOf) SetDiskNil(b bool)`

 SetDiskNil sets the value for Disk to be an explicit nil

### UnsetDisk
`func (o *NiatelemetryNiaInventoryAllOf) UnsetDisk()`

UnsetDisk ensures that no value is present for Disk, not even an explicit nil
### GetDn

`func (o *NiatelemetryNiaInventoryAllOf) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *NiatelemetryNiaInventoryAllOf) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *NiatelemetryNiaInventoryAllOf) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *NiatelemetryNiaInventoryAllOf) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetFabricName

`func (o *NiatelemetryNiaInventoryAllOf) GetFabricName() string`

GetFabricName returns the FabricName field if non-nil, zero value otherwise.

### GetFabricNameOk

`func (o *NiatelemetryNiaInventoryAllOf) GetFabricNameOk() (*string, bool)`

GetFabricNameOk returns a tuple with the FabricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricName

`func (o *NiatelemetryNiaInventoryAllOf) SetFabricName(v string)`

SetFabricName sets FabricName field to given value.

### HasFabricName

`func (o *NiatelemetryNiaInventoryAllOf) HasFabricName() bool`

HasFabricName returns a boolean if a field has been set.

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

### GetInfraWiNodeCount

`func (o *NiatelemetryNiaInventoryAllOf) GetInfraWiNodeCount() int64`

GetInfraWiNodeCount returns the InfraWiNodeCount field if non-nil, zero value otherwise.

### GetInfraWiNodeCountOk

`func (o *NiatelemetryNiaInventoryAllOf) GetInfraWiNodeCountOk() (*int64, bool)`

GetInfraWiNodeCountOk returns a tuple with the InfraWiNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraWiNodeCount

`func (o *NiatelemetryNiaInventoryAllOf) SetInfraWiNodeCount(v int64)`

SetInfraWiNodeCount sets InfraWiNodeCount field to given value.

### HasInfraWiNodeCount

`func (o *NiatelemetryNiaInventoryAllOf) HasInfraWiNodeCount() bool`

HasInfraWiNodeCount returns a boolean if a field has been set.

### GetIpAddress

`func (o *NiatelemetryNiaInventoryAllOf) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *NiatelemetryNiaInventoryAllOf) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *NiatelemetryNiaInventoryAllOf) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *NiatelemetryNiaInventoryAllOf) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetIsVirtualNode

`func (o *NiatelemetryNiaInventoryAllOf) GetIsVirtualNode() string`

GetIsVirtualNode returns the IsVirtualNode field if non-nil, zero value otherwise.

### GetIsVirtualNodeOk

`func (o *NiatelemetryNiaInventoryAllOf) GetIsVirtualNodeOk() (*string, bool)`

GetIsVirtualNodeOk returns a tuple with the IsVirtualNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVirtualNode

`func (o *NiatelemetryNiaInventoryAllOf) SetIsVirtualNode(v string)`

SetIsVirtualNode sets IsVirtualNode field to given value.

### HasIsVirtualNode

`func (o *NiatelemetryNiaInventoryAllOf) HasIsVirtualNode() bool`

HasIsVirtualNode returns a boolean if a field has been set.

### GetLastRebootTime

`func (o *NiatelemetryNiaInventoryAllOf) GetLastRebootTime() string`

GetLastRebootTime returns the LastRebootTime field if non-nil, zero value otherwise.

### GetLastRebootTimeOk

`func (o *NiatelemetryNiaInventoryAllOf) GetLastRebootTimeOk() (*string, bool)`

GetLastRebootTimeOk returns a tuple with the LastRebootTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRebootTime

`func (o *NiatelemetryNiaInventoryAllOf) SetLastRebootTime(v string)`

SetLastRebootTime sets LastRebootTime field to given value.

### HasLastRebootTime

`func (o *NiatelemetryNiaInventoryAllOf) HasLastRebootTime() bool`

HasLastRebootTime returns a boolean if a field has been set.

### GetLastResetReason

`func (o *NiatelemetryNiaInventoryAllOf) GetLastResetReason() string`

GetLastResetReason returns the LastResetReason field if non-nil, zero value otherwise.

### GetLastResetReasonOk

`func (o *NiatelemetryNiaInventoryAllOf) GetLastResetReasonOk() (*string, bool)`

GetLastResetReasonOk returns a tuple with the LastResetReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResetReason

`func (o *NiatelemetryNiaInventoryAllOf) SetLastResetReason(v string)`

SetLastResetReason sets LastResetReason field to given value.

### HasLastResetReason

`func (o *NiatelemetryNiaInventoryAllOf) HasLastResetReason() bool`

HasLastResetReason returns a boolean if a field has been set.

### GetLicenseType

`func (o *NiatelemetryNiaInventoryAllOf) GetLicenseType() string`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *NiatelemetryNiaInventoryAllOf) GetLicenseTypeOk() (*string, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *NiatelemetryNiaInventoryAllOf) SetLicenseType(v string)`

SetLicenseType sets LicenseType field to given value.

### HasLicenseType

`func (o *NiatelemetryNiaInventoryAllOf) HasLicenseType() bool`

HasLicenseType returns a boolean if a field has been set.

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

### GetNodeId

`func (o *NiatelemetryNiaInventoryAllOf) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *NiatelemetryNiaInventoryAllOf) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *NiatelemetryNiaInventoryAllOf) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *NiatelemetryNiaInventoryAllOf) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetNxosBgpMvpn

`func (o *NiatelemetryNiaInventoryAllOf) GetNxosBgpMvpn() NiatelemetryNxosBgpMvpn`

GetNxosBgpMvpn returns the NxosBgpMvpn field if non-nil, zero value otherwise.

### GetNxosBgpMvpnOk

`func (o *NiatelemetryNiaInventoryAllOf) GetNxosBgpMvpnOk() (*NiatelemetryNxosBgpMvpn, bool)`

GetNxosBgpMvpnOk returns a tuple with the NxosBgpMvpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxosBgpMvpn

`func (o *NiatelemetryNiaInventoryAllOf) SetNxosBgpMvpn(v NiatelemetryNxosBgpMvpn)`

SetNxosBgpMvpn sets NxosBgpMvpn field to given value.

### HasNxosBgpMvpn

`func (o *NiatelemetryNiaInventoryAllOf) HasNxosBgpMvpn() bool`

HasNxosBgpMvpn returns a boolean if a field has been set.

### SetNxosBgpMvpnNil

`func (o *NiatelemetryNiaInventoryAllOf) SetNxosBgpMvpnNil(b bool)`

 SetNxosBgpMvpnNil sets the value for NxosBgpMvpn to be an explicit nil

### UnsetNxosBgpMvpn
`func (o *NiatelemetryNiaInventoryAllOf) UnsetNxosBgpMvpn()`

UnsetNxosBgpMvpn ensures that no value is present for NxosBgpMvpn, not even an explicit nil
### GetNxosBootflashDetails

`func (o *NiatelemetryNiaInventoryAllOf) GetNxosBootflashDetails() NiatelemetryBootflashDetails`

GetNxosBootflashDetails returns the NxosBootflashDetails field if non-nil, zero value otherwise.

### GetNxosBootflashDetailsOk

`func (o *NiatelemetryNiaInventoryAllOf) GetNxosBootflashDetailsOk() (*NiatelemetryBootflashDetails, bool)`

GetNxosBootflashDetailsOk returns a tuple with the NxosBootflashDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxosBootflashDetails

`func (o *NiatelemetryNiaInventoryAllOf) SetNxosBootflashDetails(v NiatelemetryBootflashDetails)`

SetNxosBootflashDetails sets NxosBootflashDetails field to given value.

### HasNxosBootflashDetails

`func (o *NiatelemetryNiaInventoryAllOf) HasNxosBootflashDetails() bool`

HasNxosBootflashDetails returns a boolean if a field has been set.

### SetNxosBootflashDetailsNil

`func (o *NiatelemetryNiaInventoryAllOf) SetNxosBootflashDetailsNil(b bool)`

 SetNxosBootflashDetailsNil sets the value for NxosBootflashDetails to be an explicit nil

### UnsetNxosBootflashDetails
`func (o *NiatelemetryNiaInventoryAllOf) UnsetNxosBootflashDetails()`

UnsetNxosBootflashDetails ensures that no value is present for NxosBootflashDetails, not even an explicit nil
### GetNxosDciInterfaceStatus

`func (o *NiatelemetryNiaInventoryAllOf) GetNxosDciInterfaceStatus() string`

GetNxosDciInterfaceStatus returns the NxosDciInterfaceStatus field if non-nil, zero value otherwise.

### GetNxosDciInterfaceStatusOk

`func (o *NiatelemetryNiaInventoryAllOf) GetNxosDciInterfaceStatusOk() (*string, bool)`

GetNxosDciInterfaceStatusOk returns a tuple with the NxosDciInterfaceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxosDciInterfaceStatus

`func (o *NiatelemetryNiaInventoryAllOf) SetNxosDciInterfaceStatus(v string)`

SetNxosDciInterfaceStatus sets NxosDciInterfaceStatus field to given value.

### HasNxosDciInterfaceStatus

`func (o *NiatelemetryNiaInventoryAllOf) HasNxosDciInterfaceStatus() bool`

HasNxosDciInterfaceStatus returns a boolean if a field has been set.

### GetNxosInterfaceBrief

`func (o *NiatelemetryNiaInventoryAllOf) GetNxosInterfaceBrief() NiatelemetryInterface`

GetNxosInterfaceBrief returns the NxosInterfaceBrief field if non-nil, zero value otherwise.

### GetNxosInterfaceBriefOk

`func (o *NiatelemetryNiaInventoryAllOf) GetNxosInterfaceBriefOk() (*NiatelemetryInterface, bool)`

GetNxosInterfaceBriefOk returns a tuple with the NxosInterfaceBrief field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxosInterfaceBrief

`func (o *NiatelemetryNiaInventoryAllOf) SetNxosInterfaceBrief(v NiatelemetryInterface)`

SetNxosInterfaceBrief sets NxosInterfaceBrief field to given value.

### HasNxosInterfaceBrief

`func (o *NiatelemetryNiaInventoryAllOf) HasNxosInterfaceBrief() bool`

HasNxosInterfaceBrief returns a boolean if a field has been set.

### SetNxosInterfaceBriefNil

`func (o *NiatelemetryNiaInventoryAllOf) SetNxosInterfaceBriefNil(b bool)`

 SetNxosInterfaceBriefNil sets the value for NxosInterfaceBrief to be an explicit nil

### UnsetNxosInterfaceBrief
`func (o *NiatelemetryNiaInventoryAllOf) UnsetNxosInterfaceBrief()`

UnsetNxosInterfaceBrief ensures that no value is present for NxosInterfaceBrief, not even an explicit nil
### GetNxosNveInterfaceStatus

`func (o *NiatelemetryNiaInventoryAllOf) GetNxosNveInterfaceStatus() string`

GetNxosNveInterfaceStatus returns the NxosNveInterfaceStatus field if non-nil, zero value otherwise.

### GetNxosNveInterfaceStatusOk

`func (o *NiatelemetryNiaInventoryAllOf) GetNxosNveInterfaceStatusOk() (*string, bool)`

GetNxosNveInterfaceStatusOk returns a tuple with the NxosNveInterfaceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxosNveInterfaceStatus

`func (o *NiatelemetryNiaInventoryAllOf) SetNxosNveInterfaceStatus(v string)`

SetNxosNveInterfaceStatus sets NxosNveInterfaceStatus field to given value.

### HasNxosNveInterfaceStatus

`func (o *NiatelemetryNiaInventoryAllOf) HasNxosNveInterfaceStatus() bool`

HasNxosNveInterfaceStatus returns a boolean if a field has been set.

### GetNxosNvePacketCounters

`func (o *NiatelemetryNiaInventoryAllOf) GetNxosNvePacketCounters() NiatelemetryNvePacketCounters`

GetNxosNvePacketCounters returns the NxosNvePacketCounters field if non-nil, zero value otherwise.

### GetNxosNvePacketCountersOk

`func (o *NiatelemetryNiaInventoryAllOf) GetNxosNvePacketCountersOk() (*NiatelemetryNvePacketCounters, bool)`

GetNxosNvePacketCountersOk returns a tuple with the NxosNvePacketCounters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxosNvePacketCounters

`func (o *NiatelemetryNiaInventoryAllOf) SetNxosNvePacketCounters(v NiatelemetryNvePacketCounters)`

SetNxosNvePacketCounters sets NxosNvePacketCounters field to given value.

### HasNxosNvePacketCounters

`func (o *NiatelemetryNiaInventoryAllOf) HasNxosNvePacketCounters() bool`

HasNxosNvePacketCounters returns a boolean if a field has been set.

### SetNxosNvePacketCountersNil

`func (o *NiatelemetryNiaInventoryAllOf) SetNxosNvePacketCountersNil(b bool)`

 SetNxosNvePacketCountersNil sets the value for NxosNvePacketCounters to be an explicit nil

### UnsetNxosNvePacketCounters
`func (o *NiatelemetryNiaInventoryAllOf) UnsetNxosNvePacketCounters()`

UnsetNxosNvePacketCounters ensures that no value is present for NxosNvePacketCounters, not even an explicit nil
### GetNxosNveVni

`func (o *NiatelemetryNiaInventoryAllOf) GetNxosNveVni() NiatelemetryNveVni`

GetNxosNveVni returns the NxosNveVni field if non-nil, zero value otherwise.

### GetNxosNveVniOk

`func (o *NiatelemetryNiaInventoryAllOf) GetNxosNveVniOk() (*NiatelemetryNveVni, bool)`

GetNxosNveVniOk returns a tuple with the NxosNveVni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxosNveVni

`func (o *NiatelemetryNiaInventoryAllOf) SetNxosNveVni(v NiatelemetryNveVni)`

SetNxosNveVni sets NxosNveVni field to given value.

### HasNxosNveVni

`func (o *NiatelemetryNiaInventoryAllOf) HasNxosNveVni() bool`

HasNxosNveVni returns a boolean if a field has been set.

### SetNxosNveVniNil

`func (o *NiatelemetryNiaInventoryAllOf) SetNxosNveVniNil(b bool)`

 SetNxosNveVniNil sets the value for NxosNveVni to be an explicit nil

### UnsetNxosNveVni
`func (o *NiatelemetryNiaInventoryAllOf) UnsetNxosNveVni()`

UnsetNxosNveVni ensures that no value is present for NxosNveVni, not even an explicit nil
### GetNxosOspfNeighbors

`func (o *NiatelemetryNiaInventoryAllOf) GetNxosOspfNeighbors() int64`

GetNxosOspfNeighbors returns the NxosOspfNeighbors field if non-nil, zero value otherwise.

### GetNxosOspfNeighborsOk

`func (o *NiatelemetryNiaInventoryAllOf) GetNxosOspfNeighborsOk() (*int64, bool)`

GetNxosOspfNeighborsOk returns a tuple with the NxosOspfNeighbors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxosOspfNeighbors

`func (o *NiatelemetryNiaInventoryAllOf) SetNxosOspfNeighbors(v int64)`

SetNxosOspfNeighbors sets NxosOspfNeighbors field to given value.

### HasNxosOspfNeighbors

`func (o *NiatelemetryNiaInventoryAllOf) HasNxosOspfNeighbors() bool`

HasNxosOspfNeighbors returns a boolean if a field has been set.

### GetNxosPimNeighbors

`func (o *NiatelemetryNiaInventoryAllOf) GetNxosPimNeighbors() string`

GetNxosPimNeighbors returns the NxosPimNeighbors field if non-nil, zero value otherwise.

### GetNxosPimNeighborsOk

`func (o *NiatelemetryNiaInventoryAllOf) GetNxosPimNeighborsOk() (*string, bool)`

GetNxosPimNeighborsOk returns a tuple with the NxosPimNeighbors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxosPimNeighbors

`func (o *NiatelemetryNiaInventoryAllOf) SetNxosPimNeighbors(v string)`

SetNxosPimNeighbors sets NxosPimNeighbors field to given value.

### HasNxosPimNeighbors

`func (o *NiatelemetryNiaInventoryAllOf) HasNxosPimNeighbors() bool`

HasNxosPimNeighbors returns a boolean if a field has been set.

### GetNxosTelnet

`func (o *NiatelemetryNiaInventoryAllOf) GetNxosTelnet() string`

GetNxosTelnet returns the NxosTelnet field if non-nil, zero value otherwise.

### GetNxosTelnetOk

`func (o *NiatelemetryNiaInventoryAllOf) GetNxosTelnetOk() (*string, bool)`

GetNxosTelnetOk returns a tuple with the NxosTelnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxosTelnet

`func (o *NiatelemetryNiaInventoryAllOf) SetNxosTelnet(v string)`

SetNxosTelnet sets NxosTelnet field to given value.

### HasNxosTelnet

`func (o *NiatelemetryNiaInventoryAllOf) HasNxosTelnet() bool`

HasNxosTelnet returns a boolean if a field has been set.

### GetNxosTotalRoutes

`func (o *NiatelemetryNiaInventoryAllOf) GetNxosTotalRoutes() int64`

GetNxosTotalRoutes returns the NxosTotalRoutes field if non-nil, zero value otherwise.

### GetNxosTotalRoutesOk

`func (o *NiatelemetryNiaInventoryAllOf) GetNxosTotalRoutesOk() (*int64, bool)`

GetNxosTotalRoutesOk returns a tuple with the NxosTotalRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxosTotalRoutes

`func (o *NiatelemetryNiaInventoryAllOf) SetNxosTotalRoutes(v int64)`

SetNxosTotalRoutes sets NxosTotalRoutes field to given value.

### HasNxosTotalRoutes

`func (o *NiatelemetryNiaInventoryAllOf) HasNxosTotalRoutes() bool`

HasNxosTotalRoutes returns a boolean if a field has been set.

### GetNxosVtp

`func (o *NiatelemetryNiaInventoryAllOf) GetNxosVtp() NiatelemetryNxosVtp`

GetNxosVtp returns the NxosVtp field if non-nil, zero value otherwise.

### GetNxosVtpOk

`func (o *NiatelemetryNiaInventoryAllOf) GetNxosVtpOk() (*NiatelemetryNxosVtp, bool)`

GetNxosVtpOk returns a tuple with the NxosVtp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxosVtp

`func (o *NiatelemetryNiaInventoryAllOf) SetNxosVtp(v NiatelemetryNxosVtp)`

SetNxosVtp sets NxosVtp field to given value.

### HasNxosVtp

`func (o *NiatelemetryNiaInventoryAllOf) HasNxosVtp() bool`

HasNxosVtp returns a boolean if a field has been set.

### SetNxosVtpNil

`func (o *NiatelemetryNiaInventoryAllOf) SetNxosVtpNil(b bool)`

 SetNxosVtpNil sets the value for NxosVtp to be an explicit nil

### UnsetNxosVtp
`func (o *NiatelemetryNiaInventoryAllOf) UnsetNxosVtp()`

UnsetNxosVtp ensures that no value is present for NxosVtp, not even an explicit nil
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

### GetSiteName

`func (o *NiatelemetryNiaInventoryAllOf) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetryNiaInventoryAllOf) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetryNiaInventoryAllOf) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetryNiaInventoryAllOf) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetSmartAccountId

`func (o *NiatelemetryNiaInventoryAllOf) GetSmartAccountId() int64`

GetSmartAccountId returns the SmartAccountId field if non-nil, zero value otherwise.

### GetSmartAccountIdOk

`func (o *NiatelemetryNiaInventoryAllOf) GetSmartAccountIdOk() (*int64, bool)`

GetSmartAccountIdOk returns a tuple with the SmartAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartAccountId

`func (o *NiatelemetryNiaInventoryAllOf) SetSmartAccountId(v int64)`

SetSmartAccountId sets SmartAccountId field to given value.

### HasSmartAccountId

`func (o *NiatelemetryNiaInventoryAllOf) HasSmartAccountId() bool`

HasSmartAccountId returns a boolean if a field has been set.

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

### GetSystemUpTime

`func (o *NiatelemetryNiaInventoryAllOf) GetSystemUpTime() string`

GetSystemUpTime returns the SystemUpTime field if non-nil, zero value otherwise.

### GetSystemUpTimeOk

`func (o *NiatelemetryNiaInventoryAllOf) GetSystemUpTimeOk() (*string, bool)`

GetSystemUpTimeOk returns a tuple with the SystemUpTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemUpTime

`func (o *NiatelemetryNiaInventoryAllOf) SetSystemUpTime(v string)`

SetSystemUpTime sets SystemUpTime field to given value.

### HasSystemUpTime

`func (o *NiatelemetryNiaInventoryAllOf) HasSystemUpTime() bool`

HasSystemUpTime returns a boolean if a field has been set.

### GetTotalCriticalFaults

`func (o *NiatelemetryNiaInventoryAllOf) GetTotalCriticalFaults() int64`

GetTotalCriticalFaults returns the TotalCriticalFaults field if non-nil, zero value otherwise.

### GetTotalCriticalFaultsOk

`func (o *NiatelemetryNiaInventoryAllOf) GetTotalCriticalFaultsOk() (*int64, bool)`

GetTotalCriticalFaultsOk returns a tuple with the TotalCriticalFaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCriticalFaults

`func (o *NiatelemetryNiaInventoryAllOf) SetTotalCriticalFaults(v int64)`

SetTotalCriticalFaults sets TotalCriticalFaults field to given value.

### HasTotalCriticalFaults

`func (o *NiatelemetryNiaInventoryAllOf) HasTotalCriticalFaults() bool`

HasTotalCriticalFaults returns a boolean if a field has been set.

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


