# NiatelemetryNexusDashboards

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.NexusDashboards"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.NexusDashboards"]
**BandwidthUsageMonitoring** | Pointer to **bool** | Feature operational state of bandwidth monitoring. | [optional] [readonly] 
**ChangeApprovalCount** | Pointer to **int64** | Count of total Change Control tickets that have been approved and completed. | [optional] [readonly] 
**ChangeRollbackCount** | Pointer to **int64** | Count the number of change control tickets that have been rolled back. | [optional] [readonly] 
**ClusterName** | Pointer to **string** | Nexus Dashboard can onboard multiple APIC clusters/sites. | [optional] 
**ClusterUuid** | Pointer to **string** | UUID of the Nexus Dashboard cluster. | [optional] 
**DeviceSnapshotsCount** | Pointer to **int64** | Count of number of image snapshots taken. | [optional] [readonly] 
**Dn** | Pointer to **string** | Dn of the objects present for Nexus Dashboard devices. | [optional] 
**FabricImagePoliciesCount** | Pointer to **int64** | Count of number of devices with attached image policies. | [optional] [readonly] 
**FeatureOperStatus** | Pointer to **bool** | Feature Operation status of change management. | [optional] 
**ImageFileStagingCount** | Pointer to **int64** | Count of number of image operations of type stage. | [optional] [readonly] 
**IpamOperState** | Pointer to **string** | Feature Operation status of Integration with IPAM. | [optional] [readonly] 
**IsClusterHealthy** | Pointer to **string** | Health of Nexus Dashboard cluster. | [optional] 
**K8VisualizerAdminState** | Pointer to **string** | Feature Operation status of Kubernetes Visualizer. | [optional] [readonly] 
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
**NumberOfVxlanFabricSitesInMso** | Pointer to **int64** | Number of sites with vxLan type fabric in Multi-Site Orchestrator. | [optional] 
**PerformanceMonitoring** | Pointer to **bool** | Feature operational state of performance Monitoring. | [optional] [readonly] 
**PostUpgradeReportGenerationCount** | Pointer to **int64** | Count of post upgrade report generation. | [optional] [readonly] 
**PreUpgradeReportGenerationCount** | Pointer to **int64** | Count of pre upgrade report generation. | [optional] [readonly] 
**PtpAdminState** | Pointer to **string** | Feature Operation status of Precision Time Protocol Monitoring. | [optional] [readonly] 
**RecordType** | Pointer to **string** | Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. | [optional] 
**TypeOfSiteInMso** | Pointer to **string** | Type of site added to Multi-Site Orchestrator. | [optional] 
**VmmVisualizerAdminState** | Pointer to **string** | Feature Operation status of VMM Visualizer. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

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


### GetBandwidthUsageMonitoring

`func (o *NiatelemetryNexusDashboards) GetBandwidthUsageMonitoring() bool`

GetBandwidthUsageMonitoring returns the BandwidthUsageMonitoring field if non-nil, zero value otherwise.

### GetBandwidthUsageMonitoringOk

`func (o *NiatelemetryNexusDashboards) GetBandwidthUsageMonitoringOk() (*bool, bool)`

GetBandwidthUsageMonitoringOk returns a tuple with the BandwidthUsageMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthUsageMonitoring

`func (o *NiatelemetryNexusDashboards) SetBandwidthUsageMonitoring(v bool)`

SetBandwidthUsageMonitoring sets BandwidthUsageMonitoring field to given value.

### HasBandwidthUsageMonitoring

`func (o *NiatelemetryNexusDashboards) HasBandwidthUsageMonitoring() bool`

HasBandwidthUsageMonitoring returns a boolean if a field has been set.

### GetChangeApprovalCount

`func (o *NiatelemetryNexusDashboards) GetChangeApprovalCount() int64`

GetChangeApprovalCount returns the ChangeApprovalCount field if non-nil, zero value otherwise.

### GetChangeApprovalCountOk

`func (o *NiatelemetryNexusDashboards) GetChangeApprovalCountOk() (*int64, bool)`

GetChangeApprovalCountOk returns a tuple with the ChangeApprovalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeApprovalCount

