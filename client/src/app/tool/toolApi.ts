import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

// Helper function to extract playlist ID from Spotify URL
const extractPlaylistId = (url: string): string | null => {
  const match = url.match(/playlist\/([a-zA-Z0-9]+)/);
  return match ? match[1] : null;
};

// Define a type for the conversion request payload
interface ConversionRequest {
  playlistUrl: string;
}

// Define a type for the conversion response
interface ConversionResponse {
  message: string;
  youtubePlaylistUrl?: string;
}

export const toolApi = createApi({
  reducerPath: "toolApi",
  baseQuery: fetchBaseQuery({
    baseUrl: "http://localhost:8080",
    credentials: "include",
  }),
  endpoints: (builder) => ({
    convertSpotifyToYoutube: builder.mutation<
      ConversionResponse,
      ConversionRequest
    >({
      query: (conversionData) => {
        const playlistId = extractPlaylistId(conversionData.playlistUrl);
        if (!playlistId) {
          throw new Error("Invalid Spotify playlist URL");
        }
        return {
          url: "/tool/convert",
          method: "POST",
          body: { playlistId },
        };
      },
    }),
  }),
});

export const { useConvertSpotifyToYoutubeMutation } = toolApi;
