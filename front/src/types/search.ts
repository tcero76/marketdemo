export type SearchRequest = {
    mention: string;
    hashtag: string;
    text: string[];
}

export type SearchProps = {
    onSearch: (search:SearchRequest) => void;
}