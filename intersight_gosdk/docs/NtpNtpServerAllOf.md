# NtpNtpServerAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "ntp.NtpServer"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "ntp.NtpServer"]
**Poll** | Pointer to **int64** | The default polling interval of the NTP server in seconds. | [optional] [readonly] 
**ServerIpAddress** | Pointer to **string** | The IP address of the NTP server. Supports both IPv4 or IPv6 address. | [optional] [readonly] 
**Stratum** | Pointer to **int64** | The stratum level of the NTP server. | [optional] [readonly] 
**Type** | Pointer to **string** | It determines whether the IP address configured is server or peer. * &#x60;Server&#x60; - NTP configured is server type. * &#x60;Peer&#x60; - NTP configured is peer type. | [optional] [readonly] [default to "Server"]
**VrfName** | Pointer to **string** | VRF name to be used by NTP Server. | [optional] [readonly] 
**NetworkElement** | Pointer to [**NetworkElementRelationship**](NetworkElementRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNtpNtpServerAllOf

`func NewNtpNtpServerAllOf(classId string, objectType string, ) *NtpNtpServerAllOf`

NewNtpNtpServerAllOf instantiates a new NtpNtpServerAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNtpNtpServerAllOfWithDefaults

`func NewNtpNtpServerAllOfWithDefaults() *NtpNtpServerAllOf`

NewNtpNtpServerAllOfWithDefaults instantiates a new NtpNtpServerAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NtpNtpServerAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NtpNtpServerAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NtpNtpServerAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NtpNtpServerAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NtpNtpServerAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NtpNtpServerAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPoll

`func (o *NtpNtpServerAllOf) GetPoll() int64`

GetPoll returns the Poll field if non-nil, zero value otherwise.

### GetPollOk

`func (o *NtpNtpServerAllOf) GetPollOk() (*int64, bool)`

GetPollOk returns a tuple with the Poll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoll

`func (o *NtpNtpServerAllOf) SetPoll(v int64)`

SetPoll sets Poll field to given value.

### HasPoll

`func (o *NtpNtpServerAllOf) HasPoll() bool`

HasPoll returns a boolean if a field has been set.

### GetServerIpAddress

`func (o *NtpNtpServerAllOf) GetServerIpAddress() string`

GetServerIpAddress returns the ServerIpAddress field if non-nil, zero value otherwise.

### GetServerIpAddressOk

`func (o *NtpNtpServerAllOf) GetServerIpAddressOk() (*string, bool)`

GetServerIpAddressOk returns a tuple with the ServerIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIpAddress

`func (o *NtpNtpServerAllOf) SetServerIpAddress(v string)`

SetServerIpAddress sets ServerIpAddress field to given value.

### HasServerIpAddress

`func (o *NtpNtpServerAllOf) HasServerIpAddress() bool`

HasServerIpAddress returns a boolean if a field has been set.

### GetStratum

`func (o *NtpNtpServerAllOf) GetStratum() int64`

GetStratum returns the Stratum field if non-nil, zero value otherwise.

### GetStratumOk

`func (o *NtpNtpServerAllOf) GetStratumOk() (*int64, bool)`

GetStratumOk returns a tuple with the Stratum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStratum

`func (o *NtpNtpServerAllOf) SetStratum(v int64)`

SetStratum sets Stratum field to given value.

### HasStratum

`func (o *NtpNtpServerAllOf) HasStratum() bool`

HasStratum returns a boolean if a field has been set.

### GetType

`func (o *NtpNtpServerAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NtpNtpServerAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NtpNtpServerAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NtpNtpServerAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVrfName

`func (o *NtpNtpServerAllOf) GetVrfName() string`

GetVrfName returns the VrfName field if non-nil, zero value otherwise.

### GetVrfNameOk

`func (o *NtpNtpServerAllOf) GetVrfNameOk() (*string, bool)`

GetVrfNameOk returns a tuple with the VrfName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfName

`func (o *NtpNtpServerAllOf) SetVrfName(v string)`

SetVrfName sets VrfName field to given value.

### HasVrfName

`func (o *NtpNtpServerAllOf) HasVrfName() bool`

HasVrfName returns a boolean if a field has been set.

### GetNetworkElement

`func (o *NtpNtpServerAllOf) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *NtpNtpServerAllOf) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *NtpNtpServerAllOf) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *NtpNtpServerAllOf) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NtpNtpServerAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NtpNtpServerAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NtpNtpServerAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NtpNtpServerAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


