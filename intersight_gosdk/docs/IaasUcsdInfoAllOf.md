# IaasUcsdInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **string** | Moid of the UCS Director device connector&#39;s asset.DeviceRegistration. | [optional] [readonly] 
**Guid** | Pointer to **string** | Unique ID of UCS Director getting registerd with Intersight. | [optional] [readonly] 
**HostName** | Pointer to **string** | The UCS Director hostname for management. | [optional] [readonly] 
**Ip** | Pointer to **string** | The UCS Director IP address for management. | [optional] [readonly] 
**LastBackup** | Pointer to [**time.Time**](time.Time.md) | Last successful backup created for this UCS Director appliance if backup is configured. | [optional] [readonly] 
**NodeType** | Pointer to **string** | NodeType specifies if UCS Director is deployed in Stand-alone or Multi Node. | [optional] [readonly] 
**ProductName** | Pointer to **string** | The UCS Director product name. | [optional] [readonly] 
**ProductVendor** | Pointer to **string** | The UCS Director product vendor. | [optional] [readonly] 
**ProductVersion** | Pointer to **string** | The UCS Director product/platform version. | [optional] [readonly] 
**Status** | Pointer to **string** | The UCS Director status. Possible values are Active, Inactive, Unknown. | [optional] [readonly] 
**ConnectorPack** | Pointer to [**[]IaasConnectorPackRelationship**](iaas.ConnectorPack.Relationship.md) | An array of relationships to iaasConnectorPack resources. | [optional] [readonly] 
**DeviceStatus** | Pointer to [**[]IaasDeviceStatusRelationship**](iaas.DeviceStatus.Relationship.md) | An array of relationships to iaasDeviceStatus resources. | [optional] [readonly] 
**LicenseInfo** | Pointer to [**IaasLicenseInfoRelationship**](iaas.LicenseInfo.Relationship.md) |  | [optional] 
**MostRunTasks** | Pointer to [**[]IaasMostRunTasksRelationship**](iaas.MostRunTasks.Relationship.md) | An array of relationships to iaasMostRunTasks resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**UcsdManagedInfra** | Pointer to [**IaasUcsdManagedInfraRelationship**](iaas.UcsdManagedInfra.Relationship.md) |  | [optional] 

## Methods

### NewIaasUcsdInfoAllOf

`func NewIaasUcsdInfoAllOf() *IaasUcsdInfoAllOf`

NewIaasUcsdInfoAllOf instantiates a new IaasUcsdInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIaasUcsdInfoAllOfWithDefaults

`func NewIaasUcsdInfoAllOfWithDefaults() *IaasUcsdInfoAllOf`

NewIaasUcsdInfoAllOfWithDefaults instantiates a new IaasUcsdInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *IaasUcsdInfoAllOf) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *IaasUcsdInfoAllOf) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *IaasUcsdInfoAllOf) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *IaasUcsdInfoAllOf) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetGuid

`func (o *IaasUcsdInfoAllOf) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *IaasUcsdInfoAllOf) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *IaasUcsdInfoAllOf) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *IaasUcsdInfoAllOf) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetHostName

`func (o *IaasUcsdInfoAllOf) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *IaasUcsdInfoAllOf) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *IaasUcsdInfoAllOf) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *IaasUcsdInfoAllOf) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetIp

`func (o *IaasUcsdInfoAllOf) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *IaasUcsdInfoAllOf) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *IaasUcsdInfoAllOf) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *IaasUcsdInfoAllOf) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetLastBackup

`func (o *IaasUcsdInfoAllOf) GetLastBackup() time.Time`

GetLastBackup returns the LastBackup field if non-nil, zero value otherwise.

### GetLastBackupOk

`func (o *IaasUcsdInfoAllOf) GetLastBackupOk() (*time.Time, bool)`

GetLastBackupOk returns a tuple with the LastBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBackup

`func (o *IaasUcsdInfoAllOf) SetLastBackup(v time.Time)`

SetLastBackup sets LastBackup field to given value.

### HasLastBackup

`func (o *IaasUcsdInfoAllOf) HasLastBackup() bool`

HasLastBackup returns a boolean if a field has been set.

### GetNodeType

`func (o *IaasUcsdInfoAllOf) GetNodeType() string`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *IaasUcsdInfoAllOf) GetNodeTypeOk() (*string, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *IaasUcsdInfoAllOf) SetNodeType(v string)`

