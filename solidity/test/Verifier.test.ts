import * as fs from "fs";
import { deployments, ethers, getUnnamedAccounts } from "hardhat";
import { Verifier } from "../typechain-types";
import { expect } from "./chai-setup";
import { setupUsers } from "./utils";

const setup = deployments.createFixture(async () => {
    await deployments.fixture("Verifier");
    const contracts = {
        Verifier: <Verifier>await ethers.getContract("Verifier"),
    };
    const users = await setupUsers(await getUnnamedAccounts(), contracts);
    return { ...contracts, users };
});

describe("Verifier contract", function () {
    it("Verify should success", async function () {
        const proof = fs.readFileSync("./proof");
        const proofBytes = Uint8Array.from(proof);
        const fpSize = 4 * 8;

        // proof.Ar, proof.Bs, proof.Krs
        const a0 = proofBytes.slice(fpSize * 0, fpSize * 1);
        const a1 = proofBytes.slice(fpSize * 1, fpSize * 2);
        const b00 = proofBytes.slice(fpSize * 2, fpSize * 3);
        const b01 = proofBytes.slice(fpSize * 3, fpSize * 4);
        const b10 = proofBytes.slice(fpSize * 4, fpSize * 5);
        const b11 = proofBytes.slice(fpSize * 5, fpSize * 6);
        const c0 = proofBytes.slice(fpSize * 6, fpSize * 7);
        const c1 = proofBytes.slice(fpSize * 7, fpSize * 8);
        const input = 4;
        console.log("input: ", [input]);

        const { users, Verifier } = await setup();
        const result = await Verifier.verifyProof(
            [a0, a1],
            [
                [b00, b01],
                [b10, b11],
            ],
            [c0, c1],
            [input]
        );

        console.log("Verify result: ", result);
        expect(result).to.equal(true);
    });
    it("Verify should fail", async function () {
        const proof = fs.readFileSync("./proof");
        const proofBytes = Uint8Array.from(proof);
        const fpSize = 4 * 8;

        // proof.Ar, proof.Bs, proof.Krs
        const a0 = proofBytes.slice(fpSize * 0, fpSize * 1);
        const a1 = proofBytes.slice(fpSize * 1, fpSize * 2);
        const b00 = proofBytes.slice(fpSize * 2, fpSize * 3);
        const b01 = proofBytes.slice(fpSize * 3, fpSize * 4);
        const b10 = proofBytes.slice(fpSize * 4, fpSize * 5);
        const b11 = proofBytes.slice(fpSize * 5, fpSize * 6);
        const c0 = proofBytes.slice(fpSize * 6, fpSize * 7);
        const c1 = proofBytes.slice(fpSize * 7, fpSize * 8);

        // wrong input
        const input = 5;
        console.log("input: ", [input]);

        const { users, Verifier } = await setup();
        const result = await Verifier.verifyProof(
            [a0, a1],
            [
                [b00, b01],
                [b10, b11],
            ],
            [c0, c1],
            [input]
        );

        console.log("Verify result: ", result);
        expect(result).to.equal(false);
    });
});
