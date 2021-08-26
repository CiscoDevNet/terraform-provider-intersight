# StorageHitachiHostAllOf

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

### NewStorageHitachiHostAllOf

`func NewStorageHitachiHostAllOf(classId string, objectType string, ) *StorageHitachiHostAllOf`

NewStorageHitachiHostAllOf instantiates a new StorageHitachiHostAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageHitachiHostAllOfWithDefaults

`func NewStorageHitachiHostAllOfWithDefaults() *StorageHitachiHostAllOf`

NewStorageHitachiHostAllOfWithDefaults instantiates a new StorageHitachiHostAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageHitachiHostAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageHitachiHostAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageHitachiHostAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageHitachiHostAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageHitachiHostAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageHitachiHostAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAuthenticationMode

`func (o *StorageHitachiHostAllOf) GetAuthenticationMode() string`

GetAuthenticationMode returns the AuthenticationMode field if non-nil, zero value otherwise.

### GetAuthenticationModeOk

`func (o *StorageHitachiHostAllOf) GetAuthenticationModeOk() (*string, bool)`

GetAuthenticationModeOk returns a tuple with the AuthenticationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMode

`func (o *StorageHitachiHostAllOf) SetAuthenticationMode(v string)`

SetAuthenticationMode sets AuthenticationMode field to given value.

### HasAuthenticationMode

`func (o *StorageHitachiHostAllOf) HasAuthenticationMode() bool`

HasAuthenticationMode returns a boolean if a field has been set.

### GetHostGroupId

`func (o *StorageHitachiHostAllOf) GetHostGroupId() string`

GetHostGroupId returns the HostGroupId field if non-nil, zero value otherwise.

### GetHostGroupIdOk

`func (o *StorageHitachiHostAllOf) GetHostGroupIdOk() (*string, bool)`

GetHostGroupIdOk returns a tuple with the HostGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostGroupId

`func (o *StorageHitachiHostAllOf) SetHostGroupId(v string)`

SetHostGroupId sets HostGroupId field to given value.

### HasHostGroupId

`func (o *StorageHitachiHostAllOf) HasHostGroupId() bool`

HasHostGroupId returns a boolean if a field has been set.

### GetHostGroupNumber

`func (o *StorageHitachiHostAllOf) GetHostGroupNumber() int64`

GetHostGroupNumber returns the HostGroupNumber field if non-nil, zero value otherwise.

### GetHostGroupNumberOk

`func (o *StorageHitachiHostAllOf) GetHostGroupNumberOk() (*int64, bool)`

GetHostGroupNumberOk returns a tuple with the HostGroupNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostGroupNumber

`func (o *StorageHitachiHostAllOf) SetHostGroupNumber(v int64)`

SetHostGroupNumber sets HostGroupNumber field to given value.

### HasHostGroupNumber

`func (o *StorageHitachiHostAllOf) HasHostGroupNumber() bool`

HasHostGroupNumber returns a boolean if a field has been set.

### GetHostModeOptions

`func (o *StorageHitachiHostAllOf) GetHostModeOptions() []int64`

GetHostModeOptions returns the HostModeOptions field if non-nil, zero value otherwise.

### GetHostModeOptionsOk

`func (o *StorageHitachiHostAllOf) GetHostModeOptionsOk() (*[]int64, bool)`

GetHostModeOptionsOk returns a tuple with the HostModeOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostModeOptions

`func (o *StorageHitachiHostAllOf) SetHostModeOptions(v []int64)`

SetHostModeOptions sets HostModeOptions field to given value.

### HasHostModeOptions

`func (o *StorageHitachiHostAllOf) HasHostModeOptions() bool`

HasHostModeOptions returns a boolean if a field has been set.

### SetHostModeOptionsNil

`func (o *StorageHitachiHostAllOf) SetHostModeOptionsNil(b bool)`

 SetHostModeOptionsNil sets the value for HostModeOptions to be an explicit nil

### UnsetHostModeOptions
`func (o *StorageHitachiHostAllOf) UnsetHostModeOptions()`

UnsetHostModeOptions ensures that no value is present for HostModeOptions, not even an explicit nil
### GetIsChapMutual

`func (o *StorageHitachiHostAllOf) GetIsChapMutual() bool`

GetIsChapMutual returns the IsChapMutual field if non-nil, zero value otherwise.

### GetIsChapMutualOk

`func (o *StorageHitachiHostAllOf) GetIsChapMutualOk() (*bool, bool)`

GetIsChapMutualOk returns a tuple with the IsChapMutual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsChapMutual

`func (o *StorageHitachiHostAllOf) SetIsChapMutual(v bool)`

SetIsChapMutual sets IsChapMutual field to given value.

### HasIsChapMutual

`func (o *StorageHitachiHostAllOf) HasIsChapMutual() bool`

HasIsChapMutual returns a boolean if a field has been set.

### GetIscsiName

`func (o *StorageHitachiHostAllOf) GetIscsiName() string`

GetIscsiName returns the IscsiName field if non-nil, zero value otherwise.

### GetIscsiNameOk

`func (o *StorageHitachiHostAllOf) GetIscsiNameOk() (*string, bool)`

GetIscsiNameOk returns a tuple with the IscsiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiName

`func (o *StorageHitachiHostAllOf) SetIscsiName(v string)`

SetIscsiName sets IscsiName field to given value.

### HasIscsiName

`func (o *StorageHitachiHostAllOf) HasIscsiName() bool`

HasIscsiName returns a boolean if a field has been set.

### GetPortId

`func (o *StorageHitachiHostAllOf) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *StorageHitachiHostAllOf) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *StorageHitachiHostAllOf) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *StorageHitachiHostAllOf) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetPortLunSecurity

`func (o *StorageHitachiHostAllOf) GetPortLunSecurity() bool`

GetPortLunSecurity returns the PortLunSecurity field if non-nil, zero value otherwise.

### GetPortLunSecurityOk

`func (o *StorageHitachiHostAllOf) GetPortLunSecurityOk() (*bool, bool)`

GetPortLunSecurityOk returns a tuple with the PortLunSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortLunSecurity

`func (o *StorageHitachiHostAllOf) SetPortLunSecurity(v bool)`

SetPortLunSecurity sets PortLunSecurity field to given value.

### HasPortLunSecurity

`func (o *StorageHitachiHostAllOf) HasPortLunSecurity() bool`

HasPortLunSecurity returns a boolean if a field has been set.

### GetType

`func (o *StorageHitachiHostAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StorageHitachiHostAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StorageHitachiHostAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StorageHitachiHostAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWwn

`func (o *StorageHitachiHostAllOf) GetWwn() string`

GetWwn returns the Wwn field if non-nil, zero value otherwise.

### GetWwnOk

`func (o *StorageHitachiHostAllOf) GetWwnOk() (*string, bool)`

GetWwnOk returns a tuple with the Wwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwn

`func (o *StorageHitachiHostAllOf) SetWwn(v string)`

SetWwn sets Wwn field to given value.

### HasWwn

`func (o *StorageHitachiHostAllOf) HasWwn() bool`

HasWwn returns a boolean if a field has been set.

### GetArray

`func (o *StorageHitachiHostAllOf) GetArray() StorageHitachiArrayRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StorageHitachiHostAllOf) GetArrayOk() (*StorageHitachiArrayRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StorageHitachiHostAllOf) SetArray(v StorageHitachiArrayRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StorageHitachiHostAllOf) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageHitachiHostAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageHitachiHostAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageHitachiHostAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageHitachiHostAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


