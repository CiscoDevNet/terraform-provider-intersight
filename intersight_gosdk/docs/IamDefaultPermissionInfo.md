# IamDefaultPermissionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.DefaultPermissionInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.DefaultPermissionInfo"]
**Interface** | Pointer to **string** | The shared login interface coverage for the stored default permission. * &#x60;All&#x60; - The stored default permission applies to all currently supported login interfaces. | [optional] [default to "All"]
**Permission** | Pointer to [**MoMoRef**](MoMoRef.md) |  | [optional] 

## Methods

### NewIamDefaultPermissionInfo

`func NewIamDefaultPermissionInfo(classId string, objectType string, ) *IamDefaultPermissionInfo`

NewIamDefaultPermissionInfo instantiates a new IamDefaultPermissionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamDefaultPermissionInfoWithDefaults

`func NewIamDefaultPermissionInfoWithDefaults() *IamDefaultPermissionInfo`

NewIamDefaultPermissionInfoWithDefaults instantiates a new IamDefaultPermissionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamDefaultPermissionInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamDefaultPermissionInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamDefaultPermissionInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamDefaultPermissionInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamDefaultPermissionInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamDefaultPermissionInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetInterface

`func (o *IamDefaultPermissionInfo) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *IamDefaultPermissionInfo) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *IamDefaultPermissionInfo) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *IamDefaultPermissionInfo) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetPermission

`func (o *IamDefaultPermissionInfo) GetPermission() MoMoRef`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *IamDefaultPermissionInfo) GetPermissionOk() (*MoMoRef, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *IamDefaultPermissionInfo) SetPermission(v MoMoRef)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *IamDefaultPermissionInfo) HasPermission() bool`

HasPermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


