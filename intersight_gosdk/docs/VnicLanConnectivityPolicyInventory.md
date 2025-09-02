# VnicLanConnectivityPolicyInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.LanConnectivityPolicyInventory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.LanConnectivityPolicyInventory"]
**AzureQosEnabled** | Pointer to **bool** | Enabling AzureStack-Host QoS on an adapter allows the user to carve out traffic classes for RDMA traffic which ensures that a desired portion of the bandwidth is allocated to it. | [optional] [readonly] [default to false]
**IqnAllocationType** | Pointer to **string** | Allocation Type of iSCSI Qualified Name - Static/Pool/None. * &#x60;None&#x60; - Type indicates that there is no IQN associated to an interface. * &#x60;Static&#x60; - Type represents that static IQN is associated to an interface. * &#x60;Pool&#x60; - Type indicates that IQN value is sourced from an associated pool. | [optional] [readonly] [default to "None"]
**PlacementMode** | Pointer to **string** | The mode used for placement of vNICs on network adapters. It can either be Auto or Custom. * &#x60;custom&#x60; - The placement of the vNICs / vHBAs on network adapters is manually chosen by the user. * &#x60;auto&#x60; - The placement of the vNICs / vHBAs on network adapters is automatically determined by the system. | [optional] [readonly] [default to "custom"]
**StaticIqnName** | Pointer to **string** | User provided static iSCSI Qualified Name (IQN) for use as initiator identifiers by iSCSI vNICs in a Fabric Interconnect domain. | [optional] [readonly] 
**TargetPlatform** | Pointer to **string** | The platform for which the server profile is applicable. It can either be a server that is operating in standalone mode or which is attached to a Fabric Interconnect managed by Intersight. * &#x60;Standalone&#x60; - Servers which are operating in standalone mode i.e. not connected to a Fabric Interconnected. * &#x60;FIAttached&#x60; - Servers which are connected to a Fabric Interconnect that is managed by Intersight. * &#x60;UnifiedEdgeServer&#x60; - Unified Edge sleds that is managed by Intersight. | [optional] [readonly] [default to "Standalone"]
**EthIfs** | Pointer to [**[]VnicEthIfInventoryRelationship**](VnicEthIfInventoryRelationship.md) | An array of relationships to vnicEthIfInventory resources. | [optional] [readonly] 
**IqnPool** | Pointer to [**NullableIqnpoolPoolRelationship**](IqnpoolPoolRelationship.md) |  | [optional] 
**TargetMo** | Pointer to [**NullableMoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewVnicLanConnectivityPolicyInventory

`func NewVnicLanConnectivityPolicyInventory(classId string, objectType string, ) *VnicLanConnectivityPolicyInventory`

NewVnicLanConnectivityPolicyInventory instantiates a new VnicLanConnectivityPolicyInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicLanConnectivityPolicyInventoryWithDefaults

`func NewVnicLanConnectivityPolicyInventoryWithDefaults() *VnicLanConnectivityPolicyInventory`

NewVnicLanConnectivityPolicyInventoryWithDefaults instantiates a new VnicLanConnectivityPolicyInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicLanConnectivityPolicyInventory) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicLanConnectivityPolicyInventory) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicLanConnectivityPolicyInventory) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicLanConnectivityPolicyInventory) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicLanConnectivityPolicyInventory) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicLanConnectivityPolicyInventory) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAzureQosEnabled

`func (o *VnicLanConnectivityPolicyInventory) GetAzureQosEnabled() bool`

GetAzureQosEnabled returns the AzureQosEnabled field if non-nil, zero value otherwise.

### GetAzureQosEnabledOk

`func (o *VnicLanConnectivityPolicyInventory) GetAzureQosEnabledOk() (*bool, bool)`

GetAzureQosEnabledOk returns a tuple with the AzureQosEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureQosEnabled

`func (o *VnicLanConnectivityPolicyInventory) SetAzureQosEnabled(v bool)`

SetAzureQosEnabled sets AzureQosEnabled field to given value.

### HasAzureQosEnabled

`func (o *VnicLanConnectivityPolicyInventory) HasAzureQosEnabled() bool`

HasAzureQosEnabled returns a boolean if a field has been set.

### GetIqnAllocationType

`func (o *VnicLanConnectivityPolicyInventory) GetIqnAllocationType() string`

GetIqnAllocationType returns the IqnAllocationType field if non-nil, zero value otherwise.

### GetIqnAllocationTypeOk

`func (o *VnicLanConnectivityPolicyInventory) GetIqnAllocationTypeOk() (*string, bool)`

GetIqnAllocationTypeOk returns a tuple with the IqnAllocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIqnAllocationType

`func (o *VnicLanConnectivityPolicyInventory) SetIqnAllocationType(v string)`

SetIqnAllocationType sets IqnAllocationType field to given value.

### HasIqnAllocationType

`func (o *VnicLanConnectivityPolicyInventory) HasIqnAllocationType() bool`

HasIqnAllocationType returns a boolean if a field has been set.

### GetPlacementMode

`func (o *VnicLanConnectivityPolicyInventory) GetPlacementMode() string`

GetPlacementMode returns the PlacementMode field if non-nil, zero value otherwise.

### GetPlacementModeOk

`func (o *VnicLanConnectivityPolicyInventory) GetPlacementModeOk() (*string, bool)`

GetPlacementModeOk returns a tuple with the PlacementMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementMode

