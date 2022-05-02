# NiatelemetryNiaInventoryFabric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.NiaInventoryFabric"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.NiaInventoryFabric"]
**AnycastGwMac** | Pointer to **string** | Returns the aycast gateway mac. | [optional] 
**BgpEstablishedInterfaceCount** | Pointer to **int64** | Counts the number of BGP interfaces that are in established state. | [optional] 
**BgwCount** | Pointer to **int64** | Returns number of bgw switches in the fabric. | [optional] 
**BgwInterfaceUpCount** | Pointer to **int64** | Count number of active interfaces on border gateways. | [optional] 
**BorderGatewaySpineCount** | Pointer to **int64** | Count number of border gateway spines in the fabric inventory. | [optional] 
**BorderLeafCount** | Pointer to **int64** | Count number of border leafs in the fabric inventory. | [optional] 
**CloudsecAutoconfig** | Pointer to **bool** | Cloudsec autoconfig details on the fabric. | [optional] 
**DciSubnetRange** | Pointer to **string** | Returns the dci subnet range. | [optional] 
**DciSubnetTargetMask** | Pointer to **string** | Returns the dci subnet target mask. | [optional] 
**DcnmtrackerEnabled** | Pointer to **bool** | Returns the value of the dcnmtrackerEnabled field. | [optional] 
**EbgpEvpnLinkUpCount** | Pointer to **int64** | Count number of ebgp evpn active interfaces. | [optional] 
**FabricId** | Pointer to **string** | Uniquely identifies a fabric. | [optional] 
**FabricName** | Pointer to **string** | Returns the value of the Name of a fabric. | [optional] 
**FabricParent** | Pointer to **string** | Parent of the fabric on DCNM. | [optional] 
**FabricTechnology** | Pointer to **string** | Fabric Technology details on the fabric. | [optional] 
**FabricType** | Pointer to **string** | Fabric type information string. | [optional] 
**FeaturePtp** | Pointer to **string** | PTP feature details on the fabric. | [optional] 
**IsBgwPresent** | Pointer to **bool** | Checks if border gateway is present in the fabric inventory. | [optional] 
**IsEnableNxapiHttp** | Pointer to **bool** | Check if NXAPI HTTP is enabled or not on the fabric. | [optional] 
**IsEnableRealTimeBackup** | Pointer to **bool** | Check if real time backup is enabled or not on the fabric. | [optional] 
**IsNgoamEnabled** | Pointer to **bool** | Returns if ngoam is enabled. | [optional] 
**IsScheduledBackUpEnabled** | Pointer to **bool** | Returns if the scheduled backup is enabled. | [optional] 
**IsTrmEnabled** | Pointer to **bool** | Is TRM enabled for the fabric. | [optional] 
**LeafCount** | Pointer to **int64** | Returns total number of leafs in the fabric. | [optional] 
**LinkStateRouting** | Pointer to **string** | Link state routing details on the fabric. | [optional] 
**LinkType** | Pointer to **string** | Fabric oper status information. | [optional] 
**LogicalLinks** | Pointer to [**[]NiatelemetryLogicalLink**](NiatelemetryLogicalLink.md) |  | [optional] 
**NetworkDeploymentCount** | Pointer to **int64** | No of networks deployed on a fabric. | [optional] 
**NetworkDeploymentStatus** | Pointer to [**[]NiatelemetryDeploymentStatus**](NiatelemetryDeploymentStatus.md) |  | [optional] 
**NtpServerIpList** | Pointer to **string** | NTP server IP List on the fabric. | [optional] 
**NxosVniBwSitesCount** | Pointer to **int64** | Returns the count of vnis between sites. | [optional] 
**NxosVrfBwSitesCount** | Pointer to **int64** | Returns the count of vrfs between sites. | [optional] 
**NxosVrfCount** | Pointer to **int64** | Returns the value of the nxosVrfCount field. | [optional] 
**OperStatus** | Pointer to **string** | Fabric oper status information. | [optional] 
**ReplicationMode** | Pointer to **string** | Replication mode details on the fabric. | [optional] 
**RpMode** | Pointer to **string** | RP Mode details on the fabric. | [optional] 
**Serial** | Pointer to **string** | Serial number of device being inventoried. The serial number is unique per device. | [optional] 
**SiteName** | Pointer to **string** | Name of fabric domain of the controller. | [optional] 
**SoftwareImage** | Pointer to **string** | Software image details on the fabric. | [optional] 
**SpineCount** | Pointer to **int64** | Returns total number of spines in the fabric. | [optional] 
**SyslogServerIpList** | Pointer to **string** | Syslog server IP list on DCNM. | [optional] 
**SyslogSev** | Pointer to **string** | Syslog sev details on the fabric. | [optional] 
**TemplateName** | Pointer to **string** | Template name of the fabric on DCNM. | [optional] 
**VlanVniMappings** | Pointer to **string** | VLAN to VNI mappings configured in the DCNM. | [optional] 
**VniIpCount** | Pointer to **int64** | Count number of IP addresses configured in the DCNM networks. | [optional] 
**VpcDetails** | Pointer to [**[]NiatelemetryVpcDetails**](NiatelemetryVpcDetails.md) |  | [optional] 
**VrfDeploymentCount** | Pointer to **int64** | No of vrfs deployed on a fabric. | [optional] 
**VrfDeploymentStatus** | Pointer to [**[]NiatelemetryDeploymentStatus**](NiatelemetryDeploymentStatus.md) |  | [optional] 
**XsiteNetworkCount** | Pointer to **int64** | Returns deployed network count for bgw/bgws switches in the MSD fabric. | [optional] 
**XsiteVrfCount** | Pointer to **int64** | Returns deployed vrf count for bgw/bgws switches in the MSD fabric. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetryNiaInventoryFabric

`func NewNiatelemetryNiaInventoryFabric(classId string, objectType string, ) *NiatelemetryNiaInventoryFabric`

NewNiatelemetryNiaInventoryFabric instantiates a new NiatelemetryNiaInventoryFabric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryNiaInventoryFabricWithDefaults

