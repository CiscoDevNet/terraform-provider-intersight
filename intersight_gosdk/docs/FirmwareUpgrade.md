# FirmwareUpgrade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "firmware.Upgrade"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "firmware.Upgrade"]
**ExcludeComponentList** | Pointer to **[]string** |  | [optional] 
**Device** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**Server** | Pointer to [**ComputePhysicalRelationship**](ComputePhysicalRelationship.md) |  | [optional] 

## Methods

### NewFirmwareUpgrade

`func NewFirmwareUpgrade(classId string, objectType string, ) *FirmwareUpgrade`

NewFirmwareUpgrade instantiates a new FirmwareUpgrade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareUpgradeWithDefaults

`func NewFirmwareUpgradeWithDefaults() *FirmwareUpgrade`

NewFirmwareUpgradeWithDefaults instantiates a new FirmwareUpgrade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FirmwareUpgrade) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FirmwareUpgrade) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FirmwareUpgrade) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FirmwareUpgrade) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FirmwareUpgrade) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FirmwareUpgrade) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetExcludeComponentList

`func (o *FirmwareUpgrade) GetExcludeComponentList() []string`

GetExcludeComponentList returns the ExcludeComponentList field if non-nil, zero value otherwise.

### GetExcludeComponentListOk

`func (o *FirmwareUpgrade) GetExcludeComponentListOk() (*[]string, bool)`

GetExcludeComponentListOk returns a tuple with the ExcludeComponentList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeComponentList

`func (o *FirmwareUpgrade) SetExcludeComponentList(v []string)`

SetExcludeComponentList sets ExcludeComponentList field to given value.

### HasExcludeComponentList

`func (o *FirmwareUpgrade) HasExcludeComponentList() bool`

HasExcludeComponentList returns a boolean if a field has been set.

### SetExcludeComponentListNil

`func (o *FirmwareUpgrade) SetExcludeComponentListNil(b bool)`

 SetExcludeComponentListNil sets the value for ExcludeComponentList to be an explicit nil

### UnsetExcludeComponentList
`func (o *FirmwareUpgrade) UnsetExcludeComponentList()`

UnsetExcludeComponentList ensures that no value is present for ExcludeComponentList, not even an explicit nil
### GetDevice

`func (o *FirmwareUpgrade) GetDevice() AssetDeviceRegistrationRelationship`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *FirmwareUpgrade) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *FirmwareUpgrade) SetDevice(v AssetDeviceRegistrationRelationship)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *FirmwareUpgrade) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetServer

`func (o *FirmwareUpgrade) GetServer() ComputePhysicalRelationship`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *FirmwareUpgrade) GetServerOk() (*ComputePhysicalRelationship, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *FirmwareUpgrade) SetServer(v ComputePhysicalRelationship)`

SetServer sets Server field to given value.

### HasServer

`func (o *FirmwareUpgrade) HasServer() bool`

HasServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