SetNodeType sets NodeType field to given value.

### HasNodeType

`func (o *IaasUcsdInfoAllOf) HasNodeType() bool`

HasNodeType returns a boolean if a field has been set.

### GetProductName

`func (o *IaasUcsdInfoAllOf) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *IaasUcsdInfoAllOf) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *IaasUcsdInfoAllOf) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *IaasUcsdInfoAllOf) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetProductVendor

`func (o *IaasUcsdInfoAllOf) GetProductVendor() string`

GetProductVendor returns the ProductVendor field if non-nil, zero value otherwise.

### GetProductVendorOk

`func (o *IaasUcsdInfoAllOf) GetProductVendorOk() (*string, bool)`

GetProductVendorOk returns a tuple with the ProductVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVendor

`func (o *IaasUcsdInfoAllOf) SetProductVendor(v string)`

SetProductVendor sets ProductVendor field to given value.

### HasProductVendor

`func (o *IaasUcsdInfoAllOf) HasProductVendor() bool`

HasProductVendor returns a boolean if a field has been set.

### GetProductVersion

`func (o *IaasUcsdInfoAllOf) GetProductVersion() string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *IaasUcsdInfoAllOf) GetProductVersionOk() (*string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *IaasUcsdInfoAllOf) SetProductVersion(v string)`

SetProductVersion sets ProductVersion field to given value.

### HasProductVersion

`func (o *IaasUcsdInfoAllOf) HasProductVersion() bool`

HasProductVersion returns a boolean if a field has been set.

### GetStatus

`func (o *IaasUcsdInfoAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IaasUcsdInfoAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IaasUcsdInfoAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IaasUcsdInfoAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetConnectorPack

`func (o *IaasUcsdInfoAllOf) GetConnectorPack() []IaasConnectorPackRelationship`

GetConnectorPack returns the ConnectorPack field if non-nil, zero value otherwise.

### GetConnectorPackOk

`func (o *IaasUcsdInfoAllOf) GetConnectorPackOk() (*[]IaasConnectorPackRelationship, bool)`

GetConnectorPackOk returns a tuple with the ConnectorPack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorPack

`func (o *IaasUcsdInfoAllOf) SetConnectorPack(v []IaasConnectorPackRelationship)`

SetConnectorPack sets ConnectorPack field to given value.

### HasConnectorPack

`func (o *IaasUcsdInfoAllOf) HasConnectorPack() bool`

HasConnectorPack returns a boolean if a field has been set.

### SetConnectorPackNil

`func (o *IaasUcsdInfoAllOf) SetConnectorPackNil(b bool)`

 SetConnectorPackNil sets the value for ConnectorPack to be an explicit nil

### UnsetConnectorPack
`func (o *IaasUcsdInfoAllOf) UnsetConnectorPack()`

UnsetConnectorPack ensures that no value is present for ConnectorPack, not even an explicit nil
### GetDeviceStatus

`func (o *IaasUcsdInfoAllOf) GetDeviceStatus() []IaasDeviceStatusRelationship`

GetDeviceStatus returns the DeviceStatus field if non-nil, zero value otherwise.

### GetDeviceStatusOk

`func (o *IaasUcsdInfoAllOf) GetDeviceStatusOk() (*[]IaasDeviceStatusRelationship, bool)`

GetDeviceStatusOk returns a tuple with the DeviceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceStatus

`func (o *IaasUcsdInfoAllOf) SetDeviceStatus(v []IaasDeviceStatusRelationship)`

SetDeviceStatus sets DeviceStatus field to given value.

### HasDeviceStatus

`func (o *IaasUcsdInfoAllOf) HasDeviceStatus() bool`

HasDeviceStatus returns a boolean if a field has been set.

### SetDeviceStatusNil

`func (o *IaasUcsdInfoAllOf) SetDeviceStatusNil(b bool)`

 SetDeviceStatusNil sets the value for DeviceStatus to be an explicit nil

### UnsetDeviceStatus
`func (o *IaasUcsdInfoAllOf) UnsetDeviceStatus()`

UnsetDeviceStatus ensures that no value is present for DeviceStatus, not even an explicit nil
### GetLicenseInfo

`func (o *IaasUcsdInfoAllOf) GetLicenseInfo() IaasLicenseInfoRelationship`

GetLicenseInfo returns the LicenseInfo field if non-nil, zero value otherwise.

### GetLicenseInfoOk

`func (o *IaasUcsdInfoAllOf) GetLicenseInfoOk() (*IaasLicenseInfoRelationship, bool)`

GetLicenseInfoOk returns a tuple with the LicenseInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseInfo

`func (o *IaasUcsdInfoAllOf) SetLicenseInfo(v IaasLicenseInfoRelationship)`

SetLicenseInfo sets LicenseInfo field to given value.

### HasLicenseInfo

`func (o *IaasUcsdInfoAllOf) HasLicenseInfo() bool`

HasLicenseInfo returns a boolean if a field has been set.

### GetMostRunTasks

`func (o *IaasUcsdInfoAllOf) GetMostRunTasks() []IaasMostRunTasksRelationship`

GetMostRunTasks returns the MostRunTasks field if non-nil, zero value otherwise.

### GetMostRunTasksOk

`func (o *IaasUcsdInfoAllOf) GetMostRunTasksOk() (*[]IaasMostRunTasksRelationship, bool)`

GetMostRunTasksOk returns a tuple with the MostRunTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMostRunTasks

`func (o *IaasUcsdInfoAllOf) SetMostRunTasks(v []IaasMostRunTasksRelationship)`

SetMostRunTasks sets MostRunTasks field to given value.

### HasMostRunTasks

`func (o *IaasUcsdInfoAllOf) HasMostRunTasks() bool`

HasMostRunTasks returns a boolean if a field has been set.

### SetMostRunTasksNil

`func (o *IaasUcsdInfoAllOf) SetMostRunTasksNil(b bool)`

 SetMostRunTasksNil sets the value for MostRunTasks to be an explicit nil

### UnsetMostRunTasks
`func (o *IaasUcsdInfoAllOf) UnsetMostRunTasks()`

UnsetMostRunTasks ensures that no value is present for MostRunTasks, not even an explicit nil
### GetRegisteredDevice

`func (o *IaasUcsdInfoAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *IaasUcsdInfoAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *IaasUcsdInfoAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *IaasUcsdInfoAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetUcsdManagedInfra

`func (o *IaasUcsdInfoAllOf) GetUcsdManagedInfra() IaasUcsdManagedInfraRelationship`

GetUcsdManagedInfra returns the UcsdManagedInfra field if non-nil, zero value otherwise.

### GetUcsdManagedInfraOk

`func (o *IaasUcsdInfoAllOf) GetUcsdManagedInfraOk() (*IaasUcsdManagedInfraRelationship, bool)`

GetUcsdManagedInfraOk returns a tuple with the UcsdManagedInfra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUcsdManagedInfra

`func (o *IaasUcsdInfoAllOf) SetUcsdManagedInfra(v IaasUcsdManagedInfraRelationship)`

SetUcsdManagedInfra sets UcsdManagedInfra field to given value.

### HasUcsdManagedInfra

`func (o *IaasUcsdInfoAllOf) HasUcsdManagedInfra() bool`

HasUcsdManagedInfra returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


