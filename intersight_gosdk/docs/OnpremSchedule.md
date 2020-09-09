# OnpremSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DayOfMonth** | Pointer to **int64** | Schedule a task on a specific day of the month. Valid values are 1 through 31. If monthOfYear is specified, then dayOfMonth value must be valid for that month. DayOfMonth may not be set when dayOfWeek is specfied. | [optional] 
**DayOfWeek** | Pointer to **int64** | Schedule a task on a specific day of the week. Valid values are 1 through 7, with 1 being Sunday. DayOfWeek may not be specfied when dayOfMonth is specified. | [optional] 
**MonthOfYear** | Pointer to **int64** | Schedule a task on a specific month of the year. Valid values are 1 through 12, with 1 being January. | [optional] 
**RepeatInterval** | Pointer to **int64** | Schedule a task to run periodically at an interval. Default unit of the RepeatInterval is in minutes. If the RepeateInterval value is set, then all other properties are ignored by the scheduler. RepeateInterval constraints are enforced by the services that use the schedule. Each service has pre-configured service specific properties for enforcing minimum and maximum values of the RepeatInterval. | [optional] 
**TimeOfDay** | Pointer to **int64** | Time of the day in seconds. TimeOfDay is required for all schedule configurations, except when the RepeateInterval field is specified. | [optional] 
**TimeZone** | Pointer to **string** | Timezone to use for the schedule calculation. If a timezone value is not speficied, then the schedule calculation will be based on UTC. * &#x60;Pacific/Niue&#x60; -  * &#x60;Pacific/Pago_Pago&#x60; -  * &#x60;Pacific/Honolulu&#x60; -  * &#x60;Pacific/Rarotonga&#x60; -  * &#x60;Pacific/Tahiti&#x60; -  * &#x60;Pacific/Marquesas&#x60; -  * &#x60;America/Anchorage&#x60; -  * &#x60;Pacific/Gambier&#x60; -  * &#x60;America/Los_Angeles&#x60; -  * &#x60;America/Tijuana&#x60; -  * &#x60;America/Vancouver&#x60; -  * &#x60;America/Whitehorse&#x60; -  * &#x60;Pacific/Pitcairn&#x60; -  * &#x60;America/Dawson_Creek&#x60; -  * &#x60;America/Denver&#x60; -  * &#x60;America/Edmonton&#x60; -  * &#x60;America/Hermosillo&#x60; -  * &#x60;America/Mazatlan&#x60; -  * &#x60;America/Phoenix&#x60; -  * &#x60;America/Yellowknife&#x60; -  * &#x60;America/Belize&#x60; -  * &#x60;America/Chicago&#x60; -  * &#x60;America/Costa_Rica&#x60; -  * &#x60;America/El_Salvador&#x60; -  * &#x60;America/Guatemala&#x60; -  * &#x60;America/Managua&#x60; -  * &#x60;America/Mexico_City&#x60; -  * &#x60;America/Regina&#x60; -  * &#x60;America/Tegucigalpa&#x60; -  * &#x60;America/Winnipeg&#x60; -  * &#x60;Pacific/Galapagos&#x60; -  * &#x60;America/Bogota&#x60; -  * &#x60;America/Cancun&#x60; -  * &#x60;America/Cayman&#x60; -  * &#x60;America/Guayaquil&#x60; -  * &#x60;America/Havana&#x60; -  * &#x60;America/Iqaluit&#x60; -  * &#x60;America/Jamaica&#x60; -  * &#x60;America/Lima&#x60; -  * &#x60;America/Nassau&#x60; -  * &#x60;America/New_York&#x60; -  * &#x60;America/Panama&#x60; -  * &#x60;America/Port-au-Prince&#x60; -  * &#x60;America/Rio_Branco&#x60; -  * &#x60;America/Toronto&#x60; -  * &#x60;Pacific/Easter&#x60; -  * &#x60;America/Caracas&#x60; -  * &#x60;America/Asuncion&#x60; -  * &#x60;America/Barbados&#x60; -  * &#x60;America/Boa_Vista&#x60; -  * &#x60;America/Campo_Grande&#x60; -  * &#x60;America/Cuiaba&#x60; -  * &#x60;America/Curacao&#x60; -  * &#x60;America/Grand_Turk&#x60; -  * &#x60;America/Guyana&#x60; -  * &#x60;America/Halifax&#x60; -  * &#x60;America/La_Paz&#x60; -  * &#x60;America/Manaus&#x60; -  * &#x60;America/Martinique&#x60; -  * &#x60;America/Port_of_Spain&#x60; -  * &#x60;America/Porto_Velho&#x60; -  * &#x60;America/Puerto_Rico&#x60; -  * &#x60;America/Santo_Domingo&#x60; -  * &#x60;America/Thule&#x60; -  * &#x60;Atlantic/Bermuda&#x60; -  * &#x60;America/St_Johns&#x60; -  * &#x60;America/Araguaina&#x60; -  * &#x60;America/Argentina/Buenos_Aires&#x60; -  * &#x60;America/Bahia&#x60; -  * &#x60;America/Belem&#x60; -  * &#x60;America/Cayenne&#x60; -  * &#x60;America/Fortaleza&#x60; -  * &#x60;America/Godthab&#x60; -  * &#x60;America/Maceio&#x60; -  * &#x60;America/Miquelon&#x60; -  * &#x60;America/Montevideo&#x60; -  * &#x60;America/Paramaribo&#x60; -  * &#x60;America/Recife&#x60; -  * &#x60;America/Santiago&#x60; -  * &#x60;America/Sao_Paulo&#x60; -  * &#x60;Antarctica/Palmer&#x60; -  * &#x60;Antarctica/Rothera&#x60; -  * &#x60;Atlantic/Stanley&#x60; -  * &#x60;America/Noronha&#x60; -  * &#x60;Atlantic/South_Georgia&#x60; -  * &#x60;America/Scoresbysund&#x60; -  * &#x60;Atlantic/Azores&#x60; -  * &#x60;Atlantic/Cape_Verde&#x60; -  * &#x60;Africa/Abidjan&#x60; -  * &#x60;Africa/Accra&#x60; -  * &#x60;Africa/Bissau&#x60; -  * &#x60;Africa/Casablanca&#x60; -  * &#x60;Africa/El_Aaiun&#x60; -  * &#x60;Africa/Monrovia&#x60; -  * &#x60;America/Danmarkshavn&#x60; -  * &#x60;Atlantic/Canary&#x60; -  * &#x60;Atlantic/Faroe&#x60; -  * &#x60;Atlantic/Reykjavik&#x60; -  * &#x60;Etc/GMT&#x60; -  * &#x60;Europe/Dublin&#x60; -  * &#x60;Europe/Lisbon&#x60; -  * &#x60;Europe/London&#x60; -  * &#x60;Africa/Algiers&#x60; -  * &#x60;Africa/Ceuta&#x60; -  * &#x60;Africa/Lagos&#x60; -  * &#x60;Africa/Ndjamena&#x60; -  * &#x60;Africa/Tunis&#x60; -  * &#x60;Africa/Windhoek&#x60; -  * &#x60;Europe/Amsterdam&#x60; -  * &#x60;Europe/Andorra&#x60; -  * &#x60;Europe/Belgrade&#x60; -  * &#x60;Europe/Berlin&#x60; -  * &#x60;Europe/Brussels&#x60; -  * &#x60;Europe/Budapest&#x60; -  * &#x60;Europe/Copenhagen&#x60; -  * &#x60;Europe/Gibraltar&#x60; -  * &#x60;Europe/Luxembourg&#x60; -  * &#x60;Europe/Madrid&#x60; -  * &#x60;Europe/Malta&#x60; -  * &#x60;Europe/Monaco&#x60; -  * &#x60;Europe/Oslo&#x60; -  * &#x60;Europe/Paris&#x60; -  * &#x60;Europe/Prague&#x60; -  * &#x60;Europe/Rome&#x60; -  * &#x60;Europe/Stockholm&#x60; -  * &#x60;Europe/Tirane&#x60; -  * &#x60;Europe/Vienna&#x60; -  * &#x60;Europe/Warsaw&#x60; -  * &#x60;Europe/Zurich&#x60; -  * &#x60;Africa/Cairo&#x60; -  * &#x60;Africa/Johannesburg&#x60; -  * &#x60;Africa/Maputo&#x60; -  * &#x60;Africa/Tripoli&#x60; -  * &#x60;Asia/Amman&#x60; -  * &#x60;Asia/Beirut&#x60; -  * &#x60;Asia/Damascus&#x60; -  * &#x60;Asia/Gaza&#x60; -  * &#x60;Asia/Jerusalem&#x60; -  * &#x60;Asia/Nicosia&#x60; -  * &#x60;Europe/Athens&#x60; -  * &#x60;Europe/Bucharest&#x60; -  * &#x60;Europe/Chisinau&#x60; -  * &#x60;Europe/Helsinki&#x60; -  * &#x60;Europe/Istanbul&#x60; -  * &#x60;Europe/Kaliningrad&#x60; -  * &#x60;Europe/Kiev&#x60; -  * &#x60;Europe/Riga&#x60; -  * &#x60;Europe/Sofia&#x60; -  * &#x60;Europe/Tallinn&#x60; -  * &#x60;Europe/Vilnius&#x60; -  * &#x60;Africa/Khartoum&#x60; -  * &#x60;Africa/Nairobi&#x60; -  * &#x60;Antarctica/Syowa&#x60; -  * &#x60;Asia/Baghdad&#x60; -  * &#x60;Asia/Qatar&#x60; -  * &#x60;Asia/Riyadh&#x60; -  * &#x60;Europe/Minsk&#x60; -  * &#x60;Europe/Moscow&#x60; -  * &#x60;Asia/Tehran&#x60; -  * &#x60;Asia/Baku&#x60; -  * &#x60;Asia/Dubai&#x60; -  * &#x60;Asia/Tbilisi&#x60; -  * &#x60;Asia/Yerevan&#x60; -  * &#x60;Europe/Samara&#x60; -  * &#x60;Indian/Mahe&#x60; -  * &#x60;Indian/Mauritius&#x60; -  * &#x60;Indian/Reunion&#x60; -  * &#x60;Asia/Kabul&#x60; -  * &#x60;Antarctica/Mawson&#x60; -  * &#x60;Asia/Aqtau&#x60; -  * &#x60;Asia/Aqtobe&#x60; -  * &#x60;Asia/Ashgabat&#x60; -  * &#x60;Asia/Dushanbe&#x60; -  * &#x60;Asia/Karachi&#x60; -  * &#x60;Asia/Tashkent&#x60; -  * &#x60;Asia/Yekaterinburg&#x60; -  * &#x60;Indian/Kerguelen&#x60; -  * &#x60;Indian/Maldives&#x60; -  * &#x60;Asia/Calcutta&#x60; -  * &#x60;Asia/Kolkata&#x60; -  * &#x60;Asia/Colombo&#x60; -  * &#x60;Asia/Katmandu&#x60; -  * &#x60;Antarctica/Vostok&#x60; -  * &#x60;Asia/Almaty&#x60; -  * &#x60;Asia/Bishkek&#x60; -  * &#x60;Asia/Dhaka&#x60; -  * &#x60;Asia/Omsk&#x60; -  * &#x60;Asia/Thimphu&#x60; -  * &#x60;Indian/Chagos&#x60; -  * &#x60;Asia/Rangoon&#x60; -  * &#x60;Indian/Cocos&#x60; -  * &#x60;Antarctica/Davis&#x60; -  * &#x60;Asia/Bangkok&#x60; -  * &#x60;Asia/Hovd&#x60; -  * &#x60;Asia/Jakarta&#x60; -  * &#x60;Asia/Krasnoyarsk&#x60; -  * &#x60;Asia/Saigon&#x60; -  * &#x60;Indian/Christmas&#x60; -  * &#x60;Antarctica/Casey&#x60; -  * &#x60;Asia/Brunei&#x60; -  * &#x60;Asia/Choibalsan&#x60; -  * &#x60;Asia/Hong_Kong&#x60; -  * &#x60;Asia/Irkutsk&#x60; -  * &#x60;Asia/Kuala_Lumpur&#x60; -  * &#x60;Asia/Macau&#x60; -  * &#x60;Asia/Makassar&#x60; -  * &#x60;Asia/Manila&#x60; -  * &#x60;Asia/Shanghai&#x60; -  * &#x60;Asia/Singapore&#x60; -  * &#x60;Asia/Taipei&#x60; -  * &#x60;Asia/Ulaanbaatar&#x60; -  * &#x60;Australia/Perth&#x60; -  * &#x60;Asia/Pyongyang&#x60; -  * &#x60;Asia/Dili&#x60; -  * &#x60;Asia/Jayapura&#x60; -  * &#x60;Asia/Seoul&#x60; -  * &#x60;Asia/Tokyo&#x60; -  * &#x60;Asia/Yakutsk&#x60; -  * &#x60;Pacific/Palau&#x60; -  * &#x60;Australia/Adelaide&#x60; -  * &#x60;Australia/Darwin&#x60; -  * &#x60;Antarctica/DumontDUrville&#x60; -  * &#x60;Asia/Magadan&#x60; -  * &#x60;Asia/Vladivostok&#x60; -  * &#x60;Australia/Brisbane&#x60; -  * &#x60;Australia/Hobart&#x60; -  * &#x60;Australia/Sydney&#x60; -  * &#x60;Pacific/Chuuk&#x60; -  * &#x60;Pacific/Guam&#x60; -  * &#x60;Pacific/Port_Moresby&#x60; -  * &#x60;Pacific/Efate&#x60; -  * &#x60;Pacific/Guadalcanal&#x60; -  * &#x60;Pacific/Kosrae&#x60; -  * &#x60;Pacific/Norfolk&#x60; -  * &#x60;Pacific/Noumea&#x60; -  * &#x60;Pacific/Pohnpei&#x60; -  * &#x60;Asia/Kamchatka&#x60; -  * &#x60;Pacific/Auckland&#x60; -  * &#x60;Pacific/Fiji&#x60; -  * &#x60;Pacific/Funafuti&#x60; -  * &#x60;Pacific/Kwajalein&#x60; -  * &#x60;Pacific/Majuro&#x60; -  * &#x60;Pacific/Nauru&#x60; -  * &#x60;Pacific/Tarawa&#x60; -  * &#x60;Pacific/Wake&#x60; -  * &#x60;Pacific/Wallis&#x60; -  * &#x60;Pacific/Apia&#x60; -  * &#x60;Pacific/Enderbury&#x60; -  * &#x60;Pacific/Fakaofo&#x60; -  * &#x60;Pacific/Tongatapu&#x60; -  * &#x60;Pacific/Kiritimati&#x60; - | [optional] [default to "Pacific/Niue"]
**WeekOfMonth** | Pointer to **int64** | Schedule a task on a specific week of the month. Valid values are 1 through 5. Value of 5 means last week of the month. WeekOfMonth may not be set when dayOfMonth is specified. | [optional] 

