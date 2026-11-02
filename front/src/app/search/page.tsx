'use client';

import { useUIContext } from '@/context/UIContext';
import { useSearchParams } from 'next/navigation';
import {  useEffect, useMemo } from 'react';
import { Virtuoso } from 'react-virtuoso';
import Row from './Row';
import { useSearchProductsQuery } from '@/http/api';
import { getProductRows } from '@/utils/products-rows';
import { type SortBy } from '@/types/demo';

export default function Page() {
  const searchParams = useSearchParams();
  const uiContext = useUIContext()
  const hashtag = searchParams.get("hashtag") ?? "";
  const sortBy = searchParams.get("sort_by") as SortBy;
  const text = searchParams.get("text") ?? "";
  const { data , isLoading } = useSearchProductsQuery({
    mention:"",
    hashtag,
    text:[text]
  }, 
    { skip: !text && !hashtag});
  const rows = useMemo(() => {
    return getProductRows(data, sortBy)
  }, [data,sortBy]);
  useEffect(() => {
    if (isLoading) {
        uiContext.showSpinner();
    } else {
        uiContext.hideSpinner();
    }
    return () => {
        uiContext.hideSpinner();
    };
  }, [isLoading, uiContext]);
  return (
    <div className="h-full w-full">
        <Virtuoso
          data={rows}
          itemContent={(index, item) => <Row key={item.key} row={item.row}/>}
          computeItemKey={(index) => index}
        >
        </Virtuoso>
    </div>
    )
}


