import { PropsWithChildren } from "react";

export const Main = ({ children }: PropsWithChildren<{}>) => {
  return (
    <div className="w-full flex">
      <main className="mx-auto w-full max-w-[1440px] p-10 min-h-[60vh] max-lg:!min-h-[50vh] max-lg:px-4 max-lg:pt-4 max-lg:pb-8">
        {children}
      </main>
    </div>
  );
};
