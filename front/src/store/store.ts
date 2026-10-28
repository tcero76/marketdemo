import { configureStore } from '@reduxjs/toolkit';

import { authSlice } from './AuthSlice'
import { cartSlice } from './CartSlice'
import { api, hydra, demo } from '@/http/api';

export const store = configureStore({
  reducer: {
    auth: authSlice.reducer,
    cart: cartSlice.reducer,
    [api.reducerPath]: api.reducer,
    [hydra.reducerPath]:  hydra.reducer,
    [demo.reducerPath]: demo.reducer
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware().concat(api.middleware).concat(hydra.middleware).concat(demo.middleware),
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
