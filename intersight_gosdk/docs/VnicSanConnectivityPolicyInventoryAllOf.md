# VnicSanConnectivityPolicyInventoryAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.SanConnectivityPolicyInventory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.SanConnectivityPolicyInventory"]
**PlacementMode** | Pointer to **string** | The mode used for placement of vNICs on network adapters. It can either be Auto or Custom. * &#x60;custom&#x60; - The placement of the vNICs / vHBAs on network adapters is manually chosen by the user. * &#x60;auto&#x60; - The placement of the vNICs / vHBAs on network adapters is automatically determined by the system. | [optional] [readonly] [default to "custom"]
**StaticWwnnAddress** | Pointer to **string** | The WWNN address for the server node must be in hexadecimal format xx:xx:xx:xx:xx:xx:xx:xx. Allowed ranges are 20:00:00:00:00:00:00:00 to 20:FF:FF:FF:FF:FF:FF:FF or from 50:00:00:00:00:00:00:00 to 5F:FF:FF:FF:FF:FF:FF:FF. To ensure uniqueness of WWN&#39;s in the SAN fabric, you are strongly encouraged to use the WWN prefix - 20:00:00:25:B5:xx:xx:xx. | [optional] [readonly] 
**TargetPlatform** | Pointer to **string** | The platform for which the server profile is applicable. It can either be a server that is operating in standalone mode or which is attached to a Fabric Interconnect managed by Intersight. * &#x60;Standalone&#x60; - Servers which are operating in standalone mode i.e. not connected to a Fabric Interconnected. * &#x60;FIAttached&#x60; - Servers which are connected to a Fabric Interconnect that is managed by Intersight. | [optional] [readonly] [default to "Standalone"]
**WwnnAddressType** | Pointer to **string** | Type of allocation selected to assign a WWNN address for the server node. * &#x60;POOL&#x60; - The user selects a pool from which the mac/wwn address will be leased for the Virtual Interface. * &#x60;STATIC&#x60; - The user assigns a static mac/wwn address for the Virtual Interface. | [optional] [readonly] [default to "POOL"]
**FcIfs** | Pointer to [**[]VnicFcIfInventoryRelationship**](VnicFcIfInventoryRelationship.md) | An array of relationships to vnicFcIfInventory resources. | [optional] [readonly] 
**TargetMo** | Pointer to [**MoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 
**WwnnPool** | Pointer to [**FcpoolPoolRelationship**](FcpoolPoolRelationship.md) |  | [optional] 

## Methods

### NewVnicSanConnectivityPolicyInventoryAllOf

`func NewVnicSanConnectivityPolicyInventoryAllOf(classId string, objectType string, ) *VnicSanConnectivityPolicyInventoryAllOf`

NewVnicSanConnectivityPolicyInventoryAllOf instantiates a new VnicSanConnectivityPolicyInventoryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicSanConnectivityPolicyInventoryAllOfWithDefaults

`func NewVnicSanConnectivityPolicyInventoryAllOfWithDefaults() *VnicSanConnectivityPolicyInventoryAllOf`

NewVnicSanConnectivityPolicyInventoryAllOfWithDefaults instantiates a new VnicSanConnectivityPolicyInventoryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicSanConnectivityPolicyInventoryAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicSanConnectivityPolicyInventoryAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicSanConnectivityPolicyInventoryAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicSanConnectivityPolicyInventoryAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicSanConnectivityPolicyInventoryAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicSanConnectivityPolicyInventoryAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPlacementMode

`func (o *VnicSanConnectivityPolicyInventoryAllOf) GetPlacementMode() string`

GetPlacementMode returns the PlacementMode field if non-nil, zero value otherwise.

### GetPlacementModeOk

`func (o *VnicSanConnectivityPolicyInventoryAllOf) GetPlacementModeOk() (*string, bool)`

GetPlacementModeOk returns a tuple with the PlacementMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementMode

`func (o *VnicSanConnectivityPolicyInventoryAllOf) SetPlacementMode(v string)`

SetPlacementMode sets PlacementMode field to given value.

### HasPlacementMode

`func (o *VnicSanConnectivityPolicyInventoryAllOf) HasPlacementMode() bool`

HasPlacementMode returns a boolean if a field has been set.

### GetStaticWwnnAddress

`func (o *VnicSanConnectivityPolicyInventoryAllOf) GetStaticWwnnAddress() string`

GetStaticWwnnAddress returns the StaticWwnnAddress field if non-nil, zero value otherwise.

### GetStaticWwnnAddressOk

`func (o *VnicSanConnectivityPolicyInventoryAllOf) GetStaticWwnnAddressOk() (*string, bool)`

GetStaticWwnnAddressOk returns a tuple with the StaticWwnnAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticWwnnAddress

`func (o *VnicSanConnectivityPolicyInventoryAllOf) SetStaticWwnnAddress(v string)`

SetStaticWwnnAddress sets StaticWwnnAddress field to given value.

### HasStaticWwnnAddress

`func (o *VnicSanConnectivityPolicyInventoryAllOf) HasStaticWwnnAddress() bool`

HasStaticWwnnAddress returns a boolean if a field has been set.

### GetTargetPlatform

`func (o *VnicSanConnectivityPolicyInventoryAllOf) GetTargetPlatform() string`

GetTargetPlatform returns the TargetPlatform field if non-nil, zero value otherwise.

### GetTargetPlatformOk

`func (o *VnicSanConnectivityPolicyInventoryAllOf) GetTargetPlatformOk() (*string, bool)`

GetTargetPlatformOk returns a tuple with the TargetPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPlatform

`func (o *VnicSanConnectivityPolicyInventoryAllOf) SetTargetPlatform(v string)`

SetTargetPlatform sets TargetPlatform field to given value.

### HasTargetPlatform

`func (o *VnicSanConnectivityPolicyInventoryAllOf) HasTargetPlatform() bool`

HasTargetPlatform returns a boolean if a field has been set.

### GetWwnnAddressType

`func (o *VnicSanConnectivityPolicyInventoryAllOf) GetWwnnAddressType() string`

GetWwnnAddressType returns the WwnnAddressType field if non-nil, zero value otherwise.

### GetWwnnAddressTypeOk

`func (o *VnicSanConnectivityPolicyInventoryAllOf) GetWwnnAddressTypeOk() (*string, bool)`

GetWwnnAddressTypeOk returns a tuple with the WwnnAddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwnnAddressType

`func (o *VnicSanConnectivityPolicyInventoryAllOf) SetWwnnAddressType(v string)`

SetWwnnAddressType sets WwnnAddressType field to given value.

### HasWwnnAddressType

`func (o *VnicSanConnectivityPolicyInventoryAllOf) HasWwnnAddressType() bool`

HasWwnnAddressType returns a boolean if a field has been set.

### GetFcIfs

`func (o *VnicSanConnectivityPolicyInventoryAllOf) GetFcIfs() []VnicFcIfInventoryRelationship`

GetFcIfs returns the FcIfs field if non-nil, zero value otherwise.

### GetFcIfsOk

`func (o *VnicSanConnectivityPolicyInventoryAllOf) GetFcIfsOk() (*[]VnicFcIfInventoryRelationship, bool)`

GetFcIfsOk returns a tuple with the FcIfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcIfs

`func (o *VnicSanConnectivityPolicyInventoryAllOf) SetFcIfs(v []VnicFcIfInventoryRelationship)`

SetFcIfs sets FcIfs field to given value.

### HasFcIfs

`func (o *VnicSanConnectivityPolicyInventoryAllOf) HasFcIfs() bool`

HasFcIfs returns a boolean if a field has been set.

### SetFcIfsNil

`func (o *VnicSanConnectivityPolicyInventoryAllOf) SetFcIfsNil(b bool)`

 SetFcIfsNil sets the value for FcIfs to be an explicit nil

### UnsetFcIfs
`func (o *VnicSanConnectivityPolicyInventoryAllOf) UnsetFcIfs()`

UnsetFcIfs ensures that no value is present for FcIfs, not even an explicit nil
### GetTargetMo

`func (o *VnicSanConnectivityPolicyInventoryAllOf) GetTargetMo() MoBaseMoRelationship`

GetTargetMo returns the TargetMo field if non-nil, zero value otherwise.

### GetTargetMoOk

`func (o *VnicSanConnectivityPolicyInventoryAllOf) GetTargetMoOk() (*MoBaseMoRelationship, bool)`

GetTargetMoOk returns a tuple with the TargetMo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMo

`func (o *VnicSanConnectivityPolicyInventoryAllOf) SetTargetMo(v MoBaseMoRelationship)`

SetTargetMo sets TargetMo field to given value.

### HasTargetMo

`func (o *VnicSanConnectivityPolicyInventoryAllOf) HasTargetMo() bool`

HasTargetMo returns a boolean if a field has been set.

### GetWwnnPool

`func (o *VnicSanConnectivityPolicyInventoryAllOf) GetWwnnPool() FcpoolPoolRelationship`

GetWwnnPool returns the WwnnPool field if non-nil, zero value otherwise.

### GetWwnnPoolOk

`func (o *VnicSanConnectivityPolicyInventoryAllOf) GetWwnnPoolOk() (*FcpoolPoolRelationship, bool)`

GetWwnnPoolOk returns a tuple with the WwnnPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwnnPool

`func (o *VnicSanConnectivityPolicyInventoryAllOf) SetWwnnPool(v FcpoolPoolRelationship)`

SetWwnnPool sets WwnnPool field to given value.

### HasWwnnPool

`func (o *VnicSanConnectivityPolicyInventoryAllOf) HasWwnnPool() bool`

HasWwnnPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


