# HciManagementServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hci.ManagementServer"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hci.ManagementServer"]
**DrsEnabled** | Pointer to **bool** | Is DRS enabled on the management server. | [optional] [readonly] 
**InUse** | Pointer to **bool** | Indicates if the management server is in use. | [optional] [readonly] 
**Ip** | Pointer to [**NullableHciIpAddress**](HciIpAddress.md) |  | [optional] 
**IsRegistered** | Pointer to **bool** | Is the management server registered. | [optional] [readonly] 
**Type** | Pointer to **string** | The fully qualified domain name of the management server. | [optional] [readonly] 

## Methods

### NewHciManagementServer

`func NewHciManagementServer(classId string, objectType string, ) *HciManagementServer`

NewHciManagementServer instantiates a new HciManagementServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHciManagementServerWithDefaults

`func NewHciManagementServerWithDefaults() *HciManagementServer`

NewHciManagementServerWithDefaults instantiates a new HciManagementServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HciManagementServer) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HciManagementServer) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HciManagementServer) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HciManagementServer) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HciManagementServer) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HciManagementServer) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDrsEnabled

`func (o *HciManagementServer) GetDrsEnabled() bool`

GetDrsEnabled returns the DrsEnabled field if non-nil, zero value otherwise.

### GetDrsEnabledOk

`func (o *HciManagementServer) GetDrsEnabledOk() (*bool, bool)`

GetDrsEnabledOk returns a tuple with the DrsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrsEnabled

`func (o *HciManagementServer) SetDrsEnabled(v bool)`

SetDrsEnabled sets DrsEnabled field to given value.

### HasDrsEnabled

`func (o *HciManagementServer) HasDrsEnabled() bool`

HasDrsEnabled returns a boolean if a field has been set.

### GetInUse

`func (o *HciManagementServer) GetInUse() bool`

GetInUse returns the InUse field if non-nil, zero value otherwise.

### GetInUseOk

`func (o *HciManagementServer) GetInUseOk() (*bool, bool)`

GetInUseOk returns a tuple with the InUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUse

`func (o *HciManagementServer) SetInUse(v bool)`

SetInUse sets InUse field to given value.

### HasInUse

`func (o *HciManagementServer) HasInUse() bool`

HasInUse returns a boolean if a field has been set.

### GetIp

`func (o *HciManagementServer) GetIp() HciIpAddress`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *HciManagementServer) GetIpOk() (*HciIpAddress, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *HciManagementServer) SetIp(v HciIpAddress)`

SetIp sets Ip field to given value.

### HasIp

`func (o *HciManagementServer) HasIp() bool`

HasIp returns a boolean if a field has been set.

### SetIpNil

`func (o *HciManagementServer) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *HciManagementServer) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil
### GetIsRegistered

`func (o *HciManagementServer) GetIsRegistered() bool`

GetIsRegistered returns the IsRegistered field if non-nil, zero value otherwise.

### GetIsRegisteredOk

`func (o *HciManagementServer) GetIsRegisteredOk() (*bool, bool)`

GetIsRegisteredOk returns a tuple with the IsRegistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRegistered

`func (o *HciManagementServer) SetIsRegistered(v bool)`

SetIsRegistered sets IsRegistered field to given value.

### HasIsRegistered

`func (o *HciManagementServer) HasIsRegistered() bool`

HasIsRegistered returns a boolean if a field has been set.

### GetType

`func (o *HciManagementServer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HciManagementServer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HciManagementServer) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HciManagementServer) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


