import { describe, expect, it } from 'vitest';
import { arrobaHighlighter, hashtagHighlighter, httpsHighlighter } from './highlighters';

describe('Highlighter', () => {
    it('should extract hashtags', () => {
        const result = hashtagHighlighter('Hola #mundo');

        expect(result.metaRaw.hashtags).toEqual(['#mundo']);
    });

    it('should extract arrobas', () => {
        const result = arrobaHighlighter('Hola @mundo');

        expect(result.metaRaw.mentions).toEqual(['@mundo']);
    });

    it('should extract https links', () => {
        const result = httpsHighlighter('Hola https://mundo.com');

        expect(result.html).toEqual('Hola <span class="font-bold text-blue-500" contenteditable="true">https://mundo.com</span>');
    });
});