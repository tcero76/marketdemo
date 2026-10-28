
"use client";

import { useState } from "react";
import { usePathname, useRouter, useSearchParams } from "next/navigation";
import { Search, ShoppingBag } from "lucide-react";

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogDescription,
} from "@/components/ui/dialog";
import {
  Sheet,
  SheetContent,
  SheetHeader,
  SheetTitle,
  SheetDescription,
} from "@/components/ui/sheet";

export function CatalogActions() {
  const [searchOpen, setSearchOpen] = useState(false);
  const [cartOpen, setCartOpen] = useState(false);
  const [query, setQuery] = useState("");

  const router = useRouter();
  const pathname = usePathname();
  const searchParams = useSearchParams();

  function submitSearch(event: React.FormEvent<HTMLFormElement>) {
    console.log("Entra a submit...")
    event.preventDefault();

    const params = new URLSearchParams(searchParams.toString());
    const term = query.trim();

    if (term) {
      params.set("q", term);
    } else {
      params.delete("q");
    }

    const queryString = params.toString();
    console.log("queryString: ", queryString)
    console.log("pathname: ", pathname)
    router.push(
      queryString ? `${pathname}?${queryString}` : pathname,
      { scroll: false }
    );

    setSearchOpen(false);
  }

  return (
    <>
      <Button
        variant="ghost"
        size="icon"
        aria-label="Search products"
        onClick={() => setSearchOpen(true)}
        className="size-9"
      >
        <Search className="size-5" />
      </Button>
      <Dialog open={searchOpen} onOpenChange={setSearchOpen}>
        <DialogContent className="sm:max-w-lg">
          <DialogHeader>
            <DialogTitle>Search products</DialogTitle>
            <DialogDescription>
              Find products in our catalog.
            </DialogDescription>
          </DialogHeader>

          <form onSubmit={submitSearch} className="flex gap-2">
            <Input
              autoFocus
              placeholder="Search products..."
              value={query}
              onChange={(event) => setQuery(event.target.value)}
            />
            <Button type="submit" aria-label="Submit search">
              <Search className="size-4" />
            </Button>
          </form>
        </DialogContent>
      </Dialog>
    </>
  );
}
