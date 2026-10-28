import {
    useDispatch,
    useSelector,
    type TypedUseSelectorHook,
  } from 'react-redux';
  
import type { AppDispatch, RootState } from './store';
  
  export type DispatchFunction = () => AppDispatch;
  
export const useAppDispatch = () => useDispatch<AppDispatch>();
export const useAppSelector: TypedUseSelectorHook<RootState> = useSelector;