`func NewNiatelemetryNiaInventoryFabricWithDefaults() *NiatelemetryNiaInventoryFabric`

NewNiatelemetryNiaInventoryFabricWithDefaults instantiates a new NiatelemetryNiaInventoryFabric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryNiaInventoryFabric) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryNiaInventoryFabric) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryNiaInventoryFabric) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryNiaInventoryFabric) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryNiaInventoryFabric) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryNiaInventoryFabric) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAnycastGwMac

`func (o *NiatelemetryNiaInventoryFabric) GetAnycastGwMac() string`

GetAnycastGwMac returns the AnycastGwMac field if non-nil, zero value otherwise.

### GetAnycastGwMacOk

`func (o *NiatelemetryNiaInventoryFabric) GetAnycastGwMacOk() (*string, bool)`

GetAnycastGwMacOk returns a tuple with the AnycastGwMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnycastGwMac

`func (o *NiatelemetryNiaInventoryFabric) SetAnycastGwMac(v string)`

SetAnycastGwMac sets AnycastGwMac field to given value.

### HasAnycastGwMac

`func (o *NiatelemetryNiaInventoryFabric) HasAnycastGwMac() bool`

HasAnycastGwMac returns a boolean if a field has been set.

### GetBgpEstablishedInterfaceCount

`func (o *NiatelemetryNiaInventoryFabric) GetBgpEstablishedInterfaceCount() int64`

GetBgpEstablishedInterfaceCount returns the BgpEstablishedInterfaceCount field if non-nil, zero value otherwise.

### GetBgpEstablishedInterfaceCountOk

`func (o *NiatelemetryNiaInventoryFabric) GetBgpEstablishedInterfaceCountOk() (*int64, bool)`

GetBgpEstablishedInterfaceCountOk returns a tuple with the BgpEstablishedInterfaceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpEstablishedInterfaceCount

`func (o *NiatelemetryNiaInventoryFabric) SetBgpEstablishedInterfaceCount(v int64)`

SetBgpEstablishedInterfaceCount sets BgpEstablishedInterfaceCount field to given value.

### HasBgpEstablishedInterfaceCount

`func (o *NiatelemetryNiaInventoryFabric) HasBgpEstablishedInterfaceCount() bool`

HasBgpEstablishedInterfaceCount returns a boolean if a field has been set.

### GetBgwCount

`func (o *NiatelemetryNiaInventoryFabric) GetBgwCount() int64`

GetBgwCount returns the BgwCount field if non-nil, zero value otherwise.

### GetBgwCountOk

`func (o *NiatelemetryNiaInventoryFabric) GetBgwCountOk() (*int64, bool)`

GetBgwCountOk returns a tuple with the BgwCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgwCount

`func (o *NiatelemetryNiaInventoryFabric) SetBgwCount(v int64)`

SetBgwCount sets BgwCount field to given value.

### HasBgwCount

`func (o *NiatelemetryNiaInventoryFabric) HasBgwCount() bool`

HasBgwCount returns a boolean if a field has been set.

### GetBgwInterfaceUpCount

`func (o *NiatelemetryNiaInventoryFabric) GetBgwInterfaceUpCount() int64`

GetBgwInterfaceUpCount returns the BgwInterfaceUpCount field if non-nil, zero value otherwise.

### GetBgwInterfaceUpCountOk

`func (o *NiatelemetryNiaInventoryFabric) GetBgwInterfaceUpCountOk() (*int64, bool)`

GetBgwInterfaceUpCountOk returns a tuple with the BgwInterfaceUpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgwInterfaceUpCount

`func (o *NiatelemetryNiaInventoryFabric) SetBgwInterfaceUpCount(v int64)`

SetBgwInterfaceUpCount sets BgwInterfaceUpCount field to given value.

### HasBgwInterfaceUpCount

`func (o *NiatelemetryNiaInventoryFabric) HasBgwInterfaceUpCount() bool`

HasBgwInterfaceUpCount returns a boolean if a field has been set.

### GetBorderGatewaySpineCount

`func (o *NiatelemetryNiaInventoryFabric) GetBorderGatewaySpineCount() int64`

GetBorderGatewaySpineCount returns the BorderGatewaySpineCount field if non-nil, zero value otherwise.

### GetBorderGatewaySpineCountOk

`func (o *NiatelemetryNiaInventoryFabric) GetBorderGatewaySpineCountOk() (*int64, bool)`

GetBorderGatewaySpineCountOk returns a tuple with the BorderGatewaySpineCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorderGatewaySpineCount

`func (o *NiatelemetryNiaInventoryFabric) SetBorderGatewaySpineCount(v int64)`

SetBorderGatewaySpineCount sets BorderGatewaySpineCount field to given value.

### HasBorderGatewaySpineCount

`func (o *NiatelemetryNiaInventoryFabric) HasBorderGatewaySpineCount() bool`

HasBorderGatewaySpineCount returns a boolean if a field has been set.

### GetBorderLeafCount

`func (o *NiatelemetryNiaInventoryFabric) GetBorderLeafCount() int64`

GetBorderLeafCount returns the BorderLeafCount field if non-nil, zero value otherwise.

### GetBorderLeafCountOk

`func (o *NiatelemetryNiaInventoryFabric) GetBorderLeafCountOk() (*int64, bool)`

GetBorderLeafCountOk returns a tuple with the BorderLeafCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorderLeafCount

`func (o *NiatelemetryNiaInventoryFabric) SetBorderLeafCount(v int64)`

SetBorderLeafCount sets BorderLeafCount field to given value.

### HasBorderLeafCount

`func (o *NiatelemetryNiaInventoryFabric) HasBorderLeafCount() bool`

HasBorderLeafCount returns a boolean if a field has been set.

### GetCloudsecAutoconfig

`func (o *NiatelemetryNiaInventoryFabric) GetCloudsecAutoconfig() bool`

GetCloudsecAutoconfig returns the CloudsecAutoconfig field if non-nil, zero value otherwise.

### GetCloudsecAutoconfigOk

