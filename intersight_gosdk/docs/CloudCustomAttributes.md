# CloudCustomAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cloud.CustomAttributes"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cloud.CustomAttributes"]
**AttributeName** | Pointer to **string** | The name of an attribute. If used as a key-value pair then this field represents the key. | [optional] 
**AttributeType** | Pointer to **string** | The data type for attributeValue. For e.g. string, int, float. | [optional] 
**AttributeValue** | Pointer to **string** | The attribute value. If used as a key-value pair then this field represents the value. | [optional] 

## Methods

### NewCloudCustomAttributes

`func NewCloudCustomAttributes(classId string, objectType string, ) *CloudCustomAttributes`

NewCloudCustomAttributes instantiates a new CloudCustomAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCustomAttributesWithDefaults

`func NewCloudCustomAttributesWithDefaults() *CloudCustomAttributes`

NewCloudCustomAttributesWithDefaults instantiates a new CloudCustomAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudCustomAttributes) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudCustomAttributes) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudCustomAttributes) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudCustomAttributes) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudCustomAttributes) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudCustomAttributes) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAttributeName

`func (o *CloudCustomAttributes) GetAttributeName() string`

GetAttributeName returns the AttributeName field if non-nil, zero value otherwise.

### GetAttributeNameOk

`func (o *CloudCustomAttributes) GetAttributeNameOk() (*string, bool)`

GetAttributeNameOk returns a tuple with the AttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeName

`func (o *CloudCustomAttributes) SetAttributeName(v string)`

SetAttributeName sets AttributeName field to given value.

### HasAttributeName

`func (o *CloudCustomAttributes) HasAttributeName() bool`

HasAttributeName returns a boolean if a field has been set.

### GetAttributeType

`func (o *CloudCustomAttributes) GetAttributeType() string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *CloudCustomAttributes) GetAttributeTypeOk() (*string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *CloudCustomAttributes) SetAttributeType(v string)`

SetAttributeType sets AttributeType field to given value.

### HasAttributeType

`func (o *CloudCustomAttributes) HasAttributeType() bool`

HasAttributeType returns a boolean if a field has been set.

### GetAttributeValue

`func (o *CloudCustomAttributes) GetAttributeValue() string`

GetAttributeValue returns the AttributeValue field if non-nil, zero value otherwise.

### GetAttributeValueOk

`func (o *CloudCustomAttributes) GetAttributeValueOk() (*string, bool)`

GetAttributeValueOk returns a tuple with the AttributeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeValue

`func (o *CloudCustomAttributes) SetAttributeValue(v string)`

SetAttributeValue sets AttributeValue field to given value.

### HasAttributeValue

`func (o *CloudCustomAttributes) HasAttributeValue() bool`

HasAttributeValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


