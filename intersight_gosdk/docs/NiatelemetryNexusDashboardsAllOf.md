# NiatelemetryNexusDashboardsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.NexusDashboards"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.NexusDashboards"]
**ClusterName** | Pointer to **string** | Nexus Dashboard can onboard multiple APIC clusters/sites. | [optional] 
**Dn** | Pointer to **string** | Dn of the objects present for Nexus Dashboard devices. | [optional] 
**IsClusterHealthy** | Pointer to **string** | Health of Nexus Dashboard cluster. | [optional] 
**NdClusterSize** | Pointer to **int64** | Number of nodes in Nexus Dashboard cluster. | [optional] 
**NdType** | Pointer to **string** | Node type in Nexus Dashboard cluster. | [optional] 
**NdVersion** | Pointer to **string** | Version running on Nexus Dashboard. | [optional] 
**NumberOfApps** | Pointer to **int64** | Number of applications installed in the Nexus Dashboard. | [optional] 
**NumberOfSchemasInMso** | Pointer to **int64** | Number of total schemas in Multi-Site Orchestrator. | [optional] 
**NumberOfSitesInMso** | Pointer to **int64** | Number of sites in Multi-Site Orchestrator. | [optional] 
**NumberOfSitesServiced** | Pointer to **int64** | Number of sites serviced by ND. | [optional] 
**NumberOfTenantsInMso** | Pointer to **int64** | Number of total tenants in Multi-Site Orchestrator. | [optional] 
**RecordType** | Pointer to **string** | Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. | [optional] 
**TypeOfSiteInMso** | Pointer to **string** | Type of site added to Multi-Site Orchestrator. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetryNexusDashboardsAllOf

`func NewNiatelemetryNexusDashboardsAllOf(classId string, objectType string, ) *NiatelemetryNexusDashboardsAllOf`

NewNiatelemetryNexusDashboardsAllOf instantiates a new NiatelemetryNexusDashboardsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryNexusDashboardsAllOfWithDefaults

`func NewNiatelemetryNexusDashboardsAllOfWithDefaults() *NiatelemetryNexusDashboardsAllOf`

NewNiatelemetryNexusDashboardsAllOfWithDefaults instantiates a new NiatelemetryNexusDashboardsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryNexusDashboardsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryNexusDashboardsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryNexusDashboardsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryNexusDashboardsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryNexusDashboardsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryNexusDashboardsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClusterName

`func (o *NiatelemetryNexusDashboardsAllOf) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *NiatelemetryNexusDashboardsAllOf) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *NiatelemetryNexusDashboardsAllOf) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *NiatelemetryNexusDashboardsAllOf) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetDn

`func (o *NiatelemetryNexusDashboardsAllOf) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *NiatelemetryNexusDashboardsAllOf) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *NiatelemetryNexusDashboardsAllOf) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *NiatelemetryNexusDashboardsAllOf) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetIsClusterHealthy

`func (o *NiatelemetryNexusDashboardsAllOf) GetIsClusterHealthy() string`

GetIsClusterHealthy returns the IsClusterHealthy field if non-nil, zero value otherwise.

### GetIsClusterHealthyOk

`func (o *NiatelemetryNexusDashboardsAllOf) GetIsClusterHealthyOk() (*string, bool)`

GetIsClusterHealthyOk returns a tuple with the IsClusterHealthy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClusterHealthy

`func (o *NiatelemetryNexusDashboardsAllOf) SetIsClusterHealthy(v string)`

SetIsClusterHealthy sets IsClusterHealthy field to given value.

### HasIsClusterHealthy

`func (o *NiatelemetryNexusDashboardsAllOf) HasIsClusterHealthy() bool`

HasIsClusterHealthy returns a boolean if a field has been set.

### GetNdClusterSize

`func (o *NiatelemetryNexusDashboardsAllOf) GetNdClusterSize() int64`

GetNdClusterSize returns the NdClusterSize field if non-nil, zero value otherwise.

