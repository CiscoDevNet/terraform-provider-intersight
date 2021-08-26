# VnicIscsiBootPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.IscsiBootPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.IscsiBootPolicy"]
**AutoTargetvendorName** | Pointer to **string** | Auto target interface that is represented via the Initiator name or the DHCP vendor ID. The vendor ID can be up to 32 alphanumeric characters. | [optional] 
**Chap** | Pointer to [**NullableVnicIscsiAuthProfile**](VnicIscsiAuthProfile.md) |  | [optional] 
**InitiatorIpSource** | Pointer to **string** | Source Type of Initiator IP Address - Auto/Static/Pool. * &#x60;DHCP&#x60; - The IP address is assigned using DHCP, if available. * &#x60;Static&#x60; - Static IPv4 address is assigned to the iSCSI boot interface based on the information entered in this area. * &#x60;Pool&#x60; - An IPv4 address is assigned to the iSCSI boot interface from the management IP address pool. | [optional] [default to "DHCP"]
**InitiatorStaticIpV4Address** | Pointer to **string** | Static IP address provided for iSCSI Initiator. | [optional] 
**InitiatorStaticIpV4Config** | Pointer to [**NullableIppoolIpV4Config**](IppoolIpV4Config.md) |  | [optional] 
**MutualChap** | Pointer to [**NullableVnicIscsiAuthProfile**](VnicIscsiAuthProfile.md) |  | [optional] 
**TargetSourceType** | Pointer to **string** | Source Type of Targets - Auto/Static. * &#x60;Static&#x60; - Type indicates that static target interface is assigned to iSCSI boot. * &#x60;Auto&#x60; - Type indicates that the system selects the target interface automatically during iSCSI boot. | [optional] [default to "Static"]
**InitiatorIpPool** | Pointer to [**IppoolPoolRelationship**](IppoolPoolRelationship.md) |  | [optional] 
**IscsiAdapterPolicy** | Pointer to [**VnicIscsiAdapterPolicyRelationship**](VnicIscsiAdapterPolicyRelationship.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**PrimaryTargetPolicy** | Pointer to [**VnicIscsiStaticTargetPolicyRelationship**](VnicIscsiStaticTargetPolicyRelationship.md) |  | [optional] 
**SecondaryTargetPolicy** | Pointer to [**VnicIscsiStaticTargetPolicyRelationship**](VnicIscsiStaticTargetPolicyRelationship.md) |  | [optional] 

## Methods

### NewVnicIscsiBootPolicyAllOf

`func NewVnicIscsiBootPolicyAllOf(classId string, objectType string, ) *VnicIscsiBootPolicyAllOf`

NewVnicIscsiBootPolicyAllOf instantiates a new VnicIscsiBootPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicIscsiBootPolicyAllOfWithDefaults

`func NewVnicIscsiBootPolicyAllOfWithDefaults() *VnicIscsiBootPolicyAllOf`

NewVnicIscsiBootPolicyAllOfWithDefaults instantiates a new VnicIscsiBootPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicIscsiBootPolicyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicIscsiBootPolicyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicIscsiBootPolicyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicIscsiBootPolicyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicIscsiBootPolicyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicIscsiBootPolicyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAutoTargetvendorName

`func (o *VnicIscsiBootPolicyAllOf) GetAutoTargetvendorName() string`

GetAutoTargetvendorName returns the AutoTargetvendorName field if non-nil, zero value otherwise.

### GetAutoTargetvendorNameOk

`func (o *VnicIscsiBootPolicyAllOf) GetAutoTargetvendorNameOk() (*string, bool)`

GetAutoTargetvendorNameOk returns a tuple with the AutoTargetvendorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTargetvendorName

`func (o *VnicIscsiBootPolicyAllOf) SetAutoTargetvendorName(v string)`

SetAutoTargetvendorName sets AutoTargetvendorName field to given value.

### HasAutoTargetvendorName

`func (o *VnicIscsiBootPolicyAllOf) HasAutoTargetvendorName() bool`

HasAutoTargetvendorName returns a boolean if a field has been set.

### GetChap

`func (o *VnicIscsiBootPolicyAllOf) GetChap() VnicIscsiAuthProfile`

GetChap returns the Chap field if non-nil, zero value otherwise.

### GetChapOk

`func (o *VnicIscsiBootPolicyAllOf) GetChapOk() (*VnicIscsiAuthProfile, bool)`

GetChapOk returns a tuple with the Chap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChap

`func (o *VnicIscsiBootPolicyAllOf) SetChap(v VnicIscsiAuthProfile)`

SetChap sets Chap field to given value.

### HasChap

`func (o *VnicIscsiBootPolicyAllOf) HasChap() bool`

HasChap returns a boolean if a field has been set.

### SetChapNil

`func (o *VnicIscsiBootPolicyAllOf) SetChapNil(b bool)`

 SetChapNil sets the value for Chap to be an explicit nil

### UnsetChap
`func (o *VnicIscsiBootPolicyAllOf) UnsetChap()`

UnsetChap ensures that no value is present for Chap, not even an explicit nil
### GetInitiatorIpSource

`func (o *VnicIscsiBootPolicyAllOf) GetInitiatorIpSource() string`

GetInitiatorIpSource returns the InitiatorIpSource field if non-nil, zero value otherwise.

### GetInitiatorIpSourceOk

`func (o *VnicIscsiBootPolicyAllOf) GetInitiatorIpSourceOk() (*string, bool)`

GetInitiatorIpSourceOk returns a tuple with the InitiatorIpSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorIpSource

`func (o *VnicIscsiBootPolicyAllOf) SetInitiatorIpSource(v string)`

SetInitiatorIpSource sets InitiatorIpSource field to given value.

### HasInitiatorIpSource

`func (o *VnicIscsiBootPolicyAllOf) HasInitiatorIpSource() bool`

HasInitiatorIpSource returns a boolean if a field has been set.

### GetInitiatorStaticIpV4Address

`func (o *VnicIscsiBootPolicyAllOf) GetInitiatorStaticIpV4Address() string`

GetInitiatorStaticIpV4Address returns the InitiatorStaticIpV4Address field if non-nil, zero value otherwise.

### GetInitiatorStaticIpV4AddressOk

`func (o *VnicIscsiBootPolicyAllOf) GetInitiatorStaticIpV4AddressOk() (*string, bool)`

GetInitiatorStaticIpV4AddressOk returns a tuple with the InitiatorStaticIpV4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorStaticIpV4Address

`func (o *VnicIscsiBootPolicyAllOf) SetInitiatorStaticIpV4Address(v string)`

SetInitiatorStaticIpV4Address sets InitiatorStaticIpV4Address field to given value.

### HasInitiatorStaticIpV4Address

`func (o *VnicIscsiBootPolicyAllOf) HasInitiatorStaticIpV4Address() bool`

HasInitiatorStaticIpV4Address returns a boolean if a field has been set.

### GetInitiatorStaticIpV4Config

`func (o *VnicIscsiBootPolicyAllOf) GetInitiatorStaticIpV4Config() IppoolIpV4Config`

GetInitiatorStaticIpV4Config returns the InitiatorStaticIpV4Config field if non-nil, zero value otherwise.

### GetInitiatorStaticIpV4ConfigOk

`func (o *VnicIscsiBootPolicyAllOf) GetInitiatorStaticIpV4ConfigOk() (*IppoolIpV4Config, bool)`

GetInitiatorStaticIpV4ConfigOk returns a tuple with the InitiatorStaticIpV4Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorStaticIpV4Config

`func (o *VnicIscsiBootPolicyAllOf) SetInitiatorStaticIpV4Config(v IppoolIpV4Config)`

SetInitiatorStaticIpV4Config sets InitiatorStaticIpV4Config field to given value.

### HasInitiatorStaticIpV4Config

`func (o *VnicIscsiBootPolicyAllOf) HasInitiatorStaticIpV4Config() bool`

HasInitiatorStaticIpV4Config returns a boolean if a field has been set.

### SetInitiatorStaticIpV4ConfigNil

`func (o *VnicIscsiBootPolicyAllOf) SetInitiatorStaticIpV4ConfigNil(b bool)`

 SetInitiatorStaticIpV4ConfigNil sets the value for InitiatorStaticIpV4Config to be an explicit nil

### UnsetInitiatorStaticIpV4Config
`func (o *VnicIscsiBootPolicyAllOf) UnsetInitiatorStaticIpV4Config()`

UnsetInitiatorStaticIpV4Config ensures that no value is present for InitiatorStaticIpV4Config, not even an explicit nil
### GetMutualChap

`func (o *VnicIscsiBootPolicyAllOf) GetMutualChap() VnicIscsiAuthProfile`

GetMutualChap returns the MutualChap field if non-nil, zero value otherwise.

### GetMutualChapOk

`func (o *VnicIscsiBootPolicyAllOf) GetMutualChapOk() (*VnicIscsiAuthProfile, bool)`

GetMutualChapOk returns a tuple with the MutualChap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutualChap

`func (o *VnicIscsiBootPolicyAllOf) SetMutualChap(v VnicIscsiAuthProfile)`

SetMutualChap sets MutualChap field to given value.

### HasMutualChap

`func (o *VnicIscsiBootPolicyAllOf) HasMutualChap() bool`

HasMutualChap returns a boolean if a field has been set.

### SetMutualChapNil

`func (o *VnicIscsiBootPolicyAllOf) SetMutualChapNil(b bool)`

 SetMutualChapNil sets the value for MutualChap to be an explicit nil

### UnsetMutualChap
`func (o *VnicIscsiBootPolicyAllOf) UnsetMutualChap()`

UnsetMutualChap ensures that no value is present for MutualChap, not even an explicit nil
### GetTargetSourceType

`func (o *VnicIscsiBootPolicyAllOf) GetTargetSourceType() string`

GetTargetSourceType returns the TargetSourceType field if non-nil, zero value otherwise.

### GetTargetSourceTypeOk

`func (o *VnicIscsiBootPolicyAllOf) GetTargetSourceTypeOk() (*string, bool)`

GetTargetSourceTypeOk returns a tuple with the TargetSourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSourceType

`func (o *VnicIscsiBootPolicyAllOf) SetTargetSourceType(v string)`

SetTargetSourceType sets TargetSourceType field to given value.

### HasTargetSourceType

`func (o *VnicIscsiBootPolicyAllOf) HasTargetSourceType() bool`

HasTargetSourceType returns a boolean if a field has been set.

### GetInitiatorIpPool

`func (o *VnicIscsiBootPolicyAllOf) GetInitiatorIpPool() IppoolPoolRelationship`

GetInitiatorIpPool returns the InitiatorIpPool field if non-nil, zero value otherwise.

### GetInitiatorIpPoolOk

`func (o *VnicIscsiBootPolicyAllOf) GetInitiatorIpPoolOk() (*IppoolPoolRelationship, bool)`

GetInitiatorIpPoolOk returns a tuple with the InitiatorIpPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorIpPool

`func (o *VnicIscsiBootPolicyAllOf) SetInitiatorIpPool(v IppoolPoolRelationship)`

SetInitiatorIpPool sets InitiatorIpPool field to given value.

### HasInitiatorIpPool

`func (o *VnicIscsiBootPolicyAllOf) HasInitiatorIpPool() bool`

HasInitiatorIpPool returns a boolean if a field has been set.

### GetIscsiAdapterPolicy

`func (o *VnicIscsiBootPolicyAllOf) GetIscsiAdapterPolicy() VnicIscsiAdapterPolicyRelationship`

GetIscsiAdapterPolicy returns the IscsiAdapterPolicy field if non-nil, zero value otherwise.

### GetIscsiAdapterPolicyOk

`func (o *VnicIscsiBootPolicyAllOf) GetIscsiAdapterPolicyOk() (*VnicIscsiAdapterPolicyRelationship, bool)`

GetIscsiAdapterPolicyOk returns a tuple with the IscsiAdapterPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiAdapterPolicy

`func (o *VnicIscsiBootPolicyAllOf) SetIscsiAdapterPolicy(v VnicIscsiAdapterPolicyRelationship)`

SetIscsiAdapterPolicy sets IscsiAdapterPolicy field to given value.

### HasIscsiAdapterPolicy

`func (o *VnicIscsiBootPolicyAllOf) HasIscsiAdapterPolicy() bool`

HasIscsiAdapterPolicy returns a boolean if a field has been set.

### GetOrganization

`func (o *VnicIscsiBootPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *VnicIscsiBootPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *VnicIscsiBootPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *VnicIscsiBootPolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetPrimaryTargetPolicy

`func (o *VnicIscsiBootPolicyAllOf) GetPrimaryTargetPolicy() VnicIscsiStaticTargetPolicyRelationship`

GetPrimaryTargetPolicy returns the PrimaryTargetPolicy field if non-nil, zero value otherwise.

### GetPrimaryTargetPolicyOk

`func (o *VnicIscsiBootPolicyAllOf) GetPrimaryTargetPolicyOk() (*VnicIscsiStaticTargetPolicyRelationship, bool)`

GetPrimaryTargetPolicyOk returns a tuple with the PrimaryTargetPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryTargetPolicy

`func (o *VnicIscsiBootPolicyAllOf) SetPrimaryTargetPolicy(v VnicIscsiStaticTargetPolicyRelationship)`

SetPrimaryTargetPolicy sets PrimaryTargetPolicy field to given value.

### HasPrimaryTargetPolicy

`func (o *VnicIscsiBootPolicyAllOf) HasPrimaryTargetPolicy() bool`

HasPrimaryTargetPolicy returns a boolean if a field has been set.

### GetSecondaryTargetPolicy

`func (o *VnicIscsiBootPolicyAllOf) GetSecondaryTargetPolicy() VnicIscsiStaticTargetPolicyRelationship`

GetSecondaryTargetPolicy returns the SecondaryTargetPolicy field if non-nil, zero value otherwise.

### GetSecondaryTargetPolicyOk

`func (o *VnicIscsiBootPolicyAllOf) GetSecondaryTargetPolicyOk() (*VnicIscsiStaticTargetPolicyRelationship, bool)`

GetSecondaryTargetPolicyOk returns a tuple with the SecondaryTargetPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryTargetPolicy

`func (o *VnicIscsiBootPolicyAllOf) SetSecondaryTargetPolicy(v VnicIscsiStaticTargetPolicyRelationship)`

SetSecondaryTargetPolicy sets SecondaryTargetPolicy field to given value.

### HasSecondaryTargetPolicy

`func (o *VnicIscsiBootPolicyAllOf) HasSecondaryTargetPolicy() bool`

HasSecondaryTargetPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


