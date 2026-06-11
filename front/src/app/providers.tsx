'use client'

import { store } from '../store/store'
import { Provider } from 'react-redux';
import { ReactQueryDevtools } from "@tanstack/react-query-devtools/production";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { UIProvider } from '@/context/UIContext';

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      refetchOnWindowFocus: false,
      retry: 1,
    },
  },
});
export function Providers({ children }: { children: React.ReactNode }) {
  return (
      <Provider store={store}>
        <UIProvider>
          <QueryClientProvider client={queryClient}>
              {process.env.NEXT_PUBLIC_MOCK === "true" && <ReactQueryDevtools initialIsOpen={false} />}
                          {children}
          </QueryClientProvider>
        </UIProvider>
      </Provider>
  ) 
}