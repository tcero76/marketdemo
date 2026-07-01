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