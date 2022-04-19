# VnicIscsiBootPolicyInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.IscsiBootPolicyInventory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.IscsiBootPolicyInventory"]
**AutoTargetvendorName** | Pointer to **string** | Auto target interface that is represented via the Initiator name or the DHCP vendor ID. The vendor ID can be up to 32 alphanumeric characters. | [optional] [readonly] 
**Chap** | Pointer to [**NullableVnicIscsiAuthProfile**](VnicIscsiAuthProfile.md) |  | [optional] 
**InitiatorIpSource** | Pointer to **string** | Source Type of Initiator IP Address - Auto/Static/Pool. * &#x60;DHCP&#x60; - The IP address is assigned using DHCP, if available. * &#x60;Static&#x60; - Static IPv4 address is assigned to the iSCSI boot interface based on the information entered in this area. * &#x60;Pool&#x60; - An IPv4 address is assigned to the iSCSI boot interface from the management IP address pool. | [optional] [readonly] [default to "DHCP"]
**InitiatorStaticIpV4Address** | Pointer to **string** | Static IP address provided for iSCSI Initiator. | [optional] [readonly] 
**InitiatorStaticIpV4Config** | Pointer to [**NullableIppoolIpV4Config**](IppoolIpV4Config.md) |  | [optional] 
**MutualChap** | Pointer to [**NullableVnicIscsiAuthProfile**](VnicIscsiAuthProfile.md) |  | [optional] 
**TargetSourceType** | Pointer to **string** | Source Type of Targets - Auto/Static. * &#x60;Static&#x60; - Type indicates that static target interface is assigned to iSCSI boot. * &#x60;Auto&#x60; - Type indicates that the system selects the target interface automatically during iSCSI boot. | [optional] [readonly] [default to "Static"]
**InitiatorIpPool** | Pointer to [**IppoolPoolRelationship**](IppoolPoolRelationship.md) |  | [optional] 
**IscsiAdapterPolicy** | Pointer to [**VnicIscsiAdapterPolicyInventoryRelationship**](VnicIscsiAdapterPolicyInventoryRelationship.md) |  | [optional] 
**PrimaryTargetPolicy** | Pointer to [**VnicIscsiStaticTargetPolicyInventoryRelationship**](VnicIscsiStaticTargetPolicyInventoryRelationship.md) |  | [optional] 
**SecondaryTargetPolicy** | Pointer to [**VnicIscsiStaticTargetPolicyInventoryRelationship**](VnicIscsiStaticTargetPolicyInventoryRelationship.md) |  | [optional] 
**TargetMo** | Pointer to [**MoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewVnicIscsiBootPolicyInventory

`func NewVnicIscsiBootPolicyInventory(classId string, objectType string, ) *VnicIscsiBootPolicyInventory`

NewVnicIscsiBootPolicyInventory instantiates a new VnicIscsiBootPolicyInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicIscsiBootPolicyInventoryWithDefaults

`func NewVnicIscsiBootPolicyInventoryWithDefaults() *VnicIscsiBootPolicyInventory`

NewVnicIscsiBootPolicyInventoryWithDefaults instantiates a new VnicIscsiBootPolicyInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicIscsiBootPolicyInventory) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicIscsiBootPolicyInventory) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicIscsiBootPolicyInventory) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicIscsiBootPolicyInventory) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicIscsiBootPolicyInventory) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicIscsiBootPolicyInventory) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAutoTargetvendorName

`func (o *VnicIscsiBootPolicyInventory) GetAutoTargetvendorName() string`

GetAutoTargetvendorName returns the AutoTargetvendorName field if non-nil, zero value otherwise.

### GetAutoTargetvendorNameOk

`func (o *VnicIscsiBootPolicyInventory) GetAutoTargetvendorNameOk() (*string, bool)`

GetAutoTargetvendorNameOk returns a tuple with the AutoTargetvendorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTargetvendorName

`func (o *VnicIscsiBootPolicyInventory) SetAutoTargetvendorName(v string)`

SetAutoTargetvendorName sets AutoTargetvendorName field to given value.

### HasAutoTargetvendorName

`func (o *VnicIscsiBootPolicyInventory) HasAutoTargetvendorName() bool`

HasAutoTargetvendorName returns a boolean if a field has been set.

### GetChap

`func (o *VnicIscsiBootPolicyInventory) GetChap() VnicIscsiAuthProfile`

GetChap returns the Chap field if non-nil, zero value otherwise.

### GetChapOk

`func (o *VnicIscsiBootPolicyInventory) GetChapOk() (*VnicIscsiAuthProfile, bool)`

GetChapOk returns a tuple with the Chap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChap

`func (o *VnicIscsiBootPolicyInventory) SetChap(v VnicIscsiAuthProfile)`

SetChap sets Chap field to given value.

### HasChap

`func (o *VnicIscsiBootPolicyInventory) HasChap() bool`

HasChap returns a boolean if a field has been set.

### SetChapNil

`func (o *VnicIscsiBootPolicyInventory) SetChapNil(b bool)`

 SetChapNil sets the value for Chap to be an explicit nil

### UnsetChap
`func (o *VnicIscsiBootPolicyInventory) UnsetChap()`

UnsetChap ensures that no value is present for Chap, not even an explicit nil
### GetInitiatorIpSource

`func (o *VnicIscsiBootPolicyInventory) GetInitiatorIpSource() string`

GetInitiatorIpSource returns the InitiatorIpSource field if non-nil, zero value otherwise.

### GetInitiatorIpSourceOk

`func (o *VnicIscsiBootPolicyInventory) GetInitiatorIpSourceOk() (*string, bool)`

GetInitiatorIpSourceOk returns a tuple with the InitiatorIpSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorIpSource

`func (o *VnicIscsiBootPolicyInventory) SetInitiatorIpSource(v string)`

SetInitiatorIpSource sets InitiatorIpSource field to given value.

### HasInitiatorIpSource

`func (o *VnicIscsiBootPolicyInventory) HasInitiatorIpSource() bool`

HasInitiatorIpSource returns a boolean if a field has been set.

### GetInitiatorStaticIpV4Address

`func (o *VnicIscsiBootPolicyInventory) GetInitiatorStaticIpV4Address() string`

GetInitiatorStaticIpV4Address returns the InitiatorStaticIpV4Address field if non-nil, zero value otherwise.

### GetInitiatorStaticIpV4AddressOk

`func (o *VnicIscsiBootPolicyInventory) GetInitiatorStaticIpV4AddressOk() (*string, bool)`

GetInitiatorStaticIpV4AddressOk returns a tuple with the InitiatorStaticIpV4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorStaticIpV4Address

`func (o *VnicIscsiBootPolicyInventory) SetInitiatorStaticIpV4Address(v string)`

SetInitiatorStaticIpV4Address sets InitiatorStaticIpV4Address field to given value.

### HasInitiatorStaticIpV4Address

`func (o *VnicIscsiBootPolicyInventory) HasInitiatorStaticIpV4Address() bool`

HasInitiatorStaticIpV4Address returns a boolean if a field has been set.

### GetInitiatorStaticIpV4Config

`func (o *VnicIscsiBootPolicyInventory) GetInitiatorStaticIpV4Config() IppoolIpV4Config`

GetInitiatorStaticIpV4Config returns the InitiatorStaticIpV4Config field if non-nil, zero value otherwise.

### GetInitiatorStaticIpV4ConfigOk

`func (o *VnicIscsiBootPolicyInventory) GetInitiatorStaticIpV4ConfigOk() (*IppoolIpV4Config, bool)`

GetInitiatorStaticIpV4ConfigOk returns a tuple with the InitiatorStaticIpV4Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorStaticIpV4Config

`func (o *VnicIscsiBootPolicyInventory) SetInitiatorStaticIpV4Config(v IppoolIpV4Config)`

