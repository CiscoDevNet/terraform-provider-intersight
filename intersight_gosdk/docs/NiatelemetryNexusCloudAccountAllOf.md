# NiatelemetryNexusCloudAccountAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.NexusCloudAccount"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.NexusCloudAccount"]
**AciCount** | Pointer to **int64** | Count of ACI-type site devices. | [optional] 
**NxosCount** | Pointer to **int64** | Count of NXOS-type site devices. | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetryNexusCloudAccountAllOf

`func NewNiatelemetryNexusCloudAccountAllOf(classId string, objectType string, ) *NiatelemetryNexusCloudAccountAllOf`

NewNiatelemetryNexusCloudAccountAllOf instantiates a new NiatelemetryNexusCloudAccountAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryNexusCloudAccountAllOfWithDefaults

`func NewNiatelemetryNexusCloudAccountAllOfWithDefaults() *NiatelemetryNexusCloudAccountAllOf`

NewNiatelemetryNexusCloudAccountAllOfWithDefaults instantiates a new NiatelemetryNexusCloudAccountAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryNexusCloudAccountAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryNexusCloudAccountAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryNexusCloudAccountAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryNexusCloudAccountAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryNexusCloudAccountAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryNexusCloudAccountAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAciCount

`func (o *NiatelemetryNexusCloudAccountAllOf) GetAciCount() int64`

GetAciCount returns the AciCount field if non-nil, zero value otherwise.

### GetAciCountOk

`func (o *NiatelemetryNexusCloudAccountAllOf) GetAciCountOk() (*int64, bool)`

GetAciCountOk returns a tuple with the AciCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAciCount

`func (o *NiatelemetryNexusCloudAccountAllOf) SetAciCount(v int64)`

SetAciCount sets AciCount field to given value.

### HasAciCount

`func (o *NiatelemetryNexusCloudAccountAllOf) HasAciCount() bool`

HasAciCount returns a boolean if a field has been set.

### GetNxosCount

`func (o *NiatelemetryNexusCloudAccountAllOf) GetNxosCount() int64`

GetNxosCount returns the NxosCount field if non-nil, zero value otherwise.

### GetNxosCountOk

`func (o *NiatelemetryNexusCloudAccountAllOf) GetNxosCountOk() (*int64, bool)`

GetNxosCountOk returns a tuple with the NxosCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxosCount

`func (o *NiatelemetryNexusCloudAccountAllOf) SetNxosCount(v int64)`

SetNxosCount sets NxosCount field to given value.

### HasNxosCount

`func (o *NiatelemetryNexusCloudAccountAllOf) HasNxosCount() bool`

HasNxosCount returns a boolean if a field has been set.

### GetAccount

`func (o *NiatelemetryNexusCloudAccountAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *NiatelemetryNexusCloudAccountAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *NiatelemetryNexusCloudAccountAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *NiatelemetryNexusCloudAccountAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


