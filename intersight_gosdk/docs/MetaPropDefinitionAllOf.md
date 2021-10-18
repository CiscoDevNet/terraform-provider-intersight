# MetaPropDefinitionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "meta.PropDefinition"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "meta.PropDefinition"]
**ApiAccess** | Pointer to **string** | API access control for the property. Examples are NoAccess, ReadOnly, CreateOnly etc. * &#x60;NoAccess&#x60; - The property is not accessible from the API. * &#x60;ReadOnly&#x60; - The value of the property is read-only.An HTTP 4xx status code is returned when the user sends a POST/PUT/PATCH request that containsa ReadOnly property. * &#x60;CreateOnly&#x60; - The value of the property can be set when the REST resource is created. It cannot be changed after object creation.An HTTP 4xx status code is returned when the user sends a POST/PUT/PATCH request that containsa CreateOnly property.CreateOnly properties are returned in the response body of HTTP GET requests. * &#x60;ReadWrite&#x60; - The property has read/write access. * &#x60;WriteOnly&#x60; - The value of the property can be set but it is never returned in the response body of supported HTTP methods.This settings is used for sensitive properties such as passwords. * &#x60;ReadOnCreate&#x60; - The value of the property is returned in the HTTP POST response body when the REST resource is created.The property is not writeable and cannot be queried through a GET request after the resource has been created. | [optional] [readonly] [default to "NoAccess"]
**Default** | Pointer to **interface{}** | The default value of the property. Not applicable when IsComplexType is True. | [optional] 
**IsCollection** | Pointer to **bool** | Indicates whether the property is a collection (i.e. a JSON array), otherwise it is a single value. | [optional] [readonly] 
**IsComplexType** | Pointer to **bool** | Indicates whether the property is a complex type, otherwise it is a basic scalar type. | [optional] [readonly] 
**Kind** | Pointer to **string** | The kind of the property. * &#x60;Unknown&#x60; - The property kind is unknown. * &#x60;Boolean&#x60; - The &#39;Boolean&#39; kind of property. * &#x60;String&#x60; - The &#39;String&#39; kind of property. * &#x60;Integer&#x60; - The &#39;Integer&#39; kind of property. * &#x60;Int32&#x60; - The &#39;Int32&#39; kind of property. * &#x60;Int64&#x60; - The &#39;Int64&#39; kind of property. * &#x60;Float&#x60; - The &#39;Float&#39; kind of property. * &#x60;Double&#x60; - The &#39;Double&#39; kind of property. * &#x60;Date&#x60; - The &#39;Date&#39; kind of property. * &#x60;Duration&#x60; - The &#39;Duration&#39; kind of property. * &#x60;Time&#x60; - The &#39;Time&#39; kind of property. * &#x60;Json&#x60; - The &#39;JSON&#39; kind of property. * &#x60;Binary&#x60; - The &#39;Binary&#39; kind of property. * &#x60;EnumString&#x60; - The &#39;EnumString&#39; kind of property. * &#x60;EnumInteger&#x60; - The &#39;EnumInteger&#39; kind of property. * &#x60;ComplexType&#x60; - The &#39;ComplexType&#39; kind of property. | [optional] [readonly] [default to "Unknown"]
**Name** | Pointer to **string** | The name of the property. | [optional] [readonly] 
**OpSecurity** | Pointer to **string** | The data-at-rest security protection applied to this property when a Managed Object is persisted. For example, Encrypted or Cleartext. * &#x60;ClearText&#x60; - Data at rest is not encrypted using account specific keys.Note that data is always protected using volume encryption. ClearText propertiesare queryable and searchable. * &#x60;Encrypted&#x60; - The value of the property is encrypted with account-specific cryptographic keys.This setting is used for properties that contain sensitive data. Encrypted propertiesare not queryable and searchable. * &#x60;Pbkdf2&#x60; - The value of the property is hashed using the pbkdf2 key derivation functions thattakes a password, a salt, and a cost factor as inputs then generates a password hash.Its purpose is to make each password guessing trial by an attacker who has obtaineda password hash file expensive and therefore the cost of a guessing attack high or prohibitive. * &#x60;Bcrypt&#x60; - The value of the property is hashed using the bcrypt key derivation function. * &#x60;Sha512crypt&#x60; - The value of the property is hashed using the sha512crypt key derivation function. * &#x60;Argon2id&#x60; - The value of the property is hashed using the argon2id key derivation function. | [optional] [readonly] [default to "ClearText"]
**SearchWeight** | Pointer to **float32** | Enables the property to be searchable from global search. A value of 0 means this property is not globally searchable. | [optional] [readonly] 
**Type** | Pointer to **string** | The type of the property. In case of collection properties the type is that of individual elements. | [optional] [readonly] 

## Methods

### NewMetaPropDefinitionAllOf

`func NewMetaPropDefinitionAllOf(classId string, objectType string, ) *MetaPropDefinitionAllOf`

NewMetaPropDefinitionAllOf instantiates a new MetaPropDefinitionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaPropDefinitionAllOfWithDefaults

