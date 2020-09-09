# ConnectorXmlApiMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WithAuth** | Pointer to **bool** | Flag to disable authentication bypassing. If set to true it is expected a valid cookie/login is provided within the XML API request body. | [optional] 
**XmlRequest** | Pointer to **string** | The XML request body to proxy to the platform. | [optional] 

## Methods

### NewConnectorXmlApiMessage

`func NewConnectorXmlApiMessage() *ConnectorXmlApiMessage`

NewConnectorXmlApiMessage instantiates a new ConnectorXmlApiMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorXmlApiMessageWithDefaults

`func NewConnectorXmlApiMessageWithDefaults() *ConnectorXmlApiMessage`

NewConnectorXmlApiMessageWithDefaults instantiates a new ConnectorXmlApiMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWithAuth

`func (o *ConnectorXmlApiMessage) GetWithAuth() bool`

GetWithAuth returns the WithAuth field if non-nil, zero value otherwise.

### GetWithAuthOk

`func (o *ConnectorXmlApiMessage) GetWithAuthOk() (*bool, bool)`

GetWithAuthOk returns a tuple with the WithAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithAuth

`func (o *ConnectorXmlApiMessage) SetWithAuth(v bool)`

SetWithAuth sets WithAuth field to given value.

### HasWithAuth

`func (o *ConnectorXmlApiMessage) HasWithAuth() bool`

HasWithAuth returns a boolean if a field has been set.

### GetXmlRequest

`func (o *ConnectorXmlApiMessage) GetXmlRequest() string`

GetXmlRequest returns the XmlRequest field if non-nil, zero value otherwise.

### GetXmlRequestOk

`func (o *ConnectorXmlApiMessage) GetXmlRequestOk() (*string, bool)`

GetXmlRequestOk returns a tuple with the XmlRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXmlRequest

`func (o *ConnectorXmlApiMessage) SetXmlRequest(v string)`

SetXmlRequest sets XmlRequest field to given value.

### HasXmlRequest

`func (o *ConnectorXmlApiMessage) HasXmlRequest() bool`

HasXmlRequest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


