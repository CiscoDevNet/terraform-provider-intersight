# CapabilityServerComponentConstraint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.ServerComponentConstraint"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.ServerComponentConstraint"]
**MinSupportedVersion** | Pointer to **string** | The version below which the component is not supported. | [optional] [readonly] 
**ServerModel** | Pointer to **string** | The server model this constraint is to be enforced upon. | [optional] [readonly] 

## Methods

### NewCapabilityServerComponentConstraint

`func NewCapabilityServerComponentConstraint(classId string, objectType string, ) *CapabilityServerComponentConstraint`

NewCapabilityServerComponentConstraint instantiates a new CapabilityServerComponentConstraint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityServerComponentConstraintWithDefaults

`func NewCapabilityServerComponentConstraintWithDefaults() *CapabilityServerComponentConstraint`

NewCapabilityServerComponentConstraintWithDefaults instantiates a new CapabilityServerComponentConstraint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityServerComponentConstraint) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityServerComponentConstraint) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityServerComponentConstraint) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityServerComponentConstraint) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityServerComponentConstraint) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityServerComponentConstraint) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMinSupportedVersion

`func (o *CapabilityServerComponentConstraint) GetMinSupportedVersion() string`

GetMinSupportedVersion returns the MinSupportedVersion field if non-nil, zero value otherwise.

### GetMinSupportedVersionOk

`func (o *CapabilityServerComponentConstraint) GetMinSupportedVersionOk() (*string, bool)`

GetMinSupportedVersionOk returns a tuple with the MinSupportedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSupportedVersion

`func (o *CapabilityServerComponentConstraint) SetMinSupportedVersion(v string)`

SetMinSupportedVersion sets MinSupportedVersion field to given value.

### HasMinSupportedVersion

`func (o *CapabilityServerComponentConstraint) HasMinSupportedVersion() bool`

HasMinSupportedVersion returns a boolean if a field has been set.

### GetServerModel

`func (o *CapabilityServerComponentConstraint) GetServerModel() string`

GetServerModel returns the ServerModel field if non-nil, zero value otherwise.

### GetServerModelOk

`func (o *CapabilityServerComponentConstraint) GetServerModelOk() (*string, bool)`

GetServerModelOk returns a tuple with the ServerModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerModel

`func (o *CapabilityServerComponentConstraint) SetServerModel(v string)`

SetServerModel sets ServerModel field to given value.

### HasServerModel

`func (o *CapabilityServerComponentConstraint) HasServerModel() bool`

HasServerModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


