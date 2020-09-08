# NiatelemetryNiaFeatureUsageAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApicCount** | Pointer to **int64** | Number of APIC controllers. This determines the value of controllers for the fabric. | [optional] 
**AppCenterCount** | Pointer to **int64** | ACI APPs feature usage scale. | [optional] 
**Ave** | Pointer to **string** | AVE feature usage. This determines if ACI virtual edge feature is enabled or disabled. | [optional] 
**BdCount** | Pointer to **int64** | Number of BDs. This determines the total number of Broadcast Domains across the fabric. | [optional] 
**CloudSecPeerCount** | Pointer to **int64** | Number of Cloudsec SA peers. | [optional] 
**ConsistencyCheckerApp** | Pointer to **string** | Consistency checker application usage. This determines if the fabric has Consistency checker application installed. | [optional] 
**ContractCount** | Pointer to **int64** | Number of contracts. This determines the total number of Contracts configured across the fabric. | [optional] 
**DnsCount** | Pointer to **int64** | DNS feature usage. This determines the total number of DNS configurations across the fabric. | [optional] 
**EigrpCount** | Pointer to **int64** | Eigrp feature usage. This determines the total number of EIGRP sessions across the fabric. | [optional] 
**EpgCount** | Pointer to **int64** | Number of End Point Groups. This determines the total number of End Point Groups across the fabric. | [optional] 
**FcoeNportCount** | Pointer to **int64** | Total number of FCoE N-Port for DOM, VSAn, and VLAN. | [optional] 
**FcoeNportDomCount** | Pointer to **int64** | Number of FCoE N-Port DOM. | [optional] 
**FcoeNportVlanCount** | Pointer to **int64** | Number of FCoE N-Port VLAN. | [optional] 
**FcoeNportVsanCount** | Pointer to **int64** | Number of FCoE N-Port VSAN. | [optional] 
**HsrpCount** | Pointer to **int64** | Hsrp feature usage. This determines the total number of HSRP sessions across the fabric. | [optional] 
**IbgpCount** | Pointer to **int64** | Ibgp feature usage. This determines the total number of BGP sessions across the fabric. | [optional] 
**IgmpAccessListCount** | Pointer to **int64** | IGMP Access List feature usage. This determines the total number of IGMP access lists configured across the fabric. | [optional] 
**IgmpSnoop** | Pointer to **string** | IGMP Snooping feature usage. This determines if this feature is enabled or disabled. | [optional] 
**IpEpgCount** | Pointer to **int64** | Number of IP based End Point Groups. This determines the total number of IP End Point Groups across the fabric. | [optional] 
**IsisCount** | Pointer to **int64** | Isis feature usage. This determines the total number of ISIS sessions across the fabric. | [optional] 
**L2Multicast** | Pointer to **string** | L2Multicast feature usage. This determines if this Layer 2 Multicast feature is being enabled / disabled on the fabric. | [optional] 
**LeafCount** | Pointer to **int64** | Number of Leafs. This determines the total number of Leaf switches in the fabric. | [optional] 
**MaintenanceModeCount** | Pointer to **int64** | Maintenance Mode feature usage. This determines the number of switches that are currently in maintenance mode. | [optional] 
**ManagementOverV6Count** | Pointer to **int64** | Management over IPv6 feature usage. This determines the total number of IPv6 configurtaions in the fabric. | [optional] 
**MicrosoftUsegVmmEpPdCount** | Pointer to **int64** | Microsoft microsegmentation VmmEpPD scale. Ensures that Microsoft was configured. | [optional] 
**NetFlowCount** | Pointer to **int64** | Number of Netflow monitor policies. | [optional] 
**Nir** | Pointer to **string** | NIR application usage. This determines if the fabric has NIR application installed. | [optional] 
**OpenStack** | Pointer to **string** | Open stack feature usage. | [optional] 
**OpflexKubernetesCount** | Pointer to **int64** | Opflex for Kubernetes feature usage. This determines the total number of VMM sessions of type kubernetes. | [optional] 
**OspfCount** | Pointer to **int64** | Ospf feature usage. This determines the total number of OSPF sessions across the fabric. | [optional] 
**PoeCount** | Pointer to **int64** | POE feature usage. This determines the total number of POE configurations across the fabric. | [optional] 
**PortSecurityCount** | Pointer to **int64** | Port Security count scale. Non-Zero value indicates the object as enabled. | [optional] 
**QinVniTunnelCount** | Pointer to **int64** | QinVniTunnel feature usage. This determines if the qinVniTunnel feature is being used on the fabric and the scale of it. | [optional] 
**RemoteLeafCount** | Pointer to **int64** | Number of remote Leafs. This determines the total number of remote leaf switches in the fabric. | [optional] 
**ScvmmCount** | Pointer to **int64** | SCVMM feature usage. This determines the total number of SCVMM configurations in the fabric. | [optional] 
**SharedL3OutCount** | Pointer to **int64** | SharedL3Out feature usage. This determines the total number of Shared L3 out configured across the fabric. | [optional] 
**SiteName** | Pointer to **string** | The Site name represents an APIC cluster. Service Engine can onboard multiple APIC clusters / sites. There will be a feature usage object per site in Multi site scenario. In multi-site scenario the site name is available in all the requests being made. | [optional] 
**SmartCallHome** | Pointer to **string** | Smart callhome feature usage. This determines if this feature is being enabled or disabled. | [optional] 
**Snmp** | Pointer to **string** | SNMP feature usage. This determines if this feature is enabled or disabled. | [optional] 
**SpanCount** | Pointer to **int64** | Number of Span Sources and Destinations. | [optional] 
**SpanDstCount** | Pointer to **int64** | Number of Span Destinations with valid state. | [optional] 
**SpanSrcCount** | Pointer to **int64** | Number of Span Sources with valid state. | [optional] 
**SpineCount** | Pointer to **int64** | Number of Spines. This determines the total number of spine switches in the fabric. | [optional] 
**SshOverV6Count** | Pointer to **int64** | Ssh over IPv6 feature usage. This determines the total number of IPv6 configurtaions in the fabric. | [optional] 
**SyslogOverV6Count** | Pointer to **int64** | Syslog over IPv6 feature usage. This determines the total number of IPv6 configurtaions in the fabric. | [optional] 
**TenantCount** | Pointer to **int64** | Number of tenants. This determines the total number of tenants configured across the fabric. | [optional] 
**TierTwoLeafCount** | Pointer to **int64** | Number of tier 2 Leafs. This determines the total number of tier 2 Leaf switches in the fabric. | [optional] 
**Twamp** | Pointer to **string** | TWAMP feature usage. This determines if this feature is enabled or disabled. | [optional] 
**Useg** | Pointer to **string** | VMM uSegmentation feature usage. This determines if microsegmentation feature is enabled or disabled. | [optional] 
**VmWareVdsCount** | Pointer to **int64** | VmWare vCenter 6.5 support count scale. Checks the controller revision value. | [optional] 
**VmmCtrlrpCount** | Pointer to **int64** | Gets the scale for Virtual Machine Monitor controller policy for VMware vCenter. | [optional] 
**VmmDompCount** | Pointer to **int64** | Obtains the scale for Virtual Machine Monitor domain policy model for VMware vCenter. | [optional] 
**VmmEpPdCount** | Pointer to **int64** | Microsegmentation Distributed Virtual Switch feature usage. Gets the scale for VMware vCenter. | [optional] 
**VnsmDevCount** | Pointer to **int64** | L4-L7 Device Package Import count scale. Checks for the vendor and the model. | [optional] 
**VpodCount** | Pointer to **int64** | Virtual pod feature usage. This determines the total number of virtual POD configurations in the fabrics. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewNiatelemetryNiaFeatureUsageAllOf

