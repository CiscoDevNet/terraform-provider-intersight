# OnpremClusterInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "onprem.ClusterInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "onprem.ClusterInfo"]
**Nodes** | Pointer to [**[]OnpremClusterNode**](OnpremClusterNode.md) |  | [optional] 
**QuorumOk** | Pointer to **bool** | Indicates if the cluster quorum requirement is met. This field has a value of &#39;true&#39; when at least (n + 1) / 2 number of nodes are up and running. | [optional] [readonly] 
**ReportedTime** | Pointer to **time.Time** | Date on which the cluster information was collected and reported. | [optional] [readonly] 
**SetupTime** | Pointer to **time.Time** | Initial setup date and time of the Intersight Appliance. | [optional] [readonly] 
**TimeZone** | Pointer to **string** | Timezone of the Intersight Appliance. The default value is in UTC. * &#x60;Pacific/Niue&#x60; -  * &#x60;Africa/Abidjan&#x60; -  * &#x60;Africa/Accra&#x60; -  * &#x60;Africa/Addis_Ababa&#x60; -  * &#x60;Africa/Algiers&#x60; -  * &#x60;Africa/Asmara&#x60; -  * &#x60;Africa/Bamako&#x60; -  * &#x60;Africa/Bangui&#x60; -  * &#x60;Africa/Banjul&#x60; -  * &#x60;Africa/Bissau&#x60; -  * &#x60;Africa/Blantyre&#x60; -  * &#x60;Africa/Brazzaville&#x60; -  * &#x60;Africa/Bujumbura&#x60; -  * &#x60;Africa/Cairo&#x60; -  * &#x60;Africa/Casablanca&#x60; -  * &#x60;Africa/Ceuta&#x60; -  * &#x60;Africa/Conakry&#x60; -  * &#x60;Africa/Dakar&#x60; -  * &#x60;Africa/Dar_es_Salaam&#x60; -  * &#x60;Africa/Djibouti&#x60; -  * &#x60;Africa/Douala&#x60; -  * &#x60;Africa/El_Aaiun&#x60; -  * &#x60;Africa/Freetown&#x60; -  * &#x60;Africa/Gaborone&#x60; -  * &#x60;Africa/Harare&#x60; -  * &#x60;Africa/Johannesburg&#x60; -  * &#x60;Africa/Juba&#x60; -  * &#x60;Africa/Kampala&#x60; -  * &#x60;Africa/Khartoum&#x60; -  * &#x60;Africa/Kigali&#x60; -  * &#x60;Africa/Kinshasa&#x60; -  * &#x60;Africa/Lagos&#x60; -  * &#x60;Africa/Libreville&#x60; -  * &#x60;Africa/Lome&#x60; -  * &#x60;Africa/Luanda&#x60; -  * &#x60;Africa/Lubumbashi&#x60; -  * &#x60;Africa/Lusaka&#x60; -  * &#x60;Africa/Malabo&#x60; -  * &#x60;Africa/Maputo&#x60; -  * &#x60;Africa/Maseru&#x60; -  * &#x60;Africa/Mbabane&#x60; -  * &#x60;Africa/Mogadishu&#x60; -  * &#x60;Africa/Monrovia&#x60; -  * &#x60;Africa/Nairobi&#x60; -  * &#x60;Africa/Ndjamena&#x60; -  * &#x60;Africa/Niamey&#x60; -  * &#x60;Africa/Nouakchott&#x60; -  * &#x60;Africa/Ouagadougou&#x60; -  * &#x60;Africa/Porto-Novo&#x60; -  * &#x60;Africa/Sao_Tome&#x60; -  * &#x60;Africa/Tripoli&#x60; -  * &#x60;Africa/Tunis&#x60; -  * &#x60;Africa/Windhoek&#x60; -  * &#x60;America/Adak&#x60; -  * &#x60;America/Anchorage&#x60; -  * &#x60;America/Anguilla&#x60; -  * &#x60;America/Antigua&#x60; -  * &#x60;America/Araguaina&#x60; -  * &#x60;America/Argentina/Buenos_Aires&#x60; -  * &#x60;America/Argentina/Catamarca&#x60; -  * &#x60;America/Argentina/Cordoba&#x60; -  * &#x60;America/Argentina/Jujuy&#x60; -  * &#x60;America/Argentina/La_Rioja&#x60; -  * &#x60;America/Argentina/Mendoza&#x60; -  * &#x60;America/Argentina/Rio_Gallegos&#x60; -  * &#x60;America/Argentina/Salta&#x60; -  * &#x60;America/Argentina/San_Juan&#x60; -  * &#x60;America/Argentina/San_Luis&#x60; -  * &#x60;America/Argentina/Tucuman&#x60; -  * &#x60;America/Argentina/Ushuaia&#x60; -  * &#x60;America/Aruba&#x60; -  * &#x60;America/Asuncion&#x60; -  * &#x60;America/Atikokan&#x60; -  * &#x60;America/Bahia&#x60; -  * &#x60;America/Bahia_Banderas&#x60; -  * &#x60;America/Barbados&#x60; -  * &#x60;America/Belem&#x60; -  * &#x60;America/Belize&#x60; -  * &#x60;America/Blanc-Sablon&#x60; -  * &#x60;America/Boa_Vista&#x60; -  * &#x60;America/Bogota&#x60; -  * &#x60;America/Boise&#x60; -  * &#x60;America/Cambridge_Bay&#x60; -  * &#x60;America/Campo_Grande&#x60; -  * &#x60;America/Cancun&#x60; -  * &#x60;America/Caracas&#x60; -  * &#x60;America/Cayenne&#x60; -  * &#x60;America/Cayman&#x60; -  * &#x60;America/Chicago&#x60; -  * &#x60;America/Chihuahua&#x60; -  * &#x60;America/Costa_Rica&#x60; -  * &#x60;America/Creston&#x60; -  * &#x60;America/Cuiaba&#x60; -  * &#x60;America/Curacao&#x60; -  * &#x60;America/Danmarkshavn&#x60; -  * &#x60;America/Dawson&#x60; -  * &#x60;America/Dawson_Creek&#x60; -  * &#x60;America/Denver&#x60; -  * &#x60;America/Detroit&#x60; -  * &#x60;America/Dominica&#x60; -  * &#x60;America/Edmonton&#x60; -  * &#x60;America/Eirunepe&#x60; -  * &#x60;America/El_Salvador&#x60; -  * &#x60;America/Fortaleza&#x60; -  * &#x60;America/Glace_Bay&#x60; -  * &#x60;America/Godthab&#x60; -  * &#x60;America/Goose_Bay&#x60; -  * &#x60;America/Grand_Turk&#x60; -  * &#x60;America/Grenada&#x60; -  * &#x60;America/Guadeloupe&#x60; -  * &#x60;America/Guatemala&#x60; -  * &#x60;America/Guayaquil&#x60; -  * &#x60;America/Guyana&#x60; -  * &#x60;America/Halifax&#x60; -  * &#x60;America/Havana&#x60; -  * &#x60;America/Hermosillo&#x60; -  * &#x60;America/Indiana/Indianapolis&#x60; -  * &#x60;America/Indiana/Knox&#x60; -  * &#x60;America/Indiana/Marengo&#x60; -  * &#x60;America/Indiana/Petersburg&#x60; -  * &#x60;America/Indiana/Tell_City&#x60; -  * &#x60;America/Indiana/Vevay&#x60; -  * &#x60;America/Indiana/Vincennes&#x60; -  * &#x60;America/Indiana/Winamac&#x60; -  * &#x60;America/Inuvik&#x60; -  * &#x60;America/Iqaluit&#x60; -  * &#x60;America/Jamaica&#x60; -  * &#x60;America/Juneau&#x60; -  * &#x60;America/Kentucky/Louisville&#x60; -  * &#x60;America/Kentucky/Monticello&#x60; -  * &#x60;America/Kralendijk&#x60; -  * &#x60;America/La_Paz&#x60; -  * &#x60;America/Lima&#x60; -  * &#x60;America/Los_Angeles&#x60; -  * &#x60;America/Lower_Princes&#x60; -  * &#x60;America/Maceio&#x60; -  * &#x60;America/Managua&#x60; -  * &#x60;America/Manaus&#x60; -  * &#x60;America/Marigot&#x60; -  * &#x60;America/Martinique&#x60; -  * &#x60;America/Matamoros&#x60; -  * &#x60;America/Mazatlan&#x60; -  * &#x60;America/Menominee&#x60; -  * &#x60;America/Merida&#x60; -  * &#x60;America/Metlakatla&#x60; -  * &#x60;America/Mexico_City&#x60; -  * &#x60;America/Miquelon&#x60; -  * &#x60;America/Moncton&#x60; -  * &#x60;America/Monterrey&#x60; -  * &#x60;America/Montevideo&#x60; -  * &#x60;America/Montreal&#x60; -  * &#x60;America/Montserrat&#x60; -  * &#x60;America/Nassau&#x60; -  * &#x60;America/New_York&#x60; -  * &#x60;America/Nipigon&#x60; -  * &#x60;America/Nome&#x60; -  * &#x60;America/Noronha&#x60; -  * &#x60;America/North_Dakota/Beulah&#x60; -  * &#x60;America/North_Dakota/Center&#x60; -  * &#x60;America/North_Dakota/New_Salem&#x60; -  * &#x60;America/Ojinaga&#x60; -  * &#x60;America/Panama&#x60; -  * &#x60;America/Pangnirtung&#x60; -  * &#x60;America/Paramaribo&#x60; -  * &#x60;America/Phoenix&#x60; -  * &#x60;America/Port-au-Prince&#x60; -  * &#x60;America/Port_of_Spain&#x60; -  * &#x60;America/Porto_Velho&#x60; -  * &#x60;America/Puerto_Rico&#x60; -  * &#x60;America/Rainy_River&#x60; -  * &#x60;America/Rankin_Inlet&#x60; -  * &#x60;America/Recife&#x60; -  * &#x60;America/Regina&#x60; -  * &#x60;America/Resolute&#x60; -  * &#x60;America/Rio_Branco&#x60; -  * &#x60;America/Santa_Isabel&#x60; -  * &#x60;America/Santarem&#x60; -  * &#x60;America/Santiago&#x60; -  * &#x60;America/Santo_Domingo&#x60; -  * &#x60;America/Sao_Paulo&#x60; -  * &#x60;America/Scoresbysund&#x60; -  * &#x60;America/Shiprock&#x60; -  * &#x60;America/Sitka&#x60; -  * &#x60;America/St_Barthelemy&#x60; -  * &#x60;America/St_Johns&#x60; -  * &#x60;America/St_Kitts&#x60; -  * &#x60;America/St_Lucia&#x60; -  * &#x60;America/St_Thomas&#x60; -  * &#x60;America/St_Vincent&#x60; -  * &#x60;America/Swift_Current&#x60; -  * &#x60;America/Tegucigalpa&#x60; -  * &#x60;America/Thule&#x60; -  * &#x60;America/Thunder_Bay&#x60; -  * &#x60;America/Tijuana&#x60; -  * &#x60;America/Toronto&#x60; -  * &#x60;America/Tortola&#x60; -  * &#x60;America/Vancouver&#x60; -  * &#x60;America/Whitehorse&#x60; -  * &#x60;America/Winnipeg&#x60; -  * &#x60;America/Yakutat&#x60; -  * &#x60;America/Yellowknife&#x60; -  * &#x60;Antarctica/Casey&#x60; -  * &#x60;Antarctica/Davis&#x60; -  * &#x60;Antarctica/DumontDUrville&#x60; -  * &#x60;Antarctica/Macquarie&#x60; -  * &#x60;Antarctica/Mawson&#x60; -  * &#x60;Antarctica/McMurdo&#x60; -  * &#x60;Antarctica/Palmer&#x60; -  * &#x60;Antarctica/Rothera&#x60; -  * &#x60;Antarctica/South_Pole&#x60; -  * &#x60;Antarctica/Syowa&#x60; -  * &#x60;Antarctica/Troll&#x60; -  * &#x60;Antarctica/Vostok&#x60; -  * &#x60;Arctic/Longyearbyen&#x60; -  * &#x60;Asia/Aden&#x60; -  * &#x60;Asia/Almaty&#x60; -  * &#x60;Asia/Amman&#x60; -  * &#x60;Asia/Anadyr&#x60; -  * &#x60;Asia/Aqtau&#x60; -  * &#x60;Asia/Aqtobe&#x60; -  * &#x60;Asia/Ashgabat&#x60; -  * &#x60;Asia/Baghdad&#x60; -  * &#x60;Asia/Bahrain&#x60; -  * &#x60;Asia/Baku&#x60; -  * &#x60;Asia/Bangkok&#x60; -  * &#x60;Asia/Beirut&#x60; -  * &#x60;Asia/Bishkek&#x60; -  * &#x60;Asia/Brunei&#x60; -  * &#x60;Asia/Calcutta&#x60; -  * &#x60;Asia/Choibalsan&#x60; -  * &#x60;Asia/Chongqing&#x60; -  * &#x60;Asia/Colombo&#x60; -  * &#x60;Asia/Damascus&#x60; -  * &#x60;Asia/Dhaka&#x60; -  * &#x60;Asia/Dili&#x60; -  * &#x60;Asia/Dubai&#x60; -  * &#x60;Asia/Dushanbe&#x60; -  * &#x60;Asia/Gaza&#x60; -  * &#x60;Asia/Harbin&#x60; -  * &#x60;Asia/Hebron&#x60; -  * &#x60;Asia/Ho_Chi_Minh&#x60; -  * &#x60;Asia/Hong_Kong&#x60; -  * &#x60;Asia/Hovd&#x60; -  * &#x60;Asia/Irkutsk&#x60; -  * &#x60;Asia/Jakarta&#x60; -  * &#x60;Asia/Jayapura&#x60; -  * &#x60;Asia/Jerusalem&#x60; -  * &#x60;Asia/Kabul&#x60; -  * &#x60;Asia/Kamchatka&#x60; -  * &#x60;Asia/Karachi&#x60; -  * &#x60;Asia/Kashgar&#x60; -  * &#x60;Asia/Kathmandu&#x60; -  * &#x60;Asia/Katmandu&#x60; -  * &#x60;Asia/Khandyga&#x60; -  * &#x60;Asia/Kolkata&#x60; -  * &#x60;Asia/Krasnoyarsk&#x60; -  * &#x60;Asia/Kuala_Lumpur&#x60; -  * &#x60;Asia/Kuching&#x60; -  * &#x60;Asia/Kuwait&#x60; -  * &#x60;Asia/Macau&#x60; -  * &#x60;Asia/Magadan&#x60; -  * &#x60;Asia/Makassar&#x60; -  * &#x60;Asia/Manila&#x60; -  * &#x60;Asia/Muscat&#x60; -  * &#x60;Asia/Nicosia&#x60; -  * &#x60;Asia/Novokuznetsk&#x60; -  * &#x60;Asia/Novosibirsk&#x60; -  * &#x60;Asia/Omsk&#x60; -  * &#x60;Asia/Oral&#x60; -  * &#x60;Asia/Phnom_Penh&#x60; -  * &#x60;Asia/Pontianak&#x60; -  * &#x60;Asia/Pyongyang&#x60; -  * &#x60;Asia/Qatar&#x60; -  * &#x60;Asia/Qyzylorda&#x60; -  * &#x60;Asia/Rangoon&#x60; -  * &#x60;Asia/Riyadh&#x60; -  * &#x60;Asia/Saigon&#x60; -  * &#x60;Asia/Sakhalin&#x60; -  * &#x60;Asia/Samarkand&#x60; -  * &#x60;Asia/Seoul&#x60; -  * &#x60;Asia/Shanghai&#x60; -  * &#x60;Asia/Singapore&#x60; -  * &#x60;Asia/Taipei&#x60; -  * &#x60;Asia/Tashkent&#x60; -  * &#x60;Asia/Tbilisi&#x60; -  * &#x60;Asia/Tehran&#x60; -  * &#x60;Asia/Thimphu&#x60; -  * &#x60;Asia/Tokyo&#x60; -  * &#x60;Asia/Ulaanbaatar&#x60; -  * &#x60;Asia/Urumqi&#x60; -  * &#x60;Asia/Ust-Nera&#x60; -  * &#x60;Asia/Vientiane&#x60; -  * &#x60;Asia/Vladivostok&#x60; -  * &#x60;Asia/Yakutsk&#x60; -  * &#x60;Asia/Yekaterinburg&#x60; -  * &#x60;Asia/Yerevan&#x60; -  * &#x60;Atlantic/Azores&#x60; -  * &#x60;Atlantic/Bermuda&#x60; -  * &#x60;Atlantic/Canary&#x60; -  * &#x60;Atlantic/Cape_Verde&#x60; -  * &#x60;Atlantic/Faroe&#x60; -  * &#x60;Atlantic/Madeira&#x60; -  * &#x60;Atlantic/Reykjavik&#x60; -  * &#x60;Atlantic/South_Georgia&#x60; -  * &#x60;Atlantic/St_Helena&#x60; -  * &#x60;Atlantic/Stanley&#x60; -  * &#x60;Australia/Adelaide&#x60; -  * &#x60;Australia/Brisbane&#x60; -  * &#x60;Australia/Broken_Hill&#x60; -  * &#x60;Australia/Currie&#x60; -  * &#x60;Australia/Darwin&#x60; -  * &#x60;Australia/Eucla&#x60; -  * &#x60;Australia/Hobart&#x60; -  * &#x60;Australia/Lindeman&#x60; -  * &#x60;Australia/Lord_Howe&#x60; -  * &#x60;Australia/Melbourne&#x60; -  * &#x60;Australia/Perth&#x60; -  * &#x60;Australia/Sydney&#x60; -  * &#x60;Etc/GMT&#x60; -  * &#x60;Europe/Amsterdam&#x60; -  * &#x60;Europe/Andorra&#x60; -  * &#x60;Europe/Athens&#x60; -  * &#x60;Europe/Belgrade&#x60; -  * &#x60;Europe/Berlin&#x60; -  * &#x60;Europe/Bratislava&#x60; -  * &#x60;Europe/Brussels&#x60; -  * &#x60;Europe/Bucharest&#x60; -  * &#x60;Europe/Budapest&#x60; -  * &#x60;Europe/Busingen&#x60; -  * &#x60;Europe/Chisinau&#x60; -  * &#x60;Europe/Copenhagen&#x60; -  * &#x60;Europe/Dublin&#x60; -  * &#x60;Europe/Gibraltar&#x60; -  * &#x60;Europe/Guernsey&#x60; -  * &#x60;Europe/Helsinki&#x60; -  * &#x60;Europe/Isle_of_Man&#x60; -  * &#x60;Europe/Istanbul&#x60; -  * &#x60;Europe/Jersey&#x60; -  * &#x60;Europe/Kaliningrad&#x60; -  * &#x60;Europe/Kiev&#x60; -  * &#x60;Europe/Lisbon&#x60; -  * &#x60;Europe/Ljubljana&#x60; -  * &#x60;Europe/London&#x60; -  * &#x60;Europe/Luxembourg&#x60; -  * &#x60;Europe/Madrid&#x60; -  * &#x60;Europe/Malta&#x60; -  * &#x60;Europe/Mariehamn&#x60; -  * &#x60;Europe/Minsk&#x60; -  * &#x60;Europe/Monaco&#x60; -  * &#x60;Europe/Moscow&#x60; -  * &#x60;Europe/Oslo&#x60; -  * &#x60;Europe/Paris&#x60; -  * &#x60;Europe/Podgorica&#x60; -  * &#x60;Europe/Prague&#x60; -  * &#x60;Europe/Riga&#x60; -  * &#x60;Europe/Rome&#x60; -  * &#x60;Europe/Samara&#x60; -  * &#x60;Europe/San_Marino&#x60; -  * &#x60;Europe/Sarajevo&#x60; -  * &#x60;Europe/Simferopol&#x60; -  * &#x60;Europe/Skopje&#x60; -  * &#x60;Europe/Sofia&#x60; -  * &#x60;Europe/Stockholm&#x60; -  * &#x60;Europe/Tallinn&#x60; -  * &#x60;Europe/Tirane&#x60; -  * &#x60;Europe/Uzhgorod&#x60; -  * &#x60;Europe/Vaduz&#x60; -  * &#x60;Europe/Vatican&#x60; -  * &#x60;Europe/Vienna&#x60; -  * &#x60;Europe/Vilnius&#x60; -  * &#x60;Europe/Volgograd&#x60; -  * &#x60;Europe/Warsaw&#x60; -  * &#x60;Europe/Zagreb&#x60; -  * &#x60;Europe/Zaporozhye&#x60; -  * &#x60;Europe/Zurich&#x60; -  * &#x60;Indian/Antananarivo&#x60; -  * &#x60;Indian/Chagos&#x60; -  * &#x60;Indian/Christmas&#x60; -  * &#x60;Indian/Cocos&#x60; -  * &#x60;Indian/Comoro&#x60; -  * &#x60;Indian/Kerguelen&#x60; -  * &#x60;Indian/Mahe&#x60; -  * &#x60;Indian/Maldives&#x60; -  * &#x60;Indian/Mauritius&#x60; -  * &#x60;Indian/Mayotte&#x60; -  * &#x60;Indian/Reunion&#x60; -  * &#x60;Pacific/Apia&#x60; -  * &#x60;Pacific/Auckland&#x60; -  * &#x60;Pacific/Chatham&#x60; -  * &#x60;Pacific/Chuuk&#x60; -  * &#x60;Pacific/Easter&#x60; -  * &#x60;Pacific/Efate&#x60; -  * &#x60;Pacific/Enderbury&#x60; -  * &#x60;Pacific/Fakaofo&#x60; -  * &#x60;Pacific/Fiji&#x60; -  * &#x60;Pacific/Funafuti&#x60; -  * &#x60;Pacific/Galapagos&#x60; -  * &#x60;Pacific/Gambier&#x60; -  * &#x60;Pacific/Guadalcanal&#x60; -  * &#x60;Pacific/Guam&#x60; -  * &#x60;Pacific/Honolulu&#x60; -  * &#x60;Pacific/Johnston&#x60; -  * &#x60;Pacific/Kiritimati&#x60; -  * &#x60;Pacific/Kosrae&#x60; -  * &#x60;Pacific/Kwajalein&#x60; -  * &#x60;Pacific/Majuro&#x60; -  * &#x60;Pacific/Marquesas&#x60; -  * &#x60;Pacific/Midway&#x60; -  * &#x60;Pacific/Nauru&#x60; -  * &#x60;Pacific/Norfolk&#x60; -  * &#x60;Pacific/Noumea&#x60; -  * &#x60;Pacific/Pago_Pago&#x60; -  * &#x60;Pacific/Palau&#x60; -  * &#x60;Pacific/Pitcairn&#x60; -  * &#x60;Pacific/Pohnpei&#x60; -  * &#x60;Pacific/Port_Moresby&#x60; -  * &#x60;Pacific/Rarotonga&#x60; -  * &#x60;Pacific/Saipan&#x60; -  * &#x60;Pacific/Tahiti&#x60; -  * &#x60;Pacific/Tarawa&#x60; -  * &#x60;Pacific/Tongatapu&#x60; -  * &#x60;Pacific/Wake&#x60; -  * &#x60;Pacific/Wallis&#x60; -  * &#x60;UTC&#x60; - | [optional] [readonly] [default to "Pacific/Niue"]

