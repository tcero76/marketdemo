import { useRef } from 'react'
import { useRouter } from 'next/navigation'
import { useGlobalKeyboard } from '@/hooks/useKeyboardListener';
import { useGetCategoriesQuery } from '@/http/api';
import { useAutocomplete } from '@/hooks/useAutocomplete';
import { restoreCaretPosition, storeCaretPosition } from '@/utils/caret';
import { composeHighlighters } from '@/utils/highlights/composeHighlighters';
import { hashtagHighlighter } from '@/utils/highlights/highlighters';
import { matchBy } from '@/lib/meta/enrichMeta';

type SearchProps = {
} & React.HTMLAttributes<HTMLDivElement>

const Search = ({ className, ...props }: SearchProps) => {
    const { data:categories } = useGetCategoriesQuery()
    const hashtagAutocomplete = useAutocomplete({
        trigger: '#',
        categories: categories ?? [] });
    const router = useRouter();
    useGlobalKeyboard({
        FocusSearch: () => {
          inputRef.current?.focus();
        },
      }, { "Control+k": "FocusSearch" });
    const inputRef = useRef<HTMLDivElement>(null);
    const highlight =  (texto:string) => {
        const { metaRaw, cleanText } = composeHighlighters(
            hashtagHighlighter)(texto);
        return { metaRaw, cleanText };
    }
    const onChangeSearch = (e:React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
        e.preventDefault();
        const { metaRaw, cleanText } = highlight(inputRef.current?.innerText || '')
        const meta = matchBy<string>(metaRaw.hashtags, categories ?? [], c => c)
        const params = new URLSearchParams();
        meta.forEach(m => params.append("hashtag", m));
        if (cleanText.length > 0) params.append("text", cleanText);
        router.push(`/demo/search?${params.toString()}`);
    }
    const onKeyUp = (e: React.KeyboardEvent<HTMLDivElement>) => {
            hashtagAutocomplete.onKey(e);
        if(e.key === "Enter") {
            onChangeSearch(e as any)
        }
    }
    const onInput = async (ev: React.FormEvent<HTMLDivElement>) => {
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
            <div
                onInput={onInput}
                onKeyUp={onKeyUp}
                contentEditable
                className={`min-w-[150px] w-full px-5 py-1 bg-zinc-50 dark:bg-zinc-950 border border-zinc-200 dark:border-zinc-800
                            rounded-md text-[17px] leading-relaxed text-zinc-900 dark:text-zinc-100
                            placeholder-zinc-400 focus:outline-none focus:border-blue-500 dark:focus:border-blue-600
                            focus:ring-2 focus:ring-blue-500/20 transition-all duration-300
                            empty:before:content-[attr(data-placeholder)]
                            empty:before:text-zinc-400 dark:empty:before:text-zinc-500
                            empty:before:select-none ${className}`}
                data-placeholder="Ctrl+k"
                suppressContentEditableWarning
                ref={inputRef} {...props}/>
            <hashtagAutocomplete.AutocompleteList/>
            </>
    )
}
export default Search