# ResourceProcessorQualifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "resource.ProcessorQualifier"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "resource.ProcessorQualifier"]
**CpuCoresRange** | Pointer to [**ResourceCpuCoreRangeFilter**](ResourceCpuCoreRangeFilter.md) |  | [optional] 
**Pids** | Pointer to **[]string** |  | [optional] 
**SpeedRange** | Pointer to [**ResourceCpuSpeedRangeFilter**](ResourceCpuSpeedRangeFilter.md) |  | [optional] 
**Vendor** | Pointer to **string** | The qualification of resources is based on the vendor of the processor. | [optional] 

## Methods

### NewResourceProcessorQualifier

`func NewResourceProcessorQualifier(classId string, objectType string, ) *ResourceProcessorQualifier`

NewResourceProcessorQualifier instantiates a new ResourceProcessorQualifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceProcessorQualifierWithDefaults

`func NewResourceProcessorQualifierWithDefaults() *ResourceProcessorQualifier`

NewResourceProcessorQualifierWithDefaults instantiates a new ResourceProcessorQualifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ResourceProcessorQualifier) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ResourceProcessorQualifier) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ResourceProcessorQualifier) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ResourceProcessorQualifier) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ResourceProcessorQualifier) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ResourceProcessorQualifier) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCpuCoresRange

`func (o *ResourceProcessorQualifier) GetCpuCoresRange() ResourceCpuCoreRangeFilter`

GetCpuCoresRange returns the CpuCoresRange field if non-nil, zero value otherwise.

### GetCpuCoresRangeOk

`func (o *ResourceProcessorQualifier) GetCpuCoresRangeOk() (*ResourceCpuCoreRangeFilter, bool)`

GetCpuCoresRangeOk returns a tuple with the CpuCoresRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCoresRange

`func (o *ResourceProcessorQualifier) SetCpuCoresRange(v ResourceCpuCoreRangeFilter)`

SetCpuCoresRange sets CpuCoresRange field to given value.

### HasCpuCoresRange

`func (o *ResourceProcessorQualifier) HasCpuCoresRange() bool`

HasCpuCoresRange returns a boolean if a field has been set.

### GetPids

`func (o *ResourceProcessorQualifier) GetPids() []string`

GetPids returns the Pids field if non-nil, zero value otherwise.

### GetPidsOk

`func (o *ResourceProcessorQualifier) GetPidsOk() (*[]string, bool)`

GetPidsOk returns a tuple with the Pids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPids

`func (o *ResourceProcessorQualifier) SetPids(v []string)`

SetPids sets Pids field to given value.

### HasPids

`func (o *ResourceProcessorQualifier) HasPids() bool`

HasPids returns a boolean if a field has been set.

### SetPidsNil

`func (o *ResourceProcessorQualifier) SetPidsNil(b bool)`

 SetPidsNil sets the value for Pids to be an explicit nil

### UnsetPids
`func (o *ResourceProcessorQualifier) UnsetPids()`

UnsetPids ensures that no value is present for Pids, not even an explicit nil
### GetSpeedRange

`func (o *ResourceProcessorQualifier) GetSpeedRange() ResourceCpuSpeedRangeFilter`

GetSpeedRange returns the SpeedRange field if non-nil, zero value otherwise.

### GetSpeedRangeOk

`func (o *ResourceProcessorQualifier) GetSpeedRangeOk() (*ResourceCpuSpeedRangeFilter, bool)`

GetSpeedRangeOk returns a tuple with the SpeedRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedRange

`func (o *ResourceProcessorQualifier) SetSpeedRange(v ResourceCpuSpeedRangeFilter)`

SetSpeedRange sets SpeedRange field to given value.

### HasSpeedRange

`func (o *ResourceProcessorQualifier) HasSpeedRange() bool`

HasSpeedRange returns a boolean if a field has been set.

### GetVendor

`func (o *ResourceProcessorQualifier) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *ResourceProcessorQualifier) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *ResourceProcessorQualifier) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *ResourceProcessorQualifier) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


