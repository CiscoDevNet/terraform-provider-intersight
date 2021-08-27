# InventoryRequestAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "inventory.Request"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "inventory.Request"]
**Mos** | Pointer to [**[]InventoryInventoryMo**](InventoryInventoryMo.md) |  | [optional] 
**Device** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewInventoryRequestAllOf

`func NewInventoryRequestAllOf(classId string, objectType string, ) *InventoryRequestAllOf`

NewInventoryRequestAllOf instantiates a new InventoryRequestAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryRequestAllOfWithDefaults

`func NewInventoryRequestAllOfWithDefaults() *InventoryRequestAllOf`

NewInventoryRequestAllOfWithDefaults instantiates a new InventoryRequestAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *InventoryRequestAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *InventoryRequestAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *InventoryRequestAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *InventoryRequestAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *InventoryRequestAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *InventoryRequestAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMos

`func (o *InventoryRequestAllOf) GetMos() []InventoryInventoryMo`

GetMos returns the Mos field if non-nil, zero value otherwise.

### GetMosOk

`func (o *InventoryRequestAllOf) GetMosOk() (*[]InventoryInventoryMo, bool)`

GetMosOk returns a tuple with the Mos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMos

`func (o *InventoryRequestAllOf) SetMos(v []InventoryInventoryMo)`

SetMos sets Mos field to given value.

### HasMos

`func (o *InventoryRequestAllOf) HasMos() bool`

HasMos returns a boolean if a field has been set.

### SetMosNil

`func (o *InventoryRequestAllOf) SetMosNil(b bool)`

 SetMosNil sets the value for Mos to be an explicit nil

### UnsetMos
`func (o *InventoryRequestAllOf) UnsetMos()`

UnsetMos ensures that no value is present for Mos, not even an explicit nil
### GetDevice

`func (o *InventoryRequestAllOf) GetDevice() AssetDeviceRegistrationRelationship`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *InventoryRequestAllOf) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *InventoryRequestAllOf) SetDevice(v AssetDeviceRegistrationRelationship)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *InventoryRequestAllOf) HasDevice() bool`

HasDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


