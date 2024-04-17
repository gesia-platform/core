"use client";

import Image from "next/image";
import { NavLink } from "./nav-link";
import Link from "next/link";
import { Fragment, useState } from "react";

export const Nav = ({}) => {
  const [open, setOpen] = useState(false);
  const [navHeight, setNavHeight] = useState(0);

  const links = (
    <Fragment>
      <NavLink href="/">Home</NavLink>
      <NavLink href="/chains">Chains</NavLink>
      <NavLink href="/blocks">Blocks</NavLink>
      <NavLink href="/txs">Transactions</NavLink>
      <NavLink href="https://gesia.gitbook.io/gesia-project-kr">Docs</NavLink>
    </Fragment>
  );
  return (
    <div className="w-full flex border-b border-b-[#EAECED] bg-white sticky top-0 z-20">
      <nav
        className="mx-auto w-full max-w-[1440px] px-10 py-4 flex items-center max-md:!px-4 max-md:py-5"
        ref={(ref) => {
          if (ref) {
            setNavHeight(ref.clientHeight);
          }
        }}
      >
        <Link
          href={"/"}
          className="w-[234px] h-[26px] max-md:!w-[180px] max-md:!h-[20px] relative"
        >
          <Image src="/logo.svg" alt="Logo" fill />
        </Link>
        <div className="ml-auto">
          <div className="hidden max-md:flex">
            <button
              className="w-[30px] h-[30px] relative"
              onClick={() => {
                setOpen((open) => !open);
              }}
            >
              <Image src="/menu.png" fill alt="Menu" />
            </button>
          </div>

          <div className="max-md:hidden">{links}</div>
        </div>
      </nav>

      {open && (
        <div
          className="hidden max-md:flex fixed left-0 right-0 bottom-0 bg-[#00000033]"
          style={{
            top: navHeight,
          }}
          onClick={() => {
            setOpen(false);
          }}
        >
          <div className="bg-white flex flex-col w-full self-start">
            {links}
          </div>
        </div>
      )}
    </div>
  );
};