`func (o *NiatelemetryNiaInventoryFabric) GetCloudsecAutoconfigOk() (*bool, bool)`

GetCloudsecAutoconfigOk returns a tuple with the CloudsecAutoconfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudsecAutoconfig

`func (o *NiatelemetryNiaInventoryFabric) SetCloudsecAutoconfig(v bool)`

SetCloudsecAutoconfig sets CloudsecAutoconfig field to given value.

### HasCloudsecAutoconfig

`func (o *NiatelemetryNiaInventoryFabric) HasCloudsecAutoconfig() bool`

HasCloudsecAutoconfig returns a boolean if a field has been set.

### GetDciSubnetRange

`func (o *NiatelemetryNiaInventoryFabric) GetDciSubnetRange() string`

GetDciSubnetRange returns the DciSubnetRange field if non-nil, zero value otherwise.

### GetDciSubnetRangeOk

`func (o *NiatelemetryNiaInventoryFabric) GetDciSubnetRangeOk() (*string, bool)`

GetDciSubnetRangeOk returns a tuple with the DciSubnetRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDciSubnetRange

`func (o *NiatelemetryNiaInventoryFabric) SetDciSubnetRange(v string)`

SetDciSubnetRange sets DciSubnetRange field to given value.

### HasDciSubnetRange

`func (o *NiatelemetryNiaInventoryFabric) HasDciSubnetRange() bool`

HasDciSubnetRange returns a boolean if a field has been set.

### GetDciSubnetTargetMask

`func (o *NiatelemetryNiaInventoryFabric) GetDciSubnetTargetMask() string`

GetDciSubnetTargetMask returns the DciSubnetTargetMask field if non-nil, zero value otherwise.

### GetDciSubnetTargetMaskOk

`func (o *NiatelemetryNiaInventoryFabric) GetDciSubnetTargetMaskOk() (*string, bool)`

GetDciSubnetTargetMaskOk returns a tuple with the DciSubnetTargetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDciSubnetTargetMask

`func (o *NiatelemetryNiaInventoryFabric) SetDciSubnetTargetMask(v string)`

SetDciSubnetTargetMask sets DciSubnetTargetMask field to given value.

### HasDciSubnetTargetMask

`func (o *NiatelemetryNiaInventoryFabric) HasDciSubnetTargetMask() bool`

HasDciSubnetTargetMask returns a boolean if a field has been set.

### GetDcnmtrackerEnabled

`func (o *NiatelemetryNiaInventoryFabric) GetDcnmtrackerEnabled() bool`

GetDcnmtrackerEnabled returns the DcnmtrackerEnabled field if non-nil, zero value otherwise.

### GetDcnmtrackerEnabledOk

`func (o *NiatelemetryNiaInventoryFabric) GetDcnmtrackerEnabledOk() (*bool, bool)`

GetDcnmtrackerEnabledOk returns a tuple with the DcnmtrackerEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcnmtrackerEnabled

`func (o *NiatelemetryNiaInventoryFabric) SetDcnmtrackerEnabled(v bool)`

SetDcnmtrackerEnabled sets DcnmtrackerEnabled field to given value.

### HasDcnmtrackerEnabled

`func (o *NiatelemetryNiaInventoryFabric) HasDcnmtrackerEnabled() bool`

HasDcnmtrackerEnabled returns a boolean if a field has been set.

### GetEbgpEvpnLinkUpCount

`func (o *NiatelemetryNiaInventoryFabric) GetEbgpEvpnLinkUpCount() int64`

GetEbgpEvpnLinkUpCount returns the EbgpEvpnLinkUpCount field if non-nil, zero value otherwise.

### GetEbgpEvpnLinkUpCountOk

`func (o *NiatelemetryNiaInventoryFabric) GetEbgpEvpnLinkUpCountOk() (*int64, bool)`

GetEbgpEvpnLinkUpCountOk returns a tuple with the EbgpEvpnLinkUpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbgpEvpnLinkUpCount

`func (o *NiatelemetryNiaInventoryFabric) SetEbgpEvpnLinkUpCount(v int64)`

SetEbgpEvpnLinkUpCount sets EbgpEvpnLinkUpCount field to given value.

### HasEbgpEvpnLinkUpCount

`func (o *NiatelemetryNiaInventoryFabric) HasEbgpEvpnLinkUpCount() bool`

HasEbgpEvpnLinkUpCount returns a boolean if a field has been set.

### GetFabricId

`func (o *NiatelemetryNiaInventoryFabric) GetFabricId() string`

GetFabricId returns the FabricId field if non-nil, zero value otherwise.

### GetFabricIdOk

`func (o *NiatelemetryNiaInventoryFabric) GetFabricIdOk() (*string, bool)`

GetFabricIdOk returns a tuple with the FabricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricId

`func (o *NiatelemetryNiaInventoryFabric) SetFabricId(v string)`

SetFabricId sets FabricId field to given value.

### HasFabricId

`func (o *NiatelemetryNiaInventoryFabric) HasFabricId() bool`

HasFabricId returns a boolean if a field has been set.

### GetFabricName

`func (o *NiatelemetryNiaInventoryFabric) GetFabricName() string`

GetFabricName returns the FabricName field if non-nil, zero value otherwise.

### GetFabricNameOk

`func (o *NiatelemetryNiaInventoryFabric) GetFabricNameOk() (*string, bool)`

GetFabricNameOk returns a tuple with the FabricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricName

`func (o *NiatelemetryNiaInventoryFabric) SetFabricName(v string)`

SetFabricName sets FabricName field to given value.

### HasFabricName

`func (o *NiatelemetryNiaInventoryFabric) HasFabricName() bool`

HasFabricName returns a boolean if a field has been set.

### GetFabricParent

`func (o *NiatelemetryNiaInventoryFabric) GetFabricParent() string`

GetFabricParent returns the FabricParent field if non-nil, zero value otherwise.

### GetFabricParentOk

`func (o *NiatelemetryNiaInventoryFabric) GetFabricParentOk() (*string, bool)`

GetFabricParentOk returns a tuple with the FabricParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricParent

