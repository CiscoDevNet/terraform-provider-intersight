# NiatelemetryCloudRegionsElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.CloudRegionsElement"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.CloudRegionsElement"]
**AdminState** | Pointer to **string** | Return value of adminState attribute. | [optional] 
**CapicDeployed** | Pointer to **string** | Return whether CAPIC is deployed in the cloud region or not. | [optional] 
**InUse** | Pointer to **string** | Return whether any user deployment is configured in the cloud region or not. | [optional] 
**Name** | Pointer to **string** | Return value of name of the cloud region. | [optional] 

## Methods

### NewNiatelemetryCloudRegionsElement

`func NewNiatelemetryCloudRegionsElement(classId string, objectType string, ) *NiatelemetryCloudRegionsElement`

NewNiatelemetryCloudRegionsElement instantiates a new NiatelemetryCloudRegionsElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryCloudRegionsElementWithDefaults

`func NewNiatelemetryCloudRegionsElementWithDefaults() *NiatelemetryCloudRegionsElement`

NewNiatelemetryCloudRegionsElementWithDefaults instantiates a new NiatelemetryCloudRegionsElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryCloudRegionsElement) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryCloudRegionsElement) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryCloudRegionsElement) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryCloudRegionsElement) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryCloudRegionsElement) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryCloudRegionsElement) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminState

`func (o *NiatelemetryCloudRegionsElement) GetAdminState() string`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *NiatelemetryCloudRegionsElement) GetAdminStateOk() (*string, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *NiatelemetryCloudRegionsElement) SetAdminState(v string)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *NiatelemetryCloudRegionsElement) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetCapicDeployed

`func (o *NiatelemetryCloudRegionsElement) GetCapicDeployed() string`

GetCapicDeployed returns the CapicDeployed field if non-nil, zero value otherwise.

### GetCapicDeployedOk

`func (o *NiatelemetryCloudRegionsElement) GetCapicDeployedOk() (*string, bool)`

GetCapicDeployedOk returns a tuple with the CapicDeployed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapicDeployed

`func (o *NiatelemetryCloudRegionsElement) SetCapicDeployed(v string)`

SetCapicDeployed sets CapicDeployed field to given value.

### HasCapicDeployed

`func (o *NiatelemetryCloudRegionsElement) HasCapicDeployed() bool`

HasCapicDeployed returns a boolean if a field has been set.

### GetInUse

`func (o *NiatelemetryCloudRegionsElement) GetInUse() string`

GetInUse returns the InUse field if non-nil, zero value otherwise.

### GetInUseOk

`func (o *NiatelemetryCloudRegionsElement) GetInUseOk() (*string, bool)`

GetInUseOk returns a tuple with the InUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUse

`func (o *NiatelemetryCloudRegionsElement) SetInUse(v string)`

SetInUse sets InUse field to given value.

### HasInUse

`func (o *NiatelemetryCloudRegionsElement) HasInUse() bool`

HasInUse returns a boolean if a field has been set.

### GetName

`func (o *NiatelemetryCloudRegionsElement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NiatelemetryCloudRegionsElement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NiatelemetryCloudRegionsElement) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NiatelemetryCloudRegionsElement) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


