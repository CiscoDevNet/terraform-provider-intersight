# CatalystsdwanChassisIdentifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "catalystsdwan.ChassisIdentifier"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "catalystsdwan.ChassisIdentifier"]
**Model** | Pointer to **string** | Hardware model of the Unified Edge chassis enclosure. | [optional] 
**Moid** | Pointer to **string** | Moid of the equipment.Chassis MO that represents the Unified Edge chassis enclosure. | [optional] 
**Serial** | Pointer to **string** | Serial number of the Unified Edge chassis enclosure. | [optional] 

## Methods

### NewCatalystsdwanChassisIdentifier

`func NewCatalystsdwanChassisIdentifier(classId string, objectType string, ) *CatalystsdwanChassisIdentifier`

NewCatalystsdwanChassisIdentifier instantiates a new CatalystsdwanChassisIdentifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalystsdwanChassisIdentifierWithDefaults

`func NewCatalystsdwanChassisIdentifierWithDefaults() *CatalystsdwanChassisIdentifier`

NewCatalystsdwanChassisIdentifierWithDefaults instantiates a new CatalystsdwanChassisIdentifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CatalystsdwanChassisIdentifier) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CatalystsdwanChassisIdentifier) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CatalystsdwanChassisIdentifier) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CatalystsdwanChassisIdentifier) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CatalystsdwanChassisIdentifier) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CatalystsdwanChassisIdentifier) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetModel

`func (o *CatalystsdwanChassisIdentifier) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CatalystsdwanChassisIdentifier) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CatalystsdwanChassisIdentifier) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *CatalystsdwanChassisIdentifier) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetMoid

`func (o *CatalystsdwanChassisIdentifier) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *CatalystsdwanChassisIdentifier) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *CatalystsdwanChassisIdentifier) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *CatalystsdwanChassisIdentifier) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetSerial

`func (o *CatalystsdwanChassisIdentifier) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *CatalystsdwanChassisIdentifier) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *CatalystsdwanChassisIdentifier) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *CatalystsdwanChassisIdentifier) HasSerial() bool`

HasSerial returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


