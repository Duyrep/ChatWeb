"use client"

import Link from "next/link";
import { usePathname } from "next/navigation";

export default function Navigation() {
  const pathname = usePathname();

  return (
    <nav className="px-2 flex flex-wrap justify-between content-center fixed top-0 left-0 w-full h-12 bg-yellow-500">
      <div className="flex flex-wrap content-center">
        <b className="text-2xl"><Link href="/">DuyrepWeb</Link></b>
        <Link
          href="/chat"
          className={`flex flex-wrap content-center ml-4 px-2 rounded-md duration-300 ${pathname === "/chat" ? "after:opacity-100" : "hover:after:opacity-100"} active:bg-yellow-600 after:opacity-0 after:duration-300 after:h-1 after:w-full after:bg-black after:rounded-md`}
        >
          Chat
        </Link>
      </div>
      <div className="flex">
        <img
          className="h-10 rounded-3xl"
          src="https://cdn.discordapp.com/embed/avatars/0.png"
          alt="avatar"
        />
        <div className="flex flex-wrap pl-2 content-center">Duyrep</div>
      </div>
    </nav>
  );
}