# ApicExternalRoutedLayerThreeDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "apic.ExternalRoutedLayerThreeDomain"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "apic.ExternalRoutedLayerThreeDomain"]
**Dn** | Pointer to **string** | Distinguished Name (DN) of an object within the Cisco Application Policy Infrastructure Controller (APIC) GUI. | [optional] 
**Name** | Pointer to **string** | Name of an object within the Cisco Application Policy Infrastructure Controller (APIC) GUI. | [optional] 

## Methods

### NewApicExternalRoutedLayerThreeDomain

`func NewApicExternalRoutedLayerThreeDomain(classId string, objectType string, ) *ApicExternalRoutedLayerThreeDomain`

NewApicExternalRoutedLayerThreeDomain instantiates a new ApicExternalRoutedLayerThreeDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApicExternalRoutedLayerThreeDomainWithDefaults

`func NewApicExternalRoutedLayerThreeDomainWithDefaults() *ApicExternalRoutedLayerThreeDomain`

NewApicExternalRoutedLayerThreeDomainWithDefaults instantiates a new ApicExternalRoutedLayerThreeDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApicExternalRoutedLayerThreeDomain) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApicExternalRoutedLayerThreeDomain) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApicExternalRoutedLayerThreeDomain) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApicExternalRoutedLayerThreeDomain) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApicExternalRoutedLayerThreeDomain) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApicExternalRoutedLayerThreeDomain) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDn

`func (o *ApicExternalRoutedLayerThreeDomain) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *ApicExternalRoutedLayerThreeDomain) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *ApicExternalRoutedLayerThreeDomain) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *ApicExternalRoutedLayerThreeDomain) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetName

`func (o *ApicExternalRoutedLayerThreeDomain) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApicExternalRoutedLayerThreeDomain) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApicExternalRoutedLayerThreeDomain) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApicExternalRoutedLayerThreeDomain) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


