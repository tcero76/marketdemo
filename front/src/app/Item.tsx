'use client'

import { FC } from 'react'
import { Button } from "@/components/ui/button"
import {
  Card,
  CardDescription,
  CardTitle,
} from "@/components/ui/card"
import Link from 'next/dist/client/link';
import { useGetProductQuery } from '@/http/api';
import { ItemProductProps, Product } from '@/types/demo';
import { useAppDispatch } from '@/store/hooks';
import { addItem } from "@/store/CartSlice";
import { ShoppingCart } from 'lucide-react';

const Item:FC<ItemProductProps> = ({ item }:ItemProductProps) => {
  const dispatch = useAppDispatch();
  const productId = item;
  const { data } = useGetProductQuery(
    productId,
    { skip: !productId }
  );
  const onClickCart = (product: Product | undefined) => {
  if (!product) return;
  dispatch(
    addItem({
      productId: product?.id,
      title: product?.title,
      price: product?.price,
      quantity: 1,
      thumbnail: product?.thumbnail,
    })
  );
};
  return (
    <Card className="w-full h-full flex flex-col">
      <Link href={`/product/${data?.id}`}>
        <img
          src={data?.thumbnail}
          alt="Event cover"
          className="relative z-20 aspect-video w-full object-cover brightness-60 grayscale dark:brightness-40"
        />
      </Link>
      <div className="flex flex-col h-full space-y-4 p-4">
        <div className="space-y-1">
          <CardTitle className="text-base font-semibold leading-tight">
            {data?.title}
          </CardTitle>
          <CardDescription className="text-sm text-muted-foreground">
            {data?.category}
          </CardDescription>
        </div>
        <div className="flex-1">
          <p className="text-sm text-muted-foreground line-clamp-3">
            {data?.description}
          </p>
        </div>
        <div className="pt-2">
          <Button variant="outline" size="sm" className="w-full" onClick={() => onClickCart(data)}>
               <ShoppingCart size={24} strokeWidth={1.8} />
          </Button>
        </div>
      </div>
    </Card>
  )
}

export default Item