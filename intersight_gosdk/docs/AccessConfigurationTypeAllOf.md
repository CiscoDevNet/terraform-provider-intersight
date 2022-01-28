# AccessConfigurationTypeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "access.ConfigurationType"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "access.ConfigurationType"]
**ConfigureInband** | Pointer to **bool** | This flag enables the use of In-Band configuration for end-point access. | [optional] [default to true]
**ConfigureOutOfBand** | Pointer to **bool** | This flag enables the use of Out-Of-Band configuration for end-point access. | [optional] [default to false]

## Methods

### NewAccessConfigurationTypeAllOf

`func NewAccessConfigurationTypeAllOf(classId string, objectType string, ) *AccessConfigurationTypeAllOf`

NewAccessConfigurationTypeAllOf instantiates a new AccessConfigurationTypeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessConfigurationTypeAllOfWithDefaults

`func NewAccessConfigurationTypeAllOfWithDefaults() *AccessConfigurationTypeAllOf`

NewAccessConfigurationTypeAllOfWithDefaults instantiates a new AccessConfigurationTypeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AccessConfigurationTypeAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AccessConfigurationTypeAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AccessConfigurationTypeAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AccessConfigurationTypeAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AccessConfigurationTypeAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AccessConfigurationTypeAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConfigureInband

`func (o *AccessConfigurationTypeAllOf) GetConfigureInband() bool`

GetConfigureInband returns the ConfigureInband field if non-nil, zero value otherwise.

### GetConfigureInbandOk

`func (o *AccessConfigurationTypeAllOf) GetConfigureInbandOk() (*bool, bool)`

GetConfigureInbandOk returns a tuple with the ConfigureInband field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigureInband

`func (o *AccessConfigurationTypeAllOf) SetConfigureInband(v bool)`

SetConfigureInband sets ConfigureInband field to given value.

### HasConfigureInband

`func (o *AccessConfigurationTypeAllOf) HasConfigureInband() bool`

HasConfigureInband returns a boolean if a field has been set.

### GetConfigureOutOfBand

`func (o *AccessConfigurationTypeAllOf) GetConfigureOutOfBand() bool`

GetConfigureOutOfBand returns the ConfigureOutOfBand field if non-nil, zero value otherwise.

### GetConfigureOutOfBandOk

`func (o *AccessConfigurationTypeAllOf) GetConfigureOutOfBandOk() (*bool, bool)`

GetConfigureOutOfBandOk returns a tuple with the ConfigureOutOfBand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigureOutOfBand

`func (o *AccessConfigurationTypeAllOf) SetConfigureOutOfBand(v bool)`

SetConfigureOutOfBand sets ConfigureOutOfBand field to given value.

### HasConfigureOutOfBand

`func (o *AccessConfigurationTypeAllOf) HasConfigureOutOfBand() bool`

HasConfigureOutOfBand returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


