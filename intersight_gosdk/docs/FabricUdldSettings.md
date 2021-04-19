# FabricUdldSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.UdldSettings"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.UdldSettings"]
**AdminState** | Pointer to **string** | Admin configured UDLD State for this port. * &#x60;Disabled&#x60; - Admin configured Disabled State. * &#x60;Enabled&#x60; - Admin configured Enabled State. | [optional] [default to "Disabled"]
**Mode** | Pointer to **string** | Admin configured UDLD Mode for this port. * &#x60;normal&#x60; - Admin configured &#39;normal&#39; UDLD mode. * &#x60;aggressive&#x60; - Admin configured &#39;aggressive&#39; UDLD mode. | [optional] [default to "normal"]

## Methods

### NewFabricUdldSettings

`func NewFabricUdldSettings(classId string, objectType string, ) *FabricUdldSettings`

NewFabricUdldSettings instantiates a new FabricUdldSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricUdldSettingsWithDefaults

`func NewFabricUdldSettingsWithDefaults() *FabricUdldSettings`

NewFabricUdldSettingsWithDefaults instantiates a new FabricUdldSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricUdldSettings) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricUdldSettings) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricUdldSettings) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricUdldSettings) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricUdldSettings) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricUdldSettings) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminState

`func (o *FabricUdldSettings) GetAdminState() string`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *FabricUdldSettings) GetAdminStateOk() (*string, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *FabricUdldSettings) SetAdminState(v string)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *FabricUdldSettings) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetMode

`func (o *FabricUdldSettings) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *FabricUdldSettings) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *FabricUdldSettings) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *FabricUdldSettings) HasMode() bool`

HasMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