SetInitiatorStaticIpV4Config sets InitiatorStaticIpV4Config field to given value.

### HasInitiatorStaticIpV4Config

`func (o *VnicIscsiBootPolicyInventory) HasInitiatorStaticIpV4Config() bool`

HasInitiatorStaticIpV4Config returns a boolean if a field has been set.

### SetInitiatorStaticIpV4ConfigNil

`func (o *VnicIscsiBootPolicyInventory) SetInitiatorStaticIpV4ConfigNil(b bool)`

 SetInitiatorStaticIpV4ConfigNil sets the value for InitiatorStaticIpV4Config to be an explicit nil

### UnsetInitiatorStaticIpV4Config
`func (o *VnicIscsiBootPolicyInventory) UnsetInitiatorStaticIpV4Config()`

UnsetInitiatorStaticIpV4Config ensures that no value is present for InitiatorStaticIpV4Config, not even an explicit nil
### GetMutualChap

`func (o *VnicIscsiBootPolicyInventory) GetMutualChap() VnicIscsiAuthProfile`

GetMutualChap returns the MutualChap field if non-nil, zero value otherwise.

### GetMutualChapOk

`func (o *VnicIscsiBootPolicyInventory) GetMutualChapOk() (*VnicIscsiAuthProfile, bool)`

GetMutualChapOk returns a tuple with the MutualChap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutualChap

`func (o *VnicIscsiBootPolicyInventory) SetMutualChap(v VnicIscsiAuthProfile)`

SetMutualChap sets MutualChap field to given value.

### HasMutualChap

`func (o *VnicIscsiBootPolicyInventory) HasMutualChap() bool`

HasMutualChap returns a boolean if a field has been set.

### SetMutualChapNil

`func (o *VnicIscsiBootPolicyInventory) SetMutualChapNil(b bool)`

 SetMutualChapNil sets the value for MutualChap to be an explicit nil

### UnsetMutualChap
`func (o *VnicIscsiBootPolicyInventory) UnsetMutualChap()`

UnsetMutualChap ensures that no value is present for MutualChap, not even an explicit nil
### GetTargetSourceType

`func (o *VnicIscsiBootPolicyInventory) GetTargetSourceType() string`

GetTargetSourceType returns the TargetSourceType field if non-nil, zero value otherwise.

### GetTargetSourceTypeOk

`func (o *VnicIscsiBootPolicyInventory) GetTargetSourceTypeOk() (*string, bool)`

GetTargetSourceTypeOk returns a tuple with the TargetSourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSourceType

`func (o *VnicIscsiBootPolicyInventory) SetTargetSourceType(v string)`

SetTargetSourceType sets TargetSourceType field to given value.

### HasTargetSourceType

`func (o *VnicIscsiBootPolicyInventory) HasTargetSourceType() bool`

HasTargetSourceType returns a boolean if a field has been set.

### GetInitiatorIpPool

`func (o *VnicIscsiBootPolicyInventory) GetInitiatorIpPool() IppoolPoolRelationship`

GetInitiatorIpPool returns the InitiatorIpPool field if non-nil, zero value otherwise.

### GetInitiatorIpPoolOk

`func (o *VnicIscsiBootPolicyInventory) GetInitiatorIpPoolOk() (*IppoolPoolRelationship, bool)`

GetInitiatorIpPoolOk returns a tuple with the InitiatorIpPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorIpPool

`func (o *VnicIscsiBootPolicyInventory) SetInitiatorIpPool(v IppoolPoolRelationship)`

SetInitiatorIpPool sets InitiatorIpPool field to given value.

### HasInitiatorIpPool

`func (o *VnicIscsiBootPolicyInventory) HasInitiatorIpPool() bool`

HasInitiatorIpPool returns a boolean if a field has been set.

### GetIscsiAdapterPolicy

`func (o *VnicIscsiBootPolicyInventory) GetIscsiAdapterPolicy() VnicIscsiAdapterPolicyInventoryRelationship`

