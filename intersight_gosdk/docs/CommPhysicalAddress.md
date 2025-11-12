# CommPhysicalAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "comm.PhysicalAddress"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "comm.PhysicalAddress"]
**Address1** | Pointer to **string** | The primary street address. | [optional] 
**Address2** | Pointer to **string** | Additional address information, such as suite number or floor. | [optional] 
**City** | Pointer to **string** | The city where the address is located. | [optional] 
**Country** | Pointer to **string** | The country code in ISO 3166-1 alpha-2 format. * &#x60;Unknown&#x60; - The value Unknown is used when the country code is not known or applicable. * &#x60;AD&#x60; - The country code for Andorra. * &#x60;AE&#x60; - The country code for United Arab Emirates. * &#x60;AF&#x60; - The country code for Afghanistan. * &#x60;AG&#x60; - The country code for Antigua and Barbuda. * &#x60;AI&#x60; - The country code for Anguilla. * &#x60;AL&#x60; - The country code for Albania. * &#x60;AM&#x60; - The country code for Armenia. * &#x60;AO&#x60; - The country code for Angola. * &#x60;AQ&#x60; - The country code for Antarctica. * &#x60;AR&#x60; - The country code for Argentina. * &#x60;AS&#x60; - The country code for American Samoa. * &#x60;AT&#x60; - The country code for Austria. * &#x60;AU&#x60; - The country code for Australia. * &#x60;AW&#x60; - The country code for Aruba. * &#x60;AX&#x60; - The country code for Åland Islands. * &#x60;AZ&#x60; - The country code for Azerbaijan. * &#x60;BA&#x60; - The country code for Bosnia and Herzegovina. * &#x60;BB&#x60; - The country code for Barbados. * &#x60;BD&#x60; - The country code for Bangladesh. * &#x60;BE&#x60; - The country code for Belgium. * &#x60;BF&#x60; - The country code for Burkina Faso. * &#x60;BG&#x60; - The country code for Bulgaria. * &#x60;BH&#x60; - The country code for Bahrain. * &#x60;BI&#x60; - The country code for Burundi. * &#x60;BJ&#x60; - The country code for Benin. * &#x60;BL&#x60; - The country code for Saint Barthélemy. * &#x60;BM&#x60; - The country code for Bermuda. * &#x60;BN&#x60; - The country code for Brunei Darussalam. * &#x60;BO&#x60; - The country code for Plurinational State of Bolivia. * &#x60;BQ&#x60; - The country code for Sint Eustatius and Saba Bonaire. * &#x60;BR&#x60; - The country code for Brazil. * &#x60;BS&#x60; - The country code for Bahamas. * &#x60;BT&#x60; - The country code for Bhutan. * &#x60;BV&#x60; - The country code for Bouvet Island. * &#x60;BW&#x60; - The country code for Botswana. * &#x60;BZ&#x60; - The country code for Belize. * &#x60;CA&#x60; - The country code for Canada. * &#x60;CC&#x60; - The country code for Cocos (Keeling) Islands. * &#x60;CD&#x60; - The country code for Democratic Republic of the Congo. * &#x60;CF&#x60; - The country code for Central African Republic. * &#x60;CG&#x60; - The country code for Congo. * &#x60;CH&#x60; - The country code for Switzerland. * &#x60;CI&#x60; - The country code for Côte d&#39;Ivoire. * &#x60;CK&#x60; - The country code for Cook Islands. * &#x60;CL&#x60; - The country code for Chile. * &#x60;CM&#x60; - The country code for Cameroon. * &#x60;CN&#x60; - The country code for China. * &#x60;CO&#x60; - The country code for Colombia. * &#x60;CR&#x60; - The country code for Costa Rica. * &#x60;CV&#x60; - The country code for Cabo Verde. * &#x60;CW&#x60; - The country code for Curaçao. * &#x60;CX&#x60; - The country code for Christmas Island. * &#x60;CY&#x60; - The country code for Cyprus. * &#x60;CZ&#x60; - The country code for Czechia. * &#x60;DE&#x60; - The country code for Germany. * &#x60;DJ&#x60; - The country code for Djibouti. * &#x60;DK&#x60; - The country code for Denmark. * &#x60;DM&#x60; - The country code for Dominica. * &#x60;DO&#x60; - The country code for Dominican Republic. * &#x60;DZ&#x60; - The country code for Algeria. * &#x60;EC&#x60; - The country code for Ecuador. * &#x60;EE&#x60; - The country code for Estonia. * &#x60;EG&#x60; - The country code for Egypt. * &#x60;EH&#x60; - The country code for Western Sahara. * &#x60;ER&#x60; - The country code for Eritrea. * &#x60;ES&#x60; - The country code for Spain. * &#x60;ET&#x60; - The country code for Ethiopia. * &#x60;FI&#x60; - The country code for Finland. * &#x60;FJ&#x60; - The country code for Fiji. * &#x60;FK&#x60; - The country code for Falkland Islands (Malvinas). * &#x60;FM&#x60; - The country code for Federated States of Micronesia. * &#x60;FO&#x60; - The country code for Faroe Islands. * &#x60;FR&#x60; - The country code for France. * &#x60;GA&#x60; - The country code for Gabon. * &#x60;GB&#x60; - The country code for United Kingdom of Great Britain and Northern Ireland. * &#x60;GD&#x60; - The country code for Grenada. * &#x60;GE&#x60; - The country code for Georgia. * &#x60;GF&#x60; - The country code for French Guiana. * &#x60;GG&#x60; - The country code for Guernsey. * &#x60;GH&#x60; - The country code for Ghana. * &#x60;GI&#x60; - The country code for Gibraltar. * &#x60;GL&#x60; - The country code for Greenland. * &#x60;GM&#x60; - The country code for Gambia. * &#x60;GN&#x60; - The country code for Guinea. * &#x60;GP&#x60; - The country code for Guadeloupe. * &#x60;GQ&#x60; - The country code for Equatorial Guinea. * &#x60;GR&#x60; - The country code for Greece. * &#x60;GS&#x60; - The country code for South Georgia and the South Sandwich Islands. * &#x60;GT&#x60; - The country code for Guatemala. * &#x60;GU&#x60; - The country code for Guam. * &#x60;GW&#x60; - The country code for Guinea-Bissau. * &#x60;GY&#x60; - The country code for Guyana. * &#x60;HK&#x60; - The country code for Hong Kong. * &#x60;HM&#x60; - The country code for Heard Island and McDonald Islands. * &#x60;HN&#x60; - The country code for Honduras. * &#x60;HR&#x60; - The country code for Croatia. * &#x60;HT&#x60; - The country code for Haiti. * &#x60;HU&#x60; - The country code for Hungary. * &#x60;ID&#x60; - The country code for Indonesia. * &#x60;IE&#x60; - The country code for Ireland. * &#x60;IL&#x60; - The country code for Israel. * &#x60;IM&#x60; - The country code for Isle of Man. * &#x60;IN&#x60; - The country code for India. * &#x60;IO&#x60; - The country code for British Indian Ocean Territory. * &#x60;IQ&#x60; - The country code for Iraq. * &#x60;IS&#x60; - The country code for Iceland. * &#x60;IT&#x60; - The country code for Italy. * &#x60;JE&#x60; - The country code for Jersey. * &#x60;JM&#x60; - The country code for Jamaica. * &#x60;JO&#x60; - The country code for Jordan. * &#x60;JP&#x60; - The country code for Japan. * &#x60;KE&#x60; - The country code for Kenya. * &#x60;KG&#x60; - The country code for Kyrgyzstan. * &#x60;KH&#x60; - The country code for Cambodia. * &#x60;KI&#x60; - The country code for Kiribati. * &#x60;KM&#x60; - The country code for Comoros. * &#x60;KN&#x60; - The country code for Saint Kitts and Nevis. * &#x60;KR&#x60; - The country code for Republic of Korea. * &#x60;KW&#x60; - The country code for Kuwait. * &#x60;KY&#x60; - The country code for Cayman Islands. * &#x60;KZ&#x60; - The country code for Kazakhstan. * &#x60;LA&#x60; - The country code for Lao People&#39;s Democratic Republic. * &#x60;LB&#x60; - The country code for Lebanon. * &#x60;LC&#x60; - The country code for Saint Lucia. * &#x60;LI&#x60; - The country code for Liechtenstein. * &#x60;LK&#x60; - The country code for Sri Lanka. * &#x60;LR&#x60; - The country code for Liberia. * &#x60;LS&#x60; - The country code for Lesotho. * &#x60;LT&#x60; - The country code for Lithuania. * &#x60;LU&#x60; - The country code for Luxembourg. * &#x60;LV&#x60; - The country code for Latvia. * &#x60;LY&#x60; - The country code for Libya. * &#x60;MA&#x60; - The country code for Morocco. * &#x60;MC&#x60; - The country code for Monaco. * &#x60;MD&#x60; - The country code for Republic of Moldova. * &#x60;ME&#x60; - The country code for Montenegro. * &#x60;MF&#x60; - The country code for Saint Martin (French part). * &#x60;MG&#x60; - The country code for Madagascar. * &#x60;MH&#x60; - The country code for Marshall Islands. * &#x60;MK&#x60; - The country code for North Macedonia. * &#x60;ML&#x60; - The country code for Mali. * &#x60;MM&#x60; - The country code for Myanmar. * &#x60;MN&#x60; - The country code for Mongolia. * &#x60;MO&#x60; - The country code for Macao. * &#x60;MP&#x60; - The country code for Northern Mariana Islands. * &#x60;MQ&#x60; - The country code for Martinique. * &#x60;MR&#x60; - The country code for Mauritania. * &#x60;MS&#x60; - The country code for Montserrat. * &#x60;MT&#x60; - The country code for Malta. * &#x60;MU&#x60; - The country code for Mauritius. * &#x60;MV&#x60; - The country code for Maldives. * &#x60;MW&#x60; - The country code for Malawi. * &#x60;MX&#x60; - The country code for Mexico. * &#x60;MY&#x60; - The country code for Malaysia. * &#x60;MZ&#x60; - The country code for Mozambique. * &#x60;NA&#x60; - The country code for Namibia. * &#x60;NC&#x60; - The country code for New Caledonia. * &#x60;NE&#x60; - The country code for Niger. * &#x60;NF&#x60; - The country code for Norfolk Island. * &#x60;NG&#x60; - The country code for Nigeria. * &#x60;NI&#x60; - The country code for Nicaragua. * &#x60;NL&#x60; - The country code for Kingdom of the Netherlands. * &#x60;NO&#x60; - The country code for Norway. * &#x60;NP&#x60; - The country code for Nepal. * &#x60;NR&#x60; - The country code for Nauru. * &#x60;NU&#x60; - The country code for Niue. * &#x60;NZ&#x60; - The country code for New Zealand. * &#x60;OM&#x60; - The country code for Oman. * &#x60;PA&#x60; - The country code for Panama. * &#x60;PE&#x60; - The country code for Peru. * &#x60;PF&#x60; - The country code for French Polynesia. * &#x60;PG&#x60; - The country code for Papua New Guinea. * &#x60;PH&#x60; - The country code for Philippines. * &#x60;PK&#x60; - The country code for Pakistan. * &#x60;PL&#x60; - The country code for Poland. * &#x60;PM&#x60; - The country code for Saint Pierre and Miquelon. * &#x60;PN&#x60; - The country code for Pitcairn. * &#x60;PR&#x60; - The country code for Puerto Rico. * &#x60;PS&#x60; - The country code for State of Palestine. * &#x60;PT&#x60; - The country code for Portugal. * &#x60;PW&#x60; - The country code for Palau. * &#x60;PY&#x60; - The country code for Paraguay. * &#x60;QA&#x60; - The country code for Qatar. * &#x60;RE&#x60; - The country code for Réunion. * &#x60;RO&#x60; - The country code for Romania. * &#x60;RS&#x60; - The country code for Serbia. * &#x60;RW&#x60; - The country code for Rwanda. * &#x60;SA&#x60; - The country code for Saudi Arabia. * &#x60;SB&#x60; - The country code for Solomon Islands. * &#x60;SC&#x60; - The country code for Seychelles. * &#x60;SD&#x60; - The country code for Sudan. * &#x60;SE&#x60; - The country code for Sweden. * &#x60;SG&#x60; - The country code for Singapore. * &#x60;SH&#x60; - The country code for Ascension and Tristan da Cunha Saint Helena. * &#x60;SI&#x60; - The country code for Slovenia. * &#x60;SJ&#x60; - The country code for Svalbard and Jan Mayen. * &#x60;SK&#x60; - The country code for Slovakia. * &#x60;SL&#x60; - The country code for Sierra Leone. * &#x60;SM&#x60; - The country code for San Marino. * &#x60;SN&#x60; - The country code for Senegal. * &#x60;SO&#x60; - The country code for Somalia. * &#x60;SR&#x60; - The country code for Suriname. * &#x60;SS&#x60; - The country code for South Sudan. * &#x60;ST&#x60; - The country code for Sao Tome and Principe. * &#x60;SV&#x60; - The country code for El Salvador. * &#x60;SX&#x60; - The country code for Sint Maarten (Dutch part). * &#x60;SZ&#x60; - The country code for Eswatini. * &#x60;TC&#x60; - The country code for Turks and Caicos Islands. * &#x60;TD&#x60; - The country code for Chad. * &#x60;TF&#x60; - The country code for French Southern Territories. * &#x60;TG&#x60; - The country code for Togo. * &#x60;TH&#x60; - The country code for Thailand. * &#x60;TJ&#x60; - The country code for Tajikistan. * &#x60;TK&#x60; - The country code for Tokelau. * &#x60;TL&#x60; - The country code for Timor-Leste. * &#x60;TM&#x60; - The country code for Turkmenistan. * &#x60;TN&#x60; - The country code for Tunisia. * &#x60;TO&#x60; - The country code for Tonga. * &#x60;TR&#x60; - The country code for Türkiye. * &#x60;TT&#x60; - The country code for Trinidad and Tobago. * &#x60;TV&#x60; - The country code for Tuvalu. * &#x60;TW&#x60; - The country code for Province of China Taiwan. * &#x60;TZ&#x60; - The country code for United Republic of Tanzania. * &#x60;UA&#x60; - The country code for Ukraine. * &#x60;UG&#x60; - The country code for Uganda. * &#x60;UM&#x60; - The country code for United States Minor Outlying Islands. * &#x60;US&#x60; - The country code for United States of America. * &#x60;UY&#x60; - The country code for Uruguay. * &#x60;UZ&#x60; - The country code for Uzbekistan. * &#x60;VA&#x60; - The country code for Holy See. * &#x60;VC&#x60; - The country code for Saint Vincent and the Grenadines. * &#x60;VE&#x60; - The country code for Bolivarian Republic of Venezuela. * &#x60;VG&#x60; - The country code for Virgin Islands (British). * &#x60;VI&#x60; - The country code for Virgin Islands (U.S.). * &#x60;VN&#x60; - The country code for Viet Nam. * &#x60;VU&#x60; - The country code for Vanuatu. * &#x60;WF&#x60; - The country code for Wallis and Futuna. * &#x60;WS&#x60; - The country code for Samoa. * &#x60;YE&#x60; - The country code for Yemen. * &#x60;YT&#x60; - The country code for Mayotte. * &#x60;ZA&#x60; - The country code for South Africa. * &#x60;ZM&#x60; - The country code for Zambia. * &#x60;ZW&#x60; - The country code for Zimbabwe. | [optional] [default to "Unknown"]
**PostalCode** | Pointer to **string** | The postal or ZIP code for the address. | [optional] 
**StateProvince** | Pointer to **string** | The state or province where the address is located. | [optional] 

