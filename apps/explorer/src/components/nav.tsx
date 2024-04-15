"use client";

import Image from "next/image";
import { NavLink } from "./nav-link";
import Link from "next/link";

export const Nav = ({}) => {
  return (
    <div className="w-full flex border-b border-b-[#EAECED] bg-white sticky top-0 z-20">
      <nav className="mx-auto w-full max-w-[1440px] px-10 py-4 flex items-center">
        <div>
          <Link href={"/"}>
            <Image src="/logo.svg" alt="Logo" width={234} height={26} />
          </Link>
        </div>
        <div className="ml-auto">
          <NavLink href="/">Home</NavLink>
          <NavLink href="/chains">Chains</NavLink>
          <NavLink href="/blocks">Blocks</NavLink>
          <NavLink href="/txs">Transactions</NavLink>
          <NavLink href="https://gesia.gitbook.io/gesia-project-kr">
            Docs
          </NavLink>
        </div>
      </nav>
    </div>
  );
};