## Methods

### NewOnpremSchedule

`func NewOnpremSchedule() *OnpremSchedule`

NewOnpremSchedule instantiates a new OnpremSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnpremScheduleWithDefaults

`func NewOnpremScheduleWithDefaults() *OnpremSchedule`

NewOnpremScheduleWithDefaults instantiates a new OnpremSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDayOfMonth

`func (o *OnpremSchedule) GetDayOfMonth() int64`

GetDayOfMonth returns the DayOfMonth field if non-nil, zero value otherwise.

### GetDayOfMonthOk

`func (o *OnpremSchedule) GetDayOfMonthOk() (*int64, bool)`

GetDayOfMonthOk returns a tuple with the DayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfMonth

`func (o *OnpremSchedule) SetDayOfMonth(v int64)`

SetDayOfMonth sets DayOfMonth field to given value.

### HasDayOfMonth

`func (o *OnpremSchedule) HasDayOfMonth() bool`

HasDayOfMonth returns a boolean if a field has been set.

### GetDayOfWeek

`func (o *OnpremSchedule) GetDayOfWeek() int64`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *OnpremSchedule) GetDayOfWeekOk() (*int64, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *OnpremSchedule) SetDayOfWeek(v int64)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *OnpremSchedule) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### GetMonthOfYear

`func (o *OnpremSchedule) GetMonthOfYear() int64`

