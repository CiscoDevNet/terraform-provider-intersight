# CrdCustomResourceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "crd.CustomResource"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "crd.CustomResource"]
**DcLauncher** | Pointer to **string** | Type of custom resource or &#39;kind&#39; in K8s. | [optional] 
**Image** | Pointer to **string** | Docker image URL for public cloud DC. | [optional] 
**Name** | Pointer to **string** | A string property called name which is used as part of a uniqueness constraint. See &#39;identity&#39; specification in this MO definition. | [optional] 
**Namespace** | Pointer to **string** | Namespace to launch the deployment associated with the custom resource. | [optional] 
**Port** | Pointer to **int64** | Port used for public cloud DC. | [optional] 
**Properties** | Pointer to [**[]CrdCustomResourceConfigProperty**](CrdCustomResourceConfigProperty.md) |  | [optional] 
**TargetId** | Pointer to **string** | Target ID for public cloud DC. | [optional] 
**TargetMoid** | Pointer to **string** | Target Moid for public cloud DC. | [optional] 
**TargetType** | Pointer to **string** | Target type for public cloud DC. | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 

## Methods

### NewCrdCustomResourceAllOf

`func NewCrdCustomResourceAllOf(classId string, objectType string, ) *CrdCustomResourceAllOf`

NewCrdCustomResourceAllOf instantiates a new CrdCustomResourceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrdCustomResourceAllOfWithDefaults

`func NewCrdCustomResourceAllOfWithDefaults() *CrdCustomResourceAllOf`

NewCrdCustomResourceAllOfWithDefaults instantiates a new CrdCustomResourceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CrdCustomResourceAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CrdCustomResourceAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CrdCustomResourceAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CrdCustomResourceAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CrdCustomResourceAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CrdCustomResourceAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDcLauncher

`func (o *CrdCustomResourceAllOf) GetDcLauncher() string`

GetDcLauncher returns the DcLauncher field if non-nil, zero value otherwise.

### GetDcLauncherOk

`func (o *CrdCustomResourceAllOf) GetDcLauncherOk() (*string, bool)`

GetDcLauncherOk returns a tuple with the DcLauncher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcLauncher

`func (o *CrdCustomResourceAllOf) SetDcLauncher(v string)`

SetDcLauncher sets DcLauncher field to given value.

### HasDcLauncher

`func (o *CrdCustomResourceAllOf) HasDcLauncher() bool`

HasDcLauncher returns a boolean if a field has been set.

### GetImage

`func (o *CrdCustomResourceAllOf) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *CrdCustomResourceAllOf) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *CrdCustomResourceAllOf) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *CrdCustomResourceAllOf) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetName

`func (o *CrdCustomResourceAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CrdCustomResourceAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CrdCustomResourceAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CrdCustomResourceAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *CrdCustomResourceAllOf) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *CrdCustomResourceAllOf) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *CrdCustomResourceAllOf) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *CrdCustomResourceAllOf) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetPort

`func (o *CrdCustomResourceAllOf) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CrdCustomResourceAllOf) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CrdCustomResourceAllOf) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *CrdCustomResourceAllOf) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProperties

`func (o *CrdCustomResourceAllOf) GetProperties() []CrdCustomResourceConfigProperty`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CrdCustomResourceAllOf) GetPropertiesOk() (*[]CrdCustomResourceConfigProperty, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CrdCustomResourceAllOf) SetProperties(v []CrdCustomResourceConfigProperty)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CrdCustomResourceAllOf) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *CrdCustomResourceAllOf) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *CrdCustomResourceAllOf) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetTargetId

`func (o *CrdCustomResourceAllOf) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *CrdCustomResourceAllOf) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *CrdCustomResourceAllOf) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *CrdCustomResourceAllOf) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### GetTargetMoid

`func (o *CrdCustomResourceAllOf) GetTargetMoid() string`

GetTargetMoid returns the TargetMoid field if non-nil, zero value otherwise.

### GetTargetMoidOk

`func (o *CrdCustomResourceAllOf) GetTargetMoidOk() (*string, bool)`

GetTargetMoidOk returns a tuple with the TargetMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMoid

`func (o *CrdCustomResourceAllOf) SetTargetMoid(v string)`

SetTargetMoid sets TargetMoid field to given value.

### HasTargetMoid

`func (o *CrdCustomResourceAllOf) HasTargetMoid() bool`

HasTargetMoid returns a boolean if a field has been set.

### GetTargetType

`func (o *CrdCustomResourceAllOf) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *CrdCustomResourceAllOf) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *CrdCustomResourceAllOf) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *CrdCustomResourceAllOf) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### GetAccount

`func (o *CrdCustomResourceAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CrdCustomResourceAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CrdCustomResourceAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CrdCustomResourceAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


