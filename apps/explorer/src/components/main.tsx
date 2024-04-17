import { PropsWithChildren } from "react";

export const Main = ({ children }: PropsWithChildren<{}>) => {
  return (
    <div className="w-full flex">
      <main className="mx-auto w-full max-w-[1440px] p-10 min-h-[60vh] max-md:!min-h-[50vh] max-md:px-4 max-md:pt-4 max-md:pb-8">
        {children}
      </main>
    </div>
  );
};
