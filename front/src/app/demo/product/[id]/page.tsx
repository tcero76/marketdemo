'use client'

import { useEffect, useState } from 'react'
import CreatePost from './post/CreatePost'
import { useParams } from 'next/navigation'
import { Modelo } from '@/types'
import { publishWebRTCOffer } from '../webRTC'
import { useUIContext } from '@/context/UIContext'
import { useGetProductQuery } from '@/http/api'

const initModelo:Modelo = {
  descripcion: '',
  modelo: '', id: 0,
  fecharegistro: new Date(),
  idJob: 0,
  idModelos: 0
}

const Page = () => {
  const uiContext = useUIContext()
  const params = useParams() 
  const id = parseInt(params.id as string)
  const [accessToken, setAcessToken] = useState<string>('');
  const src = `${process.env.NEXT_PUBLIC_HOST}/hls/streams/index.m3u8`
  const { data, isLoading } = useGetProductQuery(id)
  // const pcRef = useRef<RTCPeerConnection>(new RTCPeerConnection());
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
  useEffect(() => {
    let isMounted = true;
    async function initWebRTC() {
      try {
        pcRef.current = new RTCPeerConnection();
        const res = await fetch(`${process.env.NEXT_PUBLIC_HOST}/bff/token`, { method: "POST" });
        const jwt = (await res.json()).access_token;
        setAcessToken(jwt);
        if (!isMounted) return;
        await publishWebRTCOffer(jwt, pcRef.current);
      } catch (err) {
        console.error("WebRTC error:", err);
      }
    }
  },[])
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