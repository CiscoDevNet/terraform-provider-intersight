# HyperflexServerFirmwareVersionEntryAllOf

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

### NewHyperflexServerFirmwareVersionEntryAllOf

`func NewHyperflexServerFirmwareVersionEntryAllOf(classId string, objectType string, ) *HyperflexServerFirmwareVersionEntryAllOf`

NewHyperflexServerFirmwareVersionEntryAllOf instantiates a new HyperflexServerFirmwareVersionEntryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexServerFirmwareVersionEntryAllOfWithDefaults

`func NewHyperflexServerFirmwareVersionEntryAllOfWithDefaults() *HyperflexServerFirmwareVersionEntryAllOf`

NewHyperflexServerFirmwareVersionEntryAllOfWithDefaults instantiates a new HyperflexServerFirmwareVersionEntryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexServerFirmwareVersionEntryAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexServerFirmwareVersionEntryAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexServerFirmwareVersionEntryAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexServerFirmwareVersionEntryAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexServerFirmwareVersionEntryAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexServerFirmwareVersionEntryAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConstraint

`func (o *HyperflexServerFirmwareVersionEntryAllOf) GetConstraint() HyperflexAppSettingConstraint`

GetConstraint returns the Constraint field if non-nil, zero value otherwise.

### GetConstraintOk

`func (o *HyperflexServerFirmwareVersionEntryAllOf) GetConstraintOk() (*HyperflexAppSettingConstraint, bool)`

GetConstraintOk returns a tuple with the Constraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraint

`func (o *HyperflexServerFirmwareVersionEntryAllOf) SetConstraint(v HyperflexAppSettingConstraint)`

SetConstraint sets Constraint field to given value.

### HasConstraint

`func (o *HyperflexServerFirmwareVersionEntryAllOf) HasConstraint() bool`

HasConstraint returns a boolean if a field has been set.

### SetConstraintNil

`func (o *HyperflexServerFirmwareVersionEntryAllOf) SetConstraintNil(b bool)`

 SetConstraintNil sets the value for Constraint to be an explicit nil

### UnsetConstraint
`func (o *HyperflexServerFirmwareVersionEntryAllOf) UnsetConstraint()`

UnsetConstraint ensures that no value is present for Constraint, not even an explicit nil
### GetServerPlatform

`func (o *HyperflexServerFirmwareVersionEntryAllOf) GetServerPlatform() string`

GetServerPlatform returns the ServerPlatform field if non-nil, zero value otherwise.

### GetServerPlatformOk

`func (o *HyperflexServerFirmwareVersionEntryAllOf) GetServerPlatformOk() (*string, bool)`

GetServerPlatformOk returns a tuple with the ServerPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPlatform

`func (o *HyperflexServerFirmwareVersionEntryAllOf) SetServerPlatform(v string)`

SetServerPlatform sets ServerPlatform field to given value.

### HasServerPlatform

`func (o *HyperflexServerFirmwareVersionEntryAllOf) HasServerPlatform() bool`

HasServerPlatform returns a boolean if a field has been set.

### GetVersion

`func (o *HyperflexServerFirmwareVersionEntryAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HyperflexServerFirmwareVersionEntryAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HyperflexServerFirmwareVersionEntryAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HyperflexServerFirmwareVersionEntryAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetServerFirmwareVersion

`func (o *HyperflexServerFirmwareVersionEntryAllOf) GetServerFirmwareVersion() HyperflexServerFirmwareVersionRelationship`

GetServerFirmwareVersion returns the ServerFirmwareVersion field if non-nil, zero value otherwise.

### GetServerFirmwareVersionOk

`func (o *HyperflexServerFirmwareVersionEntryAllOf) GetServerFirmwareVersionOk() (*HyperflexServerFirmwareVersionRelationship, bool)`

GetServerFirmwareVersionOk returns a tuple with the ServerFirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareVersion

`func (o *HyperflexServerFirmwareVersionEntryAllOf) SetServerFirmwareVersion(v HyperflexServerFirmwareVersionRelationship)`

SetServerFirmwareVersion sets ServerFirmwareVersion field to given value.

### HasServerFirmwareVersion

`func (o *HyperflexServerFirmwareVersionEntryAllOf) HasServerFirmwareVersion() bool`

HasServerFirmwareVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