`func (o *NiatelemetryNiaInventoryFabric) SetFabricParent(v string)`

SetFabricParent sets FabricParent field to given value.

### HasFabricParent

`func (o *NiatelemetryNiaInventoryFabric) HasFabricParent() bool`

HasFabricParent returns a boolean if a field has been set.

### GetFabricTechnology

`func (o *NiatelemetryNiaInventoryFabric) GetFabricTechnology() string`

GetFabricTechnology returns the FabricTechnology field if non-nil, zero value otherwise.

### GetFabricTechnologyOk

`func (o *NiatelemetryNiaInventoryFabric) GetFabricTechnologyOk() (*string, bool)`

GetFabricTechnologyOk returns a tuple with the FabricTechnology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricTechnology

`func (o *NiatelemetryNiaInventoryFabric) SetFabricTechnology(v string)`

SetFabricTechnology sets FabricTechnology field to given value.

### HasFabricTechnology

`func (o *NiatelemetryNiaInventoryFabric) HasFabricTechnology() bool`

HasFabricTechnology returns a boolean if a field has been set.

### GetFabricType

`func (o *NiatelemetryNiaInventoryFabric) GetFabricType() string`

GetFabricType returns the FabricType field if non-nil, zero value otherwise.

### GetFabricTypeOk

`func (o *NiatelemetryNiaInventoryFabric) GetFabricTypeOk() (*string, bool)`

GetFabricTypeOk returns a tuple with the FabricType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricType

`func (o *NiatelemetryNiaInventoryFabric) SetFabricType(v string)`

SetFabricType sets FabricType field to given value.

### HasFabricType

`func (o *NiatelemetryNiaInventoryFabric) HasFabricType() bool`

HasFabricType returns a boolean if a field has been set.

### GetFeaturePtp

`func (o *NiatelemetryNiaInventoryFabric) GetFeaturePtp() string`

GetFeaturePtp returns the FeaturePtp field if non-nil, zero value otherwise.

### GetFeaturePtpOk

`func (o *NiatelemetryNiaInventoryFabric) GetFeaturePtpOk() (*string, bool)`

GetFeaturePtpOk returns a tuple with the FeaturePtp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturePtp

`func (o *NiatelemetryNiaInventoryFabric) SetFeaturePtp(v string)`

SetFeaturePtp sets FeaturePtp field to given value.

### HasFeaturePtp

`func (o *NiatelemetryNiaInventoryFabric) HasFeaturePtp() bool`

HasFeaturePtp returns a boolean if a field has been set.

### GetIsBgwPresent

`func (o *NiatelemetryNiaInventoryFabric) GetIsBgwPresent() bool`

GetIsBgwPresent returns the IsBgwPresent field if non-nil, zero value otherwise.

### GetIsBgwPresentOk

`func (o *NiatelemetryNiaInventoryFabric) GetIsBgwPresentOk() (*bool, bool)`

GetIsBgwPresentOk returns a tuple with the IsBgwPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBgwPresent

`func (o *NiatelemetryNiaInventoryFabric) SetIsBgwPresent(v bool)`

SetIsBgwPresent sets IsBgwPresent field to given value.

### HasIsBgwPresent

`func (o *NiatelemetryNiaInventoryFabric) HasIsBgwPresent() bool`

HasIsBgwPresent returns a boolean if a field has been set.

### GetIsEnableNxapiHttp

`func (o *NiatelemetryNiaInventoryFabric) GetIsEnableNxapiHttp() bool`

GetIsEnableNxapiHttp returns the IsEnableNxapiHttp field if non-nil, zero value otherwise.

### GetIsEnableNxapiHttpOk

`func (o *NiatelemetryNiaInventoryFabric) GetIsEnableNxapiHttpOk() (*bool, bool)`

GetIsEnableNxapiHttpOk returns a tuple with the IsEnableNxapiHttp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnableNxapiHttp

`func (o *NiatelemetryNiaInventoryFabric) SetIsEnableNxapiHttp(v bool)`

SetIsEnableNxapiHttp sets IsEnableNxapiHttp field to given value.

### HasIsEnableNxapiHttp

`func (o *NiatelemetryNiaInventoryFabric) HasIsEnableNxapiHttp() bool`

HasIsEnableNxapiHttp returns a boolean if a field has been set.

### GetIsEnableRealTimeBackup

`func (o *NiatelemetryNiaInventoryFabric) GetIsEnableRealTimeBackup() bool`

GetIsEnableRealTimeBackup returns the IsEnableRealTimeBackup field if non-nil, zero value otherwise.

### GetIsEnableRealTimeBackupOk

`func (o *NiatelemetryNiaInventoryFabric) GetIsEnableRealTimeBackupOk() (*bool, bool)`

GetIsEnableRealTimeBackupOk returns a tuple with the IsEnableRealTimeBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnableRealTimeBackup

`func (o *NiatelemetryNiaInventoryFabric) SetIsEnableRealTimeBackup(v bool)`

SetIsEnableRealTimeBackup sets IsEnableRealTimeBackup field to given value.

### HasIsEnableRealTimeBackup

`func (o *NiatelemetryNiaInventoryFabric) HasIsEnableRealTimeBackup() bool`

HasIsEnableRealTimeBackup returns a boolean if a field has been set.

### GetIsNgoamEnabled

`func (o *NiatelemetryNiaInventoryFabric) GetIsNgoamEnabled() bool`

GetIsNgoamEnabled returns the IsNgoamEnabled field if non-nil, zero value otherwise.

### GetIsNgoamEnabledOk

`func (o *NiatelemetryNiaInventoryFabric) GetIsNgoamEnabledOk() (*bool, bool)`

GetIsNgoamEnabledOk returns a tuple with the IsNgoamEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNgoamEnabled

`func (o *NiatelemetryNiaInventoryFabric) SetIsNgoamEnabled(v bool)`

SetIsNgoamEnabled sets IsNgoamEnabled field to given value.

### HasIsNgoamEnabled

`func (o *NiatelemetryNiaInventoryFabric) HasIsNgoamEnabled() bool`

