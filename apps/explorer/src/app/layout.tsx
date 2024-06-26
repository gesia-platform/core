import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import { Nav } from "@/components/nav";
import { Footer } from "@/components/footer";
import ChainProvider from "@/providers/chain-provider";
import ScrollTopProvider from "@/providers/scroll-top-provider";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "GESIA Blockchain Explorer",
  description: "GESIA Carbon Blockchain Explorer",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={inter.className}>
        <ScrollTopProvider />
        <ChainProvider>
          <Nav />
          {children}
          <Footer />
        </ChainProvider>
      </body>
    </html>
  );
}
