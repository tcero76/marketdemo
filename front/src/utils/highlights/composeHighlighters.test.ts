import { describe, expect, it } from 'vitest';
import { composeHighlighters } from './composeHighlighters';

describe('composeHighlighters', () => { 
    it('should compose multiple highlighters', () => {
        const first = (text: string) => ({
        html: text.replace('hola', '<b>hola</b>'),
        metaRaw: {},
        cleanText: '',
        });

        const second = (text: string) => ({
        html: text.replace('mundo', '<i>mundo</i>'),
        metaRaw: {},
        cleanText: '',
        });
        const highlighter = composeHighlighters(first, second);
        const result = highlighter('hola mundo');
        expect(result.html).toBe('<b>hola</b> <i>mundo</i>');
        expect(result.cleanText).toBe('hola mundo');
    });

    it('should merge metadata from all highlighters', () => {
        const hashtags = (text: string) => ({
            html: text, 
            metaRaw: {
            hashtags: ['#gato'],
            },
            cleanText: '',
        });

        const mentions = (text: string) => ({
            html: text,
            metaRaw: {
            mentions: ['@leo'],
            },
            cleanText: '',
        });

        const highlighter = composeHighlighters(
            hashtags,
            mentions
        );

        const result = highlighter('Hola #gato @leo');

        expect(result.metaRaw).toEqual({
            hashtags: ['#gato'],
            mentions: ['@leo'],
        });
    });

    it('should remove highlighted values from cleanText', () => {
        const hashtags = (text: string) => ({
            html: text,
            metaRaw: {
            hashtags: ['#gato'],
            },
            cleanText: '',
        });

        const highlighter = composeHighlighters(hashtags);

        const result = highlighter('Hola #gato y #gato');

        expect(result.cleanText).toBe('Hola y');
    });

    it('should remove different highlighted values from cleanText', () => {
        const highlighter = composeHighlighters(
            (text: string) => ({
            html: text,
            metaRaw: {
                hashtags: ['#gato', '#perro'],
            },
            cleanText: '',
            })
        );

        const result = highlighter('Hola #gato y #perro');

        expect(result.cleanText).toBe('Hola y');
    });
    
    it('should normalize whitespace in cleanText', () => {
        const highlighter = composeHighlighters(
            (text: string) => ({
            html: text,
            metaRaw: {
                hashtags: ['#gato'],
            },
            cleanText: '',
            })
        );

        const result = highlighter('  Hola   #gato    mundo  ');

        expect(result.cleanText).toBe('Hola mundo');
    });

    it('should preserve text when there are no highlighted values', () => {
        const highlighter = composeHighlighters(
            (text: string) => ({
            html: text,
            metaRaw: {},
            cleanText: '',
            })
        );

        const result = highlighter('Hola mundo');

        expect(result.cleanText).toBe('Hola mundo');
    });

    it('should ignore objects without nombre', () => {
        const highlighter = composeHighlighters(
            (text: string) => ({
                html: text,
                metaRaw: {
                    products: [{ id: 123 }],
                },
                cleanText: '',
                })
            );

        const result = highlighter('Hola mundo');

        expect(result.cleanText).toBe('Hola mundo');
    });
});