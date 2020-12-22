# BootDeviceBaseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**Enabled** | Pointer to **bool** | Specifies if the boot device is enabled or disabled. | [optional] [default to false]
**Name** | Pointer to **string** | A name that helps identify a boot device. It can be any string that adheres to the following constraints. It should start and end with an alphanumeric character. It can have underscores and hyphens. It cannot be more than 30 characters. | [optional] 

## Methods

### NewBootDeviceBaseAllOf

`func NewBootDeviceBaseAllOf(classId string, objectType string, ) *BootDeviceBaseAllOf`

NewBootDeviceBaseAllOf instantiates a new BootDeviceBaseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBootDeviceBaseAllOfWithDefaults

`func NewBootDeviceBaseAllOfWithDefaults() *BootDeviceBaseAllOf`

NewBootDeviceBaseAllOfWithDefaults instantiates a new BootDeviceBaseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BootDeviceBaseAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BootDeviceBaseAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BootDeviceBaseAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BootDeviceBaseAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BootDeviceBaseAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BootDeviceBaseAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEnabled

`func (o *BootDeviceBaseAllOf) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BootDeviceBaseAllOf) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BootDeviceBaseAllOf) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BootDeviceBaseAllOf) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetName

`func (o *BootDeviceBaseAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BootDeviceBaseAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BootDeviceBaseAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BootDeviceBaseAllOf) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


