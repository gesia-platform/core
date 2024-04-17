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
      target={href.startsWith("http") ? "_blank" : undefined}
      className={`h-full px-5 font-[14px/18px] 
      ${focused ? "text-[#00C1B3] font-medium" : "text-[#1C1E20]"}
      max-md:h-auto max-md:!px-5 max-md:py-5`}
    >
      {children}
    </Link>
  );
};
