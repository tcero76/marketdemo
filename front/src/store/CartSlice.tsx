import { createSlice, PayloadAction } from "@reduxjs/toolkit";

interface CartItem {
  productId: number;
  title: string;
  price: number;
  quantity: number;
  thumbnail: string;
}

interface CartState {
  items: CartItem[];
}

const initialState: CartState = {
  items: [],
};

export const cartSlice = createSlice({
  name: "cart",
  initialState,
  reducers: {
    addItem(state, action: PayloadAction<CartItem>) {
      const item = state.items.find(
        (item) => item.productId === action.payload.productId
      );
      if (item) {
        item.quantity += action.payload.quantity;
      } else {
        state.items.push(action.payload);
      }
    },

    removeItem(state, action: PayloadAction<number>) {
      state.items = state.items.filter(
        (item) => item.productId !== action.payload
      );
    },

    dropItem(state, action: PayloadAction<{
        productId: number;
      }>) {
      const item = state.items.find(
        (item) => item.productId === action.payload.productId
      );

      if (item && item?.quantity == 1) {
        state.items = state.items.filter(
          (i) => i.productId !== item.productId
        );
      } else if (item ) {
        --item.quantity;
      }
      },

    updateQuantity(
      state,
      action: PayloadAction<{
        productId: number;
        quantity: number;
      }>
    ) {
      const item = state.items.find(
        (item) => item.productId === action.payload.productId
      );

      if (item && action.payload.quantity > 0) {
        item.quantity = action.payload.quantity;
      }
    },

    clearCart(state) {
      state.items = [];
    },
  },
});

export const {
  addItem,
  removeItem,
  dropItem,
  updateQuantity,
  clearCart,
} = cartSlice.actions;

export default cartSlice.reducer;