# PartnerintegrationDocIssues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "partnerintegration.DocIssues"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "partnerintegration.DocIssues"]
**DocumentationIssues** | Pointer to **interface{}** | List of documentation issues. | [optional] [readonly] 
**Inventory** | Pointer to [**NullablePartnerintegrationInventoryRelationship**](PartnerintegrationInventoryRelationship.md) |  | [optional] 

## Methods

### NewPartnerintegrationDocIssues

`func NewPartnerintegrationDocIssues(classId string, objectType string, ) *PartnerintegrationDocIssues`

NewPartnerintegrationDocIssues instantiates a new PartnerintegrationDocIssues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerintegrationDocIssuesWithDefaults

`func NewPartnerintegrationDocIssuesWithDefaults() *PartnerintegrationDocIssues`

NewPartnerintegrationDocIssuesWithDefaults instantiates a new PartnerintegrationDocIssues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PartnerintegrationDocIssues) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PartnerintegrationDocIssues) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PartnerintegrationDocIssues) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PartnerintegrationDocIssues) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PartnerintegrationDocIssues) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PartnerintegrationDocIssues) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDocumentationIssues

`func (o *PartnerintegrationDocIssues) GetDocumentationIssues() interface{}`

GetDocumentationIssues returns the DocumentationIssues field if non-nil, zero value otherwise.

### GetDocumentationIssuesOk

`func (o *PartnerintegrationDocIssues) GetDocumentationIssuesOk() (*interface{}, bool)`

GetDocumentationIssuesOk returns a tuple with the DocumentationIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationIssues

`func (o *PartnerintegrationDocIssues) SetDocumentationIssues(v interface{})`

SetDocumentationIssues sets DocumentationIssues field to given value.

### HasDocumentationIssues

`func (o *PartnerintegrationDocIssues) HasDocumentationIssues() bool`

HasDocumentationIssues returns a boolean if a field has been set.

### SetDocumentationIssuesNil

`func (o *PartnerintegrationDocIssues) SetDocumentationIssuesNil(b bool)`

 SetDocumentationIssuesNil sets the value for DocumentationIssues to be an explicit nil

### UnsetDocumentationIssues
`func (o *PartnerintegrationDocIssues) UnsetDocumentationIssues()`

UnsetDocumentationIssues ensures that no value is present for DocumentationIssues, not even an explicit nil
### GetInventory

`func (o *PartnerintegrationDocIssues) GetInventory() PartnerintegrationInventoryRelationship`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *PartnerintegrationDocIssues) GetInventoryOk() (*PartnerintegrationInventoryRelationship, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *PartnerintegrationDocIssues) SetInventory(v PartnerintegrationInventoryRelationship)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *PartnerintegrationDocIssues) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### SetInventoryNil

`func (o *PartnerintegrationDocIssues) SetInventoryNil(b bool)`

 SetInventoryNil sets the value for Inventory to be an explicit nil

### UnsetInventory
`func (o *PartnerintegrationDocIssues) UnsetInventory()`

UnsetInventory ensures that no value is present for Inventory, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


