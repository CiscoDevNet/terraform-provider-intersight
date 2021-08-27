# IaasUcsdMessages

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iaas.UcsdMessages"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iaas.UcsdMessages"]
**StatusId** | Pointer to **string** | Last checked time of the alerts. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewIaasUcsdMessages

`func NewIaasUcsdMessages(classId string, objectType string, ) *IaasUcsdMessages`

NewIaasUcsdMessages instantiates a new IaasUcsdMessages object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIaasUcsdMessagesWithDefaults

`func NewIaasUcsdMessagesWithDefaults() *IaasUcsdMessages`

NewIaasUcsdMessagesWithDefaults instantiates a new IaasUcsdMessages object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IaasUcsdMessages) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IaasUcsdMessages) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IaasUcsdMessages) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IaasUcsdMessages) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IaasUcsdMessages) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IaasUcsdMessages) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetStatusId

`func (o *IaasUcsdMessages) GetStatusId() string`

GetStatusId returns the StatusId field if non-nil, zero value otherwise.

### GetStatusIdOk

`func (o *IaasUcsdMessages) GetStatusIdOk() (*string, bool)`

GetStatusIdOk returns a tuple with the StatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusId

`func (o *IaasUcsdMessages) SetStatusId(v string)`

SetStatusId sets StatusId field to given value.

### HasStatusId

`func (o *IaasUcsdMessages) HasStatusId() bool`

HasStatusId returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *IaasUcsdMessages) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *IaasUcsdMessages) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *IaasUcsdMessages) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *IaasUcsdMessages) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


