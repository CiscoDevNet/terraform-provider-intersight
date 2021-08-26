# HyperflexHxapDatacenterAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.HxapDatacenter"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.HxapDatacenter"]
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 
**HypervisorManager** | Pointer to [**HyperflexCiscoHypervisorManagerRelationship**](HyperflexCiscoHypervisorManagerRelationship.md) |  | [optional] 

## Methods

### NewHyperflexHxapDatacenterAllOf

`func NewHyperflexHxapDatacenterAllOf(classId string, objectType string, ) *HyperflexHxapDatacenterAllOf`

NewHyperflexHxapDatacenterAllOf instantiates a new HyperflexHxapDatacenterAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexHxapDatacenterAllOfWithDefaults

`func NewHyperflexHxapDatacenterAllOfWithDefaults() *HyperflexHxapDatacenterAllOf`

NewHyperflexHxapDatacenterAllOfWithDefaults instantiates a new HyperflexHxapDatacenterAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexHxapDatacenterAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexHxapDatacenterAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexHxapDatacenterAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexHxapDatacenterAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexHxapDatacenterAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexHxapDatacenterAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAccount

`func (o *HyperflexHxapDatacenterAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *HyperflexHxapDatacenterAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *HyperflexHxapDatacenterAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *HyperflexHxapDatacenterAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetHypervisorManager

`func (o *HyperflexHxapDatacenterAllOf) GetHypervisorManager() HyperflexCiscoHypervisorManagerRelationship`

GetHypervisorManager returns the HypervisorManager field if non-nil, zero value otherwise.

### GetHypervisorManagerOk

`func (o *HyperflexHxapDatacenterAllOf) GetHypervisorManagerOk() (*HyperflexCiscoHypervisorManagerRelationship, bool)`

GetHypervisorManagerOk returns a tuple with the HypervisorManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorManager

`func (o *HyperflexHxapDatacenterAllOf) SetHypervisorManager(v HyperflexCiscoHypervisorManagerRelationship)`

SetHypervisorManager sets HypervisorManager field to given value.

### HasHypervisorManager

`func (o *HyperflexHxapDatacenterAllOf) HasHypervisorManager() bool`

HasHypervisorManager returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


