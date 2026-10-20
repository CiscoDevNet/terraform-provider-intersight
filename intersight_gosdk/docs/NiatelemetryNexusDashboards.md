# NiatelemetryNexusDashboards

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.NexusDashboards"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.NexusDashboards"]
**BackupStatus** | Pointer to **[]string** |  | [optional] 
**BandwidthUsageMonitoring** | Pointer to **bool** | Feature operational state of bandwidth monitoring. | [optional] [readonly] 
**ChangeApprovalCount** | Pointer to **int64** | Count of total Change Control tickets that have been approved and completed. | [optional] [readonly] 
**ChangeRollbackCount** | Pointer to **int64** | Count the number of change control tickets that have been rolled back. | [optional] [readonly] 
**ClusterName** | Pointer to **string** | Nexus Dashboard can onboard multiple APIC clusters/sites. | [optional] 
**ClusterUuid** | Pointer to **string** | UUID of the Nexus Dashboard cluster. | [optional] 
**ComplianceRules** | Pointer to **int64** | Number of compliance rules on the fabric. | [optional] [readonly] 
**DashboardCount** | Pointer to **int64** | Number of custom dashboard in the fabric. | [optional] [readonly] 
**DeviceSnapshotsCount** | Pointer to **int64** | Count of number of image snapshots taken. | [optional] [readonly] 
**Dn** | Pointer to **string** | Dn of the objects present for Nexus Dashboard devices. | [optional] 
**EndpointCount** | Pointer to **int64** | Total number of endpoints on fabric. | [optional] [readonly] 
**FabricImagePoliciesCount** | Pointer to **int64** | Count of number of devices with attached image policies. | [optional] [readonly] 
**FeatureOperStatus** | Pointer to **bool** | Feature Operation status of change management. | [optional] 
**ImageFileStagingCount** | Pointer to **int64** | Count of number of image operations of type stage. | [optional] [readonly] 
**IpamOperState** | Pointer to **string** | Feature Operation status of Integration with IPAM. | [optional] [readonly] 
**IsClusterHealthy** | Pointer to **string** | Health of Nexus Dashboard cluster. | [optional] 
**K8VisualizerAdminState** | Pointer to **string** | Feature Operation status of Kubernetes Visualizer. | [optional] [readonly] 
**LatestVersionList** | Pointer to **[]string** |  | [optional] 
**LiveProtectEnabledCount** | Pointer to **int64** | Count of devices with Live Protect shield status enabled. | [optional] 
**MulticastRouteCount** | Pointer to **int64** | Number of multicast routes on fabric. | [optional] [readonly] 
**NdClusterSize** | Pointer to **int64** | Number of nodes in Nexus Dashboard cluster. | [optional] 
**NdHealthy** | Pointer to **bool** | Health status of the Nexus Dashboard cluster. | [optional] [readonly] 
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
**OamEnabled** | Pointer to **[]bool** |  | [optional] 
**PerformanceMonitoring** | Pointer to **bool** | Feature operational state of performance Monitoring. | [optional] [readonly] 
**PostUpgradeReportGenerationCount** | Pointer to **int64** | Count of post upgrade report generation. | [optional] [readonly] 
**PreUpgradeReportGenerationCount** | Pointer to **int64** | Count of pre upgrade report generation. | [optional] [readonly] 
**PreupgradeValidationCount** | Pointer to **int64** | Number of pre-upgrade validations on the fabric. | [optional] [readonly] 
**PtpAdminState** | Pointer to **string** | Feature Operation status of Precision Time Protocol Monitoring. | [optional] [readonly] 
**RecVersionList** | Pointer to **[]string** |  | [optional] 
**RecordType** | Pointer to **string** | Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. | [optional] 
**ReleaseVersion** | Pointer to **[]string** |  | [optional] 
**SustainabilityReportStatus** | Pointer to **string** | Status of sustainability report on fabric. | [optional] [readonly] 
**TypeOfSiteInMso** | Pointer to **string** | Type of site added to Multi-Site Orchestrator. | [optional] 
**VcenterCount** | Pointer to **int64** | Number of vCenters integrated into the fabric. | [optional] [readonly] 
**VmmVisualizerAdminState** | Pointer to **string** | Feature Operation status of VMM Visualizer. | [optional] [readonly] 
**VxLanFabCount** | Pointer to **int64** | The total number of active VXLAN-managed fabrics that have both leaf and spine switches configured. | [optional] [readonly] 
**VxLanFabNames** | Pointer to **[]string** |  | [optional] 
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


