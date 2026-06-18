'use client'

import { useEffect } from 'react'
import CreatePost from './post/CreatePost'
import { useParams } from 'next/navigation'
import { Modelo } from '@/types'
import { useUIContext } from '@/context/UIContext'
import { useGetProductQuery } from '@/http/api'

const Page = () => {
  const uiContext = useUIContext()
  const params = useParams() 
  const id = params.id as string
  const { data, isLoading } = useGetProductQuery(parseInt(id))
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
    <div className="container mx-auto px-4 py-6 max-w-7xl">
      <div className="mb-6">
        <h1 className="text-3xl font-bold tracking-tight sm:text-4xl">
          {data?.title || 'Modelo'}
        </h1>
        <p className="mt-3 text-lg text-gray-600 dark:text-gray-400">
          {data?.description || 'Descripción del modelo...'}
        </p>
      </div>
        <div className="bg-gray-800/30 rounded-xl p-6 border border-gray-700 shadow-sm">
          <CreatePost id={id ?? 0}/>
        </div>
    </div>
      )
}
export default Page