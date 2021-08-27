# AdapterUnitExpander

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "adapter.UnitExpander"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "adapter.UnitExpander"]
**PartNumber** | Pointer to **string** | This field identifies the partNumber of the given component. | [optional] [readonly] 
**Vid** | Pointer to **string** | This field identifies the virtual id of the given component. | [optional] [readonly] 
**AdapterUnit** | Pointer to [**AdapterUnitRelationship**](AdapterUnitRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewAdapterUnitExpander

`func NewAdapterUnitExpander(classId string, objectType string, ) *AdapterUnitExpander`

NewAdapterUnitExpander instantiates a new AdapterUnitExpander object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdapterUnitExpanderWithDefaults

`func NewAdapterUnitExpanderWithDefaults() *AdapterUnitExpander`

NewAdapterUnitExpanderWithDefaults instantiates a new AdapterUnitExpander object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AdapterUnitExpander) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AdapterUnitExpander) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AdapterUnitExpander) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AdapterUnitExpander) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AdapterUnitExpander) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AdapterUnitExpander) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPartNumber

`func (o *AdapterUnitExpander) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *AdapterUnitExpander) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *AdapterUnitExpander) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *AdapterUnitExpander) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetVid

`func (o *AdapterUnitExpander) GetVid() string`

GetVid returns the Vid field if non-nil, zero value otherwise.

### GetVidOk

`func (o *AdapterUnitExpander) GetVidOk() (*string, bool)`

GetVidOk returns a tuple with the Vid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVid

`func (o *AdapterUnitExpander) SetVid(v string)`

SetVid sets Vid field to given value.

### HasVid

`func (o *AdapterUnitExpander) HasVid() bool`

HasVid returns a boolean if a field has been set.

### GetAdapterUnit

`func (o *AdapterUnitExpander) GetAdapterUnit() AdapterUnitRelationship`

GetAdapterUnit returns the AdapterUnit field if non-nil, zero value otherwise.

### GetAdapterUnitOk

`func (o *AdapterUnitExpander) GetAdapterUnitOk() (*AdapterUnitRelationship, bool)`

GetAdapterUnitOk returns a tuple with the AdapterUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterUnit

`func (o *AdapterUnitExpander) SetAdapterUnit(v AdapterUnitRelationship)`

SetAdapterUnit sets AdapterUnit field to given value.

### HasAdapterUnit

`func (o *AdapterUnitExpander) HasAdapterUnit() bool`

HasAdapterUnit returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *AdapterUnitExpander) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *AdapterUnitExpander) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *AdapterUnitExpander) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *AdapterUnitExpander) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


