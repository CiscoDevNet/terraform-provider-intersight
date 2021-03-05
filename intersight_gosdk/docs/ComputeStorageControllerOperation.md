# ComputeStorageControllerOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "compute.StorageControllerOperation"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "compute.StorageControllerOperation"]
**AdminAction** | Pointer to **string** | Administrative actions that can be performed on the Storage Controller. * &#x60;None&#x60; - No action on the selected Storage Controller. * &#x60;Import&#x60; - Import Foreign config action on the selected Storage Controller. * &#x60;Clear&#x60; - Clear Foreign config action on the selected Storage Controller. | [optional] [default to "None"]
**ControllerId** | Pointer to **string** | Storage Controller Id of the server. | [optional] 

## Methods

### NewComputeStorageControllerOperation

`func NewComputeStorageControllerOperation(classId string, objectType string, ) *ComputeStorageControllerOperation`

NewComputeStorageControllerOperation instantiates a new ComputeStorageControllerOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeStorageControllerOperationWithDefaults

`func NewComputeStorageControllerOperationWithDefaults() *ComputeStorageControllerOperation`

NewComputeStorageControllerOperationWithDefaults instantiates a new ComputeStorageControllerOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ComputeStorageControllerOperation) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ComputeStorageControllerOperation) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ComputeStorageControllerOperation) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ComputeStorageControllerOperation) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ComputeStorageControllerOperation) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ComputeStorageControllerOperation) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminAction

`func (o *ComputeStorageControllerOperation) GetAdminAction() string`

GetAdminAction returns the AdminAction field if non-nil, zero value otherwise.

### GetAdminActionOk

`func (o *ComputeStorageControllerOperation) GetAdminActionOk() (*string, bool)`

GetAdminActionOk returns a tuple with the AdminAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminAction

`func (o *ComputeStorageControllerOperation) SetAdminAction(v string)`

SetAdminAction sets AdminAction field to given value.

### HasAdminAction

`func (o *ComputeStorageControllerOperation) HasAdminAction() bool`

HasAdminAction returns a boolean if a field has been set.

### GetControllerId

`func (o *ComputeStorageControllerOperation) GetControllerId() string`

GetControllerId returns the ControllerId field if non-nil, zero value otherwise.

### GetControllerIdOk

`func (o *ComputeStorageControllerOperation) GetControllerIdOk() (*string, bool)`

GetControllerIdOk returns a tuple with the ControllerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerId

`func (o *ComputeStorageControllerOperation) SetControllerId(v string)`

SetControllerId sets ControllerId field to given value.

### HasControllerId

`func (o *ComputeStorageControllerOperation) HasControllerId() bool`

HasControllerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


