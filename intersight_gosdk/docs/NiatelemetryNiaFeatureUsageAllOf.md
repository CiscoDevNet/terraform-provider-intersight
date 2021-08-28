# NiatelemetryNiaFeatureUsageAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.NiaFeatureUsage"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.NiaFeatureUsage"]
**AaaLdapProviderCount** | Pointer to **int64** | Returns the total number of AAA Ldap Providers. | [optional] 
**AaaRadiusProviderCount** | Pointer to **int64** | Returns the total number of AAA Radius Providers. | [optional] 
**AaaTacacsProviderCount** | Pointer to **int64** | Returns the total number of AAA Tacacs Providers. | [optional] 
**ApicCount** | Pointer to **int64** | Number of APIC controllers. This determines the value of controllers for the fabric. | [optional] 
**ApicIsTelnetEnabled** | Pointer to **bool** | Returns if telnet is enabled on APIC. | [optional] 
**ApicNtpCount** | Pointer to **int64** | Count of NTP servers configured on APIC. | [optional] 
**ApicSnmpCommunityCount** | Pointer to **int64** | Number of SNMP communities configured on APIC. | [optional] 
**ApicSysLogGrpCount** | Pointer to **int64** | Number of logging groups configured on APIC. | [optional] 
**ApicSysLogSrcCount** | Pointer to **int64** | Number of logging sources configured on APIC. | [optional] 
**AppCenterCount** | Pointer to **int64** | ACI APPs feature usage scale. | [optional] 
**Ave** | Pointer to **string** | AVE feature usage. This determines if ACI virtual edge feature is enabled or disabled. | [optional] 
**BdCount** | Pointer to **int64** | Number of BDs. This determines the total number of Broadcast Domains across the fabric. | [optional] 
**CallhomeSmartGroupCount** | Pointer to **int64** | Number of call home smart monitoring policies on the fabric. | [optional] 
**CloudSecPeerCount** | Pointer to **int64** | Number of Cloudsec SA peers. | [optional] 
**CompHvCount** | Pointer to **int64** | Number of compute hypervisors on the fabric. | [optional] 
**ConfigExportpCount** | Pointer to **int64** | Number of system backup configure export policies on the fabric. | [optional] 
**ConfigJobCount** | Pointer to **int64** | Number of system backup configure jobs on the fabric. | [optional] 
**ConsistencyCheckerApp** | Pointer to **string** | Consistency checker application usage. This determines if the fabric has Consistency checker application installed. | [optional] 
**ContractCount** | Pointer to **int64** | Number of contracts. This determines the total number of Contracts configured across the fabric. | [optional] 
**DnsCount** | Pointer to **int64** | DNS feature usage. This determines the total number of DNS configurations across the fabric. | [optional] 
**EigrpCount** | Pointer to **int64** | Eigrp feature usage. This determines the total number of EIGRP sessions across the fabric. | [optional] 
**EpgCount** | Pointer to **int64** | Number of End Point Groups. This determines the total number of End Point Groups across the fabric. | [optional] 
**FabricModuleCount** | Pointer to **int64** | Returns the total number of fabric module slots. | [optional] 
**FabricSetuppCount** | Pointer to **int64** | Number of Multi-Pods per fabric. | [optional] 
**FcoeNportCount** | Pointer to **int64** | Total number of FCoE N-Port for DOM, VSAn, and VLAN. | [optional] 
**FcoeNportDomCount** | Pointer to **int64** | Number of FCoE N-Port DOM. | [optional] 
**FcoeNportVlanCount** | Pointer to **int64** | Number of FCoE N-Port VLAN. | [optional] 
**FcoeNportVsanCount** | Pointer to **int64** | Number of FCoE N-Port VSAN. | [optional] 
**FvSlaDefCount** | Pointer to **int64** | Number of Internet Protocol Service Level Agreements Monitoring policy objects for object tracking. | [optional] 
**HsrpCount** | Pointer to **int64** | Hsrp feature usage. This determines the total number of HSRP sessions across the fabric. | [optional] 
**IbgpCount** | Pointer to **int64** | Ibgp feature usage. This determines the total number of BGP sessions across the fabric. | [optional] 
**IgmpAccessListCount** | Pointer to **int64** | IGMP Access List feature usage. This determines the total number of IGMP access lists configured across the fabric. | [optional] 
**IgmpSnoop** | Pointer to **string** | IGMP Snooping feature usage. This determines if this feature is enabled or disabled. | [optional] 
**IpEpgCount** | Pointer to **int64** | Number of IP based End Point Groups. This determines the total number of IP End Point Groups across the fabric. | [optional] 
**IsBgpRouteReflectorsFeatureUsed** | Pointer to **bool** | BGP route reflector usage on APIC. | [optional] 
**IsBridgeDomainsFeatureUsed** | Pointer to **bool** | Brodge domains feature usage on APIC controller. | [optional] 
**IsCommonLocalUserName** | Pointer to **bool** | Returns value of isCommonLocalUserName field. | [optional] 
**IsContractsFeatureUsed** | Pointer to **bool** | Contracts feature usage on APIC controller. | [optional] 
**IsEpgFeatureUsed** | Pointer to **bool** | EPG feature usage on APIC controller. | [optional] 
**IsFiltersFeatureUsed** | Pointer to **bool** | Filters feature usage on APIC. | [optional] 
**IsHttpConfigured** | Pointer to **bool** | Returns if HTTP is configured. | [optional] 
**IsHttpsConfigured** | Pointer to **bool** | Returns if HTTPS is configured. | [optional] 
**IsNtpFeatureUsed** | Pointer to **bool** | NTP feature usage on APIC controller. | [optional] 
**IsPtpFeatureUsed** | Pointer to **bool** | Ptp feature usage on APIC. | [optional] 
**IsSynceFeatureUsed** | Pointer to **bool** | Synce feature usage on APIC. | [optional] 
**IsTechSupportCollected** | Pointer to **string** | Status of techsupport collection. | [optional] 
**IsTenantsFeatureUsed** | Pointer to **bool** | Tenants feature usage on APIC. | [optional] 
**IsVrfsFeatureUsed** | Pointer to **bool** | VRF feature usage on APIC controller. | [optional] 
**IsisCount** | Pointer to **int64** | Isis feature usage. This determines the total number of ISIS sessions across the fabric. | [optional] 
**L2Multicast** | Pointer to **string** | L2Multicast feature usage. This determines if this Layer 2 Multicast feature is being enabled / disabled on the fabric. | [optional] 
**LeafCount** | Pointer to **int64** | Number of Leafs. This determines the total number of Leaf switches in the fabric. | [optional] 
**LocalUsernameCount** | Pointer to **int64** | Returns count of local users. | [optional] 
**LoginBlockDuration** | Pointer to **int64** | Returns login block duration value. | [optional] 
**LoginMaxFailedAttempts** | Pointer to **int64** | Returns the maximum failed attempts on login. | [optional] 
**LoginMaxFailedAttemptsWindow** | Pointer to **int64** | Returns the maximum failed attempt windows on login. | [optional] 
**MaintenanceModeCount** | Pointer to **int64** | Maintenance Mode feature usage. This determines the number of switches that are currently in maintenance mode. | [optional] 
**ManagementOverV6Count** | Pointer to **int64** | Management over IPv6 feature usage. This determines the total number of IPv6 configurtaions in the fabric. | [optional] 
**MicrosoftUsegVmmEpPdCount** | Pointer to **int64** | Number of Microsoft microsegmentation VmmEpPD objects. Ensures that Microsoft was configured. | [optional] 
**NetFlowCount** | Pointer to **int64** | Number of Netflow monitor policies. | [optional] 
**Nir** | Pointer to **string** | NIR application usage. This determines if the fabric has NIR application installed. | [optional] 
**OpenStack** | Pointer to **string** | Open stack feature usage. | [optional] 
**OpflexKubernetesCount** | Pointer to **int64** | Opflex for Kubernetes feature usage. This determines the total number of VMM sessions of type kubernetes. | [optional] 
**OspfCount** | Pointer to **int64** | Ospf feature usage. This determines the total number of OSPF sessions across the fabric. | [optional] 
**PasswordHistoryCount** | Pointer to **int64** | Returns count of passwords. | [optional] 
**PasswordStrengthCheck** | Pointer to **string** | Returns if the password is strong or not. | [optional] 
**PasswordStrengthProfileCount** | Pointer to **int64** | Returns the number of password strength profile. | [optional] 
**PoeCount** | Pointer to **int64** | POE feature usage. This determines the total number of POE configurations across the fabric. | [optional] 
**PortSecurityCount** | Pointer to **int64** | Number of objects with Port Security enabled. Non-Zero value indicates the object as enabled. | [optional] 
**QinVniTunnelCount** | Pointer to **int64** | QinVniTunnel feature usage. This determines if the qinVniTunnel feature is being used on the fabric and the scale of it. | [optional] 
**QosCongCount** | Pointer to **int64** | Number of Quality Of Service congestion class. | [optional] 
**QosPfcPolCount** | Pointer to **int64** | Number of Quality Of Service class. | [optional] 
**RealmCount** | Pointer to **int64** | Returns the value of count of realms. | [optional] 
**RecordType** | Pointer to **string** | Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. | [optional] 
**RecordVersion** | Pointer to **string** | Version of record being pushed. This determines what was the API version for data available from the device. | [optional] 
**RemoteLeafCount** | Pointer to **int64** | Number of remote Leafs. This determines the total number of remote leaf switches in the fabric. | [optional] 
**ScvmmCount** | Pointer to **int64** | SCVMM feature usage. This determines the total number of SCVMM configurations in the fabric. | [optional] 
**SharedL3OutCount** | Pointer to **int64** | SharedL3Out feature usage. This determines the total number of Shared L3 out configured across the fabric. | [optional] 
**SiteName** | Pointer to **string** | The Site name represents an APIC cluster. Service Engine can onboard multiple APIC clusters / sites. There will be a feature usage object per site in Multi site scenario. In multi-site scenario the site name is available in all the requests being made. | [optional] 
**SmartCallHome** | Pointer to **string** | Smart callhome feature usage. This determines if this feature is being enabled or disabled. | [optional] 
**SmartLicense** | Pointer to [**NullableNiatelemetrySmartLicense**](NiatelemetrySmartLicense.md) |  | [optional] 
**SnapshotCount** | Pointer to **int64** | Returns count of snapshots. | [optional] 
**Snmp** | Pointer to **string** | SNMP feature usage. This determines if this feature is enabled or disabled. | [optional] 
**SnmpCommunityAccessCount** | Pointer to **int64** | Returns count of SNMP Community Access. | [optional] 
**SnmpGroupCount** | Pointer to **int64** | Number of SNMP monitoring policies on the fabric. | [optional] 
**SnmpTrapCount** | Pointer to **int64** | Returns count of SNMP trap. | [optional] 
**SnmpV3Count** | Pointer to **int64** | Returns count of SNMP V3 on the fabric. | [optional] 
**SpanCount** | Pointer to **int64** | Number of Span Sources and Destinations. | [optional] 
**SpanDstCount** | Pointer to **int64** | Number of Span Destinations with valid state. | [optional] 
**SpanSrcCount** | Pointer to **int64** | Number of Span Sources with valid state. | [optional] 
**SpineCount** | Pointer to **int64** | Number of Spines. This determines the total number of spine switches in the fabric. | [optional] 
**SshOverV6Count** | Pointer to **int64** | Ssh over IPv6 feature usage. This determines the total number of IPv6 configurtaions in the fabric. | [optional] 
**SshV2Count** | Pointer to **int64** | Returns count of ssh V2 on the fabric. | [optional] 
**SupervisorModuleCount** | Pointer to **int64** | Returns the total number of supervisor module slots. | [optional] 
**SyslogGroupCount** | Pointer to **int64** | Number of syslog monitoring policies on the fabric. | [optional] 
**SyslogOverV6Count** | Pointer to **int64** | Syslog over IPv6 feature usage. This determines the total number of IPv6 configurtaions in the fabric. | [optional] 
**SystemControllerCount** | Pointer to **int64** | Returns the total number of system controller slots. | [optional] 
**TacacsGroupCount** | Pointer to **int64** | Number of tacacs monitoring policies on the fabric. | [optional] 
**TenantCount** | Pointer to **int64** | Number of tenants. This determines the total number of tenants configured across the fabric. | [optional] 
**TierTwoLeafCount** | Pointer to **int64** | Number of tier 2 Leafs. This determines the total number of tier 2 Leaf switches in the fabric. | [optional] 
**TotalCriticalFaults** | Pointer to **int64** | Returns the total number of critical faults. | [optional] 
**Twamp** | Pointer to **string** | TWAMP feature usage. This determines if this feature is enabled or disabled. | [optional] 
**Useg** | Pointer to **string** | VMM uSegmentation feature usage. This determines if microsegmentation feature is enabled or disabled. | [optional] 
**VmWareVdsCount** | Pointer to **int64** | Number of objects with VmWare vCenter 6.5 support. Checks the controller revision value. | [optional] 
**VmmCtrlrpCount** | Pointer to **int64** | Number of Virtual Machine Monitor controller policy objects for VMware vCenter. | [optional] 
**VmmDompCount** | Pointer to **int64** | Number of Virtual Machine Monitor domain policy model objects for VMware vCenter. | [optional] 
**VmmEpPdCount** | Pointer to **int64** | Microsegmentation Distributed Virtual Switch feature usage. Gets the number of objects associated to VMware vCenter. | [optional] 
**VnsmDevCount** | Pointer to **int64** | Number of objects with L4-L7 Device Package Import enabled. Checks for the vendor and the model. | [optional] 
**VpodCount** | Pointer to **int64** | Virtual pod feature usage. This determines the total number of virtual POD configurations in the fabrics. | [optional] 
**WebtokenTimeoutSeconds** | Pointer to **int64** | Timeout for web token in seconds. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetryNiaFeatureUsageAllOf

`func NewNiatelemetryNiaFeatureUsageAllOf(classId string, objectType string, ) *NiatelemetryNiaFeatureUsageAllOf`

NewNiatelemetryNiaFeatureUsageAllOf instantiates a new NiatelemetryNiaFeatureUsageAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryNiaFeatureUsageAllOfWithDefaults

`func NewNiatelemetryNiaFeatureUsageAllOfWithDefaults() *NiatelemetryNiaFeatureUsageAllOf`

NewNiatelemetryNiaFeatureUsageAllOfWithDefaults instantiates a new NiatelemetryNiaFeatureUsageAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAaaLdapProviderCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetAaaLdapProviderCount() int64`

GetAaaLdapProviderCount returns the AaaLdapProviderCount field if non-nil, zero value otherwise.

### GetAaaLdapProviderCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetAaaLdapProviderCountOk() (*int64, bool)`

GetAaaLdapProviderCountOk returns a tuple with the AaaLdapProviderCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAaaLdapProviderCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetAaaLdapProviderCount(v int64)`

SetAaaLdapProviderCount sets AaaLdapProviderCount field to given value.

### HasAaaLdapProviderCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasAaaLdapProviderCount() bool`

HasAaaLdapProviderCount returns a boolean if a field has been set.

### GetAaaRadiusProviderCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetAaaRadiusProviderCount() int64`

GetAaaRadiusProviderCount returns the AaaRadiusProviderCount field if non-nil, zero value otherwise.

### GetAaaRadiusProviderCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetAaaRadiusProviderCountOk() (*int64, bool)`

GetAaaRadiusProviderCountOk returns a tuple with the AaaRadiusProviderCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAaaRadiusProviderCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetAaaRadiusProviderCount(v int64)`

SetAaaRadiusProviderCount sets AaaRadiusProviderCount field to given value.

### HasAaaRadiusProviderCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasAaaRadiusProviderCount() bool`

HasAaaRadiusProviderCount returns a boolean if a field has been set.

### GetAaaTacacsProviderCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetAaaTacacsProviderCount() int64`

GetAaaTacacsProviderCount returns the AaaTacacsProviderCount field if non-nil, zero value otherwise.

### GetAaaTacacsProviderCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetAaaTacacsProviderCountOk() (*int64, bool)`

GetAaaTacacsProviderCountOk returns a tuple with the AaaTacacsProviderCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAaaTacacsProviderCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetAaaTacacsProviderCount(v int64)`

SetAaaTacacsProviderCount sets AaaTacacsProviderCount field to given value.

### HasAaaTacacsProviderCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasAaaTacacsProviderCount() bool`

HasAaaTacacsProviderCount returns a boolean if a field has been set.

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

### GetApicIsTelnetEnabled

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetApicIsTelnetEnabled() bool`

GetApicIsTelnetEnabled returns the ApicIsTelnetEnabled field if non-nil, zero value otherwise.

### GetApicIsTelnetEnabledOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetApicIsTelnetEnabledOk() (*bool, bool)`

GetApicIsTelnetEnabledOk returns a tuple with the ApicIsTelnetEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApicIsTelnetEnabled

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetApicIsTelnetEnabled(v bool)`

SetApicIsTelnetEnabled sets ApicIsTelnetEnabled field to given value.

### HasApicIsTelnetEnabled

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasApicIsTelnetEnabled() bool`

HasApicIsTelnetEnabled returns a boolean if a field has been set.

### GetApicNtpCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetApicNtpCount() int64`

GetApicNtpCount returns the ApicNtpCount field if non-nil, zero value otherwise.

### GetApicNtpCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetApicNtpCountOk() (*int64, bool)`

GetApicNtpCountOk returns a tuple with the ApicNtpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApicNtpCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetApicNtpCount(v int64)`

SetApicNtpCount sets ApicNtpCount field to given value.

### HasApicNtpCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasApicNtpCount() bool`

HasApicNtpCount returns a boolean if a field has been set.

### GetApicSnmpCommunityCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetApicSnmpCommunityCount() int64`

GetApicSnmpCommunityCount returns the ApicSnmpCommunityCount field if non-nil, zero value otherwise.

### GetApicSnmpCommunityCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetApicSnmpCommunityCountOk() (*int64, bool)`