HasIsNgoamEnabled returns a boolean if a field has been set.

### GetIsScheduledBackUpEnabled

`func (o *NiatelemetryNiaInventoryFabric) GetIsScheduledBackUpEnabled() bool`

GetIsScheduledBackUpEnabled returns the IsScheduledBackUpEnabled field if non-nil, zero value otherwise.

### GetIsScheduledBackUpEnabledOk

`func (o *NiatelemetryNiaInventoryFabric) GetIsScheduledBackUpEnabledOk() (*bool, bool)`

GetIsScheduledBackUpEnabledOk returns a tuple with the IsScheduledBackUpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsScheduledBackUpEnabled

`func (o *NiatelemetryNiaInventoryFabric) SetIsScheduledBackUpEnabled(v bool)`

SetIsScheduledBackUpEnabled sets IsScheduledBackUpEnabled field to given value.

### HasIsScheduledBackUpEnabled

`func (o *NiatelemetryNiaInventoryFabric) HasIsScheduledBackUpEnabled() bool`

HasIsScheduledBackUpEnabled returns a boolean if a field has been set.

### GetIsTrmEnabled

`func (o *NiatelemetryNiaInventoryFabric) GetIsTrmEnabled() bool`

GetIsTrmEnabled returns the IsTrmEnabled field if non-nil, zero value otherwise.

### GetIsTrmEnabledOk

`func (o *NiatelemetryNiaInventoryFabric) GetIsTrmEnabledOk() (*bool, bool)`

GetIsTrmEnabledOk returns a tuple with the IsTrmEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTrmEnabled

`func (o *NiatelemetryNiaInventoryFabric) SetIsTrmEnabled(v bool)`

SetIsTrmEnabled sets IsTrmEnabled field to given value.

### HasIsTrmEnabled

`func (o *NiatelemetryNiaInventoryFabric) HasIsTrmEnabled() bool`

HasIsTrmEnabled returns a boolean if a field has been set.

### GetLeafCount

`func (o *NiatelemetryNiaInventoryFabric) GetLeafCount() int64`

GetLeafCount returns the LeafCount field if non-nil, zero value otherwise.

### GetLeafCountOk

`func (o *NiatelemetryNiaInventoryFabric) GetLeafCountOk() (*int64, bool)`

GetLeafCountOk returns a tuple with the LeafCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeafCount

`func (o *NiatelemetryNiaInventoryFabric) SetLeafCount(v int64)`

SetLeafCount sets LeafCount field to given value.

### HasLeafCount

`func (o *NiatelemetryNiaInventoryFabric) HasLeafCount() bool`

HasLeafCount returns a boolean if a field has been set.

### GetLinkStateRouting

`func (o *NiatelemetryNiaInventoryFabric) GetLinkStateRouting() string`

GetLinkStateRouting returns the LinkStateRouting field if non-nil, zero value otherwise.

### GetLinkStateRoutingOk

`func (o *NiatelemetryNiaInventoryFabric) GetLinkStateRoutingOk() (*string, bool)`

GetLinkStateRoutingOk returns a tuple with the LinkStateRouting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkStateRouting

`func (o *NiatelemetryNiaInventoryFabric) SetLinkStateRouting(v string)`

SetLinkStateRouting sets LinkStateRouting field to given value.

### HasLinkStateRouting

`func (o *NiatelemetryNiaInventoryFabric) HasLinkStateRouting() bool`

HasLinkStateRouting returns a boolean if a field has been set.

### GetLinkType

`func (o *NiatelemetryNiaInventoryFabric) GetLinkType() string`

GetLinkType returns the LinkType field if non-nil, zero value otherwise.

### GetLinkTypeOk

`func (o *NiatelemetryNiaInventoryFabric) GetLinkTypeOk() (*string, bool)`

GetLinkTypeOk returns a tuple with the LinkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkType

`func (o *NiatelemetryNiaInventoryFabric) SetLinkType(v string)`

SetLinkType sets LinkType field to given value.

### HasLinkType

`func (o *NiatelemetryNiaInventoryFabric) HasLinkType() bool`

HasLinkType returns a boolean if a field has been set.

### GetLogicalLinks

`func (o *NiatelemetryNiaInventoryFabric) GetLogicalLinks() []NiatelemetryLogicalLink`

GetLogicalLinks returns the LogicalLinks field if non-nil, zero value otherwise.

### GetLogicalLinksOk

`func (o *NiatelemetryNiaInventoryFabric) GetLogicalLinksOk() (*[]NiatelemetryLogicalLink, bool)`

GetLogicalLinksOk returns a tuple with the LogicalLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalLinks

`func (o *NiatelemetryNiaInventoryFabric) SetLogicalLinks(v []NiatelemetryLogicalLink)`

SetLogicalLinks sets LogicalLinks field to given value.

### HasLogicalLinks

`func (o *NiatelemetryNiaInventoryFabric) HasLogicalLinks() bool`

HasLogicalLinks returns a boolean if a field has been set.

### SetLogicalLinksNil

`func (o *NiatelemetryNiaInventoryFabric) SetLogicalLinksNil(b bool)`

 SetLogicalLinksNil sets the value for LogicalLinks to be an explicit nil

### UnsetLogicalLinks
`func (o *NiatelemetryNiaInventoryFabric) UnsetLogicalLinks()`

UnsetLogicalLinks ensures that no value is present for LogicalLinks, not even an explicit nil
### GetNetworkDeploymentCount

`func (o *NiatelemetryNiaInventoryFabric) GetNetworkDeploymentCount() int64`

GetNetworkDeploymentCount returns the NetworkDeploymentCount field if non-nil, zero value otherwise.

### GetNetworkDeploymentCountOk

`func (o *NiatelemetryNiaInventoryFabric) GetNetworkDeploymentCountOk() (*int64, bool)`

GetNetworkDeploymentCountOk returns a tuple with the NetworkDeploymentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeploymentCount

`func (o *NiatelemetryNiaInventoryFabric) SetNetworkDeploymentCount(v int64)`

