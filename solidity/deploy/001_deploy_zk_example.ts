import { DeployFunction } from "hardhat-deploy/types";
import { HardhatRuntimeEnvironment } from "hardhat/types";

const deployZkExampleVerifier: DeployFunction = async function (
    hre: HardhatRuntimeEnvironment
) {
    const { deployments, getNamedAccounts } = hre;
    const { deploy, log } = deployments;

    const { deployer } = await getNamedAccounts();
    const zkExampleVerifier = await deploy("Verifier", {
        from: deployer,
        args: [],
        log: true,
    });
    log(`Deployed zk example verifier to address ${zkExampleVerifier.address}`);
};

export default deployZkExampleVerifier;
deployZkExampleVerifier.tags = ["Verifier"];
