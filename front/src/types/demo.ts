import { ComponentPropsWithoutRef } from "react"

export type IdxProductProps = {
  idItem:number,
}

export type Product = {
  id: number,
  title: string,
  description: string,
  price: number,
  discountPercentage: number,
  rating: number,
  stock: number,
  brand: string,
  category: string,
  thumbnail: string
}

export type productImage = {
  id: number,
  productId: number,
  url: string
}

export type ItemProductProps = {
  item:number
}

export const sortOptions = [
  {
    value: "featured",
    label: "Featured",
    compare: null,
  },
  {
    value: "price-ascending",
    label: "Price: low to high",
    compare: (a: Product, b: Product) => a.price - b.price,
  },
  {
    value: "price-descending",
    label: "Price: high to low",
    compare: (a: Product, b: Product) => b.price - a.price,
  },
  {
    value: "name-ascending",
    label: "Name: A to Z",
    compare: (a: Product, b: Product) => a.title.localeCompare(b.title),
  },
  {
    value: "name-descending",
    label: "Name: Z to A",
    compare: (a: Product, b: Product) => b.title.localeCompare(a.title),
  },
] as const;

export type SortBy = (typeof sortOptions)[number]["value"];

export type PosteoBase = {
  id:string
  texto:string
  userId:string
}

export type PosteoCreateBase = {
  id:string
  texto:string
  productId:number
}

export type PosteoCreate = PosteoCreateBase & {
  meta:Meta
}

export type Meta = {
  hashtags: Categories[]
}
export type Categories = {
  nombre: string
}

export type EditorProps = {
  posteo:PosteoCreate | null
  productId:string
  onChangePosteo:(posteo:PosteoCreate) => void
} & ComponentPropsWithoutRef<'div'>