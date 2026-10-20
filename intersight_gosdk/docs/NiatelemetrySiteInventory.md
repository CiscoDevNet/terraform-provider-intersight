# NiatelemetrySiteInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.SiteInventory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.SiteInventory"]
**Apps** | Pointer to **[]string** |  | [optional] 
**ConfigurationChangeTrackingCount** | Pointer to **int64** | Count of configuration change tracking. | [optional] [readonly] 
**ConnectivityAnalysisCount** | Pointer to **int64** | Returns the total number of connectivity Analysis run for EPs in NDFC Fabrics. | [optional] 
**EndpointLocatorCount** | Pointer to **int64** | Count of total Endpoint Locators. | [optional] [readonly] 
**FabricTechnology** | Pointer to **string** | Fabric technology reported by the onboarded DCNM site. | [optional] 
**FabricUpdateStatus** | Pointer to **string** | Status of the fabric update. | [optional] [readonly] 
**FirmwareVersion** | Pointer to **string** | Version of the specified site. | [optional] 
**InstallType** | Pointer to **string** | Fine-grained type DCNM either SAN or LAN. | [optional] 
**IpAddress** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** | Name of the APIC / DCNM site onboarded. | [optional] 
**NexusDashboard** | Pointer to **string** | Name of ND on which site has been onboarded. | [optional] 
**Nodes** | Pointer to **int64** | Number of nodes the site contains. | [optional] 
**OnDemandBackups** | Pointer to **bool** | Count of number of manual backups. | [optional] [readonly] 
**PerimeterService** | Pointer to **int64** | Count of service functions configured with use case Perimeter Service. | [optional] [readonly] 
**RecordType** | Pointer to **string** | Specifies whether Site object is DCNM or APIC or ND. | [optional] 
**RedirectToServiceChain** | Pointer to **int64** | Count of service functions configured with use case Service Chain Redirection. | [optional] [readonly] 
**ScheduledBackups** | Pointer to **bool** | Count of number of scheduled backups. | [optional] [readonly] 
**SecurityGroupsCount** | Pointer to **int64** | Count of total security groups. | [optional] [readonly] 
**ServiceAsGateway** | Pointer to **int64** | Count of service functions configured with use case Service As Default Gateway. | [optional] [readonly] 
**ServiceClustersCount** | Pointer to **int64** | Count of total Service Clusters. | [optional] [readonly] 
**ServiceFunctionCount** | Pointer to **int64** | Count of total Service Functions configured. | [optional] [readonly] 
**ServiceInsertionCount** | Pointer to **int64** | Count of total Service Function Insertions enabled. | [optional] [readonly] 
**SumCount** | Pointer to **int64** | Sum of latestVersionCount and recommendedVersionCount. | [optional] [readonly] 
**SwitchCount** | Pointer to **int64** | Count of switches in the fabric. | [optional] [readonly] 
**TotalNetworks** | Pointer to **int64** | Count of total Networks on the fabric. | [optional] [readonly] 
**TotalVrfs** | Pointer to **int64** | Count of total VRFs on the fabric. | [optional] [readonly] 
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
### GetConfigurationChangeTrackingCount

`func (o *NiatelemetrySiteInventory) GetConfigurationChangeTrackingCount() int64`

GetConfigurationChangeTrackingCount returns the ConfigurationChangeTrackingCount field if non-nil, zero value otherwise.

### GetConfigurationChangeTrackingCountOk

`func (o *NiatelemetrySiteInventory) GetConfigurationChangeTrackingCountOk() (*int64, bool)`

GetConfigurationChangeTrackingCountOk returns a tuple with the ConfigurationChangeTrackingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationChangeTrackingCount

`func (o *NiatelemetrySiteInventory) SetConfigurationChangeTrackingCount(v int64)`

SetConfigurationChangeTrackingCount sets ConfigurationChangeTrackingCount field to given value.

### HasConfigurationChangeTrackingCount

`func (o *NiatelemetrySiteInventory) HasConfigurationChangeTrackingCount() bool`

HasConfigurationChangeTrackingCount returns a boolean if a field has been set.

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

