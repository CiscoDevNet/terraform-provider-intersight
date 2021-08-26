# StorageNetAppIpInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppIpInterface"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppIpInterface"]
**Enabled** | Pointer to **string** | IP interface is enabled or not. | [optional] [readonly] 
**HomeNode** | Pointer to **string** | Name of home node of IP interface. | [optional] [readonly] 
**HomePort** | Pointer to **string** | Name of home port of IP interface. | [optional] [readonly] 
**IpAddress** | Pointer to **string** | IP address of inteface. | [optional] [readonly] 
**IpFamily** | Pointer to **string** | IP address family of inteface. * &#x60;IPv4&#x60; - IPv4 Address type. * &#x60;IPv6&#x60; - IPv6 Address type. | [optional] [readonly] [default to "IPv4"]
**Name** | Pointer to **string** | Name of IP interface. | [optional] [readonly] 
**Netmask** | Pointer to **string** | Netmask of Interface. | [optional] [readonly] 
**ServicePolicyName** | Pointer to **string** | Services of IP interface. | [optional] [readonly] 
**ServicePolicyUuid** | Pointer to **string** | Services of IP interface. | [optional] [readonly] 
**Services** | Pointer to **[]string** |  | [optional] 
**State** | Pointer to **string** | State of IP interface. * &#x60;down&#x60; - An inactive port is listed as Down. * &#x60;up&#x60; - An active port is listed as Up. * &#x60;present&#x60; - An active port is listed as present. | [optional] [readonly] [default to "down"]
**Uuid** | Pointer to **string** | Uuid of  NetApp IP Interface. | [optional] [readonly] 
**ArrayController** | Pointer to [**StorageNetAppNodeRelationship**](StorageNetAppNodeRelationship.md) |  | [optional] 
**NetAppEthernetPort** | Pointer to [**StorageNetAppEthernetPortRelationship**](StorageNetAppEthernetPortRelationship.md) |  | [optional] 
**Tenant** | Pointer to [**StorageNetAppStorageVmRelationship**](StorageNetAppStorageVmRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppIpInterface

`func NewStorageNetAppIpInterface(classId string, objectType string, ) *StorageNetAppIpInterface`

NewStorageNetAppIpInterface instantiates a new StorageNetAppIpInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppIpInterfaceWithDefaults

`func NewStorageNetAppIpInterfaceWithDefaults() *StorageNetAppIpInterface`

NewStorageNetAppIpInterfaceWithDefaults instantiates a new StorageNetAppIpInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppIpInterface) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppIpInterface) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppIpInterface) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppIpInterface) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppIpInterface) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppIpInterface) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEnabled

`func (o *StorageNetAppIpInterface) GetEnabled() string`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *StorageNetAppIpInterface) GetEnabledOk() (*string, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *StorageNetAppIpInterface) SetEnabled(v string)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *StorageNetAppIpInterface) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHomeNode

`func (o *StorageNetAppIpInterface) GetHomeNode() string`

GetHomeNode returns the HomeNode field if non-nil, zero value otherwise.

### GetHomeNodeOk

`func (o *StorageNetAppIpInterface) GetHomeNodeOk() (*string, bool)`

GetHomeNodeOk returns a tuple with the HomeNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeNode

`func (o *StorageNetAppIpInterface) SetHomeNode(v string)`

SetHomeNode sets HomeNode field to given value.

### HasHomeNode

`func (o *StorageNetAppIpInterface) HasHomeNode() bool`

HasHomeNode returns a boolean if a field has been set.

### GetHomePort

`func (o *StorageNetAppIpInterface) GetHomePort() string`

GetHomePort returns the HomePort field if non-nil, zero value otherwise.

### GetHomePortOk

`func (o *StorageNetAppIpInterface) GetHomePortOk() (*string, bool)`

GetHomePortOk returns a tuple with the HomePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomePort

`func (o *StorageNetAppIpInterface) SetHomePort(v string)`

SetHomePort sets HomePort field to given value.

### HasHomePort

`func (o *StorageNetAppIpInterface) HasHomePort() bool`

HasHomePort returns a boolean if a field has been set.

### GetIpAddress

`func (o *StorageNetAppIpInterface) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *StorageNetAppIpInterface) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *StorageNetAppIpInterface) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *StorageNetAppIpInterface) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetIpFamily

`func (o *StorageNetAppIpInterface) GetIpFamily() string`

GetIpFamily returns the IpFamily field if non-nil, zero value otherwise.

### GetIpFamilyOk

`func (o *StorageNetAppIpInterface) GetIpFamilyOk() (*string, bool)`

GetIpFamilyOk returns a tuple with the IpFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpFamily

`func (o *StorageNetAppIpInterface) SetIpFamily(v string)`

SetIpFamily sets IpFamily field to given value.

### HasIpFamily

`func (o *StorageNetAppIpInterface) HasIpFamily() bool`

HasIpFamily returns a boolean if a field has been set.

### GetName

`func (o *StorageNetAppIpInterface) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageNetAppIpInterface) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageNetAppIpInterface) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageNetAppIpInterface) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetmask

