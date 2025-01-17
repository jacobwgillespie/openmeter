// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BalanceSnapshotsColumns holds the columns for the "balance_snapshots" table.
	BalanceSnapshotsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "namespace", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "grant_balances", Type: field.TypeJSON, SchemaType: map[string]string{"postgres": "jsonb"}},
		{Name: "balance", Type: field.TypeFloat64, SchemaType: map[string]string{"postgres": "numeric"}},
		{Name: "overage", Type: field.TypeFloat64, SchemaType: map[string]string{"postgres": "numeric"}},
		{Name: "at", Type: field.TypeTime},
		{Name: "owner_id", Type: field.TypeString, SchemaType: map[string]string{"postgres": "char(26)"}},
	}
	// BalanceSnapshotsTable holds the schema information for the "balance_snapshots" table.
	BalanceSnapshotsTable = &schema.Table{
		Name:       "balance_snapshots",
		Columns:    BalanceSnapshotsColumns,
		PrimaryKey: []*schema.Column{BalanceSnapshotsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "balance_snapshots_entitlements_balance_snapshot",
				Columns:    []*schema.Column{BalanceSnapshotsColumns[9]},
				RefColumns: []*schema.Column{EntitlementsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "balancesnapshot_namespace_at",
				Unique:  false,
				Columns: []*schema.Column{BalanceSnapshotsColumns[1], BalanceSnapshotsColumns[8]},
			},
			{
				Name:    "balancesnapshot_namespace_balance",
				Unique:  false,
				Columns: []*schema.Column{BalanceSnapshotsColumns[1], BalanceSnapshotsColumns[6]},
			},
			{
				Name:    "balancesnapshot_namespace_balance_at",
				Unique:  false,
				Columns: []*schema.Column{BalanceSnapshotsColumns[1], BalanceSnapshotsColumns[6], BalanceSnapshotsColumns[8]},
			},
		},
	}
	// EntitlementsColumns holds the columns for the "entitlements" table.
	EntitlementsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, SchemaType: map[string]string{"postgres": "char(26)"}},
		{Name: "namespace", Type: field.TypeString},
		{Name: "metadata", Type: field.TypeJSON, Nullable: true, SchemaType: map[string]string{"postgres": "jsonb"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "entitlement_type", Type: field.TypeEnum, Enums: []string{"metered", "static", "boolean"}},
		{Name: "feature_key", Type: field.TypeString},
		{Name: "subject_key", Type: field.TypeString},
		{Name: "measure_usage_from", Type: field.TypeTime, Nullable: true},
		{Name: "issue_after_reset", Type: field.TypeFloat64, Nullable: true},
		{Name: "issue_after_reset_priority", Type: field.TypeUint8, Nullable: true},
		{Name: "is_soft_limit", Type: field.TypeBool, Nullable: true},
		{Name: "preserve_overage_at_reset", Type: field.TypeBool, Nullable: true},
		{Name: "config", Type: field.TypeJSON, Nullable: true, SchemaType: map[string]string{"postgres": "jsonb"}},
		{Name: "usage_period_interval", Type: field.TypeEnum, Nullable: true, Enums: []string{"DAY", "WEEK", "MONTH", "YEAR"}},
		{Name: "usage_period_anchor", Type: field.TypeTime, Nullable: true},
		{Name: "current_usage_period_start", Type: field.TypeTime, Nullable: true},
		{Name: "current_usage_period_end", Type: field.TypeTime, Nullable: true},
		{Name: "feature_id", Type: field.TypeString, SchemaType: map[string]string{"postgres": "char(26)"}},
	}
	// EntitlementsTable holds the schema information for the "entitlements" table.
	EntitlementsTable = &schema.Table{
		Name:       "entitlements",
		Columns:    EntitlementsColumns,
		PrimaryKey: []*schema.Column{EntitlementsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "entitlements_features_entitlement",
				Columns:    []*schema.Column{EntitlementsColumns[19]},
				RefColumns: []*schema.Column{FeaturesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "entitlement_id",
				Unique:  false,
				Columns: []*schema.Column{EntitlementsColumns[0]},
			},
			{
				Name:    "entitlement_namespace_id",
				Unique:  false,
				Columns: []*schema.Column{EntitlementsColumns[1], EntitlementsColumns[0]},
			},
			{
				Name:    "entitlement_namespace_subject_key",
				Unique:  false,
				Columns: []*schema.Column{EntitlementsColumns[1], EntitlementsColumns[8]},
			},
			{
				Name:    "entitlement_namespace_id_subject_key",
				Unique:  false,
				Columns: []*schema.Column{EntitlementsColumns[1], EntitlementsColumns[0], EntitlementsColumns[8]},
			},
			{
				Name:    "entitlement_namespace_feature_id_id",
				Unique:  false,
				Columns: []*schema.Column{EntitlementsColumns[1], EntitlementsColumns[19], EntitlementsColumns[0]},
			},
			{
				Name:    "entitlement_namespace_current_usage_period_end",
				Unique:  false,
				Columns: []*schema.Column{EntitlementsColumns[1], EntitlementsColumns[18]},
			},
		},
	}
	// FeaturesColumns holds the columns for the "features" table.
	FeaturesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, SchemaType: map[string]string{"postgres": "char(26)"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "metadata", Type: field.TypeJSON, Nullable: true, SchemaType: map[string]string{"postgres": "jsonb"}},
		{Name: "namespace", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "key", Type: field.TypeString},
		{Name: "meter_slug", Type: field.TypeString, Nullable: true},
		{Name: "meter_group_by_filters", Type: field.TypeJSON, Nullable: true},
		{Name: "archived_at", Type: field.TypeTime, Nullable: true},
	}
	// FeaturesTable holds the schema information for the "features" table.
	FeaturesTable = &schema.Table{
		Name:       "features",
		Columns:    FeaturesColumns,
		PrimaryKey: []*schema.Column{FeaturesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "feature_id",
				Unique:  false,
				Columns: []*schema.Column{FeaturesColumns[0]},
			},
			{
				Name:    "feature_namespace_id",
				Unique:  false,
				Columns: []*schema.Column{FeaturesColumns[5], FeaturesColumns[0]},
			},
		},
	}
	// GrantsColumns holds the columns for the "grants" table.
	GrantsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, SchemaType: map[string]string{"postgres": "char(26)"}},
		{Name: "namespace", Type: field.TypeString},
		{Name: "metadata", Type: field.TypeJSON, Nullable: true, SchemaType: map[string]string{"postgres": "jsonb"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "amount", Type: field.TypeFloat64, SchemaType: map[string]string{"postgres": "numeric"}},
		{Name: "priority", Type: field.TypeUint8, Default: 0},
		{Name: "effective_at", Type: field.TypeTime},
		{Name: "expiration", Type: field.TypeJSON, SchemaType: map[string]string{"postgres": "jsonb"}},
		{Name: "expires_at", Type: field.TypeTime},
		{Name: "voided_at", Type: field.TypeTime, Nullable: true},
		{Name: "reset_max_rollover", Type: field.TypeFloat64, SchemaType: map[string]string{"postgres": "numeric"}},
		{Name: "reset_min_rollover", Type: field.TypeFloat64, SchemaType: map[string]string{"postgres": "numeric"}},
		{Name: "recurrence_period", Type: field.TypeEnum, Nullable: true, Enums: []string{"DAY", "WEEK", "MONTH", "YEAR"}},
		{Name: "recurrence_anchor", Type: field.TypeTime, Nullable: true},
		{Name: "owner_id", Type: field.TypeString, SchemaType: map[string]string{"postgres": "char(26)"}},
	}
	// GrantsTable holds the schema information for the "grants" table.
	GrantsTable = &schema.Table{
		Name:       "grants",
		Columns:    GrantsColumns,
		PrimaryKey: []*schema.Column{GrantsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "grants_entitlements_grant",
				Columns:    []*schema.Column{GrantsColumns[16]},
				RefColumns: []*schema.Column{EntitlementsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "grant_id",
				Unique:  false,
				Columns: []*schema.Column{GrantsColumns[0]},
			},
			{
				Name:    "grant_namespace_owner_id",
				Unique:  false,
				Columns: []*schema.Column{GrantsColumns[1], GrantsColumns[16]},
			},
			{
				Name:    "grant_effective_at_expires_at",
				Unique:  false,
				Columns: []*schema.Column{GrantsColumns[8], GrantsColumns[10]},
			},
		},
	}
	// NotificationChannelsColumns holds the columns for the "notification_channels" table.
	NotificationChannelsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, SchemaType: map[string]string{"postgres": "char(26)"}},
		{Name: "namespace", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"WEBHOOK"}},
		{Name: "name", Type: field.TypeString},
		{Name: "disabled", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "config", Type: field.TypeString, SchemaType: map[string]string{"postgres": "jsonb"}},
	}
	// NotificationChannelsTable holds the schema information for the "notification_channels" table.
	NotificationChannelsTable = &schema.Table{
		Name:       "notification_channels",
		Columns:    NotificationChannelsColumns,
		PrimaryKey: []*schema.Column{NotificationChannelsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "notificationchannel_id",
				Unique:  false,
				Columns: []*schema.Column{NotificationChannelsColumns[0]},
			},
			{
				Name:    "notificationchannel_namespace_id",
				Unique:  false,
				Columns: []*schema.Column{NotificationChannelsColumns[1], NotificationChannelsColumns[0]},
			},
			{
				Name:    "notificationchannel_namespace_type",
				Unique:  false,
				Columns: []*schema.Column{NotificationChannelsColumns[1], NotificationChannelsColumns[5]},
			},
		},
	}
	// NotificationEventsColumns holds the columns for the "notification_events" table.
	NotificationEventsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, SchemaType: map[string]string{"postgres": "char(26)"}},
		{Name: "namespace", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"entitlements.balance.threshold"}},
		{Name: "payload", Type: field.TypeString, SchemaType: map[string]string{"postgres": "jsonb"}},
		{Name: "rule_id", Type: field.TypeString, SchemaType: map[string]string{"postgres": "char(26)"}},
	}
	// NotificationEventsTable holds the schema information for the "notification_events" table.
	NotificationEventsTable = &schema.Table{
		Name:       "notification_events",
		Columns:    NotificationEventsColumns,
		PrimaryKey: []*schema.Column{NotificationEventsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "notification_events_notification_rules_events",
				Columns:    []*schema.Column{NotificationEventsColumns[5]},
				RefColumns: []*schema.Column{NotificationRulesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "notificationevent_id",
				Unique:  false,
				Columns: []*schema.Column{NotificationEventsColumns[0]},
			},
			{
				Name:    "notificationevent_namespace_id",
				Unique:  false,
				Columns: []*schema.Column{NotificationEventsColumns[1], NotificationEventsColumns[0]},
			},
			{
				Name:    "notificationevent_namespace_type",
				Unique:  false,
				Columns: []*schema.Column{NotificationEventsColumns[1], NotificationEventsColumns[3]},
			},
		},
	}
	// NotificationEventDeliveryStatusColumns holds the columns for the "notification_event_delivery_status" table.
	NotificationEventDeliveryStatusColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, SchemaType: map[string]string{"postgres": "char(26)"}},
		{Name: "namespace", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "event_id", Type: field.TypeString},
		{Name: "channel_id", Type: field.TypeString},
		{Name: "state", Type: field.TypeEnum, Enums: []string{"SUCCESS", "FAILED", "SENDING", "PENDING"}, Default: "PENDING"},
		{Name: "reason", Type: field.TypeString, Nullable: true},
	}
	// NotificationEventDeliveryStatusTable holds the schema information for the "notification_event_delivery_status" table.
	NotificationEventDeliveryStatusTable = &schema.Table{
		Name:       "notification_event_delivery_status",
		Columns:    NotificationEventDeliveryStatusColumns,
		PrimaryKey: []*schema.Column{NotificationEventDeliveryStatusColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "notificationeventdeliverystatus_id",
				Unique:  false,
				Columns: []*schema.Column{NotificationEventDeliveryStatusColumns[0]},
			},
			{
				Name:    "notificationeventdeliverystatus_namespace_id",
				Unique:  false,
				Columns: []*schema.Column{NotificationEventDeliveryStatusColumns[1], NotificationEventDeliveryStatusColumns[0]},
			},
			{
				Name:    "notificationeventdeliverystatus_namespace_event_id_channel_id",
				Unique:  false,
				Columns: []*schema.Column{NotificationEventDeliveryStatusColumns[1], NotificationEventDeliveryStatusColumns[4], NotificationEventDeliveryStatusColumns[5]},
			},
			{
				Name:    "notificationeventdeliverystatus_namespace_state",
				Unique:  false,
				Columns: []*schema.Column{NotificationEventDeliveryStatusColumns[1], NotificationEventDeliveryStatusColumns[6]},
			},
		},
	}
	// NotificationRulesColumns holds the columns for the "notification_rules" table.
	NotificationRulesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, SchemaType: map[string]string{"postgres": "char(26)"}},
		{Name: "namespace", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"entitlements.balance.threshold"}},
		{Name: "name", Type: field.TypeString},
		{Name: "disabled", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "config", Type: field.TypeString, SchemaType: map[string]string{"postgres": "jsonb"}},
	}
	// NotificationRulesTable holds the schema information for the "notification_rules" table.
	NotificationRulesTable = &schema.Table{
		Name:       "notification_rules",
		Columns:    NotificationRulesColumns,
		PrimaryKey: []*schema.Column{NotificationRulesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "notificationrule_id",
				Unique:  false,
				Columns: []*schema.Column{NotificationRulesColumns[0]},
			},
			{
				Name:    "notificationrule_namespace_id",
				Unique:  false,
				Columns: []*schema.Column{NotificationRulesColumns[1], NotificationRulesColumns[0]},
			},
			{
				Name:    "notificationrule_namespace_type",
				Unique:  false,
				Columns: []*schema.Column{NotificationRulesColumns[1], NotificationRulesColumns[5]},
			},
		},
	}
	// UsageResetsColumns holds the columns for the "usage_resets" table.
	UsageResetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, SchemaType: map[string]string{"postgres": "char(26)"}},
		{Name: "namespace", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "reset_time", Type: field.TypeTime},
		{Name: "entitlement_id", Type: field.TypeString, SchemaType: map[string]string{"postgres": "char(26)"}},
	}
	// UsageResetsTable holds the schema information for the "usage_resets" table.
	UsageResetsTable = &schema.Table{
		Name:       "usage_resets",
		Columns:    UsageResetsColumns,
		PrimaryKey: []*schema.Column{UsageResetsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "usage_resets_entitlements_usage_reset",
				Columns:    []*schema.Column{UsageResetsColumns[6]},
				RefColumns: []*schema.Column{EntitlementsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "usagereset_id",
				Unique:  false,
				Columns: []*schema.Column{UsageResetsColumns[0]},
			},
			{
				Name:    "usagereset_namespace_entitlement_id",
				Unique:  false,
				Columns: []*schema.Column{UsageResetsColumns[1], UsageResetsColumns[6]},
			},
			{
				Name:    "usagereset_namespace_entitlement_id_reset_time",
				Unique:  false,
				Columns: []*schema.Column{UsageResetsColumns[1], UsageResetsColumns[6], UsageResetsColumns[5]},
			},
		},
	}
	// NotificationChannelRulesColumns holds the columns for the "notification_channel_rules" table.
	NotificationChannelRulesColumns = []*schema.Column{
		{Name: "notification_channel_id", Type: field.TypeString, SchemaType: map[string]string{"postgres": "char(26)"}},
		{Name: "notification_rule_id", Type: field.TypeString, SchemaType: map[string]string{"postgres": "char(26)"}},
	}
	// NotificationChannelRulesTable holds the schema information for the "notification_channel_rules" table.
	NotificationChannelRulesTable = &schema.Table{
		Name:       "notification_channel_rules",
		Columns:    NotificationChannelRulesColumns,
		PrimaryKey: []*schema.Column{NotificationChannelRulesColumns[0], NotificationChannelRulesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "notification_channel_rules_notification_channel_id",
				Columns:    []*schema.Column{NotificationChannelRulesColumns[0]},
				RefColumns: []*schema.Column{NotificationChannelsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "notification_channel_rules_notification_rule_id",
				Columns:    []*schema.Column{NotificationChannelRulesColumns[1]},
				RefColumns: []*schema.Column{NotificationRulesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// NotificationEventDeliveryStatusEventsColumns holds the columns for the "notification_event_delivery_status_events" table.
	NotificationEventDeliveryStatusEventsColumns = []*schema.Column{
		{Name: "notification_event_delivery_status_id", Type: field.TypeString, SchemaType: map[string]string{"postgres": "char(26)"}},
		{Name: "notification_event_id", Type: field.TypeString, SchemaType: map[string]string{"postgres": "char(26)"}},
	}
	// NotificationEventDeliveryStatusEventsTable holds the schema information for the "notification_event_delivery_status_events" table.
	NotificationEventDeliveryStatusEventsTable = &schema.Table{
		Name:       "notification_event_delivery_status_events",
		Columns:    NotificationEventDeliveryStatusEventsColumns,
		PrimaryKey: []*schema.Column{NotificationEventDeliveryStatusEventsColumns[0], NotificationEventDeliveryStatusEventsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "notification_event_delivery_status_events_notification_event_delivery_status_id",
				Columns:    []*schema.Column{NotificationEventDeliveryStatusEventsColumns[0]},
				RefColumns: []*schema.Column{NotificationEventDeliveryStatusColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "notification_event_delivery_status_events_notification_event_id",
				Columns:    []*schema.Column{NotificationEventDeliveryStatusEventsColumns[1]},
				RefColumns: []*schema.Column{NotificationEventsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BalanceSnapshotsTable,
		EntitlementsTable,
		FeaturesTable,
		GrantsTable,
		NotificationChannelsTable,
		NotificationEventsTable,
		NotificationEventDeliveryStatusTable,
		NotificationRulesTable,
		UsageResetsTable,
		NotificationChannelRulesTable,
		NotificationEventDeliveryStatusEventsTable,
	}
)

func init() {
	BalanceSnapshotsTable.ForeignKeys[0].RefTable = EntitlementsTable
	EntitlementsTable.ForeignKeys[0].RefTable = FeaturesTable
	GrantsTable.ForeignKeys[0].RefTable = EntitlementsTable
	NotificationEventsTable.ForeignKeys[0].RefTable = NotificationRulesTable
	UsageResetsTable.ForeignKeys[0].RefTable = EntitlementsTable
	NotificationChannelRulesTable.ForeignKeys[0].RefTable = NotificationChannelsTable
	NotificationChannelRulesTable.ForeignKeys[1].RefTable = NotificationRulesTable
	NotificationEventDeliveryStatusEventsTable.ForeignKeys[0].RefTable = NotificationEventDeliveryStatusTable
	NotificationEventDeliveryStatusEventsTable.ForeignKeys[1].RefTable = NotificationEventsTable
}
