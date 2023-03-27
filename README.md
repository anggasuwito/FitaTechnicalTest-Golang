# FitaTechnicalTest-Golang

1. Setup your local configuration in file .env
2. Please migrate database structure and example data at directory "migration"
3. Run command ```go run main.go```
4. Example Request :
```
METHOD      GET
URL         127.0.0.1:8008/v1/shop/item-list
RESPONSE    
{
    "data": [
        {
            "id": "7e3e5be4-cc31-11ed-9138-00ffc2ef2645",
            "sku": "A304SD",
            "name": "Alexa Speaker",
            "price": 109.5,
            "quantity": 6
        },
        {
            "id": "7e3e036f-cc31-11ed-9138-00ffc2ef2645",
            "sku": "120P90",
            "name": "Google Home",
            "price": 49.99,
            "quantity": 6
        },
        {
            "id": "7e3e451e-cc31-11ed-9138-00ffc2ef2645",
            "sku": "43N23P",
            "name": "MacBook Pro",
            "price": 5399.99,
            "quantity": 9
        },
        {
            "id": "7e3e451e-cc31-11ed-9138-00ffc2ef2647",
            "sku": "234234",
            "name": "Raspberry Pi B",
            "price": 30,
            "quantity": 8
        }
    ],
    "err_response": null,
    "message": "success",
    "meta": null,
    "status": true
}
```

```
METHOD      POST
URL         127.0.0.1:8008/v1/shop/cart
BODY
{
    "user_id": "angga123",
    "item_sku": "A304SD",
    "quantity": 2
}
RESPONSE    
{
    "data": "Success add to your cart",
    "err_response": null,
    "message": "success",
    "meta": null,
    "status": true
}
```

```
METHOD      GET
URL         127.0.0.1:8008/v1/shop/cart/{userID}
            127.0.0.1:8008/v1/shop/cart/angga123
RESPONSE    
{
    "data": {
        "checkout_total_price": 5974.16,
        "items": [
            {
                "cart_id": "ee8ce8f5-cc92-11ed-9138-00ffc2ef2645",
                "user_id": "angga123",
                "item_sku": "A304SD",
                "item_name": "Alexa Speaker",
                "item_price": 109.5,
                "cart_quantity": 4
            },
            {
                "cart_id": "b9475547-cc91-11ed-9138-00ffc2ef2645",
                "user_id": "angga123",
                "item_sku": "120P90",
                "item_name": "Google Home",
                "item_price": 49.99,
                "cart_quantity": 4
            },
            {
                "cart_id": "6c1ba9aa-cc3b-11ed-9138-00ffc2ef2645",
                "user_id": "angga123",
                "item_sku": "43N23P",
                "item_name": "MacBook Pro",
                "item_price": 5399.99,
                "cart_quantity": 1
            },
            {
                "cart_id": "4f6ef3c4-cc3f-11ed-9138-00ffc2ef2645",
                "user_id": "angga123",
                "item_sku": "234234",
                "item_name": "Raspberry Pi B",
                "item_price": 30,
                "cart_quantity": 2
            }
        ]
    },
    "err_response": null,
    "message": "success",
    "meta": null,
    "status": true
}
```

```
METHOD      POST
NOTE        this API will run transaction query
            1. insert to transaction
            2. insert to transaction_history
            3. check availability item
            4. update quantity item
            5. delete all items on user current cart
URL         127.0.0.1:8008/v1/shop/checkout/{userID}
            127.0.0.1:8008/v1/shop/checkout/angga123
RESPONSE    
{
    "data": {
        "checkout_total_price": 5974.16,
        "items": [
            {
                "cart_id": "ee8ce8f5-cc92-11ed-9138-00ffc2ef2645",
                "user_id": "angga123",
                "item_sku": "A304SD",
                "item_name": "Alexa Speaker",
                "item_price": 109.5,
                "cart_quantity": 4
            },
            {
                "cart_id": "b9475547-cc91-11ed-9138-00ffc2ef2645",
                "user_id": "angga123",
                "item_sku": "120P90",
                "item_name": "Google Home",
                "item_price": 49.99,
                "cart_quantity": 4
            },
            {
                "cart_id": "6c1ba9aa-cc3b-11ed-9138-00ffc2ef2645",
                "user_id": "angga123",
                "item_sku": "43N23P",
                "item_name": "MacBook Pro",
                "item_price": 5399.99,
                "cart_quantity": 1
            },
            {
                "cart_id": "4f6ef3c4-cc3f-11ed-9138-00ffc2ef2645",
                "user_id": "angga123",
                "item_sku": "234234",
                "item_name": "Raspberry Pi B",
                "item_price": 30,
                "cart_quantity": 2
            }
        ]
    },
    "err_response": null,
    "message": "success",
    "meta": null,
    "status": true
}
```