## Methods

### NewCommPhysicalAddress

`func NewCommPhysicalAddress(classId string, objectType string, ) *CommPhysicalAddress`

NewCommPhysicalAddress instantiates a new CommPhysicalAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommPhysicalAddressWithDefaults

`func NewCommPhysicalAddressWithDefaults() *CommPhysicalAddress`

NewCommPhysicalAddressWithDefaults instantiates a new CommPhysicalAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CommPhysicalAddress) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CommPhysicalAddress) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CommPhysicalAddress) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CommPhysicalAddress) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CommPhysicalAddress) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CommPhysicalAddress) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAddress1

`func (o *CommPhysicalAddress) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *CommPhysicalAddress) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *CommPhysicalAddress) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.

### HasAddress1

`func (o *CommPhysicalAddress) HasAddress1() bool`

HasAddress1 returns a boolean if a field has been set.

### GetAddress2

`func (o *CommPhysicalAddress) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *CommPhysicalAddress) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *CommPhysicalAddress) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *CommPhysicalAddress) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### GetCity

`func (o *CommPhysicalAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CommPhysicalAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CommPhysicalAddress) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *CommPhysicalAddress) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountry

`func (o *CommPhysicalAddress) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CommPhysicalAddress) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CommPhysicalAddress) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CommPhysicalAddress) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetPostalCode

`func (o *CommPhysicalAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *CommPhysicalAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *CommPhysicalAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *CommPhysicalAddress) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetStateProvince

`func (o *CommPhysicalAddress) GetStateProvince() string`

GetStateProvince returns the StateProvince field if non-nil, zero value otherwise.

### GetStateProvinceOk

`func (o *CommPhysicalAddress) GetStateProvinceOk() (*string, bool)`

GetStateProvinceOk returns a tuple with the StateProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateProvince

`func (o *CommPhysicalAddress) SetStateProvince(v string)`

SetStateProvince sets StateProvince field to given value.

### HasStateProvince

`func (o *CommPhysicalAddress) HasStateProvince() bool`

HasStateProvince returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


