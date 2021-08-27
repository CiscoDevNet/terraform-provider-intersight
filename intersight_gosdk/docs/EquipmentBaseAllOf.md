# EquipmentBaseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**Model** | Pointer to **string** | This field identifies the model of the given component. | [optional] [readonly] 
**Presence** | Pointer to **string** | This field identifies the presence (equipped) or absence of the given component. | [optional] [readonly] 
**Revision** | Pointer to **string** | This field identifies the revision of the given component. | [optional] [readonly] 
**Serial** | Pointer to **string** | This field identifies the serial of the given component. | [optional] [readonly] 
**Vendor** | Pointer to **string** | This field identifies the vendor of the given component. | [optional] [readonly] 
**PreviousFru** | Pointer to [**EquipmentFruRelationship**](EquipmentFruRelationship.md) |  | [optional] 

## Methods

### NewEquipmentBaseAllOf

`func NewEquipmentBaseAllOf(classId string, objectType string, ) *EquipmentBaseAllOf`

NewEquipmentBaseAllOf instantiates a new EquipmentBaseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentBaseAllOfWithDefaults

`func NewEquipmentBaseAllOfWithDefaults() *EquipmentBaseAllOf`

NewEquipmentBaseAllOfWithDefaults instantiates a new EquipmentBaseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EquipmentBaseAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EquipmentBaseAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EquipmentBaseAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EquipmentBaseAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EquipmentBaseAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EquipmentBaseAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetModel

`func (o *EquipmentBaseAllOf) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *EquipmentBaseAllOf) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *EquipmentBaseAllOf) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *EquipmentBaseAllOf) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetPresence

`func (o *EquipmentBaseAllOf) GetPresence() string`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *EquipmentBaseAllOf) GetPresenceOk() (*string, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *EquipmentBaseAllOf) SetPresence(v string)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *EquipmentBaseAllOf) HasPresence() bool`

HasPresence returns a boolean if a field has been set.

### GetRevision

`func (o *EquipmentBaseAllOf) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *EquipmentBaseAllOf) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *EquipmentBaseAllOf) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *EquipmentBaseAllOf) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSerial

`func (o *EquipmentBaseAllOf) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *EquipmentBaseAllOf) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *EquipmentBaseAllOf) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *EquipmentBaseAllOf) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetVendor

`func (o *EquipmentBaseAllOf) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *EquipmentBaseAllOf) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *EquipmentBaseAllOf) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *EquipmentBaseAllOf) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetPreviousFru

`func (o *EquipmentBaseAllOf) GetPreviousFru() EquipmentFruRelationship`

GetPreviousFru returns the PreviousFru field if non-nil, zero value otherwise.

### GetPreviousFruOk

`func (o *EquipmentBaseAllOf) GetPreviousFruOk() (*EquipmentFruRelationship, bool)`

GetPreviousFruOk returns a tuple with the PreviousFru field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousFru

`func (o *EquipmentBaseAllOf) SetPreviousFru(v EquipmentFruRelationship)`

SetPreviousFru sets PreviousFru field to given value.

### HasPreviousFru

`func (o *EquipmentBaseAllOf) HasPreviousFru() bool`

HasPreviousFru returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


