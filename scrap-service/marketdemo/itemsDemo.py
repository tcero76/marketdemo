import scrapy

class ProductItem(scrapy.Item):
    __table__ = "marketplacedemo.products_stage"
    id = scrapy.Field()
    title = scrapy.Field()
    description = scrapy.Field()
    price = scrapy.Field()
    discount_percentage = scrapy.Field()
    rating = scrapy.Field()
    stock = scrapy.Field()
    brand = scrapy.Field()
    category = scrapy.Field()
    thumbnail = scrapy.Field()

class ProductImageItem(scrapy.Item):
    __table__ = "marketplacedemo.product_images_stage"
    id = scrapy.Field()
    product_id = scrapy.Field()
    url = scrapy.Field()