SetNetworkDeploymentCount sets NetworkDeploymentCount field to given value.

### HasNetworkDeploymentCount

`func (o *NiatelemetryNiaInventoryFabric) HasNetworkDeploymentCount() bool`

HasNetworkDeploymentCount returns a boolean if a field has been set.

### GetNetworkDeploymentStatus

`func (o *NiatelemetryNiaInventoryFabric) GetNetworkDeploymentStatus() []NiatelemetryDeploymentStatus`

GetNetworkDeploymentStatus returns the NetworkDeploymentStatus field if non-nil, zero value otherwise.

### GetNetworkDeploymentStatusOk

`func (o *NiatelemetryNiaInventoryFabric) GetNetworkDeploymentStatusOk() (*[]NiatelemetryDeploymentStatus, bool)`

GetNetworkDeploymentStatusOk returns a tuple with the NetworkDeploymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeploymentStatus

`func (o *NiatelemetryNiaInventoryFabric) SetNetworkDeploymentStatus(v []NiatelemetryDeploymentStatus)`

SetNetworkDeploymentStatus sets NetworkDeploymentStatus field to given value.

### HasNetworkDeploymentStatus

`func (o *NiatelemetryNiaInventoryFabric) HasNetworkDeploymentStatus() bool`

HasNetworkDeploymentStatus returns a boolean if a field has been set.

### SetNetworkDeploymentStatusNil

`func (o *NiatelemetryNiaInventoryFabric) SetNetworkDeploymentStatusNil(b bool)`

 SetNetworkDeploymentStatusNil sets the value for NetworkDeploymentStatus to be an explicit nil

### UnsetNetworkDeploymentStatus
`func (o *NiatelemetryNiaInventoryFabric) UnsetNetworkDeploymentStatus()`

UnsetNetworkDeploymentStatus ensures that no value is present for NetworkDeploymentStatus, not even an explicit nil
### GetNtpServerIpList

`func (o *NiatelemetryNiaInventoryFabric) GetNtpServerIpList() string`

GetNtpServerIpList returns the NtpServerIpList field if non-nil, zero value otherwise.

### GetNtpServerIpListOk

`func (o *NiatelemetryNiaInventoryFabric) GetNtpServerIpListOk() (*string, bool)`

GetNtpServerIpListOk returns a tuple with the NtpServerIpList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpServerIpList

`func (o *NiatelemetryNiaInventoryFabric) SetNtpServerIpList(v string)`

SetNtpServerIpList sets NtpServerIpList field to given value.

### HasNtpServerIpList

`func (o *NiatelemetryNiaInventoryFabric) HasNtpServerIpList() bool`

HasNtpServerIpList returns a boolean if a field has been set.

### GetNxosVniBwSitesCount

`func (o *NiatelemetryNiaInventoryFabric) GetNxosVniBwSitesCount() int64`

GetNxosVniBwSitesCount returns the NxosVniBwSitesCount field if non-nil, zero value otherwise.

### GetNxosVniBwSitesCountOk

`func (o *NiatelemetryNiaInventoryFabric) GetNxosVniBwSitesCountOk() (*int64, bool)`

GetNxosVniBwSitesCountOk returns a tuple with the NxosVniBwSitesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxosVniBwSitesCount

`func (o *NiatelemetryNiaInventoryFabric) SetNxosVniBwSitesCount(v int64)`

SetNxosVniBwSitesCount sets NxosVniBwSitesCount field to given value.

### HasNxosVniBwSitesCount

`func (o *NiatelemetryNiaInventoryFabric) HasNxosVniBwSitesCount() bool`

HasNxosVniBwSitesCount returns a boolean if a field has been set.

### GetNxosVrfBwSitesCount

`func (o *NiatelemetryNiaInventoryFabric) GetNxosVrfBwSitesCount() int64`

GetNxosVrfBwSitesCount returns the NxosVrfBwSitesCount field if non-nil, zero value otherwise.

### GetNxosVrfBwSitesCountOk

`func (o *NiatelemetryNiaInventoryFabric) GetNxosVrfBwSitesCountOk() (*int64, bool)`

GetNxosVrfBwSitesCountOk returns a tuple with the NxosVrfBwSitesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxosVrfBwSitesCount

`func (o *NiatelemetryNiaInventoryFabric) SetNxosVrfBwSitesCount(v int64)`

SetNxosVrfBwSitesCount sets NxosVrfBwSitesCount field to given value.

### HasNxosVrfBwSitesCount

`func (o *NiatelemetryNiaInventoryFabric) HasNxosVrfBwSitesCount() bool`

HasNxosVrfBwSitesCount returns a boolean if a field has been set.

### GetNxosVrfCount

`func (o *NiatelemetryNiaInventoryFabric) GetNxosVrfCount() int64`

GetNxosVrfCount returns the NxosVrfCount field if non-nil, zero value otherwise.

### GetNxosVrfCountOk

`func (o *NiatelemetryNiaInventoryFabric) GetNxosVrfCountOk() (*int64, bool)`

GetNxosVrfCountOk returns a tuple with the NxosVrfCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxosVrfCount

`func (o *NiatelemetryNiaInventoryFabric) SetNxosVrfCount(v int64)`

SetNxosVrfCount sets NxosVrfCount field to given value.

### HasNxosVrfCount

`func (o *NiatelemetryNiaInventoryFabric) HasNxosVrfCount() bool`

HasNxosVrfCount returns a boolean if a field has been set.

### GetOperStatus

`func (o *NiatelemetryNiaInventoryFabric) GetOperStatus() string`

GetOperStatus returns the OperStatus field if non-nil, zero value otherwise.

### GetOperStatusOk

`func (o *NiatelemetryNiaInventoryFabric) GetOperStatusOk() (*string, bool)`

GetOperStatusOk returns a tuple with the OperStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperStatus

`func (o *NiatelemetryNiaInventoryFabric) SetOperStatus(v string)`

SetOperStatus sets OperStatus field to given value.

### HasOperStatus

`func (o *NiatelemetryNiaInventoryFabric) HasOperStatus() bool`

