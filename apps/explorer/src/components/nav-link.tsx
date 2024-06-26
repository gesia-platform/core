"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";
import { ReactNode } from "react";

export const NavLink = ({
  href,
  children,
}: {
  href: string;
  children: ReactNode;
}) => {
  const pathname = usePathname();

  const focused = pathname === href;

  return (
    <Link
      href={href}
      className={`h-full px-5 font-[14px/18px] 
      ${focused ? "text-[#00C1B3] font-medium" : "text-[#1C1E20]"}
      max-lg:h-auto max-lg:!px-5 max-lg:py-5`}
    >
      {children}
    </Link>
  );
};
