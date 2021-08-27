# HyperflexServerFirmwareVersionEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.ServerFirmwareVersionEntry"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.ServerFirmwareVersionEntry"]
**Constraint** | Pointer to [**NullableHyperflexAppSettingConstraint**](HyperflexAppSettingConstraint.md) |  | [optional] 
**ServerPlatform** | Pointer to **string** | The server platform type that is applicable for the server firmware bundle version. * &#x60;M5&#x60; - M5 generation of UCS server. * &#x60;M3&#x60; - M3 generation of UCS server. * &#x60;M4&#x60; - M4 generation of UCS server. | [optional] [default to "M5"]
**Version** | Pointer to **string** | The server firmware bundle version. | [optional] 
**ServerFirmwareVersion** | Pointer to [**HyperflexServerFirmwareVersionRelationship**](HyperflexServerFirmwareVersionRelationship.md) |  | [optional] 

## Methods

### NewHyperflexServerFirmwareVersionEntry

`func NewHyperflexServerFirmwareVersionEntry(classId string, objectType string, ) *HyperflexServerFirmwareVersionEntry`

NewHyperflexServerFirmwareVersionEntry instantiates a new HyperflexServerFirmwareVersionEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexServerFirmwareVersionEntryWithDefaults

`func NewHyperflexServerFirmwareVersionEntryWithDefaults() *HyperflexServerFirmwareVersionEntry`

NewHyperflexServerFirmwareVersionEntryWithDefaults instantiates a new HyperflexServerFirmwareVersionEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexServerFirmwareVersionEntry) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexServerFirmwareVersionEntry) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexServerFirmwareVersionEntry) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexServerFirmwareVersionEntry) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexServerFirmwareVersionEntry) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexServerFirmwareVersionEntry) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConstraint

`func (o *HyperflexServerFirmwareVersionEntry) GetConstraint() HyperflexAppSettingConstraint`

GetConstraint returns the Constraint field if non-nil, zero value otherwise.

### GetConstraintOk

`func (o *HyperflexServerFirmwareVersionEntry) GetConstraintOk() (*HyperflexAppSettingConstraint, bool)`

GetConstraintOk returns a tuple with the Constraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraint

`func (o *HyperflexServerFirmwareVersionEntry) SetConstraint(v HyperflexAppSettingConstraint)`

SetConstraint sets Constraint field to given value.

### HasConstraint

`func (o *HyperflexServerFirmwareVersionEntry) HasConstraint() bool`

HasConstraint returns a boolean if a field has been set.

### SetConstraintNil

`func (o *HyperflexServerFirmwareVersionEntry) SetConstraintNil(b bool)`

 SetConstraintNil sets the value for Constraint to be an explicit nil

### UnsetConstraint
`func (o *HyperflexServerFirmwareVersionEntry) UnsetConstraint()`

UnsetConstraint ensures that no value is present for Constraint, not even an explicit nil
### GetServerPlatform

`func (o *HyperflexServerFirmwareVersionEntry) GetServerPlatform() string`

GetServerPlatform returns the ServerPlatform field if non-nil, zero value otherwise.

### GetServerPlatformOk

`func (o *HyperflexServerFirmwareVersionEntry) GetServerPlatformOk() (*string, bool)`

GetServerPlatformOk returns a tuple with the ServerPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPlatform

`func (o *HyperflexServerFirmwareVersionEntry) SetServerPlatform(v string)`

SetServerPlatform sets ServerPlatform field to given value.

### HasServerPlatform

`func (o *HyperflexServerFirmwareVersionEntry) HasServerPlatform() bool`

HasServerPlatform returns a boolean if a field has been set.

### GetVersion

`func (o *HyperflexServerFirmwareVersionEntry) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HyperflexServerFirmwareVersionEntry) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HyperflexServerFirmwareVersionEntry) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HyperflexServerFirmwareVersionEntry) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetServerFirmwareVersion

`func (o *HyperflexServerFirmwareVersionEntry) GetServerFirmwareVersion() HyperflexServerFirmwareVersionRelationship`

GetServerFirmwareVersion returns the ServerFirmwareVersion field if non-nil, zero value otherwise.

### GetServerFirmwareVersionOk

`func (o *HyperflexServerFirmwareVersionEntry) GetServerFirmwareVersionOk() (*HyperflexServerFirmwareVersionRelationship, bool)`

GetServerFirmwareVersionOk returns a tuple with the ServerFirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareVersion

`func (o *HyperflexServerFirmwareVersionEntry) SetServerFirmwareVersion(v HyperflexServerFirmwareVersionRelationship)`

SetServerFirmwareVersion sets ServerFirmwareVersion field to given value.

### HasServerFirmwareVersion

`func (o *HyperflexServerFirmwareVersionEntry) HasServerFirmwareVersion() bool`

HasServerFirmwareVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