### GetEndpointLocatorCount

`func (o *NiatelemetrySiteInventory) GetEndpointLocatorCount() int64`

GetEndpointLocatorCount returns the EndpointLocatorCount field if non-nil, zero value otherwise.

### GetEndpointLocatorCountOk

`func (o *NiatelemetrySiteInventory) GetEndpointLocatorCountOk() (*int64, bool)`

GetEndpointLocatorCountOk returns a tuple with the EndpointLocatorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointLocatorCount

`func (o *NiatelemetrySiteInventory) SetEndpointLocatorCount(v int64)`

SetEndpointLocatorCount sets EndpointLocatorCount field to given value.

### HasEndpointLocatorCount

`func (o *NiatelemetrySiteInventory) HasEndpointLocatorCount() bool`

HasEndpointLocatorCount returns a boolean if a field has been set.

### GetFabricTechnology

`func (o *NiatelemetrySiteInventory) GetFabricTechnology() string`

GetFabricTechnology returns the FabricTechnology field if non-nil, zero value otherwise.

### GetFabricTechnologyOk

`func (o *NiatelemetrySiteInventory) GetFabricTechnologyOk() (*string, bool)`

GetFabricTechnologyOk returns a tuple with the FabricTechnology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricTechnology

`func (o *NiatelemetrySiteInventory) SetFabricTechnology(v string)`

SetFabricTechnology sets FabricTechnology field to given value.

### HasFabricTechnology

`func (o *NiatelemetrySiteInventory) HasFabricTechnology() bool`

HasFabricTechnology returns a boolean if a field has been set.

### GetFabricUpdateStatus

`func (o *NiatelemetrySiteInventory) GetFabricUpdateStatus() string`

GetFabricUpdateStatus returns the FabricUpdateStatus field if non-nil, zero value otherwise.

### GetFabricUpdateStatusOk

`func (o *NiatelemetrySiteInventory) GetFabricUpdateStatusOk() (*string, bool)`

GetFabricUpdateStatusOk returns a tuple with the FabricUpdateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricUpdateStatus

`func (o *NiatelemetrySiteInventory) SetFabricUpdateStatus(v string)`

SetFabricUpdateStatus sets FabricUpdateStatus field to given value.

### HasFabricUpdateStatus

`func (o *NiatelemetrySiteInventory) HasFabricUpdateStatus() bool`

HasFabricUpdateStatus returns a boolean if a field has been set.

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

### GetOnDemandBackups

`func (o *NiatelemetrySiteInventory) GetOnDemandBackups() bool`

GetOnDemandBackups returns the OnDemandBackups field if non-nil, zero value otherwise.

### GetOnDemandBackupsOk

`func (o *NiatelemetrySiteInventory) GetOnDemandBackupsOk() (*bool, bool)`

GetOnDemandBackupsOk returns a tuple with the OnDemandBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDemandBackups

`func (o *NiatelemetrySiteInventory) SetOnDemandBackups(v bool)`

SetOnDemandBackups sets OnDemandBackups field to given value.

### HasOnDemandBackups

`func (o *NiatelemetrySiteInventory) HasOnDemandBackups() bool`

HasOnDemandBackups returns a boolean if a field has been set.

### GetPerimeterService

`func (o *NiatelemetrySiteInventory) GetPerimeterService() int64`

GetPerimeterService returns the PerimeterService field if non-nil, zero value otherwise.

### GetPerimeterServiceOk

`func (o *NiatelemetrySiteInventory) GetPerimeterServiceOk() (*int64, bool)`

GetPerimeterServiceOk returns a tuple with the PerimeterService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerimeterService

`func (o *NiatelemetrySiteInventory) SetPerimeterService(v int64)`

SetPerimeterService sets PerimeterService field to given value.

### HasPerimeterService

`func (o *NiatelemetrySiteInventory) HasPerimeterService() bool`

HasPerimeterService returns a boolean if a field has been set.

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

### GetRedirectToServiceChain

`func (o *NiatelemetrySiteInventory) GetRedirectToServiceChain() int64`

GetRedirectToServiceChain returns the RedirectToServiceChain field if non-nil, zero value otherwise.

