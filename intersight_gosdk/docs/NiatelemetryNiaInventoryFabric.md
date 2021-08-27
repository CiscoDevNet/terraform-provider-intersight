# NiatelemetryNiaInventoryFabric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.NiaInventoryFabric"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.NiaInventoryFabric"]
**AnycastGwMac** | Pointer to **string** | Returns the aycast gateway mac. | [optional] 
**BgpEstablishedInterfaceCount** | Pointer to **int64** | Counts the number of BGP interfaces that are in established state. | [optional] 
**BgwInterfaceUpCount** | Pointer to **int64** | Count number of active interfaces on border gateways. | [optional] 
**BorderGatewaySpineCount** | Pointer to **int64** | Count number of border gateway spines in the fabric inventory. | [optional] 
**BorderLeafCount** | Pointer to **int64** | Count number of border leafs in the fabric inventory. | [optional] 
**DciSubnetRange** | Pointer to **string** | Returns the dci subnet range. | [optional] 
**DciSubnetTargetMask** | Pointer to **string** | Returns the dci subnet target mask. | [optional] 
**DcnmtrackerEnabled** | Pointer to **bool** | Returns the value of the dcnmtrackerEnabled field. | [optional] 
**EbgpEvpnLinkUpCount** | Pointer to **int64** | Count number of ebgp evpn active interfaces. | [optional] 
**FabricId** | Pointer to **string** | Uniquely identifies a fabric. | [optional] 
**FabricName** | Pointer to **string** | Returns the value of the Name of a fabric. | [optional] 
**IsBgwPresent** | Pointer to **bool** | Checks if border gateway is present in the fabric inventory. | [optional] 
**IsNgoamEnabled** | Pointer to **bool** | Returns if ngoam is enabled. | [optional] 
**IsScheduledBackUpEnabled** | Pointer to **bool** | Returns if the scheduled backup is enabled. | [optional] 
**LeafCount** | Pointer to **int64** | Returns total number of leafs in the fabric. | [optional] 
**LogicalLinks** | Pointer to [**[]NiatelemetryLogicalLink**](NiatelemetryLogicalLink.md) |  | [optional] 
**NxosVniBwSitesCount** | Pointer to **int64** | Returns the count of vnis between sites. | [optional] 
**NxosVrfBwSitesCount** | Pointer to **int64** | Returns the count of vrfs between sites. | [optional] 
**NxosVrfCount** | Pointer to **int64** | Returns the value of the nxosVrfCount field. | [optional] 
**Serial** | Pointer to **string** | Serial number of device being inventoried. The serial number is unique per device. | [optional] 
**SiteName** | Pointer to **string** | Name of fabric domain of the controller. | [optional] 
**SpineCount** | Pointer to **int64** | Returns total number of spines in the fabric. | [optional] 
**VlanVniMappings** | Pointer to **string** | VLAN to VNI mappings configured in the DCNM. | [optional] 
**VniIpCount** | Pointer to **int64** | Count number of IP addresses configured in the DCNM networks. | [optional] 
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