GetMonthOfYear returns the MonthOfYear field if non-nil, zero value otherwise.

### GetMonthOfYearOk

`func (o *OnpremSchedule) GetMonthOfYearOk() (*int64, bool)`

GetMonthOfYearOk returns a tuple with the MonthOfYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthOfYear

`func (o *OnpremSchedule) SetMonthOfYear(v int64)`

SetMonthOfYear sets MonthOfYear field to given value.

### HasMonthOfYear

`func (o *OnpremSchedule) HasMonthOfYear() bool`

HasMonthOfYear returns a boolean if a field has been set.

### GetRepeatInterval

`func (o *OnpremSchedule) GetRepeatInterval() int64`

GetRepeatInterval returns the RepeatInterval field if non-nil, zero value otherwise.

### GetRepeatIntervalOk

`func (o *OnpremSchedule) GetRepeatIntervalOk() (*int64, bool)`

GetRepeatIntervalOk returns a tuple with the RepeatInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatInterval

`func (o *OnpremSchedule) SetRepeatInterval(v int64)`

SetRepeatInterval sets RepeatInterval field to given value.

### HasRepeatInterval

`func (o *OnpremSchedule) HasRepeatInterval() bool`

HasRepeatInterval returns a boolean if a field has been set.

### GetTimeOfDay

