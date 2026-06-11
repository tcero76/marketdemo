import scrapy

from marketdemo.itemsDemo import ProductItem, ProductImageItem

class DummySpider(scrapy.Spider):
    name = "dummy"
    allowed_domains = ["dummyjson.com"]
    start_urls = ["https://dummyjson.com/products"]

    def parse(self, response):
        response_json = response.json()
        item = ProductItem()
        for product in response_json["products"]:
            item["id"] = product["id"]
            item["title"] = product["title"]
            item["description"] = product["description"]
            item["price"] = product["price"]
            item["discount_percentage"] = product["discountPercentage"]
            item["rating"] = product["rating"]
            item["stock"] = product["stock"]
            item["category"] = product["category"]
            item["thumbnail"] = product["thumbnail"]
            item["brand"] = product.get("brand")
            yield item
            for image in product["images"]:
                image_item = ProductImageItem()
                image_item["product_id"] = product["id"]
                image_item["url"] = image
                yield image_item
