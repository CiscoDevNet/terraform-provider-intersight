# EquipmentDeviceLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "equipment.DeviceLocation"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "equipment.DeviceLocation"]
**Address** | Pointer to **string** | The information about the address. | [optional] [readonly] 
**Latitude** | Pointer to **float32** | Location latitude in decimal degrees format. | [optional] [readonly] 
**Longitude** | Pointer to **float32** | Location longitude in decimal degrees format. | [optional] [readonly] 

## Methods

### NewEquipmentDeviceLocation

`func NewEquipmentDeviceLocation(classId string, objectType string, ) *EquipmentDeviceLocation`

NewEquipmentDeviceLocation instantiates a new EquipmentDeviceLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentDeviceLocationWithDefaults

`func NewEquipmentDeviceLocationWithDefaults() *EquipmentDeviceLocation`

NewEquipmentDeviceLocationWithDefaults instantiates a new EquipmentDeviceLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EquipmentDeviceLocation) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EquipmentDeviceLocation) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EquipmentDeviceLocation) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EquipmentDeviceLocation) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EquipmentDeviceLocation) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EquipmentDeviceLocation) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAddress

`func (o *EquipmentDeviceLocation) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *EquipmentDeviceLocation) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *EquipmentDeviceLocation) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *EquipmentDeviceLocation) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetLatitude

`func (o *EquipmentDeviceLocation) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *EquipmentDeviceLocation) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *EquipmentDeviceLocation) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *EquipmentDeviceLocation) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *EquipmentDeviceLocation) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *EquipmentDeviceLocation) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *EquipmentDeviceLocation) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *EquipmentDeviceLocation) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


