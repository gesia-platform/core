import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import { Nav } from "@/components/nav";
import { Footer } from "@/components/footer";
import ChainProvider from "@/providers/chain-provider";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Carbon Explorer | GESIA",
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
        <ChainProvider>
          <Nav />
          {children}
          <Footer />
        </ChainProvider>
      </body>
    </html>
  );
}
