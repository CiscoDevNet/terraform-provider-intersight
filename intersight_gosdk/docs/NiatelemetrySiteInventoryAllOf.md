# NiatelemetrySiteInventoryAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.SiteInventory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.SiteInventory"]
**Apps** | Pointer to **[]string** |  | [optional] 
**FirmwareVersion** | Pointer to **string** | Version of the specified site. | [optional] 
**InstallType** | Pointer to **string** | Fine-grained type DCNM either SAN or LAN. | [optional] 
**IpAddress** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** | Name of the APIC / DCNM site onboarded. | [optional] 
**NexusDashboard** | Pointer to **string** | Name of ND on which site has been onboarded. | [optional] 
**Nodes** | Pointer to **int64** | Number of nodes the site contains. | [optional] 
**Type** | Pointer to **string** | Type of site onboarded either APIC or DCNM. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetrySiteInventoryAllOf

`func NewNiatelemetrySiteInventoryAllOf(classId string, objectType string, ) *NiatelemetrySiteInventoryAllOf`

NewNiatelemetrySiteInventoryAllOf instantiates a new NiatelemetrySiteInventoryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetrySiteInventoryAllOfWithDefaults

`func NewNiatelemetrySiteInventoryAllOfWithDefaults() *NiatelemetrySiteInventoryAllOf`

NewNiatelemetrySiteInventoryAllOfWithDefaults instantiates a new NiatelemetrySiteInventoryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetrySiteInventoryAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetrySiteInventoryAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetrySiteInventoryAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetrySiteInventoryAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetrySiteInventoryAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetrySiteInventoryAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetApps

`func (o *NiatelemetrySiteInventoryAllOf) GetApps() []string`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *NiatelemetrySiteInventoryAllOf) GetAppsOk() (*[]string, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *NiatelemetrySiteInventoryAllOf) SetApps(v []string)`

SetApps sets Apps field to given value.

### HasApps

`func (o *NiatelemetrySiteInventoryAllOf) HasApps() bool`

HasApps returns a boolean if a field has been set.

### SetAppsNil

`func (o *NiatelemetrySiteInventoryAllOf) SetAppsNil(b bool)`

 SetAppsNil sets the value for Apps to be an explicit nil

### UnsetApps
`func (o *NiatelemetrySiteInventoryAllOf) UnsetApps()`

UnsetApps ensures that no value is present for Apps, not even an explicit nil
### GetFirmwareVersion

`func (o *NiatelemetrySiteInventoryAllOf) GetFirmwareVersion() string`

GetFirmwareVersion returns the FirmwareVersion field if non-nil, zero value otherwise.

### GetFirmwareVersionOk

`func (o *NiatelemetrySiteInventoryAllOf) GetFirmwareVersionOk() (*string, bool)`

GetFirmwareVersionOk returns a tuple with the FirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareVersion

`func (o *NiatelemetrySiteInventoryAllOf) SetFirmwareVersion(v string)`

SetFirmwareVersion sets FirmwareVersion field to given value.

### HasFirmwareVersion

`func (o *NiatelemetrySiteInventoryAllOf) HasFirmwareVersion() bool`

HasFirmwareVersion returns a boolean if a field has been set.

### GetInstallType

`func (o *NiatelemetrySiteInventoryAllOf) GetInstallType() string`

GetInstallType returns the InstallType field if non-nil, zero value otherwise.

### GetInstallTypeOk

`func (o *NiatelemetrySiteInventoryAllOf) GetInstallTypeOk() (*string, bool)`

GetInstallTypeOk returns a tuple with the InstallType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallType

`func (o *NiatelemetrySiteInventoryAllOf) SetInstallType(v string)`

SetInstallType sets InstallType field to given value.

### HasInstallType

`func (o *NiatelemetrySiteInventoryAllOf) HasInstallType() bool`

HasInstallType returns a boolean if a field has been set.

### GetIpAddress

`func (o *NiatelemetrySiteInventoryAllOf) GetIpAddress() []string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *NiatelemetrySiteInventoryAllOf) GetIpAddressOk() (*[]string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *NiatelemetrySiteInventoryAllOf) SetIpAddress(v []string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *NiatelemetrySiteInventoryAllOf) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *NiatelemetrySiteInventoryAllOf) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *NiatelemetrySiteInventoryAllOf) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetName

`func (o *NiatelemetrySiteInventoryAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NiatelemetrySiteInventoryAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NiatelemetrySiteInventoryAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NiatelemetrySiteInventoryAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNexusDashboard

`func (o *NiatelemetrySiteInventoryAllOf) GetNexusDashboard() string`

GetNexusDashboard returns the NexusDashboard field if non-nil, zero value otherwise.

### GetNexusDashboardOk

`func (o *NiatelemetrySiteInventoryAllOf) GetNexusDashboardOk() (*string, bool)`

GetNexusDashboardOk returns a tuple with the NexusDashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNexusDashboard

`func (o *NiatelemetrySiteInventoryAllOf) SetNexusDashboard(v string)`

SetNexusDashboard sets NexusDashboard field to given value.

### HasNexusDashboard

`func (o *NiatelemetrySiteInventoryAllOf) HasNexusDashboard() bool`

HasNexusDashboard returns a boolean if a field has been set.

### GetNodes

`func (o *NiatelemetrySiteInventoryAllOf) GetNodes() int64`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *NiatelemetrySiteInventoryAllOf) GetNodesOk() (*int64, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *NiatelemetrySiteInventoryAllOf) SetNodes(v int64)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *NiatelemetrySiteInventoryAllOf) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetType

`func (o *NiatelemetrySiteInventoryAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NiatelemetrySiteInventoryAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NiatelemetrySiteInventoryAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NiatelemetrySiteInventoryAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetrySiteInventoryAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetrySiteInventoryAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetrySiteInventoryAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetrySiteInventoryAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


