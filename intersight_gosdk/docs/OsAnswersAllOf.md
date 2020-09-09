# OsAnswersAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnswerFile** | Pointer to **string** | If the source of the answers is a static file, the content of the file is stored as value in this property. The value is mandatory only when the &#39;Source&#39; property has been set to &#39;File&#39;. | [optional] 
**Hostname** | Pointer to **string** | Hostname to be configured for the server in the OS. | [optional] 
**IpConfigType** | Pointer to **string** | IP configuration type. Values are Static or Dynamic configuration of IP. In case of static IP configuration, IP address, gateway and other details need to be populated. In case of dynamic the IP configuration is obtained dynamically from DHCP. * &#x60;static&#x60; - In case of static IP configuraton, provide the details such as IP address, netmask, and gateway. * &#x60;DHCP&#x60; - In case of dynamic IP configuration, the IP address, netmask and gateway detailsare obtained from DHCP. | [optional] [default to "static"]
**IpConfiguration** | Pointer to [**OsIpConfiguration**](os.IpConfiguration.md) |  | [optional] 
**IsAnswerFileSet** | Pointer to **bool** | Indicates whether the value of the &#39;answerFile&#39; property has been set. | [optional] [readonly] 
**IsRootPasswordCrypted** | Pointer to **bool** | Enable to indicate Root Password provided is encrypted. | [optional] 
**IsRootPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;rootPassword&#39; property has been set. | [optional] [readonly] 
**Nameserver** | Pointer to **string** | IP address of the name server to be configured in the OS. | [optional] 
**ProductKey** | Pointer to **string** | The product key to be used for a specific version of Windows installation. | [optional] 
**RootPassword** | Pointer to **string** | Password configured for the root / administrator user in the OS. You can enter a plain text or an encrypted password. Intersight encrypts the plaintext password. Enable the Encrypted Password option to provide an encrypted password. For more details on encrypting passwords, see Help Center. | [optional] 
**Source** | Pointer to **string** | Answer values can be provided from three sources - Embedded in OS image, static file, or as placeholder values for an answer file template. Source of the answers is given as value, Embedded/File/Template. &#39;Embedded&#39; option indicates that the answer file is embedded within the OS Image. &#39;File&#39; option indicates that the answers are provided as a file. &#39;Template&#39; indicates that the placeholders in the selected os.ConfigurationFile MO are replaced with values provided as os.Answers MO. * &#x60;None&#x60; - Indicates that answers is not sent and values must be populated from Install Template.   * &#x60;Embedded&#x60; - Indicates that the answer file is embedded within OS image. * &#x60;File&#x60; - Indicates that the answer file is a static content that has all thevalues populated. * &#x60;Template&#x60; - Indicates that the given answers are used to populate the answer filetemplate. The template allows the users to refer some server specificanswers as fields/placeholders and replace these placeholders with theactual values for each Server during OS installation using &#39;Answers&#39; and&#39;AdditionalParameters&#39; properties in os.Install MO.The answer file templates can be created by users as os.ConfigurationFile objects. | [optional] [default to "None"]

## Methods

### NewOsAnswersAllOf

`func NewOsAnswersAllOf() *OsAnswersAllOf`

NewOsAnswersAllOf instantiates a new OsAnswersAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsAnswersAllOfWithDefaults

`func NewOsAnswersAllOfWithDefaults() *OsAnswersAllOf`

NewOsAnswersAllOfWithDefaults instantiates a new OsAnswersAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswerFile

`func (o *OsAnswersAllOf) GetAnswerFile() string`

GetAnswerFile returns the AnswerFile field if non-nil, zero value otherwise.

### GetAnswerFileOk

`func (o *OsAnswersAllOf) GetAnswerFileOk() (*string, bool)`

GetAnswerFileOk returns a tuple with the AnswerFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerFile

`func (o *OsAnswersAllOf) SetAnswerFile(v string)`

SetAnswerFile sets AnswerFile field to given value.

### HasAnswerFile

`func (o *OsAnswersAllOf) HasAnswerFile() bool`

HasAnswerFile returns a boolean if a field has been set.

### GetHostname

