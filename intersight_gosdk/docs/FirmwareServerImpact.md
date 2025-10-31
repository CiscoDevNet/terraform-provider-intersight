# FirmwareServerImpact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "firmware.ServerImpact"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "firmware.ServerImpact"]
**DisplayName** | Pointer to **string** | Display name of the server impacted by the upgrade. | [optional] 
**Name** | Pointer to **string** | Name of the server impacted by the upgrade. | [optional] 
**RebootAt** | Pointer to **string** | The reboot impact for the server endpoint during the upgrade operation. * &#x60;Phase1&#x60; - Server will be rebooted in the first phase of upgrade. * &#x60;Phase2&#x60; - Server will be rebooted in the second phase of upgrade. | [optional] [default to "Phase1"]
**Server** | Pointer to [**MoMoRef**](MoMoRef.md) |  | [optional] 

## Methods

### NewFirmwareServerImpact

`func NewFirmwareServerImpact(classId string, objectType string, ) *FirmwareServerImpact`

NewFirmwareServerImpact instantiates a new FirmwareServerImpact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareServerImpactWithDefaults

`func NewFirmwareServerImpactWithDefaults() *FirmwareServerImpact`

NewFirmwareServerImpactWithDefaults instantiates a new FirmwareServerImpact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FirmwareServerImpact) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FirmwareServerImpact) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FirmwareServerImpact) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FirmwareServerImpact) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FirmwareServerImpact) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FirmwareServerImpact) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplayName

`func (o *FirmwareServerImpact) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *FirmwareServerImpact) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *FirmwareServerImpact) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *FirmwareServerImpact) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetName

`func (o *FirmwareServerImpact) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FirmwareServerImpact) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FirmwareServerImpact) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FirmwareServerImpact) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRebootAt

`func (o *FirmwareServerImpact) GetRebootAt() string`

GetRebootAt returns the RebootAt field if non-nil, zero value otherwise.

### GetRebootAtOk

`func (o *FirmwareServerImpact) GetRebootAtOk() (*string, bool)`

GetRebootAtOk returns a tuple with the RebootAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootAt

`func (o *FirmwareServerImpact) SetRebootAt(v string)`

SetRebootAt sets RebootAt field to given value.

### HasRebootAt

`func (o *FirmwareServerImpact) HasRebootAt() bool`

HasRebootAt returns a boolean if a field has been set.

### GetServer

`func (o *FirmwareServerImpact) GetServer() MoMoRef`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *FirmwareServerImpact) GetServerOk() (*MoMoRef, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *FirmwareServerImpact) SetServer(v MoMoRef)`

SetServer sets Server field to given value.

### HasServer

`func (o *FirmwareServerImpact) HasServer() bool`

HasServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