GetApicSnmpCommunityCountOk returns a tuple with the ApicSnmpCommunityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApicSnmpCommunityCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetApicSnmpCommunityCount(v int64)`

SetApicSnmpCommunityCount sets ApicSnmpCommunityCount field to given value.

### HasApicSnmpCommunityCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasApicSnmpCommunityCount() bool`

HasApicSnmpCommunityCount returns a boolean if a field has been set.

### GetApicSysLogGrpCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetApicSysLogGrpCount() int64`

GetApicSysLogGrpCount returns the ApicSysLogGrpCount field if non-nil, zero value otherwise.

### GetApicSysLogGrpCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetApicSysLogGrpCountOk() (*int64, bool)`

GetApicSysLogGrpCountOk returns a tuple with the ApicSysLogGrpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApicSysLogGrpCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetApicSysLogGrpCount(v int64)`

SetApicSysLogGrpCount sets ApicSysLogGrpCount field to given value.

### HasApicSysLogGrpCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasApicSysLogGrpCount() bool`

HasApicSysLogGrpCount returns a boolean if a field has been set.

### GetApicSysLogSrcCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetApicSysLogSrcCount() int64`

GetApicSysLogSrcCount returns the ApicSysLogSrcCount field if non-nil, zero value otherwise.

### GetApicSysLogSrcCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetApicSysLogSrcCountOk() (*int64, bool)`

GetApicSysLogSrcCountOk returns a tuple with the ApicSysLogSrcCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApicSysLogSrcCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetApicSysLogSrcCount(v int64)`

SetApicSysLogSrcCount sets ApicSysLogSrcCount field to given value.

### HasApicSysLogSrcCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasApicSysLogSrcCount() bool`

HasApicSysLogSrcCount returns a boolean if a field has been set.

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

### GetCallhomeSmartGroupCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetCallhomeSmartGroupCount() int64`

GetCallhomeSmartGroupCount returns the CallhomeSmartGroupCount field if non-nil, zero value otherwise.

### GetCallhomeSmartGroupCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetCallhomeSmartGroupCountOk() (*int64, bool)`

GetCallhomeSmartGroupCountOk returns a tuple with the CallhomeSmartGroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallhomeSmartGroupCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetCallhomeSmartGroupCount(v int64)`

SetCallhomeSmartGroupCount sets CallhomeSmartGroupCount field to given value.

### HasCallhomeSmartGroupCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasCallhomeSmartGroupCount() bool`

HasCallhomeSmartGroupCount returns a boolean if a field has been set.

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

### GetCompHvCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetCompHvCount() int64`

GetCompHvCount returns the CompHvCount field if non-nil, zero value otherwise.

### GetCompHvCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetCompHvCountOk() (*int64, bool)`

GetCompHvCountOk returns a tuple with the CompHvCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompHvCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetCompHvCount(v int64)`

SetCompHvCount sets CompHvCount field to given value.

### HasCompHvCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasCompHvCount() bool`

HasCompHvCount returns a boolean if a field has been set.

### GetConfigExportpCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetConfigExportpCount() int64`

GetConfigExportpCount returns the ConfigExportpCount field if non-nil, zero value otherwise.

### GetConfigExportpCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetConfigExportpCountOk() (*int64, bool)`

GetConfigExportpCountOk returns a tuple with the ConfigExportpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigExportpCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetConfigExportpCount(v int64)`

SetConfigExportpCount sets ConfigExportpCount field to given value.

### HasConfigExportpCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasConfigExportpCount() bool`

HasConfigExportpCount returns a boolean if a field has been set.

### GetConfigJobCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetConfigJobCount() int64`

GetConfigJobCount returns the ConfigJobCount field if non-nil, zero value otherwise.

### GetConfigJobCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetConfigJobCountOk() (*int64, bool)`

GetConfigJobCountOk returns a tuple with the ConfigJobCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigJobCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetConfigJobCount(v int64)`

SetConfigJobCount sets ConfigJobCount field to given value.

### HasConfigJobCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasConfigJobCount() bool`

HasConfigJobCount returns a boolean if a field has been set.

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

### GetFabricModuleCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetFabricModuleCount() int64`

GetFabricModuleCount returns the FabricModuleCount field if non-nil, zero value otherwise.

### GetFabricModuleCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetFabricModuleCountOk() (*int64, bool)`

GetFabricModuleCountOk returns a tuple with the FabricModuleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricModuleCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetFabricModuleCount(v int64)`

SetFabricModuleCount sets FabricModuleCount field to given value.

### HasFabricModuleCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasFabricModuleCount() bool`

HasFabricModuleCount returns a boolean if a field has been set.

### GetFabricSetuppCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetFabricSetuppCount() int64`

GetFabricSetuppCount returns the FabricSetuppCount field if non-nil, zero value otherwise.

### GetFabricSetuppCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetFabricSetuppCountOk() (*int64, bool)`

GetFabricSetuppCountOk returns a tuple with the FabricSetuppCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricSetuppCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetFabricSetuppCount(v int64)`

SetFabricSetuppCount sets FabricSetuppCount field to given value.

### HasFabricSetuppCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasFabricSetuppCount() bool`

HasFabricSetuppCount returns a boolean if a field has been set.

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

### GetFvSlaDefCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetFvSlaDefCount() int64`

GetFvSlaDefCount returns the FvSlaDefCount field if non-nil, zero value otherwise.

### GetFvSlaDefCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetFvSlaDefCountOk() (*int64, bool)`

GetFvSlaDefCountOk returns a tuple with the FvSlaDefCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFvSlaDefCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetFvSlaDefCount(v int64)`

SetFvSlaDefCount sets FvSlaDefCount field to given value.

### HasFvSlaDefCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasFvSlaDefCount() bool`

HasFvSlaDefCount returns a boolean if a field has been set.

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

### GetIsBgpRouteReflectorsFeatureUsed

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetIsBgpRouteReflectorsFeatureUsed() bool`

GetIsBgpRouteReflectorsFeatureUsed returns the IsBgpRouteReflectorsFeatureUsed field if non-nil, zero value otherwise.

### GetIsBgpRouteReflectorsFeatureUsedOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetIsBgpRouteReflectorsFeatureUsedOk() (*bool, bool)`