## Methods

### NewOnpremClusterInfo

`func NewOnpremClusterInfo(classId string, objectType string, ) *OnpremClusterInfo`

NewOnpremClusterInfo instantiates a new OnpremClusterInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnpremClusterInfoWithDefaults

`func NewOnpremClusterInfoWithDefaults() *OnpremClusterInfo`

NewOnpremClusterInfoWithDefaults instantiates a new OnpremClusterInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OnpremClusterInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OnpremClusterInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OnpremClusterInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OnpremClusterInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OnpremClusterInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OnpremClusterInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetNodes

`func (o *OnpremClusterInfo) GetNodes() []OnpremClusterNode`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *OnpremClusterInfo) GetNodesOk() (*[]OnpremClusterNode, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *OnpremClusterInfo) SetNodes(v []OnpremClusterNode)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *OnpremClusterInfo) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### SetNodesNil

`func (o *OnpremClusterInfo) SetNodesNil(b bool)`

 SetNodesNil sets the value for Nodes to be an explicit nil

### UnsetNodes
`func (o *OnpremClusterInfo) UnsetNodes()`

UnsetNodes ensures that no value is present for Nodes, not even an explicit nil
### GetQuorumOk

`func (o *OnpremClusterInfo) GetQuorumOk() bool`

