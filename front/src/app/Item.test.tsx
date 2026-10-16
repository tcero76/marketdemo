import { describe, expect, it, vi } from 'vitest';
import { render, screen } from '@testing-library/react';

import { useGetProductQuery } from '@/http/api';
import Item from './Item';

vi.mock('@/http/api', () => ({
  useGetProductQuery: vi.fn(),
}));

describe('Item', () => {
    it('should display the product returned by the API', () => {
        vi.mocked(useGetProductQuery).mockReturnValue({
        data: {
            id: 1,
            title: 'iPhone 15',
            category: 'Smartphones',
            description: 'Un teléfono',
            stock: 10,
        },
        } as any);
        render(<Item item={1} />);
        expect(screen.getByText('iPhone 15')).toBeInTheDocument();
        expect(screen.getByText('Smartphones')).toBeInTheDocument();
        expect(screen.getByText('Un teléfono')).toBeInTheDocument();
        expect(screen.getByText('10')).toBeInTheDocument();
    });

    it('should link to the product detail page', () => {
        vi.mocked(useGetProductQuery).mockReturnValue({
            data: {
            id: 42,
            title: 'iPhone 15',
            category: 'Smartphones',
            description: 'Un teléfono',
            stock: 10,
            },
        } as any);
        render(<Item item={42} />);
        expect(
        document.querySelector('a[href="/demo/product/42"]')
        ).toBeTruthy();
    });

    it('should skip the query when there is no product id', () => {
        render(<Item item={0} />);
        expect(useGetProductQuery).toHaveBeenCalledWith(
            0,
            { skip: true }
        );
    });
});