### GetRedirectToServiceChainOk

`func (o *NiatelemetrySiteInventory) GetRedirectToServiceChainOk() (*int64, bool)`

GetRedirectToServiceChainOk returns a tuple with the RedirectToServiceChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectToServiceChain

`func (o *NiatelemetrySiteInventory) SetRedirectToServiceChain(v int64)`

SetRedirectToServiceChain sets RedirectToServiceChain field to given value.

### HasRedirectToServiceChain

`func (o *NiatelemetrySiteInventory) HasRedirectToServiceChain() bool`

HasRedirectToServiceChain returns a boolean if a field has been set.

### GetScheduledBackups

`func (o *NiatelemetrySiteInventory) GetScheduledBackups() bool`

GetScheduledBackups returns the ScheduledBackups field if non-nil, zero value otherwise.

### GetScheduledBackupsOk

`func (o *NiatelemetrySiteInventory) GetScheduledBackupsOk() (*bool, bool)`

GetScheduledBackupsOk returns a tuple with the ScheduledBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledBackups

`func (o *NiatelemetrySiteInventory) SetScheduledBackups(v bool)`

SetScheduledBackups sets ScheduledBackups field to given value.

### HasScheduledBackups

`func (o *NiatelemetrySiteInventory) HasScheduledBackups() bool`

HasScheduledBackups returns a boolean if a field has been set.

### GetSecurityGroupsCount

`func (o *NiatelemetrySiteInventory) GetSecurityGroupsCount() int64`

GetSecurityGroupsCount returns the SecurityGroupsCount field if non-nil, zero value otherwise.

### GetSecurityGroupsCountOk

`func (o *NiatelemetrySiteInventory) GetSecurityGroupsCountOk() (*int64, bool)`

GetSecurityGroupsCountOk returns a tuple with the SecurityGroupsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupsCount

`func (o *NiatelemetrySiteInventory) SetSecurityGroupsCount(v int64)`

SetSecurityGroupsCount sets SecurityGroupsCount field to given value.

### HasSecurityGroupsCount

`func (o *NiatelemetrySiteInventory) HasSecurityGroupsCount() bool`

HasSecurityGroupsCount returns a boolean if a field has been set.

### GetServiceAsGateway

`func (o *NiatelemetrySiteInventory) GetServiceAsGateway() int64`

GetServiceAsGateway returns the ServiceAsGateway field if non-nil, zero value otherwise.

### GetServiceAsGatewayOk

`func (o *NiatelemetrySiteInventory) GetServiceAsGatewayOk() (*int64, bool)`

GetServiceAsGatewayOk returns a tuple with the ServiceAsGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAsGateway

`func (o *NiatelemetrySiteInventory) SetServiceAsGateway(v int64)`

SetServiceAsGateway sets ServiceAsGateway field to given value.

### HasServiceAsGateway

`func (o *NiatelemetrySiteInventory) HasServiceAsGateway() bool`

HasServiceAsGateway returns a boolean if a field has been set.

### GetServiceClustersCount

`func (o *NiatelemetrySiteInventory) GetServiceClustersCount() int64`

GetServiceClustersCount returns the ServiceClustersCount field if non-nil, zero value otherwise.

### GetServiceClustersCountOk

`func (o *NiatelemetrySiteInventory) GetServiceClustersCountOk() (*int64, bool)`

GetServiceClustersCountOk returns a tuple with the ServiceClustersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceClustersCount

`func (o *NiatelemetrySiteInventory) SetServiceClustersCount(v int64)`

SetServiceClustersCount sets ServiceClustersCount field to given value.

### HasServiceClustersCount

`func (o *NiatelemetrySiteInventory) HasServiceClustersCount() bool`

HasServiceClustersCount returns a boolean if a field has been set.

### GetServiceFunctionCount

`func (o *NiatelemetrySiteInventory) GetServiceFunctionCount() int64`

GetServiceFunctionCount returns the ServiceFunctionCount field if non-nil, zero value otherwise.

### GetServiceFunctionCountOk

