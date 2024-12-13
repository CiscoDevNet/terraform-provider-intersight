# ResourceDomainQualifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "resource.DomainQualifier"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "resource.DomainQualifier"]
**DomainNames** | Pointer to **[]string** |  | [optional] 
**FabricInterConnectPids** | Pointer to **[]string** |  | [optional] 

## Methods

### NewResourceDomainQualifier

`func NewResourceDomainQualifier(classId string, objectType string, ) *ResourceDomainQualifier`

NewResourceDomainQualifier instantiates a new ResourceDomainQualifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceDomainQualifierWithDefaults

`func NewResourceDomainQualifierWithDefaults() *ResourceDomainQualifier`

NewResourceDomainQualifierWithDefaults instantiates a new ResourceDomainQualifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ResourceDomainQualifier) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ResourceDomainQualifier) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ResourceDomainQualifier) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ResourceDomainQualifier) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ResourceDomainQualifier) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ResourceDomainQualifier) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDomainNames

`func (o *ResourceDomainQualifier) GetDomainNames() []string`

GetDomainNames returns the DomainNames field if non-nil, zero value otherwise.

### GetDomainNamesOk

`func (o *ResourceDomainQualifier) GetDomainNamesOk() (*[]string, bool)`

GetDomainNamesOk returns a tuple with the DomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNames

`func (o *ResourceDomainQualifier) SetDomainNames(v []string)`

SetDomainNames sets DomainNames field to given value.

### HasDomainNames

`func (o *ResourceDomainQualifier) HasDomainNames() bool`

HasDomainNames returns a boolean if a field has been set.

### SetDomainNamesNil

`func (o *ResourceDomainQualifier) SetDomainNamesNil(b bool)`

 SetDomainNamesNil sets the value for DomainNames to be an explicit nil

### UnsetDomainNames
`func (o *ResourceDomainQualifier) UnsetDomainNames()`

UnsetDomainNames ensures that no value is present for DomainNames, not even an explicit nil
### GetFabricInterConnectPids

`func (o *ResourceDomainQualifier) GetFabricInterConnectPids() []string`

GetFabricInterConnectPids returns the FabricInterConnectPids field if non-nil, zero value otherwise.

### GetFabricInterConnectPidsOk

`func (o *ResourceDomainQualifier) GetFabricInterConnectPidsOk() (*[]string, bool)`

GetFabricInterConnectPidsOk returns a tuple with the FabricInterConnectPids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricInterConnectPids

`func (o *ResourceDomainQualifier) SetFabricInterConnectPids(v []string)`

SetFabricInterConnectPids sets FabricInterConnectPids field to given value.

### HasFabricInterConnectPids

`func (o *ResourceDomainQualifier) HasFabricInterConnectPids() bool`

HasFabricInterConnectPids returns a boolean if a field has been set.

### SetFabricInterConnectPidsNil

`func (o *ResourceDomainQualifier) SetFabricInterConnectPidsNil(b bool)`

 SetFabricInterConnectPidsNil sets the value for FabricInterConnectPids to be an explicit nil

### UnsetFabricInterConnectPids
`func (o *ResourceDomainQualifier) UnsetFabricInterConnectPids()`

UnsetFabricInterConnectPids ensures that no value is present for FabricInterConnectPids, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


