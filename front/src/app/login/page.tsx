'use client'

import { } from 'react'
import LoginModal from '@/layout/LoginModal'

export default function Page() {
  return (
    <div className="flex relative w-full items-center justify-center p-6 md:p-10">
      <div className="w-full max-w-sm translate-y-[25%]">
        <LoginModal/>
      </div>
    </div>
  )
}