# ConvergedinfraPodComplianceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "convergedinfra.PodComplianceInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "convergedinfra.PodComplianceInfo"]
**Details** | Pointer to [**[]ConvergedinfraBaseComplianceDetailsRelationship**](ConvergedinfraBaseComplianceDetailsRelationship.md) | An array of relationships to convergedinfraBaseComplianceDetails resources. | [optional] [readonly] 
**Pod** | Pointer to [**NullableConvergedinfraPodRelationship**](ConvergedinfraPodRelationship.md) |  | [optional] 

## Methods

### NewConvergedinfraPodComplianceInfo

`func NewConvergedinfraPodComplianceInfo(classId string, objectType string, ) *ConvergedinfraPodComplianceInfo`

NewConvergedinfraPodComplianceInfo instantiates a new ConvergedinfraPodComplianceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConvergedinfraPodComplianceInfoWithDefaults

`func NewConvergedinfraPodComplianceInfoWithDefaults() *ConvergedinfraPodComplianceInfo`

NewConvergedinfraPodComplianceInfoWithDefaults instantiates a new ConvergedinfraPodComplianceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ConvergedinfraPodComplianceInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ConvergedinfraPodComplianceInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ConvergedinfraPodComplianceInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ConvergedinfraPodComplianceInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConvergedinfraPodComplianceInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConvergedinfraPodComplianceInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDetails

`func (o *ConvergedinfraPodComplianceInfo) GetDetails() []ConvergedinfraBaseComplianceDetailsRelationship`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ConvergedinfraPodComplianceInfo) GetDetailsOk() (*[]ConvergedinfraBaseComplianceDetailsRelationship, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ConvergedinfraPodComplianceInfo) SetDetails(v []ConvergedinfraBaseComplianceDetailsRelationship)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ConvergedinfraPodComplianceInfo) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *ConvergedinfraPodComplianceInfo) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *ConvergedinfraPodComplianceInfo) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetPod

`func (o *ConvergedinfraPodComplianceInfo) GetPod() ConvergedinfraPodRelationship`

GetPod returns the Pod field if non-nil, zero value otherwise.

### GetPodOk

`func (o *ConvergedinfraPodComplianceInfo) GetPodOk() (*ConvergedinfraPodRelationship, bool)`

GetPodOk returns a tuple with the Pod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPod

`func (o *ConvergedinfraPodComplianceInfo) SetPod(v ConvergedinfraPodRelationship)`

SetPod sets Pod field to given value.

### HasPod

`func (o *ConvergedinfraPodComplianceInfo) HasPod() bool`

HasPod returns a boolean if a field has been set.

### SetPodNil

`func (o *ConvergedinfraPodComplianceInfo) SetPodNil(b bool)`

 SetPodNil sets the value for Pod to be an explicit nil

### UnsetPod
`func (o *ConvergedinfraPodComplianceInfo) UnsetPod()`

UnsetPod ensures that no value is present for Pod, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


