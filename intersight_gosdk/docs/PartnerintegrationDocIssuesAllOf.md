# PartnerintegrationDocIssuesAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "partnerintegration.DocIssues"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "partnerintegration.DocIssues"]
**DocumentationIssues** | Pointer to **interface{}** | List of documentation issues. | [optional] [readonly] 
**Inventory** | Pointer to [**PartnerintegrationInventoryRelationship**](PartnerintegrationInventoryRelationship.md) |  | [optional] 

## Methods

### NewPartnerintegrationDocIssuesAllOf

`func NewPartnerintegrationDocIssuesAllOf(classId string, objectType string, ) *PartnerintegrationDocIssuesAllOf`

NewPartnerintegrationDocIssuesAllOf instantiates a new PartnerintegrationDocIssuesAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerintegrationDocIssuesAllOfWithDefaults

`func NewPartnerintegrationDocIssuesAllOfWithDefaults() *PartnerintegrationDocIssuesAllOf`

NewPartnerintegrationDocIssuesAllOfWithDefaults instantiates a new PartnerintegrationDocIssuesAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PartnerintegrationDocIssuesAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PartnerintegrationDocIssuesAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PartnerintegrationDocIssuesAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PartnerintegrationDocIssuesAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PartnerintegrationDocIssuesAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PartnerintegrationDocIssuesAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDocumentationIssues

`func (o *PartnerintegrationDocIssuesAllOf) GetDocumentationIssues() interface{}`

GetDocumentationIssues returns the DocumentationIssues field if non-nil, zero value otherwise.

### GetDocumentationIssuesOk

`func (o *PartnerintegrationDocIssuesAllOf) GetDocumentationIssuesOk() (*interface{}, bool)`

GetDocumentationIssuesOk returns a tuple with the DocumentationIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationIssues

`func (o *PartnerintegrationDocIssuesAllOf) SetDocumentationIssues(v interface{})`

SetDocumentationIssues sets DocumentationIssues field to given value.

### HasDocumentationIssues

`func (o *PartnerintegrationDocIssuesAllOf) HasDocumentationIssues() bool`

HasDocumentationIssues returns a boolean if a field has been set.

### SetDocumentationIssuesNil

`func (o *PartnerintegrationDocIssuesAllOf) SetDocumentationIssuesNil(b bool)`

 SetDocumentationIssuesNil sets the value for DocumentationIssues to be an explicit nil

### UnsetDocumentationIssues
`func (o *PartnerintegrationDocIssuesAllOf) UnsetDocumentationIssues()`

UnsetDocumentationIssues ensures that no value is present for DocumentationIssues, not even an explicit nil
### GetInventory

`func (o *PartnerintegrationDocIssuesAllOf) GetInventory() PartnerintegrationInventoryRelationship`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *PartnerintegrationDocIssuesAllOf) GetInventoryOk() (*PartnerintegrationInventoryRelationship, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *PartnerintegrationDocIssuesAllOf) SetInventory(v PartnerintegrationInventoryRelationship)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *PartnerintegrationDocIssuesAllOf) HasInventory() bool`

HasInventory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


