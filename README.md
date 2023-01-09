1) Get list products
Post request with next fields:
{
    "price_sort": false,  //sort by price (true)
    "date_sort": true,    //sort by date (true)
    "up": false           //true: from old to new; from small price to high
}

3) Create product
Post request with next fields:
{
    title <= 200 symbols
    description <= 1000 symbols
    url1
    url2
    url3
    price
}