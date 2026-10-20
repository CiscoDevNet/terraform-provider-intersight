# DnacTransit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "dnac.Transit"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "dnac.Transit"]
**TransitId** | Pointer to **string** | Identification for the Transit. | [optional] 
**TransitName** | Pointer to **string** | Name identifier for the Transit. | [optional] 
**TransitType** | Pointer to **string** | Transit type for the transit. | [optional] 

## Methods

### NewDnacTransit

`func NewDnacTransit(classId string, objectType string, ) *DnacTransit`

NewDnacTransit instantiates a new DnacTransit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnacTransitWithDefaults

`func NewDnacTransitWithDefaults() *DnacTransit`

NewDnacTransitWithDefaults instantiates a new DnacTransit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *DnacTransit) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *DnacTransit) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *DnacTransit) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *DnacTransit) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *DnacTransit) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *DnacTransit) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetTransitId

`func (o *DnacTransit) GetTransitId() string`

GetTransitId returns the TransitId field if non-nil, zero value otherwise.

### GetTransitIdOk

`func (o *DnacTransit) GetTransitIdOk() (*string, bool)`

GetTransitIdOk returns a tuple with the TransitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitId

`func (o *DnacTransit) SetTransitId(v string)`

SetTransitId sets TransitId field to given value.

### HasTransitId

`func (o *DnacTransit) HasTransitId() bool`

HasTransitId returns a boolean if a field has been set.

### GetTransitName

`func (o *DnacTransit) GetTransitName() string`

GetTransitName returns the TransitName field if non-nil, zero value otherwise.

### GetTransitNameOk

`func (o *DnacTransit) GetTransitNameOk() (*string, bool)`

GetTransitNameOk returns a tuple with the TransitName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitName

`func (o *DnacTransit) SetTransitName(v string)`

SetTransitName sets TransitName field to given value.

### HasTransitName

`func (o *DnacTransit) HasTransitName() bool`

HasTransitName returns a boolean if a field has been set.

### GetTransitType

`func (o *DnacTransit) GetTransitType() string`

GetTransitType returns the TransitType field if non-nil, zero value otherwise.

### GetTransitTypeOk

`func (o *DnacTransit) GetTransitTypeOk() (*string, bool)`

GetTransitTypeOk returns a tuple with the TransitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitType

`func (o *DnacTransit) SetTransitType(v string)`

SetTransitType sets TransitType field to given value.

### HasTransitType

`func (o *DnacTransit) HasTransitType() bool`

HasTransitType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


