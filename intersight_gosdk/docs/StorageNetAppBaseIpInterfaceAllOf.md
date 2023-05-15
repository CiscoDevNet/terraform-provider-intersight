# StorageNetAppBaseIpInterfaceAllOf

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
**InterfaceIsHome** | Pointer to **string** | Reports whether the IP interface is home or has failed over to its HA peer. | [optional] [readonly] 
**InterfaceState** | Pointer to **string** | The state of the IP interface. * &#x60;Down&#x60; - The state is set to down if the interface is not enabled. * &#x60;Up&#x60; - The state is set to up if the interface is enabled. | [optional] [readonly] [default to "Down"]
**IpAddress** | Pointer to **string** | The IP address of interface. | [optional] [readonly] 
**IpFamily** | Pointer to **string** | IP address family of interface. * &#x60;IPv4&#x60; - IP address family type is IPv4. * &#x60;IPv6&#x60; - IP address family type is IP6. | [optional] [readonly] [default to "IPv4"]
**Ipspace** | Pointer to **string** | The name of the IPspace of the IP interface. | [optional] [readonly] 
**IsHome** | Pointer to **bool** | Reports whether the IP interface is home or has failed over to its HA peer. | [optional] [readonly] 
**LocationFailover** | Pointer to **string** | Defines where an interface may failover, [ home_port_only, default, home_node_only, sfo_partners_only, broadcast_domain_only ]. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the IP interface. | [optional] [readonly] 
**Netmask** | Pointer to **string** | The netmask of the interface. | [optional] [readonly] 
**ServicePolicyName** | Pointer to **string** | Service policy name of IP interface. | [optional] [readonly] 
**ServicePolicyUuid** | Pointer to **string** | Service policy UUID of IP interface. | [optional] [readonly] 
**Services** | Pointer to **[]string** |  | [optional] 
**State** | Pointer to **string** | The state of the IP interface. * &#x60;down&#x60; - An inactive port is listed as Down. * &#x60;up&#x60; - An active port is listed as Up. * &#x60;present&#x60; - An active port is listed as present. | [optional] [readonly] [default to "down"]
**SvmName** | Pointer to **string** | The storage virtual machine name for the interface. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Uuid of NetApp IP Interface. | [optional] [readonly] 

## Methods

### NewStorageNetAppBaseIpInterfaceAllOf

`func NewStorageNetAppBaseIpInterfaceAllOf(classId string, objectType string, ) *StorageNetAppBaseIpInterfaceAllOf`

NewStorageNetAppBaseIpInterfaceAllOf instantiates a new StorageNetAppBaseIpInterfaceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppBaseIpInterfaceAllOfWithDefaults

`func NewStorageNetAppBaseIpInterfaceAllOfWithDefaults() *StorageNetAppBaseIpInterfaceAllOf`

NewStorageNetAppBaseIpInterfaceAllOfWithDefaults instantiates a new StorageNetAppBaseIpInterfaceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppBaseIpInterfaceAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppBaseIpInterfaceAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppBaseIpInterfaceAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppBaseIpInterfaceAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppBaseIpInterfaceAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppBaseIpInterfaceAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCurrentNode

`func (o *StorageNetAppBaseIpInterfaceAllOf) GetCurrentNode() string`

GetCurrentNode returns the CurrentNode field if non-nil, zero value otherwise.

### GetCurrentNodeOk

`func (o *StorageNetAppBaseIpInterfaceAllOf) GetCurrentNodeOk() (*string, bool)`

GetCurrentNodeOk returns a tuple with the CurrentNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentNode

`func (o *StorageNetAppBaseIpInterfaceAllOf) SetCurrentNode(v string)`

SetCurrentNode sets CurrentNode field to given value.

### HasCurrentNode

`func (o *StorageNetAppBaseIpInterfaceAllOf) HasCurrentNode() bool`

HasCurrentNode returns a boolean if a field has been set.

### GetCurrentPort

`func (o *StorageNetAppBaseIpInterfaceAllOf) GetCurrentPort() string`

GetCurrentPort returns the CurrentPort field if non-nil, zero value otherwise.

### GetCurrentPortOk

`func (o *StorageNetAppBaseIpInterfaceAllOf) GetCurrentPortOk() (*string, bool)`

GetCurrentPortOk returns a tuple with the CurrentPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPort

`func (o *StorageNetAppBaseIpInterfaceAllOf) SetCurrentPort(v string)`

SetCurrentPort sets CurrentPort field to given value.

### HasCurrentPort

`func (o *StorageNetAppBaseIpInterfaceAllOf) HasCurrentPort() bool`

HasCurrentPort returns a boolean if a field has been set.

### GetEnabled

`func (o *StorageNetAppBaseIpInterfaceAllOf) GetEnabled() string`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *StorageNetAppBaseIpInterfaceAllOf) GetEnabledOk() (*string, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *StorageNetAppBaseIpInterfaceAllOf) SetEnabled(v string)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *StorageNetAppBaseIpInterfaceAllOf) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHomeNode

`func (o *StorageNetAppBaseIpInterfaceAllOf) GetHomeNode() string`

