# CondAdapterInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cond.AdapterInstance"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cond.AdapterInstance"]
**Adapter** | Pointer to [**MoMoRef**](MoMoRef.md) |  | [optional] 

## Methods

### NewCondAdapterInstance

`func NewCondAdapterInstance(classId string, objectType string, ) *CondAdapterInstance`

NewCondAdapterInstance instantiates a new CondAdapterInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCondAdapterInstanceWithDefaults

`func NewCondAdapterInstanceWithDefaults() *CondAdapterInstance`

NewCondAdapterInstanceWithDefaults instantiates a new CondAdapterInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CondAdapterInstance) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CondAdapterInstance) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CondAdapterInstance) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CondAdapterInstance) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CondAdapterInstance) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CondAdapterInstance) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdapter

`func (o *CondAdapterInstance) GetAdapter() MoMoRef`

GetAdapter returns the Adapter field if non-nil, zero value otherwise.

### GetAdapterOk

`func (o *CondAdapterInstance) GetAdapterOk() (*MoMoRef, bool)`

GetAdapterOk returns a tuple with the Adapter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapter

`func (o *CondAdapterInstance) SetAdapter(v MoMoRef)`

SetAdapter sets Adapter field to given value.

### HasAdapter

`func (o *CondAdapterInstance) HasAdapter() bool`

HasAdapter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


