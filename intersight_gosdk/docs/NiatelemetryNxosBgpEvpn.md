# NiatelemetryNxosBgpEvpn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.NxosBgpEvpn"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.NxosBgpEvpn"]
**NxosEvpnMacCount** | Pointer to **string** | Returns the EVPN mac count. | [optional] 
**TotalNetworks** | Pointer to **int64** | Returns the BGP EVPN total networks. | [optional] 
**TotalPaths** | Pointer to **int64** | Returns the BGP EVPN total paths. | [optional] 

## Methods

### NewNiatelemetryNxosBgpEvpn

`func NewNiatelemetryNxosBgpEvpn(classId string, objectType string, ) *NiatelemetryNxosBgpEvpn`

NewNiatelemetryNxosBgpEvpn instantiates a new NiatelemetryNxosBgpEvpn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryNxosBgpEvpnWithDefaults

`func NewNiatelemetryNxosBgpEvpnWithDefaults() *NiatelemetryNxosBgpEvpn`

NewNiatelemetryNxosBgpEvpnWithDefaults instantiates a new NiatelemetryNxosBgpEvpn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryNxosBgpEvpn) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryNxosBgpEvpn) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryNxosBgpEvpn) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryNxosBgpEvpn) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryNxosBgpEvpn) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryNxosBgpEvpn) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetNxosEvpnMacCount

`func (o *NiatelemetryNxosBgpEvpn) GetNxosEvpnMacCount() string`

GetNxosEvpnMacCount returns the NxosEvpnMacCount field if non-nil, zero value otherwise.

### GetNxosEvpnMacCountOk

`func (o *NiatelemetryNxosBgpEvpn) GetNxosEvpnMacCountOk() (*string, bool)`

GetNxosEvpnMacCountOk returns a tuple with the NxosEvpnMacCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxosEvpnMacCount

`func (o *NiatelemetryNxosBgpEvpn) SetNxosEvpnMacCount(v string)`

SetNxosEvpnMacCount sets NxosEvpnMacCount field to given value.

### HasNxosEvpnMacCount

`func (o *NiatelemetryNxosBgpEvpn) HasNxosEvpnMacCount() bool`

HasNxosEvpnMacCount returns a boolean if a field has been set.

### GetTotalNetworks

`func (o *NiatelemetryNxosBgpEvpn) GetTotalNetworks() int64`

GetTotalNetworks returns the TotalNetworks field if non-nil, zero value otherwise.

### GetTotalNetworksOk

`func (o *NiatelemetryNxosBgpEvpn) GetTotalNetworksOk() (*int64, bool)`

GetTotalNetworksOk returns a tuple with the TotalNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNetworks

`func (o *NiatelemetryNxosBgpEvpn) SetTotalNetworks(v int64)`

SetTotalNetworks sets TotalNetworks field to given value.

### HasTotalNetworks

`func (o *NiatelemetryNxosBgpEvpn) HasTotalNetworks() bool`

HasTotalNetworks returns a boolean if a field has been set.

### GetTotalPaths

`func (o *NiatelemetryNxosBgpEvpn) GetTotalPaths() int64`

GetTotalPaths returns the TotalPaths field if non-nil, zero value otherwise.

### GetTotalPathsOk

`func (o *NiatelemetryNxosBgpEvpn) GetTotalPathsOk() (*int64, bool)`

GetTotalPathsOk returns a tuple with the TotalPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPaths

`func (o *NiatelemetryNxosBgpEvpn) SetTotalPaths(v int64)`

SetTotalPaths sets TotalPaths field to given value.

### HasTotalPaths

`func (o *NiatelemetryNxosBgpEvpn) HasTotalPaths() bool`

HasTotalPaths returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