`func (o *NiatelemetryNexusDashboards) SetChangeApprovalCount(v int64)`

SetChangeApprovalCount sets ChangeApprovalCount field to given value.

### HasChangeApprovalCount

`func (o *NiatelemetryNexusDashboards) HasChangeApprovalCount() bool`

HasChangeApprovalCount returns a boolean if a field has been set.

### GetChangeRollbackCount

`func (o *NiatelemetryNexusDashboards) GetChangeRollbackCount() int64`

GetChangeRollbackCount returns the ChangeRollbackCount field if non-nil, zero value otherwise.

### GetChangeRollbackCountOk

`func (o *NiatelemetryNexusDashboards) GetChangeRollbackCountOk() (*int64, bool)`

GetChangeRollbackCountOk returns a tuple with the ChangeRollbackCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeRollbackCount

`func (o *NiatelemetryNexusDashboards) SetChangeRollbackCount(v int64)`

SetChangeRollbackCount sets ChangeRollbackCount field to given value.

### HasChangeRollbackCount

`func (o *NiatelemetryNexusDashboards) HasChangeRollbackCount() bool`

HasChangeRollbackCount returns a boolean if a field has been set.

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

### GetDeviceSnapshotsCount

`func (o *NiatelemetryNexusDashboards) GetDeviceSnapshotsCount() int64`

GetDeviceSnapshotsCount returns the DeviceSnapshotsCount field if non-nil, zero value otherwise.

### GetDeviceSnapshotsCountOk

`func (o *NiatelemetryNexusDashboards) GetDeviceSnapshotsCountOk() (*int64, bool)`

GetDeviceSnapshotsCountOk returns a tuple with the DeviceSnapshotsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSnapshotsCount

`func (o *NiatelemetryNexusDashboards) SetDeviceSnapshotsCount(v int64)`

SetDeviceSnapshotsCount sets DeviceSnapshotsCount field to given value.

### HasDeviceSnapshotsCount

`func (o *NiatelemetryNexusDashboards) HasDeviceSnapshotsCount() bool`

HasDeviceSnapshotsCount returns a boolean if a field has been set.

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

### GetFabricImagePoliciesCount

`func (o *NiatelemetryNexusDashboards) GetFabricImagePoliciesCount() int64`

GetFabricImagePoliciesCount returns the FabricImagePoliciesCount field if non-nil, zero value otherwise.

### GetFabricImagePoliciesCountOk

`func (o *NiatelemetryNexusDashboards) GetFabricImagePoliciesCountOk() (*int64, bool)`

GetFabricImagePoliciesCountOk returns a tuple with the FabricImagePoliciesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricImagePoliciesCount

`func (o *NiatelemetryNexusDashboards) SetFabricImagePoliciesCount(v int64)`

SetFabricImagePoliciesCount sets FabricImagePoliciesCount field to given value.

### HasFabricImagePoliciesCount

`func (o *NiatelemetryNexusDashboards) HasFabricImagePoliciesCount() bool`

HasFabricImagePoliciesCount returns a boolean if a field has been set.

### GetFeatureOperStatus

`func (o *NiatelemetryNexusDashboards) GetFeatureOperStatus() bool`

GetFeatureOperStatus returns the FeatureOperStatus field if non-nil, zero value otherwise.

### GetFeatureOperStatusOk

`func (o *NiatelemetryNexusDashboards) GetFeatureOperStatusOk() (*bool, bool)`

GetFeatureOperStatusOk returns a tuple with the FeatureOperStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureOperStatus

`func (o *NiatelemetryNexusDashboards) SetFeatureOperStatus(v bool)`

SetFeatureOperStatus sets FeatureOperStatus field to given value.

### HasFeatureOperStatus

`func (o *NiatelemetryNexusDashboards) HasFeatureOperStatus() bool`

HasFeatureOperStatus returns a boolean if a field has been set.

### GetImageFileStagingCount

`func (o *NiatelemetryNexusDashboards) GetImageFileStagingCount() int64`

GetImageFileStagingCount returns the ImageFileStagingCount field if non-nil, zero value otherwise.

