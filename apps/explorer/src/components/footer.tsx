import Image from "next/image";
import Link from "next/link";

export const Footer = ({}) => {
  return (
    <div className="bg-[#F5F6F7] w-full flex">
      <footer className="mx-auto w-full max-w-[1440px] px-[90px] py-10">
        <div className="border-b border-b-[#D8DBDE] flex pb-[100px] items-start">
          <Image
            src="/logo-simple.svg"
            width={110}
            height={28}
            alt="Logo Simple"
          />

          <div className="ml-auto grid grid-flow-col auto-cols-max gap-x-[100px]">
            <div className="grid grid-flow-row auto-rows-min gap-y-3">
              <span className="text-[16px] font-medium">Company</span>
              <Link
                href={"https://gesia.io"}
                target="_blank"
                className="text-[12px]"
              >
                GESIA Platform
              </Link>
              <Link
                href={"https://www.gec-gesia.org/"}
                target="_blank"
                className="text-[12px]"
              >
                GEC
                <br />
                Green Earth Community
              </Link>
            </div>

            <div className="grid grid-flow-row auto-rows-min gap-y-3">
              <span className="text-[16px] font-medium">Service</span>
              <Link href={"/"} className="text-[12px]">
                Carbon Blockchain Explorer
              </Link>
              <Link
                href={"https://carbonmonster.kr"}
                target="_blank"
                className="text-[12px]"
              >
                Carbon Monster
              </Link>
            </div>

            <div className="grid grid-flow-row auto-rows-min gap-y-3">
              <span className="text-[16px] font-medium">Social</span>
              <Link
                href={"https://github.com/gesia-platform"}
                target="_blank"
                className="text-[12px]"
              >
                <Image src="/github.png" width={24} height={24} alt="Github" />
              </Link>
            </div>
            <div className="grid grid-flow-row auto-rows-auto gap-y-3"></div>
          </div>
        </div>

        <span className="mt-3 text-[#777A7D] text-[14px]">
          Copyright © 2024 GESIA All rights reserved.
        </span>
      </footer>
    </div>
  );
};
