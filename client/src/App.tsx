/* eslint-disable @typescript-eslint/no-explicit-any */
import { useState } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Checkbox } from "@/components/ui/checkbox";
import { Alert, AlertDescription, AlertTitle } from "@/components/ui/alert";
import { Loader2, Music, Youtube, AlertTriangle } from "lucide-react";
import { useConvertSpotifyToYoutubeMutation } from "@/app/tool/toolApi";
import { toast } from "react-toastify";

export default function SpotifyToYouTube() {
  const [convertSpotifyToYoutube, { isLoading }] =
    useConvertSpotifyToYoutubeMutation();

  const [playlistUrl, setPlaylistUrl] = useState("");
  const [termsAccepted, setTermsAccepted] = useState(false);

  const handleConvert = async () => {
    if (!playlistUrl) {
      toast.error("Please provide a Spotify playlist URL.");
      return;
    }

    if (!termsAccepted) {
      toast.error("Please acknowledge the terms before converting.");
      return;
    }

    try {
      const result = await convertSpotifyToYoutube({
        playlistUrl: playlistUrl,
      }).unwrap();
      toast.success(result.message);
      if (result.youtubePlaylistUrl) {
        toast.info(`YouTube Playlist URL: ${result.youtubePlaylistUrl}`);
      }
      // Reset form after successful conversion
      setPlaylistUrl("");
      setTermsAccepted(false);
    } catch (err) {
      console.error("Failed to convert playlist:", err);
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      if ((err as any).data?.message) {
        toast.error(`Conversion failed: ${(err as any).data.message}`);
      } else {
        toast.error("Failed to convert playlist. Please try again.");
      }
    }
  };

  return (
    <div className="container mx-auto p-4">
      <Card className="w-full max-w-2xl mx-auto bg-zinc-800 text-white">
        <CardHeader>
          <CardTitle className="text-2xl font-bold flex items-center">
            <Music className="w-6 h-6 mr-2 text-green-500" />
            <span>Spotify</span>
            <span className="mx-2">to</span>
            <Youtube className="w-6 h-6 mr-2 text-red-600" />
            <span>YouTube</span>
          </CardTitle>
          <CardDescription className="text-zinc-400">
            Convert your Spotify playlists to YouTube playlists
          </CardDescription>
        </CardHeader>
        <CardContent className="space-y-4">
          <Alert className="bg-yellow-900 border-yellow-700">
            <AlertTriangle className="h-4 w-4" />
            <AlertTitle>Important</AlertTitle>
            <AlertDescription>
              The Spotify playlist must be set to public for this tool to work.
            </AlertDescription>
          </Alert>
          <div className="space-y-2">
            <Label htmlFor="playlist-url">Spotify Playlist URL</Label>
            <Input
              id="playlist-url"
              placeholder="https://open.spotify.com/playlist/..."
              value={playlistUrl}
              onChange={(e) => setPlaylistUrl(e.target.value)}
              className="bg-zinc-700 border-zinc-600 text-white"
            />
          </div>
          <div className="flex items-center space-x-2">
            <Checkbox
              id="terms"
              checked={termsAccepted}
              onCheckedChange={(checked) =>
                setTermsAccepted(checked as boolean)
              }
            />
            <Label htmlFor="terms" className="text-sm">
              I acknowledge that this process may take some time depending on
              the playlist size
            </Label>
          </div>
        </CardContent>
        <CardFooter className="flex flex-col items-stretch space-y-4">
          <Button
            onClick={handleConvert}
            disabled={!playlistUrl || isLoading || !termsAccepted}
            className="w-full bg-red-600 hover:bg-red-700"
          >
            {isLoading ? (
              <>
                <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                Converting...
              </>
            ) : (
              "Convert to YouTube Playlist"
            )}
          </Button>
        </CardFooter>
      </Card>
    </div>
  );
}