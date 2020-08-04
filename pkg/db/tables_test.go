package db_test

import (
    "github.com/jyotishp/go-orders/pkg/db"
    "testing"
)

func TestOrdersTable(t *testing.T) {
    table := db.OrdersTableSchema()
    assertTableName(t, *table.TableName, db.OrdersTable)
}

func assertTableName(t *testing.T, a, b string) {
    if a != b {
        t.Errorf("table name doesn't match")
    }
}

func TestRestaurantsTable(t *testing.T) {
    table := db.RestaurantsTableSchema()
    assertTableName(t, *table.TableName, db.RestaurantsTable)
}

func TestCustomersTable(t *testing.T) {
    table := db.CustomersTableSchema()
    assertTableName(t, *table.TableName, db.CustomersTable)
}

func TestItemsTable(t *testing.T) {
    table := db.ItemsTableSchema()
    assertTableName(t, *table.TableName, "Items")
}

func TestCreateTable(t *testing.T) {
    svc := &mockDynamoDBClient{}
    err := db.CreateTable(svc, db.OrdersTableSchema())
    if err != nil {
        t.Errorf("unable to create new table: %v", err)
    }
}

func TestTableExists(t *testing.T) {
    svc := &mockDynamoDBClient{}
    db.CreateTable(svc, db.OrdersTableSchema())
    res := db.TableExists(svc, db.OrdersTable)
    if !res {
        t.Errorf("table exists but returns doesn't exist")
    }
}
