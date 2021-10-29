# RackUnitPersonalityAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "rack.UnitPersonality"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "rack.UnitPersonality"]
**AdditionalInfo** | Pointer to **string** | Additional info about the added software personality. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the software personality. | [optional] [readonly] 
**PersonalityId** | Pointer to **int64** | Unique identity of added software personality. | [optional] [readonly] 
**ComputeRackUnit** | Pointer to [**ComputeRackUnitRelationship**](ComputeRackUnitRelationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewRackUnitPersonalityAllOf

`func NewRackUnitPersonalityAllOf(classId string, objectType string, ) *RackUnitPersonalityAllOf`

NewRackUnitPersonalityAllOf instantiates a new RackUnitPersonalityAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRackUnitPersonalityAllOfWithDefaults

`func NewRackUnitPersonalityAllOfWithDefaults() *RackUnitPersonalityAllOf`

NewRackUnitPersonalityAllOfWithDefaults instantiates a new RackUnitPersonalityAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *RackUnitPersonalityAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *RackUnitPersonalityAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *RackUnitPersonalityAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *RackUnitPersonalityAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *RackUnitPersonalityAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *RackUnitPersonalityAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdditionalInfo

`func (o *RackUnitPersonalityAllOf) GetAdditionalInfo() string`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *RackUnitPersonalityAllOf) GetAdditionalInfoOk() (*string, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *RackUnitPersonalityAllOf) SetAdditionalInfo(v string)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *RackUnitPersonalityAllOf) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.

### GetName

`func (o *RackUnitPersonalityAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RackUnitPersonalityAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RackUnitPersonalityAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RackUnitPersonalityAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPersonalityId

`func (o *RackUnitPersonalityAllOf) GetPersonalityId() int64`

GetPersonalityId returns the PersonalityId field if non-nil, zero value otherwise.

### GetPersonalityIdOk

`func (o *RackUnitPersonalityAllOf) GetPersonalityIdOk() (*int64, bool)`

GetPersonalityIdOk returns a tuple with the PersonalityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalityId

`func (o *RackUnitPersonalityAllOf) SetPersonalityId(v int64)`

SetPersonalityId sets PersonalityId field to given value.

### HasPersonalityId

`func (o *RackUnitPersonalityAllOf) HasPersonalityId() bool`

HasPersonalityId returns a boolean if a field has been set.

### GetComputeRackUnit

`func (o *RackUnitPersonalityAllOf) GetComputeRackUnit() ComputeRackUnitRelationship`

GetComputeRackUnit returns the ComputeRackUnit field if non-nil, zero value otherwise.

### GetComputeRackUnitOk

`func (o *RackUnitPersonalityAllOf) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool)`

GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRackUnit

`func (o *RackUnitPersonalityAllOf) SetComputeRackUnit(v ComputeRackUnitRelationship)`

SetComputeRackUnit sets ComputeRackUnit field to given value.

### HasComputeRackUnit

`func (o *RackUnitPersonalityAllOf) HasComputeRackUnit() bool`

HasComputeRackUnit returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *RackUnitPersonalityAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *RackUnitPersonalityAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *RackUnitPersonalityAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *RackUnitPersonalityAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *RackUnitPersonalityAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *RackUnitPersonalityAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *RackUnitPersonalityAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *RackUnitPersonalityAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


