import { Footer } from "@/components/footer";
import { Main } from "@/components/main";
import { HomeCarbonStats } from "@/containers/home/carbon-stats";
import { LastChainData } from "@/containers/home/last-chain-data";
import { HomeSearch } from "@/containers/home/search";

export default function Home() {
  return (
    <Main>
      <HomeSearch />
      <HomeCarbonStats />
      <LastChainData />
    </Main>
  );
}
