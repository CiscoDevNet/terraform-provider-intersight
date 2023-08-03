# ApplianceNetworkStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.NetworkStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.NetworkStatus"]
**DestinationHostname** | Pointer to **string** | Hostname of the destination endpoint. | [optional] [readonly] 
**PingTime** | Pointer to **float32** | Time to reach the destination endpoint in milliseconds from the source endpoint. | [optional] [readonly] 
**SourceHostname** | Pointer to **string** | Hostname of the source endpoint. | [optional] [readonly] 

## Methods

### NewApplianceNetworkStatus

`func NewApplianceNetworkStatus(classId string, objectType string, ) *ApplianceNetworkStatus`

NewApplianceNetworkStatus instantiates a new ApplianceNetworkStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceNetworkStatusWithDefaults

`func NewApplianceNetworkStatusWithDefaults() *ApplianceNetworkStatus`

NewApplianceNetworkStatusWithDefaults instantiates a new ApplianceNetworkStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceNetworkStatus) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceNetworkStatus) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceNetworkStatus) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceNetworkStatus) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceNetworkStatus) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceNetworkStatus) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDestinationHostname

`func (o *ApplianceNetworkStatus) GetDestinationHostname() string`

GetDestinationHostname returns the DestinationHostname field if non-nil, zero value otherwise.

### GetDestinationHostnameOk

`func (o *ApplianceNetworkStatus) GetDestinationHostnameOk() (*string, bool)`

GetDestinationHostnameOk returns a tuple with the DestinationHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationHostname

`func (o *ApplianceNetworkStatus) SetDestinationHostname(v string)`

SetDestinationHostname sets DestinationHostname field to given value.

### HasDestinationHostname

`func (o *ApplianceNetworkStatus) HasDestinationHostname() bool`

HasDestinationHostname returns a boolean if a field has been set.

### GetPingTime

`func (o *ApplianceNetworkStatus) GetPingTime() float32`

GetPingTime returns the PingTime field if non-nil, zero value otherwise.

### GetPingTimeOk

`func (o *ApplianceNetworkStatus) GetPingTimeOk() (*float32, bool)`

GetPingTimeOk returns a tuple with the PingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingTime

`func (o *ApplianceNetworkStatus) SetPingTime(v float32)`

SetPingTime sets PingTime field to given value.

### HasPingTime

`func (o *ApplianceNetworkStatus) HasPingTime() bool`

HasPingTime returns a boolean if a field has been set.

### GetSourceHostname

`func (o *ApplianceNetworkStatus) GetSourceHostname() string`

GetSourceHostname returns the SourceHostname field if non-nil, zero value otherwise.

### GetSourceHostnameOk

`func (o *ApplianceNetworkStatus) GetSourceHostnameOk() (*string, bool)`

GetSourceHostnameOk returns a tuple with the SourceHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceHostname

`func (o *ApplianceNetworkStatus) SetSourceHostname(v string)`

SetSourceHostname sets SourceHostname field to given value.

### HasSourceHostname

`func (o *ApplianceNetworkStatus) HasSourceHostname() bool`

HasSourceHostname returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