`func (o *OsAnswersAllOf) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *OsAnswersAllOf) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *OsAnswersAllOf) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *OsAnswersAllOf) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetIpConfigType

`func (o *OsAnswersAllOf) GetIpConfigType() string`

GetIpConfigType returns the IpConfigType field if non-nil, zero value otherwise.

### GetIpConfigTypeOk

`func (o *OsAnswersAllOf) GetIpConfigTypeOk() (*string, bool)`

GetIpConfigTypeOk returns a tuple with the IpConfigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpConfigType

`func (o *OsAnswersAllOf) SetIpConfigType(v string)`

SetIpConfigType sets IpConfigType field to given value.

### HasIpConfigType

`func (o *OsAnswersAllOf) HasIpConfigType() bool`

HasIpConfigType returns a boolean if a field has been set.

### GetIpConfiguration

`func (o *OsAnswersAllOf) GetIpConfiguration() OsIpConfiguration`

GetIpConfiguration returns the IpConfiguration field if non-nil, zero value otherwise.

### GetIpConfigurationOk

`func (o *OsAnswersAllOf) GetIpConfigurationOk() (*OsIpConfiguration, bool)`

GetIpConfigurationOk returns a tuple with the IpConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpConfiguration

`func (o *OsAnswersAllOf) SetIpConfiguration(v OsIpConfiguration)`

SetIpConfiguration sets IpConfiguration field to given value.

### HasIpConfiguration

`func (o *OsAnswersAllOf) HasIpConfiguration() bool`

HasIpConfiguration returns a boolean if a field has been set.

### GetIsAnswerFileSet

`func (o *OsAnswersAllOf) GetIsAnswerFileSet() bool`

GetIsAnswerFileSet returns the IsAnswerFileSet field if non-nil, zero value otherwise.

### GetIsAnswerFileSetOk

`func (o *OsAnswersAllOf) GetIsAnswerFileSetOk() (*bool, bool)`

GetIsAnswerFileSetOk returns a tuple with the IsAnswerFileSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnswerFileSet

`func (o *OsAnswersAllOf) SetIsAnswerFileSet(v bool)`

SetIsAnswerFileSet sets IsAnswerFileSet field to given value.

### HasIsAnswerFileSet

`func (o *OsAnswersAllOf) HasIsAnswerFileSet() bool`

HasIsAnswerFileSet returns a boolean if a field has been set.

### GetIsRootPasswordCrypted

`func (o *OsAnswersAllOf) GetIsRootPasswordCrypted() bool`

GetIsRootPasswordCrypted returns the IsRootPasswordCrypted field if non-nil, zero value otherwise.

### GetIsRootPasswordCryptedOk

`func (o *OsAnswersAllOf) GetIsRootPasswordCryptedOk() (*bool, bool)`

GetIsRootPasswordCryptedOk returns a tuple with the IsRootPasswordCrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRootPasswordCrypted

`func (o *OsAnswersAllOf) SetIsRootPasswordCrypted(v bool)`

SetIsRootPasswordCrypted sets IsRootPasswordCrypted field to given value.

### HasIsRootPasswordCrypted

`func (o *OsAnswersAllOf) HasIsRootPasswordCrypted() bool`

HasIsRootPasswordCrypted returns a boolean if a field has been set.

### GetIsRootPasswordSet

`func (o *OsAnswersAllOf) GetIsRootPasswordSet() bool`

GetIsRootPasswordSet returns the IsRootPasswordSet field if non-nil, zero value otherwise.

### GetIsRootPasswordSetOk

`func (o *OsAnswersAllOf) GetIsRootPasswordSetOk() (*bool, bool)`

GetIsRootPasswordSetOk returns a tuple with the IsRootPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRootPasswordSet

`func (o *OsAnswersAllOf) SetIsRootPasswordSet(v bool)`

SetIsRootPasswordSet sets IsRootPasswordSet field to given value.

### HasIsRootPasswordSet

`func (o *OsAnswersAllOf) HasIsRootPasswordSet() bool`

HasIsRootPasswordSet returns a boolean if a field has been set.

### GetNameserver

`func (o *OsAnswersAllOf) GetNameserver() string`

GetNameserver returns the Nameserver field if non-nil, zero value otherwise.

### GetNameserverOk

`func (o *OsAnswersAllOf) GetNameserverOk() (*string, bool)`

GetNameserverOk returns a tuple with the Nameserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameserver

`func (o *OsAnswersAllOf) SetNameserver(v string)`

SetNameserver sets Nameserver field to given value.

### HasNameserver

`func (o *OsAnswersAllOf) HasNameserver() bool`

HasNameserver returns a boolean if a field has been set.

### GetProductKey

`func (o *OsAnswersAllOf) GetProductKey() string`

GetProductKey returns the ProductKey field if non-nil, zero value otherwise.

### GetProductKeyOk

`func (o *OsAnswersAllOf) GetProductKeyOk() (*string, bool)`

GetProductKeyOk returns a tuple with the ProductKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductKey

`func (o *OsAnswersAllOf) SetProductKey(v string)`

SetProductKey sets ProductKey field to given value.

### HasProductKey

`func (o *OsAnswersAllOf) HasProductKey() bool`

HasProductKey returns a boolean if a field has been set.

### GetRootPassword

`func (o *OsAnswersAllOf) GetRootPassword() string`

GetRootPassword returns the RootPassword field if non-nil, zero value otherwise.

### GetRootPasswordOk

`func (o *OsAnswersAllOf) GetRootPasswordOk() (*string, bool)`

GetRootPasswordOk returns a tuple with the RootPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPassword

`func (o *OsAnswersAllOf) SetRootPassword(v string)`

SetRootPassword sets RootPassword field to given value.

### HasRootPassword

`func (o *OsAnswersAllOf) HasRootPassword() bool`

HasRootPassword returns a boolean if a field has been set.

### GetSource

`func (o *OsAnswersAllOf) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *OsAnswersAllOf) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *OsAnswersAllOf) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *OsAnswersAllOf) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