`func NewNiatelemetryNiaFeatureUsageAllOf() *NiatelemetryNiaFeatureUsageAllOf`

NewNiatelemetryNiaFeatureUsageAllOf instantiates a new NiatelemetryNiaFeatureUsageAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryNiaFeatureUsageAllOfWithDefaults

`func NewNiatelemetryNiaFeatureUsageAllOfWithDefaults() *NiatelemetryNiaFeatureUsageAllOf`

NewNiatelemetryNiaFeatureUsageAllOfWithDefaults instantiates a new NiatelemetryNiaFeatureUsageAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApicCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetApicCount() int64`

GetApicCount returns the ApicCount field if non-nil, zero value otherwise.

### GetApicCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetApicCountOk() (*int64, bool)`

GetApicCountOk returns a tuple with the ApicCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApicCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetApicCount(v int64)`

SetApicCount sets ApicCount field to given value.

### HasApicCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasApicCount() bool`

HasApicCount returns a boolean if a field has been set.

### GetAppCenterCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetAppCenterCount() int64`

GetAppCenterCount returns the AppCenterCount field if non-nil, zero value otherwise.

### GetAppCenterCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetAppCenterCountOk() (*int64, bool)`

GetAppCenterCountOk returns a tuple with the AppCenterCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppCenterCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetAppCenterCount(v int64)`

SetAppCenterCount sets AppCenterCount field to given value.

### HasAppCenterCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasAppCenterCount() bool`

