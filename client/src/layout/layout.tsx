import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Outlet } from "react-router-dom";

import { Sheet, SheetContent, SheetTrigger } from "@/components/ui/sheet";
import {
  Menu,
  Search,
  Youtube,
  Home,
  TrendingUp,
  SubscriptIcon,
  Library,
  User,
} from "lucide-react";
import { useGetCurrentUserQuery } from "@/app/auth/authApi";

export default function Layout() {
  const { data: user, isLoading } = useGetCurrentUserQuery({});
  console.log(user);

  return (
    <div className="flex w-full h-screen bg-zinc-900 text-white">
      {/* Sidebar */}
      <aside className="hidden w-64 bg-zinc-800 lg:flex flex-col border-r border-zinc-700">
        <div className="flex items-center p-4 border-b border-zinc-700">
          <Youtube className="h-8 w-8 text-red-600" />
          <span className="ml-2 text-xl font-bold">YT Tool</span>
        </div>
        <nav className="flex-grow mt-8">
          <Button variant="ghost" className="w-full justify-start text-lg">
            <Home className="mr-2 h-5 w-5" />
            Home
          </Button>
        </nav>
        <div className="p-4 border-t border-zinc-700">
          <div className="flex items-center space-x-2">
            <User className="h-5 w-5 text-zinc-400" />
            <span className="text-sm text-zinc-400 truncate">
              {isLoading ? "Loading..." : user?.Email || "Not logged in"}
            </span>
          </div>
        </div>
      </aside>

      <div className="flex flex-col flex-1">
        <header className="flex items-center justify-between p-4 bg-zinc-800 border-b border-zinc-700">
          <div className="flex items-center">
            <Sheet>
              <SheetTrigger asChild>
                <Button variant="ghost" size="icon" className="lg:hidden">
                  <Menu className="h-6 w-6" />
                </Button>
              </SheetTrigger>
              <SheetContent side="left" className="w-64 bg-zinc-800 p-0">
                <div className="flex items-center p-4 border-b border-zinc-700">
                  <Youtube className="h-8 w-8 text-red-600" />
                  <span className="ml-2 text-xl font-bold">YT Tool</span>
                </div>
                <nav className="mt-8">
                  <Button
                    variant="ghost"
                    className="w-full justify-start text-lg"
                  >
                    <Home className="mr-2 h-5 w-5" />
                    Home
                  </Button>
                  <Button
                    variant="ghost"
                    className="w-full justify-start text-lg"
                  >
                    <TrendingUp className="mr-2 h-5 w-5" />
                    Trending
                  </Button>
                  <Button
                    variant="ghost"
                    className="w-full justify-start text-lg"
                  >
                    <SubscriptIcon className="mr-2 h-5 w-5" />
                    Subscriptions
                  </Button>
                  <Button
                    variant="ghost"
                    className="w-full justify-start text-lg"
                  >
                    <Library className="mr-2 h-5 w-5" />
                    Library
                  </Button>
                </nav>
                {/* User email box in mobile sidebar */}
                <div className="absolute bottom-0 left-0 right-0 p-4 border-t border-zinc-700">
                  <div className="flex items-center space-x-2">
                    <User className="h-5 w-5 text-zinc-400" />
                    <span className="text-sm text-zinc-400 truncate">
                      {isLoading
                        ? "Loading..."
                        : user?.email || "Not logged in"}
                    </span>
                  </div>
                </div>
              </SheetContent>
            </Sheet>
            <Youtube className="h-8 w-8 text-red-600 ml-2 lg:hidden" />
            <span className="ml-2 text-xl font-bold hidden lg:inline">
              YT Tool
            </span>
          </div>
          <div className="flex items-center">
            <Input
              type="search"
              placeholder="Search..."
              className="w-64 mr-2 bg-zinc-700 border-zinc-600"
            />
            <Button size="icon" variant="ghost">
              <Search className="h-5 w-5" />
            </Button>
          </div>
        </header>

        {/* Main content */}
        <main className="flex-1 overflow-auto p-6">
          <Outlet />
        </main>
      </div>
    </div>
  );
}
