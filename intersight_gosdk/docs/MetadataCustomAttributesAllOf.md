# MetadataCustomAttributesAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "metadata.CustomAttributes"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "metadata.CustomAttributes"]
**AttributeName** | Pointer to **string** | The name of an attribute. If used as a key-value pair then this field represents the key. | [optional] 
**AttributeType** | Pointer to **string** | The data type for attributeValue. For e.g. string, int, float. | [optional] 
**AttributeValue** | Pointer to **string** | The attribute value. If used as a key-value pair then this field represents the value. | [optional] 

## Methods

### NewMetadataCustomAttributesAllOf

`func NewMetadataCustomAttributesAllOf(classId string, objectType string, ) *MetadataCustomAttributesAllOf`

NewMetadataCustomAttributesAllOf instantiates a new MetadataCustomAttributesAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataCustomAttributesAllOfWithDefaults

`func NewMetadataCustomAttributesAllOfWithDefaults() *MetadataCustomAttributesAllOf`

NewMetadataCustomAttributesAllOfWithDefaults instantiates a new MetadataCustomAttributesAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MetadataCustomAttributesAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MetadataCustomAttributesAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MetadataCustomAttributesAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MetadataCustomAttributesAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MetadataCustomAttributesAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MetadataCustomAttributesAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAttributeName

`func (o *MetadataCustomAttributesAllOf) GetAttributeName() string`

GetAttributeName returns the AttributeName field if non-nil, zero value otherwise.

### GetAttributeNameOk

`func (o *MetadataCustomAttributesAllOf) GetAttributeNameOk() (*string, bool)`

GetAttributeNameOk returns a tuple with the AttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeName

`func (o *MetadataCustomAttributesAllOf) SetAttributeName(v string)`

SetAttributeName sets AttributeName field to given value.

### HasAttributeName

`func (o *MetadataCustomAttributesAllOf) HasAttributeName() bool`

HasAttributeName returns a boolean if a field has been set.

### GetAttributeType

`func (o *MetadataCustomAttributesAllOf) GetAttributeType() string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *MetadataCustomAttributesAllOf) GetAttributeTypeOk() (*string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *MetadataCustomAttributesAllOf) SetAttributeType(v string)`

SetAttributeType sets AttributeType field to given value.

### HasAttributeType

`func (o *MetadataCustomAttributesAllOf) HasAttributeType() bool`

HasAttributeType returns a boolean if a field has been set.

### GetAttributeValue

`func (o *MetadataCustomAttributesAllOf) GetAttributeValue() string`

GetAttributeValue returns the AttributeValue field if non-nil, zero value otherwise.

### GetAttributeValueOk

`func (o *MetadataCustomAttributesAllOf) GetAttributeValueOk() (*string, bool)`

GetAttributeValueOk returns a tuple with the AttributeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeValue

`func (o *MetadataCustomAttributesAllOf) SetAttributeValue(v string)`

SetAttributeValue sets AttributeValue field to given value.

### HasAttributeValue

`func (o *MetadataCustomAttributesAllOf) HasAttributeValue() bool`

HasAttributeValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


