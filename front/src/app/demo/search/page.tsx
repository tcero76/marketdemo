'use client';

import { useUIContext } from '@/context/UIContext';
import { useSearchParams } from 'next/navigation';
import {  useEffect, useMemo } from 'react';
import { Virtuoso } from 'react-virtuoso';
import Row from './Row';
import { useSearchProductsQuery } from '@/http/api';

export default function Page() {
  const searchParams = useSearchParams();
  const uiContext = useUIContext()
  const hashtag = searchParams.get("hashtag") ?? "";
  const text = searchParams.get("text") ?? "";
  const { data , isLoading } = useSearchProductsQuery({
    mention:"",
    hashtag,
    text:[text]
  }, 
    { skip: !text && !hashtag});
  const rows = useMemo(() => {
    if (!data?.length) return [];
    return Array.from(
      { length: Math.ceil(data.length / 3) },
      (_, index) => {
        const start = index * 3;
        return {
          row: data.map(d => d.id).slice(start, start + 3),
          key: start,
        };
      }
    );
  }, [data]);
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