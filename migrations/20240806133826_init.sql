-- Create "features" table
CREATE TABLE "features" ("id" character(26) NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NULL, "metadata" jsonb NULL, "namespace" character varying NOT NULL, "name" character varying NOT NULL, "key" character varying NOT NULL, "meter_slug" character varying NULL, "meter_group_by_filters" jsonb NULL, "archived_at" timestamptz NULL, PRIMARY KEY ("id"));
-- Create index "feature_id" to table: "features"
CREATE INDEX "feature_id" ON "features" ("id");
-- Create index "feature_namespace_id" to table: "features"
CREATE INDEX "feature_namespace_id" ON "features" ("namespace", "id");
-- Create "entitlements" table
CREATE TABLE "entitlements" ("id" character(26) NOT NULL, "namespace" character varying NOT NULL, "metadata" jsonb NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NULL, "entitlement_type" character varying NOT NULL, "feature_key" character varying NOT NULL, "subject_key" character varying NOT NULL, "measure_usage_from" timestamptz NULL, "issue_after_reset" double precision NULL, "issue_after_reset_priority" smallint NULL, "is_soft_limit" boolean NULL, "config" jsonb NULL, "usage_period_interval" character varying NULL, "usage_period_anchor" timestamptz NULL, "current_usage_period_start" timestamptz NULL, "current_usage_period_end" timestamptz NULL, "feature_id" character(26) NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "entitlements_features_entitlement" FOREIGN KEY ("feature_id") REFERENCES "features" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create index "entitlement_id" to table: "entitlements"
CREATE INDEX "entitlement_id" ON "entitlements" ("id");
-- Create index "entitlement_namespace_current_usage_period_end" to table: "entitlements"
CREATE INDEX "entitlement_namespace_current_usage_period_end" ON "entitlements" ("namespace", "current_usage_period_end");
-- Create index "entitlement_namespace_feature_id_id" to table: "entitlements"
CREATE INDEX "entitlement_namespace_feature_id_id" ON "entitlements" ("namespace", "feature_id", "id");
-- Create index "entitlement_namespace_id" to table: "entitlements"
CREATE INDEX "entitlement_namespace_id" ON "entitlements" ("namespace", "id");
-- Create index "entitlement_namespace_id_subject_key" to table: "entitlements"
CREATE INDEX "entitlement_namespace_id_subject_key" ON "entitlements" ("namespace", "id", "subject_key");
-- Create index "entitlement_namespace_subject_key" to table: "entitlements"
CREATE INDEX "entitlement_namespace_subject_key" ON "entitlements" ("namespace", "subject_key");
-- Create "balance_snapshots" table
CREATE TABLE "balance_snapshots" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "namespace" character varying NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NULL, "grant_balances" jsonb NOT NULL, "balance" numeric NOT NULL, "overage" numeric NOT NULL, "at" timestamptz NOT NULL, "owner_id" character(26) NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "balance_snapshots_entitlements_balance_snapshot" FOREIGN KEY ("owner_id") REFERENCES "entitlements" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create index "balancesnapshot_namespace_at" to table: "balance_snapshots"
CREATE INDEX "balancesnapshot_namespace_at" ON "balance_snapshots" ("namespace", "at");
-- Create index "balancesnapshot_namespace_balance" to table: "balance_snapshots"
CREATE INDEX "balancesnapshot_namespace_balance" ON "balance_snapshots" ("namespace", "balance");
-- Create index "balancesnapshot_namespace_balance_at" to table: "balance_snapshots"
CREATE INDEX "balancesnapshot_namespace_balance_at" ON "balance_snapshots" ("namespace", "balance", "at");
-- Create "grants" table
CREATE TABLE "grants" ("id" character(26) NOT NULL, "namespace" character varying NOT NULL, "metadata" jsonb NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NULL, "amount" numeric NOT NULL, "priority" smallint NOT NULL DEFAULT 0, "effective_at" timestamptz NOT NULL, "expiration" jsonb NOT NULL, "expires_at" timestamptz NOT NULL, "voided_at" timestamptz NULL, "reset_max_rollover" numeric NOT NULL, "reset_min_rollover" numeric NOT NULL, "recurrence_period" character varying NULL, "recurrence_anchor" timestamptz NULL, "owner_id" character(26) NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "grants_entitlements_grant" FOREIGN KEY ("owner_id") REFERENCES "entitlements" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create index "grant_effective_at_expires_at" to table: "grants"
CREATE INDEX "grant_effective_at_expires_at" ON "grants" ("effective_at", "expires_at");
-- Create index "grant_id" to table: "grants"
CREATE INDEX "grant_id" ON "grants" ("id");
-- Create index "grant_namespace_owner_id" to table: "grants"
CREATE INDEX "grant_namespace_owner_id" ON "grants" ("namespace", "owner_id");
-- Create "usage_resets" table
CREATE TABLE "usage_resets" ("id" character(26) NOT NULL, "namespace" character varying NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NULL, "reset_time" timestamptz NOT NULL, "entitlement_id" character(26) NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "usage_resets_entitlements_usage_reset" FOREIGN KEY ("entitlement_id") REFERENCES "entitlements" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create index "usagereset_id" to table: "usage_resets"
CREATE INDEX "usagereset_id" ON "usage_resets" ("id");
-- Create index "usagereset_namespace_entitlement_id" to table: "usage_resets"
CREATE INDEX "usagereset_namespace_entitlement_id" ON "usage_resets" ("namespace", "entitlement_id");
-- Create index "usagereset_namespace_entitlement_id_reset_time" to table: "usage_resets"
CREATE INDEX "usagereset_namespace_entitlement_id_reset_time" ON "usage_resets" ("namespace", "entitlement_id", "reset_time");
