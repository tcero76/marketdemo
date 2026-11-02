
"use client";

import { useRef, useState } from "react";
import { usePathname, useRouter, useSearchParams } from "next/navigation";
import { Search } from "lucide-react";

import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
} from "@/components/ui/dialog";
import { Command, CommandEmpty, CommandGroup, CommandItem, CommandList } from "@/components/ui/command";
import { useGetCategoriesQuery, useSearchProductsQuery } from "@/http/api";
import { useAutocomplete } from "@/hooks/useAutocomplete";
import { matchBy } from "@/lib/meta/enrichMeta";
import { composeHighlighters } from "@/utils/highlights/composeHighlighters";
import { hashtagHighlighter } from "@/utils/highlights/highlighters";
import { restoreCaretPosition, storeCaretPosition } from "@/utils/caret";
import { cn } from "cn";

export function  CatalogActions() {
  const [searchOpen, setSearchOpen] = useState(false);
  const [query, setQuery] = useState("");
  const router = useRouter();
  const pathname = usePathname();
  const searchParams = useSearchParams();
  const { data, isFetching } = useSearchProductsQuery({
      mention: "",
      hashtag: "",
      text: [query],
    },
    { 
      skip: query.length < 2
    }
  );



    const inputRef = useRef<HTMLDivElement | null>(null);
    const { data:categories } = useGetCategoriesQuery()
    const hashtagAutocomplete = useAutocomplete({
      trigger: '#',
      categories: categories ?? [] });
    const highlight =  (texto:string) => {
        const { metaRaw, cleanText } =
          composeHighlighters(hashtagHighlighter)(texto);
        return { metaRaw, cleanText };
    }
    const onChangeSearch = (e:React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
        const { metaRaw, cleanText } = highlight(inputRef.current?.innerText || '')
        const meta = matchBy<string>(metaRaw.hashtags, categories ?? [], c => c)
        const params = new URLSearchParams();
        meta.forEach(m => params.append("hashtag", m));
        if (cleanText.length > 0) params.append("text", cleanText);
        router.push(`/search?${params.toString()}`);
        setSearchOpen(false)
    }
    const onKeyUp = (e: React.KeyboardEvent<HTMLDivElement>) => {
      e.preventDefault()
      hashtagAutocomplete.onKey(e);
      if(e.key === "Enter") {
          onChangeSearch(e as any)
      }
    }

  function submitSearch(event: React.FormEvent<HTMLFormElement>) {
    const params = new URLSearchParams(searchParams.toString());
    const term = query.trim();
    if (term) {
      params.set("q", term);
    } else {
      params.delete("q");
    }

    const queryString = params.toString();
    router.push(
      queryString ? `${pathname}?${queryString}` : pathname,
      { scroll: false }
    );

    setSearchOpen(false);
  }
  const onClickSelect = (productId:number) => {
    router.push(`/product/${productId}`);
    setSearchOpen(false);
  }
  const onInput = async (ev: React.FormEvent<HTMLDivElement>) => {
      setQuery(inputRef?.current?.innerText!)
      const editor = inputRef.current;
      if(!editor) return
      const native = ev.nativeEvent as InputEvent;
      const inputType = native.inputType;
      if(inputType === 'insertCompositionText' || inputType === 'deleteCompositionText') return
      const pos = storeCaretPosition(editor)
      editor.innerHTML = editor.innerText
      restoreCaretPosition(editor, pos);
  };
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
          <Command shouldFilter={false} className="bg-transparent text-foreground">
            <div className="flex items-center border-b px-3 cmdk-input-wrapper" >
              <Search className="mr-2 h-4 w-4 shrink-0 opacity-50" />
            <div
                onInput={onInput}
                onKeyUp={onKeyUp}
                contentEditable
                className={cn("flex h-10 w-full rounded-md bg-transparent py-3 outline-none placeholder:text-muted-foreground disabled:opacity-50")}
                data-placeholder="Search products..."
                suppressContentEditableWarning
                ref={inputRef}>
                  </div>
            <hashtagAutocomplete.AutocompleteList/>
            </div>

            <hashtagAutocomplete.AutocompleteList/>
            <CommandList>
              <CommandEmpty>
                No products found.
              </CommandEmpty>
              <CommandGroup heading="Suggestions">
                {data?.slice(0, 3).map((product) => (
                  <CommandItem
                    key={product.id}
                    value={String(product.id)}
                    onSelect={(e) => onClickSelect(product.id)}
                  >
                    <img
                      src={product.thumbnail}
                      alt={product.title}
                      className="size-10 rounded object-cover"
                    />
                    <span>{product.title}</span>
                  </CommandItem>
                ))}
              </CommandGroup>
            </CommandList>
          </Command>
        </DialogContent>
      </Dialog>
    </>
  );
}
