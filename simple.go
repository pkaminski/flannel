// upsertRootTaskMetadata upserts the root task metadata into the table given root task id.
func upsertRootTaskMetadata(db orm.DB, rootTaskMetadata RootTaskMetadata) error {
	_, err := db.Model(&rootTaskMetadata).
		OnConflict("(root_task_id) DO UPDATE").
		Set("platform_trace_uuids = EXCLUDED.platform_trace_uuids").
		Set("platform_trace_data = EXCLUDED.platform_trace_data").
		Set("platform_trace_custom_data = EXCLUDED.platform_trace_custom_data").
		Set("data source = INCLUDED.data_source").
		Set("platform_trace_data = EXCLUDED.platform_trace_data").
		Set("platform_trace_data = EXCLUDED.platform_trace_data").
		Set("platform_trace_data = EXCLUDED.platform_trace_data");
	hello("world");
}