GetQuorumOk returns the QuorumOk field if non-nil, zero value otherwise.

### GetQuorumOkOk

`func (o *OnpremClusterInfo) GetQuorumOkOk() (*bool, bool)`

GetQuorumOkOk returns a tuple with the QuorumOk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuorumOk

`func (o *OnpremClusterInfo) SetQuorumOk(v bool)`

SetQuorumOk sets QuorumOk field to given value.

### HasQuorumOk

`func (o *OnpremClusterInfo) HasQuorumOk() bool`

HasQuorumOk returns a boolean if a field has been set.

### GetReportedTime

`func (o *OnpremClusterInfo) GetReportedTime() time.Time`

GetReportedTime returns the ReportedTime field if non-nil, zero value otherwise.

### GetReportedTimeOk

`func (o *OnpremClusterInfo) GetReportedTimeOk() (*time.Time, bool)`

GetReportedTimeOk returns a tuple with the ReportedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportedTime

`func (o *OnpremClusterInfo) SetReportedTime(v time.Time)`

SetReportedTime sets ReportedTime field to given value.

### HasReportedTime

`func (o *OnpremClusterInfo) HasReportedTime() bool`

HasReportedTime returns a boolean if a field has been set.

### GetSetupTime

`func (o *OnpremClusterInfo) GetSetupTime() time.Time`

GetSetupTime returns the SetupTime field if non-nil, zero value otherwise.

### GetSetupTimeOk

`func (o *OnpremClusterInfo) GetSetupTimeOk() (*time.Time, bool)`

GetSetupTimeOk returns a tuple with the SetupTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupTime

`func (o *OnpremClusterInfo) SetSetupTime(v time.Time)`

SetSetupTime sets SetupTime field to given value.

### HasSetupTime

`func (o *OnpremClusterInfo) HasSetupTime() bool`

HasSetupTime returns a boolean if a field has been set.

### GetTimeZone

`func (o *OnpremClusterInfo) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *OnpremClusterInfo) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *OnpremClusterInfo) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *OnpremClusterInfo) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


