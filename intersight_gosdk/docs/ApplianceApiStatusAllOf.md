# ApplianceApiStatusAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.ApiStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.ApiStatus"]
**ElapsedTime** | Pointer to **float32** | The elapsed time for query in seconds. | [optional] [readonly] 
**ObjectName** | Pointer to **string** | Name to identify the API. | [optional] [readonly] 
**Reason** | Pointer to **string** | Reason to address why the API call failed, if API call was successed, reason would be null. | [optional] [readonly] 
**Response** | Pointer to **int64** | Response code of the API call. | [optional] [readonly] 

## Methods

### NewApplianceApiStatusAllOf

`func NewApplianceApiStatusAllOf(classId string, objectType string, ) *ApplianceApiStatusAllOf`

NewApplianceApiStatusAllOf instantiates a new ApplianceApiStatusAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceApiStatusAllOfWithDefaults

`func NewApplianceApiStatusAllOfWithDefaults() *ApplianceApiStatusAllOf`

NewApplianceApiStatusAllOfWithDefaults instantiates a new ApplianceApiStatusAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceApiStatusAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceApiStatusAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceApiStatusAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceApiStatusAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceApiStatusAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceApiStatusAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetElapsedTime

`func (o *ApplianceApiStatusAllOf) GetElapsedTime() float32`

GetElapsedTime returns the ElapsedTime field if non-nil, zero value otherwise.

### GetElapsedTimeOk

`func (o *ApplianceApiStatusAllOf) GetElapsedTimeOk() (*float32, bool)`

GetElapsedTimeOk returns a tuple with the ElapsedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsedTime

`func (o *ApplianceApiStatusAllOf) SetElapsedTime(v float32)`

SetElapsedTime sets ElapsedTime field to given value.

### HasElapsedTime

`func (o *ApplianceApiStatusAllOf) HasElapsedTime() bool`

HasElapsedTime returns a boolean if a field has been set.

### GetObjectName

`func (o *ApplianceApiStatusAllOf) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *ApplianceApiStatusAllOf) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *ApplianceApiStatusAllOf) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *ApplianceApiStatusAllOf) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetReason

`func (o *ApplianceApiStatusAllOf) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ApplianceApiStatusAllOf) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ApplianceApiStatusAllOf) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ApplianceApiStatusAllOf) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetResponse

`func (o *ApplianceApiStatusAllOf) GetResponse() int64`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *ApplianceApiStatusAllOf) GetResponseOk() (*int64, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *ApplianceApiStatusAllOf) SetResponse(v int64)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *ApplianceApiStatusAllOf) HasResponse() bool`

HasResponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


