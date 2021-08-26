# StorageHitachiHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.HitachiHost"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.HitachiHost"]
**AuthenticationMode** | Pointer to **string** | Authentication mode for the iSCSI target. * &#x60;N/A&#x60; - Authentication Mode is not available. * &#x60;CHAP&#x60; - CHAP-authentication mode. * &#x60;NONE&#x60; - Authentication mode is not set. * &#x60;BOTH&#x60; - Comply with Host Setting. | [optional] [readonly] [default to "N/A"]
**HostGroupId** | Pointer to **string** | ID of the host group. | [optional] [readonly] 
**HostGroupNumber** | Pointer to **int64** | Host group number for this host. | [optional] 
**HostModeOptions** | Pointer to **[]int64** |  | [optional] 
**IsChapMutual** | Pointer to **bool** | Mutual CHAP setting that is Enabled or Disabled. | [optional] [readonly] 
**IscsiName** | Pointer to **string** | IQN (iSCSI qualified name). Can be up to 255 characters long. | [optional] [readonly] 
**PortId** | Pointer to **string** | Port ID of the host group. | [optional] [readonly] 
**PortLunSecurity** | Pointer to **bool** | LUN security setting for the port. | [optional] [readonly] 
**Type** | Pointer to **string** | Host Group type, it can be FC or iSCSI. * &#x60;FC&#x60; - Port supports fibre channel protocol. * &#x60;iSCSI&#x60; - Port supports iSCSI protocol. * &#x60;FCoE&#x60; - Port supports fibre channel over ethernet protocol. | [optional] [readonly] [default to "FC"]
**Wwn** | Pointer to **string** | World wide port name, 64 bit address represented in hexa decimal notation. | [optional] [readonly] 
**Array** | Pointer to [**StorageHitachiArrayRelationship**](StorageHitachiArrayRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewStorageHitachiHost

`func NewStorageHitachiHost(classId string, objectType string, ) *StorageHitachiHost`

NewStorageHitachiHost instantiates a new StorageHitachiHost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageHitachiHostWithDefaults

`func NewStorageHitachiHostWithDefaults() *StorageHitachiHost`

NewStorageHitachiHostWithDefaults instantiates a new StorageHitachiHost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageHitachiHost) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageHitachiHost) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageHitachiHost) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageHitachiHost) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageHitachiHost) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageHitachiHost) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAuthenticationMode

`func (o *StorageHitachiHost) GetAuthenticationMode() string`

GetAuthenticationMode returns the AuthenticationMode field if non-nil, zero value otherwise.

### GetAuthenticationModeOk

`func (o *StorageHitachiHost) GetAuthenticationModeOk() (*string, bool)`

GetAuthenticationModeOk returns a tuple with the AuthenticationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMode

`func (o *StorageHitachiHost) SetAuthenticationMode(v string)`

SetAuthenticationMode sets AuthenticationMode field to given value.

### HasAuthenticationMode

`func (o *StorageHitachiHost) HasAuthenticationMode() bool`

HasAuthenticationMode returns a boolean if a field has been set.

### GetHostGroupId

`func (o *StorageHitachiHost) GetHostGroupId() string`

GetHostGroupId returns the HostGroupId field if non-nil, zero value otherwise.

### GetHostGroupIdOk

`func (o *StorageHitachiHost) GetHostGroupIdOk() (*string, bool)`

GetHostGroupIdOk returns a tuple with the HostGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostGroupId

`func (o *StorageHitachiHost) SetHostGroupId(v string)`

SetHostGroupId sets HostGroupId field to given value.

### HasHostGroupId

`func (o *StorageHitachiHost) HasHostGroupId() bool`

HasHostGroupId returns a boolean if a field has been set.

### GetHostGroupNumber

`func (o *StorageHitachiHost) GetHostGroupNumber() int64`

GetHostGroupNumber returns the HostGroupNumber field if non-nil, zero value otherwise.

### GetHostGroupNumberOk

`func (o *StorageHitachiHost) GetHostGroupNumberOk() (*int64, bool)`

GetHostGroupNumberOk returns a tuple with the HostGroupNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostGroupNumber

`func (o *StorageHitachiHost) SetHostGroupNumber(v int64)`

SetHostGroupNumber sets HostGroupNumber field to given value.

### HasHostGroupNumber

`func (o *StorageHitachiHost) HasHostGroupNumber() bool`