HasAppCenterCount returns a boolean if a field has been set.

### GetAve

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetAve() string`

GetAve returns the Ave field if non-nil, zero value otherwise.

### GetAveOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetAveOk() (*string, bool)`

GetAveOk returns a tuple with the Ave field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAve

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetAve(v string)`

SetAve sets Ave field to given value.

### HasAve

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasAve() bool`

HasAve returns a boolean if a field has been set.

### GetBdCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetBdCount() int64`

GetBdCount returns the BdCount field if non-nil, zero value otherwise.

### GetBdCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetBdCountOk() (*int64, bool)`

GetBdCountOk returns a tuple with the BdCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBdCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetBdCount(v int64)`

SetBdCount sets BdCount field to given value.

### HasBdCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasBdCount() bool`

HasBdCount returns a boolean if a field has been set.

### GetCloudSecPeerCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetCloudSecPeerCount() int64`

GetCloudSecPeerCount returns the CloudSecPeerCount field if non-nil, zero value otherwise.

### GetCloudSecPeerCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetCloudSecPeerCountOk() (*int64, bool)`

GetCloudSecPeerCountOk returns a tuple with the CloudSecPeerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudSecPeerCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetCloudSecPeerCount(v int64)`

SetCloudSecPeerCount sets CloudSecPeerCount field to given value.

### HasCloudSecPeerCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasCloudSecPeerCount() bool`

HasCloudSecPeerCount returns a boolean if a field has been set.

### GetConsistencyCheckerApp

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetConsistencyCheckerApp() string`

GetConsistencyCheckerApp returns the ConsistencyCheckerApp field if non-nil, zero value otherwise.

### GetConsistencyCheckerAppOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetConsistencyCheckerAppOk() (*string, bool)`

GetConsistencyCheckerAppOk returns a tuple with the ConsistencyCheckerApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsistencyCheckerApp

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetConsistencyCheckerApp(v string)`

SetConsistencyCheckerApp sets ConsistencyCheckerApp field to given value.

### HasConsistencyCheckerApp

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasConsistencyCheckerApp() bool`

HasConsistencyCheckerApp returns a boolean if a field has been set.

### GetContractCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetContractCount() int64`

GetContractCount returns the ContractCount field if non-nil, zero value otherwise.

### GetContractCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetContractCountOk() (*int64, bool)`

GetContractCountOk returns a tuple with the ContractCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetContractCount(v int64)`

SetContractCount sets ContractCount field to given value.

### HasContractCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasContractCount() bool`

HasContractCount returns a boolean if a field has been set.

### GetDnsCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetDnsCount() int64`

GetDnsCount returns the DnsCount field if non-nil, zero value otherwise.

### GetDnsCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetDnsCountOk() (*int64, bool)`

GetDnsCountOk returns a tuple with the DnsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetDnsCount(v int64)`

SetDnsCount sets DnsCount field to given value.

### HasDnsCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasDnsCount() bool`

HasDnsCount returns a boolean if a field has been set.

### GetEigrpCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetEigrpCount() int64`

GetEigrpCount returns the EigrpCount field if non-nil, zero value otherwise.

### GetEigrpCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetEigrpCountOk() (*int64, bool)`

GetEigrpCountOk returns a tuple with the EigrpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEigrpCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetEigrpCount(v int64)`

SetEigrpCount sets EigrpCount field to given value.

### HasEigrpCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasEigrpCount() bool`

HasEigrpCount returns a boolean if a field has been set.

### GetEpgCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetEpgCount() int64`

GetEpgCount returns the EpgCount field if non-nil, zero value otherwise.

### GetEpgCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetEpgCountOk() (*int64, bool)`

GetEpgCountOk returns a tuple with the EpgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpgCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetEpgCount(v int64)`

SetEpgCount sets EpgCount field to given value.

### HasEpgCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasEpgCount() bool`

HasEpgCount returns a boolean if a field has been set.

### GetFcoeNportCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetFcoeNportCount() int64`

GetFcoeNportCount returns the FcoeNportCount field if non-nil, zero value otherwise.

### GetFcoeNportCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetFcoeNportCountOk() (*int64, bool)`

GetFcoeNportCountOk returns a tuple with the FcoeNportCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcoeNportCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetFcoeNportCount(v int64)`

SetFcoeNportCount sets FcoeNportCount field to given value.

### HasFcoeNportCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasFcoeNportCount() bool`

HasFcoeNportCount returns a boolean if a field has been set.

### GetFcoeNportDomCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetFcoeNportDomCount() int64`

