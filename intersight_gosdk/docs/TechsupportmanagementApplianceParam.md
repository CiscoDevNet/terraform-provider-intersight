# TechsupportmanagementApplianceParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "techsupportmanagement.ApplianceParam"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "techsupportmanagement.ApplianceParam"]
**IsApplianceRequest** | Pointer to **bool** | Specifies whether the techsupport request is from the cloud or by the appliance. | [optional] 

## Methods

### NewTechsupportmanagementApplianceParam

`func NewTechsupportmanagementApplianceParam(classId string, objectType string, ) *TechsupportmanagementApplianceParam`

NewTechsupportmanagementApplianceParam instantiates a new TechsupportmanagementApplianceParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTechsupportmanagementApplianceParamWithDefaults

`func NewTechsupportmanagementApplianceParamWithDefaults() *TechsupportmanagementApplianceParam`

NewTechsupportmanagementApplianceParamWithDefaults instantiates a new TechsupportmanagementApplianceParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TechsupportmanagementApplianceParam) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TechsupportmanagementApplianceParam) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TechsupportmanagementApplianceParam) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TechsupportmanagementApplianceParam) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TechsupportmanagementApplianceParam) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TechsupportmanagementApplianceParam) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIsApplianceRequest

`func (o *TechsupportmanagementApplianceParam) GetIsApplianceRequest() bool`

GetIsApplianceRequest returns the IsApplianceRequest field if non-nil, zero value otherwise.

### GetIsApplianceRequestOk

`func (o *TechsupportmanagementApplianceParam) GetIsApplianceRequestOk() (*bool, bool)`

GetIsApplianceRequestOk returns a tuple with the IsApplianceRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApplianceRequest

`func (o *TechsupportmanagementApplianceParam) SetIsApplianceRequest(v bool)`

SetIsApplianceRequest sets IsApplianceRequest field to given value.

### HasIsApplianceRequest

`func (o *TechsupportmanagementApplianceParam) HasIsApplianceRequest() bool`

HasIsApplianceRequest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


