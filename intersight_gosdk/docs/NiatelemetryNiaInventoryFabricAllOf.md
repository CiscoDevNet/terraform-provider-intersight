# NiatelemetryNiaInventoryFabricAllOf

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

### NewNiatelemetryNiaInventoryFabricAllOf

`func NewNiatelemetryNiaInventoryFabricAllOf(classId string, objectType string, ) *NiatelemetryNiaInventoryFabricAllOf`

NewNiatelemetryNiaInventoryFabricAllOf instantiates a new NiatelemetryNiaInventoryFabricAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryNiaInventoryFabricAllOfWithDefaults

`func NewNiatelemetryNiaInventoryFabricAllOfWithDefaults() *NiatelemetryNiaInventoryFabricAllOf`

NewNiatelemetryNiaInventoryFabricAllOfWithDefaults instantiates a new NiatelemetryNiaInventoryFabricAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryNiaInventoryFabricAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryNiaInventoryFabricAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAnycastGwMac

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetAnycastGwMac() string`

GetAnycastGwMac returns the AnycastGwMac field if non-nil, zero value otherwise.

### GetAnycastGwMacOk

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetAnycastGwMacOk() (*string, bool)`

GetAnycastGwMacOk returns a tuple with the AnycastGwMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnycastGwMac

`func (o *NiatelemetryNiaInventoryFabricAllOf) SetAnycastGwMac(v string)`

SetAnycastGwMac sets AnycastGwMac field to given value.

### HasAnycastGwMac

`func (o *NiatelemetryNiaInventoryFabricAllOf) HasAnycastGwMac() bool`

HasAnycastGwMac returns a boolean if a field has been set.

### GetBgpEstablishedInterfaceCount

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetBgpEstablishedInterfaceCount() int64`

GetBgpEstablishedInterfaceCount returns the BgpEstablishedInterfaceCount field if non-nil, zero value otherwise.

### GetBgpEstablishedInterfaceCountOk

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetBgpEstablishedInterfaceCountOk() (*int64, bool)`

GetBgpEstablishedInterfaceCountOk returns a tuple with the BgpEstablishedInterfaceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpEstablishedInterfaceCount

`func (o *NiatelemetryNiaInventoryFabricAllOf) SetBgpEstablishedInterfaceCount(v int64)`

SetBgpEstablishedInterfaceCount sets BgpEstablishedInterfaceCount field to given value.

### HasBgpEstablishedInterfaceCount

`func (o *NiatelemetryNiaInventoryFabricAllOf) HasBgpEstablishedInterfaceCount() bool`

HasBgpEstablishedInterfaceCount returns a boolean if a field has been set.

### GetBgwInterfaceUpCount

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetBgwInterfaceUpCount() int64`

GetBgwInterfaceUpCount returns the BgwInterfaceUpCount field if non-nil, zero value otherwise.

### GetBgwInterfaceUpCountOk

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetBgwInterfaceUpCountOk() (*int64, bool)`

GetBgwInterfaceUpCountOk returns a tuple with the BgwInterfaceUpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgwInterfaceUpCount

`func (o *NiatelemetryNiaInventoryFabricAllOf) SetBgwInterfaceUpCount(v int64)`

SetBgwInterfaceUpCount sets BgwInterfaceUpCount field to given value.

### HasBgwInterfaceUpCount

`func (o *NiatelemetryNiaInventoryFabricAllOf) HasBgwInterfaceUpCount() bool`

HasBgwInterfaceUpCount returns a boolean if a field has been set.

### GetBorderGatewaySpineCount

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetBorderGatewaySpineCount() int64`

GetBorderGatewaySpineCount returns the BorderGatewaySpineCount field if non-nil, zero value otherwise.

### GetBorderGatewaySpineCountOk

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetBorderGatewaySpineCountOk() (*int64, bool)`

GetBorderGatewaySpineCountOk returns a tuple with the BorderGatewaySpineCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorderGatewaySpineCount

`func (o *NiatelemetryNiaInventoryFabricAllOf) SetBorderGatewaySpineCount(v int64)`

SetBorderGatewaySpineCount sets BorderGatewaySpineCount field to given value.

### HasBorderGatewaySpineCount

`func (o *NiatelemetryNiaInventoryFabricAllOf) HasBorderGatewaySpineCount() bool`

HasBorderGatewaySpineCount returns a boolean if a field has been set.

### GetBorderLeafCount

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetBorderLeafCount() int64`

GetBorderLeafCount returns the BorderLeafCount field if non-nil, zero value otherwise.

### GetBorderLeafCountOk

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetBorderLeafCountOk() (*int64, bool)`

