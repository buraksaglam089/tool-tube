import { configureStore } from "@reduxjs/toolkit";
import { authApi } from "./auth/authApi";
import { toolApi } from "./tool/toolApi";

export const store = configureStore({
  reducer: {
    [authApi.reducerPath]: authApi.reducer,
    [toolApi.reducerPath]: toolApi.reducer,
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware().concat(authApi.middleware, toolApi.middleware),
});
