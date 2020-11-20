# MetaPropDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "meta.PropDefinition"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "meta.PropDefinition"]
**ApiAccess** | Pointer to **string** | API access control for the property. Examples are NoAccess, ReadOnly, CreateOnly etc. * &#x60;NoAccess&#x60; - The property is not accessible from the API. * &#x60;ReadOnly&#x60; - The value of the property is read-only.An HTTP 4xx status code is returned when the user sends a POST/PUT/PATCH request that containsa ReadOnly property. * &#x60;CreateOnly&#x60; - The value of the property can be set when the REST resource is created. It cannot be changed after object creation.An HTTP 4xx status code is returned when the user sends a POST/PUT/PATCH request that containsa CreateOnly property.CreateOnly properties are returned in the response body of HTTP GET requests. * &#x60;ReadWrite&#x60; - The property has read/write access. * &#x60;WriteOnly&#x60; - The value of the property can be set but it is never returned in the response body of supported HTTP methods.This settings is used for sensitive properties such as passwords. * &#x60;ReadOnCreate&#x60; - The value of the property is returned in the HTTP POST response body when the REST resource is created.The property is not writeable and cannot be queried through a GET request after the resource has been created. | [optional] [readonly] [default to "NoAccess"]
**Name** | Pointer to **string** | The name of the property. | [optional] [readonly] 
**OpSecurity** | Pointer to **string** | The data-at-rest security protection applied to this property when a Managed Object is persisted. For example, Encrypted or Cleartext. * &#x60;ClearText&#x60; - Data at rest is not encrypted using account specific keys.Note that data is always protected using volume encryption. ClearText propertiesare queryable and searchable. * &#x60;Encrypted&#x60; - The value of the property is encrypted with account-specific cryptographic keys.This setting is used for properties that contain sensitive data. Encrypted propertiesare not queryable and searchable. * &#x60;Pbkdf2&#x60; - The value of the property is hashed using the pbkdf2 key derivation functions thattakes a password, a salt, and a cost factor as inputs then generates a password hash.Its purpose is to make each password guessing trial by an attacker who has obtaineda password hash file expensive and therefore the cost of a guessing attack high or prohibitive. * &#x60;Bcrypt&#x60; - The value of the property is hashed using the bcrypt key derivation function. * &#x60;Sha512crypt&#x60; - The value of the property is hashed using the sha512crypt key derivation function. * &#x60;Argon2id&#x60; - The value of the property is hashed using the argon2id key derivation function. | [optional] [readonly] [default to "ClearText"]
**SearchWeight** | Pointer to **float32** | Enables the property to be searchable from global search. A value of 0 means this property is not globally searchable. | [optional] [readonly] 

## Methods

### NewMetaPropDefinition

`func NewMetaPropDefinition(classId string, objectType string, ) *MetaPropDefinition`

NewMetaPropDefinition instantiates a new MetaPropDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaPropDefinitionWithDefaults

`func NewMetaPropDefinitionWithDefaults() *MetaPropDefinition`

NewMetaPropDefinitionWithDefaults instantiates a new MetaPropDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MetaPropDefinition) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MetaPropDefinition) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MetaPropDefinition) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MetaPropDefinition) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MetaPropDefinition) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MetaPropDefinition) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetApiAccess

`func (o *MetaPropDefinition) GetApiAccess() string`

GetApiAccess returns the ApiAccess field if non-nil, zero value otherwise.

### GetApiAccessOk

`func (o *MetaPropDefinition) GetApiAccessOk() (*string, bool)`

GetApiAccessOk returns a tuple with the ApiAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiAccess

`func (o *MetaPropDefinition) SetApiAccess(v string)`

SetApiAccess sets ApiAccess field to given value.

### HasApiAccess

`func (o *MetaPropDefinition) HasApiAccess() bool`

HasApiAccess returns a boolean if a field has been set.

### GetName

`func (o *MetaPropDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetaPropDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetaPropDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MetaPropDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOpSecurity

`func (o *MetaPropDefinition) GetOpSecurity() string`

GetOpSecurity returns the OpSecurity field if non-nil, zero value otherwise.

### GetOpSecurityOk

`func (o *MetaPropDefinition) GetOpSecurityOk() (*string, bool)`

GetOpSecurityOk returns a tuple with the OpSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpSecurity

`func (o *MetaPropDefinition) SetOpSecurity(v string)`

SetOpSecurity sets OpSecurity field to given value.

### HasOpSecurity

`func (o *MetaPropDefinition) HasOpSecurity() bool`

HasOpSecurity returns a boolean if a field has been set.

### GetSearchWeight

`func (o *MetaPropDefinition) GetSearchWeight() float32`

GetSearchWeight returns the SearchWeight field if non-nil, zero value otherwise.

### GetSearchWeightOk

`func (o *MetaPropDefinition) GetSearchWeightOk() (*float32, bool)`

GetSearchWeightOk returns a tuple with the SearchWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchWeight

`func (o *MetaPropDefinition) SetSearchWeight(v float32)`

SetSearchWeight sets SearchWeight field to given value.

### HasSearchWeight

`func (o *MetaPropDefinition) HasSearchWeight() bool`

HasSearchWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


