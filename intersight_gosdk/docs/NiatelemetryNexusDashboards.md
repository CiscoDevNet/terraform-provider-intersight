# NiatelemetryNexusDashboards

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.NexusDashboards"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.NexusDashboards"]
**ClusterName** | Pointer to **string** | Nexus Dashboard can onboard multiple APIC clusters/sites. | [optional] 
**ClusterUuid** | Pointer to **string** | UUID of the Nexus Dashboard cluster. | [optional] 
**Dn** | Pointer to **string** | Dn of the objects present for Nexus Dashboard devices. | [optional] 
**IsClusterHealthy** | Pointer to **string** | Health of Nexus Dashboard cluster. | [optional] 
**NdClusterSize** | Pointer to **int64** | Number of nodes in Nexus Dashboard cluster. | [optional] 
**NdSites** | Pointer to [**[]NiatelemetrySites**](NiatelemetrySites.md) |  | [optional] 
**NdType** | Pointer to **string** | Node type in Nexus Dashboard cluster. | [optional] 
**NdVersion** | Pointer to **string** | Version running on Nexus Dashboard. | [optional] 
**NumberOfApps** | Pointer to **int64** | Number of applications installed in the Nexus Dashboard. | [optional] 
**NumberOfInsightGroups** | Pointer to **int64** | Number of total insight groups in ND. | [optional] 
**NumberOfNirDashboards** | Pointer to **int64** | Number of total NIR dashboards in ND. | [optional] 
**NumberOfSchemasInMso** | Pointer to **int64** | Number of total schemas in Multi-Site Orchestrator. | [optional] 
**NumberOfSitesInMso** | Pointer to **int64** | Number of sites in Multi-Site Orchestrator. | [optional] 
**NumberOfSitesServiced** | Pointer to **int64** | Number of sites serviced by ND. | [optional] 
**NumberOfTenantsInMso** | Pointer to **int64** | Number of total tenants in Multi-Site Orchestrator. | [optional] 
**RecordType** | Pointer to **string** | Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. | [optional] 
**TypeOfSiteInMso** | Pointer to **string** | Type of site added to Multi-Site Orchestrator. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetryNexusDashboards

`func NewNiatelemetryNexusDashboards(classId string, objectType string, ) *NiatelemetryNexusDashboards`

NewNiatelemetryNexusDashboards instantiates a new NiatelemetryNexusDashboards object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryNexusDashboardsWithDefaults

`func NewNiatelemetryNexusDashboardsWithDefaults() *NiatelemetryNexusDashboards`

NewNiatelemetryNexusDashboardsWithDefaults instantiates a new NiatelemetryNexusDashboards object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryNexusDashboards) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryNexusDashboards) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryNexusDashboards) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryNexusDashboards) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryNexusDashboards) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryNexusDashboards) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClusterName

`func (o *NiatelemetryNexusDashboards) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *NiatelemetryNexusDashboards) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *NiatelemetryNexusDashboards) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *NiatelemetryNexusDashboards) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetClusterUuid

`func (o *NiatelemetryNexusDashboards) GetClusterUuid() string`

GetClusterUuid returns the ClusterUuid field if non-nil, zero value otherwise.

### GetClusterUuidOk

`func (o *NiatelemetryNexusDashboards) GetClusterUuidOk() (*string, bool)`

GetClusterUuidOk returns a tuple with the ClusterUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterUuid

`func (o *NiatelemetryNexusDashboards) SetClusterUuid(v string)`

SetClusterUuid sets ClusterUuid field to given value.

### HasClusterUuid

`func (o *NiatelemetryNexusDashboards) HasClusterUuid() bool`

HasClusterUuid returns a boolean if a field has been set.

### GetDn

`func (o *NiatelemetryNexusDashboards) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *NiatelemetryNexusDashboards) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *NiatelemetryNexusDashboards) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *NiatelemetryNexusDashboards) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetIsClusterHealthy

`func (o *NiatelemetryNexusDashboards) GetIsClusterHealthy() string`

GetIsClusterHealthy returns the IsClusterHealthy field if non-nil, zero value otherwise.

### GetIsClusterHealthyOk

`func (o *NiatelemetryNexusDashboards) GetIsClusterHealthyOk() (*string, bool)`

GetIsClusterHealthyOk returns a tuple with the IsClusterHealthy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClusterHealthy

`func (o *NiatelemetryNexusDashboards) SetIsClusterHealthy(v string)`

SetIsClusterHealthy sets IsClusterHealthy field to given value.

### HasIsClusterHealthy

`func (o *NiatelemetryNexusDashboards) HasIsClusterHealthy() bool`

HasIsClusterHealthy returns a boolean if a field has been set.

### GetNdClusterSize

`func (o *NiatelemetryNexusDashboards) GetNdClusterSize() int64`

GetNdClusterSize returns the NdClusterSize field if non-nil, zero value otherwise.

### GetNdClusterSizeOk

`func (o *NiatelemetryNexusDashboards) GetNdClusterSizeOk() (*int64, bool)`