GetHomeNode returns the HomeNode field if non-nil, zero value otherwise.

### GetHomeNodeOk

`func (o *StorageNetAppBaseIpInterfaceAllOf) GetHomeNodeOk() (*string, bool)`

GetHomeNodeOk returns a tuple with the HomeNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeNode

`func (o *StorageNetAppBaseIpInterfaceAllOf) SetHomeNode(v string)`

SetHomeNode sets HomeNode field to given value.

### HasHomeNode

`func (o *StorageNetAppBaseIpInterfaceAllOf) HasHomeNode() bool`

HasHomeNode returns a boolean if a field has been set.

### GetHomePort

`func (o *StorageNetAppBaseIpInterfaceAllOf) GetHomePort() string`

GetHomePort returns the HomePort field if non-nil, zero value otherwise.

### GetHomePortOk

`func (o *StorageNetAppBaseIpInterfaceAllOf) GetHomePortOk() (*string, bool)`

GetHomePortOk returns a tuple with the HomePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomePort

`func (o *StorageNetAppBaseIpInterfaceAllOf) SetHomePort(v string)`

SetHomePort sets HomePort field to given value.

### HasHomePort

`func (o *StorageNetAppBaseIpInterfaceAllOf) HasHomePort() bool`

HasHomePort returns a boolean if a field has been set.

### GetInterfaceIsHome

`func (o *StorageNetAppBaseIpInterfaceAllOf) GetInterfaceIsHome() string`

GetInterfaceIsHome returns the InterfaceIsHome field if non-nil, zero value otherwise.

### GetInterfaceIsHomeOk

`func (o *StorageNetAppBaseIpInterfaceAllOf) GetInterfaceIsHomeOk() (*string, bool)`

GetInterfaceIsHomeOk returns a tuple with the InterfaceIsHome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceIsHome

`func (o *StorageNetAppBaseIpInterfaceAllOf) SetInterfaceIsHome(v string)`

SetInterfaceIsHome sets InterfaceIsHome field to given value.

### HasInterfaceIsHome

`func (o *StorageNetAppBaseIpInterfaceAllOf) HasInterfaceIsHome() bool`

HasInterfaceIsHome returns a boolean if a field has been set.

### GetInterfaceState

`func (o *StorageNetAppBaseIpInterfaceAllOf) GetInterfaceState() string`

GetInterfaceState returns the InterfaceState field if non-nil, zero value otherwise.

### GetInterfaceStateOk

`func (o *StorageNetAppBaseIpInterfaceAllOf) GetInterfaceStateOk() (*string, bool)`

GetInterfaceStateOk returns a tuple with the InterfaceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceState

`func (o *StorageNetAppBaseIpInterfaceAllOf) SetInterfaceState(v string)`

SetInterfaceState sets InterfaceState field to given value.

### HasInterfaceState

`func (o *StorageNetAppBaseIpInterfaceAllOf) HasInterfaceState() bool`

HasInterfaceState returns a boolean if a field has been set.

### GetIpAddress