### GetBackupStatus

`func (o *NiatelemetryNexusDashboards) GetBackupStatus() []string`

GetBackupStatus returns the BackupStatus field if non-nil, zero value otherwise.

### GetBackupStatusOk

`func (o *NiatelemetryNexusDashboards) GetBackupStatusOk() (*[]string, bool)`

GetBackupStatusOk returns a tuple with the BackupStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupStatus

`func (o *NiatelemetryNexusDashboards) SetBackupStatus(v []string)`

SetBackupStatus sets BackupStatus field to given value.

### HasBackupStatus

`func (o *NiatelemetryNexusDashboards) HasBackupStatus() bool`

HasBackupStatus returns a boolean if a field has been set.

### SetBackupStatusNil

`func (o *NiatelemetryNexusDashboards) SetBackupStatusNil(b bool)`

 SetBackupStatusNil sets the value for BackupStatus to be an explicit nil

### UnsetBackupStatus
`func (o *NiatelemetryNexusDashboards) UnsetBackupStatus()`

UnsetBackupStatus ensures that no value is present for BackupStatus, not even an explicit nil
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

### GetComplianceRules

`func (o *NiatelemetryNexusDashboards) GetComplianceRules() int64`

GetComplianceRules returns the ComplianceRules field if non-nil, zero value otherwise.

### GetComplianceRulesOk

`func (o *NiatelemetryNexusDashboards) GetComplianceRulesOk() (*int64, bool)`

GetComplianceRulesOk returns a tuple with the ComplianceRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceRules

`func (o *NiatelemetryNexusDashboards) SetComplianceRules(v int64)`

SetComplianceRules sets ComplianceRules field to given value.

### HasComplianceRules

`func (o *NiatelemetryNexusDashboards) HasComplianceRules() bool`

HasComplianceRules returns a boolean if a field has been set.

### GetDashboardCount

`func (o *NiatelemetryNexusDashboards) GetDashboardCount() int64`

GetDashboardCount returns the DashboardCount field if non-nil, zero value otherwise.

### GetDashboardCountOk

`func (o *NiatelemetryNexusDashboards) GetDashboardCountOk() (*int64, bool)`

GetDashboardCountOk returns a tuple with the DashboardCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardCount

`func (o *NiatelemetryNexusDashboards) SetDashboardCount(v int64)`

SetDashboardCount sets DashboardCount field to given value.

### HasDashboardCount

`func (o *NiatelemetryNexusDashboards) HasDashboardCount() bool`

HasDashboardCount returns a boolean if a field has been set.

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

### GetEndpointCount

`func (o *NiatelemetryNexusDashboards) GetEndpointCount() int64`

GetEndpointCount returns the EndpointCount field if non-nil, zero value otherwise.

### GetEndpointCountOk

`func (o *NiatelemetryNexusDashboards) GetEndpointCountOk() (*int64, bool)`

GetEndpointCountOk returns a tuple with the EndpointCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointCount

`func (o *NiatelemetryNexusDashboards) SetEndpointCount(v int64)`

SetEndpointCount sets EndpointCount field to given value.

### HasEndpointCount

`func (o *NiatelemetryNexusDashboards) HasEndpointCount() bool`

HasEndpointCount returns a boolean if a field has been set.

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

