package storage

import (
    pb "github.com/jyotishp/go-orders/pkg/proto"
    "os"
)

type DataHandler struct {
    OutDir             string
    OrdWriter          *JsonWriter
    Customers          map[int32]*pb.Customer
    Restaurants        map[int32]*pb.Restaurant
    CuisineStateWriter *JsonWriter
    CustOrdWriter      *JsonWriter
    OrdRestWriter      *JsonWriter
    CuisineCustWriter  *JsonWriter
}

func NewDataHandler(outDir string) *DataHandler {
    if _, err := os.Stat(outDir); os.IsNotExist(err) {
        os.Mkdir(outDir, 0755)
    }
    return &DataHandler{
        OutDir:             outDir,
        OrdWriter:          NewJsonWriter(outDir+"/orders.json", "orders"),
        Customers:          make(map[int32]*pb.Customer),
        Restaurants:        make(map[int32]*pb.Restaurant),
        CuisineStateWriter: NewJsonWriter(outDir+"/cuisine_state.json", "cuisine_state"),
        CustOrdWriter:      NewJsonWriter(outDir+"/customer_orders.json", "customer_orders"),
        OrdRestWriter:      NewJsonWriter(outDir+"/order_rests.json", "order_rests"),
        CuisineCustWriter:  NewJsonWriter(outDir+"/cuisine_custs.json", "cuisine_custs"),
    }
}

func (dh *DataHandler) Init(data []string) {
    order, custOrd, ordRest, cuisineState, cuisineCust := dh.ProcessData(data)
    dh.OrdWriter.InitWriteResource(order)
    dh.CuisineStateWriter.InitWriteResource(cuisineState)
    dh.CustOrdWriter.InitWriteResource(custOrd)
    dh.OrdRestWriter.InitWriteResource(ordRest)
    dh.CuisineCustWriter.InitWriteResource(cuisineCust)
}

func (dh *DataHandler) Write(data []string) {
    if len(data) > 0 {
        order, custOrd, ordRest, cuisineState, cuisineCust := dh.ProcessData(data)
        dh.OrdWriter.WriteResource(order)
        dh.CuisineStateWriter.WriteResource(cuisineState)
        dh.CustOrdWriter.WriteResource(custOrd)
        dh.OrdRestWriter.WriteResource(ordRest)
        dh.CuisineCustWriter.WriteResource(cuisineCust)
    }
}

func (dh *DataHandler) Close() {
    dh.OrdWriter.Close()
    dh.CuisineCustWriter.Close()
    dh.CuisineStateWriter.Close()
    dh.CustOrdWriter.Close()
    dh.OrdRestWriter.Close()
}

func (dh *DataHandler) ProcessData(data []string) (*pb.Order, CustomerOrders, OrderRestaurants, CuisineState, CuisineCustomers) {
    order := NewOrder(data)
    customer := NewCustomer(data)
    restaurant := NewRestaurant(data)
    dh.Customers[customer.Id] = customer
    dh.Restaurants[restaurant.Id] = restaurant
    custOrd := CustomerOrders{
        CustomerId: customer.Id,
        OrderId:    order.Id,
    }
    ordRest := OrderRestaurants{
        OrderId:      order.Id,
        RestaurantId: restaurant.Id,
    }
    cuisineState := CuisineState{
        Cuisine: order.Cuisine,
        State:   customer.State,
    }
    cuisineCust := CuisineCustomers{
        Cuisine:    order.Cuisine,
        CustomerId: customer.Id,
    }
    return order, custOrd, ordRest, cuisineState, cuisineCust
}