GetBorderLeafCountOk returns a tuple with the BorderLeafCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorderLeafCount

`func (o *NiatelemetryNiaInventoryFabricAllOf) SetBorderLeafCount(v int64)`

SetBorderLeafCount sets BorderLeafCount field to given value.

### HasBorderLeafCount

`func (o *NiatelemetryNiaInventoryFabricAllOf) HasBorderLeafCount() bool`

HasBorderLeafCount returns a boolean if a field has been set.

### GetDciSubnetRange

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetDciSubnetRange() string`

GetDciSubnetRange returns the DciSubnetRange field if non-nil, zero value otherwise.

### GetDciSubnetRangeOk

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetDciSubnetRangeOk() (*string, bool)`

GetDciSubnetRangeOk returns a tuple with the DciSubnetRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDciSubnetRange

`func (o *NiatelemetryNiaInventoryFabricAllOf) SetDciSubnetRange(v string)`

SetDciSubnetRange sets DciSubnetRange field to given value.

### HasDciSubnetRange

`func (o *NiatelemetryNiaInventoryFabricAllOf) HasDciSubnetRange() bool`

HasDciSubnetRange returns a boolean if a field has been set.

### GetDciSubnetTargetMask

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetDciSubnetTargetMask() string`

GetDciSubnetTargetMask returns the DciSubnetTargetMask field if non-nil, zero value otherwise.

### GetDciSubnetTargetMaskOk

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetDciSubnetTargetMaskOk() (*string, bool)`

GetDciSubnetTargetMaskOk returns a tuple with the DciSubnetTargetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDciSubnetTargetMask

`func (o *NiatelemetryNiaInventoryFabricAllOf) SetDciSubnetTargetMask(v string)`

SetDciSubnetTargetMask sets DciSubnetTargetMask field to given value.

### HasDciSubnetTargetMask

`func (o *NiatelemetryNiaInventoryFabricAllOf) HasDciSubnetTargetMask() bool`

HasDciSubnetTargetMask returns a boolean if a field has been set.

### GetDcnmtrackerEnabled

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetDcnmtrackerEnabled() bool`

GetDcnmtrackerEnabled returns the DcnmtrackerEnabled field if non-nil, zero value otherwise.

### GetDcnmtrackerEnabledOk

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetDcnmtrackerEnabledOk() (*bool, bool)`

GetDcnmtrackerEnabledOk returns a tuple with the DcnmtrackerEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcnmtrackerEnabled

`func (o *NiatelemetryNiaInventoryFabricAllOf) SetDcnmtrackerEnabled(v bool)`

SetDcnmtrackerEnabled sets DcnmtrackerEnabled field to given value.

### HasDcnmtrackerEnabled

`func (o *NiatelemetryNiaInventoryFabricAllOf) HasDcnmtrackerEnabled() bool`

HasDcnmtrackerEnabled returns a boolean if a field has been set.

### GetEbgpEvpnLinkUpCount

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetEbgpEvpnLinkUpCount() int64`

GetEbgpEvpnLinkUpCount returns the EbgpEvpnLinkUpCount field if non-nil, zero value otherwise.

### GetEbgpEvpnLinkUpCountOk

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetEbgpEvpnLinkUpCountOk() (*int64, bool)`

GetEbgpEvpnLinkUpCountOk returns a tuple with the EbgpEvpnLinkUpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbgpEvpnLinkUpCount

`func (o *NiatelemetryNiaInventoryFabricAllOf) SetEbgpEvpnLinkUpCount(v int64)`

SetEbgpEvpnLinkUpCount sets EbgpEvpnLinkUpCount field to given value.

### HasEbgpEvpnLinkUpCount

`func (o *NiatelemetryNiaInventoryFabricAllOf) HasEbgpEvpnLinkUpCount() bool`

HasEbgpEvpnLinkUpCount returns a boolean if a field has been set.

### GetFabricId

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetFabricId() string`

GetFabricId returns the FabricId field if non-nil, zero value otherwise.

### GetFabricIdOk

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetFabricIdOk() (*string, bool)`

GetFabricIdOk returns a tuple with the FabricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricId

`func (o *NiatelemetryNiaInventoryFabricAllOf) SetFabricId(v string)`

SetFabricId sets FabricId field to given value.

### HasFabricId

`func (o *NiatelemetryNiaInventoryFabricAllOf) HasFabricId() bool`

HasFabricId returns a boolean if a field has been set.

### GetFabricName

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetFabricName() string`

GetFabricName returns the FabricName field if non-nil, zero value otherwise.

### GetFabricNameOk

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetFabricNameOk() (*string, bool)`

GetFabricNameOk returns a tuple with the FabricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricName

