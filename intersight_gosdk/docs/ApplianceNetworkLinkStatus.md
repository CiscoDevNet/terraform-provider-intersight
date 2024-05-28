# ApplianceNetworkLinkStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.NetworkLinkStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.NetworkLinkStatus"]
**DestinationHostname** | Pointer to **string** | Hostname of the destination endpoint. | [optional] [readonly] 
**PingTime** | Pointer to **float32** | Time to reach the destination endpoint in milliseconds from the source endpoint. | [optional] [readonly] 
**SourceHostname** | Pointer to **string** | Hostname of the source endpoint. | [optional] [readonly] 
**NodeOpStatus** | Pointer to [**NullableApplianceNodeOpStatusRelationship**](ApplianceNodeOpStatusRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewApplianceNetworkLinkStatus

`func NewApplianceNetworkLinkStatus(classId string, objectType string, ) *ApplianceNetworkLinkStatus`

NewApplianceNetworkLinkStatus instantiates a new ApplianceNetworkLinkStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceNetworkLinkStatusWithDefaults

`func NewApplianceNetworkLinkStatusWithDefaults() *ApplianceNetworkLinkStatus`

NewApplianceNetworkLinkStatusWithDefaults instantiates a new ApplianceNetworkLinkStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceNetworkLinkStatus) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceNetworkLinkStatus) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceNetworkLinkStatus) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceNetworkLinkStatus) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceNetworkLinkStatus) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceNetworkLinkStatus) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDestinationHostname

`func (o *ApplianceNetworkLinkStatus) GetDestinationHostname() string`

GetDestinationHostname returns the DestinationHostname field if non-nil, zero value otherwise.

### GetDestinationHostnameOk

`func (o *ApplianceNetworkLinkStatus) GetDestinationHostnameOk() (*string, bool)`

GetDestinationHostnameOk returns a tuple with the DestinationHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationHostname

`func (o *ApplianceNetworkLinkStatus) SetDestinationHostname(v string)`

SetDestinationHostname sets DestinationHostname field to given value.

### HasDestinationHostname

`func (o *ApplianceNetworkLinkStatus) HasDestinationHostname() bool`

HasDestinationHostname returns a boolean if a field has been set.

### GetPingTime

`func (o *ApplianceNetworkLinkStatus) GetPingTime() float32`

GetPingTime returns the PingTime field if non-nil, zero value otherwise.

### GetPingTimeOk

`func (o *ApplianceNetworkLinkStatus) GetPingTimeOk() (*float32, bool)`

GetPingTimeOk returns a tuple with the PingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingTime

`func (o *ApplianceNetworkLinkStatus) SetPingTime(v float32)`

SetPingTime sets PingTime field to given value.

### HasPingTime

`func (o *ApplianceNetworkLinkStatus) HasPingTime() bool`

HasPingTime returns a boolean if a field has been set.

### GetSourceHostname

`func (o *ApplianceNetworkLinkStatus) GetSourceHostname() string`

GetSourceHostname returns the SourceHostname field if non-nil, zero value otherwise.

### GetSourceHostnameOk

`func (o *ApplianceNetworkLinkStatus) GetSourceHostnameOk() (*string, bool)`

GetSourceHostnameOk returns a tuple with the SourceHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceHostname

`func (o *ApplianceNetworkLinkStatus) SetSourceHostname(v string)`

SetSourceHostname sets SourceHostname field to given value.

### HasSourceHostname

`func (o *ApplianceNetworkLinkStatus) HasSourceHostname() bool`

HasSourceHostname returns a boolean if a field has been set.

### GetNodeOpStatus

`func (o *ApplianceNetworkLinkStatus) GetNodeOpStatus() ApplianceNodeOpStatusRelationship`

GetNodeOpStatus returns the NodeOpStatus field if non-nil, zero value otherwise.

### GetNodeOpStatusOk

`func (o *ApplianceNetworkLinkStatus) GetNodeOpStatusOk() (*ApplianceNodeOpStatusRelationship, bool)`

GetNodeOpStatusOk returns a tuple with the NodeOpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeOpStatus

`func (o *ApplianceNetworkLinkStatus) SetNodeOpStatus(v ApplianceNodeOpStatusRelationship)`

SetNodeOpStatus sets NodeOpStatus field to given value.

### HasNodeOpStatus

`func (o *ApplianceNetworkLinkStatus) HasNodeOpStatus() bool`

HasNodeOpStatus returns a boolean if a field has been set.

### SetNodeOpStatusNil

`func (o *ApplianceNetworkLinkStatus) SetNodeOpStatusNil(b bool)`

 SetNodeOpStatusNil sets the value for NodeOpStatus to be an explicit nil

### UnsetNodeOpStatus
`func (o *ApplianceNetworkLinkStatus) UnsetNodeOpStatus()`

UnsetNodeOpStatus ensures that no value is present for NodeOpStatus, not even an explicit nil
### GetRegisteredDevice

`func (o *ApplianceNetworkLinkStatus) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *ApplianceNetworkLinkStatus) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *ApplianceNetworkLinkStatus) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *ApplianceNetworkLinkStatus) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *ApplianceNetworkLinkStatus) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *ApplianceNetworkLinkStatus) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


