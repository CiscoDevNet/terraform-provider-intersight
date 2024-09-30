# ComputeStorageUtilityImageOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "compute.StorageUtilityImageOperation"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "compute.StorageUtilityImageOperation"]
**Action** | Pointer to **string** | Actions that can be performed by the storage utility. * &#x60;None&#x60; - No action by storage utility. * &#x60;Upload&#x60; - Upload action by storage utility. * &#x60;TurnOnImageVisibility&#x60; - Turn on image&#39;s visibility. * &#x60;TurnOffImageVisibility&#x60; - Turn off image&#39;s visibility. * &#x60;Remove&#x60; - Removes images from a server&#39;s utility storage. | [optional] [default to "None"]
**ImageName** | Pointer to **string** | The image name this action operates on. | [optional] 
**ImageType** | Pointer to **string** | The image type this action operates on. | [optional] 

## Methods

### NewComputeStorageUtilityImageOperation

`func NewComputeStorageUtilityImageOperation(classId string, objectType string, ) *ComputeStorageUtilityImageOperation`

NewComputeStorageUtilityImageOperation instantiates a new ComputeStorageUtilityImageOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeStorageUtilityImageOperationWithDefaults

`func NewComputeStorageUtilityImageOperationWithDefaults() *ComputeStorageUtilityImageOperation`

NewComputeStorageUtilityImageOperationWithDefaults instantiates a new ComputeStorageUtilityImageOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ComputeStorageUtilityImageOperation) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ComputeStorageUtilityImageOperation) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ComputeStorageUtilityImageOperation) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ComputeStorageUtilityImageOperation) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ComputeStorageUtilityImageOperation) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ComputeStorageUtilityImageOperation) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAction

`func (o *ComputeStorageUtilityImageOperation) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ComputeStorageUtilityImageOperation) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ComputeStorageUtilityImageOperation) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *ComputeStorageUtilityImageOperation) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetImageName

`func (o *ComputeStorageUtilityImageOperation) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *ComputeStorageUtilityImageOperation) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *ComputeStorageUtilityImageOperation) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *ComputeStorageUtilityImageOperation) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetImageType

`func (o *ComputeStorageUtilityImageOperation) GetImageType() string`

GetImageType returns the ImageType field if non-nil, zero value otherwise.

### GetImageTypeOk

`func (o *ComputeStorageUtilityImageOperation) GetImageTypeOk() (*string, bool)`

GetImageTypeOk returns a tuple with the ImageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageType

`func (o *ComputeStorageUtilityImageOperation) SetImageType(v string)`

SetImageType sets ImageType field to given value.

### HasImageType

`func (o *ComputeStorageUtilityImageOperation) HasImageType() bool`

HasImageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


