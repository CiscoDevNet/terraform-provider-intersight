# NiatelemetryNxosModuleInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.NxosModuleInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.NxosModuleInfo"]
**Hw** | Pointer to **string** | Hardware version reported for the module. | [optional] [readonly] 
**Mod** | Pointer to **string** | Module number or slot identifier in the switch chassis. | [optional] [readonly] 
**Model** | Pointer to **string** | Model name or product ID of the module hardware. | [optional] [readonly] 
**OnlineDiagStatus** | Pointer to **string** | Online diagnostic status of the module. | [optional] [readonly] 
**SerialNumber** | Pointer to **string** | Serial number of the module. | [optional] [readonly] 
**Slot** | Pointer to **string** | Slot type reported for the module. | [optional] [readonly] 
**Status** | Pointer to **string** | Operational status of the module. | [optional] [readonly] 
**Sw** | Pointer to **string** | Software version reported for the module. | [optional] [readonly] 

## Methods

### NewNiatelemetryNxosModuleInfo

`func NewNiatelemetryNxosModuleInfo(classId string, objectType string, ) *NiatelemetryNxosModuleInfo`

NewNiatelemetryNxosModuleInfo instantiates a new NiatelemetryNxosModuleInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryNxosModuleInfoWithDefaults

`func NewNiatelemetryNxosModuleInfoWithDefaults() *NiatelemetryNxosModuleInfo`

NewNiatelemetryNxosModuleInfoWithDefaults instantiates a new NiatelemetryNxosModuleInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryNxosModuleInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryNxosModuleInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryNxosModuleInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryNxosModuleInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryNxosModuleInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryNxosModuleInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetHw

`func (o *NiatelemetryNxosModuleInfo) GetHw() string`

GetHw returns the Hw field if non-nil, zero value otherwise.

### GetHwOk

`func (o *NiatelemetryNxosModuleInfo) GetHwOk() (*string, bool)`

GetHwOk returns a tuple with the Hw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHw

`func (o *NiatelemetryNxosModuleInfo) SetHw(v string)`

SetHw sets Hw field to given value.

### HasHw

`func (o *NiatelemetryNxosModuleInfo) HasHw() bool`

HasHw returns a boolean if a field has been set.

### GetMod

`func (o *NiatelemetryNxosModuleInfo) GetMod() string`

GetMod returns the Mod field if non-nil, zero value otherwise.

### GetModOk

`func (o *NiatelemetryNxosModuleInfo) GetModOk() (*string, bool)`

GetModOk returns a tuple with the Mod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMod

`func (o *NiatelemetryNxosModuleInfo) SetMod(v string)`

SetMod sets Mod field to given value.

### HasMod

`func (o *NiatelemetryNxosModuleInfo) HasMod() bool`

HasMod returns a boolean if a field has been set.

### GetModel

`func (o *NiatelemetryNxosModuleInfo) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *NiatelemetryNxosModuleInfo) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *NiatelemetryNxosModuleInfo) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *NiatelemetryNxosModuleInfo) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetOnlineDiagStatus

`func (o *NiatelemetryNxosModuleInfo) GetOnlineDiagStatus() string`

GetOnlineDiagStatus returns the OnlineDiagStatus field if non-nil, zero value otherwise.

### GetOnlineDiagStatusOk

`func (o *NiatelemetryNxosModuleInfo) GetOnlineDiagStatusOk() (*string, bool)`

GetOnlineDiagStatusOk returns a tuple with the OnlineDiagStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineDiagStatus

`func (o *NiatelemetryNxosModuleInfo) SetOnlineDiagStatus(v string)`

SetOnlineDiagStatus sets OnlineDiagStatus field to given value.

### HasOnlineDiagStatus

`func (o *NiatelemetryNxosModuleInfo) HasOnlineDiagStatus() bool`

HasOnlineDiagStatus returns a boolean if a field has been set.

### GetSerialNumber

`func (o *NiatelemetryNxosModuleInfo) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *NiatelemetryNxosModuleInfo) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *NiatelemetryNxosModuleInfo) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *NiatelemetryNxosModuleInfo) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSlot

`func (o *NiatelemetryNxosModuleInfo) GetSlot() string`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *NiatelemetryNxosModuleInfo) GetSlotOk() (*string, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *NiatelemetryNxosModuleInfo) SetSlot(v string)`

SetSlot sets Slot field to given value.

### HasSlot

`func (o *NiatelemetryNxosModuleInfo) HasSlot() bool`

HasSlot returns a boolean if a field has been set.

### GetStatus

`func (o *NiatelemetryNxosModuleInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NiatelemetryNxosModuleInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NiatelemetryNxosModuleInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NiatelemetryNxosModuleInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSw

`func (o *NiatelemetryNxosModuleInfo) GetSw() string`

GetSw returns the Sw field if non-nil, zero value otherwise.

### GetSwOk

`func (o *NiatelemetryNxosModuleInfo) GetSwOk() (*string, bool)`

GetSwOk returns a tuple with the Sw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSw

`func (o *NiatelemetryNxosModuleInfo) SetSw(v string)`

SetSw sets Sw field to given value.

### HasSw

`func (o *NiatelemetryNxosModuleInfo) HasSw() bool`

HasSw returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


