# TelemetryDruidQueryContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrandTotal** | Pointer to **bool** | Druid can include an extra \&quot;grand totals\&quot; row as the last row of a timeseries result set. To enable this, set \&quot;grandTotal\&quot; to true. The grand totals row will appear as the last row in the result array, and will have no timestamp. It will be the last row even if the query is run in \&quot;descending\&quot; mode. Post-aggregations in the grand totals row will be computed based upon the grand total aggregations. | [optional] 
**SkipEmptyBuckets** | Pointer to **bool** | Timeseries queries normally fill empty interior time buckets with zeroes. Time buckets that lie completely outside the data interval are not zero-filled. You can disable all zero-filling with this flag. In this mode, the data point for empty buckets are omitted from the results. | [optional] 
**Timeout** | Pointer to **int32** | Query timeout in milliseconds, beyond which unfinished queries will be cancelled. 0 timeout means no timeout. | [optional] 
**Priority** | Pointer to **int32** | Query Priority. Queries with higher priority get precedence for computational resources. | [optional] 
**QueryId** | Pointer to **string** | Unique identifier given to this query. If a query ID is set or known, this can be used to cancel the query. | [optional] 
**UseCache** | Pointer to **bool** | Flag indicating whether to leverage the query cache for this query. When set to false, it disables reading from the query cache for this query. When set to true, Apache Druid uses druid.broker.cache.useCache or druid.historical.cache.useCache to determine whether or not to read from the query cache. | [optional] 
**PopulateCache** | Pointer to **bool** | Flag indicating whether to save the results of the query to the query cache. Primarily used for debugging. When set to false, it disables saving the results of this query to the query cache. When set to true, Druid uses druid.broker.cache.populateCache or druid.historical.cache.populateCache to determine whether or not to save the results of this query to the query cache. | [optional] 
**UseResultLevelCache** | Pointer to **bool** | Flag indicating whether to leverage the result level cache for this query. When set to false, it disables reading from the query cache for this query. When set to true, Druid uses druid.broker.cache.useResultLevelCache to determine whether or not to read from the result-level query cache. | [optional] 
**PopulateResultLevelCache** | Pointer to **bool** | Flag indicating whether to save the results of the query to the result level cache. Primarily used for debugging. When set to false, it disables saving the results of this query to the query cache. When set to true, Druid uses druid.broker.cache.populateResultLevelCache to determine whether or not to save the results of this query to the result-level query cache. | [optional] 
**BySegment** | Pointer to **bool** | Return \&quot;by segment\&quot; results. Primarily used for debugging, setting it to true returns results associated with the data segment they came from. | [optional] 
**Finalize** | Pointer to **bool** | Flag indicating whether to \&quot;finalize\&quot; aggregation results. Primarily used for debugging. For instance, the hyperUnique aggregator will return the full HyperLogLog sketch instead of the estimated cardinality when this flag is set to false. | [optional] 
**ChunkPeriod** | Pointer to **string** | At the Broker process level, long interval queries (of any type) may be broken into shorter interval queries to parallelize merging more than normal. Broken up queries will use a larger share of cluster resources, but, if you use groupBy \&quot;v1, it may be able to complete faster as a result. Use ISO 8601 periods. For example, if this property is set to P1M (one month), then a query covering a year would be broken into 12 smaller queries. The broker uses its query processing executor service to initiate processing for query chunks, so make sure druid.processing.numThreads is configured appropriately on the broker. groupBy queries do not support chunkPeriod by default, although they do if using the legacy \&quot;v1\&quot; engine. This context is deprecated since it&#39;s only useful for groupBy \&quot;v1\&quot;, and will be removed in the future releases. | [optional] 
**MaxScatterGatherBytes** | Pointer to **int32** | Maximum number of bytes gathered from data processes such as Historicals and realtime processes to execute a query. This parameter can be used to further reduce maxScatterGatherBytes limit at query time. | [optional] 
**MaxQueuedBytes** | Pointer to **int32** | Maximum number of bytes queued per query before exerting backpressure on the channel to the data server. Similar to maxScatterGatherBytes, except unlike that configuration, this one will trigger backpressure rather than query failure. Zero means disabled. | [optional] 
**SerializeDateTimeAsLong** | Pointer to **bool** | If true, DateTime is serialized as long in the result returned by Broker and the data transportation between Broker and compute process. | [optional] 
**SerializeDateTimeAsLongInner** | Pointer to **bool** | If true, DateTime is serialized as long in the data transportation between Broker and compute process. | [optional] 
**EnableParallelMerge** | Pointer to **bool** | Enable parallel result merging on the Broker. Note that druid.processing.merge.useParallelMergePool must be enabled for this setting to be set to true. | [optional] 
**ParallelMergeParallelism** | Pointer to **int32** | Maximum number of parallel threads to use for parallel result merging on the Broker. | [optional] 
**ParallelMergeInitialYieldRows** | Pointer to **int32** | Number of rows to yield per ForkJoinPool merge task for parallel result merging on the Broker, before forking off a new task to continue merging sequences. | [optional] 
**ParallelMergeSmallBatchRows** | Pointer to **int32** | Size of result batches to operate on in ForkJoinPool merge tasks for parallel result merging on the Broker. | [optional] 

