# NiatelemetryDiskinfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.Diskinfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.Diskinfo"]
**Free** | Pointer to **int64** | The free disk capacity, currently the type of this field is set to integer. This determines how much memory is free in Bytes. | [optional] 
**Name** | Pointer to **string** | Disk Name used to identified the disk usage record. This determines the name of the disk partition that is inventoried. | [optional] 
**Total** | Pointer to **int64** | The total disk capacity, it should be the sum of free and used, currently the type of this field is set to integer. This determines the total memory for this partition. | [optional] 
**Used** | Pointer to **int64** | The used disk capacity, currently the type of this field is set to integer. This determines how much memory is used in Bytes. | [optional] 

## Methods

### NewNiatelemetryDiskinfoAllOf

`func NewNiatelemetryDiskinfoAllOf(classId string, objectType string, ) *NiatelemetryDiskinfoAllOf`

NewNiatelemetryDiskinfoAllOf instantiates a new NiatelemetryDiskinfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryDiskinfoAllOfWithDefaults

`func NewNiatelemetryDiskinfoAllOfWithDefaults() *NiatelemetryDiskinfoAllOf`

NewNiatelemetryDiskinfoAllOfWithDefaults instantiates a new NiatelemetryDiskinfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryDiskinfoAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryDiskinfoAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryDiskinfoAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryDiskinfoAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryDiskinfoAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryDiskinfoAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFree

`func (o *NiatelemetryDiskinfoAllOf) GetFree() int64`

GetFree returns the Free field if non-nil, zero value otherwise.

### GetFreeOk

`func (o *NiatelemetryDiskinfoAllOf) GetFreeOk() (*int64, bool)`

GetFreeOk returns a tuple with the Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree

`func (o *NiatelemetryDiskinfoAllOf) SetFree(v int64)`

SetFree sets Free field to given value.

### HasFree

`func (o *NiatelemetryDiskinfoAllOf) HasFree() bool`

HasFree returns a boolean if a field has been set.

### GetName

`func (o *NiatelemetryDiskinfoAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NiatelemetryDiskinfoAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NiatelemetryDiskinfoAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NiatelemetryDiskinfoAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTotal

`func (o *NiatelemetryDiskinfoAllOf) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *NiatelemetryDiskinfoAllOf) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *NiatelemetryDiskinfoAllOf) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *NiatelemetryDiskinfoAllOf) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUsed

`func (o *NiatelemetryDiskinfoAllOf) GetUsed() int64`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *NiatelemetryDiskinfoAllOf) GetUsedOk() (*int64, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *NiatelemetryDiskinfoAllOf) SetUsed(v int64)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *NiatelemetryDiskinfoAllOf) HasUsed() bool`

HasUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


