# HciKeyManagementDeviceToCertStatusInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hci.KeyManagementDeviceToCertStatusInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hci.KeyManagementDeviceToCertStatusInfo"]
**KeyManagementServerName** | Pointer to **string** | The name of the key management server. | [optional] [readonly] 
**Status** | Pointer to **bool** | The status of the certificate. | [optional] [readonly] 

## Methods

### NewHciKeyManagementDeviceToCertStatusInfo

`func NewHciKeyManagementDeviceToCertStatusInfo(classId string, objectType string, ) *HciKeyManagementDeviceToCertStatusInfo`

NewHciKeyManagementDeviceToCertStatusInfo instantiates a new HciKeyManagementDeviceToCertStatusInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHciKeyManagementDeviceToCertStatusInfoWithDefaults

`func NewHciKeyManagementDeviceToCertStatusInfoWithDefaults() *HciKeyManagementDeviceToCertStatusInfo`

NewHciKeyManagementDeviceToCertStatusInfoWithDefaults instantiates a new HciKeyManagementDeviceToCertStatusInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HciKeyManagementDeviceToCertStatusInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HciKeyManagementDeviceToCertStatusInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HciKeyManagementDeviceToCertStatusInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HciKeyManagementDeviceToCertStatusInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HciKeyManagementDeviceToCertStatusInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HciKeyManagementDeviceToCertStatusInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetKeyManagementServerName

`func (o *HciKeyManagementDeviceToCertStatusInfo) GetKeyManagementServerName() string`

GetKeyManagementServerName returns the KeyManagementServerName field if non-nil, zero value otherwise.

### GetKeyManagementServerNameOk

`func (o *HciKeyManagementDeviceToCertStatusInfo) GetKeyManagementServerNameOk() (*string, bool)`

GetKeyManagementServerNameOk returns a tuple with the KeyManagementServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyManagementServerName

`func (o *HciKeyManagementDeviceToCertStatusInfo) SetKeyManagementServerName(v string)`

SetKeyManagementServerName sets KeyManagementServerName field to given value.

### HasKeyManagementServerName

`func (o *HciKeyManagementDeviceToCertStatusInfo) HasKeyManagementServerName() bool`

HasKeyManagementServerName returns a boolean if a field has been set.

### GetStatus

`func (o *HciKeyManagementDeviceToCertStatusInfo) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HciKeyManagementDeviceToCertStatusInfo) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HciKeyManagementDeviceToCertStatusInfo) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HciKeyManagementDeviceToCertStatusInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


