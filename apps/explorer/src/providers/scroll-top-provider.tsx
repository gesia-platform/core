"use client";

import { usePathname } from "next/navigation";
import { useEffect } from "react";

export default function ScrollTopProvider() {
  const pathname = usePathname();

  useEffect(() => {
    window.scrollTo({ top: 0 });
  }, [pathname]);

  return null;
}
