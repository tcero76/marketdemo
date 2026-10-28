/**
 * A right drawer shopping cart with three items, remove buttons, an order
 * summary and Checkout and Continue Shopping buttons.
 */
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

const RightDrawerShoppingCart = () => (
  <Drawer swipeDirection="right">
    <DrawerContent>
      <DrawerHeader>
        <DrawerTitle>Shopping Cart</DrawerTitle>
        <DrawerDescription>3 items in your cart</DrawerDescription>
      </DrawerHeader>
      <div className="flex min-h-0 flex-1 flex-col gap-4 overflow-y-auto p-4">
        <ul className="flex flex-col gap-4">
          {CART_ITEMS.map((item) => (
            <li className="flex gap-4 border-b pb-4" key={item.id}>
              <div className="size-20 shrink-0 rounded-md bg-muted" />
              <div className="flex flex-1 flex-col gap-1">
                <h3 className="font-medium text-sm">{item.name}</h3>
                <p className="text-muted-foreground text-xs">{item.variant}</p>
                <div className="flex items-center justify-between">
                  <span className="font-medium text-sm">{item.price}</span>
                  <Button
                    aria-label={`Remove ${item.name}`}
                    size="lg"
                    variant="ghost"
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
            <dd>$147.00</dd>
          </div>
          <div className="flex justify-between">
            <dt className="text-muted-foreground">Shipping</dt>
            <dd>$10.00</dd>
          </div>
          <div className="flex justify-between border-t pt-2 font-medium">
            <dt>Total</dt>
            <dd>$157.00</dd>
          </div>
        </dl>
      </div>
      <DrawerFooter>
        <Button>Checkout</Button>
        <DrawerClose
        render={<Button variant="outline">Continue Shopping</Button>}
        />
      </DrawerFooter>
    </DrawerContent>
  </Drawer>
);