GetIscsiAdapterPolicy returns the IscsiAdapterPolicy field if non-nil, zero value otherwise.

### GetIscsiAdapterPolicyOk

`func (o *VnicIscsiBootPolicyInventory) GetIscsiAdapterPolicyOk() (*VnicIscsiAdapterPolicyInventoryRelationship, bool)`

GetIscsiAdapterPolicyOk returns a tuple with the IscsiAdapterPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiAdapterPolicy

`func (o *VnicIscsiBootPolicyInventory) SetIscsiAdapterPolicy(v VnicIscsiAdapterPolicyInventoryRelationship)`

SetIscsiAdapterPolicy sets IscsiAdapterPolicy field to given value.

### HasIscsiAdapterPolicy

`func (o *VnicIscsiBootPolicyInventory) HasIscsiAdapterPolicy() bool`

HasIscsiAdapterPolicy returns a boolean if a field has been set.

### GetPrimaryTargetPolicy

`func (o *VnicIscsiBootPolicyInventory) GetPrimaryTargetPolicy() VnicIscsiStaticTargetPolicyInventoryRelationship`

GetPrimaryTargetPolicy returns the PrimaryTargetPolicy field if non-nil, zero value otherwise.

### GetPrimaryTargetPolicyOk

`func (o *VnicIscsiBootPolicyInventory) GetPrimaryTargetPolicyOk() (*VnicIscsiStaticTargetPolicyInventoryRelationship, bool)`

GetPrimaryTargetPolicyOk returns a tuple with the PrimaryTargetPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryTargetPolicy

`func (o *VnicIscsiBootPolicyInventory) SetPrimaryTargetPolicy(v VnicIscsiStaticTargetPolicyInventoryRelationship)`

SetPrimaryTargetPolicy sets PrimaryTargetPolicy field to given value.

### HasPrimaryTargetPolicy

`func (o *VnicIscsiBootPolicyInventory) HasPrimaryTargetPolicy() bool`

HasPrimaryTargetPolicy returns a boolean if a field has been set.

### GetSecondaryTargetPolicy

`func (o *VnicIscsiBootPolicyInventory) GetSecondaryTargetPolicy() VnicIscsiStaticTargetPolicyInventoryRelationship`

GetSecondaryTargetPolicy returns the SecondaryTargetPolicy field if non-nil, zero value otherwise.

### GetSecondaryTargetPolicyOk

`func (o *VnicIscsiBootPolicyInventory) GetSecondaryTargetPolicyOk() (*VnicIscsiStaticTargetPolicyInventoryRelationship, bool)`

GetSecondaryTargetPolicyOk returns a tuple with the SecondaryTargetPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryTargetPolicy

`func (o *VnicIscsiBootPolicyInventory) SetSecondaryTargetPolicy(v VnicIscsiStaticTargetPolicyInventoryRelationship)`

SetSecondaryTargetPolicy sets SecondaryTargetPolicy field to given value.

### HasSecondaryTargetPolicy

`func (o *VnicIscsiBootPolicyInventory) HasSecondaryTargetPolicy() bool`

HasSecondaryTargetPolicy returns a boolean if a field has been set.

### GetTargetMo

`func (o *VnicIscsiBootPolicyInventory) GetTargetMo() MoBaseMoRelationship`

GetTargetMo returns the TargetMo field if non-nil, zero value otherwise.

### GetTargetMoOk

`func (o *VnicIscsiBootPolicyInventory) GetTargetMoOk() (*MoBaseMoRelationship, bool)`

GetTargetMoOk returns a tuple with the TargetMo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMo

`func (o *VnicIscsiBootPolicyInventory) SetTargetMo(v MoBaseMoRelationship)`

SetTargetMo sets TargetMo field to given value.

### HasTargetMo

`func (o *VnicIscsiBootPolicyInventory) HasTargetMo() bool`

HasTargetMo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


