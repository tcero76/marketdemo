import { Suspense } from 'react'
import SearchPage from './SearchPage.tsx'

export default function Page() {
  return (
    <Suspense fallback={<div>Cargando...</div>}>
      <SearchPage />
    </Suspense>
  )
}