### GetNdClusterSizeOk

`func (o *NiatelemetryNexusDashboardsAllOf) GetNdClusterSizeOk() (*int64, bool)`

GetNdClusterSizeOk returns a tuple with the NdClusterSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNdClusterSize

`func (o *NiatelemetryNexusDashboardsAllOf) SetNdClusterSize(v int64)`

SetNdClusterSize sets NdClusterSize field to given value.

### HasNdClusterSize

`func (o *NiatelemetryNexusDashboardsAllOf) HasNdClusterSize() bool`

HasNdClusterSize returns a boolean if a field has been set.

### GetNdType

`func (o *NiatelemetryNexusDashboardsAllOf) GetNdType() string`

GetNdType returns the NdType field if non-nil, zero value otherwise.

### GetNdTypeOk

`func (o *NiatelemetryNexusDashboardsAllOf) GetNdTypeOk() (*string, bool)`

GetNdTypeOk returns a tuple with the NdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNdType

`func (o *NiatelemetryNexusDashboardsAllOf) SetNdType(v string)`

SetNdType sets NdType field to given value.

### HasNdType

`func (o *NiatelemetryNexusDashboardsAllOf) HasNdType() bool`

HasNdType returns a boolean if a field has been set.

### GetNdVersion

`func (o *NiatelemetryNexusDashboardsAllOf) GetNdVersion() string`

GetNdVersion returns the NdVersion field if non-nil, zero value otherwise.

### GetNdVersionOk

`func (o *NiatelemetryNexusDashboardsAllOf) GetNdVersionOk() (*string, bool)`

GetNdVersionOk returns a tuple with the NdVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNdVersion

`func (o *NiatelemetryNexusDashboardsAllOf) SetNdVersion(v string)`

SetNdVersion sets NdVersion field to given value.

### HasNdVersion

`func (o *NiatelemetryNexusDashboardsAllOf) HasNdVersion() bool`

HasNdVersion returns a boolean if a field has been set.

### GetNumberOfApps

`func (o *NiatelemetryNexusDashboardsAllOf) GetNumberOfApps() int64`

GetNumberOfApps returns the NumberOfApps field if non-nil, zero value otherwise.

### GetNumberOfAppsOk

`func (o *NiatelemetryNexusDashboardsAllOf) GetNumberOfAppsOk() (*int64, bool)`

GetNumberOfAppsOk returns a tuple with the NumberOfApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfApps

`func (o *NiatelemetryNexusDashboardsAllOf) SetNumberOfApps(v int64)`

SetNumberOfApps sets NumberOfApps field to given value.

### HasNumberOfApps

`func (o *NiatelemetryNexusDashboardsAllOf) HasNumberOfApps() bool`

HasNumberOfApps returns a boolean if a field has been set.

### GetNumberOfSchemasInMso

`func (o *NiatelemetryNexusDashboardsAllOf) GetNumberOfSchemasInMso() int64`

GetNumberOfSchemasInMso returns the NumberOfSchemasInMso field if non-nil, zero value otherwise.

### GetNumberOfSchemasInMsoOk

`func (o *NiatelemetryNexusDashboardsAllOf) GetNumberOfSchemasInMsoOk() (*int64, bool)`

GetNumberOfSchemasInMsoOk returns a tuple with the NumberOfSchemasInMso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfSchemasInMso

`func (o *NiatelemetryNexusDashboardsAllOf) SetNumberOfSchemasInMso(v int64)`

SetNumberOfSchemasInMso sets NumberOfSchemasInMso field to given value.

### HasNumberOfSchemasInMso

`func (o *NiatelemetryNexusDashboardsAllOf) HasNumberOfSchemasInMso() bool`

HasNumberOfSchemasInMso returns a boolean if a field has been set.

### GetNumberOfSitesInMso

`func (o *NiatelemetryNexusDashboardsAllOf) GetNumberOfSitesInMso() int64`

