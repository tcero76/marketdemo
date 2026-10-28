
"use client";

import { useState } from "react";
import { usePathname, useRouter, useSearchParams } from "next/navigation";
import { SlidersHorizontal, X } from "lucide-react";

import { Button } from "@/components/ui/button";
import {
  Sheet,
  SheetContent,
  SheetDescription,
  SheetHeader,
  SheetTitle,
  SheetTrigger,
  SheetFooter,
  SheetClose,
} from "@/components/ui/sheet";
import { Checkbox } from "@/components/ui/checkbox";
import { Separator } from "@/components/ui/separator";

type Filters = {
  categories: string[];
  colors: string[];
  inStock: boolean;
};

const categories = [
  { label: "Beauty", value: "beauty" },
  { label: "Groceries", value: "groceries" },
  { label: "Fragrances", value: "fragrances" },
];

const colors = [
  { label: "Black", value: "black" },
  { label: "White", value: "white" },
  { label: "Blue", value: "blue" },
  { label: "Green", value: "green" },
];

function readFilters(params: URLSearchParams): Filters {
  return {
    categories: params.getAll("category"),
    colors: params.getAll("color"),
    inStock: params.get("in_stock") === "true",
  };
}

function countFilters(filters: Filters) {
  return (
    filters.categories.length +
    filters.colors.length +
    Number(filters.inStock)
  );
}

export function ProductFilters() {
  const router = useRouter();
  const pathname = usePathname();
  const searchParams = useSearchParams();

  const [open, setOpen] = useState(false);

  // Filtros aplicados en la URL.
  const [applied, setApplied] = useState<Filters>(() =>
    readFilters(new URLSearchParams(searchParams.toString()))
  );

  // Filtros que el usuario está editando en el panel.
  const [draft, setDraft] = useState<Filters>(() =>
    readFilters(new URLSearchParams(searchParams.toString()))
  );

  const activeCount = countFilters(applied);

  function openFilters() {
    setDraft({
      categories: [...applied.categories],
      colors: [...applied.colors],
      inStock: applied.inStock,
    });
    setOpen(true);
  }

  function toggleValue(
    key: "categories" | "colors",
    value: string,
    checked: boolean
  ) {
    setDraft((current) => ({
      ...current,
      [key]: checked
        ? [...current[key].filter((item) => item !== value), value]
        : current[key].filter((item) => item !== value),
    }));
  }

  function applyFilters() {
    const params = new URLSearchParams(searchParams.toString());

    // Reemplaza solamente los parámetros administrados por este componente.
    params.delete("hashtag");
    params.delete("color");
    params.delete("in_stock");

    draft.categories.forEach((value) =>
      params.append("hashtag", value)
    );

    draft.colors.forEach((value) =>
      params.append("color", value)
    );

    if (draft.inStock) {
      params.set("in_stock", "true");
    }

    const query = params.toString();
    const url = query ? `/search?${query}` : pathname;

    setApplied({
      categories: [...draft.categories],
      colors: [...draft.colors],
      inStock: draft.inStock,
    });

    setOpen(false);
    router.push(url, { scroll: false });
  }

  function clearFilters() {
    setDraft({
      categories: [],
      colors: [],
      inStock: false,
    });
  }

  return (
    <Sheet
      open={open}
      onOpenChange={(nextOpen) => {
        if (nextOpen) {
          openFilters();
        } else {
          setOpen(false);
        }
      }}
    >
      <SheetTrigger asChild>
        <Button
          variant="ghost"
          className="gap-2 px-2"
          onClick={openFilters}
        >
          <SlidersHorizontal className="size-4" />
          <span>Filters</span>

          {activeCount > 0 && (
            <span className="flex size-5 items-center justify-center rounded-full bg-foreground text-xs text-background">
              {activeCount}
            </span>
          )}
        </Button>
      </SheetTrigger>

      <SheetContent
        side="left"
        className="flex w-full flex-col sm:max-w-md"
      >
        <SheetHeader>
          <div className="flex items-center justify-between">
            <SheetTitle>Filters</SheetTitle>

            {countFilters(draft) > 0 && (
              <Button
                variant="ghost"
                size="sm"
                onClick={clearFilters}
              >
                Clear all
              </Button>
            )}
          </div>

          <SheetDescription>
            Refine your selection to find the right products.
          </SheetDescription>
        </SheetHeader>

        <div className="flex-1 space-y-6 overflow-y-auto px-4 py-4">
          {/* Categorías */}
          <section className="space-y-3">
            <h3 className="font-medium">Categoria</h3>

            {categories.map((hashtag) => (
              <label
                key={hashtag.value}
                className="flex cursor-pointer items-center gap-3 text-sm"
              >
                <Checkbox
                  checked={draft.categories.includes(hashtag.value)}
                  onCheckedChange={(checked) =>
                    toggleValue(
                      "categories",
                      hashtag.value,
                      checked === true
                    )
                  }
                />
                {hashtag.label}
              </label>
            ))}
          </section>
          <Separator />
        </div>
        <SheetFooter className="flex-row gap-2 border-t">
          <SheetClose asChild>
            <Button variant="outline" className="flex-1">
              Cancel
            </Button>
          </SheetClose>

          <Button
            className="flex-1"
            onClick={applyFilters}
          >
            Apply filters ({countFilters(draft)})
          </Button>
        </SheetFooter>
      </SheetContent>
    </Sheet>
  );
}