### GetLatestVersionList

`func (o *NiatelemetryNexusDashboards) GetLatestVersionList() []string`

GetLatestVersionList returns the LatestVersionList field if non-nil, zero value otherwise.

### GetLatestVersionListOk

`func (o *NiatelemetryNexusDashboards) GetLatestVersionListOk() (*[]string, bool)`

GetLatestVersionListOk returns a tuple with the LatestVersionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersionList

`func (o *NiatelemetryNexusDashboards) SetLatestVersionList(v []string)`

SetLatestVersionList sets LatestVersionList field to given value.

### HasLatestVersionList

`func (o *NiatelemetryNexusDashboards) HasLatestVersionList() bool`

HasLatestVersionList returns a boolean if a field has been set.

### SetLatestVersionListNil

`func (o *NiatelemetryNexusDashboards) SetLatestVersionListNil(b bool)`

 SetLatestVersionListNil sets the value for LatestVersionList to be an explicit nil

### UnsetLatestVersionList
`func (o *NiatelemetryNexusDashboards) UnsetLatestVersionList()`

UnsetLatestVersionList ensures that no value is present for LatestVersionList, not even an explicit nil
### GetLiveProtectEnabledCount

`func (o *NiatelemetryNexusDashboards) GetLiveProtectEnabledCount() int64`

GetLiveProtectEnabledCount returns the LiveProtectEnabledCount field if non-nil, zero value otherwise.

### GetLiveProtectEnabledCountOk

`func (o *NiatelemetryNexusDashboards) GetLiveProtectEnabledCountOk() (*int64, bool)`

GetLiveProtectEnabledCountOk returns a tuple with the LiveProtectEnabledCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveProtectEnabledCount

`func (o *NiatelemetryNexusDashboards) SetLiveProtectEnabledCount(v int64)`

SetLiveProtectEnabledCount sets LiveProtectEnabledCount field to given value.

### HasLiveProtectEnabledCount

`func (o *NiatelemetryNexusDashboards) HasLiveProtectEnabledCount() bool`

HasLiveProtectEnabledCount returns a boolean if a field has been set.

### GetMulticastRouteCount

`func (o *NiatelemetryNexusDashboards) GetMulticastRouteCount() int64`

GetMulticastRouteCount returns the MulticastRouteCount field if non-nil, zero value otherwise.

### GetMulticastRouteCountOk

`func (o *NiatelemetryNexusDashboards) GetMulticastRouteCountOk() (*int64, bool)`

GetMulticastRouteCountOk returns a tuple with the MulticastRouteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulticastRouteCount

`func (o *NiatelemetryNexusDashboards) SetMulticastRouteCount(v int64)`

SetMulticastRouteCount sets MulticastRouteCount field to given value.

### HasMulticastRouteCount

`func (o *NiatelemetryNexusDashboards) HasMulticastRouteCount() bool`

HasMulticastRouteCount returns a boolean if a field has been set.

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

### GetNdHealthy

`func (o *NiatelemetryNexusDashboards) GetNdHealthy() bool`

GetNdHealthy returns the NdHealthy field if non-nil, zero value otherwise.

### GetNdHealthyOk

`func (o *NiatelemetryNexusDashboards) GetNdHealthyOk() (*bool, bool)`

GetNdHealthyOk returns a tuple with the NdHealthy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNdHealthy

`func (o *NiatelemetryNexusDashboards) SetNdHealthy(v bool)`

SetNdHealthy sets NdHealthy field to given value.

### HasNdHealthy

`func (o *NiatelemetryNexusDashboards) HasNdHealthy() bool`

HasNdHealthy returns a boolean if a field has been set.

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

### GetOamEnabled

`func (o *NiatelemetryNexusDashboards) GetOamEnabled() []bool`

GetOamEnabled returns the OamEnabled field if non-nil, zero value otherwise.

