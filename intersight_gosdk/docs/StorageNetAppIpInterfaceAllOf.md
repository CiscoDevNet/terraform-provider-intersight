# StorageNetAppIpInterfaceAllOf

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

### NewStorageNetAppIpInterfaceAllOf

`func NewStorageNetAppIpInterfaceAllOf(classId string, objectType string, ) *StorageNetAppIpInterfaceAllOf`

NewStorageNetAppIpInterfaceAllOf instantiates a new StorageNetAppIpInterfaceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppIpInterfaceAllOfWithDefaults

`func NewStorageNetAppIpInterfaceAllOfWithDefaults() *StorageNetAppIpInterfaceAllOf`

NewStorageNetAppIpInterfaceAllOfWithDefaults instantiates a new StorageNetAppIpInterfaceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppIpInterfaceAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppIpInterfaceAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppIpInterfaceAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppIpInterfaceAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppIpInterfaceAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppIpInterfaceAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEnabled

`func (o *StorageNetAppIpInterfaceAllOf) GetEnabled() string`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *StorageNetAppIpInterfaceAllOf) GetEnabledOk() (*string, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *StorageNetAppIpInterfaceAllOf) SetEnabled(v string)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *StorageNetAppIpInterfaceAllOf) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHomeNode

`func (o *StorageNetAppIpInterfaceAllOf) GetHomeNode() string`

GetHomeNode returns the HomeNode field if non-nil, zero value otherwise.

### GetHomeNodeOk

`func (o *StorageNetAppIpInterfaceAllOf) GetHomeNodeOk() (*string, bool)`

GetHomeNodeOk returns a tuple with the HomeNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeNode

`func (o *StorageNetAppIpInterfaceAllOf) SetHomeNode(v string)`

SetHomeNode sets HomeNode field to given value.

### HasHomeNode

`func (o *StorageNetAppIpInterfaceAllOf) HasHomeNode() bool`

HasHomeNode returns a boolean if a field has been set.

### GetHomePort

`func (o *StorageNetAppIpInterfaceAllOf) GetHomePort() string`

GetHomePort returns the HomePort field if non-nil, zero value otherwise.

### GetHomePortOk

`func (o *StorageNetAppIpInterfaceAllOf) GetHomePortOk() (*string, bool)`

GetHomePortOk returns a tuple with the HomePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomePort

`func (o *StorageNetAppIpInterfaceAllOf) SetHomePort(v string)`

SetHomePort sets HomePort field to given value.

### HasHomePort

`func (o *StorageNetAppIpInterfaceAllOf) HasHomePort() bool`

HasHomePort returns a boolean if a field has been set.

### GetIpAddress

