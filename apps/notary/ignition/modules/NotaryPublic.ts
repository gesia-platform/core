import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";

const NotaryPublicModule = buildModule("NotaryPublicModule", (m) => {
  const notaryPublic = m.contract("NotaryPublic", [], {});

  return { notaryPublic };
});

export default NotaryPublicModule;