HasOperStatus returns a boolean if a field has been set.

### GetReplicationMode

`func (o *NiatelemetryNiaInventoryFabric) GetReplicationMode() string`

GetReplicationMode returns the ReplicationMode field if non-nil, zero value otherwise.

### GetReplicationModeOk

`func (o *NiatelemetryNiaInventoryFabric) GetReplicationModeOk() (*string, bool)`

GetReplicationModeOk returns a tuple with the ReplicationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationMode

`func (o *NiatelemetryNiaInventoryFabric) SetReplicationMode(v string)`

SetReplicationMode sets ReplicationMode field to given value.

### HasReplicationMode

`func (o *NiatelemetryNiaInventoryFabric) HasReplicationMode() bool`

HasReplicationMode returns a boolean if a field has been set.

### GetRpMode

`func (o *NiatelemetryNiaInventoryFabric) GetRpMode() string`

GetRpMode returns the RpMode field if non-nil, zero value otherwise.

### GetRpModeOk

`func (o *NiatelemetryNiaInventoryFabric) GetRpModeOk() (*string, bool)`

GetRpModeOk returns a tuple with the RpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpMode

`func (o *NiatelemetryNiaInventoryFabric) SetRpMode(v string)`

SetRpMode sets RpMode field to given value.

### HasRpMode

`func (o *NiatelemetryNiaInventoryFabric) HasRpMode() bool`

HasRpMode returns a boolean if a field has been set.

### GetSerial

`func (o *NiatelemetryNiaInventoryFabric) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *NiatelemetryNiaInventoryFabric) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *NiatelemetryNiaInventoryFabric) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *NiatelemetryNiaInventoryFabric) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetryNiaInventoryFabric) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetryNiaInventoryFabric) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetryNiaInventoryFabric) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetryNiaInventoryFabric) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetSoftwareImage

`func (o *NiatelemetryNiaInventoryFabric) GetSoftwareImage() string`

GetSoftwareImage returns the SoftwareImage field if non-nil, zero value otherwise.

### GetSoftwareImageOk

`func (o *NiatelemetryNiaInventoryFabric) GetSoftwareImageOk() (*string, bool)`

GetSoftwareImageOk returns a tuple with the SoftwareImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareImage

`func (o *NiatelemetryNiaInventoryFabric) SetSoftwareImage(v string)`

SetSoftwareImage sets SoftwareImage field to given value.

### HasSoftwareImage

`func (o *NiatelemetryNiaInventoryFabric) HasSoftwareImage() bool`

HasSoftwareImage returns a boolean if a field has been set.

### GetSpineCount

`func (o *NiatelemetryNiaInventoryFabric) GetSpineCount() int64`

GetSpineCount returns the SpineCount field if non-nil, zero value otherwise.

### GetSpineCountOk

`func (o *NiatelemetryNiaInventoryFabric) GetSpineCountOk() (*int64, bool)`

GetSpineCountOk returns a tuple with the SpineCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpineCount

`func (o *NiatelemetryNiaInventoryFabric) SetSpineCount(v int64)`

SetSpineCount sets SpineCount field to given value.

### HasSpineCount

`func (o *NiatelemetryNiaInventoryFabric) HasSpineCount() bool`

HasSpineCount returns a boolean if a field has been set.

### GetSyslogServerIpList

`func (o *NiatelemetryNiaInventoryFabric) GetSyslogServerIpList() string`

GetSyslogServerIpList returns the SyslogServerIpList field if non-nil, zero value otherwise.

### GetSyslogServerIpListOk

`func (o *NiatelemetryNiaInventoryFabric) GetSyslogServerIpListOk() (*string, bool)`

GetSyslogServerIpListOk returns a tuple with the SyslogServerIpList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogServerIpList

`func (o *NiatelemetryNiaInventoryFabric) SetSyslogServerIpList(v string)`

SetSyslogServerIpList sets SyslogServerIpList field to given value.

### HasSyslogServerIpList

`func (o *NiatelemetryNiaInventoryFabric) HasSyslogServerIpList() bool`

HasSyslogServerIpList returns a boolean if a field has been set.

### GetSyslogSev

`func (o *NiatelemetryNiaInventoryFabric) GetSyslogSev() string`

GetSyslogSev returns the SyslogSev field if non-nil, zero value otherwise.

### GetSyslogSevOk

`func (o *NiatelemetryNiaInventoryFabric) GetSyslogSevOk() (*string, bool)`

GetSyslogSevOk returns a tuple with the SyslogSev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogSev

`func (o *NiatelemetryNiaInventoryFabric) SetSyslogSev(v string)`

SetSyslogSev sets SyslogSev field to given value.

### HasSyslogSev

`func (o *NiatelemetryNiaInventoryFabric) HasSyslogSev() bool`

HasSyslogSev returns a boolean if a field has been set.

### GetTemplateName

`func (o *NiatelemetryNiaInventoryFabric) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *NiatelemetryNiaInventoryFabric) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *NiatelemetryNiaInventoryFabric) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *NiatelemetryNiaInventoryFabric) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.

### GetVlanVniMappings

`func (o *NiatelemetryNiaInventoryFabric) GetVlanVniMappings() string`

GetVlanVniMappings returns the VlanVniMappings field if non-nil, zero value otherwise.

### GetVlanVniMappingsOk

`func (o *NiatelemetryNiaInventoryFabric) GetVlanVniMappingsOk() (*string, bool)`

GetVlanVniMappingsOk returns a tuple with the VlanVniMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanVniMappings

`func (o *NiatelemetryNiaInventoryFabric) SetVlanVniMappings(v string)`

SetVlanVniMappings sets VlanVniMappings field to given value.

### HasVlanVniMappings

`func (o *NiatelemetryNiaInventoryFabric) HasVlanVniMappings() bool`

HasVlanVniMappings returns a boolean if a field has been set.

### GetVniIpCount

`func (o *NiatelemetryNiaInventoryFabric) GetVniIpCount() int64`

GetVniIpCount returns the VniIpCount field if non-nil, zero value otherwise.