GetIsBgpRouteReflectorsFeatureUsedOk returns a tuple with the IsBgpRouteReflectorsFeatureUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBgpRouteReflectorsFeatureUsed

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetIsBgpRouteReflectorsFeatureUsed(v bool)`

SetIsBgpRouteReflectorsFeatureUsed sets IsBgpRouteReflectorsFeatureUsed field to given value.

### HasIsBgpRouteReflectorsFeatureUsed

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasIsBgpRouteReflectorsFeatureUsed() bool`

HasIsBgpRouteReflectorsFeatureUsed returns a boolean if a field has been set.

### GetIsBridgeDomainsFeatureUsed

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetIsBridgeDomainsFeatureUsed() bool`

GetIsBridgeDomainsFeatureUsed returns the IsBridgeDomainsFeatureUsed field if non-nil, zero value otherwise.

### GetIsBridgeDomainsFeatureUsedOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetIsBridgeDomainsFeatureUsedOk() (*bool, bool)`

GetIsBridgeDomainsFeatureUsedOk returns a tuple with the IsBridgeDomainsFeatureUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBridgeDomainsFeatureUsed

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetIsBridgeDomainsFeatureUsed(v bool)`

SetIsBridgeDomainsFeatureUsed sets IsBridgeDomainsFeatureUsed field to given value.

### HasIsBridgeDomainsFeatureUsed

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasIsBridgeDomainsFeatureUsed() bool`

HasIsBridgeDomainsFeatureUsed returns a boolean if a field has been set.

### GetIsCommonLocalUserName

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetIsCommonLocalUserName() bool`

GetIsCommonLocalUserName returns the IsCommonLocalUserName field if non-nil, zero value otherwise.

### GetIsCommonLocalUserNameOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetIsCommonLocalUserNameOk() (*bool, bool)`

GetIsCommonLocalUserNameOk returns a tuple with the IsCommonLocalUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCommonLocalUserName

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetIsCommonLocalUserName(v bool)`

SetIsCommonLocalUserName sets IsCommonLocalUserName field to given value.

### HasIsCommonLocalUserName

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasIsCommonLocalUserName() bool`

HasIsCommonLocalUserName returns a boolean if a field has been set.

### GetIsContractsFeatureUsed

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetIsContractsFeatureUsed() bool`

GetIsContractsFeatureUsed returns the IsContractsFeatureUsed field if non-nil, zero value otherwise.

### GetIsContractsFeatureUsedOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetIsContractsFeatureUsedOk() (*bool, bool)`

GetIsContractsFeatureUsedOk returns a tuple with the IsContractsFeatureUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsContractsFeatureUsed

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetIsContractsFeatureUsed(v bool)`

SetIsContractsFeatureUsed sets IsContractsFeatureUsed field to given value.

### HasIsContractsFeatureUsed

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasIsContractsFeatureUsed() bool`

HasIsContractsFeatureUsed returns a boolean if a field has been set.

### GetIsEpgFeatureUsed

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetIsEpgFeatureUsed() bool`

GetIsEpgFeatureUsed returns the IsEpgFeatureUsed field if non-nil, zero value otherwise.

### GetIsEpgFeatureUsedOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetIsEpgFeatureUsedOk() (*bool, bool)`

GetIsEpgFeatureUsedOk returns a tuple with the IsEpgFeatureUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEpgFeatureUsed

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetIsEpgFeatureUsed(v bool)`

SetIsEpgFeatureUsed sets IsEpgFeatureUsed field to given value.

### HasIsEpgFeatureUsed

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasIsEpgFeatureUsed() bool`

HasIsEpgFeatureUsed returns a boolean if a field has been set.

### GetIsFiltersFeatureUsed

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetIsFiltersFeatureUsed() bool`

GetIsFiltersFeatureUsed returns the IsFiltersFeatureUsed field if non-nil, zero value otherwise.

### GetIsFiltersFeatureUsedOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetIsFiltersFeatureUsedOk() (*bool, bool)`

GetIsFiltersFeatureUsedOk returns a tuple with the IsFiltersFeatureUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFiltersFeatureUsed

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetIsFiltersFeatureUsed(v bool)`

SetIsFiltersFeatureUsed sets IsFiltersFeatureUsed field to given value.

### HasIsFiltersFeatureUsed

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasIsFiltersFeatureUsed() bool`

HasIsFiltersFeatureUsed returns a boolean if a field has been set.

### GetIsHttpConfigured

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetIsHttpConfigured() bool`

GetIsHttpConfigured returns the IsHttpConfigured field if non-nil, zero value otherwise.

### GetIsHttpConfiguredOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetIsHttpConfiguredOk() (*bool, bool)`

GetIsHttpConfiguredOk returns a tuple with the IsHttpConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHttpConfigured

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetIsHttpConfigured(v bool)`

SetIsHttpConfigured sets IsHttpConfigured field to given value.

### HasIsHttpConfigured

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasIsHttpConfigured() bool`

HasIsHttpConfigured returns a boolean if a field has been set.

### GetIsHttpsConfigured

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetIsHttpsConfigured() bool`

GetIsHttpsConfigured returns the IsHttpsConfigured field if non-nil, zero value otherwise.

### GetIsHttpsConfiguredOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetIsHttpsConfiguredOk() (*bool, bool)`

GetIsHttpsConfiguredOk returns a tuple with the IsHttpsConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHttpsConfigured

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetIsHttpsConfigured(v bool)`

SetIsHttpsConfigured sets IsHttpsConfigured field to given value.

### HasIsHttpsConfigured

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasIsHttpsConfigured() bool`

HasIsHttpsConfigured returns a boolean if a field has been set.

### GetIsNtpFeatureUsed

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetIsNtpFeatureUsed() bool`

GetIsNtpFeatureUsed returns the IsNtpFeatureUsed field if non-nil, zero value otherwise.

### GetIsNtpFeatureUsedOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetIsNtpFeatureUsedOk() (*bool, bool)`

GetIsNtpFeatureUsedOk returns a tuple with the IsNtpFeatureUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNtpFeatureUsed

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetIsNtpFeatureUsed(v bool)`

SetIsNtpFeatureUsed sets IsNtpFeatureUsed field to given value.

### HasIsNtpFeatureUsed

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasIsNtpFeatureUsed() bool`

HasIsNtpFeatureUsed returns a boolean if a field has been set.

### GetIsPtpFeatureUsed

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetIsPtpFeatureUsed() bool`

GetIsPtpFeatureUsed returns the IsPtpFeatureUsed field if non-nil, zero value otherwise.

### GetIsPtpFeatureUsedOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetIsPtpFeatureUsedOk() (*bool, bool)`

GetIsPtpFeatureUsedOk returns a tuple with the IsPtpFeatureUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPtpFeatureUsed

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetIsPtpFeatureUsed(v bool)`

SetIsPtpFeatureUsed sets IsPtpFeatureUsed field to given value.

### HasIsPtpFeatureUsed

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasIsPtpFeatureUsed() bool`

HasIsPtpFeatureUsed returns a boolean if a field has been set.

### GetIsSynceFeatureUsed

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetIsSynceFeatureUsed() bool`

GetIsSynceFeatureUsed returns the IsSynceFeatureUsed field if non-nil, zero value otherwise.

### GetIsSynceFeatureUsedOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetIsSynceFeatureUsedOk() (*bool, bool)`

GetIsSynceFeatureUsedOk returns a tuple with the IsSynceFeatureUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSynceFeatureUsed

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetIsSynceFeatureUsed(v bool)`

SetIsSynceFeatureUsed sets IsSynceFeatureUsed field to given value.

### HasIsSynceFeatureUsed

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasIsSynceFeatureUsed() bool`

HasIsSynceFeatureUsed returns a boolean if a field has been set.

### GetIsTechSupportCollected

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetIsTechSupportCollected() string`

GetIsTechSupportCollected returns the IsTechSupportCollected field if non-nil, zero value otherwise.

### GetIsTechSupportCollectedOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetIsTechSupportCollectedOk() (*string, bool)`

GetIsTechSupportCollectedOk returns a tuple with the IsTechSupportCollected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTechSupportCollected

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetIsTechSupportCollected(v string)`

SetIsTechSupportCollected sets IsTechSupportCollected field to given value.

### HasIsTechSupportCollected

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasIsTechSupportCollected() bool`

HasIsTechSupportCollected returns a boolean if a field has been set.

### GetIsTenantsFeatureUsed

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetIsTenantsFeatureUsed() bool`

GetIsTenantsFeatureUsed returns the IsTenantsFeatureUsed field if non-nil, zero value otherwise.

### GetIsTenantsFeatureUsedOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetIsTenantsFeatureUsedOk() (*bool, bool)`

GetIsTenantsFeatureUsedOk returns a tuple with the IsTenantsFeatureUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTenantsFeatureUsed

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetIsTenantsFeatureUsed(v bool)`

SetIsTenantsFeatureUsed sets IsTenantsFeatureUsed field to given value.

### HasIsTenantsFeatureUsed

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasIsTenantsFeatureUsed() bool`

HasIsTenantsFeatureUsed returns a boolean if a field has been set.

### GetIsVrfsFeatureUsed

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetIsVrfsFeatureUsed() bool`

GetIsVrfsFeatureUsed returns the IsVrfsFeatureUsed field if non-nil, zero value otherwise.

### GetIsVrfsFeatureUsedOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetIsVrfsFeatureUsedOk() (*bool, bool)`

GetIsVrfsFeatureUsedOk returns a tuple with the IsVrfsFeatureUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVrfsFeatureUsed

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetIsVrfsFeatureUsed(v bool)`

SetIsVrfsFeatureUsed sets IsVrfsFeatureUsed field to given value.

### HasIsVrfsFeatureUsed

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasIsVrfsFeatureUsed() bool`

HasIsVrfsFeatureUsed returns a boolean if a field has been set.

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

### GetLocalUsernameCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetLocalUsernameCount() int64`

GetLocalUsernameCount returns the LocalUsernameCount field if non-nil, zero value otherwise.

### GetLocalUsernameCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetLocalUsernameCountOk() (*int64, bool)`

GetLocalUsernameCountOk returns a tuple with the LocalUsernameCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalUsernameCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetLocalUsernameCount(v int64)`

SetLocalUsernameCount sets LocalUsernameCount field to given value.

### HasLocalUsernameCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasLocalUsernameCount() bool`

HasLocalUsernameCount returns a boolean if a field has been set.

### GetLoginBlockDuration

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetLoginBlockDuration() int64`

GetLoginBlockDuration returns the LoginBlockDuration field if non-nil, zero value otherwise.

### GetLoginBlockDurationOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetLoginBlockDurationOk() (*int64, bool)`

GetLoginBlockDurationOk returns a tuple with the LoginBlockDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginBlockDuration

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetLoginBlockDuration(v int64)`

SetLoginBlockDuration sets LoginBlockDuration field to given value.

### HasLoginBlockDuration

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasLoginBlockDuration() bool`

HasLoginBlockDuration returns a boolean if a field has been set.

### GetLoginMaxFailedAttempts

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetLoginMaxFailedAttempts() int64`

GetLoginMaxFailedAttempts returns the LoginMaxFailedAttempts field if non-nil, zero value otherwise.

### GetLoginMaxFailedAttemptsOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetLoginMaxFailedAttemptsOk() (*int64, bool)`

GetLoginMaxFailedAttemptsOk returns a tuple with the LoginMaxFailedAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginMaxFailedAttempts

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetLoginMaxFailedAttempts(v int64)`

SetLoginMaxFailedAttempts sets LoginMaxFailedAttempts field to given value.

### HasLoginMaxFailedAttempts

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasLoginMaxFailedAttempts() bool`

HasLoginMaxFailedAttempts returns a boolean if a field has been set.

### GetLoginMaxFailedAttemptsWindow

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetLoginMaxFailedAttemptsWindow() int64`

GetLoginMaxFailedAttemptsWindow returns the LoginMaxFailedAttemptsWindow field if non-nil, zero value otherwise.

### GetLoginMaxFailedAttemptsWindowOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetLoginMaxFailedAttemptsWindowOk() (*int64, bool)`

