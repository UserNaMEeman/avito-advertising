1) Get list products
Post request with next fields: (path: /getproducts)
{
    "price_sort": false,  //sort by price (true)
    "date_sort": true,    //sort by date (true)
    "up": false           //true: from old to new; from small price to high
}

2) Get product by id
Post request with next fields: (path: /getproduct)
{
    id
    descr bool
    photo bool
}
3) Create product
Post request with next fields: (path: /newproduct)
{
    title <= 200 symbols
    description <= 1000 symbols
    url1
    url2
    url3
    price
}