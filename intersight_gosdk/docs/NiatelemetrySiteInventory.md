# NiatelemetrySiteInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.SiteInventory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.SiteInventory"]
**Apps** | Pointer to **[]string** |  | [optional] 
**ConnectivityAnalysisCount** | Pointer to **int64** | Returns the total number of connectivity Analysis run for EPs in NDFC Fabrics. | [optional] 
**FirmwareVersion** | Pointer to **string** | Version of the specified site. | [optional] 
**InstallType** | Pointer to **string** | Fine-grained type DCNM either SAN or LAN. | [optional] 
**IpAddress** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** | Name of the APIC / DCNM site onboarded. | [optional] 
**NexusDashboard** | Pointer to **string** | Name of ND on which site has been onboarded. | [optional] 
**Nodes** | Pointer to **int64** | Number of nodes the site contains. | [optional] 
**RecordType** | Pointer to **string** | Specifies whether Site object is DCNM or APIC or ND. | [optional] 
**Type** | Pointer to **string** | Type of site onboarded either APIC or DCNM. | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetrySiteInventory

`func NewNiatelemetrySiteInventory(classId string, objectType string, ) *NiatelemetrySiteInventory`

NewNiatelemetrySiteInventory instantiates a new NiatelemetrySiteInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetrySiteInventoryWithDefaults

`func NewNiatelemetrySiteInventoryWithDefaults() *NiatelemetrySiteInventory`

NewNiatelemetrySiteInventoryWithDefaults instantiates a new NiatelemetrySiteInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetrySiteInventory) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetrySiteInventory) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetrySiteInventory) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetrySiteInventory) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetrySiteInventory) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetrySiteInventory) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetApps

`func (o *NiatelemetrySiteInventory) GetApps() []string`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *NiatelemetrySiteInventory) GetAppsOk() (*[]string, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *NiatelemetrySiteInventory) SetApps(v []string)`

SetApps sets Apps field to given value.

### HasApps

`func (o *NiatelemetrySiteInventory) HasApps() bool`

HasApps returns a boolean if a field has been set.

### SetAppsNil

`func (o *NiatelemetrySiteInventory) SetAppsNil(b bool)`

 SetAppsNil sets the value for Apps to be an explicit nil

### UnsetApps
`func (o *NiatelemetrySiteInventory) UnsetApps()`

UnsetApps ensures that no value is present for Apps, not even an explicit nil
### GetConnectivityAnalysisCount

`func (o *NiatelemetrySiteInventory) GetConnectivityAnalysisCount() int64`

GetConnectivityAnalysisCount returns the ConnectivityAnalysisCount field if non-nil, zero value otherwise.

### GetConnectivityAnalysisCountOk

`func (o *NiatelemetrySiteInventory) GetConnectivityAnalysisCountOk() (*int64, bool)`

GetConnectivityAnalysisCountOk returns a tuple with the ConnectivityAnalysisCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectivityAnalysisCount

`func (o *NiatelemetrySiteInventory) SetConnectivityAnalysisCount(v int64)`

SetConnectivityAnalysisCount sets ConnectivityAnalysisCount field to given value.

### HasConnectivityAnalysisCount

`func (o *NiatelemetrySiteInventory) HasConnectivityAnalysisCount() bool`

HasConnectivityAnalysisCount returns a boolean if a field has been set.

### GetFirmwareVersion

`func (o *NiatelemetrySiteInventory) GetFirmwareVersion() string`

GetFirmwareVersion returns the FirmwareVersion field if non-nil, zero value otherwise.

### GetFirmwareVersionOk

`func (o *NiatelemetrySiteInventory) GetFirmwareVersionOk() (*string, bool)`

GetFirmwareVersionOk returns a tuple with the FirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareVersion

`func (o *NiatelemetrySiteInventory) SetFirmwareVersion(v string)`

SetFirmwareVersion sets FirmwareVersion field to given value.

### HasFirmwareVersion

`func (o *NiatelemetrySiteInventory) HasFirmwareVersion() bool`

HasFirmwareVersion returns a boolean if a field has been set.

### GetInstallType

`func (o *NiatelemetrySiteInventory) GetInstallType() string`

GetInstallType returns the InstallType field if non-nil, zero value otherwise.

### GetInstallTypeOk

`func (o *NiatelemetrySiteInventory) GetInstallTypeOk() (*string, bool)`

GetInstallTypeOk returns a tuple with the InstallType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallType

`func (o *NiatelemetrySiteInventory) SetInstallType(v string)`

SetInstallType sets InstallType field to given value.

### HasInstallType

`func (o *NiatelemetrySiteInventory) HasInstallType() bool`

HasInstallType returns a boolean if a field has been set.

### GetIpAddress

`func (o *NiatelemetrySiteInventory) GetIpAddress() []string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *NiatelemetrySiteInventory) GetIpAddressOk() (*[]string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *NiatelemetrySiteInventory) SetIpAddress(v []string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *NiatelemetrySiteInventory) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *NiatelemetrySiteInventory) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *NiatelemetrySiteInventory) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetName

`func (o *NiatelemetrySiteInventory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NiatelemetrySiteInventory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NiatelemetrySiteInventory) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NiatelemetrySiteInventory) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNexusDashboard

`func (o *NiatelemetrySiteInventory) GetNexusDashboard() string`

GetNexusDashboard returns the NexusDashboard field if non-nil, zero value otherwise.

### GetNexusDashboardOk

`func (o *NiatelemetrySiteInventory) GetNexusDashboardOk() (*string, bool)`

GetNexusDashboardOk returns a tuple with the NexusDashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNexusDashboard

`func (o *NiatelemetrySiteInventory) SetNexusDashboard(v string)`

SetNexusDashboard sets NexusDashboard field to given value.

### HasNexusDashboard

`func (o *NiatelemetrySiteInventory) HasNexusDashboard() bool`

HasNexusDashboard returns a boolean if a field has been set.

### GetNodes

`func (o *NiatelemetrySiteInventory) GetNodes() int64`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *NiatelemetrySiteInventory) GetNodesOk() (*int64, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *NiatelemetrySiteInventory) SetNodes(v int64)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *NiatelemetrySiteInventory) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetrySiteInventory) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetrySiteInventory) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetrySiteInventory) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetrySiteInventory) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetType

`func (o *NiatelemetrySiteInventory) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NiatelemetrySiteInventory) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NiatelemetrySiteInventory) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NiatelemetrySiteInventory) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetrySiteInventory) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetrySiteInventory) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetrySiteInventory) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetrySiteInventory) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *NiatelemetrySiteInventory) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *NiatelemetrySiteInventory) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


