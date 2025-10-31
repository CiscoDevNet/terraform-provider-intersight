# CapabilityDeviceInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.DeviceInventory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.DeviceInventory"]
**DeviceObjectType** | Pointer to **string** | The type classification of the endpoint device. | [optional] 
**MaxVersion** | Pointer to **string** | The maximum firmware or software version supported by this endpoint. | [optional] 
**MinVersion** | Pointer to **string** | The minimum firmware or software version supported by this endpoint. | [optional] 
**Model** | Pointer to **string** | The specific model identifier of the endpoint device. | [optional] 

## Methods

### NewCapabilityDeviceInventory

`func NewCapabilityDeviceInventory(classId string, objectType string, ) *CapabilityDeviceInventory`

NewCapabilityDeviceInventory instantiates a new CapabilityDeviceInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityDeviceInventoryWithDefaults

`func NewCapabilityDeviceInventoryWithDefaults() *CapabilityDeviceInventory`

NewCapabilityDeviceInventoryWithDefaults instantiates a new CapabilityDeviceInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityDeviceInventory) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityDeviceInventory) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityDeviceInventory) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityDeviceInventory) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityDeviceInventory) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityDeviceInventory) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDeviceObjectType

`func (o *CapabilityDeviceInventory) GetDeviceObjectType() string`

GetDeviceObjectType returns the DeviceObjectType field if non-nil, zero value otherwise.

### GetDeviceObjectTypeOk

`func (o *CapabilityDeviceInventory) GetDeviceObjectTypeOk() (*string, bool)`

GetDeviceObjectTypeOk returns a tuple with the DeviceObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceObjectType

`func (o *CapabilityDeviceInventory) SetDeviceObjectType(v string)`

SetDeviceObjectType sets DeviceObjectType field to given value.

### HasDeviceObjectType

`func (o *CapabilityDeviceInventory) HasDeviceObjectType() bool`

HasDeviceObjectType returns a boolean if a field has been set.

### GetMaxVersion

`func (o *CapabilityDeviceInventory) GetMaxVersion() string`

GetMaxVersion returns the MaxVersion field if non-nil, zero value otherwise.

### GetMaxVersionOk

`func (o *CapabilityDeviceInventory) GetMaxVersionOk() (*string, bool)`

GetMaxVersionOk returns a tuple with the MaxVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersion

`func (o *CapabilityDeviceInventory) SetMaxVersion(v string)`

SetMaxVersion sets MaxVersion field to given value.

### HasMaxVersion

`func (o *CapabilityDeviceInventory) HasMaxVersion() bool`

HasMaxVersion returns a boolean if a field has been set.

### GetMinVersion

`func (o *CapabilityDeviceInventory) GetMinVersion() string`

GetMinVersion returns the MinVersion field if non-nil, zero value otherwise.

### GetMinVersionOk

`func (o *CapabilityDeviceInventory) GetMinVersionOk() (*string, bool)`

GetMinVersionOk returns a tuple with the MinVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinVersion

`func (o *CapabilityDeviceInventory) SetMinVersion(v string)`

SetMinVersion sets MinVersion field to given value.

### HasMinVersion

`func (o *CapabilityDeviceInventory) HasMinVersion() bool`

HasMinVersion returns a boolean if a field has been set.

### GetModel

`func (o *CapabilityDeviceInventory) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CapabilityDeviceInventory) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CapabilityDeviceInventory) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *CapabilityDeviceInventory) HasModel() bool`

HasModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


