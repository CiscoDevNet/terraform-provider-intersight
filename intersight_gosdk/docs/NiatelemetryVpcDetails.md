# NiatelemetryVpcDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.VpcDetails"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.VpcDetails"]
**IsVpcConfigured** | Pointer to **bool** | Returns boolean if VPC is configured on switch or not. | [optional] 
**PeerSwitchDbId** | Pointer to **int64** | Returns peer switch id if VPC configured. | [optional] 
**SwitchDbId** | Pointer to **int64** | Returns the switch id of the switch. | [optional] 

## Methods

### NewNiatelemetryVpcDetails

`func NewNiatelemetryVpcDetails(classId string, objectType string, ) *NiatelemetryVpcDetails`

NewNiatelemetryVpcDetails instantiates a new NiatelemetryVpcDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryVpcDetailsWithDefaults

`func NewNiatelemetryVpcDetailsWithDefaults() *NiatelemetryVpcDetails`

NewNiatelemetryVpcDetailsWithDefaults instantiates a new NiatelemetryVpcDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryVpcDetails) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryVpcDetails) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryVpcDetails) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryVpcDetails) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryVpcDetails) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryVpcDetails) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIsVpcConfigured

`func (o *NiatelemetryVpcDetails) GetIsVpcConfigured() bool`

GetIsVpcConfigured returns the IsVpcConfigured field if non-nil, zero value otherwise.

### GetIsVpcConfiguredOk

`func (o *NiatelemetryVpcDetails) GetIsVpcConfiguredOk() (*bool, bool)`

GetIsVpcConfiguredOk returns a tuple with the IsVpcConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVpcConfigured

`func (o *NiatelemetryVpcDetails) SetIsVpcConfigured(v bool)`

SetIsVpcConfigured sets IsVpcConfigured field to given value.

### HasIsVpcConfigured

`func (o *NiatelemetryVpcDetails) HasIsVpcConfigured() bool`

HasIsVpcConfigured returns a boolean if a field has been set.

### GetPeerSwitchDbId

`func (o *NiatelemetryVpcDetails) GetPeerSwitchDbId() int64`

GetPeerSwitchDbId returns the PeerSwitchDbId field if non-nil, zero value otherwise.

### GetPeerSwitchDbIdOk

`func (o *NiatelemetryVpcDetails) GetPeerSwitchDbIdOk() (*int64, bool)`

GetPeerSwitchDbIdOk returns a tuple with the PeerSwitchDbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerSwitchDbId

`func (o *NiatelemetryVpcDetails) SetPeerSwitchDbId(v int64)`

SetPeerSwitchDbId sets PeerSwitchDbId field to given value.

### HasPeerSwitchDbId

`func (o *NiatelemetryVpcDetails) HasPeerSwitchDbId() bool`

HasPeerSwitchDbId returns a boolean if a field has been set.

### GetSwitchDbId

`func (o *NiatelemetryVpcDetails) GetSwitchDbId() int64`

GetSwitchDbId returns the SwitchDbId field if non-nil, zero value otherwise.

### GetSwitchDbIdOk

`func (o *NiatelemetryVpcDetails) GetSwitchDbIdOk() (*int64, bool)`

GetSwitchDbIdOk returns a tuple with the SwitchDbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchDbId

`func (o *NiatelemetryVpcDetails) SetSwitchDbId(v int64)`

SetSwitchDbId sets SwitchDbId field to given value.

### HasSwitchDbId

`func (o *NiatelemetryVpcDetails) HasSwitchDbId() bool`

HasSwitchDbId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


