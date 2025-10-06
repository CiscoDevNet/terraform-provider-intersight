# CommGeoLocationDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "comm.GeoLocationDetails"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "comm.GeoLocationDetails"]
**Address** | Pointer to [**NullableCommPhysicalAddress**](CommPhysicalAddress.md) |  | [optional] 
**Coordinates** | Pointer to [**NullableCommGeoPoint**](CommGeoPoint.md) |  | [optional] 
**Name** | Pointer to **string** | A user provided name for the location. | [optional] [readonly] 

## Methods

### NewCommGeoLocationDetails

`func NewCommGeoLocationDetails(classId string, objectType string, ) *CommGeoLocationDetails`

NewCommGeoLocationDetails instantiates a new CommGeoLocationDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommGeoLocationDetailsWithDefaults

`func NewCommGeoLocationDetailsWithDefaults() *CommGeoLocationDetails`

NewCommGeoLocationDetailsWithDefaults instantiates a new CommGeoLocationDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CommGeoLocationDetails) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CommGeoLocationDetails) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CommGeoLocationDetails) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CommGeoLocationDetails) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CommGeoLocationDetails) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CommGeoLocationDetails) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAddress

`func (o *CommGeoLocationDetails) GetAddress() CommPhysicalAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CommGeoLocationDetails) GetAddressOk() (*CommPhysicalAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CommGeoLocationDetails) SetAddress(v CommPhysicalAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *CommGeoLocationDetails) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *CommGeoLocationDetails) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *CommGeoLocationDetails) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetCoordinates

`func (o *CommGeoLocationDetails) GetCoordinates() CommGeoPoint`

GetCoordinates returns the Coordinates field if non-nil, zero value otherwise.

### GetCoordinatesOk

`func (o *CommGeoLocationDetails) GetCoordinatesOk() (*CommGeoPoint, bool)`

GetCoordinatesOk returns a tuple with the Coordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinates

`func (o *CommGeoLocationDetails) SetCoordinates(v CommGeoPoint)`

SetCoordinates sets Coordinates field to given value.

### HasCoordinates

`func (o *CommGeoLocationDetails) HasCoordinates() bool`

HasCoordinates returns a boolean if a field has been set.

### SetCoordinatesNil

`func (o *CommGeoLocationDetails) SetCoordinatesNil(b bool)`

 SetCoordinatesNil sets the value for Coordinates to be an explicit nil

### UnsetCoordinates
`func (o *CommGeoLocationDetails) UnsetCoordinates()`

UnsetCoordinates ensures that no value is present for Coordinates, not even an explicit nil
### GetName

`func (o *CommGeoLocationDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CommGeoLocationDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CommGeoLocationDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CommGeoLocationDetails) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