`func NewMetaPropDefinitionAllOfWithDefaults() *MetaPropDefinitionAllOf`

NewMetaPropDefinitionAllOfWithDefaults instantiates a new MetaPropDefinitionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MetaPropDefinitionAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MetaPropDefinitionAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MetaPropDefinitionAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MetaPropDefinitionAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MetaPropDefinitionAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MetaPropDefinitionAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetApiAccess

`func (o *MetaPropDefinitionAllOf) GetApiAccess() string`

GetApiAccess returns the ApiAccess field if non-nil, zero value otherwise.

### GetApiAccessOk

`func (o *MetaPropDefinitionAllOf) GetApiAccessOk() (*string, bool)`

GetApiAccessOk returns a tuple with the ApiAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiAccess

`func (o *MetaPropDefinitionAllOf) SetApiAccess(v string)`

SetApiAccess sets ApiAccess field to given value.

### HasApiAccess

`func (o *MetaPropDefinitionAllOf) HasApiAccess() bool`

HasApiAccess returns a boolean if a field has been set.

### GetDefault

`func (o *MetaPropDefinitionAllOf) GetDefault() interface{}`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *MetaPropDefinitionAllOf) GetDefaultOk() (*interface{}, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *MetaPropDefinitionAllOf) SetDefault(v interface{})`

SetDefault sets Default field to given value.

### HasDefault

`func (o *MetaPropDefinitionAllOf) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### SetDefaultNil

`func (o *MetaPropDefinitionAllOf) SetDefaultNil(b bool)`

 SetDefaultNil sets the value for Default to be an explicit nil

### UnsetDefault
`func (o *MetaPropDefinitionAllOf) UnsetDefault()`

UnsetDefault ensures that no value is present for Default, not even an explicit nil
### GetIsCollection

`func (o *MetaPropDefinitionAllOf) GetIsCollection() bool`

GetIsCollection returns the IsCollection field if non-nil, zero value otherwise.

### GetIsCollectionOk

`func (o *MetaPropDefinitionAllOf) GetIsCollectionOk() (*bool, bool)`

GetIsCollectionOk returns a tuple with the IsCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCollection

`func (o *MetaPropDefinitionAllOf) SetIsCollection(v bool)`

SetIsCollection sets IsCollection field to given value.

### HasIsCollection

`func (o *MetaPropDefinitionAllOf) HasIsCollection() bool`

HasIsCollection returns a boolean if a field has been set.

### GetIsComplexType

`func (o *MetaPropDefinitionAllOf) GetIsComplexType() bool`

GetIsComplexType returns the IsComplexType field if non-nil, zero value otherwise.

### GetIsComplexTypeOk

`func (o *MetaPropDefinitionAllOf) GetIsComplexTypeOk() (*bool, bool)`

GetIsComplexTypeOk returns a tuple with the IsComplexType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsComplexType

`func (o *MetaPropDefinitionAllOf) SetIsComplexType(v bool)`

SetIsComplexType sets IsComplexType field to given value.

### HasIsComplexType

`func (o *MetaPropDefinitionAllOf) HasIsComplexType() bool`

HasIsComplexType returns a boolean if a field has been set.

### GetKind

`func (o *MetaPropDefinitionAllOf) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *MetaPropDefinitionAllOf) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *MetaPropDefinitionAllOf) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *MetaPropDefinitionAllOf) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *MetaPropDefinitionAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetaPropDefinitionAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetaPropDefinitionAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MetaPropDefinitionAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOpSecurity

`func (o *MetaPropDefinitionAllOf) GetOpSecurity() string`

GetOpSecurity returns the OpSecurity field if non-nil, zero value otherwise.

### GetOpSecurityOk

`func (o *MetaPropDefinitionAllOf) GetOpSecurityOk() (*string, bool)`

GetOpSecurityOk returns a tuple with the OpSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpSecurity

`func (o *MetaPropDefinitionAllOf) SetOpSecurity(v string)`

SetOpSecurity sets OpSecurity field to given value.

### HasOpSecurity

`func (o *MetaPropDefinitionAllOf) HasOpSecurity() bool`

HasOpSecurity returns a boolean if a field has been set.

### GetSearchWeight

`func (o *MetaPropDefinitionAllOf) GetSearchWeight() float32`

GetSearchWeight returns the SearchWeight field if non-nil, zero value otherwise.

### GetSearchWeightOk

`func (o *MetaPropDefinitionAllOf) GetSearchWeightOk() (*float32, bool)`

GetSearchWeightOk returns a tuple with the SearchWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchWeight

`func (o *MetaPropDefinitionAllOf) SetSearchWeight(v float32)`

SetSearchWeight sets SearchWeight field to given value.

### HasSearchWeight

`func (o *MetaPropDefinitionAllOf) HasSearchWeight() bool`

HasSearchWeight returns a boolean if a field has been set.

### GetType

`func (o *MetaPropDefinitionAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetaPropDefinitionAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetaPropDefinitionAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MetaPropDefinitionAllOf) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


