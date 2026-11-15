import scrapy

class ProductItem(scrapy.Item):
    __table__ = "scrap.products_stage"
    id = scrapy.Field()
    sku = scrapy.Field()
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
    __table__ = "scrap.product_images_stage"
    id = scrapy.Field()
    product_id = scrapy.Field()
    url = scrapy.Field()