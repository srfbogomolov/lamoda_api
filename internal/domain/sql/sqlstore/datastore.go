package sqlstore

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"

	"github.com/jmoiron/sqlx"
)

const (
	ProductTable      = "products"
	ProductIdColumn   = "id"
	ProductCodeColumn = "code"
	ProductNameColumn = "name"
	ProductSizeColumn = "size"

	WarehouseTable             = "warehouses"
	WarehouseIdColumn          = "id"
	WarehouseNameColumn        = "name"
	WarehouseIsAvailableColumn = "is_available"

	PlacementTable             = "placements"
	PlacementIdColumn          = "id"
	PlacementWarehouseIdColumn = "warehouse_id"
	PlacementProductIdColumn   = "product_id"
	PlacementQTYColumn         = "qty"

	ReservationTable             = "reservations"
	ReservationIdColumn          = "id"
	ReservationPlacementIdColumn = "placement_id"
	ReservationQTYColumn         = "qty"
)

var (
	dialect = goqu.Dialect("postgres")

	prodT          = goqu.T(ProductTable)
	prodIdCol      = prodT.Col(ProductIdColumn)
	prodCodeCol    = prodT.Col(ProductCodeColumn)
	prodNameCol    = prodT.Col(ProductNameColumn)
	prodSizeCol    = prodT.Col(ProductSizeColumn)
	prodPreparedDs = dialect.From(prodT).Prepared(true)

	wrhsT              = goqu.T(WarehouseTable)
	wrhsIdCol          = wrhsT.Col(WarehouseIdColumn)
	wrhsNameCol        = wrhsT.Col(WarehouseNameColumn)
	wrhsIsAvailableCol = wrhsT.Col(WarehouseIsAvailableColumn)
	wrhsPreparedDs     = dialect.From(wrhsT).Prepared(true)

	plcT              = goqu.T(PlacementTable)
	plcIdCol          = plcT.Col(PlacementIdColumn)
	plcWarehouseIdCol = plcT.Col(PlacementWarehouseIdColumn)
	plcProductIdCol   = plcT.Col(PlacementProductIdColumn)
	plcQTYCol         = plcT.Col(PlacementQTYColumn)
	plcPreparedDs     = dialect.From(plcT).Prepared(true)

	rsvT              = goqu.T(ReservationTable)
	rsvIdCol          = rsvT.Col(ReservationIdColumn)
	rsvPlacementIdCol = rsvT.Col(ReservationPlacementIdColumn)
	rsvQTYCol         = rsvT.Col(ReservationQTYColumn)
	rsvPreparedDs     = dialect.From(rsvT).Prepared(true)
)

type SqlxDatabase interface {
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}