`func (o *OnpremSchedule) GetTimeOfDay() int64`

GetTimeOfDay returns the TimeOfDay field if non-nil, zero value otherwise.

### GetTimeOfDayOk

`func (o *OnpremSchedule) GetTimeOfDayOk() (*int64, bool)`

GetTimeOfDayOk returns a tuple with the TimeOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOfDay

`func (o *OnpremSchedule) SetTimeOfDay(v int64)`

SetTimeOfDay sets TimeOfDay field to given value.

### HasTimeOfDay

`func (o *OnpremSchedule) HasTimeOfDay() bool`

HasTimeOfDay returns a boolean if a field has been set.

### GetTimeZone

`func (o *OnpremSchedule) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *OnpremSchedule) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *OnpremSchedule) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *OnpremSchedule) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetWeekOfMonth

`func (o *OnpremSchedule) GetWeekOfMonth() int64`

GetWeekOfMonth returns the WeekOfMonth field if non-nil, zero value otherwise.

### GetWeekOfMonthOk

`func (o *OnpremSchedule) GetWeekOfMonthOk() (*int64, bool)`

GetWeekOfMonthOk returns a tuple with the WeekOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekOfMonth

`func (o *OnpremSchedule) SetWeekOfMonth(v int64)`

SetWeekOfMonth sets WeekOfMonth field to given value.

### HasWeekOfMonth

`func (o *OnpremSchedule) HasWeekOfMonth() bool`

HasWeekOfMonth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


