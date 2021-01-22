# StorageBaseTenant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "storage.NetAppStorageVm"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "storage.NetAppStorageVm"]
**Name** | Pointer to **string** | Name of the tenant in storage array. | [optional] [readonly] 
**State** | Pointer to **string** | The state of this tenant. * &#x60;Unknown&#x60; - Component state is not available. * &#x60;Starting&#x60; - Component is being started. * &#x60;Running&#x60; - Component is currently running. * &#x60;Stopping&#x60; - Component is being stopped. * &#x60;Stopped&#x60; - Component has been stopped. * &#x60;Deleting&#x60; - Component deletion is in progress. | [optional] [readonly] [default to "Unknown"]
**Uuid** | Pointer to **string** | The uuid of this tenant in storage array. | [optional] [readonly] 

## Methods

### NewStorageBaseTenant

`func NewStorageBaseTenant(classId string, objectType string, ) *StorageBaseTenant`

NewStorageBaseTenant instantiates a new StorageBaseTenant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageBaseTenantWithDefaults

`func NewStorageBaseTenantWithDefaults() *StorageBaseTenant`

NewStorageBaseTenantWithDefaults instantiates a new StorageBaseTenant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageBaseTenant) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageBaseTenant) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageBaseTenant) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageBaseTenant) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageBaseTenant) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageBaseTenant) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *StorageBaseTenant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageBaseTenant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageBaseTenant) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageBaseTenant) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *StorageBaseTenant) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StorageBaseTenant) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StorageBaseTenant) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *StorageBaseTenant) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUuid

`func (o *StorageBaseTenant) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StorageBaseTenant) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StorageBaseTenant) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *StorageBaseTenant) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