`func (o *StorageNetAppIpInterface) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *StorageNetAppIpInterface) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *StorageNetAppIpInterface) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *StorageNetAppIpInterface) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetServicePolicyName

`func (o *StorageNetAppIpInterface) GetServicePolicyName() string`

GetServicePolicyName returns the ServicePolicyName field if non-nil, zero value otherwise.

### GetServicePolicyNameOk

`func (o *StorageNetAppIpInterface) GetServicePolicyNameOk() (*string, bool)`

GetServicePolicyNameOk returns a tuple with the ServicePolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePolicyName

`func (o *StorageNetAppIpInterface) SetServicePolicyName(v string)`

SetServicePolicyName sets ServicePolicyName field to given value.

### HasServicePolicyName

`func (o *StorageNetAppIpInterface) HasServicePolicyName() bool`

HasServicePolicyName returns a boolean if a field has been set.

### GetServicePolicyUuid

`func (o *StorageNetAppIpInterface) GetServicePolicyUuid() string`

GetServicePolicyUuid returns the ServicePolicyUuid field if non-nil, zero value otherwise.

### GetServicePolicyUuidOk

`func (o *StorageNetAppIpInterface) GetServicePolicyUuidOk() (*string, bool)`

GetServicePolicyUuidOk returns a tuple with the ServicePolicyUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePolicyUuid

`func (o *StorageNetAppIpInterface) SetServicePolicyUuid(v string)`

SetServicePolicyUuid sets ServicePolicyUuid field to given value.

### HasServicePolicyUuid

`func (o *StorageNetAppIpInterface) HasServicePolicyUuid() bool`

HasServicePolicyUuid returns a boolean if a field has been set.

### GetServices

`func (o *StorageNetAppIpInterface) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *StorageNetAppIpInterface) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *StorageNetAppIpInterface) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *StorageNetAppIpInterface) HasServices() bool`

HasServices returns a boolean if a field has been set.

### SetServicesNil

`func (o *StorageNetAppIpInterface) SetServicesNil(b bool)`

 SetServicesNil sets the value for Services to be an explicit nil

### UnsetServices
`func (o *StorageNetAppIpInterface) UnsetServices()`

UnsetServices ensures that no value is present for Services, not even an explicit nil
### GetState

`func (o *StorageNetAppIpInterface) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StorageNetAppIpInterface) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StorageNetAppIpInterface) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *StorageNetAppIpInterface) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUuid

`func (o *StorageNetAppIpInterface) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StorageNetAppIpInterface) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StorageNetAppIpInterface) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *StorageNetAppIpInterface) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetArrayController

`func (o *StorageNetAppIpInterface) GetArrayController() StorageNetAppNodeRelationship`

GetArrayController returns the ArrayController field if non-nil, zero value otherwise.

### GetArrayControllerOk

`func (o *StorageNetAppIpInterface) GetArrayControllerOk() (*StorageNetAppNodeRelationship, bool)`

GetArrayControllerOk returns a tuple with the ArrayController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayController

`func (o *StorageNetAppIpInterface) SetArrayController(v StorageNetAppNodeRelationship)`

SetArrayController sets ArrayController field to given value.

### HasArrayController

`func (o *StorageNetAppIpInterface) HasArrayController() bool`

HasArrayController returns a boolean if a field has been set.

### GetNetAppEthernetPort

`func (o *StorageNetAppIpInterface) GetNetAppEthernetPort() StorageNetAppEthernetPortRelationship`

GetNetAppEthernetPort returns the NetAppEthernetPort field if non-nil, zero value otherwise.

### GetNetAppEthernetPortOk

`func (o *StorageNetAppIpInterface) GetNetAppEthernetPortOk() (*StorageNetAppEthernetPortRelationship, bool)`

GetNetAppEthernetPortOk returns a tuple with the NetAppEthernetPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAppEthernetPort

`func (o *StorageNetAppIpInterface) SetNetAppEthernetPort(v StorageNetAppEthernetPortRelationship)`

SetNetAppEthernetPort sets NetAppEthernetPort field to given value.

### HasNetAppEthernetPort

`func (o *StorageNetAppIpInterface) HasNetAppEthernetPort() bool`

HasNetAppEthernetPort returns a boolean if a field has been set.

### GetTenant

`func (o *StorageNetAppIpInterface) GetTenant() StorageNetAppStorageVmRelationship`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *StorageNetAppIpInterface) GetTenantOk() (*StorageNetAppStorageVmRelationship, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *StorageNetAppIpInterface) SetTenant(v StorageNetAppStorageVmRelationship)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *StorageNetAppIpInterface) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


