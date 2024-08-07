# HyperflexVmImportOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.VmImportOperation"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.VmImportOperation"]
**DeviceMoid** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**Organization** | Pointer to [**NullableOrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewHyperflexVmImportOperation

`func NewHyperflexVmImportOperation(classId string, objectType string, ) *HyperflexVmImportOperation`

NewHyperflexVmImportOperation instantiates a new HyperflexVmImportOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexVmImportOperationWithDefaults

`func NewHyperflexVmImportOperationWithDefaults() *HyperflexVmImportOperation`

NewHyperflexVmImportOperationWithDefaults instantiates a new HyperflexVmImportOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexVmImportOperation) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexVmImportOperation) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexVmImportOperation) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexVmImportOperation) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexVmImportOperation) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexVmImportOperation) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDeviceMoid

`func (o *HyperflexVmImportOperation) GetDeviceMoid() AssetDeviceRegistrationRelationship`

GetDeviceMoid returns the DeviceMoid field if non-nil, zero value otherwise.

### GetDeviceMoidOk

`func (o *HyperflexVmImportOperation) GetDeviceMoidOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceMoidOk returns a tuple with the DeviceMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMoid

`func (o *HyperflexVmImportOperation) SetDeviceMoid(v AssetDeviceRegistrationRelationship)`

SetDeviceMoid sets DeviceMoid field to given value.

### HasDeviceMoid

`func (o *HyperflexVmImportOperation) HasDeviceMoid() bool`

HasDeviceMoid returns a boolean if a field has been set.

### SetDeviceMoidNil

`func (o *HyperflexVmImportOperation) SetDeviceMoidNil(b bool)`

 SetDeviceMoidNil sets the value for DeviceMoid to be an explicit nil

### UnsetDeviceMoid
`func (o *HyperflexVmImportOperation) UnsetDeviceMoid()`

UnsetDeviceMoid ensures that no value is present for DeviceMoid, not even an explicit nil
### GetOrganization

`func (o *HyperflexVmImportOperation) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *HyperflexVmImportOperation) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *HyperflexVmImportOperation) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *HyperflexVmImportOperation) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *HyperflexVmImportOperation) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *HyperflexVmImportOperation) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


