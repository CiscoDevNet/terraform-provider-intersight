# OsInstallTargetResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**SourceMo** | Pointer to [**MoMoRef**](mo.MoRef.md) |  | [optional] 

## Methods

### NewOsInstallTargetResponseAllOf

`func NewOsInstallTargetResponseAllOf(classId string, objectType string, ) *OsInstallTargetResponseAllOf`

NewOsInstallTargetResponseAllOf instantiates a new OsInstallTargetResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsInstallTargetResponseAllOfWithDefaults

`func NewOsInstallTargetResponseAllOfWithDefaults() *OsInstallTargetResponseAllOf`

NewOsInstallTargetResponseAllOfWithDefaults instantiates a new OsInstallTargetResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OsInstallTargetResponseAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OsInstallTargetResponseAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OsInstallTargetResponseAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OsInstallTargetResponseAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OsInstallTargetResponseAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OsInstallTargetResponseAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSourceMo

`func (o *OsInstallTargetResponseAllOf) GetSourceMo() MoMoRef`

GetSourceMo returns the SourceMo field if non-nil, zero value otherwise.

### GetSourceMoOk

`func (o *OsInstallTargetResponseAllOf) GetSourceMoOk() (*MoMoRef, bool)`

GetSourceMoOk returns a tuple with the SourceMo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMo

`func (o *OsInstallTargetResponseAllOf) SetSourceMo(v MoMoRef)`

SetSourceMo sets SourceMo field to given value.

### HasSourceMo

`func (o *OsInstallTargetResponseAllOf) HasSourceMo() bool`

HasSourceMo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


