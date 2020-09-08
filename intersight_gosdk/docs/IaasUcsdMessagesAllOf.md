# IaasUcsdMessagesAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusId** | Pointer to **string** | Last checked time of the alerts. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewIaasUcsdMessagesAllOf

`func NewIaasUcsdMessagesAllOf() *IaasUcsdMessagesAllOf`

NewIaasUcsdMessagesAllOf instantiates a new IaasUcsdMessagesAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIaasUcsdMessagesAllOfWithDefaults

`func NewIaasUcsdMessagesAllOfWithDefaults() *IaasUcsdMessagesAllOf`

NewIaasUcsdMessagesAllOfWithDefaults instantiates a new IaasUcsdMessagesAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusId

`func (o *IaasUcsdMessagesAllOf) GetStatusId() string`

GetStatusId returns the StatusId field if non-nil, zero value otherwise.

### GetStatusIdOk

`func (o *IaasUcsdMessagesAllOf) GetStatusIdOk() (*string, bool)`

GetStatusIdOk returns a tuple with the StatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusId

`func (o *IaasUcsdMessagesAllOf) SetStatusId(v string)`

SetStatusId sets StatusId field to given value.

### HasStatusId

`func (o *IaasUcsdMessagesAllOf) HasStatusId() bool`

HasStatusId returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *IaasUcsdMessagesAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *IaasUcsdMessagesAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *IaasUcsdMessagesAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *IaasUcsdMessagesAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