GetLoginMaxFailedAttemptsWindowOk returns a tuple with the LoginMaxFailedAttemptsWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginMaxFailedAttemptsWindow

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetLoginMaxFailedAttemptsWindow(v int64)`

SetLoginMaxFailedAttemptsWindow sets LoginMaxFailedAttemptsWindow field to given value.

### HasLoginMaxFailedAttemptsWindow

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasLoginMaxFailedAttemptsWindow() bool`

HasLoginMaxFailedAttemptsWindow returns a boolean if a field has been set.

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

### GetPasswordHistoryCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetPasswordHistoryCount() int64`

GetPasswordHistoryCount returns the PasswordHistoryCount field if non-nil, zero value otherwise.

### GetPasswordHistoryCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetPasswordHistoryCountOk() (*int64, bool)`

GetPasswordHistoryCountOk returns a tuple with the PasswordHistoryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHistoryCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetPasswordHistoryCount(v int64)`

SetPasswordHistoryCount sets PasswordHistoryCount field to given value.

### HasPasswordHistoryCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasPasswordHistoryCount() bool`

HasPasswordHistoryCount returns a boolean if a field has been set.

### GetPasswordStrengthCheck

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetPasswordStrengthCheck() string`

GetPasswordStrengthCheck returns the PasswordStrengthCheck field if non-nil, zero value otherwise.

### GetPasswordStrengthCheckOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetPasswordStrengthCheckOk() (*string, bool)`

GetPasswordStrengthCheckOk returns a tuple with the PasswordStrengthCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordStrengthCheck

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetPasswordStrengthCheck(v string)`

SetPasswordStrengthCheck sets PasswordStrengthCheck field to given value.

### HasPasswordStrengthCheck

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasPasswordStrengthCheck() bool`

HasPasswordStrengthCheck returns a boolean if a field has been set.

### GetPasswordStrengthProfileCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetPasswordStrengthProfileCount() int64`

GetPasswordStrengthProfileCount returns the PasswordStrengthProfileCount field if non-nil, zero value otherwise.

### GetPasswordStrengthProfileCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetPasswordStrengthProfileCountOk() (*int64, bool)`

GetPasswordStrengthProfileCountOk returns a tuple with the PasswordStrengthProfileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordStrengthProfileCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetPasswordStrengthProfileCount(v int64)`

SetPasswordStrengthProfileCount sets PasswordStrengthProfileCount field to given value.

### HasPasswordStrengthProfileCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasPasswordStrengthProfileCount() bool`

HasPasswordStrengthProfileCount returns a boolean if a field has been set.

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

### GetQosCongCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetQosCongCount() int64`

GetQosCongCount returns the QosCongCount field if non-nil, zero value otherwise.

### GetQosCongCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetQosCongCountOk() (*int64, bool)`

GetQosCongCountOk returns a tuple with the QosCongCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosCongCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetQosCongCount(v int64)`

SetQosCongCount sets QosCongCount field to given value.

### HasQosCongCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasQosCongCount() bool`

HasQosCongCount returns a boolean if a field has been set.

### GetQosPfcPolCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetQosPfcPolCount() int64`

GetQosPfcPolCount returns the QosPfcPolCount field if non-nil, zero value otherwise.

### GetQosPfcPolCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetQosPfcPolCountOk() (*int64, bool)`

GetQosPfcPolCountOk returns a tuple with the QosPfcPolCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosPfcPolCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetQosPfcPolCount(v int64)`

SetQosPfcPolCount sets QosPfcPolCount field to given value.

### HasQosPfcPolCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasQosPfcPolCount() bool`

HasQosPfcPolCount returns a boolean if a field has been set.

### GetRealmCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetRealmCount() int64`

GetRealmCount returns the RealmCount field if non-nil, zero value otherwise.

### GetRealmCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetRealmCountOk() (*int64, bool)`

GetRealmCountOk returns a tuple with the RealmCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetRealmCount(v int64)`

SetRealmCount sets RealmCount field to given value.

### HasRealmCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasRealmCount() bool`

HasRealmCount returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

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

### GetSmartLicense

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetSmartLicense() NiatelemetrySmartLicense`

GetSmartLicense returns the SmartLicense field if non-nil, zero value otherwise.

### GetSmartLicenseOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetSmartLicenseOk() (*NiatelemetrySmartLicense, bool)`

GetSmartLicenseOk returns a tuple with the SmartLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartLicense

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetSmartLicense(v NiatelemetrySmartLicense)`

SetSmartLicense sets SmartLicense field to given value.

### HasSmartLicense

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasSmartLicense() bool`

HasSmartLicense returns a boolean if a field has been set.

### SetSmartLicenseNil

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetSmartLicenseNil(b bool)`

 SetSmartLicenseNil sets the value for SmartLicense to be an explicit nil

### UnsetSmartLicense
`func (o *NiatelemetryNiaFeatureUsageAllOf) UnsetSmartLicense()`

UnsetSmartLicense ensures that no value is present for SmartLicense, not even an explicit nil
### GetSnapshotCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetSnapshotCount() int64`

GetSnapshotCount returns the SnapshotCount field if non-nil, zero value otherwise.

### GetSnapshotCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetSnapshotCountOk() (*int64, bool)`

GetSnapshotCountOk returns a tuple with the SnapshotCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetSnapshotCount(v int64)`

SetSnapshotCount sets SnapshotCount field to given value.

### HasSnapshotCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasSnapshotCount() bool`

HasSnapshotCount returns a boolean if a field has been set.

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

### GetSnmpCommunityAccessCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetSnmpCommunityAccessCount() int64`

GetSnmpCommunityAccessCount returns the SnmpCommunityAccessCount field if non-nil, zero value otherwise.

### GetSnmpCommunityAccessCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetSnmpCommunityAccessCountOk() (*int64, bool)`

GetSnmpCommunityAccessCountOk returns a tuple with the SnmpCommunityAccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpCommunityAccessCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetSnmpCommunityAccessCount(v int64)`

SetSnmpCommunityAccessCount sets SnmpCommunityAccessCount field to given value.

### HasSnmpCommunityAccessCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasSnmpCommunityAccessCount() bool`

HasSnmpCommunityAccessCount returns a boolean if a field has been set.

### GetSnmpGroupCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetSnmpGroupCount() int64`