### GetOamEnabledOk

`func (o *NiatelemetryNexusDashboards) GetOamEnabledOk() (*[]bool, bool)`

GetOamEnabledOk returns a tuple with the OamEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOamEnabled

`func (o *NiatelemetryNexusDashboards) SetOamEnabled(v []bool)`

SetOamEnabled sets OamEnabled field to given value.

### HasOamEnabled

`func (o *NiatelemetryNexusDashboards) HasOamEnabled() bool`

HasOamEnabled returns a boolean if a field has been set.

### SetOamEnabledNil

`func (o *NiatelemetryNexusDashboards) SetOamEnabledNil(b bool)`

 SetOamEnabledNil sets the value for OamEnabled to be an explicit nil

### UnsetOamEnabled
`func (o *NiatelemetryNexusDashboards) UnsetOamEnabled()`

UnsetOamEnabled ensures that no value is present for OamEnabled, not even an explicit nil
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

### GetPreupgradeValidationCount

`func (o *NiatelemetryNexusDashboards) GetPreupgradeValidationCount() int64`

GetPreupgradeValidationCount returns the PreupgradeValidationCount field if non-nil, zero value otherwise.

### GetPreupgradeValidationCountOk

`func (o *NiatelemetryNexusDashboards) GetPreupgradeValidationCountOk() (*int64, bool)`

GetPreupgradeValidationCountOk returns a tuple with the PreupgradeValidationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreupgradeValidationCount

`func (o *NiatelemetryNexusDashboards) SetPreupgradeValidationCount(v int64)`

SetPreupgradeValidationCount sets PreupgradeValidationCount field to given value.

### HasPreupgradeValidationCount

`func (o *NiatelemetryNexusDashboards) HasPreupgradeValidationCount() bool`

HasPreupgradeValidationCount returns a boolean if a field has been set.

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

### GetRecVersionList

`func (o *NiatelemetryNexusDashboards) GetRecVersionList() []string`

GetRecVersionList returns the RecVersionList field if non-nil, zero value otherwise.

### GetRecVersionListOk

`func (o *NiatelemetryNexusDashboards) GetRecVersionListOk() (*[]string, bool)`

GetRecVersionListOk returns a tuple with the RecVersionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecVersionList

`func (o *NiatelemetryNexusDashboards) SetRecVersionList(v []string)`

SetRecVersionList sets RecVersionList field to given value.

### HasRecVersionList

`func (o *NiatelemetryNexusDashboards) HasRecVersionList() bool`

HasRecVersionList returns a boolean if a field has been set.

### SetRecVersionListNil

`func (o *NiatelemetryNexusDashboards) SetRecVersionListNil(b bool)`

 SetRecVersionListNil sets the value for RecVersionList to be an explicit nil

### UnsetRecVersionList
`func (o *NiatelemetryNexusDashboards) UnsetRecVersionList()`

UnsetRecVersionList ensures that no value is present for RecVersionList, not even an explicit nil
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

### GetReleaseVersion

`func (o *NiatelemetryNexusDashboards) GetReleaseVersion() []string`

GetReleaseVersion returns the ReleaseVersion field if non-nil, zero value otherwise.

### GetReleaseVersionOk

`func (o *NiatelemetryNexusDashboards) GetReleaseVersionOk() (*[]string, bool)`

GetReleaseVersionOk returns a tuple with the ReleaseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseVersion

`func (o *NiatelemetryNexusDashboards) SetReleaseVersion(v []string)`

SetReleaseVersion sets ReleaseVersion field to given value.

### HasReleaseVersion

`func (o *NiatelemetryNexusDashboards) HasReleaseVersion() bool`

HasReleaseVersion returns a boolean if a field has been set.

### SetReleaseVersionNil

`func (o *NiatelemetryNexusDashboards) SetReleaseVersionNil(b bool)`

 SetReleaseVersionNil sets the value for ReleaseVersion to be an explicit nil

