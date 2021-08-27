# KubernetesNodeProfileAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "kubernetes.VirtualMachineNodeProfile"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "kubernetes.VirtualMachineNodeProfile"]
**CloudProvider** | Pointer to **string** | Cloud provider for this node profile. * &#x60;noProvider&#x60; - Enables the use of no cloud provider. * &#x60;external&#x60; - Out of tree cloud provider, e.g. CPI for vsphere. | [optional] [default to "noProvider"]
**ConfigResult** | Pointer to [**KubernetesConfigResultRelationship**](KubernetesConfigResultRelationship.md) |  | [optional] 
**NodeGroup** | Pointer to [**KubernetesNodeGroupProfileRelationship**](KubernetesNodeGroupProfileRelationship.md) |  | [optional] 
**Target** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**Version** | Pointer to [**KubernetesVersionRelationship**](KubernetesVersionRelationship.md) |  | [optional] 

## Methods

### NewKubernetesNodeProfileAllOf

`func NewKubernetesNodeProfileAllOf(classId string, objectType string, ) *KubernetesNodeProfileAllOf`

NewKubernetesNodeProfileAllOf instantiates a new KubernetesNodeProfileAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesNodeProfileAllOfWithDefaults

`func NewKubernetesNodeProfileAllOfWithDefaults() *KubernetesNodeProfileAllOf`

NewKubernetesNodeProfileAllOfWithDefaults instantiates a new KubernetesNodeProfileAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesNodeProfileAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesNodeProfileAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesNodeProfileAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesNodeProfileAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesNodeProfileAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesNodeProfileAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCloudProvider

`func (o *KubernetesNodeProfileAllOf) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *KubernetesNodeProfileAllOf) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *KubernetesNodeProfileAllOf) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *KubernetesNodeProfileAllOf) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetConfigResult

`func (o *KubernetesNodeProfileAllOf) GetConfigResult() KubernetesConfigResultRelationship`

GetConfigResult returns the ConfigResult field if non-nil, zero value otherwise.

### GetConfigResultOk

`func (o *KubernetesNodeProfileAllOf) GetConfigResultOk() (*KubernetesConfigResultRelationship, bool)`

GetConfigResultOk returns a tuple with the ConfigResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigResult

`func (o *KubernetesNodeProfileAllOf) SetConfigResult(v KubernetesConfigResultRelationship)`

SetConfigResult sets ConfigResult field to given value.

### HasConfigResult

`func (o *KubernetesNodeProfileAllOf) HasConfigResult() bool`

HasConfigResult returns a boolean if a field has been set.

### GetNodeGroup

`func (o *KubernetesNodeProfileAllOf) GetNodeGroup() KubernetesNodeGroupProfileRelationship`

GetNodeGroup returns the NodeGroup field if non-nil, zero value otherwise.

### GetNodeGroupOk

`func (o *KubernetesNodeProfileAllOf) GetNodeGroupOk() (*KubernetesNodeGroupProfileRelationship, bool)`

GetNodeGroupOk returns a tuple with the NodeGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeGroup

`func (o *KubernetesNodeProfileAllOf) SetNodeGroup(v KubernetesNodeGroupProfileRelationship)`

SetNodeGroup sets NodeGroup field to given value.

### HasNodeGroup

`func (o *KubernetesNodeProfileAllOf) HasNodeGroup() bool`

HasNodeGroup returns a boolean if a field has been set.

### GetTarget

`func (o *KubernetesNodeProfileAllOf) GetTarget() AssetDeviceRegistrationRelationship`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *KubernetesNodeProfileAllOf) GetTargetOk() (*AssetDeviceRegistrationRelationship, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *KubernetesNodeProfileAllOf) SetTarget(v AssetDeviceRegistrationRelationship)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *KubernetesNodeProfileAllOf) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetVersion

`func (o *KubernetesNodeProfileAllOf) GetVersion() KubernetesVersionRelationship`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *KubernetesNodeProfileAllOf) GetVersionOk() (*KubernetesVersionRelationship, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *KubernetesNodeProfileAllOf) SetVersion(v KubernetesVersionRelationship)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *KubernetesNodeProfileAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


