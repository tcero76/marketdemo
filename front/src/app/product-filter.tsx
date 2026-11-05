
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
import { RadioGroup, RadioGroupItem } from "@/components/ui/radio-group";
import { Label } from "@/components/ui/label";

type Filters = {
  categories: string[];
};

const categories = [
  { label: "Beauty", value: "beauty" },
  { label: "Groceries", value: "groceries" },
  { label: "Fragrances", value: "fragrances" },
  { label: "Furnitures", value: "furniture"}
];

function readFilters(params: URLSearchParams): Filters {
  return {
    categories: params.getAll("hashtag"),
  };
}

function countFilters(filters: Filters) {
  return filters.categories.length;
}

export function ProductFilters() {
  const router = useRouter();
  const pathname = usePathname();
  const searchParams = useSearchParams();

  const [open, setOpen] = useState(false);

  const [draft, setDraft] = useState<Filters>(() =>
    readFilters(new URLSearchParams(searchParams.toString()))
  );

  const applied = readFilters(
    new URLSearchParams(searchParams.toString())
  );

  const activeCount = countFilters(applied);

  function openFilters() {
    setDraft({
      categories: [...applied.categories],
    });
    setOpen(true);
  }

  function selectCategory(value: string) {
    setDraft((current) => ({
      ...current,
      categories: value ? [value] : [],
    }));
  }

  function applyFilters() {
    const params = new URLSearchParams(searchParams.toString());
    params.delete("hashtag");
    draft.categories.forEach((value) =>
      params.set("hashtag", value)
    );
    const query = params.toString();
    router.push(
      query ? `/search?${query}` : pathname,
      { scroll: false }
    );
    setOpen(false);
  }

  function clearFilters() {
    setDraft({
      categories: [],
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

          <RadioGroup defaultValue="comfortable" className="w-fit" value={draft.categories[0] ?? ""} onValueChange={selectCategory}>
            {categories.map((hashtag) => (
              <div className="flex items-center gap-3" key={hashtag.value}>
                <RadioGroupItem value={hashtag.value} id="r1" />
                <Label htmlFor="r1">{hashtag.label}</Label>
              </div>
            ))}
            </RadioGroup>
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