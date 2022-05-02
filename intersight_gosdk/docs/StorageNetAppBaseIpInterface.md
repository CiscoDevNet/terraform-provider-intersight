# StorageNetAppBaseIpInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**CurrentNode** | Pointer to **string** | Name of the current node of IP interface. | [optional] [readonly] 
**CurrentPort** | Pointer to **string** | Name of the current port of IP interface. | [optional] [readonly] 
**Enabled** | Pointer to **string** | IP interface is enabled or not. | [optional] [readonly] 
**HomeNode** | Pointer to **string** | Name of home node of IP interface. | [optional] [readonly] 
**HomePort** | Pointer to **string** | Name of home port of IP interface. | [optional] [readonly] 
**IpAddress** | Pointer to **string** | The IP address of interface. | [optional] [readonly] 
**IpFamily** | Pointer to **string** | IP address family of interface. * &#x60;IPv4&#x60; - IP address family type is IPv4. * &#x60;IPv6&#x60; - IP address family type is IP6. | [optional] [readonly] [default to "IPv4"]
**Ipspace** | Pointer to **string** | The name of the IPspace of the IP interface. | [optional] [readonly] 
**IsHome** | Pointer to **bool** | Reports whether the IP interface is home or has failed over to its HA peer. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the IP interface. | [optional] [readonly] 
**Netmask** | Pointer to **string** | The netmask of the interface. | [optional] [readonly] 
**ServicePolicyName** | Pointer to **string** | Service policy name of IP interface. | [optional] [readonly] 
**ServicePolicyUuid** | Pointer to **string** | Service policy UUID of IP interface. | [optional] [readonly] 
**Services** | Pointer to **[]string** |  | [optional] 
**State** | Pointer to **string** | The state of the IP interface. * &#x60;down&#x60; - An inactive port is listed as Down. * &#x60;up&#x60; - An active port is listed as Up. * &#x60;present&#x60; - An active port is listed as present. | [optional] [readonly] [default to "down"]
**Uuid** | Pointer to **string** | Uuid of NetApp IP Interface. | [optional] [readonly] 

## Methods

### NewStorageNetAppBaseIpInterface

`func NewStorageNetAppBaseIpInterface(classId string, objectType string, ) *StorageNetAppBaseIpInterface`

NewStorageNetAppBaseIpInterface instantiates a new StorageNetAppBaseIpInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppBaseIpInterfaceWithDefaults

`func NewStorageNetAppBaseIpInterfaceWithDefaults() *StorageNetAppBaseIpInterface`

NewStorageNetAppBaseIpInterfaceWithDefaults instantiates a new StorageNetAppBaseIpInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppBaseIpInterface) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppBaseIpInterface) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppBaseIpInterface) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppBaseIpInterface) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppBaseIpInterface) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppBaseIpInterface) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCurrentNode

`func (o *StorageNetAppBaseIpInterface) GetCurrentNode() string`

GetCurrentNode returns the CurrentNode field if non-nil, zero value otherwise.

### GetCurrentNodeOk

`func (o *StorageNetAppBaseIpInterface) GetCurrentNodeOk() (*string, bool)`

GetCurrentNodeOk returns a tuple with the CurrentNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentNode

`func (o *StorageNetAppBaseIpInterface) SetCurrentNode(v string)`

SetCurrentNode sets CurrentNode field to given value.

### HasCurrentNode

`func (o *StorageNetAppBaseIpInterface) HasCurrentNode() bool`

HasCurrentNode returns a boolean if a field has been set.

### GetCurrentPort

`func (o *StorageNetAppBaseIpInterface) GetCurrentPort() string`

GetCurrentPort returns the CurrentPort field if non-nil, zero value otherwise.

### GetCurrentPortOk

`func (o *StorageNetAppBaseIpInterface) GetCurrentPortOk() (*string, bool)`

GetCurrentPortOk returns a tuple with the CurrentPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPort

`func (o *StorageNetAppBaseIpInterface) SetCurrentPort(v string)`

SetCurrentPort sets CurrentPort field to given value.

### HasCurrentPort

`func (o *StorageNetAppBaseIpInterface) HasCurrentPort() bool`

HasCurrentPort returns a boolean if a field has been set.

### GetEnabled

`func (o *StorageNetAppBaseIpInterface) GetEnabled() string`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *StorageNetAppBaseIpInterface) GetEnabledOk() (*string, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *StorageNetAppBaseIpInterface) SetEnabled(v string)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *StorageNetAppBaseIpInterface) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHomeNode

`func (o *StorageNetAppBaseIpInterface) GetHomeNode() string`

GetHomeNode returns the HomeNode field if non-nil, zero value otherwise.

### GetHomeNodeOk

`func (o *StorageNetAppBaseIpInterface) GetHomeNodeOk() (*string, bool)`

GetHomeNodeOk returns a tuple with the HomeNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeNode

`func (o *StorageNetAppBaseIpInterface) SetHomeNode(v string)`

SetHomeNode sets HomeNode field to given value.

### HasHomeNode

`func (o *StorageNetAppBaseIpInterface) HasHomeNode() bool`

HasHomeNode returns a boolean if a field has been set.

### GetHomePort

`func (o *StorageNetAppBaseIpInterface) GetHomePort() string`

GetHomePort returns the HomePort field if non-nil, zero value otherwise.

### GetHomePortOk

`func (o *StorageNetAppBaseIpInterface) GetHomePortOk() (*string, bool)`