GetSnmpGroupCount returns the SnmpGroupCount field if non-nil, zero value otherwise.

### GetSnmpGroupCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetSnmpGroupCountOk() (*int64, bool)`

GetSnmpGroupCountOk returns a tuple with the SnmpGroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpGroupCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetSnmpGroupCount(v int64)`

SetSnmpGroupCount sets SnmpGroupCount field to given value.

### HasSnmpGroupCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasSnmpGroupCount() bool`

HasSnmpGroupCount returns a boolean if a field has been set.

### GetSnmpTrapCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetSnmpTrapCount() int64`

GetSnmpTrapCount returns the SnmpTrapCount field if non-nil, zero value otherwise.

### GetSnmpTrapCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetSnmpTrapCountOk() (*int64, bool)`

GetSnmpTrapCountOk returns a tuple with the SnmpTrapCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpTrapCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetSnmpTrapCount(v int64)`

SetSnmpTrapCount sets SnmpTrapCount field to given value.

### HasSnmpTrapCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasSnmpTrapCount() bool`

HasSnmpTrapCount returns a boolean if a field has been set.

### GetSnmpV3Count

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetSnmpV3Count() int64`

GetSnmpV3Count returns the SnmpV3Count field if non-nil, zero value otherwise.

### GetSnmpV3CountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetSnmpV3CountOk() (*int64, bool)`

GetSnmpV3CountOk returns a tuple with the SnmpV3Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpV3Count

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetSnmpV3Count(v int64)`

SetSnmpV3Count sets SnmpV3Count field to given value.

### HasSnmpV3Count

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasSnmpV3Count() bool`

HasSnmpV3Count returns a boolean if a field has been set.

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

### GetSshV2Count

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetSshV2Count() int64`

GetSshV2Count returns the SshV2Count field if non-nil, zero value otherwise.

### GetSshV2CountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetSshV2CountOk() (*int64, bool)`

GetSshV2CountOk returns a tuple with the SshV2Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshV2Count

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetSshV2Count(v int64)`

SetSshV2Count sets SshV2Count field to given value.

### HasSshV2Count

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasSshV2Count() bool`

HasSshV2Count returns a boolean if a field has been set.

### GetSupervisorModuleCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetSupervisorModuleCount() int64`

GetSupervisorModuleCount returns the SupervisorModuleCount field if non-nil, zero value otherwise.

### GetSupervisorModuleCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetSupervisorModuleCountOk() (*int64, bool)`

GetSupervisorModuleCountOk returns a tuple with the SupervisorModuleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupervisorModuleCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetSupervisorModuleCount(v int64)`

SetSupervisorModuleCount sets SupervisorModuleCount field to given value.

### HasSupervisorModuleCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasSupervisorModuleCount() bool`

HasSupervisorModuleCount returns a boolean if a field has been set.

### GetSyslogGroupCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetSyslogGroupCount() int64`

GetSyslogGroupCount returns the SyslogGroupCount field if non-nil, zero value otherwise.

### GetSyslogGroupCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetSyslogGroupCountOk() (*int64, bool)`

GetSyslogGroupCountOk returns a tuple with the SyslogGroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogGroupCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetSyslogGroupCount(v int64)`

SetSyslogGroupCount sets SyslogGroupCount field to given value.

### HasSyslogGroupCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasSyslogGroupCount() bool`

HasSyslogGroupCount returns a boolean if a field has been set.

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

### GetSystemControllerCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetSystemControllerCount() int64`

GetSystemControllerCount returns the SystemControllerCount field if non-nil, zero value otherwise.

### GetSystemControllerCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetSystemControllerCountOk() (*int64, bool)`

GetSystemControllerCountOk returns a tuple with the SystemControllerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemControllerCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetSystemControllerCount(v int64)`

SetSystemControllerCount sets SystemControllerCount field to given value.

### HasSystemControllerCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasSystemControllerCount() bool`

HasSystemControllerCount returns a boolean if a field has been set.

### GetTacacsGroupCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetTacacsGroupCount() int64`

GetTacacsGroupCount returns the TacacsGroupCount field if non-nil, zero value otherwise.

### GetTacacsGroupCountOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetTacacsGroupCountOk() (*int64, bool)`

GetTacacsGroupCountOk returns a tuple with the TacacsGroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTacacsGroupCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetTacacsGroupCount(v int64)`

SetTacacsGroupCount sets TacacsGroupCount field to given value.

### HasTacacsGroupCount

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasTacacsGroupCount() bool`

HasTacacsGroupCount returns a boolean if a field has been set.

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

### GetTotalCriticalFaults

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetTotalCriticalFaults() int64`

GetTotalCriticalFaults returns the TotalCriticalFaults field if non-nil, zero value otherwise.

### GetTotalCriticalFaultsOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetTotalCriticalFaultsOk() (*int64, bool)`

GetTotalCriticalFaultsOk returns a tuple with the TotalCriticalFaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCriticalFaults

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetTotalCriticalFaults(v int64)`

SetTotalCriticalFaults sets TotalCriticalFaults field to given value.

### HasTotalCriticalFaults

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasTotalCriticalFaults() bool`

HasTotalCriticalFaults returns a boolean if a field has been set.

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

### GetWebtokenTimeoutSeconds

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetWebtokenTimeoutSeconds() int64`

GetWebtokenTimeoutSeconds returns the WebtokenTimeoutSeconds field if non-nil, zero value otherwise.

### GetWebtokenTimeoutSecondsOk

`func (o *NiatelemetryNiaFeatureUsageAllOf) GetWebtokenTimeoutSecondsOk() (*int64, bool)`

GetWebtokenTimeoutSecondsOk returns a tuple with the WebtokenTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebtokenTimeoutSeconds

`func (o *NiatelemetryNiaFeatureUsageAllOf) SetWebtokenTimeoutSeconds(v int64)`

SetWebtokenTimeoutSeconds sets WebtokenTimeoutSeconds field to given value.

### HasWebtokenTimeoutSeconds

`func (o *NiatelemetryNiaFeatureUsageAllOf) HasWebtokenTimeoutSeconds() bool`

HasWebtokenTimeoutSeconds returns a boolean if a field has been set.

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