### GetImageFileStagingCountOk

`func (o *NiatelemetryNexusDashboards) GetImageFileStagingCountOk() (*int64, bool)`

GetImageFileStagingCountOk returns a tuple with the ImageFileStagingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFileStagingCount

`func (o *NiatelemetryNexusDashboards) SetImageFileStagingCount(v int64)`

SetImageFileStagingCount sets ImageFileStagingCount field to given value.

### HasImageFileStagingCount

`func (o *NiatelemetryNexusDashboards) HasImageFileStagingCount() bool`

HasImageFileStagingCount returns a boolean if a field has been set.

### GetIpamOperState

`func (o *NiatelemetryNexusDashboards) GetIpamOperState() string`

GetIpamOperState returns the IpamOperState field if non-nil, zero value otherwise.

### GetIpamOperStateOk

`func (o *NiatelemetryNexusDashboards) GetIpamOperStateOk() (*string, bool)`

GetIpamOperStateOk returns a tuple with the IpamOperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpamOperState

`func (o *NiatelemetryNexusDashboards) SetIpamOperState(v string)`

SetIpamOperState sets IpamOperState field to given value.

### HasIpamOperState

`func (o *NiatelemetryNexusDashboards) HasIpamOperState() bool`

HasIpamOperState returns a boolean if a field has been set.

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

### GetK8VisualizerAdminState

`func (o *NiatelemetryNexusDashboards) GetK8VisualizerAdminState() string`

GetK8VisualizerAdminState returns the K8VisualizerAdminState field if non-nil, zero value otherwise.

### GetK8VisualizerAdminStateOk

`func (o *NiatelemetryNexusDashboards) GetK8VisualizerAdminStateOk() (*string, bool)`

GetK8VisualizerAdminStateOk returns a tuple with the K8VisualizerAdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8VisualizerAdminState

`func (o *NiatelemetryNexusDashboards) SetK8VisualizerAdminState(v string)`

SetK8VisualizerAdminState sets K8VisualizerAdminState field to given value.

### HasK8VisualizerAdminState

`func (o *NiatelemetryNexusDashboards) HasK8VisualizerAdminState() bool`

HasK8VisualizerAdminState returns a boolean if a field has been set.

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

### GetNumberOfVxlanFabricSitesInMso

`func (o *NiatelemetryNexusDashboards) GetNumberOfVxlanFabricSitesInMso() int64`

GetNumberOfVxlanFabricSitesInMso returns the NumberOfVxlanFabricSitesInMso field if non-nil, zero value otherwise.

### GetNumberOfVxlanFabricSitesInMsoOk

`func (o *NiatelemetryNexusDashboards) GetNumberOfVxlanFabricSitesInMsoOk() (*int64, bool)`

GetNumberOfVxlanFabricSitesInMsoOk returns a tuple with the NumberOfVxlanFabricSitesInMso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfVxlanFabricSitesInMso

`func (o *NiatelemetryNexusDashboards) SetNumberOfVxlanFabricSitesInMso(v int64)`

SetNumberOfVxlanFabricSitesInMso sets NumberOfVxlanFabricSitesInMso field to given value.

### HasNumberOfVxlanFabricSitesInMso

`func (o *NiatelemetryNexusDashboards) HasNumberOfVxlanFabricSitesInMso() bool`

HasNumberOfVxlanFabricSitesInMso returns a boolean if a field has been set.

### GetPerformanceMonitoring

`func (o *NiatelemetryNexusDashboards) GetPerformanceMonitoring() bool`

GetPerformanceMonitoring returns the PerformanceMonitoring field if non-nil, zero value otherwise.

### GetPerformanceMonitoringOk

`func (o *NiatelemetryNexusDashboards) GetPerformanceMonitoringOk() (*bool, bool)`

GetPerformanceMonitoringOk returns a tuple with the PerformanceMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformanceMonitoring

`func (o *NiatelemetryNexusDashboards) SetPerformanceMonitoring(v bool)`

SetPerformanceMonitoring sets PerformanceMonitoring field to given value.

### HasPerformanceMonitoring