`func (o *StorageNetAppIpInterfaceAllOf) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *StorageNetAppIpInterfaceAllOf) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *StorageNetAppIpInterfaceAllOf) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *StorageNetAppIpInterfaceAllOf) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetIpFamily

`func (o *StorageNetAppIpInterfaceAllOf) GetIpFamily() string`

GetIpFamily returns the IpFamily field if non-nil, zero value otherwise.

### GetIpFamilyOk

`func (o *StorageNetAppIpInterfaceAllOf) GetIpFamilyOk() (*string, bool)`

GetIpFamilyOk returns a tuple with the IpFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpFamily

`func (o *StorageNetAppIpInterfaceAllOf) SetIpFamily(v string)`

SetIpFamily sets IpFamily field to given value.

### HasIpFamily

`func (o *StorageNetAppIpInterfaceAllOf) HasIpFamily() bool`

HasIpFamily returns a boolean if a field has been set.

### GetName

`func (o *StorageNetAppIpInterfaceAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageNetAppIpInterfaceAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageNetAppIpInterfaceAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageNetAppIpInterfaceAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetmask

`func (o *StorageNetAppIpInterfaceAllOf) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *StorageNetAppIpInterfaceAllOf) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *StorageNetAppIpInterfaceAllOf) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *StorageNetAppIpInterfaceAllOf) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetServicePolicyName

`func (o *StorageNetAppIpInterfaceAllOf) GetServicePolicyName() string`

GetServicePolicyName returns the ServicePolicyName field if non-nil, zero value otherwise.

### GetServicePolicyNameOk

`func (o *StorageNetAppIpInterfaceAllOf) GetServicePolicyNameOk() (*string, bool)`

GetServicePolicyNameOk returns a tuple with the ServicePolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePolicyName

`func (o *StorageNetAppIpInterfaceAllOf) SetServicePolicyName(v string)`

SetServicePolicyName sets ServicePolicyName field to given value.

### HasServicePolicyName

`func (o *StorageNetAppIpInterfaceAllOf) HasServicePolicyName() bool`

HasServicePolicyName returns a boolean if a field has been set.

### GetServicePolicyUuid

`func (o *StorageNetAppIpInterfaceAllOf) GetServicePolicyUuid() string`

GetServicePolicyUuid returns the ServicePolicyUuid field if non-nil, zero value otherwise.

### GetServicePolicyUuidOk

`func (o *StorageNetAppIpInterfaceAllOf) GetServicePolicyUuidOk() (*string, bool)`

GetServicePolicyUuidOk returns a tuple with the ServicePolicyUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePolicyUuid

`func (o *StorageNetAppIpInterfaceAllOf) SetServicePolicyUuid(v string)`

SetServicePolicyUuid sets ServicePolicyUuid field to given value.

### HasServicePolicyUuid

`func (o *StorageNetAppIpInterfaceAllOf) HasServicePolicyUuid() bool`

HasServicePolicyUuid returns a boolean if a field has been set.

### GetServices

`func (o *StorageNetAppIpInterfaceAllOf) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *StorageNetAppIpInterfaceAllOf) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *StorageNetAppIpInterfaceAllOf) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *StorageNetAppIpInterfaceAllOf) HasServices() bool`

HasServices returns a boolean if a field has been set.

### SetServicesNil

`func (o *StorageNetAppIpInterfaceAllOf) SetServicesNil(b bool)`

 SetServicesNil sets the value for Services to be an explicit nil

### UnsetServices
`func (o *StorageNetAppIpInterfaceAllOf) UnsetServices()`

UnsetServices ensures that no value is present for Services, not even an explicit nil
### GetState

`func (o *StorageNetAppIpInterfaceAllOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StorageNetAppIpInterfaceAllOf) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StorageNetAppIpInterfaceAllOf) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *StorageNetAppIpInterfaceAllOf) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUuid

`func (o *StorageNetAppIpInterfaceAllOf) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StorageNetAppIpInterfaceAllOf) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StorageNetAppIpInterfaceAllOf) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *StorageNetAppIpInterfaceAllOf) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetArrayController

`func (o *StorageNetAppIpInterfaceAllOf) GetArrayController() StorageNetAppNodeRelationship`

GetArrayController returns the ArrayController field if non-nil, zero value otherwise.

### GetArrayControllerOk

`func (o *StorageNetAppIpInterfaceAllOf) GetArrayControllerOk() (*StorageNetAppNodeRelationship, bool)`

GetArrayControllerOk returns a tuple with the ArrayController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayController

`func (o *StorageNetAppIpInterfaceAllOf) SetArrayController(v StorageNetAppNodeRelationship)`

SetArrayController sets ArrayController field to given value.

### HasArrayController

`func (o *StorageNetAppIpInterfaceAllOf) HasArrayController() bool`

HasArrayController returns a boolean if a field has been set.

### GetNetAppEthernetPort

`func (o *StorageNetAppIpInterfaceAllOf) GetNetAppEthernetPort() StorageNetAppEthernetPortRelationship`

GetNetAppEthernetPort returns the NetAppEthernetPort field if non-nil, zero value otherwise.

### GetNetAppEthernetPortOk

`func (o *StorageNetAppIpInterfaceAllOf) GetNetAppEthernetPortOk() (*StorageNetAppEthernetPortRelationship, bool)`

GetNetAppEthernetPortOk returns a tuple with the NetAppEthernetPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAppEthernetPort

`func (o *StorageNetAppIpInterfaceAllOf) SetNetAppEthernetPort(v StorageNetAppEthernetPortRelationship)`

SetNetAppEthernetPort sets NetAppEthernetPort field to given value.

### HasNetAppEthernetPort

`func (o *StorageNetAppIpInterfaceAllOf) HasNetAppEthernetPort() bool`

HasNetAppEthernetPort returns a boolean if a field has been set.

### GetTenant

`func (o *StorageNetAppIpInterfaceAllOf) GetTenant() StorageNetAppStorageVmRelationship`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *StorageNetAppIpInterfaceAllOf) GetTenantOk() (*StorageNetAppStorageVmRelationship, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *StorageNetAppIpInterfaceAllOf) SetTenant(v StorageNetAppStorageVmRelationship)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *StorageNetAppIpInterfaceAllOf) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


