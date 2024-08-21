# OnpremResourceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "onprem.ResourceInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "onprem.ResourceInfo"]
**Avail** | Pointer to **int64** | Available memory or storage in bytes. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the resource. In case of disk, it is the mount name of the disk. | [optional] [readonly] 
**Total** | Pointer to **int64** | Total memory or storage in bytes. | [optional] [readonly] 

## Methods

### NewOnpremResourceInfo

`func NewOnpremResourceInfo(classId string, objectType string, ) *OnpremResourceInfo`

NewOnpremResourceInfo instantiates a new OnpremResourceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnpremResourceInfoWithDefaults

`func NewOnpremResourceInfoWithDefaults() *OnpremResourceInfo`

NewOnpremResourceInfoWithDefaults instantiates a new OnpremResourceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OnpremResourceInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OnpremResourceInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OnpremResourceInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OnpremResourceInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OnpremResourceInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OnpremResourceInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAvail

`func (o *OnpremResourceInfo) GetAvail() int64`

GetAvail returns the Avail field if non-nil, zero value otherwise.

### GetAvailOk

`func (o *OnpremResourceInfo) GetAvailOk() (*int64, bool)`

GetAvailOk returns a tuple with the Avail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvail

`func (o *OnpremResourceInfo) SetAvail(v int64)`

SetAvail sets Avail field to given value.

### HasAvail

`func (o *OnpremResourceInfo) HasAvail() bool`

HasAvail returns a boolean if a field has been set.

### GetName

`func (o *OnpremResourceInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OnpremResourceInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OnpremResourceInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OnpremResourceInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTotal

`func (o *OnpremResourceInfo) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *OnpremResourceInfo) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *OnpremResourceInfo) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *OnpremResourceInfo) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


