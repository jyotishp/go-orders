package storage

type CustomerOrders struct {
    CustomerId int32
    OrderId int32
}

type OrderRestaurants struct {
    OrderId int32
    RestaurantId int32
}

type CuisineCustomers struct {
    Cuisine string
    CustomerId int32
}

type CuisineState struct {
    Cuisine string
    State string
}