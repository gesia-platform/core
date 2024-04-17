import { Footer } from "@/components/footer";
import { Main } from "@/components/main";
import { HomeCarbonStats } from "@/containers/home/carbon-stats";
import { LastChainData } from "@/containers/home/last-chain-data";
import { HomeSearch } from "@/containers/home/search";

export default function Home() {
  return (
    <Main>
      <h1 className="text-[24px] leading-[32px] font-medium text-[#1C1E20] mt-5">
        The GESIA Blockchain Explorer
      </h1>

      <HomeSearch />
      <HomeCarbonStats />
      <LastChainData />
    </Main>
  );
}
