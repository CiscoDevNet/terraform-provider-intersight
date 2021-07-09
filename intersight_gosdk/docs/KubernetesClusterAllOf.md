# KubernetesClusterAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.Cluster"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.Cluster"]
**ConnectionStatus** | Pointer to **string** | Status of the endpoint connection of this Kubernetes cluster. * &#x60;&#x60; - The target details have been persisted but Intersight has not yet attempted to connect to the target. * &#x60;Connected&#x60; - Intersight is able to establish a connection to the target and initiate management activities. * &#x60;NotConnected&#x60; - Intersight is unable to establish a connection to the target. * &#x60;ClaimInProgress&#x60; - Claim of the target is in progress. A connection to the target has not been fully established. * &#x60;Unclaimed&#x60; - The device was un-claimed from the users account by an Administrator of the device. Also indicates the failure to claim Targets of type HTTP Endpoint in Intersight. * &#x60;Claimed&#x60; - Target of type HTTP Endpoint is successfully claimed in Intersight. Currently no validation is performed to verify the Target connectivity from Intersight at the time of creation. However invoking API from Intersight Orchestrator fails if this Target is not reachable from Intersight or if Target API credentials are incorrect. | [optional] [default to ""]
**KubeConfig** | Pointer to **string** | Kubeconfig for the cluster to collect inventory for. | [optional] 
**Name** | Pointer to **string** | Name of the Kubernetes cluster. | [optional] 
**ClusterAddonProfile** | Pointer to [**KubernetesClusterAddonProfileRelationship**](kubernetes.ClusterAddonProfile.Relationship.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 
**RegisteredDevices** | Pointer to [**[]AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) | An array of relationships to assetDeviceRegistration resources. | [optional] 

## Methods

### NewKubernetesClusterAllOf

`func NewKubernetesClusterAllOf(classId string, objectType string, ) *KubernetesClusterAllOf`

NewKubernetesClusterAllOf instantiates a new KubernetesClusterAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesClusterAllOfWithDefaults

`func NewKubernetesClusterAllOfWithDefaults() *KubernetesClusterAllOf`

NewKubernetesClusterAllOfWithDefaults instantiates a new KubernetesClusterAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesClusterAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesClusterAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesClusterAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesClusterAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesClusterAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesClusterAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConnectionStatus

`func (o *KubernetesClusterAllOf) GetConnectionStatus() string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *KubernetesClusterAllOf) GetConnectionStatusOk() (*string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *KubernetesClusterAllOf) SetConnectionStatus(v string)`

SetConnectionStatus sets ConnectionStatus field to given value.

### HasConnectionStatus

`func (o *KubernetesClusterAllOf) HasConnectionStatus() bool`

HasConnectionStatus returns a boolean if a field has been set.

### GetKubeConfig

`func (o *KubernetesClusterAllOf) GetKubeConfig() string`

GetKubeConfig returns the KubeConfig field if non-nil, zero value otherwise.

### GetKubeConfigOk

`func (o *KubernetesClusterAllOf) GetKubeConfigOk() (*string, bool)`

GetKubeConfigOk returns a tuple with the KubeConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeConfig

`func (o *KubernetesClusterAllOf) SetKubeConfig(v string)`

SetKubeConfig sets KubeConfig field to given value.

### HasKubeConfig

`func (o *KubernetesClusterAllOf) HasKubeConfig() bool`

HasKubeConfig returns a boolean if a field has been set.

### GetName

`func (o *KubernetesClusterAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesClusterAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesClusterAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KubernetesClusterAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetClusterAddonProfile

`func (o *KubernetesClusterAllOf) GetClusterAddonProfile() KubernetesClusterAddonProfileRelationship`

GetClusterAddonProfile returns the ClusterAddonProfile field if non-nil, zero value otherwise.

### GetClusterAddonProfileOk

`func (o *KubernetesClusterAllOf) GetClusterAddonProfileOk() (*KubernetesClusterAddonProfileRelationship, bool)`

GetClusterAddonProfileOk returns a tuple with the ClusterAddonProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterAddonProfile

`func (o *KubernetesClusterAllOf) SetClusterAddonProfile(v KubernetesClusterAddonProfileRelationship)`

SetClusterAddonProfile sets ClusterAddonProfile field to given value.

### HasClusterAddonProfile

`func (o *KubernetesClusterAllOf) HasClusterAddonProfile() bool`

HasClusterAddonProfile returns a boolean if a field has been set.

### GetOrganization

`func (o *KubernetesClusterAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *KubernetesClusterAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *KubernetesClusterAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *KubernetesClusterAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetRegisteredDevices

`func (o *KubernetesClusterAllOf) GetRegisteredDevices() []AssetDeviceRegistrationRelationship`

GetRegisteredDevices returns the RegisteredDevices field if non-nil, zero value otherwise.

### GetRegisteredDevicesOk

`func (o *KubernetesClusterAllOf) GetRegisteredDevicesOk() (*[]AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDevicesOk returns a tuple with the RegisteredDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevices

`func (o *KubernetesClusterAllOf) SetRegisteredDevices(v []AssetDeviceRegistrationRelationship)`

SetRegisteredDevices sets RegisteredDevices field to given value.

### HasRegisteredDevices

`func (o *KubernetesClusterAllOf) HasRegisteredDevices() bool`

HasRegisteredDevices returns a boolean if a field has been set.

### SetRegisteredDevicesNil

`func (o *KubernetesClusterAllOf) SetRegisteredDevicesNil(b bool)`

 SetRegisteredDevicesNil sets the value for RegisteredDevices to be an explicit nil

### UnsetRegisteredDevices
`func (o *KubernetesClusterAllOf) UnsetRegisteredDevices()`

UnsetRegisteredDevices ensures that no value is present for RegisteredDevices, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


