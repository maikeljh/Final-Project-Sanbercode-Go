# Final-Project-Sanbercode-Go-Batch-41
## Dokumentasi API

### - Login (POST)
    https://final-project-sanbercode-go-production.up.railway.app/login
**Body**
- username : string
- password : string

Description : Melakukan proses autentikasi login

### - Logout (POST)
    https://final-project-sanbercode-go-production.up.railway.app/logout

Description : Melakukan proses logout

### - Register (POST)
    https://final-project-sanbercode-go-production.up.railway.app/users
**Body**
- name : string
- address : string
- phone_number : string
- username : string
- password : string

Description : Melakukan proses register

### - Get Categories
    https://final-project-sanbercode-go-production.up.railway.app/categories
**Authorization**
- Basic Auth (Admin) : Username and Password
- Session (Login) : Username and Password

Description : Mengembalikan list of category

### - Post Category
    https://final-project-sanbercode-go-production.up.railway.app/categories
**Authorization**
- Basic Auth (Admin) : Username and Password

**Body**
- name : string

Description : Menambahkan category

### - Put Category
    https://final-project-sanbercode-go-production.up.railway.app/categories/{{id}}
**Authorization**
- Basic Auth (Admin) : Username and Password

**Body**
- name : string

Description : Mengupdate category berdasarkan id

### - Delete Category
    https://final-project-sanbercode-go-production.up.railway.app/categories/{{id}}
**Authorization**
- Basic Auth (Admin) : Username and Password

Description : Men-delete category berdasarkan id

### - Get Products By Category ID
    https://final-project-sanbercode-go-production.up.railway.app/categories/{{category_id}}/products
**Authorization**
- Basic Auth (Admin) : Username and Password
- Session (Login) : Username and Password

Description : Mengembalikan list of product berdasarkan category_id

### - Get Users
    https://final-project-sanbercode-go-production.up.railway.app/users
**Authorization**
- Basic Auth (Admin) : Username and Password

Description : Mengembalikan list of user

### - Put User
    https://final-project-sanbercode-go-production.up.railway.app/users/{{id}}
**Authorization**
- Basic Auth : Username and Password
- Session (Login) : Username and Password

**Body**
- name : string
- address : string
- phone_number : string

Description : Mengupdate user berdasarkan id

### - Delete User
    https://final-project-sanbercode-go-production.up.railway.app/users/{{id}}
**Authorization**
- Basic Auth : Username and Password
- Session (Login) : Username and Password

Description : Men-delete user berdasarkan id

### - Get Products
    https://final-project-sanbercode-go-production.up.railway.app/products
**Authorization**
- Basic Auth (Admin) : Username and Password
- Session (Login) : Username and Password

Description : Mengembalikan list of product

### - Post Product
    https://final-project-sanbercode-go-production.up.railway.app/products
**Authorization**
- Basic Auth (Admin) : Username and Password

**Body**
- name : string
- category_id : int
- price : int
- description : string

Description : Melakukan penambahan product

### - Put Product
    https://final-project-sanbercode-go-production.up.railway.app/products/{{id}}
**Authorization**
- Basic Auth : Username and Password

**Body**
- name : string
- category_id : int
- price : int
- description : string

Description : Mengupdate product berdasarkan id

### - Delete Product
    https://final-project-sanbercode-go-production.up.railway.app/products/{{id}}
**Authorization**
- Basic Auth : Username and Password

Description : Men-delete product berdasarkan id

### - Get Carts
    https://final-project-sanbercode-go-production.up.railway.app/carts
**Authorization**
- Basic Auth (Admin) : Username and Password
- Session (Login) : Username and Password

Description : Mengembalikan list of cart

### - Post Cart
    https://final-project-sanbercode-go-production.up.railway.app/carts
**Authorization**
- Basic Auth (Admin) : Username and Password
- Session (Login) : Username and Password

**Body**
- product_id : int
- count : int
- user_id : int

Description : Melakukan penambahan cart

### - Put Cart
    https://final-project-sanbercode-go-production.up.railway.app/carts/{{id}}
**Authorization**
- Basic Auth : Username and Password
- Session (Login) : Username and Password

**Body**
- product_id : int
- count : int
- user_id : int

Description : Mengupdate cart berdasarkan id

### - Delete Cart
    https://final-project-sanbercode-go-production.up.railway.app/carts/{{id}}
**Authorization**
- Basic Auth : Username and Password
- Session (Login) : Username and Password

Description : Men-delete cart berdasarkan id

### - Get Carts By User ID
    https://final-project-sanbercode-go-production.up.railway.app/users/{{user_id}}/carts
**Authorization**
- Basic Auth (Admin) : Username and Password
- Session (Login) : Username and Password

Description : Mengembalikan list of cart berdasarkan user_id

### - Get Orders
    https://final-project-sanbercode-go-production.up.railway.app/orders
**Authorization**
- Basic Auth (Admin) : Username and Password
- Session (Login) : Username and Password

Description : Mengembalikan list of order

### - Post Order
    https://final-project-sanbercode-go-production.up.railway.app/orders
**Authorization**
- Basic Auth (Admin) : Username and Password
- Session (Login) : Username and Password

**Body**
- product_id : int
- count : int
- user_id : int
- cart_id : int

Description : Melakukan penambahan order

### - Put Order
    https://final-project-sanbercode-go-production.up.railway.app/orders/{{id}}
**Authorization**
- Basic Auth : Username and Password

**Body**
- product_id : int
- count : int
- user_id : int
- cart_id : int

Description : Mengupdate order berdasarkan id

### - Delete Order
    https://final-project-sanbercode-go-production.up.railway.app/orders/{{id}}
**Authorization**
- Basic Auth : Username and Password

Description : Men-delete order berdasarkan id

### - Get Orders By User ID
    https://final-project-sanbercode-go-production.up.railway.app/users/{{user_id}}/orders
**Authorization**
- Basic Auth (Admin) : Username and Password
- Session (Login) : Username and Password

Description : Mengembalikan list of order berdasarkan user_id
