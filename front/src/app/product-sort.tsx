
"use client";

import { useRouter, usePathname, useSearchParams } from "next/navigation";
import { ChevronDown } from "lucide-react";

import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { sortOptions } from "@/types/demo";

export function ProductSort() {
  const router = useRouter();
  const pathname = usePathname();
  const searchParams = useSearchParams();

  const currentSort = searchParams.get("sort_by") ?? "featured";

  function handleSortChange(value: string) {
    const params = new URLSearchParams(searchParams.toString());

    params.set("sort_by", value);

    router.push(`/search?${params.toString()}`, {
      scroll: false,
    });
  }

  return (
    <Select value={currentSort} onValueChange={handleSortChange}>
      <SelectTrigger
        aria-label="Sort products"
        className="w-auto gap-2 border-0 bg-transparent px-0 shadow-none focus:ring-0"
      >
        <SelectValue placeholder="Sort" />
      </SelectTrigger>

      <SelectContent align="end">
        {sortOptions.map((option) => (
          <SelectItem key={option.value} value={option.value}>
            {option.label}
          </SelectItem>
        ))}
      </SelectContent>
    </Select>
  );
}