## Methods

### NewTelemetryDruidQueryContext

`func NewTelemetryDruidQueryContext() *TelemetryDruidQueryContext`

NewTelemetryDruidQueryContext instantiates a new TelemetryDruidQueryContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidQueryContextWithDefaults

`func NewTelemetryDruidQueryContextWithDefaults() *TelemetryDruidQueryContext`

NewTelemetryDruidQueryContextWithDefaults instantiates a new TelemetryDruidQueryContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrandTotal

`func (o *TelemetryDruidQueryContext) GetGrandTotal() bool`

GetGrandTotal returns the GrandTotal field if non-nil, zero value otherwise.

### GetGrandTotalOk

`func (o *TelemetryDruidQueryContext) GetGrandTotalOk() (*bool, bool)`

GetGrandTotalOk returns a tuple with the GrandTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrandTotal

`func (o *TelemetryDruidQueryContext) SetGrandTotal(v bool)`

SetGrandTotal sets GrandTotal field to given value.

### HasGrandTotal

`func (o *TelemetryDruidQueryContext) HasGrandTotal() bool`

HasGrandTotal returns a boolean if a field has been set.

### GetSkipEmptyBuckets

`func (o *TelemetryDruidQueryContext) GetSkipEmptyBuckets() bool`

GetSkipEmptyBuckets returns the SkipEmptyBuckets field if non-nil, zero value otherwise.

### GetSkipEmptyBucketsOk

`func (o *TelemetryDruidQueryContext) GetSkipEmptyBucketsOk() (*bool, bool)`

GetSkipEmptyBucketsOk returns a tuple with the SkipEmptyBuckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipEmptyBuckets

`func (o *TelemetryDruidQueryContext) SetSkipEmptyBuckets(v bool)`

SetSkipEmptyBuckets sets SkipEmptyBuckets field to given value.

### HasSkipEmptyBuckets

`func (o *TelemetryDruidQueryContext) HasSkipEmptyBuckets() bool`

HasSkipEmptyBuckets returns a boolean if a field has been set.

### GetTimeout

`func (o *TelemetryDruidQueryContext) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *TelemetryDruidQueryContext) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *TelemetryDruidQueryContext) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *TelemetryDruidQueryContext) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetPriority

`func (o *TelemetryDruidQueryContext) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TelemetryDruidQueryContext) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TelemetryDruidQueryContext) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *TelemetryDruidQueryContext) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetQueryId

`func (o *TelemetryDruidQueryContext) GetQueryId() string`

GetQueryId returns the QueryId field if non-nil, zero value otherwise.

### GetQueryIdOk

`func (o *TelemetryDruidQueryContext) GetQueryIdOk() (*string, bool)`

GetQueryIdOk returns a tuple with the QueryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryId

`func (o *TelemetryDruidQueryContext) SetQueryId(v string)`

SetQueryId sets QueryId field to given value.

### HasQueryId

`func (o *TelemetryDruidQueryContext) HasQueryId() bool`

HasQueryId returns a boolean if a field has been set.

### GetUseCache

