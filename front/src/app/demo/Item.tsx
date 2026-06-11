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
import { ItemProductProps } from '@/types/demo';

const Item:FC<ItemProductProps> = ({ item }:ItemProductProps) => {
const productId = item;
const { data } = useGetProductQuery(
  productId,
  { skip: !productId }
);
  return (
    <Card className="w-full h-full flex flex-col">
      <img
        src={data?.thumbnail}
        alt="Event cover"
        className="relative z-20 aspect-video w-full object-cover brightness-60 grayscale dark:brightness-40"
      />
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
          <Button asChild variant="outline" size="sm" className="w-full">
            <Link href={`/demo/product/${data?.id}`}>
              {data?.stock}
            </Link>
          </Button>
        </div>
      </div>
    </Card>
  )
}

export default Item