GetNumberOfSitesInMso returns the NumberOfSitesInMso field if non-nil, zero value otherwise.

### GetNumberOfSitesInMsoOk

`func (o *NiatelemetryNexusDashboardsAllOf) GetNumberOfSitesInMsoOk() (*int64, bool)`

GetNumberOfSitesInMsoOk returns a tuple with the NumberOfSitesInMso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfSitesInMso

`func (o *NiatelemetryNexusDashboardsAllOf) SetNumberOfSitesInMso(v int64)`

SetNumberOfSitesInMso sets NumberOfSitesInMso field to given value.

### HasNumberOfSitesInMso

`func (o *NiatelemetryNexusDashboardsAllOf) HasNumberOfSitesInMso() bool`

HasNumberOfSitesInMso returns a boolean if a field has been set.

### GetNumberOfSitesServiced

`func (o *NiatelemetryNexusDashboardsAllOf) GetNumberOfSitesServiced() int64`

GetNumberOfSitesServiced returns the NumberOfSitesServiced field if non-nil, zero value otherwise.

### GetNumberOfSitesServicedOk

`func (o *NiatelemetryNexusDashboardsAllOf) GetNumberOfSitesServicedOk() (*int64, bool)`

GetNumberOfSitesServicedOk returns a tuple with the NumberOfSitesServiced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfSitesServiced

`func (o *NiatelemetryNexusDashboardsAllOf) SetNumberOfSitesServiced(v int64)`

SetNumberOfSitesServiced sets NumberOfSitesServiced field to given value.

### HasNumberOfSitesServiced

`func (o *NiatelemetryNexusDashboardsAllOf) HasNumberOfSitesServiced() bool`

HasNumberOfSitesServiced returns a boolean if a field has been set.

### GetNumberOfTenantsInMso

`func (o *NiatelemetryNexusDashboardsAllOf) GetNumberOfTenantsInMso() int64`

GetNumberOfTenantsInMso returns the NumberOfTenantsInMso field if non-nil, zero value otherwise.

### GetNumberOfTenantsInMsoOk

`func (o *NiatelemetryNexusDashboardsAllOf) GetNumberOfTenantsInMsoOk() (*int64, bool)`

GetNumberOfTenantsInMsoOk returns a tuple with the NumberOfTenantsInMso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfTenantsInMso

`func (o *NiatelemetryNexusDashboardsAllOf) SetNumberOfTenantsInMso(v int64)`

SetNumberOfTenantsInMso sets NumberOfTenantsInMso field to given value.

### HasNumberOfTenantsInMso

`func (o *NiatelemetryNexusDashboardsAllOf) HasNumberOfTenantsInMso() bool`

HasNumberOfTenantsInMso returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetryNexusDashboardsAllOf) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryNexusDashboardsAllOf) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryNexusDashboardsAllOf) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryNexusDashboardsAllOf) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetTypeOfSiteInMso

`func (o *NiatelemetryNexusDashboardsAllOf) GetTypeOfSiteInMso() string`

GetTypeOfSiteInMso returns the TypeOfSiteInMso field if non-nil, zero value otherwise.

### GetTypeOfSiteInMsoOk

`func (o *NiatelemetryNexusDashboardsAllOf) GetTypeOfSiteInMsoOk() (*string, bool)`

GetTypeOfSiteInMsoOk returns a tuple with the TypeOfSiteInMso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeOfSiteInMso

`func (o *NiatelemetryNexusDashboardsAllOf) SetTypeOfSiteInMso(v string)`

SetTypeOfSiteInMso sets TypeOfSiteInMso field to given value.

### HasTypeOfSiteInMso

`func (o *NiatelemetryNexusDashboardsAllOf) HasTypeOfSiteInMso() bool`

HasTypeOfSiteInMso returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryNexusDashboardsAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryNexusDashboardsAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryNexusDashboardsAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryNexusDashboardsAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