GetFcoeNportDomCount returns the FcoeNportDomCount field if non-nil, zero value otherwise.

### GetFcoeNportDomCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetFcoeNportDomCountOk() (*int64, bool)`

GetFcoeNportDomCountOk returns a tuple with the FcoeNportDomCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcoeNportDomCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetFcoeNportDomCount(v int64)`

SetFcoeNportDomCount sets FcoeNportDomCount field to given value.

### HasFcoeNportDomCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasFcoeNportDomCount() bool`

HasFcoeNportDomCount returns a boolean if a field has been set.

### GetFcoeNportVlanCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetFcoeNportVlanCount() int64`

GetFcoeNportVlanCount returns the FcoeNportVlanCount field if non-nil, zero value otherwise.

### GetFcoeNportVlanCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetFcoeNportVlanCountOk() (*int64, bool)`

GetFcoeNportVlanCountOk returns a tuple with the FcoeNportVlanCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcoeNportVlanCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetFcoeNportVlanCount(v int64)`

SetFcoeNportVlanCount sets FcoeNportVlanCount field to given value.

### HasFcoeNportVlanCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasFcoeNportVlanCount() bool`

HasFcoeNportVlanCount returns a boolean if a field has been set.

### GetFcoeNportVsanCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetFcoeNportVsanCount() int64`

GetFcoeNportVsanCount returns the FcoeNportVsanCount field if non-nil, zero value otherwise.

### GetFcoeNportVsanCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetFcoeNportVsanCountOk() (*int64, bool)`

GetFcoeNportVsanCountOk returns a tuple with the FcoeNportVsanCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcoeNportVsanCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetFcoeNportVsanCount(v int64)`

SetFcoeNportVsanCount sets FcoeNportVsanCount field to given value.

### HasFcoeNportVsanCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasFcoeNportVsanCount() bool`

HasFcoeNportVsanCount returns a boolean if a field has been set.

### GetHsrpCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetHsrpCount() int64`

GetHsrpCount returns the HsrpCount field if non-nil, zero value otherwise.

### GetHsrpCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetHsrpCountOk() (*int64, bool)`

GetHsrpCountOk returns a tuple with the HsrpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsrpCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetHsrpCount(v int64)`

SetHsrpCount sets HsrpCount field to given value.

### HasHsrpCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasHsrpCount() bool`

HasHsrpCount returns a boolean if a field has been set.

### GetIbgpCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetIbgpCount() int64`

GetIbgpCount returns the IbgpCount field if non-nil, zero value otherwise.

### GetIbgpCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetIbgpCountOk() (*int64, bool)`

GetIbgpCountOk returns a tuple with the IbgpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbgpCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetIbgpCount(v int64)`

SetIbgpCount sets IbgpCount field to given value.

### HasIbgpCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasIbgpCount() bool`

HasIbgpCount returns a boolean if a field has been set.

### GetIgmpAccessListCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetIgmpAccessListCount() int64`

GetIgmpAccessListCount returns the IgmpAccessListCount field if non-nil, zero value otherwise.

### GetIgmpAccessListCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetIgmpAccessListCountOk() (*int64, bool)`

GetIgmpAccessListCountOk returns a tuple with the IgmpAccessListCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgmpAccessListCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetIgmpAccessListCount(v int64)`

SetIgmpAccessListCount sets IgmpAccessListCount field to given value.

### HasIgmpAccessListCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasIgmpAccessListCount() bool`

HasIgmpAccessListCount returns a boolean if a field has been set.

### GetIgmpSnoop

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetIgmpSnoop() string`

GetIgmpSnoop returns the IgmpSnoop field if non-nil, zero value otherwise.

### GetIgmpSnoopOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetIgmpSnoopOk() (*string, bool)`

GetIgmpSnoopOk returns a tuple with the IgmpSnoop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgmpSnoop

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetIgmpSnoop(v string)`

SetIgmpSnoop sets IgmpSnoop field to given value.

### HasIgmpSnoop

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasIgmpSnoop() bool`

HasIgmpSnoop returns a boolean if a field has been set.

### GetIpEpgCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetIpEpgCount() int64`

GetIpEpgCount returns the IpEpgCount field if non-nil, zero value otherwise.

### GetIpEpgCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetIpEpgCountOk() (*int64, bool)`

GetIpEpgCountOk returns a tuple with the IpEpgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpEpgCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetIpEpgCount(v int64)`

SetIpEpgCount sets IpEpgCount field to given value.

### HasIpEpgCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasIpEpgCount() bool`

HasIpEpgCount returns a boolean if a field has been set.

### GetIsisCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetIsisCount() int64`

GetIsisCount returns the IsisCount field if non-nil, zero value otherwise.

### GetIsisCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetIsisCountOk() (*int64, bool)`

GetIsisCountOk returns a tuple with the IsisCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsisCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetIsisCount(v int64)`

SetIsisCount sets IsisCount field to given value.

### HasIsisCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasIsisCount() bool`

HasIsisCount returns a boolean if a field has been set.

### GetL2Multicast

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetL2Multicast() string`

GetL2Multicast returns the L2Multicast field if non-nil, zero value otherwise.

### GetL2MulticastOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetL2MulticastOk() (*string, bool)`

GetL2MulticastOk returns a tuple with the L2Multicast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL2Multicast

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetL2Multicast(v string)`

SetL2Multicast sets L2Multicast field to given value.

### HasL2Multicast

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasL2Multicast() bool`

HasL2Multicast returns a boolean if a field has been set.

### GetLeafCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetLeafCount() int64`

GetLeafCount returns the LeafCount field if non-nil, zero value otherwise.

### GetLeafCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetLeafCountOk() (*int64, bool)`

GetLeafCountOk returns a tuple with the LeafCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeafCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetLeafCount(v int64)`

SetLeafCount sets LeafCount field to given value.

### HasLeafCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasLeafCount() bool`

HasLeafCount returns a boolean if a field has been set.

### GetMaintenanceModeCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetMaintenanceModeCount() int64`

GetMaintenanceModeCount returns the MaintenanceModeCount field if non-nil, zero value otherwise.

### GetMaintenanceModeCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetMaintenanceModeCountOk() (*int64, bool)`

GetMaintenanceModeCountOk returns a tuple with the MaintenanceModeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceModeCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetMaintenanceModeCount(v int64)`

SetMaintenanceModeCount sets MaintenanceModeCount field to given value.

### HasMaintenanceModeCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasMaintenanceModeCount() bool`

HasMaintenanceModeCount returns a boolean if a field has been set.

### GetManagementOverV6Count

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetManagementOverV6Count() int64`

GetManagementOverV6Count returns the ManagementOverV6Count field if non-nil, zero value otherwise.

### GetManagementOverV6CountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetManagementOverV6CountOk() (*int64, bool)`

GetManagementOverV6CountOk returns a tuple with the ManagementOverV6Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementOverV6Count

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetManagementOverV6Count(v int64)`

SetManagementOverV6Count sets ManagementOverV6Count field to given value.

### HasManagementOverV6Count

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasManagementOverV6Count() bool`

HasManagementOverV6Count returns a boolean if a field has been set.

### GetMicrosoftUsegVmmEpPdCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetMicrosoftUsegVmmEpPdCount() int64`

GetMicrosoftUsegVmmEpPdCount returns the MicrosoftUsegVmmEpPdCount field if non-nil, zero value otherwise.

### GetMicrosoftUsegVmmEpPdCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetMicrosoftUsegVmmEpPdCountOk() (*int64, bool)`

GetMicrosoftUsegVmmEpPdCountOk returns a tuple with the MicrosoftUsegVmmEpPdCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicrosoftUsegVmmEpPdCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetMicrosoftUsegVmmEpPdCount(v int64)`

SetMicrosoftUsegVmmEpPdCount sets MicrosoftUsegVmmEpPdCount field to given value.

### HasMicrosoftUsegVmmEpPdCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasMicrosoftUsegVmmEpPdCount() bool`

HasMicrosoftUsegVmmEpPdCount returns a boolean if a field has been set.

### GetNetFlowCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetNetFlowCount() int64`

GetNetFlowCount returns the NetFlowCount field if non-nil, zero value otherwise.

### GetNetFlowCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetNetFlowCountOk() (*int64, bool)`

GetNetFlowCountOk returns a tuple with the NetFlowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetFlowCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetNetFlowCount(v int64)`

SetNetFlowCount sets NetFlowCount field to given value.

### HasNetFlowCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasNetFlowCount() bool`

HasNetFlowCount returns a boolean if a field has been set.

### GetNir

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetNir() string`

GetNir returns the Nir field if non-nil, zero value otherwise.

### GetNirOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetNirOk() (*string, bool)`

GetNirOk returns a tuple with the Nir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNir

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetNir(v string)`

SetNir sets Nir field to given value.

### HasNir

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasNir() bool`

HasNir returns a boolean if a field has been set.

### GetOpenStack

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetOpenStack() string`

GetOpenStack returns the OpenStack field if non-nil, zero value otherwise.

### GetOpenStackOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetOpenStackOk() (*string, bool)`

