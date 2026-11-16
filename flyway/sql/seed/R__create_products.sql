--
-- Data for Name: category; Type: TABLE DATA; Schema: marketplacedemo; Owner: tcero
--

ALTER TABLE inventory.category DISABLE TRIGGER ALL;

INSERT INTO inventory.category (id, name) VALUES
	('7f3a9c21-8d64-4b2e-a591-3c8f0d72e614', 'beauty'),
	('e2b8f491-6c35-4d7a-9f02-8a1c5e63b790', 'groceries'),
	('3d91a6e8-2f47-4c05-b823-7e9d1a54f630', 'furniture'),
	('a5c7e302-91b8-46d4-8f61-2b3e9c70d845', 'fragrances');


ALTER TABLE inventory.category ENABLE TRIGGER ALL;

--
-- Data for Name: products; Type: TABLE DATA; Schema: marketplacedemo; Owner: tcero
--

ALTER TABLE inventory.products DISABLE TRIGGER ALL;

INSERT INTO inventory.products (id, sku, title, description, price, discount_percentage, rating, quantity, brand, categoryid, thumbnail, search_vector, deleted_at, active, created_at) VALUES
	('2657a1e2-38c5-4d71-9d7b-94664f66189b', 'BEA-NAI-NAI-005', 'Red Nail Polish', 'The Red Nail Polish offers a rich and glossy red hue for vibrant and polished nails. With a quick-drying formula, it provides a salon-quality finish at home.', 8.99, 11.44, 4.32, 79, 'Nail Couture', '7f3a9c21-8d64-4b2e-a591-3c8f0d72e614', 'https://cdn.dummyjson.com/product-images/beauty/red-nail-polish/thumbnail.webp', '''beauti'':37C ''coutur'':36C ''dri'':24B ''finish'':32B ''formula'':25B ''glossi'':12B ''home'':34B ''hue'':14B ''nail'':2A,6B,19B,35C ''offer'':8B ''polish'':3A,7B,18B ''provid'':27B ''qualiti'':31B ''quick'':23B ''quick-dri'':22B ''red'':1A,5B,13B ''rich'':10B ''salon'':30B ''salon-qu'':29B ''vibrant'':16B', NULL, true, '2026-11-14 22:15:27.504324+00'),
	('831ad86a-6956-475f-8a61-5a89d72b130d', 'BEA-CHI-LIP-004', 'Red Lipstick', 'The Red Lipstick is a classic and bold choice for adding a pop of color to your lips. With a creamy and pigmented formula, it provides a vibrant and long-lasting finish.', 12.99, 12.16, 4.36, 91, 'Chic Cosmetics', '7f3a9c21-8d64-4b2e-a591-3c8f0d72e614', 'https://cdn.dummyjson.com/product-images/beauty/red-lipstick/thumbnail.webp', '''ad'':13B ''beauti'':38C ''bold'':10B ''chic'':36C ''choic'':11B ''classic'':8B ''color'':17B ''cosmet'':37C ''creami'':23B ''finish'':35B ''formula'':26B ''last'':34B ''lip'':20B ''lipstick'':2A,5B ''long'':33B ''long-last'':32B ''pigment'':25B ''pop'':15B ''provid'':28B ''red'':1A,4B ''vibrant'':30B', NULL, true, '2026-11-14 22:15:27.504324+00'),
	('88a7c616-21ff-45bf-aa22-dbc6e3b0acb4', 'BEA-VEL-POW-003', 'Powder Canister', 'The Powder Canister is a finely milled setting powder designed to set makeup and control shine. With a lightweight and translucent formula, it provides a smooth and matte finish.', 14.99, 9.84, 4.64, 89, 'Velvet Touch', '7f3a9c21-8d64-4b2e-a591-3c8f0d72e614', 'https://cdn.dummyjson.com/product-images/beauty/powder-canister/thumbnail.webp', '''beauti'':34C ''canist'':2A,5B ''control'':17B ''design'':12B ''fine'':8B ''finish'':31B ''formula'':24B ''lightweight'':21B ''makeup'':15B ''matt'':30B ''mill'':9B ''powder'':1A,4B,11B ''provid'':26B ''set'':10B,14B ''shine'':18B ''smooth'':28B ''touch'':33C ''transluc'':23B ''velvet'':32C', NULL, true, '2026-11-14 22:15:27.504324+00'),
	('e2d666e5-506c-4b51-ba3b-290afd160bf0', 'BEA-GLA-EYE-002', 'Eyeshadow Palette with Mirror', 'The Eyeshadow Palette with Mirror offers a versatile range of eyeshadow shades for creating stunning eye looks. With a built-in mirror, it''s convenient for on-the-go makeup application.', 19.99, 18.19, 2.86, 34, 'Glamour Beauty', '7f3a9c21-8d64-4b2e-a591-3c8f0d72e614', 'https://cdn.dummyjson.com/product-images/beauty/eyeshadow-palette-with-mirror/thumbnail.webp', '''applic'':37B ''beauti'':39C,40C ''built'':25B ''built-in'':24B ''conveni'':30B ''creat'':18B ''eye'':20B ''eyeshadow'':1A,6B,15B ''glamour'':38C ''go'':35B ''look'':21B ''makeup'':36B ''mirror'':4A,9B,27B ''offer'':10B ''on-the-go'':32B ''palett'':2A,7B ''rang'':13B ''shade'':16B ''stun'':19B ''versatil'':12B', NULL, true, '2026-11-14 22:15:27.504324+00'),
	('fc1f2f0c-7d52-48d9-895a-50f8e20e0bf9', 'BEA-ESS-ESS-001', 'Essence Mascara Lash Princess', 'The Essence Mascara Lash Princess is a popular mascara known for its volumizing and lengthening effects. Achieve dramatic lashes with this long-lasting and cruelty-free formula.', 9.99, 10.48, 2.56, 99, 'Essence', '7f3a9c21-8d64-4b2e-a591-3c8f0d72e614', 'https://cdn.dummyjson.com/product-images/beauty/essence-mascara-lash-princess/thumbnail.webp', '''achiev'':21B ''beauti'':35C ''cruelti'':31B ''cruelty-fre'':30B ''dramat'':22B ''effect'':20B ''essenc'':1A,6B,34C ''formula'':33B ''free'':32B ''known'':14B ''lash'':3A,8B,23B ''last'':28B ''lengthen'':19B ''long'':27B ''long-last'':26B ''mascara'':2A,7B,13B ''popular'':12B ''princess'':4A,9B ''volum'':17B', NULL, true, '2026-11-14 22:15:27.504324+00'),
	('911ea40b-e358-4b60-92ee-a1b862e2234f', 'GRO-BRD-KIW-030', 'Kiwi', 'Nutrient-rich kiwi, perfect for snacking or adding a tropical twist to your dishes.', 2.49, 15.22, 4.93, 99, NULL, 'e2b8f491-6c35-4d7a-9f02-8a1c5e63b790', 'https://cdn.dummyjson.com/product-images/groceries/kiwi/thumbnail.webp', '''ad'':10B ''dish'':16B ''groceri'':17C ''kiwi'':1A,5B ''nutrient'':3B ''nutrient-rich'':2B ''perfect'':6B ''rich'':4B ''snack'':8B ''tropic'':12B ''twist'':13B', NULL, true, '2026-11-14 22:15:27.504324+00'),
	('76477601-96bc-4fb2-a08c-ee37a874cb31', 'GRO-BRD-JUI-029', 'Juice', 'Refreshing fruit juice, packed with vitamins and great for staying hydrated.', 3.99, 12.06, 3.94, 50, NULL, 'e2b8f491-6c35-4d7a-9f02-8a1c5e63b790', 'https://cdn.dummyjson.com/product-images/groceries/juice/thumbnail.webp', '''fruit'':3B ''great'':9B ''groceri'':13C ''hydrat'':12B ''juic'':1A,4B ''pack'':5B ''refresh'':2B ''stay'':11B ''vitamin'':7B', NULL, true, '2026-11-14 22:15:27.504324+00'),
	('5fc579d4-d64f-49d6-b215-b2db6cc239ab', 'GRO-BRD-CRE-028', 'Ice Cream', 'Creamy and delicious ice cream, available in various flavors for a delightful treat.', 5.49, 8.69, 3.39, 27, NULL, 'e2b8f491-6c35-4d7a-9f02-8a1c5e63b790', 'https://cdn.dummyjson.com/product-images/groceries/ice-cream/thumbnail.webp', '''avail'':8B ''cream'':2A,7B ''creami'':3B ''delici'':5B ''delight'':14B ''flavor'':11B ''groceri'':16C ''ice'':1A,6B ''treat'':15B ''various'':10B', NULL, true, '2026-11-14 22:15:27.504324+00'),
	('f7097af4-d9e7-4e8b-b68c-6acda085a959', 'GRO-BRD-HON-027', 'Honey Jar', 'Pure and natural honey in a convenient jar, perfect for sweetening beverages or drizzling over food.', 6.99, 14.4, 3.97, 34, NULL, 'e2b8f491-6c35-4d7a-9f02-8a1c5e63b790', 'https://cdn.dummyjson.com/product-images/groceries/honey-jar/thumbnail.webp', '''beverag'':14B ''conveni'':9B ''drizzl'':16B ''food'':18B ''groceri'':19C ''honey'':1A,6B ''jar'':2A,10B ''natur'':5B ''perfect'':11B ''pure'':3B ''sweeten'':13B', NULL, true, '2026-11-14 22:15:27.504324+00'),
	('44fc661b-849a-4c0e-acde-9d80fd67ca0d', 'GRO-BRD-GRE-026', 'Green Chili Pepper', 'Spicy green chili pepper, ideal for adding heat to your favorite recipes.', 0.99, 1, 3.66, 3, NULL, 'e2b8f491-6c35-4d7a-9f02-8a1c5e63b790', 'https://cdn.dummyjson.com/product-images/groceries/green-chili-pepper/thumbnail.webp', '''ad'':10B ''chili'':2A,6B ''favorit'':14B ''green'':1A,5B ''groceri'':16C ''heat'':11B ''ideal'':8B ''pepper'':3A,7B ''recip'':15B ''spici'':4B', NULL, true, '2026-11-14 22:15:27.504324+00'),
	('f87b350d-53cd-4e00-8caf-69fe499a2e17', 'GRO-BRD-GRE-025', 'Green Bell Pepper', 'Fresh and vibrant green bell pepper, perfect for adding color and flavor to your dishes.', 1.29, 0.16, 3.25, 33, NULL, 'e2b8f491-6c35-4d7a-9f02-8a1c5e63b790', 'https://cdn.dummyjson.com/product-images/groceries/green-bell-pepper/thumbnail.webp', '''ad'':12B ''bell'':2A,8B ''color'':13B ''dish'':18B ''flavor'':15B ''fresh'':4B ''green'':1A,7B ''groceri'':19C ''pepper'':3A,9B ''perfect'':10B ''vibrant'':6B', NULL, true, '2026-11-14 22:15:27.504324+00'),
	('fdec0dca-29d0-459f-8cab-9365bc5b12d2', 'GRO-BRD-FIS-024', 'Fish Steak', 'Quality fish steak, suitable for grilling, baking, or pan-searing.', 14.99, 4.23, 3.78, 74, NULL, 'e2b8f491-6c35-4d7a-9f02-8a1c5e63b790', 'https://cdn.dummyjson.com/product-images/groceries/fish-steak/thumbnail.webp', '''bake'':9B ''fish'':1A,4B ''grill'':8B ''groceri'':14C ''pan'':12B ''pan-sear'':11B ''qualiti'':3B ''sear'':13B ''steak'':2A,5B ''suitabl'':6B', NULL, true, '2026-11-14 22:15:27.504324+00'),
	('255828c9-35bf-471d-b9b9-d2cd7809c959', 'GRO-BRD-EGG-023', 'Eggs', 'Fresh eggs, a versatile ingredient for baking, cooking, or breakfast.', 2.99, 11.05, 2.53, 9, NULL, 'e2b8f491-6c35-4d7a-9f02-8a1c5e63b790', 'https://cdn.dummyjson.com/product-images/groceries/eggs/thumbnail.webp', '''bake'':8B ''breakfast'':11B ''cook'':9B ''egg'':1A,3B ''fresh'':2B ''groceri'':12C ''ingredi'':6B ''versatil'':5B', NULL, true, '2026-11-14 22:15:27.504324+00'),
	('43f79957-ffff-4660-b161-31a6476849b9', 'GRO-BRD-FOO-022', 'Dog Food', 'Specially formulated dog food designed to provide essential nutrients for your canine companion.', 10.99, 10.27, 4.55, 71, NULL, 'e2b8f491-6c35-4d7a-9f02-8a1c5e63b790', 'https://cdn.dummyjson.com/product-images/groceries/dog-food/thumbnail.webp', '''canin'':14B ''companion'':15B ''design'':7B ''dog'':1A,5B ''essenti'':10B ''food'':2A,6B ''formul'':4B ''groceri'':16C ''nutrient'':11B ''provid'':9B ''special'':3B', NULL, true, '2026-11-14 22:15:27.504324+00'),
	('1f999e9c-10b6-42f2-97c9-ca47eb113803', 'GRO-BRD-CUC-021', 'Cucumber', 'Crisp and hydrating cucumbers, ideal for salads, snacks, or as a refreshing side.', 1.49, 0.16, 4.07, 84, NULL, 'e2b8f491-6c35-4d7a-9f02-8a1c5e63b790', 'https://cdn.dummyjson.com/product-images/groceries/cucumber/thumbnail.webp', '''crisp'':2B ''cucumb'':1A,5B ''groceri'':15C ''hydrat'':4B ''ideal'':6B ''refresh'':13B ''salad'':8B ''side'':14B ''snack'':9B', NULL, true, '2026-11-14 22:15:27.504324+00'),
	('cc240550-b35b-487b-9373-8fc145cbaf23', 'GRO-BRD-COO-020', 'Cooking Oil', 'Versatile cooking oil suitable for frying, sautéing, and various culinary applications.', 4.99, 9.33, 4.8, 10, NULL, 'e2b8f491-6c35-4d7a-9f02-8a1c5e63b790', 'https://cdn.dummyjson.com/product-images/groceries/cooking-oil/thumbnail.webp', '''applic'':13B ''cook'':1A,4B ''culinari'':12B ''fri'':8B ''groceri'':14C ''oil'':2A,5B ''sauté'':9B ''suitabl'':6B ''various'':11B ''versatil'':3B', NULL, true, '2026-11-14 22:15:27.504324+00'),
	('22974f8e-6e1b-4d51-aaef-dbb86954615c', 'GRO-BRD-CHI-019', 'Chicken Meat', 'Fresh and tender chicken meat, suitable for various culinary preparations.', 9.99, 13.7, 3.19, 97, NULL, 'e2b8f491-6c35-4d7a-9f02-8a1c5e63b790', 'https://cdn.dummyjson.com/product-images/groceries/chicken-meat/thumbnail.webp', '''chicken'':1A,6B ''culinari'':11B ''fresh'':3B ''groceri'':13C ''meat'':2A,7B ''prepar'':12B ''suitabl'':8B ''tender'':5B ''various'':10B', NULL, true, '2026-11-14 22:15:27.504324+00'),
	('d203b0f0-af14-4ff1-b52f-546a68595f27', 'GRO-BRD-FOO-018', 'Cat Food', 'Nutritious cat food formulated to meet the dietary needs of your feline friend.', 8.99, 9.58, 3.13, 46, NULL, 'e2b8f491-6c35-4d7a-9f02-8a1c5e63b790', 'https://cdn.dummyjson.com/product-images/groceries/cat-food/thumbnail.webp', '''cat'':1A,4B ''dietari'':10B ''felin'':14B ''food'':2A,5B ''formul'':6B ''friend'':15B ''groceri'':16C ''meet'':8B ''need'':11B ''nutriti'':3B', NULL, true, '2026-11-14 22:15:27.504324+00'),
	('de4b438d-6c7c-49b6-b38e-97d45fd65d49', 'GRO-BRD-BEE-017', 'Beef Steak', 'High-quality beef steak, great for grilling or cooking to your preferred level of doneness.', 12.99, 9.61, 4.47, 86, NULL, 'e2b8f491-6c35-4d7a-9f02-8a1c5e63b790', 'https://cdn.dummyjson.com/product-images/groceries/beef-steak/thumbnail.webp', '''beef'':1A,6B ''cook'':12B ''done'':18B ''great'':8B ''grill'':10B ''groceri'':19C ''high'':4B ''high-qual'':3B ''level'':16B ''prefer'':15B ''qualiti'':5B ''steak'':2A,7B', NULL, true, '2026-11-14 22:15:27.504324+00'),
	('6c8dc39b-0867-496b-b86f-807a123af86e', 'GRO-BRD-APP-016', 'Apple', 'Fresh and crisp apples, perfect for snacking or incorporating into various recipes.', 1.99, 12.62, 4.19, 8, NULL, 'e2b8f491-6c35-4d7a-9f02-8a1c5e63b790', 'https://cdn.dummyjson.com/product-images/groceries/apple/thumbnail.webp', '''appl'':1A,5B ''crisp'':4B ''fresh'':2B ''groceri'':14C ''incorpor'':10B ''perfect'':6B ''recip'':13B ''snack'':8B ''various'':12B', NULL, true, '2026-11-14 22:15:27.504324+00'),
	('c4f47e1f-21c6-4bbb-9b47-43880c0116b5', 'FUR-BAT-WOO-015', 'Wooden Bathroom Sink With Mirror', 'The Wooden Bathroom Sink with Mirror is a unique and stylish addition to your bathroom, featuring a wooden sink countertop and a matching mirror.', 799.99, 8.8, 3.59, 7, 'Bath Trends', '3d91a6e8-2f47-4c05-b823-7e9d1a54f630', 'https://cdn.dummyjson.com/product-images/furniture/wooden-bathroom-sink-with-mirror/thumbnail.webp', '''addit'':17B ''bath'':30C ''bathroom'':2A,8B,20B ''countertop'':25B ''featur'':21B ''furnitur'':32C ''match'':28B ''mirror'':5A,11B,29B ''sink'':3A,9B,24B ''stylish'':16B ''trend'':31C ''uniqu'':14B ''wooden'':1A,7B,23B', NULL, true, '2026-11-14 22:15:27.504324+00'),
	('a637455c-4080-49a5-918c-e595c9d3b575', 'FUR-KNO-KNO-014', 'Knoll Saarinen Executive Conference Chair', 'The Knoll Saarinen Executive Conference Chair is a modern and ergonomic chair, perfect for your office or conference room with its timeless design.', 499.99, 2.01, 4.88, 26, 'Knoll', '3d91a6e8-2f47-4c05-b823-7e9d1a54f630', 'https://cdn.dummyjson.com/product-images/furniture/knoll-saarinen-executive-conference-chair/thumbnail.webp', '''chair'':5A,11B,17B ''confer'':4A,10B,23B ''design'':28B ''ergonom'':16B ''execut'':3A,9B ''furnitur'':30C ''knoll'':1A,7B,29C ''modern'':14B ''offic'':21B ''perfect'':18B ''room'':24B ''saarinen'':2A,8B ''timeless'':27B', NULL, true, '2026-11-14 22:15:27.504324+00'),
	('49eb0570-e9b4-444f-9669-8c06f32640ed', 'FUR-FUR-BED-013', 'Bedside Table African Cherry', 'The Bedside Table in African Cherry is a stylish and functional addition to your bedroom, providing convenient storage space and a touch of elegance.', 299.99, 19.09, 2.87, 64, 'Furniture Co.', '3d91a6e8-2f47-4c05-b823-7e9d1a54f630', 'https://cdn.dummyjson.com/product-images/furniture/bedside-table-african-cherry/thumbnail.webp', '''addit'':16B ''african'':3A,9B ''bedroom'':19B ''bedsid'':1A,6B ''cherri'':4A,10B ''co'':30C ''conveni'':21B ''eleg'':28B ''function'':15B ''furnitur'':29C,31C ''provid'':20B ''space'':23B ''storag'':22B ''stylish'':13B ''tabl'':2A,7B ''touch'':26B', NULL, true, '2026-11-14 22:15:27.504324+00'),
	('d7976e0e-7ce5-4e2d-aac5-b5bea054f31a', 'FUR-ANN-ANN-012', 'Annibale Colombo Sofa', 'The Annibale Colombo Sofa is a sophisticated and comfortable seating option, featuring exquisite design and premium upholstery for your living room.', 2499.99, 14.4, 3.92, 60, 'Annibale Colombo', '3d91a6e8-2f47-4c05-b823-7e9d1a54f630', 'https://cdn.dummyjson.com/product-images/furniture/annibale-colombo-sofa/thumbnail.webp', '''annibal'':1A,5B,25C ''colombo'':2A,6B,26C ''comfort'':12B ''design'':17B ''exquisit'':16B ''featur'':15B ''furnitur'':27C ''live'':23B ''option'':14B ''premium'':19B ''room'':24B ''seat'':13B ''sofa'':3A,7B ''sophist'':10B ''upholsteri'':20B', NULL, true, '2026-11-14 22:15:27.504324+00'),
	('e8ffad32-6297-4bba-935b-83e6ea1626ef', 'FUR-ANN-ANN-011', 'Annibale Colombo Bed', 'The Annibale Colombo Bed is a luxurious and elegant bed frame, crafted with high-quality materials for a comfortable and stylish bedroom.', 1899.99, 8.57, 4.77, 88, 'Annibale Colombo', '3d91a6e8-2f47-4c05-b823-7e9d1a54f630', 'https://cdn.dummyjson.com/product-images/furniture/annibale-colombo-bed/thumbnail.webp', '''annibal'':1A,5B,27C ''bed'':3A,7B,13B ''bedroom'':26B ''colombo'':2A,6B,28C ''comfort'':23B ''craft'':15B ''eleg'':12B ''frame'':14B ''furnitur'':29C ''high'':18B ''high-qual'':17B ''luxuri'':10B ''materi'':20B ''qualiti'':19B ''stylish'':25B', NULL, true, '2026-11-14 22:15:27.504324+00'),
	('6d71bb85-b2ca-4d8d-bc27-e9f334ed41d7', 'FRA-GUC-GUC-010', 'Gucci Bloom Eau de', 'Gucci Bloom by Gucci is a floral and captivating fragrance, with notes of tuberose, jasmine, and Rangoon creeper. It''s a modern and romantic scent.', 79.99, 14.39, 2.74, 91, 'Gucci', 'a5c7e302-91b8-46d4-8f61-2b3e9c70d845', 'https://cdn.dummyjson.com/product-images/fragrances/gucci-bloom-eau-de/thumbnail.webp', '''bloom'':2A,6B ''captiv'':13B ''creeper'':22B ''de'':4A ''eau'':3A ''floral'':11B ''fragranc'':14B,31C ''gucci'':1A,5B,8B,30C ''jasmin'':19B ''modern'':26B ''note'':16B ''rangoon'':21B ''romant'':28B ''scent'':29B ''tuberos'':18B', NULL, true, '2026-11-14 22:15:27.504324+00'),
	('4b2b727d-d003-47ee-9cc1-3463f6782a92', 'FRA-DOL-DOL-009', 'Dolce Shine Eau de', 'Dolce Shine by Dolce & Gabbana is a vibrant and fruity fragrance, featuring notes of mango, jasmine, and blonde woods. It''s a joyful and youthful scent.', 69.99, 0.62, 3.96, 4, 'Dolce & Gabbana', 'a5c7e302-91b8-46d4-8f61-2b3e9c70d845', 'https://cdn.dummyjson.com/product-images/fragrances/dolce-shine-eau-de/thumbnail.webp', '''blond'':22B ''de'':4A ''dolc'':1A,5B,8B,31C ''eau'':3A ''featur'':16B ''fragranc'':15B,33C ''fruiti'':14B ''gabbana'':9B,32C ''jasmin'':20B ''joy'':27B ''mango'':19B ''note'':17B ''scent'':30B ''shine'':2A,6B ''vibrant'':12B ''wood'':23B ''youth'':29B', NULL, true, '2026-11-14 22:15:27.504324+00'),
	('70f7542c-9310-4788-8ba7-106b81e60860', 'FRA-DIO-DIO-008', 'Dior J''adore', 'J''adore by Dior is a luxurious and floral fragrance, known for its blend of ylang-ylang, rose, and jasmine. It embodies femininity and sophistication.', 89.99, 14.72, 3.8, 98, 'Dior', 'a5c7e302-91b8-46d4-8f61-2b3e9c70d845', 'https://cdn.dummyjson.com/product-images/fragrances/dior-j''adore/thumbnail.webp', '''ador'':3A,5B ''blend'':17B ''dior'':1A,7B,30C ''embodi'':26B ''feminin'':27B ''floral'':12B ''fragranc'':13B,31C ''j'':2A,4B ''jasmin'':24B ''known'':14B ''luxuri'':10B ''rose'':22B ''sophist'':29B ''ylang'':20B,21B ''ylang-ylang'':19B', NULL, true, '2026-11-14 22:15:27.504324+00'),
	('77439800-fb9e-4a57-9693-bf050b53a884', 'FRA-CHA-CHA-007', 'Chanel Coco Noir Eau De', 'Coco Noir by Chanel is an elegant and mysterious fragrance, featuring notes of grapefruit, rose, and sandalwood. Perfect for evening occasions.', 129.99, 16.51, 4.26, 58, 'Chanel', 'a5c7e302-91b8-46d4-8f61-2b3e9c70d845', 'https://cdn.dummyjson.com/product-images/fragrances/chanel-coco-noir-eau-de/thumbnail.webp', '''chanel'':1A,9B,27C ''coco'':2A,6B ''de'':5A ''eau'':4A ''eleg'':12B ''even'':25B ''featur'':16B ''fragranc'':15B,28C ''grapefruit'':19B ''mysteri'':14B ''noir'':3A,7B ''note'':17B ''occas'':26B ''perfect'':23B ''rose'':20B ''sandalwood'':22B', NULL, true, '2026-11-14 22:15:27.504324+00'),
	('b1fb6809-6c45-4734-9cc8-86756513eadb', 'FRA-CAL-CAL-006', 'Calvin Klein CK One', 'CK One by Calvin Klein is a classic unisex fragrance, known for its fresh and clean scent. It''s a versatile fragrance suitable for everyday wear.', 49.99, 1.89, 4.37, 29, 'Calvin Klein', 'a5c7e302-91b8-46d4-8f61-2b3e9c70d845', 'https://cdn.dummyjson.com/product-images/fragrances/calvin-klein-ck-one/thumbnail.webp', '''calvin'':1A,8B,31C ''ck'':3A,5B ''classic'':12B ''clean'':20B ''everyday'':29B ''fragranc'':14B,26B,33C ''fresh'':18B ''klein'':2A,9B,32C ''known'':15B ''one'':4A,6B ''scent'':21B ''suitabl'':27B ''unisex'':13B ''versatil'':25B ''wear'':30B', NULL, true, '2026-11-14 22:15:27.504324+00');


ALTER TABLE inventory.products ENABLE TRIGGER ALL;


--
-- Data for Name: product_images; Type: TABLE DATA; Schema: marketplacedemo; Owner: tcero
--

ALTER TABLE inventory.product_images DISABLE TRIGGER ALL;

INSERT INTO inventory.product_images (id, product_id, url) VALUES
	(1, 'fc1f2f0c-7d52-48d9-895a-50f8e20e0bf9', 'https://cdn.dummyjson.com/product-images/beauty/essence-mascara-lash-princess/1.webp'),
	(2, 'e2d666e5-506c-4b51-ba3b-290afd160bf0', 'https://cdn.dummyjson.com/product-images/beauty/eyeshadow-palette-with-mirror/1.webp'),
	(3, '88a7c616-21ff-45bf-aa22-dbc6e3b0acb4', 'https://cdn.dummyjson.com/product-images/beauty/powder-canister/1.webp'),
	(4, '831ad86a-6956-475f-8a61-5a89d72b130d', 'https://cdn.dummyjson.com/product-images/beauty/red-lipstick/1.webp'),
	(5, '2657a1e2-38c5-4d71-9d7b-94664f66189b', 'https://cdn.dummyjson.com/product-images/beauty/red-nail-polish/1.webp'),
	(6, 'b1fb6809-6c45-4734-9cc8-86756513eadb', 'https://cdn.dummyjson.com/product-images/fragrances/calvin-klein-ck-one/1.webp'),
	(7, 'b1fb6809-6c45-4734-9cc8-86756513eadb', 'https://cdn.dummyjson.com/product-images/fragrances/calvin-klein-ck-one/2.webp'),
	(8, 'b1fb6809-6c45-4734-9cc8-86756513eadb', 'https://cdn.dummyjson.com/product-images/fragrances/calvin-klein-ck-one/3.webp'),
	(9, '77439800-fb9e-4a57-9693-bf050b53a884', 'https://cdn.dummyjson.com/product-images/fragrances/chanel-coco-noir-eau-de/1.webp'),
	(10, '77439800-fb9e-4a57-9693-bf050b53a884', 'https://cdn.dummyjson.com/product-images/fragrances/chanel-coco-noir-eau-de/2.webp'),
	(11, '77439800-fb9e-4a57-9693-bf050b53a884', 'https://cdn.dummyjson.com/product-images/fragrances/chanel-coco-noir-eau-de/3.webp'),
	(12, '70f7542c-9310-4788-8ba7-106b81e60860', 'https://cdn.dummyjson.com/product-images/fragrances/dior-j''adore/1.webp'),
	(13, '70f7542c-9310-4788-8ba7-106b81e60860', 'https://cdn.dummyjson.com/product-images/fragrances/dior-j''adore/2.webp'),
	(14, '70f7542c-9310-4788-8ba7-106b81e60860', 'https://cdn.dummyjson.com/product-images/fragrances/dior-j''adore/3.webp'),
	(15, '4b2b727d-d003-47ee-9cc1-3463f6782a92', 'https://cdn.dummyjson.com/product-images/fragrances/dolce-shine-eau-de/1.webp'),
	(16, '4b2b727d-d003-47ee-9cc1-3463f6782a92', 'https://cdn.dummyjson.com/product-images/fragrances/dolce-shine-eau-de/2.webp'),
	(17, '4b2b727d-d003-47ee-9cc1-3463f6782a92', 'https://cdn.dummyjson.com/product-images/fragrances/dolce-shine-eau-de/3.webp'),
	(18, '6d71bb85-b2ca-4d8d-bc27-e9f334ed41d7', 'https://cdn.dummyjson.com/product-images/fragrances/gucci-bloom-eau-de/1.webp'),
	(19, '6d71bb85-b2ca-4d8d-bc27-e9f334ed41d7', 'https://cdn.dummyjson.com/product-images/fragrances/gucci-bloom-eau-de/2.webp'),
	(20, '6d71bb85-b2ca-4d8d-bc27-e9f334ed41d7', 'https://cdn.dummyjson.com/product-images/fragrances/gucci-bloom-eau-de/3.webp'),
	(21, 'e8ffad32-6297-4bba-935b-83e6ea1626ef', 'https://cdn.dummyjson.com/product-images/furniture/annibale-colombo-bed/1.webp'),
	(22, 'e8ffad32-6297-4bba-935b-83e6ea1626ef', 'https://cdn.dummyjson.com/product-images/furniture/annibale-colombo-bed/2.webp'),
	(23, 'e8ffad32-6297-4bba-935b-83e6ea1626ef', 'https://cdn.dummyjson.com/product-images/furniture/annibale-colombo-bed/3.webp'),
	(24, 'd7976e0e-7ce5-4e2d-aac5-b5bea054f31a', 'https://cdn.dummyjson.com/product-images/furniture/annibale-colombo-sofa/1.webp'),
	(25, 'd7976e0e-7ce5-4e2d-aac5-b5bea054f31a', 'https://cdn.dummyjson.com/product-images/furniture/annibale-colombo-sofa/2.webp'),
	(26, 'd7976e0e-7ce5-4e2d-aac5-b5bea054f31a', 'https://cdn.dummyjson.com/product-images/furniture/annibale-colombo-sofa/3.webp'),
	(27, '49eb0570-e9b4-444f-9669-8c06f32640ed', 'https://cdn.dummyjson.com/product-images/furniture/bedside-table-african-cherry/1.webp'),
	(28, '49eb0570-e9b4-444f-9669-8c06f32640ed', 'https://cdn.dummyjson.com/product-images/furniture/bedside-table-african-cherry/2.webp'),
	(29, '49eb0570-e9b4-444f-9669-8c06f32640ed', 'https://cdn.dummyjson.com/product-images/furniture/bedside-table-african-cherry/3.webp'),
	(30, 'a637455c-4080-49a5-918c-e595c9d3b575', 'https://cdn.dummyjson.com/product-images/furniture/knoll-saarinen-executive-conference-chair/1.webp'),
	(31, 'a637455c-4080-49a5-918c-e595c9d3b575', 'https://cdn.dummyjson.com/product-images/furniture/knoll-saarinen-executive-conference-chair/2.webp'),
	(32, 'a637455c-4080-49a5-918c-e595c9d3b575', 'https://cdn.dummyjson.com/product-images/furniture/knoll-saarinen-executive-conference-chair/3.webp'),
	(33, 'c4f47e1f-21c6-4bbb-9b47-43880c0116b5', 'https://cdn.dummyjson.com/product-images/furniture/wooden-bathroom-sink-with-mirror/1.webp'),
	(34, 'c4f47e1f-21c6-4bbb-9b47-43880c0116b5', 'https://cdn.dummyjson.com/product-images/furniture/wooden-bathroom-sink-with-mirror/2.webp'),
	(35, 'c4f47e1f-21c6-4bbb-9b47-43880c0116b5', 'https://cdn.dummyjson.com/product-images/furniture/wooden-bathroom-sink-with-mirror/3.webp'),
	(36, '6c8dc39b-0867-496b-b86f-807a123af86e', 'https://cdn.dummyjson.com/product-images/groceries/apple/1.webp'),
	(37, 'de4b438d-6c7c-49b6-b38e-97d45fd65d49', 'https://cdn.dummyjson.com/product-images/groceries/beef-steak/1.webp'),
	(38, 'd203b0f0-af14-4ff1-b52f-546a68595f27', 'https://cdn.dummyjson.com/product-images/groceries/cat-food/1.webp'),
	(39, '22974f8e-6e1b-4d51-aaef-dbb86954615c', 'https://cdn.dummyjson.com/product-images/groceries/chicken-meat/1.webp'),
	(40, '22974f8e-6e1b-4d51-aaef-dbb86954615c', 'https://cdn.dummyjson.com/product-images/groceries/chicken-meat/2.webp'),
	(41, 'cc240550-b35b-487b-9373-8fc145cbaf23', 'https://cdn.dummyjson.com/product-images/groceries/cooking-oil/1.webp'),
	(42, '1f999e9c-10b6-42f2-97c9-ca47eb113803', 'https://cdn.dummyjson.com/product-images/groceries/cucumber/1.webp'),
	(43, '43f79957-ffff-4660-b161-31a6476849b9', 'https://cdn.dummyjson.com/product-images/groceries/dog-food/1.webp'),
	(44, '255828c9-35bf-471d-b9b9-d2cd7809c959', 'https://cdn.dummyjson.com/product-images/groceries/eggs/1.webp'),
	(45, 'fdec0dca-29d0-459f-8cab-9365bc5b12d2', 'https://cdn.dummyjson.com/product-images/groceries/fish-steak/1.webp'),
	(46, 'f87b350d-53cd-4e00-8caf-69fe499a2e17', 'https://cdn.dummyjson.com/product-images/groceries/green-bell-pepper/1.webp'),
	(47, '44fc661b-849a-4c0e-acde-9d80fd67ca0d', 'https://cdn.dummyjson.com/product-images/groceries/green-chili-pepper/1.webp'),
	(48, 'f7097af4-d9e7-4e8b-b68c-6acda085a959', 'https://cdn.dummyjson.com/product-images/groceries/honey-jar/1.webp'),
	(49, '5fc579d4-d64f-49d6-b215-b2db6cc239ab', 'https://cdn.dummyjson.com/product-images/groceries/ice-cream/1.webp'),
	(50, '5fc579d4-d64f-49d6-b215-b2db6cc239ab', 'https://cdn.dummyjson.com/product-images/groceries/ice-cream/2.webp'),
	(51, '5fc579d4-d64f-49d6-b215-b2db6cc239ab', 'https://cdn.dummyjson.com/product-images/groceries/ice-cream/3.webp'),
	(52, '5fc579d4-d64f-49d6-b215-b2db6cc239ab', 'https://cdn.dummyjson.com/product-images/groceries/ice-cream/4.webp'),
	(53, '76477601-96bc-4fb2-a08c-ee37a874cb31', 'https://cdn.dummyjson.com/product-images/groceries/juice/1.webp'),
	(54, '911ea40b-e358-4b60-92ee-a1b862e2234f', 'https://cdn.dummyjson.com/product-images/groceries/kiwi/1.webp');


ALTER TABLE inventory.product_images ENABLE TRIGGER ALL;

--
-- Data for Name: products_stage; Type: TABLE DATA; Schema: marketplacedemo; Owner: tcero
--

ALTER TABLE scrap.products_stage DISABLE TRIGGER ALL;

INSERT INTO scrap.products_stage (id, sku, title, description, price, discount_percentage, rating, quantity, brand, category, thumbnail, deleted_at) VALUES
	('fc1f2f0c-7d52-48d9-895a-50f8e20e0bf9', 'BEA-ESS-ESS-001', 'Essence Mascara Lash Princess', 'The Essence Mascara Lash Princess is a popular mascara known for its volumizing and lengthening effects. Achieve dramatic lashes with this long-lasting and cruelty-free formula.', 9.99, 10.48, 2.56, 99, 'Essence', 'beauty', 'https://cdn.dummyjson.com/product-images/beauty/essence-mascara-lash-princess/thumbnail.webp', NULL),
	('e2d666e5-506c-4b51-ba3b-290afd160bf0', 'BEA-GLA-EYE-002', 'Eyeshadow Palette with Mirror', 'The Eyeshadow Palette with Mirror offers a versatile range of eyeshadow shades for creating stunning eye looks. With a built-in mirror, it''s convenient for on-the-go makeup application.', 19.99, 18.19, 2.86, 34, 'Glamour Beauty', 'beauty', 'https://cdn.dummyjson.com/product-images/beauty/eyeshadow-palette-with-mirror/thumbnail.webp', NULL),
	('88a7c616-21ff-45bf-aa22-dbc6e3b0acb4', 'BEA-VEL-POW-003', 'Powder Canister', 'The Powder Canister is a finely milled setting powder designed to set makeup and control shine. With a lightweight and translucent formula, it provides a smooth and matte finish.', 14.99, 9.84, 4.64, 89, 'Velvet Touch', 'beauty', 'https://cdn.dummyjson.com/product-images/beauty/powder-canister/thumbnail.webp', NULL),
	('831ad86a-6956-475f-8a61-5a89d72b130d', 'BEA-CHI-LIP-004', 'Red Lipstick', 'The Red Lipstick is a classic and bold choice for adding a pop of color to your lips. With a creamy and pigmented formula, it provides a vibrant and long-lasting finish.', 12.99, 12.16, 4.36, 91, 'Chic Cosmetics', 'beauty', 'https://cdn.dummyjson.com/product-images/beauty/red-lipstick/thumbnail.webp', NULL),
	('2657a1e2-38c5-4d71-9d7b-94664f66189b', 'BEA-NAI-NAI-005', 'Red Nail Polish', 'The Red Nail Polish offers a rich and glossy red hue for vibrant and polished nails. With a quick-drying formula, it provides a salon-quality finish at home.', 8.99, 11.44, 4.32, 79, 'Nail Couture', 'beauty', 'https://cdn.dummyjson.com/product-images/beauty/red-nail-polish/thumbnail.webp', NULL),
	('b1fb6809-6c45-4734-9cc8-86756513eadb', 'FRA-CAL-CAL-006', 'Calvin Klein CK One', 'CK One by Calvin Klein is a classic unisex fragrance, known for its fresh and clean scent. It''s a versatile fragrance suitable for everyday wear.', 49.99, 1.89, 4.37, 29, 'Calvin Klein', 'fragrances', 'https://cdn.dummyjson.com/product-images/fragrances/calvin-klein-ck-one/thumbnail.webp', NULL),
	('77439800-fb9e-4a57-9693-bf050b53a884', 'FRA-CHA-CHA-007', 'Chanel Coco Noir Eau De', 'Coco Noir by Chanel is an elegant and mysterious fragrance, featuring notes of grapefruit, rose, and sandalwood. Perfect for evening occasions.', 129.99, 16.51, 4.26, 58, 'Chanel', 'fragrances', 'https://cdn.dummyjson.com/product-images/fragrances/chanel-coco-noir-eau-de/thumbnail.webp', NULL),
	('70f7542c-9310-4788-8ba7-106b81e60860', 'FRA-DIO-DIO-008', 'Dior J''adore', 'J''adore by Dior is a luxurious and floral fragrance, known for its blend of ylang-ylang, rose, and jasmine. It embodies femininity and sophistication.', 89.99, 14.72, 3.8, 98, 'Dior', 'fragrances', 'https://cdn.dummyjson.com/product-images/fragrances/dior-j''adore/thumbnail.webp', NULL),
	('4b2b727d-d003-47ee-9cc1-3463f6782a92', 'FRA-DOL-DOL-009', 'Dolce Shine Eau de', 'Dolce Shine by Dolce & Gabbana is a vibrant and fruity fragrance, featuring notes of mango, jasmine, and blonde woods. It''s a joyful and youthful scent.', 69.99, 0.62, 3.96, 4, 'Dolce & Gabbana', 'fragrances', 'https://cdn.dummyjson.com/product-images/fragrances/dolce-shine-eau-de/thumbnail.webp', NULL),
	('6d71bb85-b2ca-4d8d-bc27-e9f334ed41d7', 'FRA-GUC-GUC-010', 'Gucci Bloom Eau de', 'Gucci Bloom by Gucci is a floral and captivating fragrance, with notes of tuberose, jasmine, and Rangoon creeper. It''s a modern and romantic scent.', 79.99, 14.39, 2.74, 91, 'Gucci', 'fragrances', 'https://cdn.dummyjson.com/product-images/fragrances/gucci-bloom-eau-de/thumbnail.webp', NULL),
	('e8ffad32-6297-4bba-935b-83e6ea1626ef', 'FUR-ANN-ANN-011', 'Annibale Colombo Bed', 'The Annibale Colombo Bed is a luxurious and elegant bed frame, crafted with high-quality materials for a comfortable and stylish bedroom.', 1899.99, 8.57, 4.77, 88, 'Annibale Colombo', 'furniture', 'https://cdn.dummyjson.com/product-images/furniture/annibale-colombo-bed/thumbnail.webp', NULL),
	('d7976e0e-7ce5-4e2d-aac5-b5bea054f31a', 'FUR-ANN-ANN-012', 'Annibale Colombo Sofa', 'The Annibale Colombo Sofa is a sophisticated and comfortable seating option, featuring exquisite design and premium upholstery for your living room.', 2499.99, 14.4, 3.92, 60, 'Annibale Colombo', 'furniture', 'https://cdn.dummyjson.com/product-images/furniture/annibale-colombo-sofa/thumbnail.webp', NULL),
	('49eb0570-e9b4-444f-9669-8c06f32640ed', 'FUR-FUR-BED-013', 'Bedside Table African Cherry', 'The Bedside Table in African Cherry is a stylish and functional addition to your bedroom, providing convenient storage space and a touch of elegance.', 299.99, 19.09, 2.87, 64, 'Furniture Co.', 'furniture', 'https://cdn.dummyjson.com/product-images/furniture/bedside-table-african-cherry/thumbnail.webp', NULL),
	('a637455c-4080-49a5-918c-e595c9d3b575', 'FUR-KNO-KNO-014', 'Knoll Saarinen Executive Conference Chair', 'The Knoll Saarinen Executive Conference Chair is a modern and ergonomic chair, perfect for your office or conference room with its timeless design.', 499.99, 2.01, 4.88, 26, 'Knoll', 'furniture', 'https://cdn.dummyjson.com/product-images/furniture/knoll-saarinen-executive-conference-chair/thumbnail.webp', NULL),
	('c4f47e1f-21c6-4bbb-9b47-43880c0116b5', 'FUR-BAT-WOO-015', 'Wooden Bathroom Sink With Mirror', 'The Wooden Bathroom Sink with Mirror is a unique and stylish addition to your bathroom, featuring a wooden sink countertop and a matching mirror.', 799.99, 8.8, 3.59, 7, 'Bath Trends', 'furniture', 'https://cdn.dummyjson.com/product-images/furniture/wooden-bathroom-sink-with-mirror/thumbnail.webp', NULL),
	('6c8dc39b-0867-496b-b86f-807a123af86e', 'GRO-BRD-APP-016', 'Apple', 'Fresh and crisp apples, perfect for snacking or incorporating into various recipes.', 1.99, 12.62, 4.19, 8, NULL, 'groceries', 'https://cdn.dummyjson.com/product-images/groceries/apple/thumbnail.webp', NULL),
	('de4b438d-6c7c-49b6-b38e-97d45fd65d49', 'GRO-BRD-BEE-017', 'Beef Steak', 'High-quality beef steak, great for grilling or cooking to your preferred level of doneness.', 12.99, 9.61, 4.47, 86, NULL, 'groceries', 'https://cdn.dummyjson.com/product-images/groceries/beef-steak/thumbnail.webp', NULL),
	('d203b0f0-af14-4ff1-b52f-546a68595f27', 'GRO-BRD-FOO-018', 'Cat Food', 'Nutritious cat food formulated to meet the dietary needs of your feline friend.', 8.99, 9.58, 3.13, 46, NULL, 'groceries', 'https://cdn.dummyjson.com/product-images/groceries/cat-food/thumbnail.webp', NULL),
	('22974f8e-6e1b-4d51-aaef-dbb86954615c', 'GRO-BRD-CHI-019', 'Chicken Meat', 'Fresh and tender chicken meat, suitable for various culinary preparations.', 9.99, 13.7, 3.19, 97, NULL, 'groceries', 'https://cdn.dummyjson.com/product-images/groceries/chicken-meat/thumbnail.webp', NULL),
	('cc240550-b35b-487b-9373-8fc145cbaf23', 'GRO-BRD-COO-020', 'Cooking Oil', 'Versatile cooking oil suitable for frying, sautéing, and various culinary applications.', 4.99, 9.33, 4.8, 10, NULL, 'groceries', 'https://cdn.dummyjson.com/product-images/groceries/cooking-oil/thumbnail.webp', NULL),
	('1f999e9c-10b6-42f2-97c9-ca47eb113803', 'GRO-BRD-CUC-021', 'Cucumber', 'Crisp and hydrating cucumbers, ideal for salads, snacks, or as a refreshing side.', 1.49, 0.16, 4.07, 84, NULL, 'groceries', 'https://cdn.dummyjson.com/product-images/groceries/cucumber/thumbnail.webp', NULL),
	('43f79957-ffff-4660-b161-31a6476849b9', 'GRO-BRD-FOO-022', 'Dog Food', 'Specially formulated dog food designed to provide essential nutrients for your canine companion.', 10.99, 10.27, 4.55, 71, NULL, 'groceries', 'https://cdn.dummyjson.com/product-images/groceries/dog-food/thumbnail.webp', NULL),
	('255828c9-35bf-471d-b9b9-d2cd7809c959', 'GRO-BRD-EGG-023', 'Eggs', 'Fresh eggs, a versatile ingredient for baking, cooking, or breakfast.', 2.99, 11.05, 2.53, 9, NULL, 'groceries', 'https://cdn.dummyjson.com/product-images/groceries/eggs/thumbnail.webp', NULL),
	('fdec0dca-29d0-459f-8cab-9365bc5b12d2', 'GRO-BRD-FIS-024', 'Fish Steak', 'Quality fish steak, suitable for grilling, baking, or pan-searing.', 14.99, 4.23, 3.78, 74, NULL, 'groceries', 'https://cdn.dummyjson.com/product-images/groceries/fish-steak/thumbnail.webp', NULL),
	('f87b350d-53cd-4e00-8caf-69fe499a2e17', 'GRO-BRD-GRE-025', 'Green Bell Pepper', 'Fresh and vibrant green bell pepper, perfect for adding color and flavor to your dishes.', 1.29, 0.16, 3.25, 33, NULL, 'groceries', 'https://cdn.dummyjson.com/product-images/groceries/green-bell-pepper/thumbnail.webp', NULL),
	('44fc661b-849a-4c0e-acde-9d80fd67ca0d', 'GRO-BRD-GRE-026', 'Green Chili Pepper', 'Spicy green chili pepper, ideal for adding heat to your favorite recipes.', 0.99, 1, 3.66, 3, NULL, 'groceries', 'https://cdn.dummyjson.com/product-images/groceries/green-chili-pepper/thumbnail.webp', NULL),
	('f7097af4-d9e7-4e8b-b68c-6acda085a959', 'GRO-BRD-HON-027', 'Honey Jar', 'Pure and natural honey in a convenient jar, perfect for sweetening beverages or drizzling over food.', 6.99, 14.4, 3.97, 34, NULL, 'groceries', 'https://cdn.dummyjson.com/product-images/groceries/honey-jar/thumbnail.webp', NULL),
	('5fc579d4-d64f-49d6-b215-b2db6cc239ab', 'GRO-BRD-CRE-028', 'Ice Cream', 'Creamy and delicious ice cream, available in various flavors for a delightful treat.', 5.49, 8.69, 3.39, 27, NULL, 'groceries', 'https://cdn.dummyjson.com/product-images/groceries/ice-cream/thumbnail.webp', NULL),
	('76477601-96bc-4fb2-a08c-ee37a874cb31', 'GRO-BRD-JUI-029', 'Juice', 'Refreshing fruit juice, packed with vitamins and great for staying hydrated.', 3.99, 12.06, 3.94, 50, NULL, 'groceries', 'https://cdn.dummyjson.com/product-images/groceries/juice/thumbnail.webp', NULL),
	('911ea40b-e358-4b60-92ee-a1b862e2234f', 'GRO-BRD-KIW-030', 'Kiwi', 'Nutrient-rich kiwi, perfect for snacking or adding a tropical twist to your dishes.', 2.49, 15.22, 4.93, 99, NULL, 'groceries', 'https://cdn.dummyjson.com/product-images/groceries/kiwi/thumbnail.webp', NULL);


ALTER TABLE scrap.products_stage ENABLE TRIGGER ALL;

--
-- Data for Name: product_images_stage; Type: TABLE DATA; Schema: marketplacedemo; Owner: tcero
--


ALTER TABLE scrap.product_images_stage DISABLE TRIGGER ALL;

INSERT INTO scrap.product_images_stage (id, product_id, url) VALUES
	(1, 'fc1f2f0c-7d52-48d9-895a-50f8e20e0bf9', 'https://cdn.dummyjson.com/product-images/beauty/essence-mascara-lash-princess/1.webp'),
	(2, 'e2d666e5-506c-4b51-ba3b-290afd160bf0', 'https://cdn.dummyjson.com/product-images/beauty/eyeshadow-palette-with-mirror/1.webp'),
	(3, '88a7c616-21ff-45bf-aa22-dbc6e3b0acb4', 'https://cdn.dummyjson.com/product-images/beauty/powder-canister/1.webp'),
	(4, '831ad86a-6956-475f-8a61-5a89d72b130d', 'https://cdn.dummyjson.com/product-images/beauty/red-lipstick/1.webp'),
	(5, '2657a1e2-38c5-4d71-9d7b-94664f66189b', 'https://cdn.dummyjson.com/product-images/beauty/red-nail-polish/1.webp'),
	(6, 'b1fb6809-6c45-4734-9cc8-86756513eadb', 'https://cdn.dummyjson.com/product-images/fragrances/calvin-klein-ck-one/1.webp'),
	(7, 'b1fb6809-6c45-4734-9cc8-86756513eadb', 'https://cdn.dummyjson.com/product-images/fragrances/calvin-klein-ck-one/2.webp'),
	(8, 'b1fb6809-6c45-4734-9cc8-86756513eadb', 'https://cdn.dummyjson.com/product-images/fragrances/calvin-klein-ck-one/3.webp'),
	(9, '77439800-fb9e-4a57-9693-bf050b53a884', 'https://cdn.dummyjson.com/product-images/fragrances/chanel-coco-noir-eau-de/1.webp'),
	(10, '77439800-fb9e-4a57-9693-bf050b53a884', 'https://cdn.dummyjson.com/product-images/fragrances/chanel-coco-noir-eau-de/2.webp'),
	(11, '77439800-fb9e-4a57-9693-bf050b53a884', 'https://cdn.dummyjson.com/product-images/fragrances/chanel-coco-noir-eau-de/3.webp'),
	(12, '70f7542c-9310-4788-8ba7-106b81e60860', 'https://cdn.dummyjson.com/product-images/fragrances/dior-j''adore/1.webp'),
	(13, '70f7542c-9310-4788-8ba7-106b81e60860', 'https://cdn.dummyjson.com/product-images/fragrances/dior-j''adore/2.webp'),
	(14, '70f7542c-9310-4788-8ba7-106b81e60860', 'https://cdn.dummyjson.com/product-images/fragrances/dior-j''adore/3.webp'),
	(15, '4b2b727d-d003-47ee-9cc1-3463f6782a92', 'https://cdn.dummyjson.com/product-images/fragrances/dolce-shine-eau-de/1.webp'),
	(16, '4b2b727d-d003-47ee-9cc1-3463f6782a92', 'https://cdn.dummyjson.com/product-images/fragrances/dolce-shine-eau-de/2.webp'),
	(17, '4b2b727d-d003-47ee-9cc1-3463f6782a92', 'https://cdn.dummyjson.com/product-images/fragrances/dolce-shine-eau-de/3.webp'),
	(18, '6d71bb85-b2ca-4d8d-bc27-e9f334ed41d7', 'https://cdn.dummyjson.com/product-images/fragrances/gucci-bloom-eau-de/1.webp'),
	(19, '6d71bb85-b2ca-4d8d-bc27-e9f334ed41d7', 'https://cdn.dummyjson.com/product-images/fragrances/gucci-bloom-eau-de/2.webp'),
	(20, '6d71bb85-b2ca-4d8d-bc27-e9f334ed41d7', 'https://cdn.dummyjson.com/product-images/fragrances/gucci-bloom-eau-de/3.webp'),
	(21, 'e8ffad32-6297-4bba-935b-83e6ea1626ef', 'https://cdn.dummyjson.com/product-images/furniture/annibale-colombo-bed/1.webp'),
	(22, 'e8ffad32-6297-4bba-935b-83e6ea1626ef', 'https://cdn.dummyjson.com/product-images/furniture/annibale-colombo-bed/2.webp'),
	(23, 'e8ffad32-6297-4bba-935b-83e6ea1626ef', 'https://cdn.dummyjson.com/product-images/furniture/annibale-colombo-bed/3.webp'),
	(24, 'd7976e0e-7ce5-4e2d-aac5-b5bea054f31a', 'https://cdn.dummyjson.com/product-images/furniture/annibale-colombo-sofa/1.webp'),
	(25, 'd7976e0e-7ce5-4e2d-aac5-b5bea054f31a', 'https://cdn.dummyjson.com/product-images/furniture/annibale-colombo-sofa/2.webp'),
	(26, 'd7976e0e-7ce5-4e2d-aac5-b5bea054f31a', 'https://cdn.dummyjson.com/product-images/furniture/annibale-colombo-sofa/3.webp'),
	(27, '49eb0570-e9b4-444f-9669-8c06f32640ed', 'https://cdn.dummyjson.com/product-images/furniture/bedside-table-african-cherry/1.webp'),
	(28, '49eb0570-e9b4-444f-9669-8c06f32640ed', 'https://cdn.dummyjson.com/product-images/furniture/bedside-table-african-cherry/2.webp'),
	(29, '49eb0570-e9b4-444f-9669-8c06f32640ed', 'https://cdn.dummyjson.com/product-images/furniture/bedside-table-african-cherry/3.webp'),
	(30, 'a637455c-4080-49a5-918c-e595c9d3b575', 'https://cdn.dummyjson.com/product-images/furniture/knoll-saarinen-executive-conference-chair/1.webp'),
	(31, 'a637455c-4080-49a5-918c-e595c9d3b575', 'https://cdn.dummyjson.com/product-images/furniture/knoll-saarinen-executive-conference-chair/2.webp'),
	(32, 'a637455c-4080-49a5-918c-e595c9d3b575', 'https://cdn.dummyjson.com/product-images/furniture/knoll-saarinen-executive-conference-chair/3.webp'),
	(33, 'c4f47e1f-21c6-4bbb-9b47-43880c0116b5', 'https://cdn.dummyjson.com/product-images/furniture/wooden-bathroom-sink-with-mirror/1.webp'),
	(34, 'c4f47e1f-21c6-4bbb-9b47-43880c0116b5', 'https://cdn.dummyjson.com/product-images/furniture/wooden-bathroom-sink-with-mirror/2.webp'),
	(35, 'c4f47e1f-21c6-4bbb-9b47-43880c0116b5', 'https://cdn.dummyjson.com/product-images/furniture/wooden-bathroom-sink-with-mirror/3.webp'),
	(36, '6c8dc39b-0867-496b-b86f-807a123af86e', 'https://cdn.dummyjson.com/product-images/groceries/apple/1.webp'),
	(37, 'de4b438d-6c7c-49b6-b38e-97d45fd65d49', 'https://cdn.dummyjson.com/product-images/groceries/beef-steak/1.webp'),
	(38, 'd203b0f0-af14-4ff1-b52f-546a68595f27', 'https://cdn.dummyjson.com/product-images/groceries/cat-food/1.webp'),
	(39, '22974f8e-6e1b-4d51-aaef-dbb86954615c', 'https://cdn.dummyjson.com/product-images/groceries/chicken-meat/1.webp'),
	(40, '22974f8e-6e1b-4d51-aaef-dbb86954615c', 'https://cdn.dummyjson.com/product-images/groceries/chicken-meat/2.webp'),
	(41, 'cc240550-b35b-487b-9373-8fc145cbaf23', 'https://cdn.dummyjson.com/product-images/groceries/cooking-oil/1.webp'),
	(42, '1f999e9c-10b6-42f2-97c9-ca47eb113803', 'https://cdn.dummyjson.com/product-images/groceries/cucumber/1.webp'),
	(43, '43f79957-ffff-4660-b161-31a6476849b9', 'https://cdn.dummyjson.com/product-images/groceries/dog-food/1.webp'),
	(44, '255828c9-35bf-471d-b9b9-d2cd7809c959', 'https://cdn.dummyjson.com/product-images/groceries/eggs/1.webp'),
	(45, 'fdec0dca-29d0-459f-8cab-9365bc5b12d2', 'https://cdn.dummyjson.com/product-images/groceries/fish-steak/1.webp'),
	(46, 'f87b350d-53cd-4e00-8caf-69fe499a2e17', 'https://cdn.dummyjson.com/product-images/groceries/green-bell-pepper/1.webp'),
	(47, '44fc661b-849a-4c0e-acde-9d80fd67ca0d', 'https://cdn.dummyjson.com/product-images/groceries/green-chili-pepper/1.webp'),
	(48, 'f7097af4-d9e7-4e8b-b68c-6acda085a959', 'https://cdn.dummyjson.com/product-images/groceries/honey-jar/1.webp'),
	(49, '5fc579d4-d64f-49d6-b215-b2db6cc239ab', 'https://cdn.dummyjson.com/product-images/groceries/ice-cream/1.webp'),
	(50, '5fc579d4-d64f-49d6-b215-b2db6cc239ab', 'https://cdn.dummyjson.com/product-images/groceries/ice-cream/2.webp'),
	(51, '5fc579d4-d64f-49d6-b215-b2db6cc239ab', 'https://cdn.dummyjson.com/product-images/groceries/ice-cream/3.webp'),
	(52, '5fc579d4-d64f-49d6-b215-b2db6cc239ab', 'https://cdn.dummyjson.com/product-images/groceries/ice-cream/4.webp'),
	(53, '76477601-96bc-4fb2-a08c-ee37a874cb31', 'https://cdn.dummyjson.com/product-images/groceries/juice/1.webp'),
	(54, '911ea40b-e358-4b60-92ee-a1b862e2234f', 'https://cdn.dummyjson.com/product-images/groceries/kiwi/1.webp');


ALTER TABLE scrap.product_images_stage ENABLE TRIGGER ALL;

