# KubernetesClusterManagementConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.ClusterManagementConfig"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.ClusterManagementConfig"]
**EncryptedEtcd** | Pointer to **bool** | Encrypt ETCD data at rest using the etcdEncryptionKey specified in the cluster Kubernetes configuration. | [optional] 
**LoadBalancerCount** | Pointer to **int64** | Number of IP addresses to reserve for load balancer services. | [optional] 
**LoadBalancers** | Pointer to **[]string** |  | [optional] 
**MasterVip** | Pointer to **string** | VIP for the cluster Kubernetes API server. If this is empty and a cluster IP pool is specified, it will be allocated from the IP pool. | [optional] 
**SshKeys** | Pointer to **[]string** |  | [optional] 
**SshUser** | Pointer to **string** | Name of the user to SSH to nodes in a cluster. | [optional] [readonly] [default to "iksadmin"]

## Methods

### NewKubernetesClusterManagementConfig

`func NewKubernetesClusterManagementConfig(classId string, objectType string, ) *KubernetesClusterManagementConfig`

NewKubernetesClusterManagementConfig instantiates a new KubernetesClusterManagementConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesClusterManagementConfigWithDefaults

`func NewKubernetesClusterManagementConfigWithDefaults() *KubernetesClusterManagementConfig`

NewKubernetesClusterManagementConfigWithDefaults instantiates a new KubernetesClusterManagementConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesClusterManagementConfig) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesClusterManagementConfig) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesClusterManagementConfig) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesClusterManagementConfig) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesClusterManagementConfig) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesClusterManagementConfig) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEncryptedEtcd

`func (o *KubernetesClusterManagementConfig) GetEncryptedEtcd() bool`

GetEncryptedEtcd returns the EncryptedEtcd field if non-nil, zero value otherwise.

### GetEncryptedEtcdOk

`func (o *KubernetesClusterManagementConfig) GetEncryptedEtcdOk() (*bool, bool)`

GetEncryptedEtcdOk returns a tuple with the EncryptedEtcd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedEtcd

`func (o *KubernetesClusterManagementConfig) SetEncryptedEtcd(v bool)`

SetEncryptedEtcd sets EncryptedEtcd field to given value.

### HasEncryptedEtcd

`func (o *KubernetesClusterManagementConfig) HasEncryptedEtcd() bool`

HasEncryptedEtcd returns a boolean if a field has been set.

### GetLoadBalancerCount

`func (o *KubernetesClusterManagementConfig) GetLoadBalancerCount() int64`

GetLoadBalancerCount returns the LoadBalancerCount field if non-nil, zero value otherwise.

### GetLoadBalancerCountOk

`func (o *KubernetesClusterManagementConfig) GetLoadBalancerCountOk() (*int64, bool)`

GetLoadBalancerCountOk returns a tuple with the LoadBalancerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerCount

`func (o *KubernetesClusterManagementConfig) SetLoadBalancerCount(v int64)`

SetLoadBalancerCount sets LoadBalancerCount field to given value.

### HasLoadBalancerCount

`func (o *KubernetesClusterManagementConfig) HasLoadBalancerCount() bool`

HasLoadBalancerCount returns a boolean if a field has been set.

### GetLoadBalancers

`func (o *KubernetesClusterManagementConfig) GetLoadBalancers() []string`

GetLoadBalancers returns the LoadBalancers field if non-nil, zero value otherwise.

### GetLoadBalancersOk

`func (o *KubernetesClusterManagementConfig) GetLoadBalancersOk() (*[]string, bool)`

GetLoadBalancersOk returns a tuple with the LoadBalancers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancers

`func (o *KubernetesClusterManagementConfig) SetLoadBalancers(v []string)`

SetLoadBalancers sets LoadBalancers field to given value.

### HasLoadBalancers

`func (o *KubernetesClusterManagementConfig) HasLoadBalancers() bool`

HasLoadBalancers returns a boolean if a field has been set.

### SetLoadBalancersNil

`func (o *KubernetesClusterManagementConfig) SetLoadBalancersNil(b bool)`

 SetLoadBalancersNil sets the value for LoadBalancers to be an explicit nil

### UnsetLoadBalancers
`func (o *KubernetesClusterManagementConfig) UnsetLoadBalancers()`

UnsetLoadBalancers ensures that no value is present for LoadBalancers, not even an explicit nil
### GetMasterVip

`func (o *KubernetesClusterManagementConfig) GetMasterVip() string`

GetMasterVip returns the MasterVip field if non-nil, zero value otherwise.

### GetMasterVipOk

`func (o *KubernetesClusterManagementConfig) GetMasterVipOk() (*string, bool)`

GetMasterVipOk returns a tuple with the MasterVip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterVip

`func (o *KubernetesClusterManagementConfig) SetMasterVip(v string)`

SetMasterVip sets MasterVip field to given value.

### HasMasterVip

`func (o *KubernetesClusterManagementConfig) HasMasterVip() bool`

HasMasterVip returns a boolean if a field has been set.

### GetSshKeys

`func (o *KubernetesClusterManagementConfig) GetSshKeys() []string`

GetSshKeys returns the SshKeys field if non-nil, zero value otherwise.

### GetSshKeysOk

`func (o *KubernetesClusterManagementConfig) GetSshKeysOk() (*[]string, bool)`

GetSshKeysOk returns a tuple with the SshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeys

`func (o *KubernetesClusterManagementConfig) SetSshKeys(v []string)`

SetSshKeys sets SshKeys field to given value.

### HasSshKeys

`func (o *KubernetesClusterManagementConfig) HasSshKeys() bool`

HasSshKeys returns a boolean if a field has been set.

### SetSshKeysNil

`func (o *KubernetesClusterManagementConfig) SetSshKeysNil(b bool)`

 SetSshKeysNil sets the value for SshKeys to be an explicit nil

### UnsetSshKeys
`func (o *KubernetesClusterManagementConfig) UnsetSshKeys()`

UnsetSshKeys ensures that no value is present for SshKeys, not even an explicit nil
### GetSshUser

`func (o *KubernetesClusterManagementConfig) GetSshUser() string`

GetSshUser returns the SshUser field if non-nil, zero value otherwise.

### GetSshUserOk

`func (o *KubernetesClusterManagementConfig) GetSshUserOk() (*string, bool)`

GetSshUserOk returns a tuple with the SshUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUser

`func (o *KubernetesClusterManagementConfig) SetSshUser(v string)`

SetSshUser sets SshUser field to given value.

### HasSshUser

`func (o *KubernetesClusterManagementConfig) HasSshUser() bool`

HasSshUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


