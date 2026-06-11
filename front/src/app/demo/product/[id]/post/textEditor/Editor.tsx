'use client'

import { FC, useRef } from "react";
import { PosteoRaw } from "@/types";
import { usePasteImage } from "@/hooks/usePasteImage";
import useFormatText from "@/hooks/useFormatText";
import { useAutocomplete } from '@/hooks/useAutocomplete';
import { useGetCategoriesQuery } from '@/http/api';
import { matchBy } from "@/lib/meta/enrichMeta";
import { PosteoCreate, type EditorProps } from "@/types/demo";

const Editor:FC<EditorProps> = ({ onChangePosteo, posteo, ...props }) => {
  const { data:categories } = useGetCategoriesQuery()
  const handleChangePosteo = (posteoRaw: PosteoRaw) => {
    const meta = matchBy(
      posteoRaw.metaRaw.hashtags,
      categories ?? [],
      c => c
    );
    onChangePosteo({
      id: posteoRaw.id,
      meta: { 
        hashtags: meta.map(m => ({ nombre: m }))
      },
      texto: posteoRaw.texto,
      productId: posteoRaw.userId
    });
  };
  const editorRef = useRef<HTMLDivElement>(null);
  const { imageUrl, isLoading } = usePasteImage(editorRef);
  const hashtagAutocomplete = useAutocomplete({ trigger: '#',
    categories: categories?.map(c => c) ?? [] });
  const { onInput, Embeded } = useFormatText<PosteoCreate>({
    onChangePosteo:handleChangePosteo,
    editorRef,
    imageUrl,
    posteo
  })
  const onKeyUp = (e: React.KeyboardEvent<HTMLDivElement>) => {
    hashtagAutocomplete.onKey(e);
  }
  if(!posteo) return <div>Cargando...</div>
  return (
    <div>
      <div
        spellCheck="false"
        ref={editorRef}
        contentEditable
        onKeyUp={onKeyUp}
        onInput={onInput}
        className="p-[10px] min-h-[100px] border border-gray-300 rounded"
        suppressContentEditableWarning
        {...props}
      />
      {isLoading ? <div>loading</div> : Embeded }
      <hashtagAutocomplete.AutocompleteList/>
    </div>);
};

export default Editor;