`func (o *NiatelemetryNiaInventoryFabricAllOf) SetFabricName(v string)`

SetFabricName sets FabricName field to given value.

### HasFabricName

`func (o *NiatelemetryNiaInventoryFabricAllOf) HasFabricName() bool`

HasFabricName returns a boolean if a field has been set.

### GetIsBgwPresent

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetIsBgwPresent() bool`

GetIsBgwPresent returns the IsBgwPresent field if non-nil, zero value otherwise.

### GetIsBgwPresentOk

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetIsBgwPresentOk() (*bool, bool)`

GetIsBgwPresentOk returns a tuple with the IsBgwPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBgwPresent

`func (o *NiatelemetryNiaInventoryFabricAllOf) SetIsBgwPresent(v bool)`

SetIsBgwPresent sets IsBgwPresent field to given value.

### HasIsBgwPresent

`func (o *NiatelemetryNiaInventoryFabricAllOf) HasIsBgwPresent() bool`

HasIsBgwPresent returns a boolean if a field has been set.

### GetIsNgoamEnabled

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetIsNgoamEnabled() bool`

GetIsNgoamEnabled returns the IsNgoamEnabled field if non-nil, zero value otherwise.

### GetIsNgoamEnabledOk

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetIsNgoamEnabledOk() (*bool, bool)`

GetIsNgoamEnabledOk returns a tuple with the IsNgoamEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNgoamEnabled

`func (o *NiatelemetryNiaInventoryFabricAllOf) SetIsNgoamEnabled(v bool)`

SetIsNgoamEnabled sets IsNgoamEnabled field to given value.

### HasIsNgoamEnabled

`func (o *NiatelemetryNiaInventoryFabricAllOf) HasIsNgoamEnabled() bool`

HasIsNgoamEnabled returns a boolean if a field has been set.

### GetIsScheduledBackUpEnabled

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetIsScheduledBackUpEnabled() bool`

GetIsScheduledBackUpEnabled returns the IsScheduledBackUpEnabled field if non-nil, zero value otherwise.

### GetIsScheduledBackUpEnabledOk

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetIsScheduledBackUpEnabledOk() (*bool, bool)`

GetIsScheduledBackUpEnabledOk returns a tuple with the IsScheduledBackUpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsScheduledBackUpEnabled

`func (o *NiatelemetryNiaInventoryFabricAllOf) SetIsScheduledBackUpEnabled(v bool)`

SetIsScheduledBackUpEnabled sets IsScheduledBackUpEnabled field to given value.

### HasIsScheduledBackUpEnabled

`func (o *NiatelemetryNiaInventoryFabricAllOf) HasIsScheduledBackUpEnabled() bool`

HasIsScheduledBackUpEnabled returns a boolean if a field has been set.

### GetLeafCount

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetLeafCount() int64`

GetLeafCount returns the LeafCount field if non-nil, zero value otherwise.

### GetLeafCountOk

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetLeafCountOk() (*int64, bool)`

GetLeafCountOk returns a tuple with the LeafCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeafCount

`func (o *NiatelemetryNiaInventoryFabricAllOf) SetLeafCount(v int64)`

SetLeafCount sets LeafCount field to given value.

### HasLeafCount

`func (o *NiatelemetryNiaInventoryFabricAllOf) HasLeafCount() bool`

HasLeafCount returns a boolean if a field has been set.

### GetLogicalLinks

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetLogicalLinks() []NiatelemetryLogicalLink`

GetLogicalLinks returns the LogicalLinks field if non-nil, zero value otherwise.

### GetLogicalLinksOk

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetLogicalLinksOk() (*[]NiatelemetryLogicalLink, bool)`

GetLogicalLinksOk returns a tuple with the LogicalLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalLinks

`func (o *NiatelemetryNiaInventoryFabricAllOf) SetLogicalLinks(v []NiatelemetryLogicalLink)`

SetLogicalLinks sets LogicalLinks field to given value.

### HasLogicalLinks

`func (o *NiatelemetryNiaInventoryFabricAllOf) HasLogicalLinks() bool`

HasLogicalLinks returns a boolean if a field has been set.

### SetLogicalLinksNil

`func (o *NiatelemetryNiaInventoryFabricAllOf) SetLogicalLinksNil(b bool)`

 SetLogicalLinksNil sets the value for LogicalLinks to be an explicit nil

### UnsetLogicalLinks
`func (o *NiatelemetryNiaInventoryFabricAllOf) UnsetLogicalLinks()`