GetNdClusterSizeOk returns a tuple with the NdClusterSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNdClusterSize

`func (o *NiatelemetryNexusDashboards) SetNdClusterSize(v int64)`

SetNdClusterSize sets NdClusterSize field to given value.

### HasNdClusterSize

`func (o *NiatelemetryNexusDashboards) HasNdClusterSize() bool`

HasNdClusterSize returns a boolean if a field has been set.

### GetNdSites

`func (o *NiatelemetryNexusDashboards) GetNdSites() []NiatelemetrySites`

GetNdSites returns the NdSites field if non-nil, zero value otherwise.

### GetNdSitesOk

`func (o *NiatelemetryNexusDashboards) GetNdSitesOk() (*[]NiatelemetrySites, bool)`

GetNdSitesOk returns a tuple with the NdSites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNdSites

`func (o *NiatelemetryNexusDashboards) SetNdSites(v []NiatelemetrySites)`

SetNdSites sets NdSites field to given value.

### HasNdSites

`func (o *NiatelemetryNexusDashboards) HasNdSites() bool`

HasNdSites returns a boolean if a field has been set.

### SetNdSitesNil

`func (o *NiatelemetryNexusDashboards) SetNdSitesNil(b bool)`

 SetNdSitesNil sets the value for NdSites to be an explicit nil

### UnsetNdSites
`func (o *NiatelemetryNexusDashboards) UnsetNdSites()`

UnsetNdSites ensures that no value is present for NdSites, not even an explicit nil
### GetNdType

`func (o *NiatelemetryNexusDashboards) GetNdType() string`

GetNdType returns the NdType field if non-nil, zero value otherwise.

### GetNdTypeOk

`func (o *NiatelemetryNexusDashboards) GetNdTypeOk() (*string, bool)`

GetNdTypeOk returns a tuple with the NdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNdType

`func (o *NiatelemetryNexusDashboards) SetNdType(v string)`

SetNdType sets NdType field to given value.

### HasNdType

`func (o *NiatelemetryNexusDashboards) HasNdType() bool`

HasNdType returns a boolean if a field has been set.

### GetNdVersion

`func (o *NiatelemetryNexusDashboards) GetNdVersion() string`

GetNdVersion returns the NdVersion field if non-nil, zero value otherwise.

### GetNdVersionOk

`func (o *NiatelemetryNexusDashboards) GetNdVersionOk() (*string, bool)`

GetNdVersionOk returns a tuple with the NdVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNdVersion

`func (o *NiatelemetryNexusDashboards) SetNdVersion(v string)`

SetNdVersion sets NdVersion field to given value.

### HasNdVersion

`func (o *NiatelemetryNexusDashboards) HasNdVersion() bool`

HasNdVersion returns a boolean if a field has been set.

### GetNumberOfApps

`func (o *NiatelemetryNexusDashboards) GetNumberOfApps() int64`

GetNumberOfApps returns the NumberOfApps field if non-nil, zero value otherwise.

### GetNumberOfAppsOk

`func (o *NiatelemetryNexusDashboards) GetNumberOfAppsOk() (*int64, bool)`

GetNumberOfAppsOk returns a tuple with the NumberOfApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfApps

`func (o *NiatelemetryNexusDashboards) SetNumberOfApps(v int64)`

SetNumberOfApps sets NumberOfApps field to given value.

### HasNumberOfApps

`func (o *NiatelemetryNexusDashboards) HasNumberOfApps() bool`

HasNumberOfApps returns a boolean if a field has been set.

### GetNumberOfInsightGroups

`func (o *NiatelemetryNexusDashboards) GetNumberOfInsightGroups() int64`

GetNumberOfInsightGroups returns the NumberOfInsightGroups field if non-nil, zero value otherwise.

### GetNumberOfInsightGroupsOk

`func (o *NiatelemetryNexusDashboards) GetNumberOfInsightGroupsOk() (*int64, bool)`

GetNumberOfInsightGroupsOk returns a tuple with the NumberOfInsightGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfInsightGroups

`func (o *NiatelemetryNexusDashboards) SetNumberOfInsightGroups(v int64)`

SetNumberOfInsightGroups sets NumberOfInsightGroups field to given value.

### HasNumberOfInsightGroups

`func (o *NiatelemetryNexusDashboards) HasNumberOfInsightGroups() bool`

HasNumberOfInsightGroups returns a boolean if a field has been set.

### GetNumberOfNirDashboards

`func (o *NiatelemetryNexusDashboards) GetNumberOfNirDashboards() int64`

GetNumberOfNirDashboards returns the NumberOfNirDashboards field if non-nil, zero value otherwise.

### GetNumberOfNirDashboardsOk

`func (o *NiatelemetryNexusDashboards) GetNumberOfNirDashboardsOk() (*int64, bool)`

GetNumberOfNirDashboardsOk returns a tuple with the NumberOfNirDashboards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfNirDashboards

`func (o *NiatelemetryNexusDashboards) SetNumberOfNirDashboards(v int64)`

