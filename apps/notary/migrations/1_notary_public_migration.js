const NotaryPublic = artifacts.require("NotaryPublic");

module.exports = function (deployer) {
  deployer.deploy(NotaryPublic);
};
