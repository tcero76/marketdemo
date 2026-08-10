--
-- Data for Name: category; Type: TABLE DATA; Schema: marketplacedemo; Owner: tcero
--

ALTER TABLE marketplacedemo.category DISABLE TRIGGER ALL;

INSERT INTO marketplacedemo.category (id, name) VALUES
	(1, 'beauty'),
	(2, 'groceries'),
	(3, 'furniture'),
	(4, 'fragrances');


ALTER TABLE marketplacedemo.category ENABLE TRIGGER ALL;

--
-- Data for Name: products; Type: TABLE DATA; Schema: marketplacedemo; Owner: tcero
--

ALTER TABLE marketplacedemo.products DISABLE TRIGGER ALL;

INSERT INTO marketplacedemo.products (id, title, description, price, discount_percentage, rating, stock, brand, categoryid, thumbnail, search_vector, deleted_at) VALUES
	(1, 'Essence Mascara Lash Princess', 'The Essence Mascara Lash Princess is a popular mascara known for its volumizing and lengthening effects. Achieve dramatic lashes with this long-lasting and cruelty-free formula.', 9.99, 10.48, 2.56, 99, 'Essence', 1, 'https://cdn.dummyjson.com/product-images/beauty/essence-mascara-lash-princess/thumbnail.webp', '''achiev'':21B ''beauti'':35C ''cruelti'':31B ''cruelty-fre'':30B ''dramat'':22B ''effect'':20B ''essenc'':1A,6B,34C ''formula'':33B ''free'':32B ''known'':14B ''lash'':3A,8B,23B ''last'':28B ''lengthen'':19B ''long'':27B ''long-last'':26B ''mascara'':2A,7B,13B ''popular'':12B ''princess'':4A,9B ''volum'':17B', NULL),
	(23, 'Eggs', 'Fresh eggs, a versatile ingredient for baking, cooking, or breakfast.', 2.99, 11.05, 2.53, 9, NULL, 2, 'https://cdn.dummyjson.com/product-images/groceries/eggs/thumbnail.webp', '''bake'':8B ''breakfast'':11B ''cook'':9B ''egg'':1A,3B ''fresh'':2B ''groceri'':12C ''ingredi'':6B ''versatil'':5B', NULL),
	(2, 'Eyeshadow Palette with Mirror', 'The Eyeshadow Palette with Mirror offers a versatile range of eyeshadow shades for creating stunning eye looks. With a built-in mirror, it''s convenient for on-the-go makeup application.', 19.99, 18.19, 2.86, 34, 'Glamour Beauty', 1, 'https://cdn.dummyjson.com/product-images/beauty/eyeshadow-palette-with-mirror/thumbnail.webp', '''applic'':37B ''beauti'':39C,40C ''built'':25B ''built-in'':24B ''conveni'':30B ''creat'':18B ''eye'':20B ''eyeshadow'':1A,6B,15B ''glamour'':38C ''go'':35B ''look'':21B ''makeup'':36B ''mirror'':4A,9B,27B ''offer'':10B ''on-the-go'':32B ''palett'':2A,7B ''rang'':13B ''shade'':16B ''stun'':19B ''versatil'':12B', NULL),
	(3, 'Powder Canister', 'The Powder Canister is a finely milled setting powder designed to set makeup and control shine. With a lightweight and translucent formula, it provides a smooth and matte finish.', 14.99, 9.84, 4.64, 89, 'Velvet Touch', 1, 'https://cdn.dummyjson.com/product-images/beauty/powder-canister/thumbnail.webp', '''beauti'':34C ''canist'':2A,5B ''control'':17B ''design'':12B ''fine'':8B ''finish'':31B ''formula'':24B ''lightweight'':21B ''makeup'':15B ''matt'':30B ''mill'':9B ''powder'':1A,4B,11B ''provid'':26B ''set'':10B,14B ''shine'':18B ''smooth'':28B ''touch'':33C ''transluc'':23B ''velvet'':32C', NULL),
	(4, 'Red Lipstick', 'The Red Lipstick is a classic and bold choice for adding a pop of color to your lips. With a creamy and pigmented formula, it provides a vibrant and long-lasting finish.', 12.99, 12.16, 4.36, 91, 'Chic Cosmetics', 1, 'https://cdn.dummyjson.com/product-images/beauty/red-lipstick/thumbnail.webp', '''ad'':13B ''beauti'':38C ''bold'':10B ''chic'':36C ''choic'':11B ''classic'':8B ''color'':17B ''cosmet'':37C ''creami'':23B ''finish'':35B ''formula'':26B ''last'':34B ''lip'':20B ''lipstick'':2A,5B ''long'':33B ''long-last'':32B ''pigment'':25B ''pop'':15B ''provid'':28B ''red'':1A,4B ''vibrant'':30B', NULL),
	(5, 'Red Nail Polish', 'The Red Nail Polish offers a rich and glossy red hue for vibrant and polished nails. With a quick-drying formula, it provides a salon-quality finish at home.', 8.99, 11.44, 4.32, 79, 'Nail Couture', 1, 'https://cdn.dummyjson.com/product-images/beauty/red-nail-polish/thumbnail.webp', '''beauti'':37C ''coutur'':36C ''dri'':24B ''finish'':32B ''formula'':25B ''glossi'':12B ''home'':34B ''hue'':14B ''nail'':2A,6B,19B,35C ''offer'':8B ''polish'':3A,7B,18B ''provid'':27B ''qualiti'':31B ''quick'':23B ''quick-dri'':22B ''red'':1A,5B,13B ''rich'':10B ''salon'':30B ''salon-qu'':29B ''vibrant'':16B', NULL),
	(6, 'Calvin Klein CK One', 'CK One by Calvin Klein is a classic unisex fragrance, known for its fresh and clean scent. It''s a versatile fragrance suitable for everyday wear.', 49.99, 1.89, 4.37, 29, 'Calvin Klein', 4, 'https://cdn.dummyjson.com/product-images/fragrances/calvin-klein-ck-one/thumbnail.webp', '''calvin'':1A,8B,31C ''ck'':3A,5B ''classic'':12B ''clean'':20B ''everyday'':29B ''fragranc'':14B,26B,33C ''fresh'':18B ''klein'':2A,9B,32C ''known'':15B ''one'':4A,6B ''scent'':21B ''suitabl'':27B ''unisex'':13B ''versatil'':25B ''wear'':30B', NULL),
	(7, 'Chanel Coco Noir Eau De', 'Coco Noir by Chanel is an elegant and mysterious fragrance, featuring notes of grapefruit, rose, and sandalwood. Perfect for evening occasions.', 129.99, 16.51, 4.26, 58, 'Chanel', 4, 'https://cdn.dummyjson.com/product-images/fragrances/chanel-coco-noir-eau-de/thumbnail.webp', '''chanel'':1A,9B,27C ''coco'':2A,6B ''de'':5A ''eau'':4A ''eleg'':12B ''even'':25B ''featur'':16B ''fragranc'':15B,28C ''grapefruit'':19B ''mysteri'':14B ''noir'':3A,7B ''note'':17B ''occas'':26B ''perfect'':23B ''rose'':20B ''sandalwood'':22B', NULL),
	(8, 'Dior J''adore', 'J''adore by Dior is a luxurious and floral fragrance, known for its blend of ylang-ylang, rose, and jasmine. It embodies femininity and sophistication.', 89.99, 14.72, 3.8, 98, 'Dior', 4, 'https://cdn.dummyjson.com/product-images/fragrances/dior-j''adore/thumbnail.webp', '''ador'':3A,5B ''blend'':17B ''dior'':1A,7B,30C ''embodi'':26B ''feminin'':27B ''floral'':12B ''fragranc'':13B,31C ''j'':2A,4B ''jasmin'':24B ''known'':14B ''luxuri'':10B ''rose'':22B ''sophist'':29B ''ylang'':20B,21B ''ylang-ylang'':19B', NULL),
	(9, 'Dolce Shine Eau de', 'Dolce Shine by Dolce & Gabbana is a vibrant and fruity fragrance, featuring notes of mango, jasmine, and blonde woods. It''s a joyful and youthful scent.', 69.99, 0.62, 3.96, 4, 'Dolce & Gabbana', 4, 'https://cdn.dummyjson.com/product-images/fragrances/dolce-shine-eau-de/thumbnail.webp', '''blond'':22B ''de'':4A ''dolc'':1A,5B,8B,31C ''eau'':3A ''featur'':16B ''fragranc'':15B,33C ''fruiti'':14B ''gabbana'':9B,32C ''jasmin'':20B ''joy'':27B ''mango'':19B ''note'':17B ''scent'':30B ''shine'':2A,6B ''vibrant'':12B ''wood'':23B ''youth'':29B', NULL),
	(10, 'Gucci Bloom Eau de', 'Gucci Bloom by Gucci is a floral and captivating fragrance, with notes of tuberose, jasmine, and Rangoon creeper. It''s a modern and romantic scent.', 79.99, 14.39, 2.74, 91, 'Gucci', 4, 'https://cdn.dummyjson.com/product-images/fragrances/gucci-bloom-eau-de/thumbnail.webp', '''bloom'':2A,6B ''captiv'':13B ''creeper'':22B ''de'':4A ''eau'':3A ''floral'':11B ''fragranc'':14B,31C ''gucci'':1A,5B,8B,30C ''jasmin'':19B ''modern'':26B ''note'':16B ''rangoon'':21B ''romant'':28B ''scent'':29B ''tuberos'':18B', NULL),
	(11, 'Annibale Colombo Bed', 'The Annibale Colombo Bed is a luxurious and elegant bed frame, crafted with high-quality materials for a comfortable and stylish bedroom.', 1899.99, 8.57, 4.77, 88, 'Annibale Colombo', 3, 'https://cdn.dummyjson.com/product-images/furniture/annibale-colombo-bed/thumbnail.webp', '''annibal'':1A,5B,27C ''bed'':3A,7B,13B ''bedroom'':26B ''colombo'':2A,6B,28C ''comfort'':23B ''craft'':15B ''eleg'':12B ''frame'':14B ''furnitur'':29C ''high'':18B ''high-qual'':17B ''luxuri'':10B ''materi'':20B ''qualiti'':19B ''stylish'':25B', NULL),
	(12, 'Annibale Colombo Sofa', 'The Annibale Colombo Sofa is a sophisticated and comfortable seating option, featuring exquisite design and premium upholstery for your living room.', 2499.99, 14.4, 3.92, 60, 'Annibale Colombo', 3, 'https://cdn.dummyjson.com/product-images/furniture/annibale-colombo-sofa/thumbnail.webp', '''annibal'':1A,5B,25C ''colombo'':2A,6B,26C ''comfort'':12B ''design'':17B ''exquisit'':16B ''featur'':15B ''furnitur'':27C ''live'':23B ''option'':14B ''premium'':19B ''room'':24B ''seat'':13B ''sofa'':3A,7B ''sophist'':10B ''upholsteri'':20B', NULL),
	(13, 'Bedside Table African Cherry', 'The Bedside Table in African Cherry is a stylish and functional addition to your bedroom, providing convenient storage space and a touch of elegance.', 299.99, 19.09, 2.87, 64, 'Furniture Co.', 3, 'https://cdn.dummyjson.com/product-images/furniture/bedside-table-african-cherry/thumbnail.webp', '''addit'':16B ''african'':3A,9B ''bedroom'':19B ''bedsid'':1A,6B ''cherri'':4A,10B ''co'':30C ''conveni'':21B ''eleg'':28B ''function'':15B ''furnitur'':29C,31C ''provid'':20B ''space'':23B ''storag'':22B ''stylish'':13B ''tabl'':2A,7B ''touch'':26B', NULL),
	(14, 'Knoll Saarinen Executive Conference Chair', 'The Knoll Saarinen Executive Conference Chair is a modern and ergonomic chair, perfect for your office or conference room with its timeless design.', 499.99, 2.01, 4.88, 26, 'Knoll', 3, 'https://cdn.dummyjson.com/product-images/furniture/knoll-saarinen-executive-conference-chair/thumbnail.webp', '''chair'':5A,11B,17B ''confer'':4A,10B,23B ''design'':28B ''ergonom'':16B ''execut'':3A,9B ''furnitur'':30C ''knoll'':1A,7B,29C ''modern'':14B ''offic'':21B ''perfect'':18B ''room'':24B ''saarinen'':2A,8B ''timeless'':27B', NULL),
	(15, 'Wooden Bathroom Sink With Mirror', 'The Wooden Bathroom Sink with Mirror is a unique and stylish addition to your bathroom, featuring a wooden sink countertop and a matching mirror.', 799.99, 8.8, 3.59, 7, 'Bath Trends', 3, 'https://cdn.dummyjson.com/product-images/furniture/wooden-bathroom-sink-with-mirror/thumbnail.webp', '''addit'':17B ''bath'':30C ''bathroom'':2A,8B,20B ''countertop'':25B ''featur'':21B ''furnitur'':32C ''match'':28B ''mirror'':5A,11B,29B ''sink'':3A,9B,24B ''stylish'':16B ''trend'':31C ''uniqu'':14B ''wooden'':1A,7B,23B', NULL),
	(16, 'Apple', 'Fresh and crisp apples, perfect for snacking or incorporating into various recipes.', 1.99, 12.62, 4.19, 8, NULL, 2, 'https://cdn.dummyjson.com/product-images/groceries/apple/thumbnail.webp', '''appl'':1A,5B ''crisp'':4B ''fresh'':2B ''groceri'':14C ''incorpor'':10B ''perfect'':6B ''recip'':13B ''snack'':8B ''various'':12B', NULL),
	(17, 'Beef Steak', 'High-quality beef steak, great for grilling or cooking to your preferred level of doneness.', 12.99, 9.61, 4.47, 86, NULL, 2, 'https://cdn.dummyjson.com/product-images/groceries/beef-steak/thumbnail.webp', '''beef'':1A,6B ''cook'':12B ''done'':18B ''great'':8B ''grill'':10B ''groceri'':19C ''high'':4B ''high-qual'':3B ''level'':16B ''prefer'':15B ''qualiti'':5B ''steak'':2A,7B', NULL),
	(18, 'Cat Food', 'Nutritious cat food formulated to meet the dietary needs of your feline friend.', 8.99, 9.58, 3.13, 46, NULL, 2, 'https://cdn.dummyjson.com/product-images/groceries/cat-food/thumbnail.webp', '''cat'':1A,4B ''dietari'':10B ''felin'':14B ''food'':2A,5B ''formul'':6B ''friend'':15B ''groceri'':16C ''meet'':8B ''need'':11B ''nutriti'':3B', NULL),
	(19, 'Chicken Meat', 'Fresh and tender chicken meat, suitable for various culinary preparations.', 9.99, 13.7, 3.19, 97, NULL, 2, 'https://cdn.dummyjson.com/product-images/groceries/chicken-meat/thumbnail.webp', '''chicken'':1A,6B ''culinari'':11B ''fresh'':3B ''groceri'':13C ''meat'':2A,7B ''prepar'':12B ''suitabl'':8B ''tender'':5B ''various'':10B', NULL),
	(20, 'Cooking Oil', 'Versatile cooking oil suitable for frying, sautéing, and various culinary applications.', 4.99, 9.33, 4.8, 10, NULL, 2, 'https://cdn.dummyjson.com/product-images/groceries/cooking-oil/thumbnail.webp', '''applic'':13B ''cook'':1A,4B ''culinari'':12B ''fri'':8B ''groceri'':14C ''oil'':2A,5B ''sauté'':9B ''suitabl'':6B ''various'':11B ''versatil'':3B', NULL),
	(21, 'Cucumber', 'Crisp and hydrating cucumbers, ideal for salads, snacks, or as a refreshing side.', 1.49, 0.16, 4.07, 84, NULL, 2, 'https://cdn.dummyjson.com/product-images/groceries/cucumber/thumbnail.webp', '''crisp'':2B ''cucumb'':1A,5B ''groceri'':15C ''hydrat'':4B ''ideal'':6B ''refresh'':13B ''salad'':8B ''side'':14B ''snack'':9B', NULL),
	(22, 'Dog Food', 'Specially formulated dog food designed to provide essential nutrients for your canine companion.', 10.99, 10.27, 4.55, 71, NULL, 2, 'https://cdn.dummyjson.com/product-images/groceries/dog-food/thumbnail.webp', '''canin'':14B ''companion'':15B ''design'':7B ''dog'':1A,5B ''essenti'':10B ''food'':2A,6B ''formul'':4B ''groceri'':16C ''nutrient'':11B ''provid'':9B ''special'':3B', NULL),
	(24, 'Fish Steak', 'Quality fish steak, suitable for grilling, baking, or pan-searing.', 14.99, 4.23, 3.78, 74, NULL, 2, 'https://cdn.dummyjson.com/product-images/groceries/fish-steak/thumbnail.webp', '''bake'':9B ''fish'':1A,4B ''grill'':8B ''groceri'':14C ''pan'':12B ''pan-sear'':11B ''qualiti'':3B ''sear'':13B ''steak'':2A,5B ''suitabl'':6B', NULL),
	(25, 'Green Bell Pepper', 'Fresh and vibrant green bell pepper, perfect for adding color and flavor to your dishes.', 1.29, 0.16, 3.25, 33, NULL, 2, 'https://cdn.dummyjson.com/product-images/groceries/green-bell-pepper/thumbnail.webp', '''ad'':12B ''bell'':2A,8B ''color'':13B ''dish'':18B ''flavor'':15B ''fresh'':4B ''green'':1A,7B ''groceri'':19C ''pepper'':3A,9B ''perfect'':10B ''vibrant'':6B', NULL),
	(26, 'Green Chili Pepper', 'Spicy green chili pepper, ideal for adding heat to your favorite recipes.', 0.99, 1, 3.66, 3, NULL, 2, 'https://cdn.dummyjson.com/product-images/groceries/green-chili-pepper/thumbnail.webp', '''ad'':10B ''chili'':2A,6B ''favorit'':14B ''green'':1A,5B ''groceri'':16C ''heat'':11B ''ideal'':8B ''pepper'':3A,7B ''recip'':15B ''spici'':4B', NULL),
	(27, 'Honey Jar', 'Pure and natural honey in a convenient jar, perfect for sweetening beverages or drizzling over food.', 6.99, 14.4, 3.97, 34, NULL, 2, 'https://cdn.dummyjson.com/product-images/groceries/honey-jar/thumbnail.webp', '''beverag'':14B ''conveni'':9B ''drizzl'':16B ''food'':18B ''groceri'':19C ''honey'':1A,6B ''jar'':2A,10B ''natur'':5B ''perfect'':11B ''pure'':3B ''sweeten'':13B', NULL),
	(28, 'Ice Cream', 'Creamy and delicious ice cream, available in various flavors for a delightful treat.', 5.49, 8.69, 3.39, 27, NULL, 2, 'https://cdn.dummyjson.com/product-images/groceries/ice-cream/thumbnail.webp', '''avail'':8B ''cream'':2A,7B ''creami'':3B ''delici'':5B ''delight'':14B ''flavor'':11B ''groceri'':16C ''ice'':1A,6B ''treat'':15B ''various'':10B', NULL),
	(29, 'Juice', 'Refreshing fruit juice, packed with vitamins and great for staying hydrated.', 3.99, 12.06, 3.94, 50, NULL, 2, 'https://cdn.dummyjson.com/product-images/groceries/juice/thumbnail.webp', '''fruit'':3B ''great'':9B ''groceri'':13C ''hydrat'':12B ''juic'':1A,4B ''pack'':5B ''refresh'':2B ''stay'':11B ''vitamin'':7B', NULL),
	(30, 'Kiwi', 'Nutrient-rich kiwi, perfect for snacking or adding a tropical twist to your dishes.', 2.49, 15.22, 4.93, 99, NULL, 2, 'https://cdn.dummyjson.com/product-images/groceries/kiwi/thumbnail.webp', '''ad'':10B ''dish'':16B ''groceri'':17C ''kiwi'':1A,5B ''nutrient'':3B ''nutrient-rich'':2B ''perfect'':6B ''rich'':4B ''snack'':8B ''tropic'':12B ''twist'':13B', NULL);


ALTER TABLE marketplacedemo.products ENABLE TRIGGER ALL;


--
-- Data for Name: product_images; Type: TABLE DATA; Schema: marketplacedemo; Owner: tcero
--

ALTER TABLE marketplacedemo.product_images DISABLE TRIGGER ALL;

INSERT INTO marketplacedemo.product_images (id, product_id, url) VALUES
	(55, 1, 'https://cdn.dummyjson.com/product-images/beauty/essence-mascara-lash-princess/1.webp'),
	(56, 2, 'https://cdn.dummyjson.com/product-images/beauty/eyeshadow-palette-with-mirror/1.webp'),
	(57, 3, 'https://cdn.dummyjson.com/product-images/beauty/powder-canister/1.webp'),
	(58, 4, 'https://cdn.dummyjson.com/product-images/beauty/red-lipstick/1.webp'),
	(59, 5, 'https://cdn.dummyjson.com/product-images/beauty/red-nail-polish/1.webp'),
	(60, 6, 'https://cdn.dummyjson.com/product-images/fragrances/calvin-klein-ck-one/1.webp'),
	(61, 6, 'https://cdn.dummyjson.com/product-images/fragrances/calvin-klein-ck-one/2.webp'),
	(62, 6, 'https://cdn.dummyjson.com/product-images/fragrances/calvin-klein-ck-one/3.webp'),
	(63, 7, 'https://cdn.dummyjson.com/product-images/fragrances/chanel-coco-noir-eau-de/1.webp'),
	(64, 7, 'https://cdn.dummyjson.com/product-images/fragrances/chanel-coco-noir-eau-de/2.webp'),
	(65, 7, 'https://cdn.dummyjson.com/product-images/fragrances/chanel-coco-noir-eau-de/3.webp'),
	(66, 8, 'https://cdn.dummyjson.com/product-images/fragrances/dior-j''adore/1.webp'),
	(67, 8, 'https://cdn.dummyjson.com/product-images/fragrances/dior-j''adore/2.webp'),
	(68, 8, 'https://cdn.dummyjson.com/product-images/fragrances/dior-j''adore/3.webp'),
	(69, 9, 'https://cdn.dummyjson.com/product-images/fragrances/dolce-shine-eau-de/1.webp'),
	(70, 9, 'https://cdn.dummyjson.com/product-images/fragrances/dolce-shine-eau-de/2.webp'),
	(71, 9, 'https://cdn.dummyjson.com/product-images/fragrances/dolce-shine-eau-de/3.webp'),
	(72, 10, 'https://cdn.dummyjson.com/product-images/fragrances/gucci-bloom-eau-de/1.webp'),
	(73, 10, 'https://cdn.dummyjson.com/product-images/fragrances/gucci-bloom-eau-de/2.webp'),
	(74, 10, 'https://cdn.dummyjson.com/product-images/fragrances/gucci-bloom-eau-de/3.webp'),
	(75, 11, 'https://cdn.dummyjson.com/product-images/furniture/annibale-colombo-bed/1.webp'),
	(76, 11, 'https://cdn.dummyjson.com/product-images/furniture/annibale-colombo-bed/2.webp'),
	(77, 11, 'https://cdn.dummyjson.com/product-images/furniture/annibale-colombo-bed/3.webp'),
	(78, 12, 'https://cdn.dummyjson.com/product-images/furniture/annibale-colombo-sofa/1.webp'),
	(79, 12, 'https://cdn.dummyjson.com/product-images/furniture/annibale-colombo-sofa/2.webp'),
	(80, 12, 'https://cdn.dummyjson.com/product-images/furniture/annibale-colombo-sofa/3.webp'),
	(81, 13, 'https://cdn.dummyjson.com/product-images/furniture/bedside-table-african-cherry/1.webp'),
	(82, 13, 'https://cdn.dummyjson.com/product-images/furniture/bedside-table-african-cherry/2.webp'),
	(83, 13, 'https://cdn.dummyjson.com/product-images/furniture/bedside-table-african-cherry/3.webp'),
	(84, 14, 'https://cdn.dummyjson.com/product-images/furniture/knoll-saarinen-executive-conference-chair/1.webp'),
	(85, 14, 'https://cdn.dummyjson.com/product-images/furniture/knoll-saarinen-executive-conference-chair/2.webp'),
	(86, 14, 'https://cdn.dummyjson.com/product-images/furniture/knoll-saarinen-executive-conference-chair/3.webp'),
	(87, 15, 'https://cdn.dummyjson.com/product-images/furniture/wooden-bathroom-sink-with-mirror/1.webp'),
	(88, 15, 'https://cdn.dummyjson.com/product-images/furniture/wooden-bathroom-sink-with-mirror/2.webp'),
	(89, 15, 'https://cdn.dummyjson.com/product-images/furniture/wooden-bathroom-sink-with-mirror/3.webp'),
	(90, 16, 'https://cdn.dummyjson.com/product-images/groceries/apple/1.webp'),
	(91, 17, 'https://cdn.dummyjson.com/product-images/groceries/beef-steak/1.webp'),
	(92, 18, 'https://cdn.dummyjson.com/product-images/groceries/cat-food/1.webp'),
	(93, 19, 'https://cdn.dummyjson.com/product-images/groceries/chicken-meat/1.webp'),
	(94, 19, 'https://cdn.dummyjson.com/product-images/groceries/chicken-meat/2.webp'),
	(95, 20, 'https://cdn.dummyjson.com/product-images/groceries/cooking-oil/1.webp'),
	(96, 21, 'https://cdn.dummyjson.com/product-images/groceries/cucumber/1.webp'),
	(97, 22, 'https://cdn.dummyjson.com/product-images/groceries/dog-food/1.webp'),
	(98, 23, 'https://cdn.dummyjson.com/product-images/groceries/eggs/1.webp'),
	(99, 24, 'https://cdn.dummyjson.com/product-images/groceries/fish-steak/1.webp'),
	(100, 25, 'https://cdn.dummyjson.com/product-images/groceries/green-bell-pepper/1.webp'),
	(101, 26, 'https://cdn.dummyjson.com/product-images/groceries/green-chili-pepper/1.webp'),
	(102, 27, 'https://cdn.dummyjson.com/product-images/groceries/honey-jar/1.webp'),
	(103, 28, 'https://cdn.dummyjson.com/product-images/groceries/ice-cream/1.webp'),
	(104, 28, 'https://cdn.dummyjson.com/product-images/groceries/ice-cream/2.webp'),
	(105, 28, 'https://cdn.dummyjson.com/product-images/groceries/ice-cream/3.webp'),
	(106, 28, 'https://cdn.dummyjson.com/product-images/groceries/ice-cream/4.webp'),
	(107, 29, 'https://cdn.dummyjson.com/product-images/groceries/juice/1.webp'),
	(108, 30, 'https://cdn.dummyjson.com/product-images/groceries/kiwi/1.webp'),
	(109, 30, 'https://cdn.dummyjson.com/product-images/groceries/kiwi/1.webp');


ALTER TABLE marketplacedemo.product_images ENABLE TRIGGER ALL;

--
-- Data for Name: products_stage; Type: TABLE DATA; Schema: marketplacedemo; Owner: tcero
--

ALTER TABLE marketplacedemo.products_stage DISABLE TRIGGER ALL;

INSERT INTO marketplacedemo.products_stage (id, title, description, price, discount_percentage, rating, stock, brand, category, thumbnail, deleted_at) VALUES
	(1, 'Essence Mascara Lash Princess', 'The Essence Mascara Lash Princess is a popular mascara known for its volumizing and lengthening effects. Achieve dramatic lashes with this long-lasting and cruelty-free formula.', 9.99, 10.48, 2.56, 99, 'Essence', 'beauty', 'https://cdn.dummyjson.com/product-images/beauty/essence-mascara-lash-princess/thumbnail.webp', NULL),
	(2, 'Eyeshadow Palette with Mirror', 'The Eyeshadow Palette with Mirror offers a versatile range of eyeshadow shades for creating stunning eye looks. With a built-in mirror, it''s convenient for on-the-go makeup application.', 19.99, 18.19, 2.86, 34, 'Glamour Beauty', 'beauty', 'https://cdn.dummyjson.com/product-images/beauty/eyeshadow-palette-with-mirror/thumbnail.webp', NULL),
	(3, 'Powder Canister', 'The Powder Canister is a finely milled setting powder designed to set makeup and control shine. With a lightweight and translucent formula, it provides a smooth and matte finish.', 14.99, 9.84, 4.64, 89, 'Velvet Touch', 'beauty', 'https://cdn.dummyjson.com/product-images/beauty/powder-canister/thumbnail.webp', NULL),
	(4, 'Red Lipstick', 'The Red Lipstick is a classic and bold choice for adding a pop of color to your lips. With a creamy and pigmented formula, it provides a vibrant and long-lasting finish.', 12.99, 12.16, 4.36, 91, 'Chic Cosmetics', 'beauty', 'https://cdn.dummyjson.com/product-images/beauty/red-lipstick/thumbnail.webp', NULL),
	(5, 'Red Nail Polish', 'The Red Nail Polish offers a rich and glossy red hue for vibrant and polished nails. With a quick-drying formula, it provides a salon-quality finish at home.', 8.99, 11.44, 4.32, 79, 'Nail Couture', 'beauty', 'https://cdn.dummyjson.com/product-images/beauty/red-nail-polish/thumbnail.webp', NULL),
	(6, 'Calvin Klein CK One', 'CK One by Calvin Klein is a classic unisex fragrance, known for its fresh and clean scent. It''s a versatile fragrance suitable for everyday wear.', 49.99, 1.89, 4.37, 29, 'Calvin Klein', 'fragrances', 'https://cdn.dummyjson.com/product-images/fragrances/calvin-klein-ck-one/thumbnail.webp', NULL),
	(7, 'Chanel Coco Noir Eau De', 'Coco Noir by Chanel is an elegant and mysterious fragrance, featuring notes of grapefruit, rose, and sandalwood. Perfect for evening occasions.', 129.99, 16.51, 4.26, 58, 'Chanel', 'fragrances', 'https://cdn.dummyjson.com/product-images/fragrances/chanel-coco-noir-eau-de/thumbnail.webp', NULL),
	(8, 'Dior J''adore', 'J''adore by Dior is a luxurious and floral fragrance, known for its blend of ylang-ylang, rose, and jasmine. It embodies femininity and sophistication.', 89.99, 14.72, 3.8, 98, 'Dior', 'fragrances', 'https://cdn.dummyjson.com/product-images/fragrances/dior-j''adore/thumbnail.webp', NULL),
	(9, 'Dolce Shine Eau de', 'Dolce Shine by Dolce & Gabbana is a vibrant and fruity fragrance, featuring notes of mango, jasmine, and blonde woods. It''s a joyful and youthful scent.', 69.99, 0.62, 3.96, 4, 'Dolce & Gabbana', 'fragrances', 'https://cdn.dummyjson.com/product-images/fragrances/dolce-shine-eau-de/thumbnail.webp', NULL),
	(10, 'Gucci Bloom Eau de', 'Gucci Bloom by Gucci is a floral and captivating fragrance, with notes of tuberose, jasmine, and Rangoon creeper. It''s a modern and romantic scent.', 79.99, 14.39, 2.74, 91, 'Gucci', 'fragrances', 'https://cdn.dummyjson.com/product-images/fragrances/gucci-bloom-eau-de/thumbnail.webp', NULL),
	(11, 'Annibale Colombo Bed', 'The Annibale Colombo Bed is a luxurious and elegant bed frame, crafted with high-quality materials for a comfortable and stylish bedroom.', 1899.99, 8.57, 4.77, 88, 'Annibale Colombo', 'furniture', 'https://cdn.dummyjson.com/product-images/furniture/annibale-colombo-bed/thumbnail.webp', NULL),
	(12, 'Annibale Colombo Sofa', 'The Annibale Colombo Sofa is a sophisticated and comfortable seating option, featuring exquisite design and premium upholstery for your living room.', 2499.99, 14.4, 3.92, 60, 'Annibale Colombo', 'furniture', 'https://cdn.dummyjson.com/product-images/furniture/annibale-colombo-sofa/thumbnail.webp', NULL),
	(13, 'Bedside Table African Cherry', 'The Bedside Table in African Cherry is a stylish and functional addition to your bedroom, providing convenient storage space and a touch of elegance.', 299.99, 19.09, 2.87, 64, 'Furniture Co.', 'furniture', 'https://cdn.dummyjson.com/product-images/furniture/bedside-table-african-cherry/thumbnail.webp', NULL),
	(14, 'Knoll Saarinen Executive Conference Chair', 'The Knoll Saarinen Executive Conference Chair is a modern and ergonomic chair, perfect for your office or conference room with its timeless design.', 499.99, 2.01, 4.88, 26, 'Knoll', 'furniture', 'https://cdn.dummyjson.com/product-images/furniture/knoll-saarinen-executive-conference-chair/thumbnail.webp', NULL),
	(15, 'Wooden Bathroom Sink With Mirror', 'The Wooden Bathroom Sink with Mirror is a unique and stylish addition to your bathroom, featuring a wooden sink countertop and a matching mirror.', 799.99, 8.8, 3.59, 7, 'Bath Trends', 'furniture', 'https://cdn.dummyjson.com/product-images/furniture/wooden-bathroom-sink-with-mirror/thumbnail.webp', NULL),
	(16, 'Apple', 'Fresh and crisp apples, perfect for snacking or incorporating into various recipes.', 1.99, 12.62, 4.19, 8, NULL, 'groceries', 'https://cdn.dummyjson.com/product-images/groceries/apple/thumbnail.webp', NULL),
	(17, 'Beef Steak', 'High-quality beef steak, great for grilling or cooking to your preferred level of doneness.', 12.99, 9.61, 4.47, 86, NULL, 'groceries', 'https://cdn.dummyjson.com/product-images/groceries/beef-steak/thumbnail.webp', NULL),
	(18, 'Cat Food', 'Nutritious cat food formulated to meet the dietary needs of your feline friend.', 8.99, 9.58, 3.13, 46, NULL, 'groceries', 'https://cdn.dummyjson.com/product-images/groceries/cat-food/thumbnail.webp', NULL),
	(19, 'Chicken Meat', 'Fresh and tender chicken meat, suitable for various culinary preparations.', 9.99, 13.7, 3.19, 97, NULL, 'groceries', 'https://cdn.dummyjson.com/product-images/groceries/chicken-meat/thumbnail.webp', NULL),
	(20, 'Cooking Oil', 'Versatile cooking oil suitable for frying, sautéing, and various culinary applications.', 4.99, 9.33, 4.8, 10, NULL, 'groceries', 'https://cdn.dummyjson.com/product-images/groceries/cooking-oil/thumbnail.webp', NULL),
	(21, 'Cucumber', 'Crisp and hydrating cucumbers, ideal for salads, snacks, or as a refreshing side.', 1.49, 0.16, 4.07, 84, NULL, 'groceries', 'https://cdn.dummyjson.com/product-images/groceries/cucumber/thumbnail.webp', NULL),
	(22, 'Dog Food', 'Specially formulated dog food designed to provide essential nutrients for your canine companion.', 10.99, 10.27, 4.55, 71, NULL, 'groceries', 'https://cdn.dummyjson.com/product-images/groceries/dog-food/thumbnail.webp', NULL),
	(23, 'Eggs', 'Fresh eggs, a versatile ingredient for baking, cooking, or breakfast.', 2.99, 11.05, 2.53, 9, NULL, 'groceries', 'https://cdn.dummyjson.com/product-images/groceries/eggs/thumbnail.webp', NULL),
	(24, 'Fish Steak', 'Quality fish steak, suitable for grilling, baking, or pan-searing.', 14.99, 4.23, 3.78, 74, NULL, 'groceries', 'https://cdn.dummyjson.com/product-images/groceries/fish-steak/thumbnail.webp', NULL),
	(25, 'Green Bell Pepper', 'Fresh and vibrant green bell pepper, perfect for adding color and flavor to your dishes.', 1.29, 0.16, 3.25, 33, NULL, 'groceries', 'https://cdn.dummyjson.com/product-images/groceries/green-bell-pepper/thumbnail.webp', NULL),
	(26, 'Green Chili Pepper', 'Spicy green chili pepper, ideal for adding heat to your favorite recipes.', 0.99, 1, 3.66, 3, NULL, 'groceries', 'https://cdn.dummyjson.com/product-images/groceries/green-chili-pepper/thumbnail.webp', NULL),
	(27, 'Honey Jar', 'Pure and natural honey in a convenient jar, perfect for sweetening beverages or drizzling over food.', 6.99, 14.4, 3.97, 34, NULL, 'groceries', 'https://cdn.dummyjson.com/product-images/groceries/honey-jar/thumbnail.webp', NULL),
	(28, 'Ice Cream', 'Creamy and delicious ice cream, available in various flavors for a delightful treat.', 5.49, 8.69, 3.39, 27, NULL, 'groceries', 'https://cdn.dummyjson.com/product-images/groceries/ice-cream/thumbnail.webp', NULL),
	(29, 'Juice', 'Refreshing fruit juice, packed with vitamins and great for staying hydrated.', 3.99, 12.06, 3.94, 50, NULL, 'groceries', 'https://cdn.dummyjson.com/product-images/groceries/juice/thumbnail.webp', NULL),
	(30, 'Kiwi', 'Nutrient-rich kiwi, perfect for snacking or adding a tropical twist to your dishes.', 2.49, 15.22, 4.93, 99, NULL, 'groceries', 'https://cdn.dummyjson.com/product-images/groceries/kiwi/thumbnail.webp', NULL);


ALTER TABLE marketplacedemo.products_stage ENABLE TRIGGER ALL;

--
-- Data for Name: product_images_stage; Type: TABLE DATA; Schema: marketplacedemo; Owner: tcero
--

ALTER TABLE marketplacedemo.product_images_stage DISABLE TRIGGER ALL;

INSERT INTO marketplacedemo.product_images_stage (id, product_id, url) VALUES
	(1, 1, 'https://cdn.dummyjson.com/product-images/beauty/essence-mascara-lash-princess/1.webp'),
	(2, 2, 'https://cdn.dummyjson.com/product-images/beauty/eyeshadow-palette-with-mirror/1.webp'),
	(3, 3, 'https://cdn.dummyjson.com/product-images/beauty/powder-canister/1.webp'),
	(4, 4, 'https://cdn.dummyjson.com/product-images/beauty/red-lipstick/1.webp'),
	(5, 5, 'https://cdn.dummyjson.com/product-images/beauty/red-nail-polish/1.webp'),
	(6, 6, 'https://cdn.dummyjson.com/product-images/fragrances/calvin-klein-ck-one/1.webp'),
	(7, 6, 'https://cdn.dummyjson.com/product-images/fragrances/calvin-klein-ck-one/2.webp'),
	(8, 6, 'https://cdn.dummyjson.com/product-images/fragrances/calvin-klein-ck-one/3.webp'),
	(9, 7, 'https://cdn.dummyjson.com/product-images/fragrances/chanel-coco-noir-eau-de/1.webp'),
	(10, 7, 'https://cdn.dummyjson.com/product-images/fragrances/chanel-coco-noir-eau-de/2.webp'),
	(11, 7, 'https://cdn.dummyjson.com/product-images/fragrances/chanel-coco-noir-eau-de/3.webp'),
	(12, 8, 'https://cdn.dummyjson.com/product-images/fragrances/dior-j''adore/1.webp'),
	(13, 8, 'https://cdn.dummyjson.com/product-images/fragrances/dior-j''adore/2.webp'),
	(14, 8, 'https://cdn.dummyjson.com/product-images/fragrances/dior-j''adore/3.webp'),
	(15, 9, 'https://cdn.dummyjson.com/product-images/fragrances/dolce-shine-eau-de/1.webp'),
	(16, 9, 'https://cdn.dummyjson.com/product-images/fragrances/dolce-shine-eau-de/2.webp'),
	(17, 9, 'https://cdn.dummyjson.com/product-images/fragrances/dolce-shine-eau-de/3.webp'),
	(18, 10, 'https://cdn.dummyjson.com/product-images/fragrances/gucci-bloom-eau-de/1.webp'),
	(19, 10, 'https://cdn.dummyjson.com/product-images/fragrances/gucci-bloom-eau-de/2.webp'),
	(20, 10, 'https://cdn.dummyjson.com/product-images/fragrances/gucci-bloom-eau-de/3.webp'),
	(21, 11, 'https://cdn.dummyjson.com/product-images/furniture/annibale-colombo-bed/1.webp'),
	(22, 11, 'https://cdn.dummyjson.com/product-images/furniture/annibale-colombo-bed/2.webp'),
	(23, 11, 'https://cdn.dummyjson.com/product-images/furniture/annibale-colombo-bed/3.webp'),
	(24, 12, 'https://cdn.dummyjson.com/product-images/furniture/annibale-colombo-sofa/1.webp'),
	(25, 12, 'https://cdn.dummyjson.com/product-images/furniture/annibale-colombo-sofa/2.webp'),
	(26, 12, 'https://cdn.dummyjson.com/product-images/furniture/annibale-colombo-sofa/3.webp'),
	(27, 13, 'https://cdn.dummyjson.com/product-images/furniture/bedside-table-african-cherry/1.webp'),
	(28, 13, 'https://cdn.dummyjson.com/product-images/furniture/bedside-table-african-cherry/2.webp'),
	(29, 13, 'https://cdn.dummyjson.com/product-images/furniture/bedside-table-african-cherry/3.webp'),
	(30, 14, 'https://cdn.dummyjson.com/product-images/furniture/knoll-saarinen-executive-conference-chair/1.webp'),
	(31, 14, 'https://cdn.dummyjson.com/product-images/furniture/knoll-saarinen-executive-conference-chair/2.webp'),
	(32, 14, 'https://cdn.dummyjson.com/product-images/furniture/knoll-saarinen-executive-conference-chair/3.webp'),
	(33, 15, 'https://cdn.dummyjson.com/product-images/furniture/wooden-bathroom-sink-with-mirror/1.webp'),
	(34, 15, 'https://cdn.dummyjson.com/product-images/furniture/wooden-bathroom-sink-with-mirror/2.webp'),
	(35, 15, 'https://cdn.dummyjson.com/product-images/furniture/wooden-bathroom-sink-with-mirror/3.webp'),
	(36, 16, 'https://cdn.dummyjson.com/product-images/groceries/apple/1.webp'),
	(37, 17, 'https://cdn.dummyjson.com/product-images/groceries/beef-steak/1.webp'),
	(38, 18, 'https://cdn.dummyjson.com/product-images/groceries/cat-food/1.webp'),
	(39, 19, 'https://cdn.dummyjson.com/product-images/groceries/chicken-meat/1.webp'),
	(40, 19, 'https://cdn.dummyjson.com/product-images/groceries/chicken-meat/2.webp'),
	(41, 20, 'https://cdn.dummyjson.com/product-images/groceries/cooking-oil/1.webp'),
	(42, 21, 'https://cdn.dummyjson.com/product-images/groceries/cucumber/1.webp'),
	(43, 22, 'https://cdn.dummyjson.com/product-images/groceries/dog-food/1.webp'),
	(44, 23, 'https://cdn.dummyjson.com/product-images/groceries/eggs/1.webp'),
	(45, 24, 'https://cdn.dummyjson.com/product-images/groceries/fish-steak/1.webp'),
	(46, 25, 'https://cdn.dummyjson.com/product-images/groceries/green-bell-pepper/1.webp'),
	(47, 26, 'https://cdn.dummyjson.com/product-images/groceries/green-chili-pepper/1.webp'),
	(48, 27, 'https://cdn.dummyjson.com/product-images/groceries/honey-jar/1.webp'),
	(49, 28, 'https://cdn.dummyjson.com/product-images/groceries/ice-cream/1.webp'),
	(50, 28, 'https://cdn.dummyjson.com/product-images/groceries/ice-cream/2.webp'),
	(51, 28, 'https://cdn.dummyjson.com/product-images/groceries/ice-cream/3.webp'),
	(52, 28, 'https://cdn.dummyjson.com/product-images/groceries/ice-cream/4.webp'),
	(53, 29, 'https://cdn.dummyjson.com/product-images/groceries/juice/1.webp'),
	(54, 30, 'https://cdn.dummyjson.com/product-images/groceries/kiwi/1.webp'),
	(108, 30, 'https://cdn.dummyjson.com/product-images/groceries/kiwi/1.webp');


ALTER TABLE marketplacedemo.product_images_stage ENABLE TRIGGER ALL;
