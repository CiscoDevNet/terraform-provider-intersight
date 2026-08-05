# MgmtConfigObjectDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "mgmt.ConfigObjectDetails"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "mgmt.ConfigObjectDetails"]
**Description** | Pointer to **string** | Description of object backed up or restored. Used to store organization and resource group decription. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the configured object. | [optional] [readonly] 
**ObjectMeta** | Pointer to **interface{}** | Additional information about the type of configured object. | [optional] [readonly] 
**Type** | Pointer to **string** | ObjectType of the configured object. | [optional] [readonly] 

## Methods

### NewMgmtConfigObjectDetails

`func NewMgmtConfigObjectDetails(classId string, objectType string, ) *MgmtConfigObjectDetails`

NewMgmtConfigObjectDetails instantiates a new MgmtConfigObjectDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMgmtConfigObjectDetailsWithDefaults

`func NewMgmtConfigObjectDetailsWithDefaults() *MgmtConfigObjectDetails`

NewMgmtConfigObjectDetailsWithDefaults instantiates a new MgmtConfigObjectDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MgmtConfigObjectDetails) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MgmtConfigObjectDetails) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MgmtConfigObjectDetails) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MgmtConfigObjectDetails) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MgmtConfigObjectDetails) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MgmtConfigObjectDetails) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *MgmtConfigObjectDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MgmtConfigObjectDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MgmtConfigObjectDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MgmtConfigObjectDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *MgmtConfigObjectDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MgmtConfigObjectDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MgmtConfigObjectDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MgmtConfigObjectDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjectMeta

`func (o *MgmtConfigObjectDetails) GetObjectMeta() interface{}`

GetObjectMeta returns the ObjectMeta field if non-nil, zero value otherwise.

### GetObjectMetaOk

`func (o *MgmtConfigObjectDetails) GetObjectMetaOk() (*interface{}, bool)`

GetObjectMetaOk returns a tuple with the ObjectMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectMeta

`func (o *MgmtConfigObjectDetails) SetObjectMeta(v interface{})`

SetObjectMeta sets ObjectMeta field to given value.

### HasObjectMeta

`func (o *MgmtConfigObjectDetails) HasObjectMeta() bool`

HasObjectMeta returns a boolean if a field has been set.

### SetObjectMetaNil

`func (o *MgmtConfigObjectDetails) SetObjectMetaNil(b bool)`

 SetObjectMetaNil sets the value for ObjectMeta to be an explicit nil

### UnsetObjectMeta
`func (o *MgmtConfigObjectDetails) UnsetObjectMeta()`

UnsetObjectMeta ensures that no value is present for ObjectMeta, not even an explicit nil
### GetType

`func (o *MgmtConfigObjectDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MgmtConfigObjectDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MgmtConfigObjectDetails) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MgmtConfigObjectDetails) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


