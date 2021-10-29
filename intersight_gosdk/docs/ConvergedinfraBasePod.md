# ConvergedinfraBasePod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "convergedinfra.Pod"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "convergedinfra.Pod"]
**Description** | Pointer to **string** | Description of the pod. A short note about the nature or purpose of the pod. | [optional] 
**Name** | Pointer to **string** | Name of the pod. Concrete pod will be created with this name. | [optional] 
**Type** | Pointer to **string** | Defines the type of the pod. * &#x60;FlexPod&#x60; - Pod type is FlexPod, an integrated infrastructure solution developed by Cisco and NetApp. * &#x60;FlashStack&#x60; - Pod type is FlashStack, an integrated infrastructure solution developed by Cisco and Pure Storage. | [optional] [default to "FlexPod"]

## Methods

### NewConvergedinfraBasePod

`func NewConvergedinfraBasePod(classId string, objectType string, ) *ConvergedinfraBasePod`

NewConvergedinfraBasePod instantiates a new ConvergedinfraBasePod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConvergedinfraBasePodWithDefaults

`func NewConvergedinfraBasePodWithDefaults() *ConvergedinfraBasePod`

NewConvergedinfraBasePodWithDefaults instantiates a new ConvergedinfraBasePod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ConvergedinfraBasePod) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ConvergedinfraBasePod) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ConvergedinfraBasePod) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ConvergedinfraBasePod) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConvergedinfraBasePod) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConvergedinfraBasePod) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *ConvergedinfraBasePod) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConvergedinfraBasePod) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConvergedinfraBasePod) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConvergedinfraBasePod) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *ConvergedinfraBasePod) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConvergedinfraBasePod) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConvergedinfraBasePod) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConvergedinfraBasePod) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ConvergedinfraBasePod) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConvergedinfraBasePod) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConvergedinfraBasePod) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ConvergedinfraBasePod) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


