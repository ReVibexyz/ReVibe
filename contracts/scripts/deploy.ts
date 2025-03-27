import { ethers } from "hardhat";

async function main() {
  console.log("Deploying ReVibe contract...");

  const ReVibe = await ethers.getContractFactory("ReVibe");
  const revibe = await ReVibe.deploy();

  await revibe.deployed();

  console.log("ReVibe deployed to:", revibe.address);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
}); 