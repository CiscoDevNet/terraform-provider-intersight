# KubernetesNodeProfile

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

### NewKubernetesNodeProfile

`func NewKubernetesNodeProfile(classId string, objectType string, ) *KubernetesNodeProfile`

NewKubernetesNodeProfile instantiates a new KubernetesNodeProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesNodeProfileWithDefaults

`func NewKubernetesNodeProfileWithDefaults() *KubernetesNodeProfile`

NewKubernetesNodeProfileWithDefaults instantiates a new KubernetesNodeProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesNodeProfile) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesNodeProfile) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesNodeProfile) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesNodeProfile) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesNodeProfile) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesNodeProfile) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCloudProvider

`func (o *KubernetesNodeProfile) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *KubernetesNodeProfile) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *KubernetesNodeProfile) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *KubernetesNodeProfile) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetConfigResult

`func (o *KubernetesNodeProfile) GetConfigResult() KubernetesConfigResultRelationship`

GetConfigResult returns the ConfigResult field if non-nil, zero value otherwise.

### GetConfigResultOk

`func (o *KubernetesNodeProfile) GetConfigResultOk() (*KubernetesConfigResultRelationship, bool)`

GetConfigResultOk returns a tuple with the ConfigResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigResult

`func (o *KubernetesNodeProfile) SetConfigResult(v KubernetesConfigResultRelationship)`

SetConfigResult sets ConfigResult field to given value.

### HasConfigResult

`func (o *KubernetesNodeProfile) HasConfigResult() bool`

HasConfigResult returns a boolean if a field has been set.

### GetNodeGroup

`func (o *KubernetesNodeProfile) GetNodeGroup() KubernetesNodeGroupProfileRelationship`

GetNodeGroup returns the NodeGroup field if non-nil, zero value otherwise.

### GetNodeGroupOk

`func (o *KubernetesNodeProfile) GetNodeGroupOk() (*KubernetesNodeGroupProfileRelationship, bool)`

GetNodeGroupOk returns a tuple with the NodeGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeGroup

`func (o *KubernetesNodeProfile) SetNodeGroup(v KubernetesNodeGroupProfileRelationship)`

SetNodeGroup sets NodeGroup field to given value.

### HasNodeGroup

`func (o *KubernetesNodeProfile) HasNodeGroup() bool`

HasNodeGroup returns a boolean if a field has been set.

### GetTarget

`func (o *KubernetesNodeProfile) GetTarget() AssetDeviceRegistrationRelationship`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *KubernetesNodeProfile) GetTargetOk() (*AssetDeviceRegistrationRelationship, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *KubernetesNodeProfile) SetTarget(v AssetDeviceRegistrationRelationship)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *KubernetesNodeProfile) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetVersion

`func (o *KubernetesNodeProfile) GetVersion() KubernetesVersionRelationship`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *KubernetesNodeProfile) GetVersionOk() (*KubernetesVersionRelationship, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *KubernetesNodeProfile) SetVersion(v KubernetesVersionRelationship)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *KubernetesNodeProfile) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


