# NtpNtpServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "ntp.NtpServer"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "ntp.NtpServer"]
**Poll** | Pointer to **int64** | The default polling interval of the NTP server in seconds. | [optional] [readonly] 
**ServerIpAddress** | Pointer to **string** | The IP address of the NTP server. Supports both IPv4 or IPv6 address. | [optional] [readonly] 
**Stratum** | Pointer to **int64** | The stratum level of the NTP server. | [optional] [readonly] 
**Type** | Pointer to **string** | It determines whether the IP address configured is server or peer. * &#x60;Server&#x60; - NTP configured is server type. * &#x60;Peer&#x60; - NTP configured is peer type. | [optional] [readonly] [default to "Server"]
**NetworkElement** | Pointer to [**NetworkElementRelationship**](NetworkElementRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNtpNtpServer

`func NewNtpNtpServer(classId string, objectType string, ) *NtpNtpServer`

NewNtpNtpServer instantiates a new NtpNtpServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNtpNtpServerWithDefaults

`func NewNtpNtpServerWithDefaults() *NtpNtpServer`

NewNtpNtpServerWithDefaults instantiates a new NtpNtpServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NtpNtpServer) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NtpNtpServer) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NtpNtpServer) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NtpNtpServer) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NtpNtpServer) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NtpNtpServer) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPoll

`func (o *NtpNtpServer) GetPoll() int64`

GetPoll returns the Poll field if non-nil, zero value otherwise.

### GetPollOk

`func (o *NtpNtpServer) GetPollOk() (*int64, bool)`

GetPollOk returns a tuple with the Poll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoll

`func (o *NtpNtpServer) SetPoll(v int64)`

SetPoll sets Poll field to given value.

### HasPoll

`func (o *NtpNtpServer) HasPoll() bool`

HasPoll returns a boolean if a field has been set.

### GetServerIpAddress

`func (o *NtpNtpServer) GetServerIpAddress() string`

GetServerIpAddress returns the ServerIpAddress field if non-nil, zero value otherwise.

### GetServerIpAddressOk

`func (o *NtpNtpServer) GetServerIpAddressOk() (*string, bool)`

GetServerIpAddressOk returns a tuple with the ServerIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIpAddress

`func (o *NtpNtpServer) SetServerIpAddress(v string)`

SetServerIpAddress sets ServerIpAddress field to given value.

### HasServerIpAddress

`func (o *NtpNtpServer) HasServerIpAddress() bool`

HasServerIpAddress returns a boolean if a field has been set.

### GetStratum

`func (o *NtpNtpServer) GetStratum() int64`

GetStratum returns the Stratum field if non-nil, zero value otherwise.

### GetStratumOk

`func (o *NtpNtpServer) GetStratumOk() (*int64, bool)`

GetStratumOk returns a tuple with the Stratum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStratum

`func (o *NtpNtpServer) SetStratum(v int64)`

SetStratum sets Stratum field to given value.

### HasStratum

`func (o *NtpNtpServer) HasStratum() bool`

HasStratum returns a boolean if a field has been set.

### GetType

`func (o *NtpNtpServer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NtpNtpServer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NtpNtpServer) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NtpNtpServer) HasType() bool`

HasType returns a boolean if a field has been set.

### GetNetworkElement

`func (o *NtpNtpServer) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *NtpNtpServer) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *NtpNtpServer) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *NtpNtpServer) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NtpNtpServer) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NtpNtpServer) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NtpNtpServer) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NtpNtpServer) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


