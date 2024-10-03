// src/services/auth.js
import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

export const authApi = createApi({
  reducerPath: "authApi",
  baseQuery: fetchBaseQuery({
    baseUrl: "http://localhost:8080",
    credentials: "include",
  }),
  endpoints: (builder) => ({
    getCurrentUser: builder.query({
      query: () => "/auth/me",
    }),
    login: builder.mutation({
      query: () => "/auth/google",
    }),
    logout: builder.mutation({
      query: () => "/auth/logout",
    }),
  }),
});

export const { useGetCurrentUserQuery, useLoginMutation, useLogoutMutation } =
  authApi;
