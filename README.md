# Introduction
As a marketplace, we need to pay our sellers for every item that has been sold on our platform. In this task, you’ll be working with 2 main entities: Items to sell (products on the website) and Payouts instructions to send (bank transactions to seller accounts)
API call that accepts a list of sold Items and creates Payouts for the sellers.


# Docker:
This project uses mysql db to store the created payouts
# docker-compose up -d --build

mysql has a connection string to myDb database. Modify it accordingly.

# Request Body:

POST: /createPayout
Body:
[
{
"name": "Top",
"price_amount":10000,
"price_currency":"USD",
"seller_reference": 1
},
{
"name": "Bottom",
"price_amount":1000,
"price_currency":"USD",
"seller_reference": 2
}
]

# Response:
Status: 200 StatusOk
Body:
{
"no_of_transactions": 6,
"payouts": [
{
"seller_reference": 1,
"amount": 10000,
"currency": "USD"
},
{
"seller_reference": 2,
"amount": 1000,
"currency": "USD"
}
]
}
