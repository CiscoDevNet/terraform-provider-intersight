# KubernetesNodeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.Node"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.Node"]
**NodeAddresses** | Pointer to [**[]KubernetesNodeAddress**](KubernetesNodeAddress.md) |  | [optional] 
**NodeInfo** | Pointer to [**NullableKubernetesNodeInfo**](KubernetesNodeInfo.md) |  | [optional] 
**NodeSpec** | Pointer to [**NullableKubernetesNodeSpec**](KubernetesNodeSpec.md) |  | [optional] 
**NodeStatuses** | Pointer to [**[]KubernetesNodeStatus**](KubernetesNodeStatus.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewKubernetesNodeAllOf

`func NewKubernetesNodeAllOf(classId string, objectType string, ) *KubernetesNodeAllOf`

NewKubernetesNodeAllOf instantiates a new KubernetesNodeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesNodeAllOfWithDefaults

`func NewKubernetesNodeAllOfWithDefaults() *KubernetesNodeAllOf`

NewKubernetesNodeAllOfWithDefaults instantiates a new KubernetesNodeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesNodeAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesNodeAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesNodeAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesNodeAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesNodeAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesNodeAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetNodeAddresses

`func (o *KubernetesNodeAllOf) GetNodeAddresses() []KubernetesNodeAddress`

GetNodeAddresses returns the NodeAddresses field if non-nil, zero value otherwise.

### GetNodeAddressesOk

`func (o *KubernetesNodeAllOf) GetNodeAddressesOk() (*[]KubernetesNodeAddress, bool)`

GetNodeAddressesOk returns a tuple with the NodeAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeAddresses

`func (o *KubernetesNodeAllOf) SetNodeAddresses(v []KubernetesNodeAddress)`

SetNodeAddresses sets NodeAddresses field to given value.

### HasNodeAddresses

`func (o *KubernetesNodeAllOf) HasNodeAddresses() bool`

HasNodeAddresses returns a boolean if a field has been set.

### SetNodeAddressesNil

`func (o *KubernetesNodeAllOf) SetNodeAddressesNil(b bool)`

 SetNodeAddressesNil sets the value for NodeAddresses to be an explicit nil

### UnsetNodeAddresses
`func (o *KubernetesNodeAllOf) UnsetNodeAddresses()`

UnsetNodeAddresses ensures that no value is present for NodeAddresses, not even an explicit nil
### GetNodeInfo

`func (o *KubernetesNodeAllOf) GetNodeInfo() KubernetesNodeInfo`

GetNodeInfo returns the NodeInfo field if non-nil, zero value otherwise.

### GetNodeInfoOk

`func (o *KubernetesNodeAllOf) GetNodeInfoOk() (*KubernetesNodeInfo, bool)`

GetNodeInfoOk returns a tuple with the NodeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeInfo

`func (o *KubernetesNodeAllOf) SetNodeInfo(v KubernetesNodeInfo)`

SetNodeInfo sets NodeInfo field to given value.

### HasNodeInfo

`func (o *KubernetesNodeAllOf) HasNodeInfo() bool`

HasNodeInfo returns a boolean if a field has been set.

### SetNodeInfoNil

`func (o *KubernetesNodeAllOf) SetNodeInfoNil(b bool)`

 SetNodeInfoNil sets the value for NodeInfo to be an explicit nil

### UnsetNodeInfo
`func (o *KubernetesNodeAllOf) UnsetNodeInfo()`

UnsetNodeInfo ensures that no value is present for NodeInfo, not even an explicit nil
### GetNodeSpec

`func (o *KubernetesNodeAllOf) GetNodeSpec() KubernetesNodeSpec`

GetNodeSpec returns the NodeSpec field if non-nil, zero value otherwise.

### GetNodeSpecOk

`func (o *KubernetesNodeAllOf) GetNodeSpecOk() (*KubernetesNodeSpec, bool)`

GetNodeSpecOk returns a tuple with the NodeSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSpec

`func (o *KubernetesNodeAllOf) SetNodeSpec(v KubernetesNodeSpec)`

SetNodeSpec sets NodeSpec field to given value.

### HasNodeSpec

`func (o *KubernetesNodeAllOf) HasNodeSpec() bool`

HasNodeSpec returns a boolean if a field has been set.

### SetNodeSpecNil

`func (o *KubernetesNodeAllOf) SetNodeSpecNil(b bool)`

 SetNodeSpecNil sets the value for NodeSpec to be an explicit nil

### UnsetNodeSpec
`func (o *KubernetesNodeAllOf) UnsetNodeSpec()`

UnsetNodeSpec ensures that no value is present for NodeSpec, not even an explicit nil
### GetNodeStatuses

`func (o *KubernetesNodeAllOf) GetNodeStatuses() []KubernetesNodeStatus`

GetNodeStatuses returns the NodeStatuses field if non-nil, zero value otherwise.

### GetNodeStatusesOk

`func (o *KubernetesNodeAllOf) GetNodeStatusesOk() (*[]KubernetesNodeStatus, bool)`

GetNodeStatusesOk returns a tuple with the NodeStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeStatuses

`func (o *KubernetesNodeAllOf) SetNodeStatuses(v []KubernetesNodeStatus)`

SetNodeStatuses sets NodeStatuses field to given value.

### HasNodeStatuses

`func (o *KubernetesNodeAllOf) HasNodeStatuses() bool`

HasNodeStatuses returns a boolean if a field has been set.

### SetNodeStatusesNil

`func (o *KubernetesNodeAllOf) SetNodeStatusesNil(b bool)`

 SetNodeStatusesNil sets the value for NodeStatuses to be an explicit nil

### UnsetNodeStatuses
`func (o *KubernetesNodeAllOf) UnsetNodeStatuses()`

UnsetNodeStatuses ensures that no value is present for NodeStatuses, not even an explicit nil
### GetRegisteredDevice

`func (o *KubernetesNodeAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *KubernetesNodeAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *KubernetesNodeAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *KubernetesNodeAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