`func (o *NiatelemetryNexusDashboards) HasPerformanceMonitoring() bool`

HasPerformanceMonitoring returns a boolean if a field has been set.

### GetPostUpgradeReportGenerationCount

`func (o *NiatelemetryNexusDashboards) GetPostUpgradeReportGenerationCount() int64`

GetPostUpgradeReportGenerationCount returns the PostUpgradeReportGenerationCount field if non-nil, zero value otherwise.

### GetPostUpgradeReportGenerationCountOk

`func (o *NiatelemetryNexusDashboards) GetPostUpgradeReportGenerationCountOk() (*int64, bool)`

GetPostUpgradeReportGenerationCountOk returns a tuple with the PostUpgradeReportGenerationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostUpgradeReportGenerationCount

`func (o *NiatelemetryNexusDashboards) SetPostUpgradeReportGenerationCount(v int64)`

SetPostUpgradeReportGenerationCount sets PostUpgradeReportGenerationCount field to given value.

### HasPostUpgradeReportGenerationCount

`func (o *NiatelemetryNexusDashboards) HasPostUpgradeReportGenerationCount() bool`

HasPostUpgradeReportGenerationCount returns a boolean if a field has been set.

### GetPreUpgradeReportGenerationCount

`func (o *NiatelemetryNexusDashboards) GetPreUpgradeReportGenerationCount() int64`

GetPreUpgradeReportGenerationCount returns the PreUpgradeReportGenerationCount field if non-nil, zero value otherwise.

### GetPreUpgradeReportGenerationCountOk

`func (o *NiatelemetryNexusDashboards) GetPreUpgradeReportGenerationCountOk() (*int64, bool)`

GetPreUpgradeReportGenerationCountOk returns a tuple with the PreUpgradeReportGenerationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreUpgradeReportGenerationCount

`func (o *NiatelemetryNexusDashboards) SetPreUpgradeReportGenerationCount(v int64)`

SetPreUpgradeReportGenerationCount sets PreUpgradeReportGenerationCount field to given value.

### HasPreUpgradeReportGenerationCount

`func (o *NiatelemetryNexusDashboards) HasPreUpgradeReportGenerationCount() bool`

HasPreUpgradeReportGenerationCount returns a boolean if a field has been set.

### GetPtpAdminState

`func (o *NiatelemetryNexusDashboards) GetPtpAdminState() string`

GetPtpAdminState returns the PtpAdminState field if non-nil, zero value otherwise.

### GetPtpAdminStateOk

`func (o *NiatelemetryNexusDashboards) GetPtpAdminStateOk() (*string, bool)`

GetPtpAdminStateOk returns a tuple with the PtpAdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtpAdminState

`func (o *NiatelemetryNexusDashboards) SetPtpAdminState(v string)`

SetPtpAdminState sets PtpAdminState field to given value.

### HasPtpAdminState

`func (o *NiatelemetryNexusDashboards) HasPtpAdminState() bool`

HasPtpAdminState returns a boolean if a field has been set.

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

### GetVmmVisualizerAdminState

`func (o *NiatelemetryNexusDashboards) GetVmmVisualizerAdminState() string`

GetVmmVisualizerAdminState returns the VmmVisualizerAdminState field if non-nil, zero value otherwise.

### GetVmmVisualizerAdminStateOk

`func (o *NiatelemetryNexusDashboards) GetVmmVisualizerAdminStateOk() (*string, bool)`

GetVmmVisualizerAdminStateOk returns a tuple with the VmmVisualizerAdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmmVisualizerAdminState

`func (o *NiatelemetryNexusDashboards) SetVmmVisualizerAdminState(v string)`

SetVmmVisualizerAdminState sets VmmVisualizerAdminState field to given value.

### HasVmmVisualizerAdminState

`func (o *NiatelemetryNexusDashboards) HasVmmVisualizerAdminState() bool`

HasVmmVisualizerAdminState returns a boolean if a field has been set.

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

### SetRegisteredDeviceNil

`func (o *NiatelemetryNexusDashboards) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *NiatelemetryNexusDashboards) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


