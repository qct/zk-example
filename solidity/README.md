npm init

# basic
npm install --save-dev hardhat ts-node typescript @types/node ethers dotenv

# hardhat-deploy
npm install --save-dev hardhat-deploy @nomiclabs/hardhat-ethers@npm:hardhat-deploy-ethers ethers

# test
npm install --save-dev chai @types/node @types/mocha @types/chai

# typechain
npm install --save-dev typechain @typechain/hardhat @typechain/ethers-v5 @nomiclabs/hardhat-solhint

# format
npm install --save-dev eslint prettier eslint-config-airbnb-typescript-prettier   
echo -e "
module.exports = {
  extends: "airbnb-typescript-prettier"
};
" > .eslintrc1.js
