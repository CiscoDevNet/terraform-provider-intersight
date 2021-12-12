# NiatelemetryJobDetailAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.JobDetail"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.JobDetail"]
**JobId** | Pointer to **int64** | Returns the id of the job. | [optional] 
**UpgStatus** | Pointer to **string** | Returns the status of the job. | [optional] 

## Methods

### NewNiatelemetryJobDetailAllOf

`func NewNiatelemetryJobDetailAllOf(classId string, objectType string, ) *NiatelemetryJobDetailAllOf`

NewNiatelemetryJobDetailAllOf instantiates a new NiatelemetryJobDetailAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryJobDetailAllOfWithDefaults

`func NewNiatelemetryJobDetailAllOfWithDefaults() *NiatelemetryJobDetailAllOf`

NewNiatelemetryJobDetailAllOfWithDefaults instantiates a new NiatelemetryJobDetailAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryJobDetailAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryJobDetailAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryJobDetailAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryJobDetailAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryJobDetailAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryJobDetailAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetJobId

`func (o *NiatelemetryJobDetailAllOf) GetJobId() int64`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *NiatelemetryJobDetailAllOf) GetJobIdOk() (*int64, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *NiatelemetryJobDetailAllOf) SetJobId(v int64)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *NiatelemetryJobDetailAllOf) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetUpgStatus

`func (o *NiatelemetryJobDetailAllOf) GetUpgStatus() string`

GetUpgStatus returns the UpgStatus field if non-nil, zero value otherwise.

### GetUpgStatusOk

`func (o *NiatelemetryJobDetailAllOf) GetUpgStatusOk() (*string, bool)`

GetUpgStatusOk returns a tuple with the UpgStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgStatus

`func (o *NiatelemetryJobDetailAllOf) SetUpgStatus(v string)`

SetUpgStatus sets UpgStatus field to given value.

### HasUpgStatus

`func (o *NiatelemetryJobDetailAllOf) HasUpgStatus() bool`

HasUpgStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


