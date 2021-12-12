# NiatelemetryDeploymentStatusAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.DeploymentStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.DeploymentStatus"]
**Id** | Pointer to **int64** | Returns the id of network/vrf. | [optional] 
**Name** | Pointer to **string** | Returns the name of network/vrf. | [optional] 
**Status** | Pointer to **string** | Returns the deployment status of network/vrf. | [optional] 

## Methods

### NewNiatelemetryDeploymentStatusAllOf

`func NewNiatelemetryDeploymentStatusAllOf(classId string, objectType string, ) *NiatelemetryDeploymentStatusAllOf`

NewNiatelemetryDeploymentStatusAllOf instantiates a new NiatelemetryDeploymentStatusAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryDeploymentStatusAllOfWithDefaults

`func NewNiatelemetryDeploymentStatusAllOfWithDefaults() *NiatelemetryDeploymentStatusAllOf`

NewNiatelemetryDeploymentStatusAllOfWithDefaults instantiates a new NiatelemetryDeploymentStatusAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryDeploymentStatusAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryDeploymentStatusAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryDeploymentStatusAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryDeploymentStatusAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryDeploymentStatusAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryDeploymentStatusAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetId

`func (o *NiatelemetryDeploymentStatusAllOf) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NiatelemetryDeploymentStatusAllOf) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NiatelemetryDeploymentStatusAllOf) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NiatelemetryDeploymentStatusAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *NiatelemetryDeploymentStatusAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NiatelemetryDeploymentStatusAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NiatelemetryDeploymentStatusAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NiatelemetryDeploymentStatusAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *NiatelemetryDeploymentStatusAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NiatelemetryDeploymentStatusAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NiatelemetryDeploymentStatusAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NiatelemetryDeploymentStatusAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


