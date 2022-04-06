# pizza-ordering-app

Congrats, you are the proud owner of a pizza business. Your business has a very simple
workflow:
- You serve only 2 kinds of pizza (Veggie Lovers and Meat Lovers). You serve only one
size of pizza.
- After you receive an order, you “make” the pizza and then you notify customers that their
pizza is ready for pick up. Customers can be notified either through their email or text
message.

## Project Structure
```
.
├── Dockerfile
├── LICENSE
├── README.md
├── docker-compose.yml
├── go.mod
├── go.sum
├── internal
│   ├── api
│   │   ├── api_helper.go
│   │   ├── api_oder_pizza.go
│   │   ├── api_oder_pizza_test.go
│   │   ├── api_status.go
│   │   ├── model.go
│   │   ├── router.go
│   │   └── server.go
│   ├── cmd
│   │   ├── api.go
│   │   ├── cfg_helper.go
│   │   └── root.go
│   ├── data
│   │   ├── dal_helper.go
│   │   ├── dal_order_pizza.go
│   │   ├── dal_order_pizza_test.go
│   │   ├── migration
│   │   │   └── initial_migration.sql
│   │   ├── mock_db.go
│   │   ├── models.go
│   │   └── repo.go
│   └── sms
│       └── sms.go
├── main.go
├── makefile
└── wait-for-it.sh
```

## Dependencies
> Docker

## Steps to Execute 

#### 1. goto app root directory 
```bash
cd pizza-app
```
#### 2. start the application
```bash
make up
```


### API for status checking [GET]

> http://localhost:3000/api/status
```bash
Sample response body
{
    "status": "ok",
    "result": "postgres db is working perfectly"
}
```
### Add a new order [POST]

> http://localhost:3000/api/buy_pizza

Sample request body
```bash

{
    "user_id": 1,
    "pizza_id": 2,
    "pizza_size": "medium"
}
```
Sample response body
```bash

{
    "status": "ok",
    "result": {
        "pizza_id": 2,
        "pizza_size": "medium",
        "cooking_stage": "start",
        "user_id": 1,
        "start_time": "2021-10-17T12:32:51.680207885Z",
        "is_active": true
    }
}
```

### Update an order [PUT]

> http://localhost:3000/api/order

Sample request body
```bash

{
    "id": 1,
    "cooking_stage": "done"
}
```
Sample response body
```bash

{
    "status": "ok",
    "result": null
}
```

### View order status [GET]

> http://localhost:3000/api/order-status/{user_id}

Sample response body
```bash

{
    "status": "ok",
    "result": {
        "pizza_id": 2,
        "pizza_size": "medium",
        "cooking_stage": "start",
        "user_id": 1,
        "start_time": "2021-10-17T12:32:51.680207885Z",
        "is_active": true
    }
}
```

# To do
- Create API to delete the order
- Prevent updation of deleted orders
- Create API to add new pizza types
- Create API to list all pizza types
- create api to list all orders for the last x days
