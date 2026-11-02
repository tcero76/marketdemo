import { sortOptions, type Product, type SortBy } from "@/types/demo";

export function getProductRows(data: Product[] | undefined, sortBy: SortBy) {
  if (!data?.length) return [];
  var newData = [...data];


const sortOption = sortOptions.find(option => option.value === sortBy);

const sortData = sortOption?.compare
  ? [...newData].sort(sortOption.compare)
  : [...newData];



//   if (sortBy === "price-ascending") {
//     sortData.sort((d1, d2) => d1.price - d2.price);
//   } else if (sortBy === "price-descending") {
//     sortData.sort((d1, d2) => d2.price - d1.price);
//   }


  return Array.from(
    { length: Math.ceil(sortData.length / 3) },
    (_, index) => {
      const start = index * 3;

      return {
        row: sortData.slice(start, start + 3).map(d => d.id),
        key: start,
      };
    }
  );
}