# VnicBaseFcIf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**PersistentBindings** | Pointer to **bool** | Enables retention of LUN ID associations in memory until they are manually cleared. | [optional] 
**PinGroupName** | Pointer to **string** | Pingroup name associated to vfc for static pinning. SCP deploy will resolve pingroup name and fetches the correspoding uplink port/port channel to pin the vfc traffic. | [optional] 
**Type** | Pointer to **string** | VHBA Type configuration for SAN Connectivity Policy. This configuration is supported only on Cisco VIC 14XX series and higher series of adapters. * &#x60;fc-initiator&#x60; - The default value set for vHBA Type Configuration. Fc-initiator specifies vHBA as a consumer of storage. Enables SCSI commands to transfer data and status information between host and target storage systems. * &#x60;fc-nvme-initiator&#x60; - Fc-nvme-initiator specifies vHBA as a consumer of storage. Enables NVMe-based message commands to transfer data and status information between host and target storage systems. * &#x60;fc-nvme-target&#x60; - Fc-nvme-target specifies vHBA as a provider of storage volumes to initiators. Enables NVMe-based message commands to transfer data and status information between host and target storage systems. Currently tech-preview, only enabled with an asynchronous driver. * &#x60;fc-target&#x60; - Fc-target specifies vHBA as a provider of storage volumes to initiators. Enables SCSI commands to transfer data and status information between host and target storage systems. fc-target is enabled only with an asynchronous driver. | [optional] [default to "fc-initiator"]
**FcAdapterPolicy** | Pointer to [**NullableVnicFcAdapterPolicyRelationship**](VnicFcAdapterPolicyRelationship.md) |  | [optional] 
**FcNetworkPolicy** | Pointer to [**NullableVnicFcNetworkPolicyRelationship**](VnicFcNetworkPolicyRelationship.md) |  | [optional] 
**FcQosPolicy** | Pointer to [**NullableVnicFcQosPolicyRelationship**](VnicFcQosPolicyRelationship.md) |  | [optional] 
**FcZonePolicies** | Pointer to [**[]FabricFcZonePolicyRelationship**](FabricFcZonePolicyRelationship.md) | An array of relationships to fabricFcZonePolicy resources. | [optional] 
**WwpnPool** | Pointer to [**NullableFcpoolPoolRelationship**](FcpoolPoolRelationship.md) |  | [optional] 

## Methods

### NewVnicBaseFcIf

`func NewVnicBaseFcIf(classId string, objectType string, ) *VnicBaseFcIf`

NewVnicBaseFcIf instantiates a new VnicBaseFcIf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicBaseFcIfWithDefaults

`func NewVnicBaseFcIfWithDefaults() *VnicBaseFcIf`

NewVnicBaseFcIfWithDefaults instantiates a new VnicBaseFcIf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicBaseFcIf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicBaseFcIf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicBaseFcIf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicBaseFcIf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicBaseFcIf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicBaseFcIf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPersistentBindings

`func (o *VnicBaseFcIf) GetPersistentBindings() bool`

GetPersistentBindings returns the PersistentBindings field if non-nil, zero value otherwise.

### GetPersistentBindingsOk

`func (o *VnicBaseFcIf) GetPersistentBindingsOk() (*bool, bool)`

GetPersistentBindingsOk returns a tuple with the PersistentBindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentBindings

`func (o *VnicBaseFcIf) SetPersistentBindings(v bool)`

SetPersistentBindings sets PersistentBindings field to given value.

### HasPersistentBindings

`func (o *VnicBaseFcIf) HasPersistentBindings() bool`

HasPersistentBindings returns a boolean if a field has been set.

### GetPinGroupName

`func (o *VnicBaseFcIf) GetPinGroupName() string`

GetPinGroupName returns the PinGroupName field if non-nil, zero value otherwise.

### GetPinGroupNameOk

`func (o *VnicBaseFcIf) GetPinGroupNameOk() (*string, bool)`

GetPinGroupNameOk returns a tuple with the PinGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinGroupName

`func (o *VnicBaseFcIf) SetPinGroupName(v string)`

SetPinGroupName sets PinGroupName field to given value.

### HasPinGroupName

`func (o *VnicBaseFcIf) HasPinGroupName() bool`

HasPinGroupName returns a boolean if a field has been set.

### GetType