`func (o *VnicLanConnectivityPolicyInventory) SetPlacementMode(v string)`

SetPlacementMode sets PlacementMode field to given value.

### HasPlacementMode

`func (o *VnicLanConnectivityPolicyInventory) HasPlacementMode() bool`

HasPlacementMode returns a boolean if a field has been set.

### GetStaticIqnName

`func (o *VnicLanConnectivityPolicyInventory) GetStaticIqnName() string`

GetStaticIqnName returns the StaticIqnName field if non-nil, zero value otherwise.

### GetStaticIqnNameOk

`func (o *VnicLanConnectivityPolicyInventory) GetStaticIqnNameOk() (*string, bool)`

GetStaticIqnNameOk returns a tuple with the StaticIqnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticIqnName

`func (o *VnicLanConnectivityPolicyInventory) SetStaticIqnName(v string)`

SetStaticIqnName sets StaticIqnName field to given value.

### HasStaticIqnName

`func (o *VnicLanConnectivityPolicyInventory) HasStaticIqnName() bool`

HasStaticIqnName returns a boolean if a field has been set.

### GetTargetPlatform

`func (o *VnicLanConnectivityPolicyInventory) GetTargetPlatform() string`

GetTargetPlatform returns the TargetPlatform field if non-nil, zero value otherwise.

### GetTargetPlatformOk

`func (o *VnicLanConnectivityPolicyInventory) GetTargetPlatformOk() (*string, bool)`

GetTargetPlatformOk returns a tuple with the TargetPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPlatform

`func (o *VnicLanConnectivityPolicyInventory) SetTargetPlatform(v string)`

SetTargetPlatform sets TargetPlatform field to given value.

### HasTargetPlatform

`func (o *VnicLanConnectivityPolicyInventory) HasTargetPlatform() bool`

HasTargetPlatform returns a boolean if a field has been set.

### GetEthIfs

`func (o *VnicLanConnectivityPolicyInventory) GetEthIfs() []VnicEthIfInventoryRelationship`

GetEthIfs returns the EthIfs field if non-nil, zero value otherwise.

### GetEthIfsOk

`func (o *VnicLanConnectivityPolicyInventory) GetEthIfsOk() (*[]VnicEthIfInventoryRelationship, bool)`

GetEthIfsOk returns a tuple with the EthIfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthIfs

`func (o *VnicLanConnectivityPolicyInventory) SetEthIfs(v []VnicEthIfInventoryRelationship)`

SetEthIfs sets EthIfs field to given value.

### HasEthIfs

`func (o *VnicLanConnectivityPolicyInventory) HasEthIfs() bool`

HasEthIfs returns a boolean if a field has been set.

### SetEthIfsNil

`func (o *VnicLanConnectivityPolicyInventory) SetEthIfsNil(b bool)`

 SetEthIfsNil sets the value for EthIfs to be an explicit nil

### UnsetEthIfs
`func (o *VnicLanConnectivityPolicyInventory) UnsetEthIfs()`

UnsetEthIfs ensures that no value is present for EthIfs, not even an explicit nil
### GetIqnPool

`func (o *VnicLanConnectivityPolicyInventory) GetIqnPool() IqnpoolPoolRelationship`

GetIqnPool returns the IqnPool field if non-nil, zero value otherwise.

### GetIqnPoolOk

`func (o *VnicLanConnectivityPolicyInventory) GetIqnPoolOk() (*IqnpoolPoolRelationship, bool)`

GetIqnPoolOk returns a tuple with the IqnPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIqnPool

`func (o *VnicLanConnectivityPolicyInventory) SetIqnPool(v IqnpoolPoolRelationship)`

SetIqnPool sets IqnPool field to given value.

### HasIqnPool

`func (o *VnicLanConnectivityPolicyInventory) HasIqnPool() bool`

HasIqnPool returns a boolean if a field has been set.

### SetIqnPoolNil

`func (o *VnicLanConnectivityPolicyInventory) SetIqnPoolNil(b bool)`

 SetIqnPoolNil sets the value for IqnPool to be an explicit nil

### UnsetIqnPool
`func (o *VnicLanConnectivityPolicyInventory) UnsetIqnPool()`

UnsetIqnPool ensures that no value is present for IqnPool, not even an explicit nil
### GetTargetMo

`func (o *VnicLanConnectivityPolicyInventory) GetTargetMo() MoBaseMoRelationship`

GetTargetMo returns the TargetMo field if non-nil, zero value otherwise.

### GetTargetMoOk

`func (o *VnicLanConnectivityPolicyInventory) GetTargetMoOk() (*MoBaseMoRelationship, bool)`

GetTargetMoOk returns a tuple with the TargetMo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMo

`func (o *VnicLanConnectivityPolicyInventory) SetTargetMo(v MoBaseMoRelationship)`

SetTargetMo sets TargetMo field to given value.

### HasTargetMo

`func (o *VnicLanConnectivityPolicyInventory) HasTargetMo() bool`

HasTargetMo returns a boolean if a field has been set.

### SetTargetMoNil

`func (o *VnicLanConnectivityPolicyInventory) SetTargetMoNil(b bool)`

 SetTargetMoNil sets the value for TargetMo to be an explicit nil

### UnsetTargetMo
`func (o *VnicLanConnectivityPolicyInventory) UnsetTargetMo()`

UnsetTargetMo ensures that no value is present for TargetMo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


