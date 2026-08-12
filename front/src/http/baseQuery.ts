import { BaseQueryFn, FetchArgs, fetchBaseQuery, FetchBaseQueryError } from "@reduxjs/toolkit/query"
import type { RootState } from "@/store/store"
import { type Product } from "@/types/demo"

export const baseQuery = fetchBaseQuery({
  baseUrl: "/bff",
  prepareHeaders: (headers, { getState }) => {
      headers.set("Authorization", `Bearer ${sessionStorage.getItem("Access_Token")}`)
    return headers
  }
})

export const baseHydraQuery = fetchBaseQuery({
  baseUrl: "/hydra",
  credentials: "include",
})

export const fakeBaseQueryWithRefresh: BaseQueryFn<
  FetchArgs,
  unknown,
  FetchBaseQueryError
>  = async (args, api, extraOptions) => {
  const url = typeof args === 'string' ? args : args.url
  if(url === '/getAuthentication') {
      return { data: authentication }
  }
  if(url === '/usuario/getProducts') {
    return { data: products};
  }
  if(url === '/usuario/getCategories') {
    return { data: categories };
  }
  if(url === '/usuario/getRecommendations') {
    return { data: recomendations };
  }
  if(url === '/usuario/getPosteos') {
    return { data: [] }
  }
  if(url === '/usuario/createPost' && args.method === 'POST') {
    return { data: { message: 'ok' } }
  }
  if (url === '/usuario/getProduct?product=1') {
    return {
      data: {
        id: 1,
        title: "Essence Mascara Lash Princess",
        description:
          "The Essence Mascara Lash Princess is a popular mascara known for its volumizing and lengthening effects. Achieve dramatic lashes with this long-lasting and cruelty-free formula.",
        price: 9.99,
        discountPercentage: 10.48,
        rating: 2.56,
        stock: 99,
        brand: "Essence",
        category: "beauty",
        thumbnail:
          "https://cdn.dummyjson.com/product-images/beauty/essence-mascara-lash-princess/thumbnail.webp",
      },
    }
  }
  if(url === '/usuario/getProduct?product=2') {
    return {
      data: {
        "id": 2,
        "title": "Eyeshadow Palette with Mirror",
        "description": "The Eyeshadow Palette with Mirror offers a versatile range of eyeshadow shades for creating stunning eye looks. With a built-in mirror, it's convenient for on-the-go makeup application.",
        "price": 19.99,
        "discountPercentage": 18.19,
        "rating": 2.86,
        "stock": 34,
        "brand": "Glamour Beauty",
        "category": "beauty",
        "thumbnail": "https://cdn.dummyjson.com/product-images/beauty/eyeshadow-palette-with-mirror/thumbnail.webp"
      }
    }
  }
  if(url === '/usuario/getProduct?product=3') {
    return {
      data: {
        "id": 3,
        "title": "Powder Canister",
        "description": "The Powder Canister is a finely milled setting powder designed to set makeup and control shine. With a lightweight and translucent formula, it provides a smooth and matte finish.",
        "price": 14.99,
        "discountPercentage": 9.84,
        "rating": 4.64,
        "stock": 89,
        "brand": "Velvet Touch",
        "category": "beauty",
        "thumbnail": "https://cdn.dummyjson.com/product-images/beauty/powder-canister/thumbnail.webp"
      }
    }
  }
  if(url === '/usuario/getProduct?product=4') {
    return {
      data: {
        "id": 4,
        "title": "Red Lipstick",
        "description": "The Red Lipstick is a classic and bold choice for adding a pop of color to your lips. With a creamy and pigmented formula, it provides a vibrant and long-lasting finish.",
        "price": 12.99,
        "discountPercentage": 12.16,
        "rating": 4.36,
        "stock": 91,
        "brand": "Chic Cosmetics",
        "category": "beauty",
        "thumbnail": "https://cdn.dummyjson.com/product-images/beauty/red-lipstick/thumbnail.webp"
      }
    }
  }
  if(url === '/usuario/getProduct?product=5') {
    return {
      data: {
        "id": 5,
        "title": "Red Nail Polish",
        "description": "The Red Nail Polish offers a rich and glossy red hue for vibrant and polished nails. With a quick-drying formula, it provides a salon-quality finish at home.",
        "price": 8.99,
        "discountPercentage": 11.44,
        "rating": 4.32,
        "stock": 79,
        "brand": "Nail Couture",
        "category": "beauty",
        "thumbnail": "https://cdn.dummyjson.com/product-images/beauty/red-nail-polish/thumbnail.webp"
      }
    }
  }
  if(url === '/usuario/getProduct?product=6') {
    return {
      data: {
        "id": 6,
        "title": "Calvin Klein CK One",
        "description": "CK One by Calvin Klein is a classic unisex fragrance, known for its fresh and clean scent. It's a versatile fragrance suitable for everyday wear.",
        "price": 49.99,
        "discountPercentage": 1.89,
        "rating": 4.37,
        "stock": 29,
        "brand": "Calvin Klein",
        "category": "fragrances",
        "thumbnail": "https://cdn.dummyjson.com/product-images/fragrances/calvin-klein-ck-one/thumbnail.webp"
      }
    }
  }
  if(url === '/usuario/getProduct?product=7') {
    return {
      data: {
        "id": 7,
        "title": "Chanel Coco Noir Eau De",
        "description": "Coco Noir by Chanel is an elegant and mysterious fragrance, featuring notes of grapefruit, rose, and sandalwood. Perfect for evening occasions.",
        "price": 129.99,
        "discountPercentage": 16.51,
        "rating": 4.26,
        "stock": 58,
        "brand": "Chanel",
        "category": "fragrances",
        "thumbnail": "https://cdn.dummyjson.com/product-images/fragrances/chanel-coco-noir-eau-de/thumbnail.webp"
      }
    }
  }
  if(url === '/usuario/getProduct?product=8') {
    return {
      data: {
        "id": 8,
        "title": "Dior J'adore",
        "description": "J'adore by Dior is a luxurious and floral fragrance, known for its blend of ylang-ylang, rose, and jasmine. It embodies femininity and sophistication.",
        "price": 89.99,
        "discountPercentage": 14.72,
        "rating": 3.8,
        "stock": 98,
        "brand": "Dior",
        "category": "fragrances",
        "thumbnail": "https://cdn.dummyjson.com/product-images/fragrances/dior-j'adore/thumbnail.webp"
      }
    }
  }
  if(url === '/usuario/getProduct?product=9') {
    return {
      data: {
        "id": 9,
        "title": "Dolce Shine Eau de",
        "description": "Dolce Shine by Dolce \u0026 Gabbana is a vibrant and fruity fragrance, featuring notes of mango, jasmine, and blonde woods. It's a joyful and youthful scent.",
        "price": 69.99,
        "discountPercentage": 0.62,
        "rating": 3.96,
        "stock": 4,
        "brand": "Dolce \u0026 Gabbana",
        "category": "fragrances",
        "thumbnail": "https://cdn.dummyjson.com/product-images/fragrances/dolce-shine-eau-de/thumbnail.webp"
      }
    }
  }
  if(url === '/usuario/getProduct?product=10') {
    return {
      data: {
        "id": 10,
        "title": "Gucci Bloom Eau de",
        "description": "Gucci Bloom by Gucci is a floral and captivating fragrance, with notes of tuberose, jasmine, and Rangoon creeper. It's a modern and romantic scent.",
        "price": 79.99,
        "discountPercentage": 14.39,
        "rating": 2.74,
        "stock": 91,
        "brand": "Gucci",
        "category": "fragrances",
        "thumbnail": "https://cdn.dummyjson.com/product-images/fragrances/gucci-bloom-eau-de/thumbnail.webp"
      }
    }
  }
  if(url === '/usuario/getProduct?product=11') {
    return {
      data: {
        "id": 11,
        "title": "Hugo Boss Bottled Tonic",
        "description": "Hugo Boss Bottled Tonic is a fresh and vibrant fragrance, with notes of bergamot, cardamom, and vetiver. It's a modern and confident scent.",
        "price": 59.99,
        "discountPercentage": 12.5,
        "rating": 4.1,
        "stock": 67,
        "brand": "Hugo Boss",
        "category": "fragrances",
        "thumbnail": "https://cdn.dummyjson.com/product-images/fragrances/hugo-boss-bottled-tonic/thumbnail.webp"
      }
    }
  }
  if(url === '/usuario/getProduct?product=12') {
    return {
      data: {
        "id": 12,
        "title": "Light Blue Eau de",
        "description": "Light Blue by Dolce \u0026 Gabbana is a fresh and feminine fragrance, featuring notes of lemon, jasmine, and musk. It's a timeless and elegant scent.",
        "price": 49.99,
        "discountPercentage": 10.2,
        "rating": 4.0,
        "stock": 85,
        "brand": "Dolce \u0026 Gabbana",
        "category": "fragrances",
        "thumbnail": "https://cdn.dummyjson.com/product-images/fragrances/light-blue-eau-de/thumbnail.webp"
      }
    }
  }
  if(url === '/usuario/getProduct?product=13') {
    return {
      data: {
        "id": 13,
        "title": "Miss Dior Eau de",
        "description": "Miss Dior by Dior is a fresh and feminine fragrance, with notes of orange blossom, rose, and vanilla. It's a classic and elegant scent.",
        "price": 69.99,
        "discountPercentage": 15.0,
        "rating": 4.3,
        "stock": 78,
        "brand": "Dior",
        "category": "fragrances",
        "thumbnail": "https://cdn.dummyjson.com/product-images/fragrances/miss-dior-eau-de/thumbnail.webp"
      }
    }
  }
  if(url === '/usuario/getProduct?product=14') {
    return {
      data: {
        "id": 14,
        "title": "Opium Eau de",
        "description": "Opium by Yves Saint Laurent is a bold and seductive fragrance, featuring notes of coffee, vanilla, and sandalwood. It's a modern and captivating scent.",
        "price": 79.99,
        "discountPercentage": 12.0,
        "rating": 4.2,
        "stock": 52,
        "brand": "Yves Saint Laurent",
        "category": "fragrances",
        "thumbnail": "https://cdn.dummyjson.com/product-images/fragrances/opium-eau-de/thumbnail.webp"
      }
    }
  }
  return { data: { message: 'ok' } }
}

const products:Product[] = []

const categories = ["beauty","groceries","furniture","fragrances"]

const authentication = {
    "aud": [],
    "client_id": "657d0bb0-2314-4c8b-b649-1525af797d72",
    "exp": 1776168934,
    "ext": {
        "email": "",
        "family_name": "",
        "given_name": "",
        "id": "123e4567-e89b-12d3-a456-426614174000",
        "name": "leonardo",
        "picture": "https://lh3.googleusercontent.com/ogw/AF2bZyi3jj_bNU_05Z6ldnz_Xs-ofC1aGtz2o_GOyDO5RkSLHvk=s32-c-mo",
        "verified_email": false
    },
    "iat": 1776108934,
    "iss": "https://localhost/hydra",
    "jti": "2a9e8e58-9c6b-4956-86cd-b3625533ad99",
    "nbf": 1776108934,
    "scp": [
        "openid",
        "offline",
        "mediamtx:stream"
    ],
    "sub": "123e4567-e89b-12d3-a456-426614174000"
}

const recomendations = [1,23,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,24,25,26,27,28,29,30]

