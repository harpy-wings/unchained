// Code generated by ent, DO NOT EDIT.

package signer

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the signer type in the database.
	Label = "signer"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldEvm holds the string denoting the evm field in the database.
	FieldEvm = "evm"
	// FieldKey holds the string denoting the key field in the database.
	FieldKey = "key"
	// FieldShortkey holds the string denoting the shortkey field in the database.
	FieldShortkey = "shortkey"
	// FieldPoints holds the string denoting the points field in the database.
	FieldPoints = "points"
	// EdgeAssetPrice holds the string denoting the assetprice edge name in mutations.
	EdgeAssetPrice = "assetPrice"
	// EdgeEventLogs holds the string denoting the eventlogs edge name in mutations.
	EdgeEventLogs = "eventLogs"
	// Table holds the table name of the signer in the database.
	Table = "signers"
	// AssetPriceTable is the table that holds the assetPrice relation/edge. The primary key declared below.
	AssetPriceTable = "asset_price_signers"
	// AssetPriceInverseTable is the table name for the AssetPrice entity.
	// It exists in this package in order to avoid circular dependency with the "assetprice" package.
	AssetPriceInverseTable = "asset_prices"
	// EventLogsTable is the table that holds the eventLogs relation/edge. The primary key declared below.
	EventLogsTable = "event_log_signers"
	// EventLogsInverseTable is the table name for the EventLog entity.
	// It exists in this package in order to avoid circular dependency with the "eventlog" package.
	EventLogsInverseTable = "event_logs"
)

// Columns holds all SQL columns for signer fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldEvm,
	FieldKey,
	FieldShortkey,
	FieldPoints,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "signers"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"correctness_report_signers",
}

var (
	// AssetPricePrimaryKey and AssetPriceColumn2 are the table columns denoting the
	// primary key for the assetPrice relation (M2M).
	AssetPricePrimaryKey = []string{"asset_price_id", "signer_id"}
	// EventLogsPrimaryKey and EventLogsColumn2 are the table columns denoting the
	// primary key for the eventLogs relation (M2M).
	EventLogsPrimaryKey = []string{"event_log_id", "signer_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// KeyValidator is a validator for the "key" field. It is called by the builders before save.
	KeyValidator func([]byte) error
	// ShortkeyValidator is a validator for the "shortkey" field. It is called by the builders before save.
	ShortkeyValidator func([]byte) error
)

// OrderOption defines the ordering options for the Signer queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByEvm orders the results by the evm field.
func ByEvm(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEvm, opts...).ToFunc()
}

// ByPoints orders the results by the points field.
func ByPoints(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPoints, opts...).ToFunc()
}

// ByAssetPriceCount orders the results by assetPrice count.
func ByAssetPriceCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAssetPriceStep(), opts...)
	}
}

// ByAssetPrice orders the results by assetPrice terms.
func ByAssetPrice(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAssetPriceStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByEventLogsCount orders the results by eventLogs count.
func ByEventLogsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newEventLogsStep(), opts...)
	}
}

// ByEventLogs orders the results by eventLogs terms.
func ByEventLogs(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newEventLogsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newAssetPriceStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AssetPriceInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, AssetPriceTable, AssetPricePrimaryKey...),
	)
}
func newEventLogsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EventLogsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, EventLogsTable, EventLogsPrimaryKey...),
	)
}
