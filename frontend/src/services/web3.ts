import { ethers } from 'ethers';
import ReVibeABI from '../../contracts/artifacts/ReVibe.json';

const CONTRACT_ADDRESS = process.env.REACT_APP_CONTRACT_ADDRESS || '';
const CHAIN_ID = process.env.REACT_APP_CHAIN_ID || '1';

class Web3Service {
  private provider: ethers.providers.Web3Provider | null = null;
  private signer: ethers.Signer | null = null;
  private contract: ethers.Contract | null = null;

  async init() {
    if (!window.ethereum) {
      throw new Error('Please install MetaMask or another Web3 wallet');
    }

    this.provider = new ethers.providers.Web3Provider(window.ethereum);
    this.signer = this.provider.getSigner();
    this.contract = new ethers.Contract(
      CONTRACT_ADDRESS,
      ReVibeABI.abi,
      this.signer
    );
  }

  async connect() {
    if (!this.provider) {
      await this.init();
    }

    try {
      await window.ethereum.request({
        method: 'eth_requestAccounts',
      });
    } catch (error) {
      throw new Error('Failed to connect wallet');
    }
  }

  async getAccount() {
    if (!this.signer) {
      throw new Error('Wallet not connected');
    }
    return await this.signer.getAddress();
  }

  async getChainId() {
    if (!this.provider) {
      throw new Error('Provider not initialized');
    }
    const network = await this.provider.getNetwork();
    return network.chainId;
  }

  async switchNetwork() {
    if (!window.ethereum) {
      throw new Error('Please install MetaMask or another Web3 wallet');
    }

    try {
      await window.ethereum.request({
        method: 'wallet_switchEthereumChain',
        params: [{ chainId: `0x${parseInt(CHAIN_ID).toString(16)}` }],
      });
    } catch (error: any) {
      if (error.code === 4902) {
        // Chain not added to MetaMask
        throw new Error('Please add the network to your wallet');
      }
      throw error;
    }
  }

  // Contract Methods
  async listProduct(
    name: string,
    description: string,
    price: string,
    images: string[]
  ) {
    if (!this.contract) {
      throw new Error('Contract not initialized');
    }

    const tx = await this.contract.listProduct(
      name,
      description,
      ethers.utils.parseEther(price),
      images
    );
    return await tx.wait();
  }

  async buyProduct(productId: string, price: string) {
    if (!this.contract) {
      throw new Error('Contract not initialized');
    }

    const tx = await this.contract.buyProduct(productId, {
      value: ethers.utils.parseEther(price),
    });
    return await tx.wait();
  }

  async authenticateProduct(productId: string) {
    if (!this.contract) {
      throw new Error('Contract not initialized');
    }

    const tx = await this.contract.authenticateProduct(productId);
    return await tx.wait();
  }

  async updatePrice(productId: string, newPrice: string) {
    if (!this.contract) {
      throw new Error('Contract not initialized');
    }

    const tx = await this.contract.updatePrice(
      productId,
      ethers.utils.parseEther(newPrice)
    );
    return await tx.wait();
  }

  async getUserProducts(walletAddress: string) {
    if (!this.contract) {
      throw new Error('Contract not initialized');
    }

    return await this.contract.getUserProducts(walletAddress);
  }

  async getUserPurchases(walletAddress: string) {
    if (!this.contract) {
      throw new Error('Contract not initialized');
    }

    return await this.contract.getUserPurchases(walletAddress);
  }

  // Event Listeners
  onProductListed(callback: (productId: string, seller: string) => void) {
    if (!this.contract) {
      throw new Error('Contract not initialized');
    }

    this.contract.on('ProductListed', callback);
  }

  onProductSold(callback: (productId: string, buyer: string) => void) {
    if (!this.contract) {
      throw new Error('Contract not initialized');
    }

    this.contract.on('ProductSold', callback);
  }

  onProductAuthenticated(callback: (productId: string) => void) {
    if (!this.contract) {
      throw new Error('Contract not initialized');
    }

    this.contract.on('ProductAuthenticated', callback);
  }

  // Cleanup
  removeAllListeners() {
    if (this.contract) {
      this.contract.removeAllListeners();
    }
  }
}

export const web3Service = new Web3Service();
export default web3Service; 