`func (o *StorageNetAppBaseIpInterfaceAllOf) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *StorageNetAppBaseIpInterfaceAllOf) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *StorageNetAppBaseIpInterfaceAllOf) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *StorageNetAppBaseIpInterfaceAllOf) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetIpFamily

`func (o *StorageNetAppBaseIpInterfaceAllOf) GetIpFamily() string`

GetIpFamily returns the IpFamily field if non-nil, zero value otherwise.

### GetIpFamilyOk

`func (o *StorageNetAppBaseIpInterfaceAllOf) GetIpFamilyOk() (*string, bool)`

GetIpFamilyOk returns a tuple with the IpFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpFamily

`func (o *StorageNetAppBaseIpInterfaceAllOf) SetIpFamily(v string)`

SetIpFamily sets IpFamily field to given value.

### HasIpFamily

`func (o *StorageNetAppBaseIpInterfaceAllOf) HasIpFamily() bool`

HasIpFamily returns a boolean if a field has been set.

### GetIpspace

`func (o *StorageNetAppBaseIpInterfaceAllOf) GetIpspace() string`

GetIpspace returns the Ipspace field if non-nil, zero value otherwise.

### GetIpspaceOk

`func (o *StorageNetAppBaseIpInterfaceAllOf) GetIpspaceOk() (*string, bool)`

GetIpspaceOk returns a tuple with the Ipspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpspace

`func (o *StorageNetAppBaseIpInterfaceAllOf) SetIpspace(v string)`

SetIpspace sets Ipspace field to given value.

### HasIpspace

`func (o *StorageNetAppBaseIpInterfaceAllOf) HasIpspace() bool`

HasIpspace returns a boolean if a field has been set.

### GetIsHome

`func (o *StorageNetAppBaseIpInterfaceAllOf) GetIsHome() bool`

GetIsHome returns the IsHome field if non-nil, zero value otherwise.

### GetIsHomeOk

`func (o *StorageNetAppBaseIpInterfaceAllOf) GetIsHomeOk() (*bool, bool)`

GetIsHomeOk returns a tuple with the IsHome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHome

`func (o *StorageNetAppBaseIpInterfaceAllOf) SetIsHome(v bool)`

SetIsHome sets IsHome field to given value.

### HasIsHome

`func (o *StorageNetAppBaseIpInterfaceAllOf) HasIsHome() bool`

HasIsHome returns a boolean if a field has been set.

### GetLocationFailover

`func (o *StorageNetAppBaseIpInterfaceAllOf) GetLocationFailover() string`

GetLocationFailover returns the LocationFailover field if non-nil, zero value otherwise.

### GetLocationFailoverOk

`func (o *StorageNetAppBaseIpInterfaceAllOf) GetLocationFailoverOk() (*string, bool)`

GetLocationFailoverOk returns a tuple with the LocationFailover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationFailover

`func (o *StorageNetAppBaseIpInterfaceAllOf) SetLocationFailover(v string)`

SetLocationFailover sets LocationFailover field to given value.

### HasLocationFailover

`func (o *StorageNetAppBaseIpInterfaceAllOf) HasLocationFailover() bool`

HasLocationFailover returns a boolean if a field has been set.

### GetName

`func (o *StorageNetAppBaseIpInterfaceAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageNetAppBaseIpInterfaceAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageNetAppBaseIpInterfaceAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageNetAppBaseIpInterfaceAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetmask

`func (o *StorageNetAppBaseIpInterfaceAllOf) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *StorageNetAppBaseIpInterfaceAllOf) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *StorageNetAppBaseIpInterfaceAllOf) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *StorageNetAppBaseIpInterfaceAllOf) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetServicePolicyName

`func (o *StorageNetAppBaseIpInterfaceAllOf) GetServicePolicyName() string`

GetServicePolicyName returns the ServicePolicyName field if non-nil, zero value otherwise.

### GetServicePolicyNameOk

`func (o *StorageNetAppBaseIpInterfaceAllOf) GetServicePolicyNameOk() (*string, bool)`

GetServicePolicyNameOk returns a tuple with the ServicePolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePolicyName

`func (o *StorageNetAppBaseIpInterfaceAllOf) SetServicePolicyName(v string)`

SetServicePolicyName sets ServicePolicyName field to given value.

### HasServicePolicyName

`func (o *StorageNetAppBaseIpInterfaceAllOf) HasServicePolicyName() bool`

HasServicePolicyName returns a boolean if a field has been set.

### GetServicePolicyUuid

`func (o *StorageNetAppBaseIpInterfaceAllOf) GetServicePolicyUuid() string`

GetServicePolicyUuid returns the ServicePolicyUuid field if non-nil, zero value otherwise.

### GetServicePolicyUuidOk

`func (o *StorageNetAppBaseIpInterfaceAllOf) GetServicePolicyUuidOk() (*string, bool)`

GetServicePolicyUuidOk returns a tuple with the ServicePolicyUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePolicyUuid

`func (o *StorageNetAppBaseIpInterfaceAllOf) SetServicePolicyUuid(v string)`

SetServicePolicyUuid sets ServicePolicyUuid field to given value.

### HasServicePolicyUuid

`func (o *StorageNetAppBaseIpInterfaceAllOf) HasServicePolicyUuid() bool`

HasServicePolicyUuid returns a boolean if a field has been set.

### GetServices

`func (o *StorageNetAppBaseIpInterfaceAllOf) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *StorageNetAppBaseIpInterfaceAllOf) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *StorageNetAppBaseIpInterfaceAllOf) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *StorageNetAppBaseIpInterfaceAllOf) HasServices() bool`

HasServices returns a boolean if a field has been set.

### SetServicesNil

`func (o *StorageNetAppBaseIpInterfaceAllOf) SetServicesNil(b bool)`

 SetServicesNil sets the value for Services to be an explicit nil

### UnsetServices
`func (o *StorageNetAppBaseIpInterfaceAllOf) UnsetServices()`

UnsetServices ensures that no value is present for Services, not even an explicit nil
### GetState

`func (o *StorageNetAppBaseIpInterfaceAllOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StorageNetAppBaseIpInterfaceAllOf) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StorageNetAppBaseIpInterfaceAllOf) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *StorageNetAppBaseIpInterfaceAllOf) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSvmName

`func (o *StorageNetAppBaseIpInterfaceAllOf) GetSvmName() string`

GetSvmName returns the SvmName field if non-nil, zero value otherwise.

### GetSvmNameOk

`func (o *StorageNetAppBaseIpInterfaceAllOf) GetSvmNameOk() (*string, bool)`

GetSvmNameOk returns a tuple with the SvmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvmName

`func (o *StorageNetAppBaseIpInterfaceAllOf) SetSvmName(v string)`

SetSvmName sets SvmName field to given value.

### HasSvmName

`func (o *StorageNetAppBaseIpInterfaceAllOf) HasSvmName() bool`

HasSvmName returns a boolean if a field has been set.

### GetUuid

`func (o *StorageNetAppBaseIpInterfaceAllOf) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StorageNetAppBaseIpInterfaceAllOf) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StorageNetAppBaseIpInterfaceAllOf) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *StorageNetAppBaseIpInterfaceAllOf) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


