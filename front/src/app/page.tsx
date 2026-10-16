'use client'

import { useEffect, useMemo } from 'react';
import { useAuthDispatch } from '@/store/hooks';
import { Virtuoso } from 'react-virtuoso';
import {
  useGetAuthenticatedQuery,
  useGetProductsIdxQuery,
 } from '@/http/api';
import { setToken } from '@/store/AuthSlice';
import { skipToken } from '@reduxjs/toolkit/query';
import Row from './search/Row';

export default function Home() {
  const dispatch = useAuthDispatch()
  const { data:user } = useGetAuthenticatedQuery()
  const { data, isLoading:isLoadingRecomendations } = useGetProductsIdxQuery(user?.sub ? undefined : skipToken);
  useEffect(() => {
    const url = new URL(window.location.href)
    const accessToken = url.searchParams.get('accessToken');
    if (accessToken) {
      dispatch(setToken(accessToken))
      url.searchParams.delete('accessToken')
      window.history.replaceState({}, '', url.toString())
    }
  }, [dispatch]);
  const rows = useMemo(() => {
  if (!data?.length) return [];
  return Array.from(
      { length: Math.ceil(data.length / 3) },
      (_, index) => {
        const start = index * 3;
        return {
          row: data.slice(start, start + 3),
          key: start,
        };
      }
    );
  }, [data]);
  if (!user) return <div>No estás logeado</div>;
  if (isLoadingRecomendations) return <div>Cargando recomendaciones...</div>;
  return (
    <Virtuoso
      scrolling='none'
      data={rows}
      itemContent={(index, item) => <Row key={item.key} row={item.row}/>}
      computeItemKey={(index) => index}
    >
    </Virtuoso>
  );
}
