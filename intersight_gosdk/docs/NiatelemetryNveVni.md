# NiatelemetryNveVni

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.NveVni"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.NveVni"]
**CpVniCount** | Pointer to **int64** | Return value of cp vni count. | [optional] 
**CpVniDown** | Pointer to **int64** | Return value of cp vni down count. | [optional] 
**CpVniUp** | Pointer to **int64** | Return value of cp vni up count. | [optional] 
**DpVniCount** | Pointer to **int64** | Return value of dp vni count. | [optional] 
**DpVniDown** | Pointer to **int64** | Return value of cp vni down count. | [optional] 
**DpVniUp** | Pointer to **int64** | Return value of cp vni up count. | [optional] 

## Methods

### NewNiatelemetryNveVni

`func NewNiatelemetryNveVni(classId string, objectType string, ) *NiatelemetryNveVni`

NewNiatelemetryNveVni instantiates a new NiatelemetryNveVni object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryNveVniWithDefaults

`func NewNiatelemetryNveVniWithDefaults() *NiatelemetryNveVni`

NewNiatelemetryNveVniWithDefaults instantiates a new NiatelemetryNveVni object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryNveVni) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryNveVni) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryNveVni) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryNveVni) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryNveVni) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryNveVni) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCpVniCount

`func (o *NiatelemetryNveVni) GetCpVniCount() int64`

GetCpVniCount returns the CpVniCount field if non-nil, zero value otherwise.

### GetCpVniCountOk

`func (o *NiatelemetryNveVni) GetCpVniCountOk() (*int64, bool)`

GetCpVniCountOk returns a tuple with the CpVniCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpVniCount

`func (o *NiatelemetryNveVni) SetCpVniCount(v int64)`

SetCpVniCount sets CpVniCount field to given value.

### HasCpVniCount

`func (o *NiatelemetryNveVni) HasCpVniCount() bool`

HasCpVniCount returns a boolean if a field has been set.

### GetCpVniDown

`func (o *NiatelemetryNveVni) GetCpVniDown() int64`

GetCpVniDown returns the CpVniDown field if non-nil, zero value otherwise.

### GetCpVniDownOk

`func (o *NiatelemetryNveVni) GetCpVniDownOk() (*int64, bool)`

GetCpVniDownOk returns a tuple with the CpVniDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpVniDown

`func (o *NiatelemetryNveVni) SetCpVniDown(v int64)`

SetCpVniDown sets CpVniDown field to given value.

### HasCpVniDown

`func (o *NiatelemetryNveVni) HasCpVniDown() bool`

HasCpVniDown returns a boolean if a field has been set.

### GetCpVniUp

`func (o *NiatelemetryNveVni) GetCpVniUp() int64`

GetCpVniUp returns the CpVniUp field if non-nil, zero value otherwise.

### GetCpVniUpOk

`func (o *NiatelemetryNveVni) GetCpVniUpOk() (*int64, bool)`

GetCpVniUpOk returns a tuple with the CpVniUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpVniUp

`func (o *NiatelemetryNveVni) SetCpVniUp(v int64)`

SetCpVniUp sets CpVniUp field to given value.

### HasCpVniUp

`func (o *NiatelemetryNveVni) HasCpVniUp() bool`

HasCpVniUp returns a boolean if a field has been set.

### GetDpVniCount

`func (o *NiatelemetryNveVni) GetDpVniCount() int64`

GetDpVniCount returns the DpVniCount field if non-nil, zero value otherwise.

### GetDpVniCountOk

`func (o *NiatelemetryNveVni) GetDpVniCountOk() (*int64, bool)`

GetDpVniCountOk returns a tuple with the DpVniCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpVniCount

`func (o *NiatelemetryNveVni) SetDpVniCount(v int64)`

SetDpVniCount sets DpVniCount field to given value.

### HasDpVniCount

`func (o *NiatelemetryNveVni) HasDpVniCount() bool`

HasDpVniCount returns a boolean if a field has been set.

### GetDpVniDown

`func (o *NiatelemetryNveVni) GetDpVniDown() int64`

GetDpVniDown returns the DpVniDown field if non-nil, zero value otherwise.

### GetDpVniDownOk

`func (o *NiatelemetryNveVni) GetDpVniDownOk() (*int64, bool)`

GetDpVniDownOk returns a tuple with the DpVniDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpVniDown

`func (o *NiatelemetryNveVni) SetDpVniDown(v int64)`

SetDpVniDown sets DpVniDown field to given value.

### HasDpVniDown

`func (o *NiatelemetryNveVni) HasDpVniDown() bool`

HasDpVniDown returns a boolean if a field has been set.

### GetDpVniUp

`func (o *NiatelemetryNveVni) GetDpVniUp() int64`

GetDpVniUp returns the DpVniUp field if non-nil, zero value otherwise.

### GetDpVniUpOk

`func (o *NiatelemetryNveVni) GetDpVniUpOk() (*int64, bool)`

GetDpVniUpOk returns a tuple with the DpVniUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpVniUp

`func (o *NiatelemetryNveVni) SetDpVniUp(v int64)`

SetDpVniUp sets DpVniUp field to given value.

### HasDpVniUp

`func (o *NiatelemetryNveVni) HasDpVniUp() bool`

HasDpVniUp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


