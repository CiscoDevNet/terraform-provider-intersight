# NiatelemetryEqptStorageFirmwareAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.EqptStorageFirmware"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.EqptStorageFirmware"]
**Available** | Pointer to **string** | Availabe disk space in aci nodes. | [optional] 
**Used** | Pointer to **string** | Used disk space in aci nodes. | [optional] 

## Methods

### NewNiatelemetryEqptStorageFirmwareAllOf

`func NewNiatelemetryEqptStorageFirmwareAllOf(classId string, objectType string, ) *NiatelemetryEqptStorageFirmwareAllOf`

NewNiatelemetryEqptStorageFirmwareAllOf instantiates a new NiatelemetryEqptStorageFirmwareAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryEqptStorageFirmwareAllOfWithDefaults

`func NewNiatelemetryEqptStorageFirmwareAllOfWithDefaults() *NiatelemetryEqptStorageFirmwareAllOf`

NewNiatelemetryEqptStorageFirmwareAllOfWithDefaults instantiates a new NiatelemetryEqptStorageFirmwareAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryEqptStorageFirmwareAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryEqptStorageFirmwareAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryEqptStorageFirmwareAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryEqptStorageFirmwareAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryEqptStorageFirmwareAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryEqptStorageFirmwareAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAvailable

`func (o *NiatelemetryEqptStorageFirmwareAllOf) GetAvailable() string`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *NiatelemetryEqptStorageFirmwareAllOf) GetAvailableOk() (*string, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *NiatelemetryEqptStorageFirmwareAllOf) SetAvailable(v string)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *NiatelemetryEqptStorageFirmwareAllOf) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetUsed

`func (o *NiatelemetryEqptStorageFirmwareAllOf) GetUsed() string`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *NiatelemetryEqptStorageFirmwareAllOf) GetUsedOk() (*string, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *NiatelemetryEqptStorageFirmwareAllOf) SetUsed(v string)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *NiatelemetryEqptStorageFirmwareAllOf) HasUsed() bool`

HasUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