UnsetLogicalLinks ensures that no value is present for LogicalLinks, not even an explicit nil
### GetNxosVniBwSitesCount

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetNxosVniBwSitesCount() int64`

GetNxosVniBwSitesCount returns the NxosVniBwSitesCount field if non-nil, zero value otherwise.

### GetNxosVniBwSitesCountOk

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetNxosVniBwSitesCountOk() (*int64, bool)`

GetNxosVniBwSitesCountOk returns a tuple with the NxosVniBwSitesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxosVniBwSitesCount

`func (o *NiatelemetryNiaInventoryFabricAllOf) SetNxosVniBwSitesCount(v int64)`

SetNxosVniBwSitesCount sets NxosVniBwSitesCount field to given value.

### HasNxosVniBwSitesCount

`func (o *NiatelemetryNiaInventoryFabricAllOf) HasNxosVniBwSitesCount() bool`

HasNxosVniBwSitesCount returns a boolean if a field has been set.

### GetNxosVrfBwSitesCount

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetNxosVrfBwSitesCount() int64`

GetNxosVrfBwSitesCount returns the NxosVrfBwSitesCount field if non-nil, zero value otherwise.

### GetNxosVrfBwSitesCountOk

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetNxosVrfBwSitesCountOk() (*int64, bool)`

GetNxosVrfBwSitesCountOk returns a tuple with the NxosVrfBwSitesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxosVrfBwSitesCount

`func (o *NiatelemetryNiaInventoryFabricAllOf) SetNxosVrfBwSitesCount(v int64)`

SetNxosVrfBwSitesCount sets NxosVrfBwSitesCount field to given value.

### HasNxosVrfBwSitesCount

`func (o *NiatelemetryNiaInventoryFabricAllOf) HasNxosVrfBwSitesCount() bool`

HasNxosVrfBwSitesCount returns a boolean if a field has been set.

### GetNxosVrfCount

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetNxosVrfCount() int64`

GetNxosVrfCount returns the NxosVrfCount field if non-nil, zero value otherwise.

### GetNxosVrfCountOk

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetNxosVrfCountOk() (*int64, bool)`

GetNxosVrfCountOk returns a tuple with the NxosVrfCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxosVrfCount

`func (o *NiatelemetryNiaInventoryFabricAllOf) SetNxosVrfCount(v int64)`

SetNxosVrfCount sets NxosVrfCount field to given value.

### HasNxosVrfCount

`func (o *NiatelemetryNiaInventoryFabricAllOf) HasNxosVrfCount() bool`

HasNxosVrfCount returns a boolean if a field has been set.

### GetSerial

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *NiatelemetryNiaInventoryFabricAllOf) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *NiatelemetryNiaInventoryFabricAllOf) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetryNiaInventoryFabricAllOf) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetryNiaInventoryFabricAllOf) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetSpineCount

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetSpineCount() int64`

GetSpineCount returns the SpineCount field if non-nil, zero value otherwise.

### GetSpineCountOk

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetSpineCountOk() (*int64, bool)`

GetSpineCountOk returns a tuple with the SpineCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpineCount

`func (o *NiatelemetryNiaInventoryFabricAllOf) SetSpineCount(v int64)`

SetSpineCount sets SpineCount field to given value.

### HasSpineCount

`func (o *NiatelemetryNiaInventoryFabricAllOf) HasSpineCount() bool`

HasSpineCount returns a boolean if a field has been set.

### GetVlanVniMappings

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetVlanVniMappings() string`

GetVlanVniMappings returns the VlanVniMappings field if non-nil, zero value otherwise.

### GetVlanVniMappingsOk

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetVlanVniMappingsOk() (*string, bool)`

GetVlanVniMappingsOk returns a tuple with the VlanVniMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanVniMappings

`func (o *NiatelemetryNiaInventoryFabricAllOf) SetVlanVniMappings(v string)`

SetVlanVniMappings sets VlanVniMappings field to given value.

### HasVlanVniMappings

`func (o *NiatelemetryNiaInventoryFabricAllOf) HasVlanVniMappings() bool`

HasVlanVniMappings returns a boolean if a field has been set.

### GetVniIpCount

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetVniIpCount() int64`

GetVniIpCount returns the VniIpCount field if non-nil, zero value otherwise.

### GetVniIpCountOk

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetVniIpCountOk() (*int64, bool)`

GetVniIpCountOk returns a tuple with the VniIpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVniIpCount

`func (o *NiatelemetryNiaInventoryFabricAllOf) SetVniIpCount(v int64)`

SetVniIpCount sets VniIpCount field to given value.

### HasVniIpCount

`func (o *NiatelemetryNiaInventoryFabricAllOf) HasVniIpCount() bool`

HasVniIpCount returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryNiaInventoryFabricAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryNiaInventoryFabricAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