`func (o *NiatelemetrySiteInventory) GetServiceFunctionCountOk() (*int64, bool)`

GetServiceFunctionCountOk returns a tuple with the ServiceFunctionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceFunctionCount

`func (o *NiatelemetrySiteInventory) SetServiceFunctionCount(v int64)`

SetServiceFunctionCount sets ServiceFunctionCount field to given value.

### HasServiceFunctionCount

`func (o *NiatelemetrySiteInventory) HasServiceFunctionCount() bool`

HasServiceFunctionCount returns a boolean if a field has been set.

### GetServiceInsertionCount

`func (o *NiatelemetrySiteInventory) GetServiceInsertionCount() int64`

GetServiceInsertionCount returns the ServiceInsertionCount field if non-nil, zero value otherwise.

### GetServiceInsertionCountOk

`func (o *NiatelemetrySiteInventory) GetServiceInsertionCountOk() (*int64, bool)`

GetServiceInsertionCountOk returns a tuple with the ServiceInsertionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceInsertionCount

`func (o *NiatelemetrySiteInventory) SetServiceInsertionCount(v int64)`

SetServiceInsertionCount sets ServiceInsertionCount field to given value.

### HasServiceInsertionCount

`func (o *NiatelemetrySiteInventory) HasServiceInsertionCount() bool`

HasServiceInsertionCount returns a boolean if a field has been set.

### GetSumCount

`func (o *NiatelemetrySiteInventory) GetSumCount() int64`

GetSumCount returns the SumCount field if non-nil, zero value otherwise.

### GetSumCountOk

`func (o *NiatelemetrySiteInventory) GetSumCountOk() (*int64, bool)`

GetSumCountOk returns a tuple with the SumCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumCount

`func (o *NiatelemetrySiteInventory) SetSumCount(v int64)`

SetSumCount sets SumCount field to given value.

### HasSumCount

`func (o *NiatelemetrySiteInventory) HasSumCount() bool`

HasSumCount returns a boolean if a field has been set.

### GetSwitchCount

`func (o *NiatelemetrySiteInventory) GetSwitchCount() int64`

GetSwitchCount returns the SwitchCount field if non-nil, zero value otherwise.

### GetSwitchCountOk

`func (o *NiatelemetrySiteInventory) GetSwitchCountOk() (*int64, bool)`

GetSwitchCountOk returns a tuple with the SwitchCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchCount

`func (o *NiatelemetrySiteInventory) SetSwitchCount(v int64)`

SetSwitchCount sets SwitchCount field to given value.

### HasSwitchCount

`func (o *NiatelemetrySiteInventory) HasSwitchCount() bool`

HasSwitchCount returns a boolean if a field has been set.

### GetTotalNetworks

`func (o *NiatelemetrySiteInventory) GetTotalNetworks() int64`

GetTotalNetworks returns the TotalNetworks field if non-nil, zero value otherwise.

### GetTotalNetworksOk

`func (o *NiatelemetrySiteInventory) GetTotalNetworksOk() (*int64, bool)`

GetTotalNetworksOk returns a tuple with the TotalNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNetworks

`func (o *NiatelemetrySiteInventory) SetTotalNetworks(v int64)`

SetTotalNetworks sets TotalNetworks field to given value.

### HasTotalNetworks

`func (o *NiatelemetrySiteInventory) HasTotalNetworks() bool`

HasTotalNetworks returns a boolean if a field has been set.

### GetTotalVrfs

`func (o *NiatelemetrySiteInventory) GetTotalVrfs() int64`

GetTotalVrfs returns the TotalVrfs field if non-nil, zero value otherwise.

### GetTotalVrfsOk

`func (o *NiatelemetrySiteInventory) GetTotalVrfsOk() (*int64, bool)`

GetTotalVrfsOk returns a tuple with the TotalVrfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVrfs

`func (o *NiatelemetrySiteInventory) SetTotalVrfs(v int64)`

SetTotalVrfs sets TotalVrfs field to given value.

### HasTotalVrfs

`func (o *NiatelemetrySiteInventory) HasTotalVrfs() bool`

HasTotalVrfs returns a boolean if a field has been set.

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