`func (o *TelemetryDruidQueryContext) GetUseCache() bool`

GetUseCache returns the UseCache field if non-nil, zero value otherwise.

### GetUseCacheOk

`func (o *TelemetryDruidQueryContext) GetUseCacheOk() (*bool, bool)`

GetUseCacheOk returns a tuple with the UseCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCache

`func (o *TelemetryDruidQueryContext) SetUseCache(v bool)`

SetUseCache sets UseCache field to given value.

### HasUseCache

`func (o *TelemetryDruidQueryContext) HasUseCache() bool`

HasUseCache returns a boolean if a field has been set.

### GetPopulateCache

`func (o *TelemetryDruidQueryContext) GetPopulateCache() bool`

GetPopulateCache returns the PopulateCache field if non-nil, zero value otherwise.

### GetPopulateCacheOk

`func (o *TelemetryDruidQueryContext) GetPopulateCacheOk() (*bool, bool)`

GetPopulateCacheOk returns a tuple with the PopulateCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopulateCache

`func (o *TelemetryDruidQueryContext) SetPopulateCache(v bool)`

SetPopulateCache sets PopulateCache field to given value.

### HasPopulateCache

`func (o *TelemetryDruidQueryContext) HasPopulateCache() bool`

HasPopulateCache returns a boolean if a field has been set.

### GetUseResultLevelCache

`func (o *TelemetryDruidQueryContext) GetUseResultLevelCache() bool`

GetUseResultLevelCache returns the UseResultLevelCache field if non-nil, zero value otherwise.

### GetUseResultLevelCacheOk

`func (o *TelemetryDruidQueryContext) GetUseResultLevelCacheOk() (*bool, bool)`

GetUseResultLevelCacheOk returns a tuple with the UseResultLevelCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseResultLevelCache

`func (o *TelemetryDruidQueryContext) SetUseResultLevelCache(v bool)`

SetUseResultLevelCache sets UseResultLevelCache field to given value.

### HasUseResultLevelCache

`func (o *TelemetryDruidQueryContext) HasUseResultLevelCache() bool`

HasUseResultLevelCache returns a boolean if a field has been set.

### GetPopulateResultLevelCache

`func (o *TelemetryDruidQueryContext) GetPopulateResultLevelCache() bool`

GetPopulateResultLevelCache returns the PopulateResultLevelCache field if non-nil, zero value otherwise.

### GetPopulateResultLevelCacheOk

`func (o *TelemetryDruidQueryContext) GetPopulateResultLevelCacheOk() (*bool, bool)`

GetPopulateResultLevelCacheOk returns a tuple with the PopulateResultLevelCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopulateResultLevelCache

`func (o *TelemetryDruidQueryContext) SetPopulateResultLevelCache(v bool)`

SetPopulateResultLevelCache sets PopulateResultLevelCache field to given value.

### HasPopulateResultLevelCache

`func (o *TelemetryDruidQueryContext) HasPopulateResultLevelCache() bool`

HasPopulateResultLevelCache returns a boolean if a field has been set.

### GetBySegment

`func (o *TelemetryDruidQueryContext) GetBySegment() bool`

GetBySegment returns the BySegment field if non-nil, zero value otherwise.

### GetBySegmentOk

`func (o *TelemetryDruidQueryContext) GetBySegmentOk() (*bool, bool)`

GetBySegmentOk returns a tuple with the BySegment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBySegment

`func (o *TelemetryDruidQueryContext) SetBySegment(v bool)`

SetBySegment sets BySegment field to given value.

### HasBySegment

`func (o *TelemetryDruidQueryContext) HasBySegment() bool`

HasBySegment returns a boolean if a field has been set.

### GetFinalize

`func (o *TelemetryDruidQueryContext) GetFinalize() bool`

GetFinalize returns the Finalize field if non-nil, zero value otherwise.

### GetFinalizeOk

`func (o *TelemetryDruidQueryContext) GetFinalizeOk() (*bool, bool)`

GetFinalizeOk returns a tuple with the Finalize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalize

`func (o *TelemetryDruidQueryContext) SetFinalize(v bool)`

SetFinalize sets Finalize field to given value.

### HasFinalize