GetHomePortOk returns a tuple with the HomePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomePort

`func (o *StorageNetAppBaseIpInterface) SetHomePort(v string)`

SetHomePort sets HomePort field to given value.

### HasHomePort

`func (o *StorageNetAppBaseIpInterface) HasHomePort() bool`

HasHomePort returns a boolean if a field has been set.

### GetIpAddress

`func (o *StorageNetAppBaseIpInterface) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *StorageNetAppBaseIpInterface) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *StorageNetAppBaseIpInterface) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *StorageNetAppBaseIpInterface) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetIpFamily

`func (o *StorageNetAppBaseIpInterface) GetIpFamily() string`

GetIpFamily returns the IpFamily field if non-nil, zero value otherwise.

### GetIpFamilyOk

`func (o *StorageNetAppBaseIpInterface) GetIpFamilyOk() (*string, bool)`

GetIpFamilyOk returns a tuple with the IpFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpFamily

`func (o *StorageNetAppBaseIpInterface) SetIpFamily(v string)`

SetIpFamily sets IpFamily field to given value.

### HasIpFamily

`func (o *StorageNetAppBaseIpInterface) HasIpFamily() bool`

HasIpFamily returns a boolean if a field has been set.

### GetIpspace

`func (o *StorageNetAppBaseIpInterface) GetIpspace() string`

GetIpspace returns the Ipspace field if non-nil, zero value otherwise.

### GetIpspaceOk

`func (o *StorageNetAppBaseIpInterface) GetIpspaceOk() (*string, bool)`

GetIpspaceOk returns a tuple with the Ipspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpspace

`func (o *StorageNetAppBaseIpInterface) SetIpspace(v string)`

SetIpspace sets Ipspace field to given value.

### HasIpspace

`func (o *StorageNetAppBaseIpInterface) HasIpspace() bool`

HasIpspace returns a boolean if a field has been set.

### GetIsHome

`func (o *StorageNetAppBaseIpInterface) GetIsHome() bool`

GetIsHome returns the IsHome field if non-nil, zero value otherwise.

### GetIsHomeOk

`func (o *StorageNetAppBaseIpInterface) GetIsHomeOk() (*bool, bool)`

GetIsHomeOk returns a tuple with the IsHome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHome

`func (o *StorageNetAppBaseIpInterface) SetIsHome(v bool)`

SetIsHome sets IsHome field to given value.

### HasIsHome

`func (o *StorageNetAppBaseIpInterface) HasIsHome() bool`

HasIsHome returns a boolean if a field has been set.

### GetName

`func (o *StorageNetAppBaseIpInterface) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageNetAppBaseIpInterface) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageNetAppBaseIpInterface) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageNetAppBaseIpInterface) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetmask

`func (o *StorageNetAppBaseIpInterface) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *StorageNetAppBaseIpInterface) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *StorageNetAppBaseIpInterface) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *StorageNetAppBaseIpInterface) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetServicePolicyName

`func (o *StorageNetAppBaseIpInterface) GetServicePolicyName() string`

GetServicePolicyName returns the ServicePolicyName field if non-nil, zero value otherwise.

### GetServicePolicyNameOk

`func (o *StorageNetAppBaseIpInterface) GetServicePolicyNameOk() (*string, bool)`

GetServicePolicyNameOk returns a tuple with the ServicePolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePolicyName

`func (o *StorageNetAppBaseIpInterface) SetServicePolicyName(v string)`

SetServicePolicyName sets ServicePolicyName field to given value.

### HasServicePolicyName

`func (o *StorageNetAppBaseIpInterface) HasServicePolicyName() bool`

HasServicePolicyName returns a boolean if a field has been set.

### GetServicePolicyUuid

`func (o *StorageNetAppBaseIpInterface) GetServicePolicyUuid() string`

GetServicePolicyUuid returns the ServicePolicyUuid field if non-nil, zero value otherwise.

### GetServicePolicyUuidOk

`func (o *StorageNetAppBaseIpInterface) GetServicePolicyUuidOk() (*string, bool)`

GetServicePolicyUuidOk returns a tuple with the ServicePolicyUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePolicyUuid

`func (o *StorageNetAppBaseIpInterface) SetServicePolicyUuid(v string)`

SetServicePolicyUuid sets ServicePolicyUuid field to given value.

### HasServicePolicyUuid

`func (o *StorageNetAppBaseIpInterface) HasServicePolicyUuid() bool`

HasServicePolicyUuid returns a boolean if a field has been set.

### GetServices

`func (o *StorageNetAppBaseIpInterface) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *StorageNetAppBaseIpInterface) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *StorageNetAppBaseIpInterface) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *StorageNetAppBaseIpInterface) HasServices() bool`

HasServices returns a boolean if a field has been set.

### SetServicesNil

`func (o *StorageNetAppBaseIpInterface) SetServicesNil(b bool)`

 SetServicesNil sets the value for Services to be an explicit nil

### UnsetServices
`func (o *StorageNetAppBaseIpInterface) UnsetServices()`

UnsetServices ensures that no value is present for Services, not even an explicit nil
### GetState

`func (o *StorageNetAppBaseIpInterface) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StorageNetAppBaseIpInterface) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StorageNetAppBaseIpInterface) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *StorageNetAppBaseIpInterface) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUuid

`func (o *StorageNetAppBaseIpInterface) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StorageNetAppBaseIpInterface) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StorageNetAppBaseIpInterface) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *StorageNetAppBaseIpInterface) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


