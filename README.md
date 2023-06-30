# TesIndochat


First of All You Need to Start XAMPP and start Apache and MySQL services

and then create a new database "indochat" without quote and then you can start the main.go in terminal "go run main.go"

(I Use Postman here)
first off all you can register to endpoint localhost:5000/api/v1/register with name, email, password
then you can login with email and passowrd you register before there is token copy the token for authrization

and then you can create category (input the bearer token from token before) and create category woth json name

and then you can create Product dont forget to input bearer token first of all POST to endpoint localhost:5000/api/v1/product you can  input
from form-data with :
name : string
desc : string
price : int
category_id : [] (with array)
image : you can upload file


for order use localhost:5000/api/v1/order endpoint with json
  "status": "pending",
  "product_id": 2,
  "discount_code": "IC003"


for check the result
localhost:5000/api/v1/products
localhost:5000/api/v1/orders
localhost:5000/api/v1/categories


For The Discount only the discount 10%, the code discount  i was wrong to place the func discount im so sorry😢
The CSV adn limiter script at file CSV adn Limiter.txt (not yet implemented 😢)

Thank you