`func (o *TelemetryDruidQueryContext) HasFinalize() bool`

HasFinalize returns a boolean if a field has been set.

### GetChunkPeriod

`func (o *TelemetryDruidQueryContext) GetChunkPeriod() string`

GetChunkPeriod returns the ChunkPeriod field if non-nil, zero value otherwise.

### GetChunkPeriodOk

`func (o *TelemetryDruidQueryContext) GetChunkPeriodOk() (*string, bool)`

GetChunkPeriodOk returns a tuple with the ChunkPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunkPeriod

`func (o *TelemetryDruidQueryContext) SetChunkPeriod(v string)`

SetChunkPeriod sets ChunkPeriod field to given value.

### HasChunkPeriod

`func (o *TelemetryDruidQueryContext) HasChunkPeriod() bool`

HasChunkPeriod returns a boolean if a field has been set.

### GetMaxScatterGatherBytes

`func (o *TelemetryDruidQueryContext) GetMaxScatterGatherBytes() int32`

GetMaxScatterGatherBytes returns the MaxScatterGatherBytes field if non-nil, zero value otherwise.

### GetMaxScatterGatherBytesOk

`func (o *TelemetryDruidQueryContext) GetMaxScatterGatherBytesOk() (*int32, bool)`

GetMaxScatterGatherBytesOk returns a tuple with the MaxScatterGatherBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxScatterGatherBytes

`func (o *TelemetryDruidQueryContext) SetMaxScatterGatherBytes(v int32)`

SetMaxScatterGatherBytes sets MaxScatterGatherBytes field to given value.

### HasMaxScatterGatherBytes

`func (o *TelemetryDruidQueryContext) HasMaxScatterGatherBytes() bool`

HasMaxScatterGatherBytes returns a boolean if a field has been set.

### GetMaxQueuedBytes

`func (o *TelemetryDruidQueryContext) GetMaxQueuedBytes() int32`

GetMaxQueuedBytes returns the MaxQueuedBytes field if non-nil, zero value otherwise.

### GetMaxQueuedBytesOk

`func (o *TelemetryDruidQueryContext) GetMaxQueuedBytesOk() (*int32, bool)`

GetMaxQueuedBytesOk returns a tuple with the MaxQueuedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxQueuedBytes

`func (o *TelemetryDruidQueryContext) SetMaxQueuedBytes(v int32)`

SetMaxQueuedBytes sets MaxQueuedBytes field to given value.

### HasMaxQueuedBytes

`func (o *TelemetryDruidQueryContext) HasMaxQueuedBytes() bool`

HasMaxQueuedBytes returns a boolean if a field has been set.

### GetSerializeDateTimeAsLong

`func (o *TelemetryDruidQueryContext) GetSerializeDateTimeAsLong() bool`

GetSerializeDateTimeAsLong returns the SerializeDateTimeAsLong field if non-nil, zero value otherwise.

### GetSerializeDateTimeAsLongOk

`func (o *TelemetryDruidQueryContext) GetSerializeDateTimeAsLongOk() (*bool, bool)`

GetSerializeDateTimeAsLongOk returns a tuple with the SerializeDateTimeAsLong field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerializeDateTimeAsLong

`func (o *TelemetryDruidQueryContext) SetSerializeDateTimeAsLong(v bool)`

SetSerializeDateTimeAsLong sets SerializeDateTimeAsLong field to given value.

### HasSerializeDateTimeAsLong

`func (o *TelemetryDruidQueryContext) HasSerializeDateTimeAsLong() bool`

HasSerializeDateTimeAsLong returns a boolean if a field has been set.

### GetSerializeDateTimeAsLongInner

`func (o *TelemetryDruidQueryContext) GetSerializeDateTimeAsLongInner() bool`

GetSerializeDateTimeAsLongInner returns the SerializeDateTimeAsLongInner field if non-nil, zero value otherwise.

### GetSerializeDateTimeAsLongInnerOk

`func (o *TelemetryDruidQueryContext) GetSerializeDateTimeAsLongInnerOk() (*bool, bool)`

GetSerializeDateTimeAsLongInnerOk returns a tuple with the SerializeDateTimeAsLongInner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerializeDateTimeAsLongInner

