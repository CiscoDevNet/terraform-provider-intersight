# NetworkHyperFlexInternetProtocolAddressRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "network.HyperFlexInternetProtocolAddressRange"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "network.HyperFlexInternetProtocolAddressRange"]
**BeginAddress** | Pointer to **string** | Begining IP address. IPv4 only. | [optional] [readonly] 
**EndAddress** | Pointer to **string** | Ending IP address. IPv4 only. | [optional] [readonly] 

## Methods

### NewNetworkHyperFlexInternetProtocolAddressRange

`func NewNetworkHyperFlexInternetProtocolAddressRange(classId string, objectType string, ) *NetworkHyperFlexInternetProtocolAddressRange`

NewNetworkHyperFlexInternetProtocolAddressRange instantiates a new NetworkHyperFlexInternetProtocolAddressRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkHyperFlexInternetProtocolAddressRangeWithDefaults

`func NewNetworkHyperFlexInternetProtocolAddressRangeWithDefaults() *NetworkHyperFlexInternetProtocolAddressRange`

NewNetworkHyperFlexInternetProtocolAddressRangeWithDefaults instantiates a new NetworkHyperFlexInternetProtocolAddressRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NetworkHyperFlexInternetProtocolAddressRange) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NetworkHyperFlexInternetProtocolAddressRange) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NetworkHyperFlexInternetProtocolAddressRange) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NetworkHyperFlexInternetProtocolAddressRange) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NetworkHyperFlexInternetProtocolAddressRange) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NetworkHyperFlexInternetProtocolAddressRange) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBeginAddress

`func (o *NetworkHyperFlexInternetProtocolAddressRange) GetBeginAddress() string`

GetBeginAddress returns the BeginAddress field if non-nil, zero value otherwise.

### GetBeginAddressOk

`func (o *NetworkHyperFlexInternetProtocolAddressRange) GetBeginAddressOk() (*string, bool)`

GetBeginAddressOk returns a tuple with the BeginAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeginAddress

`func (o *NetworkHyperFlexInternetProtocolAddressRange) SetBeginAddress(v string)`

SetBeginAddress sets BeginAddress field to given value.

### HasBeginAddress

`func (o *NetworkHyperFlexInternetProtocolAddressRange) HasBeginAddress() bool`

HasBeginAddress returns a boolean if a field has been set.

### GetEndAddress

`func (o *NetworkHyperFlexInternetProtocolAddressRange) GetEndAddress() string`

GetEndAddress returns the EndAddress field if non-nil, zero value otherwise.

### GetEndAddressOk

`func (o *NetworkHyperFlexInternetProtocolAddressRange) GetEndAddressOk() (*string, bool)`

GetEndAddressOk returns a tuple with the EndAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAddress

`func (o *NetworkHyperFlexInternetProtocolAddressRange) SetEndAddress(v string)`

SetEndAddress sets EndAddress field to given value.

### HasEndAddress

`func (o *NetworkHyperFlexInternetProtocolAddressRange) HasEndAddress() bool`

HasEndAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


