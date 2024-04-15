import { PropsWithChildren } from "react";

export const Main = ({ children }: PropsWithChildren<{}>) => {
  return (
    <div className="w-full flex">
      <main className="mx-auto w-full max-w-[1440px] p-10">{children}</main>
    </div>
  );
};