### GetVniIpCountOk

`func (o *NiatelemetryNiaInventoryFabric) GetVniIpCountOk() (*int64, bool)`

GetVniIpCountOk returns a tuple with the VniIpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVniIpCount

`func (o *NiatelemetryNiaInventoryFabric) SetVniIpCount(v int64)`

SetVniIpCount sets VniIpCount field to given value.

### HasVniIpCount

`func (o *NiatelemetryNiaInventoryFabric) HasVniIpCount() bool`

HasVniIpCount returns a boolean if a field has been set.

### GetVpcDetails

`func (o *NiatelemetryNiaInventoryFabric) GetVpcDetails() []NiatelemetryVpcDetails`

GetVpcDetails returns the VpcDetails field if non-nil, zero value otherwise.

### GetVpcDetailsOk

`func (o *NiatelemetryNiaInventoryFabric) GetVpcDetailsOk() (*[]NiatelemetryVpcDetails, bool)`

GetVpcDetailsOk returns a tuple with the VpcDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcDetails

`func (o *NiatelemetryNiaInventoryFabric) SetVpcDetails(v []NiatelemetryVpcDetails)`

SetVpcDetails sets VpcDetails field to given value.

### HasVpcDetails

`func (o *NiatelemetryNiaInventoryFabric) HasVpcDetails() bool`

HasVpcDetails returns a boolean if a field has been set.

### SetVpcDetailsNil

`func (o *NiatelemetryNiaInventoryFabric) SetVpcDetailsNil(b bool)`

 SetVpcDetailsNil sets the value for VpcDetails to be an explicit nil

### UnsetVpcDetails
`func (o *NiatelemetryNiaInventoryFabric) UnsetVpcDetails()`

UnsetVpcDetails ensures that no value is present for VpcDetails, not even an explicit nil
### GetVrfDeploymentCount

`func (o *NiatelemetryNiaInventoryFabric) GetVrfDeploymentCount() int64`

GetVrfDeploymentCount returns the VrfDeploymentCount field if non-nil, zero value otherwise.

### GetVrfDeploymentCountOk

`func (o *NiatelemetryNiaInventoryFabric) GetVrfDeploymentCountOk() (*int64, bool)`

GetVrfDeploymentCountOk returns a tuple with the VrfDeploymentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfDeploymentCount

`func (o *NiatelemetryNiaInventoryFabric) SetVrfDeploymentCount(v int64)`

SetVrfDeploymentCount sets VrfDeploymentCount field to given value.

### HasVrfDeploymentCount

`func (o *NiatelemetryNiaInventoryFabric) HasVrfDeploymentCount() bool`

HasVrfDeploymentCount returns a boolean if a field has been set.

### GetVrfDeploymentStatus

`func (o *NiatelemetryNiaInventoryFabric) GetVrfDeploymentStatus() []NiatelemetryDeploymentStatus`

GetVrfDeploymentStatus returns the VrfDeploymentStatus field if non-nil, zero value otherwise.

### GetVrfDeploymentStatusOk

`func (o *NiatelemetryNiaInventoryFabric) GetVrfDeploymentStatusOk() (*[]NiatelemetryDeploymentStatus, bool)`

GetVrfDeploymentStatusOk returns a tuple with the VrfDeploymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfDeploymentStatus

`func (o *NiatelemetryNiaInventoryFabric) SetVrfDeploymentStatus(v []NiatelemetryDeploymentStatus)`

SetVrfDeploymentStatus sets VrfDeploymentStatus field to given value.

### HasVrfDeploymentStatus

`func (o *NiatelemetryNiaInventoryFabric) HasVrfDeploymentStatus() bool`

HasVrfDeploymentStatus returns a boolean if a field has been set.

### SetVrfDeploymentStatusNil

`func (o *NiatelemetryNiaInventoryFabric) SetVrfDeploymentStatusNil(b bool)`

 SetVrfDeploymentStatusNil sets the value for VrfDeploymentStatus to be an explicit nil

### UnsetVrfDeploymentStatus
`func (o *NiatelemetryNiaInventoryFabric) UnsetVrfDeploymentStatus()`

UnsetVrfDeploymentStatus ensures that no value is present for VrfDeploymentStatus, not even an explicit nil
### GetXsiteNetworkCount

`func (o *NiatelemetryNiaInventoryFabric) GetXsiteNetworkCount() int64`

GetXsiteNetworkCount returns the XsiteNetworkCount field if non-nil, zero value otherwise.

### GetXsiteNetworkCountOk

`func (o *NiatelemetryNiaInventoryFabric) GetXsiteNetworkCountOk() (*int64, bool)`

GetXsiteNetworkCountOk returns a tuple with the XsiteNetworkCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXsiteNetworkCount

`func (o *NiatelemetryNiaInventoryFabric) SetXsiteNetworkCount(v int64)`

SetXsiteNetworkCount sets XsiteNetworkCount field to given value.

### HasXsiteNetworkCount

`func (o *NiatelemetryNiaInventoryFabric) HasXsiteNetworkCount() bool`

HasXsiteNetworkCount returns a boolean if a field has been set.

### GetXsiteVrfCount

`func (o *NiatelemetryNiaInventoryFabric) GetXsiteVrfCount() int64`

GetXsiteVrfCount returns the XsiteVrfCount field if non-nil, zero value otherwise.

### GetXsiteVrfCountOk

`func (o *NiatelemetryNiaInventoryFabric) GetXsiteVrfCountOk() (*int64, bool)`

GetXsiteVrfCountOk returns a tuple with the XsiteVrfCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXsiteVrfCount

`func (o *NiatelemetryNiaInventoryFabric) SetXsiteVrfCount(v int64)`

SetXsiteVrfCount sets XsiteVrfCount field to given value.

### HasXsiteVrfCount

`func (o *NiatelemetryNiaInventoryFabric) HasXsiteVrfCount() bool`

HasXsiteVrfCount returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryNiaInventoryFabric) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryNiaInventoryFabric) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryNiaInventoryFabric) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryNiaInventoryFabric) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