`func (o *TelemetryDruidQueryContext) SetSerializeDateTimeAsLongInner(v bool)`

SetSerializeDateTimeAsLongInner sets SerializeDateTimeAsLongInner field to given value.

### HasSerializeDateTimeAsLongInner

`func (o *TelemetryDruidQueryContext) HasSerializeDateTimeAsLongInner() bool`

HasSerializeDateTimeAsLongInner returns a boolean if a field has been set.

### GetEnableParallelMerge

`func (o *TelemetryDruidQueryContext) GetEnableParallelMerge() bool`

GetEnableParallelMerge returns the EnableParallelMerge field if non-nil, zero value otherwise.

### GetEnableParallelMergeOk

`func (o *TelemetryDruidQueryContext) GetEnableParallelMergeOk() (*bool, bool)`

GetEnableParallelMergeOk returns a tuple with the EnableParallelMerge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableParallelMerge

`func (o *TelemetryDruidQueryContext) SetEnableParallelMerge(v bool)`

SetEnableParallelMerge sets EnableParallelMerge field to given value.

### HasEnableParallelMerge

`func (o *TelemetryDruidQueryContext) HasEnableParallelMerge() bool`

HasEnableParallelMerge returns a boolean if a field has been set.

### GetParallelMergeParallelism

`func (o *TelemetryDruidQueryContext) GetParallelMergeParallelism() int32`

GetParallelMergeParallelism returns the ParallelMergeParallelism field if non-nil, zero value otherwise.

### GetParallelMergeParallelismOk

`func (o *TelemetryDruidQueryContext) GetParallelMergeParallelismOk() (*int32, bool)`

GetParallelMergeParallelismOk returns a tuple with the ParallelMergeParallelism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelMergeParallelism

`func (o *TelemetryDruidQueryContext) SetParallelMergeParallelism(v int32)`

SetParallelMergeParallelism sets ParallelMergeParallelism field to given value.

### HasParallelMergeParallelism

`func (o *TelemetryDruidQueryContext) HasParallelMergeParallelism() bool`

HasParallelMergeParallelism returns a boolean if a field has been set.

### GetParallelMergeInitialYieldRows

`func (o *TelemetryDruidQueryContext) GetParallelMergeInitialYieldRows() int32`

GetParallelMergeInitialYieldRows returns the ParallelMergeInitialYieldRows field if non-nil, zero value otherwise.

### GetParallelMergeInitialYieldRowsOk

`func (o *TelemetryDruidQueryContext) GetParallelMergeInitialYieldRowsOk() (*int32, bool)`

GetParallelMergeInitialYieldRowsOk returns a tuple with the ParallelMergeInitialYieldRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelMergeInitialYieldRows

`func (o *TelemetryDruidQueryContext) SetParallelMergeInitialYieldRows(v int32)`

SetParallelMergeInitialYieldRows sets ParallelMergeInitialYieldRows field to given value.

### HasParallelMergeInitialYieldRows

`func (o *TelemetryDruidQueryContext) HasParallelMergeInitialYieldRows() bool`

HasParallelMergeInitialYieldRows returns a boolean if a field has been set.

### GetParallelMergeSmallBatchRows

`func (o *TelemetryDruidQueryContext) GetParallelMergeSmallBatchRows() int32`

GetParallelMergeSmallBatchRows returns the ParallelMergeSmallBatchRows field if non-nil, zero value otherwise.

### GetParallelMergeSmallBatchRowsOk

`func (o *TelemetryDruidQueryContext) GetParallelMergeSmallBatchRowsOk() (*int32, bool)`

GetParallelMergeSmallBatchRowsOk returns a tuple with the ParallelMergeSmallBatchRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelMergeSmallBatchRows

`func (o *TelemetryDruidQueryContext) SetParallelMergeSmallBatchRows(v int32)`

SetParallelMergeSmallBatchRows sets ParallelMergeSmallBatchRows field to given value.

### HasParallelMergeSmallBatchRows

`func (o *TelemetryDruidQueryContext) HasParallelMergeSmallBatchRows() bool`

HasParallelMergeSmallBatchRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


