import { AuthorizationType, AuthType, CredencialType, EmbededType,
  FetchLoginChallengeType, ImagePasteRes, LoginPayloadType, LoginResponseType,
  SearchPosts, SearchType } from "@/types";
import { createApi } from "@reduxjs/toolkit/query/react";
import { baseQueryWithRefresh } from "./baseQueryWithRefresh"
import { fetchLoginChallenge, logout } from "@/store/AuthSlice";
import { baseHydraQuery } from "./baseQuery";
import { Product, PosteoCreate } from "@/types/demo";

export const api = createApi({
  reducerPath: "api",
  baseQuery: baseQueryWithRefresh,
  tagTypes: ['Posteos'],
  endpoints: (builder) => {
    return (
      {
      getAuthenticated: builder.query<AuthType, void>({
        query: () => ({
          url: `/getAuthentication`,
          method: 'GET',
          headers: {
                  'Content-Type': 'application/json',
                }
        })
      }),
      login: builder.mutation<AuthorizationType, LoginPayloadType>({
        query: ({ user, password }:LoginPayloadType) => ({
          url: `/login?` +
            `login_challenge=${sessionStorage.getItem("login_challenge")}` +
            `&state=${sessionStorage.getItem("state")}` +
            `&idp=internal`,
          method: 'POST',
          body: { user, password }, 
          headers:  {
              'Content-Type': 'application/json'
            }
        }),
        async onQueryStarted(_, { dispatch, queryFulfilled }) {
          try {
            const { data, meta } = await queryFulfilled
            if (data.accessToken) {
              sessionStorage.setItem("Access_Token", data.accessToken)
              sessionStorage.removeItem("login_challenge")
              sessionStorage.removeItem("state")
            }
          } catch (err) {
            console.error("Login error:", err)
          }
        }
      }),
      logout: builder.mutation<string, void>({
        query: () => ({
          url: `${process.env.NEXT_PUBLIC_HOST}/bff/logout`,
          method: 'GET',
          headers: {
              'Content-Type': 'application/json',
            }
        }),
        async onQueryStarted(_, { dispatch, queryFulfilled }) {
              await queryFulfilled  
              dispatch(logout())
        }
      }),
      loginGoogle: builder.query<LoginResponseType, void>({
        query: () => ({
          url: `${process.env.NEXT_PUBLIC_HOST}/bff/login?` +
            `login_challenge=${sessionStorage.getItem("login_challenge")}` +
            `&state=${sessionStorage.getItem("state")}` +
            `&idp=google`,
          method: 'POST',
          headers: {
              'Content-Type': 'application/json'
            }
        }),
        async onQueryStarted(_, { dispatch, queryFulfilled }) {
          const { data } = await queryFulfilled
          if (data.url !== undefined) {
              window.location.href = data.url;
          }   
        }
      }),
      signUp: builder.query<string, CredencialType>({
        query: ({ user, password }: CredencialType) => ({
          url: `${process.env.NEXT_PUBLIC_HOST}/bff/signup`,
          method: 'POST',
          body: { email:user, password },
          headers: {
              'Content-Type': 'application/json'
            }
        })
      }),
      onImagePaste: builder.mutation<ImagePasteRes, File>({
        query: (file: File) => {
          const formData = new FormData()
          formData.append("image", file)
          return {
            url: "/uploadImage",
            method: "POST",
            body: formData
          }
        }
      }),
      onEmbed: builder.query<EmbededType, string>({
        query: (url: string) => ({
          url: `${process.env.NEXT_PUBLIC_HOST}/bff/embeded`,
          method: "GET",
          params: { url }
        })
      }),
      }
    );
  }
});

export const demo = createApi({
  reducerPath: "baseHydraQuery",
  baseQuery: baseQueryWithRefresh,
  tagTypes: ['PosteosDemo'],
  endpoints:(builder) => {
    return (
      {
        getProductsIdx: builder.query<number[], void>({
          query: () => ({
            url: `/usuario/product/getRecommendations`
          })
        }),
        getProduct: builder.query<Product, number>({
          query: (id:number) => ({ 
            url:`/usuario/getProduct?product=${id}`
          })
        }),
        getProducts: builder.query<Product[], void>({
          query: () => ({
            url: `/usuario/getProducts`
          })
        }),
        getCategories: builder.query<string[], void>({
          query: () => ({
            url: `/usuario/getCategories`
          })
        }),
        searchProducts: builder.query<SearchPosts[],SearchType>({
          query: (search:SearchType) => ({ 
            url:`/usuario/searchProducts`,
            method: 'POST',
            body: search
          })
        }),
        sendPostDemo: builder.mutation<string, PosteoCreate>({
          query: (posteo:PosteoCreate) => ({
            url: `/usuario/createPostDemo`,
            method: 'POST',
            body: posteo
          }),
          invalidatesTags: ['PosteosDemo']
        }),
        getPosteosDemo: builder.query<PosteoCreate[], string>({
          query: (id:string) => ({
            url: `/usuario/getPosteosDemo?productId=${id}`,
            method: 'GET',
          }),
          providesTags: ['PosteosDemo'],
        }),
      }
    )
   }
})

export const hydra = createApi({
  reducerPath: "hydra",
  baseQuery: baseHydraQuery,
  endpoints:(builder) => {
    return (
      {
      fetchLoginChallenge: builder.mutation<FetchLoginChallengeType, string>({
        query: (state) => ({
          url: `/hydra/oauth2/auth?client_id=${process.env.NEXT_PUBLIC_CLIENT_ID}` +
          `&response_type=code&scope=openid offline mediamtx:stream` +
          `&state=${state}` +
           `&redirect_uri=${process.env.NEXT_PUBLIC_HOST}/bff/callback`,
          method: 'GET'
        }),
        async onQueryStarted(state, { dispatch, queryFulfilled }) {
          const { data, meta } = await queryFulfilled
          // const responseUrl = (meta as { response: Response })?.response?.url ?? ""
          // const params = new URLSearchParams(new URL(responseUrl).search);
          // const loginChallenge = params.get("login_challenge") ?? "";
          dispatch(fetchLoginChallenge({ loginChallenge: '' , state}));
        }
      })
    })}
})

export const {
  useLoginMutation,
  useLoginGoogleQuery,
  useLogoutMutation,
  useGetAuthenticatedQuery,
  useSignUpQuery,
  useOnEmbedQuery,
  useOnImagePasteMutation,
} = api

export const {
  useFetchLoginChallengeMutation
} = hydra

export const {
  useGetProductsIdxQuery,
  useGetProductQuery,
  useGetProductsQuery,
  useSearchProductsQuery,
  useGetCategoriesQuery,
  useSendPostDemoMutation,
  useGetPosteosDemoQuery
} = demo