SetNumberOfNirDashboards sets NumberOfNirDashboards field to given value.

### HasNumberOfNirDashboards

`func (o *NiatelemetryNexusDashboards) HasNumberOfNirDashboards() bool`

HasNumberOfNirDashboards returns a boolean if a field has been set.

### GetNumberOfSchemasInMso

`func (o *NiatelemetryNexusDashboards) GetNumberOfSchemasInMso() int64`

GetNumberOfSchemasInMso returns the NumberOfSchemasInMso field if non-nil, zero value otherwise.

### GetNumberOfSchemasInMsoOk

`func (o *NiatelemetryNexusDashboards) GetNumberOfSchemasInMsoOk() (*int64, bool)`

GetNumberOfSchemasInMsoOk returns a tuple with the NumberOfSchemasInMso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfSchemasInMso

`func (o *NiatelemetryNexusDashboards) SetNumberOfSchemasInMso(v int64)`

SetNumberOfSchemasInMso sets NumberOfSchemasInMso field to given value.

### HasNumberOfSchemasInMso

`func (o *NiatelemetryNexusDashboards) HasNumberOfSchemasInMso() bool`

HasNumberOfSchemasInMso returns a boolean if a field has been set.

### GetNumberOfSitesInMso

`func (o *NiatelemetryNexusDashboards) GetNumberOfSitesInMso() int64`

GetNumberOfSitesInMso returns the NumberOfSitesInMso field if non-nil, zero value otherwise.

### GetNumberOfSitesInMsoOk

`func (o *NiatelemetryNexusDashboards) GetNumberOfSitesInMsoOk() (*int64, bool)`

GetNumberOfSitesInMsoOk returns a tuple with the NumberOfSitesInMso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfSitesInMso

`func (o *NiatelemetryNexusDashboards) SetNumberOfSitesInMso(v int64)`

SetNumberOfSitesInMso sets NumberOfSitesInMso field to given value.

### HasNumberOfSitesInMso

`func (o *NiatelemetryNexusDashboards) HasNumberOfSitesInMso() bool`

HasNumberOfSitesInMso returns a boolean if a field has been set.

### GetNumberOfSitesServiced

`func (o *NiatelemetryNexusDashboards) GetNumberOfSitesServiced() int64`

GetNumberOfSitesServiced returns the NumberOfSitesServiced field if non-nil, zero value otherwise.

### GetNumberOfSitesServicedOk

`func (o *NiatelemetryNexusDashboards) GetNumberOfSitesServicedOk() (*int64, bool)`

GetNumberOfSitesServicedOk returns a tuple with the NumberOfSitesServiced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfSitesServiced

`func (o *NiatelemetryNexusDashboards) SetNumberOfSitesServiced(v int64)`

SetNumberOfSitesServiced sets NumberOfSitesServiced field to given value.

### HasNumberOfSitesServiced

`func (o *NiatelemetryNexusDashboards) HasNumberOfSitesServiced() bool`

HasNumberOfSitesServiced returns a boolean if a field has been set.

### GetNumberOfTenantsInMso

`func (o *NiatelemetryNexusDashboards) GetNumberOfTenantsInMso() int64`

GetNumberOfTenantsInMso returns the NumberOfTenantsInMso field if non-nil, zero value otherwise.

### GetNumberOfTenantsInMsoOk

`func (o *NiatelemetryNexusDashboards) GetNumberOfTenantsInMsoOk() (*int64, bool)`

GetNumberOfTenantsInMsoOk returns a tuple with the NumberOfTenantsInMso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfTenantsInMso

`func (o *NiatelemetryNexusDashboards) SetNumberOfTenantsInMso(v int64)`

SetNumberOfTenantsInMso sets NumberOfTenantsInMso field to given value.

### HasNumberOfTenantsInMso

`func (o *NiatelemetryNexusDashboards) HasNumberOfTenantsInMso() bool`

HasNumberOfTenantsInMso returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetryNexusDashboards) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryNexusDashboards) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryNexusDashboards) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryNexusDashboards) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetTypeOfSiteInMso

`func (o *NiatelemetryNexusDashboards) GetTypeOfSiteInMso() string`

GetTypeOfSiteInMso returns the TypeOfSiteInMso field if non-nil, zero value otherwise.

### GetTypeOfSiteInMsoOk

`func (o *NiatelemetryNexusDashboards) GetTypeOfSiteInMsoOk() (*string, bool)`

GetTypeOfSiteInMsoOk returns a tuple with the TypeOfSiteInMso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeOfSiteInMso

`func (o *NiatelemetryNexusDashboards) SetTypeOfSiteInMso(v string)`

SetTypeOfSiteInMso sets TypeOfSiteInMso field to given value.

### HasTypeOfSiteInMso

`func (o *NiatelemetryNexusDashboards) HasTypeOfSiteInMso() bool`

HasTypeOfSiteInMso returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryNexusDashboards) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryNexusDashboards) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryNexusDashboards) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryNexusDashboards) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


