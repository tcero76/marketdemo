import { ShoppingCartIcon, Trash2Icon } from "lucide-react";
import { Button } from "@/components/ui/button";
import {
  Drawer,
  DrawerClose,
  DrawerContent,
  DrawerDescription,
  DrawerFooter,
  DrawerHeader,
  DrawerTitle,
  DrawerTrigger,
} from "@/components/ui/drawer";
import { useAppSelector, useAppDispatch } from "@/store/hooks";
import { currencyFormatter } from "@/utils/currency";
import type { MouseEvent } from "react";
import { dropItem } from "@/store/CartSlice";

interface CartItem {
  id: string;
  name: string;
  price: string;
  variant: string;
}

const CART_ITEMS: CartItem[] = [
  {
    id: "headphones",
    name: "Wireless Headphones",
    price: "$99.00",
    variant: "Black, Standard",
  },
  { id: "cable", name: "USB-C Cable", price: "$19.00", variant: "White, 2m" },
  { id: "case", name: "Phone Case", price: "$29.00", variant: "Clear, Slim" },
];
const RightDrawerShoppingCart = () => {
  const items = useAppSelector((state) => state.cart.items);
  const totalItems = items.reduce((total, item) => total + item.quantity,0)
  const total = items.reduce((total, item) => total + item.price * item.quantity, 0)
  const dispatch = useAppDispatch()
  const onClickRemove = (productId:number) => {
    dispatch(dropItem({ productId }))
  }
  return (
  <Drawer swipeDirection="right">
    <DrawerTrigger render={<Button variant="ghost" />}>
      <ShoppingCartIcon aria-hidden="true" data-icon="inline-start" />
      {`Cart (${totalItems})`}
    </DrawerTrigger>
    <DrawerContent>
      <DrawerHeader>
        <DrawerTitle>Shopping Cart</DrawerTitle>
        <DrawerDescription>3 items in your cart</DrawerDescription>
      </DrawerHeader>
      <div className="flex min-h-0 flex-1 flex-col gap-4 overflow-y-auto p-4">
        <ul className="flex flex-col gap-4">
          {items.map((item) => (
            <li className="flex gap-4 border-b pb-4" key={item.productId}>
              <div className="size-20 shrink-0 rounded-md bg-muted">
                <img src={item.thumbnail}/>
                </div>
              <div className="flex flex-1 flex-col gap-1">
                <h3 className="font-medium text-sm">{item.title}</h3>
                <p className="text-muted-foreground text-xs">x{item.quantity}</p>
                <div className="flex items-center justify-between">
                  <span className="font-medium text-sm">{item.price}</span>
                  <Button
                    size="sm"
                    variant="ghost"
                    onClick={() => onClickRemove(item.productId)}
                  >
                    <Trash2Icon aria-hidden="true" />
                  </Button>
                </div>
              </div>
            </li>
          ))}
        </ul>
        <dl className="flex flex-col gap-2 text-sm">
          <div className="flex justify-between">
            <dt className="text-muted-foreground">Subtotal</dt>
            <dd>{currencyFormatter.format(total)}</dd>
          </div>
          <div className="flex justify-between">
            <dt className="text-muted-foreground">Shipping</dt>
            <dd>{currencyFormatter.format(total*.2)}</dd>
          </div>
          <div className="flex justify-between border-t pt-2 font-medium">
            <dt>Total</dt>
            <dd>{currencyFormatter.format(total*1.2)}</dd>
          </div>
        </dl>
      </div>
      <DrawerFooter>
        <Button>Checkout</Button>
        <DrawerClose render={<Button variant="outline" />}>
          Continue Shopping
        </DrawerClose>
      </DrawerFooter>
    </DrawerContent>
  </Drawer>
)
};

export default RightDrawerShoppingCart;