### UnsetReleaseVersion
`func (o *NiatelemetryNexusDashboards) UnsetReleaseVersion()`

UnsetReleaseVersion ensures that no value is present for ReleaseVersion, not even an explicit nil
### GetSustainabilityReportStatus

`func (o *NiatelemetryNexusDashboards) GetSustainabilityReportStatus() string`

GetSustainabilityReportStatus returns the SustainabilityReportStatus field if non-nil, zero value otherwise.

### GetSustainabilityReportStatusOk

`func (o *NiatelemetryNexusDashboards) GetSustainabilityReportStatusOk() (*string, bool)`

GetSustainabilityReportStatusOk returns a tuple with the SustainabilityReportStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSustainabilityReportStatus

`func (o *NiatelemetryNexusDashboards) SetSustainabilityReportStatus(v string)`

SetSustainabilityReportStatus sets SustainabilityReportStatus field to given value.

### HasSustainabilityReportStatus

`func (o *NiatelemetryNexusDashboards) HasSustainabilityReportStatus() bool`

HasSustainabilityReportStatus returns a boolean if a field has been set.

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

### GetVcenterCount

`func (o *NiatelemetryNexusDashboards) GetVcenterCount() int64`

GetVcenterCount returns the VcenterCount field if non-nil, zero value otherwise.

### GetVcenterCountOk

`func (o *NiatelemetryNexusDashboards) GetVcenterCountOk() (*int64, bool)`

GetVcenterCountOk returns a tuple with the VcenterCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcenterCount

`func (o *NiatelemetryNexusDashboards) SetVcenterCount(v int64)`

SetVcenterCount sets VcenterCount field to given value.

### HasVcenterCount

`func (o *NiatelemetryNexusDashboards) HasVcenterCount() bool`

HasVcenterCount returns a boolean if a field has been set.

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

### GetVxLanFabCount

`func (o *NiatelemetryNexusDashboards) GetVxLanFabCount() int64`

GetVxLanFabCount returns the VxLanFabCount field if non-nil, zero value otherwise.

### GetVxLanFabCountOk

`func (o *NiatelemetryNexusDashboards) GetVxLanFabCountOk() (*int64, bool)`

GetVxLanFabCountOk returns a tuple with the VxLanFabCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVxLanFabCount

`func (o *NiatelemetryNexusDashboards) SetVxLanFabCount(v int64)`

SetVxLanFabCount sets VxLanFabCount field to given value.

### HasVxLanFabCount

`func (o *NiatelemetryNexusDashboards) HasVxLanFabCount() bool`

HasVxLanFabCount returns a boolean if a field has been set.

### GetVxLanFabNames

`func (o *NiatelemetryNexusDashboards) GetVxLanFabNames() []string`

GetVxLanFabNames returns the VxLanFabNames field if non-nil, zero value otherwise.

### GetVxLanFabNamesOk

`func (o *NiatelemetryNexusDashboards) GetVxLanFabNamesOk() (*[]string, bool)`

GetVxLanFabNamesOk returns a tuple with the VxLanFabNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVxLanFabNames

`func (o *NiatelemetryNexusDashboards) SetVxLanFabNames(v []string)`

SetVxLanFabNames sets VxLanFabNames field to given value.

### HasVxLanFabNames

`func (o *NiatelemetryNexusDashboards) HasVxLanFabNames() bool`

HasVxLanFabNames returns a boolean if a field has been set.

### SetVxLanFabNamesNil

`func (o *NiatelemetryNexusDashboards) SetVxLanFabNamesNil(b bool)`

 SetVxLanFabNamesNil sets the value for VxLanFabNames to be an explicit nil

### UnsetVxLanFabNames
`func (o *NiatelemetryNexusDashboards) UnsetVxLanFabNames()`

UnsetVxLanFabNames ensures that no value is present for VxLanFabNames, not even an explicit nil
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