HasHostGroupNumber returns a boolean if a field has been set.

### GetHostModeOptions

`func (o *StorageHitachiHost) GetHostModeOptions() []int64`

GetHostModeOptions returns the HostModeOptions field if non-nil, zero value otherwise.

### GetHostModeOptionsOk

`func (o *StorageHitachiHost) GetHostModeOptionsOk() (*[]int64, bool)`

GetHostModeOptionsOk returns a tuple with the HostModeOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostModeOptions

`func (o *StorageHitachiHost) SetHostModeOptions(v []int64)`

SetHostModeOptions sets HostModeOptions field to given value.

### HasHostModeOptions

`func (o *StorageHitachiHost) HasHostModeOptions() bool`

HasHostModeOptions returns a boolean if a field has been set.

### SetHostModeOptionsNil

`func (o *StorageHitachiHost) SetHostModeOptionsNil(b bool)`

 SetHostModeOptionsNil sets the value for HostModeOptions to be an explicit nil

### UnsetHostModeOptions
`func (o *StorageHitachiHost) UnsetHostModeOptions()`

UnsetHostModeOptions ensures that no value is present for HostModeOptions, not even an explicit nil
### GetIsChapMutual

`func (o *StorageHitachiHost) GetIsChapMutual() bool`

GetIsChapMutual returns the IsChapMutual field if non-nil, zero value otherwise.

### GetIsChapMutualOk

`func (o *StorageHitachiHost) GetIsChapMutualOk() (*bool, bool)`

GetIsChapMutualOk returns a tuple with the IsChapMutual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsChapMutual

`func (o *StorageHitachiHost) SetIsChapMutual(v bool)`

SetIsChapMutual sets IsChapMutual field to given value.

### HasIsChapMutual

`func (o *StorageHitachiHost) HasIsChapMutual() bool`

HasIsChapMutual returns a boolean if a field has been set.

### GetIscsiName

`func (o *StorageHitachiHost) GetIscsiName() string`

GetIscsiName returns the IscsiName field if non-nil, zero value otherwise.

### GetIscsiNameOk

`func (o *StorageHitachiHost) GetIscsiNameOk() (*string, bool)`

GetIscsiNameOk returns a tuple with the IscsiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiName

`func (o *StorageHitachiHost) SetIscsiName(v string)`

SetIscsiName sets IscsiName field to given value.

### HasIscsiName

`func (o *StorageHitachiHost) HasIscsiName() bool`

HasIscsiName returns a boolean if a field has been set.

### GetPortId

`func (o *StorageHitachiHost) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *StorageHitachiHost) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *StorageHitachiHost) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *StorageHitachiHost) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetPortLunSecurity

`func (o *StorageHitachiHost) GetPortLunSecurity() bool`

GetPortLunSecurity returns the PortLunSecurity field if non-nil, zero value otherwise.

### GetPortLunSecurityOk

`func (o *StorageHitachiHost) GetPortLunSecurityOk() (*bool, bool)`

GetPortLunSecurityOk returns a tuple with the PortLunSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortLunSecurity

`func (o *StorageHitachiHost) SetPortLunSecurity(v bool)`

SetPortLunSecurity sets PortLunSecurity field to given value.

### HasPortLunSecurity

`func (o *StorageHitachiHost) HasPortLunSecurity() bool`

HasPortLunSecurity returns a boolean if a field has been set.

### GetType

`func (o *StorageHitachiHost) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StorageHitachiHost) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StorageHitachiHost) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StorageHitachiHost) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWwn

`func (o *StorageHitachiHost) GetWwn() string`

GetWwn returns the Wwn field if non-nil, zero value otherwise.

### GetWwnOk

`func (o *StorageHitachiHost) GetWwnOk() (*string, bool)`

GetWwnOk returns a tuple with the Wwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwn

`func (o *StorageHitachiHost) SetWwn(v string)`

SetWwn sets Wwn field to given value.

### HasWwn

`func (o *StorageHitachiHost) HasWwn() bool`

HasWwn returns a boolean if a field has been set.

### GetArray

`func (o *StorageHitachiHost) GetArray() StorageHitachiArrayRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StorageHitachiHost) GetArrayOk() (*StorageHitachiArrayRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StorageHitachiHost) SetArray(v StorageHitachiArrayRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StorageHitachiHost) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageHitachiHost) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageHitachiHost) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageHitachiHost) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageHitachiHost) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