`func (o *VnicBaseFcIf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VnicBaseFcIf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VnicBaseFcIf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VnicBaseFcIf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFcAdapterPolicy

`func (o *VnicBaseFcIf) GetFcAdapterPolicy() VnicFcAdapterPolicyRelationship`

GetFcAdapterPolicy returns the FcAdapterPolicy field if non-nil, zero value otherwise.

### GetFcAdapterPolicyOk

`func (o *VnicBaseFcIf) GetFcAdapterPolicyOk() (*VnicFcAdapterPolicyRelationship, bool)`

GetFcAdapterPolicyOk returns a tuple with the FcAdapterPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcAdapterPolicy

`func (o *VnicBaseFcIf) SetFcAdapterPolicy(v VnicFcAdapterPolicyRelationship)`

SetFcAdapterPolicy sets FcAdapterPolicy field to given value.

### HasFcAdapterPolicy

`func (o *VnicBaseFcIf) HasFcAdapterPolicy() bool`

HasFcAdapterPolicy returns a boolean if a field has been set.

### SetFcAdapterPolicyNil

`func (o *VnicBaseFcIf) SetFcAdapterPolicyNil(b bool)`

 SetFcAdapterPolicyNil sets the value for FcAdapterPolicy to be an explicit nil

### UnsetFcAdapterPolicy
`func (o *VnicBaseFcIf) UnsetFcAdapterPolicy()`

UnsetFcAdapterPolicy ensures that no value is present for FcAdapterPolicy, not even an explicit nil
### GetFcNetworkPolicy

`func (o *VnicBaseFcIf) GetFcNetworkPolicy() VnicFcNetworkPolicyRelationship`

GetFcNetworkPolicy returns the FcNetworkPolicy field if non-nil, zero value otherwise.

### GetFcNetworkPolicyOk

`func (o *VnicBaseFcIf) GetFcNetworkPolicyOk() (*VnicFcNetworkPolicyRelationship, bool)`

GetFcNetworkPolicyOk returns a tuple with the FcNetworkPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcNetworkPolicy

`func (o *VnicBaseFcIf) SetFcNetworkPolicy(v VnicFcNetworkPolicyRelationship)`

SetFcNetworkPolicy sets FcNetworkPolicy field to given value.

### HasFcNetworkPolicy

`func (o *VnicBaseFcIf) HasFcNetworkPolicy() bool`

HasFcNetworkPolicy returns a boolean if a field has been set.

### SetFcNetworkPolicyNil

`func (o *VnicBaseFcIf) SetFcNetworkPolicyNil(b bool)`

 SetFcNetworkPolicyNil sets the value for FcNetworkPolicy to be an explicit nil

### UnsetFcNetworkPolicy
`func (o *VnicBaseFcIf) UnsetFcNetworkPolicy()`

UnsetFcNetworkPolicy ensures that no value is present for FcNetworkPolicy, not even an explicit nil
### GetFcQosPolicy

`func (o *VnicBaseFcIf) GetFcQosPolicy() VnicFcQosPolicyRelationship`

GetFcQosPolicy returns the FcQosPolicy field if non-nil, zero value otherwise.

### GetFcQosPolicyOk

`func (o *VnicBaseFcIf) GetFcQosPolicyOk() (*VnicFcQosPolicyRelationship, bool)`

GetFcQosPolicyOk returns a tuple with the FcQosPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcQosPolicy

`func (o *VnicBaseFcIf) SetFcQosPolicy(v VnicFcQosPolicyRelationship)`

SetFcQosPolicy sets FcQosPolicy field to given value.

### HasFcQosPolicy

`func (o *VnicBaseFcIf) HasFcQosPolicy() bool`

HasFcQosPolicy returns a boolean if a field has been set.

### SetFcQosPolicyNil

`func (o *VnicBaseFcIf) SetFcQosPolicyNil(b bool)`

 SetFcQosPolicyNil sets the value for FcQosPolicy to be an explicit nil

### UnsetFcQosPolicy
`func (o *VnicBaseFcIf) UnsetFcQosPolicy()`

UnsetFcQosPolicy ensures that no value is present for FcQosPolicy, not even an explicit nil
### GetFcZonePolicies

`func (o *VnicBaseFcIf) GetFcZonePolicies() []FabricFcZonePolicyRelationship`

GetFcZonePolicies returns the FcZonePolicies field if non-nil, zero value otherwise.

### GetFcZonePoliciesOk

`func (o *VnicBaseFcIf) GetFcZonePoliciesOk() (*[]FabricFcZonePolicyRelationship, bool)`

GetFcZonePoliciesOk returns a tuple with the FcZonePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcZonePolicies

`func (o *VnicBaseFcIf) SetFcZonePolicies(v []FabricFcZonePolicyRelationship)`

SetFcZonePolicies sets FcZonePolicies field to given value.

### HasFcZonePolicies

`func (o *VnicBaseFcIf) HasFcZonePolicies() bool`

HasFcZonePolicies returns a boolean if a field has been set.

### SetFcZonePoliciesNil

`func (o *VnicBaseFcIf) SetFcZonePoliciesNil(b bool)`

 SetFcZonePoliciesNil sets the value for FcZonePolicies to be an explicit nil

### UnsetFcZonePolicies
`func (o *VnicBaseFcIf) UnsetFcZonePolicies()`

UnsetFcZonePolicies ensures that no value is present for FcZonePolicies, not even an explicit nil
### GetWwpnPool

`func (o *VnicBaseFcIf) GetWwpnPool() FcpoolPoolRelationship`

GetWwpnPool returns the WwpnPool field if non-nil, zero value otherwise.

### GetWwpnPoolOk

`func (o *VnicBaseFcIf) GetWwpnPoolOk() (*FcpoolPoolRelationship, bool)`

GetWwpnPoolOk returns a tuple with the WwpnPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwpnPool

`func (o *VnicBaseFcIf) SetWwpnPool(v FcpoolPoolRelationship)`

SetWwpnPool sets WwpnPool field to given value.

### HasWwpnPool

`func (o *VnicBaseFcIf) HasWwpnPool() bool`

HasWwpnPool returns a boolean if a field has been set.

### SetWwpnPoolNil

`func (o *VnicBaseFcIf) SetWwpnPoolNil(b bool)`

 SetWwpnPoolNil sets the value for WwpnPool to be an explicit nil

### UnsetWwpnPool
`func (o *VnicBaseFcIf) UnsetWwpnPool()`

UnsetWwpnPool ensures that no value is present for WwpnPool, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