GetOpenStackOk returns a tuple with the OpenStack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStack

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetOpenStack(v string)`

SetOpenStack sets OpenStack field to given value.

### HasOpenStack

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasOpenStack() bool`

HasOpenStack returns a boolean if a field has been set.

### GetOpflexKubernetesCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetOpflexKubernetesCount() int64`

GetOpflexKubernetesCount returns the OpflexKubernetesCount field if non-nil, zero value otherwise.

### GetOpflexKubernetesCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetOpflexKubernetesCountOk() (*int64, bool)`

GetOpflexKubernetesCountOk returns a tuple with the OpflexKubernetesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpflexKubernetesCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetOpflexKubernetesCount(v int64)`

SetOpflexKubernetesCount sets OpflexKubernetesCount field to given value.

### HasOpflexKubernetesCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasOpflexKubernetesCount() bool`

HasOpflexKubernetesCount returns a boolean if a field has been set.

### GetOspfCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetOspfCount() int64`

GetOspfCount returns the OspfCount field if non-nil, zero value otherwise.

### GetOspfCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetOspfCountOk() (*int64, bool)`

GetOspfCountOk returns a tuple with the OspfCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspfCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetOspfCount(v int64)`

SetOspfCount sets OspfCount field to given value.

### HasOspfCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasOspfCount() bool`

HasOspfCount returns a boolean if a field has been set.

### GetPoeCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetPoeCount() int64`

GetPoeCount returns the PoeCount field if non-nil, zero value otherwise.

### GetPoeCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetPoeCountOk() (*int64, bool)`

GetPoeCountOk returns a tuple with the PoeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetPoeCount(v int64)`

SetPoeCount sets PoeCount field to given value.

### HasPoeCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasPoeCount() bool`

HasPoeCount returns a boolean if a field has been set.

### GetPortSecurityCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetPortSecurityCount() int64`

GetPortSecurityCount returns the PortSecurityCount field if non-nil, zero value otherwise.

### GetPortSecurityCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetPortSecurityCountOk() (*int64, bool)`

GetPortSecurityCountOk returns a tuple with the PortSecurityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortSecurityCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetPortSecurityCount(v int64)`

SetPortSecurityCount sets PortSecurityCount field to given value.

### HasPortSecurityCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasPortSecurityCount() bool`

HasPortSecurityCount returns a boolean if a field has been set.

### GetQinVniTunnelCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetQinVniTunnelCount() int64`

GetQinVniTunnelCount returns the QinVniTunnelCount field if non-nil, zero value otherwise.

### GetQinVniTunnelCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetQinVniTunnelCountOk() (*int64, bool)`

GetQinVniTunnelCountOk returns a tuple with the QinVniTunnelCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQinVniTunnelCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetQinVniTunnelCount(v int64)`

SetQinVniTunnelCount sets QinVniTunnelCount field to given value.

### HasQinVniTunnelCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasQinVniTunnelCount() bool`

HasQinVniTunnelCount returns a boolean if a field has been set.

### GetRemoteLeafCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetRemoteLeafCount() int64`

GetRemoteLeafCount returns the RemoteLeafCount field if non-nil, zero value otherwise.

### GetRemoteLeafCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetRemoteLeafCountOk() (*int64, bool)`

GetRemoteLeafCountOk returns a tuple with the RemoteLeafCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteLeafCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetRemoteLeafCount(v int64)`

SetRemoteLeafCount sets RemoteLeafCount field to given value.

### HasRemoteLeafCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasRemoteLeafCount() bool`

HasRemoteLeafCount returns a boolean if a field has been set.

### GetScvmmCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetScvmmCount() int64`

GetScvmmCount returns the ScvmmCount field if non-nil, zero value otherwise.

### GetScvmmCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetScvmmCountOk() (*int64, bool)`

GetScvmmCountOk returns a tuple with the ScvmmCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScvmmCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetScvmmCount(v int64)`

SetScvmmCount sets ScvmmCount field to given value.

### HasScvmmCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasScvmmCount() bool`

HasScvmmCount returns a boolean if a field has been set.

### GetSharedL3OutCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetSharedL3OutCount() int64`

GetSharedL3OutCount returns the SharedL3OutCount field if non-nil, zero value otherwise.

### GetSharedL3OutCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetSharedL3OutCountOk() (*int64, bool)`

GetSharedL3OutCountOk returns a tuple with the SharedL3OutCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedL3OutCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetSharedL3OutCount(v int64)`

SetSharedL3OutCount sets SharedL3OutCount field to given value.

### HasSharedL3OutCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasSharedL3OutCount() bool`

