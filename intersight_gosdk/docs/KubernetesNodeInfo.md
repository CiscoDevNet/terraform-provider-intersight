# KubernetesNodeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.NodeInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.NodeInfo"]
**Architecture** | Pointer to **string** | Node Operating System architecture (amd64, arm64). | [optional] 
**BootId** | Pointer to **string** | A Boot ID is an Identifier generated by the host everytimes it gets reboot. | [optional] 
**ContainerRuntimeVersion** | Pointer to **string** | To run containers in Pods, Kubernetes uses a container runtime. This field describes Container Runtime Version. | [optional] 
**KernelVersion** | Pointer to **string** | Node Operating System kernel version. | [optional] 
**KubeProxyVersion** | Pointer to **string** | The Kubernetes network proxy runs on each node. This reflects services as defined in the Kubernetes API on each node and can do simple TCP, UDP, and SCTP stream forwarding or round robin TCP, UDP, and SCTP forwarding across a set of backends. This field describes the kube-proxy version. | [optional] 
**KubeletVersion** | Pointer to **string** | The kubelet is the primary \&quot;node agent\&quot; that runs on each node. It can register the node with the apiserver using one of such as the hostname; a flag to override the hostname; or specific logic for a cloud provider. This field describes the kubelet version the node currently using. | [optional] 
**MachineId** | Pointer to **string** | It is a node identifier in Kubernetes. When the node joins a kubernetes cluster, Kubernetes will assign a machine ID to that node. Learn more from man machine-id from http://man7.org/linux/man-pages/man5/machine-id.5.html. | [optional] 
**OperatingSystem** | Pointer to **string** | Operating System installed on this Node. | [optional] 
**OsImage** | Pointer to **string** | Node current Operating System image. | [optional] 
**SystemUuid** | Pointer to **string** | SystemUUID reported by the node. For unique machine identification MachineID is preferred. This field is specific to Red Hat hosts https://access.redhat.com/documentation/en-US/Red_Hat_Subscription_Management/1/html/RHSM/getting-system-uuid.html. | [optional] 

## Methods

### NewKubernetesNodeInfo

`func NewKubernetesNodeInfo(classId string, objectType string, ) *KubernetesNodeInfo`

NewKubernetesNodeInfo instantiates a new KubernetesNodeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesNodeInfoWithDefaults

`func NewKubernetesNodeInfoWithDefaults() *KubernetesNodeInfo`

NewKubernetesNodeInfoWithDefaults instantiates a new KubernetesNodeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesNodeInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesNodeInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesNodeInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesNodeInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesNodeInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesNodeInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetArchitecture

`func (o *KubernetesNodeInfo) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *KubernetesNodeInfo) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *KubernetesNodeInfo) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *KubernetesNodeInfo) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### GetBootId

`func (o *KubernetesNodeInfo) GetBootId() string`

GetBootId returns the BootId field if non-nil, zero value otherwise.

### GetBootIdOk

`func (o *KubernetesNodeInfo) GetBootIdOk() (*string, bool)`

GetBootIdOk returns a tuple with the BootId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootId

`func (o *KubernetesNodeInfo) SetBootId(v string)`

SetBootId sets BootId field to given value.

### HasBootId

`func (o *KubernetesNodeInfo) HasBootId() bool`

HasBootId returns a boolean if a field has been set.

### GetContainerRuntimeVersion

`func (o *KubernetesNodeInfo) GetContainerRuntimeVersion() string`

GetContainerRuntimeVersion returns the ContainerRuntimeVersion field if non-nil, zero value otherwise.

### GetContainerRuntimeVersionOk

`func (o *KubernetesNodeInfo) GetContainerRuntimeVersionOk() (*string, bool)`

GetContainerRuntimeVersionOk returns a tuple with the ContainerRuntimeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerRuntimeVersion

`func (o *KubernetesNodeInfo) SetContainerRuntimeVersion(v string)`

SetContainerRuntimeVersion sets ContainerRuntimeVersion field to given value.

### HasContainerRuntimeVersion

`func (o *KubernetesNodeInfo) HasContainerRuntimeVersion() bool`

HasContainerRuntimeVersion returns a boolean if a field has been set.

### GetKernelVersion

`func (o *KubernetesNodeInfo) GetKernelVersion() string`

GetKernelVersion returns the KernelVersion field if non-nil, zero value otherwise.

### GetKernelVersionOk

`func (o *KubernetesNodeInfo) GetKernelVersionOk() (*string, bool)`

GetKernelVersionOk returns a tuple with the KernelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKernelVersion

`func (o *KubernetesNodeInfo) SetKernelVersion(v string)`

SetKernelVersion sets KernelVersion field to given value.

### HasKernelVersion

`func (o *KubernetesNodeInfo) HasKernelVersion() bool`

HasKernelVersion returns a boolean if a field has been set.

### GetKubeProxyVersion

`func (o *KubernetesNodeInfo) GetKubeProxyVersion() string`

GetKubeProxyVersion returns the KubeProxyVersion field if non-nil, zero value otherwise.

### GetKubeProxyVersionOk

`func (o *KubernetesNodeInfo) GetKubeProxyVersionOk() (*string, bool)`

GetKubeProxyVersionOk returns a tuple with the KubeProxyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeProxyVersion

`func (o *KubernetesNodeInfo) SetKubeProxyVersion(v string)`

SetKubeProxyVersion sets KubeProxyVersion field to given value.

### HasKubeProxyVersion

`func (o *KubernetesNodeInfo) HasKubeProxyVersion() bool`

HasKubeProxyVersion returns a boolean if a field has been set.

### GetKubeletVersion

`func (o *KubernetesNodeInfo) GetKubeletVersion() string`

GetKubeletVersion returns the KubeletVersion field if non-nil, zero value otherwise.

### GetKubeletVersionOk

`func (o *KubernetesNodeInfo) GetKubeletVersionOk() (*string, bool)`

GetKubeletVersionOk returns a tuple with the KubeletVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeletVersion

`func (o *KubernetesNodeInfo) SetKubeletVersion(v string)`

SetKubeletVersion sets KubeletVersion field to given value.

### HasKubeletVersion

`func (o *KubernetesNodeInfo) HasKubeletVersion() bool`

HasKubeletVersion returns a boolean if a field has been set.

### GetMachineId

`func (o *KubernetesNodeInfo) GetMachineId() string`

GetMachineId returns the MachineId field if non-nil, zero value otherwise.

### GetMachineIdOk

`func (o *KubernetesNodeInfo) GetMachineIdOk() (*string, bool)`

GetMachineIdOk returns a tuple with the MachineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineId

`func (o *KubernetesNodeInfo) SetMachineId(v string)`

SetMachineId sets MachineId field to given value.

### HasMachineId

`func (o *KubernetesNodeInfo) HasMachineId() bool`

HasMachineId returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *KubernetesNodeInfo) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *KubernetesNodeInfo) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *KubernetesNodeInfo) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *KubernetesNodeInfo) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetOsImage

`func (o *KubernetesNodeInfo) GetOsImage() string`

GetOsImage returns the OsImage field if non-nil, zero value otherwise.

### GetOsImageOk

`func (o *KubernetesNodeInfo) GetOsImageOk() (*string, bool)`

GetOsImageOk returns a tuple with the OsImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsImage

`func (o *KubernetesNodeInfo) SetOsImage(v string)`

SetOsImage sets OsImage field to given value.

### HasOsImage

`func (o *KubernetesNodeInfo) HasOsImage() bool`

HasOsImage returns a boolean if a field has been set.

### GetSystemUuid

`func (o *KubernetesNodeInfo) GetSystemUuid() string`

GetSystemUuid returns the SystemUuid field if non-nil, zero value otherwise.

### GetSystemUuidOk

`func (o *KubernetesNodeInfo) GetSystemUuidOk() (*string, bool)`

GetSystemUuidOk returns a tuple with the SystemUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemUuid

`func (o *KubernetesNodeInfo) SetSystemUuid(v string)`

SetSystemUuid sets SystemUuid field to given value.

### HasSystemUuid

`func (o *KubernetesNodeInfo) HasSystemUuid() bool`

HasSystemUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


