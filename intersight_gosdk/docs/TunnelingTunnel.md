# TunnelingTunnel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ClientUrl** | Pointer to **string** | The multiplexer URL for the client to connect on. | [optional] [readonly] 

## Methods

### NewTunnelingTunnel

`func NewTunnelingTunnel(classId string, objectType string, ) *TunnelingTunnel`

NewTunnelingTunnel instantiates a new TunnelingTunnel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTunnelingTunnelWithDefaults

`func NewTunnelingTunnelWithDefaults() *TunnelingTunnel`

NewTunnelingTunnelWithDefaults instantiates a new TunnelingTunnel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TunnelingTunnel) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TunnelingTunnel) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TunnelingTunnel) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TunnelingTunnel) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TunnelingTunnel) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TunnelingTunnel) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClientUrl

`func (o *TunnelingTunnel) GetClientUrl() string`

GetClientUrl returns the ClientUrl field if non-nil, zero value otherwise.

### GetClientUrlOk

`func (o *TunnelingTunnel) GetClientUrlOk() (*string, bool)`

GetClientUrlOk returns a tuple with the ClientUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientUrl

`func (o *TunnelingTunnel) SetClientUrl(v string)`

SetClientUrl sets ClientUrl field to given value.

### HasClientUrl

`func (o *TunnelingTunnel) HasClientUrl() bool`

HasClientUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