HasSharedL3OutCount returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetSmartCallHome

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetSmartCallHome() string`

GetSmartCallHome returns the SmartCallHome field if non-nil, zero value otherwise.

### GetSmartCallHomeOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetSmartCallHomeOk() (*string, bool)`

GetSmartCallHomeOk returns a tuple with the SmartCallHome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartCallHome

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetSmartCallHome(v string)`

SetSmartCallHome sets SmartCallHome field to given value.

### HasSmartCallHome

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasSmartCallHome() bool`

HasSmartCallHome returns a boolean if a field has been set.

### GetSnmp

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetSnmp() string`

GetSnmp returns the Snmp field if non-nil, zero value otherwise.

### GetSnmpOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetSnmpOk() (*string, bool)`

GetSnmpOk returns a tuple with the Snmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmp

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetSnmp(v string)`

SetSnmp sets Snmp field to given value.

### HasSnmp

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasSnmp() bool`

HasSnmp returns a boolean if a field has been set.

### GetSpanCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetSpanCount() int64`

GetSpanCount returns the SpanCount field if non-nil, zero value otherwise.

### GetSpanCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetSpanCountOk() (*int64, bool)`

GetSpanCountOk returns a tuple with the SpanCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetSpanCount(v int64)`

SetSpanCount sets SpanCount field to given value.

### HasSpanCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasSpanCount() bool`

HasSpanCount returns a boolean if a field has been set.

### GetSpanDstCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetSpanDstCount() int64`

GetSpanDstCount returns the SpanDstCount field if non-nil, zero value otherwise.

### GetSpanDstCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetSpanDstCountOk() (*int64, bool)`

GetSpanDstCountOk returns a tuple with the SpanDstCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanDstCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetSpanDstCount(v int64)`

SetSpanDstCount sets SpanDstCount field to given value.

### HasSpanDstCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasSpanDstCount() bool`

HasSpanDstCount returns a boolean if a field has been set.

### GetSpanSrcCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetSpanSrcCount() int64`

GetSpanSrcCount returns the SpanSrcCount field if non-nil, zero value otherwise.

### GetSpanSrcCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetSpanSrcCountOk() (*int64, bool)`

GetSpanSrcCountOk returns a tuple with the SpanSrcCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanSrcCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetSpanSrcCount(v int64)`

SetSpanSrcCount sets SpanSrcCount field to given value.

### HasSpanSrcCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasSpanSrcCount() bool`

HasSpanSrcCount returns a boolean if a field has been set.

### GetSpineCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetSpineCount() int64`

GetSpineCount returns the SpineCount field if non-nil, zero value otherwise.

### GetSpineCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetSpineCountOk() (*int64, bool)`

GetSpineCountOk returns a tuple with the SpineCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpineCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetSpineCount(v int64)`

SetSpineCount sets SpineCount field to given value.

### HasSpineCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasSpineCount() bool`

HasSpineCount returns a boolean if a field has been set.

### GetSshOverV6Count

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetSshOverV6Count() int64`

GetSshOverV6Count returns the SshOverV6Count field if non-nil, zero value otherwise.

### GetSshOverV6CountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetSshOverV6CountOk() (*int64, bool)`

GetSshOverV6CountOk returns a tuple with the SshOverV6Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshOverV6Count

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetSshOverV6Count(v int64)`

SetSshOverV6Count sets SshOverV6Count field to given value.

### HasSshOverV6Count

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasSshOverV6Count() bool`

HasSshOverV6Count returns a boolean if a field has been set.

### GetSyslogOverV6Count

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetSyslogOverV6Count() int64`

GetSyslogOverV6Count returns the SyslogOverV6Count field if non-nil, zero value otherwise.

### GetSyslogOverV6CountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetSyslogOverV6CountOk() (*int64, bool)`

GetSyslogOverV6CountOk returns a tuple with the SyslogOverV6Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogOverV6Count

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetSyslogOverV6Count(v int64)`

SetSyslogOverV6Count sets SyslogOverV6Count field to given value.

### HasSyslogOverV6Count

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasSyslogOverV6Count() bool`

HasSyslogOverV6Count returns a boolean if a field has been set.

### GetTenantCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetTenantCount() int64`

GetTenantCount returns the TenantCount field if non-nil, zero value otherwise.

### GetTenantCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetTenantCountOk() (*int64, bool)`

GetTenantCountOk returns a tuple with the TenantCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetTenantCount(v int64)`

SetTenantCount sets TenantCount field to given value.

### HasTenantCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasTenantCount() bool`

HasTenantCount returns a boolean if a field has been set.

### GetTierTwoLeafCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetTierTwoLeafCount() int64`

GetTierTwoLeafCount returns the TierTwoLeafCount field if non-nil, zero value otherwise.

### GetTierTwoLeafCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetTierTwoLeafCountOk() (*int64, bool)`

GetTierTwoLeafCountOk returns a tuple with the TierTwoLeafCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierTwoLeafCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetTierTwoLeafCount(v int64)`

SetTierTwoLeafCount sets TierTwoLeafCount field to given value.

### HasTierTwoLeafCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasTierTwoLeafCount() bool`

HasTierTwoLeafCount returns a boolean if a field has been set.

### GetTwamp

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetTwamp() string`

GetTwamp returns the Twamp field if non-nil, zero value otherwise.

### GetTwampOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetTwampOk() (*string, bool)`

GetTwampOk returns a tuple with the Twamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwamp

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetTwamp(v string)`

SetTwamp sets Twamp field to given value.

### HasTwamp

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasTwamp() bool`

HasTwamp returns a boolean if a field has been set.

### GetUseg

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetUseg() string`

GetUseg returns the Useg field if non-nil, zero value otherwise.

### GetUsegOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetUsegOk() (*string, bool)`

GetUsegOk returns a tuple with the Useg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseg

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetUseg(v string)`

SetUseg sets Useg field to given value.

### HasUseg

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasUseg() bool`

HasUseg returns a boolean if a field has been set.

### GetVmWareVdsCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetVmWareVdsCount() int64`

GetVmWareVdsCount returns the VmWareVdsCount field if non-nil, zero value otherwise.

### GetVmWareVdsCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetVmWareVdsCountOk() (*int64, bool)`

GetVmWareVdsCountOk returns a tuple with the VmWareVdsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmWareVdsCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetVmWareVdsCount(v int64)`

SetVmWareVdsCount sets VmWareVdsCount field to given value.

### HasVmWareVdsCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasVmWareVdsCount() bool`

HasVmWareVdsCount returns a boolean if a field has been set.

### GetVmmCtrlrpCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetVmmCtrlrpCount() int64`

GetVmmCtrlrpCount returns the VmmCtrlrpCount field if non-nil, zero value otherwise.

### GetVmmCtrlrpCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetVmmCtrlrpCountOk() (*int64, bool)`

GetVmmCtrlrpCountOk returns a tuple with the VmmCtrlrpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmmCtrlrpCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetVmmCtrlrpCount(v int64)`

SetVmmCtrlrpCount sets VmmCtrlrpCount field to given value.

### HasVmmCtrlrpCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasVmmCtrlrpCount() bool`

HasVmmCtrlrpCount returns a boolean if a field has been set.

### GetVmmDompCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetVmmDompCount() int64`

GetVmmDompCount returns the VmmDompCount field if non-nil, zero value otherwise.

### GetVmmDompCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetVmmDompCountOk() (*int64, bool)`

GetVmmDompCountOk returns a tuple with the VmmDompCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmmDompCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetVmmDompCount(v int64)`

SetVmmDompCount sets VmmDompCount field to given value.

### HasVmmDompCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasVmmDompCount() bool`

HasVmmDompCount returns a boolean if a field has been set.

### GetVmmEpPdCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetVmmEpPdCount() int64`

GetVmmEpPdCount returns the VmmEpPdCount field if non-nil, zero value otherwise.

### GetVmmEpPdCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetVmmEpPdCountOk() (*int64, bool)`

GetVmmEpPdCountOk returns a tuple with the VmmEpPdCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmmEpPdCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetVmmEpPdCount(v int64)`

SetVmmEpPdCount sets VmmEpPdCount field to given value.

### HasVmmEpPdCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasVmmEpPdCount() bool`

HasVmmEpPdCount returns a boolean if a field has been set.

### GetVnsmDevCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetVnsmDevCount() int64`

GetVnsmDevCount returns the VnsmDevCount field if non-nil, zero value otherwise.

### GetVnsmDevCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetVnsmDevCountOk() (*int64, bool)`

GetVnsmDevCountOk returns a tuple with the VnsmDevCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnsmDevCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetVnsmDevCount(v int64)`

SetVnsmDevCount sets VnsmDevCount field to given value.

### HasVnsmDevCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasVnsmDevCount() bool`

HasVnsmDevCount returns a boolean if a field has been set.

### GetVpodCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetVpodCount() int64`

GetVpodCount returns the VpodCount field if non-nil, zero value otherwise.

### GetVpodCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetVpodCountOk() (*int64, bool)`

GetVpodCountOk returns a tuple with the VpodCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpodCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetVpodCount(v int64)`

SetVpodCount sets VpodCount field to given value.

### HasVpodCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasVpodCount() bool`

HasVpodCount returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


