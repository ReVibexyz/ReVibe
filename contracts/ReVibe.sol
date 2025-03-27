// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/Counters.sol";

contract ReVibe is ERC721, ReentrancyGuard, Ownable {
    using Counters for Counters.Counter;
    Counters.Counter private _tokenIds;
    
    struct Product {
        uint256 id;
        address seller;
        uint256 price;
        bool isAuthenticated;
        bool isSold;
        string metadata;
    }
    
    mapping(uint256 => Product) public products;
    mapping(address => uint256[]) public userProducts;
    mapping(address => uint256[]) public userPurchases;
    
    uint256 public platformFee = 25; // 2.5%
    
    event ProductListed(uint256 indexed tokenId, address indexed seller, uint256 price);
    event ProductSold(uint256 indexed tokenId, address indexed seller, address indexed buyer, uint256 price);
    event ProductAuthenticated(uint256 indexed tokenId, bool authenticated);
    event PriceUpdated(uint256 indexed tokenId, uint256 newPrice);
    
    constructor() ERC721("ReVibe", "RVB") {}
    
    function listProduct(uint256 price, string memory metadata) external returns (uint256) {
        require(price > 0, "Price must be greater than 0");
        
        _tokenIds.increment();
        uint256 newTokenId = _tokenIds.current();
        
        _safeMint(msg.sender, newTokenId);
        
        products[newTokenId] = Product({
            id: newTokenId,
            seller: msg.sender,
            price: price,
            isAuthenticated: false,
            isSold: false,
            metadata: metadata
        });
        
        userProducts[msg.sender].push(newTokenId);
        
        emit ProductListed(newTokenId, msg.sender, price);
        
        return newTokenId;
    }
    
    function buyProduct(uint256 tokenId) external payable nonReentrant {
        Product storage product = products[tokenId];
        require(product.seller != address(0), "Product does not exist");
        require(!product.isSold, "Product already sold");
        require(msg.value >= product.price, "Insufficient payment");
        require(product.isAuthenticated, "Product not authenticated");
        
        address seller = product.seller;
        uint256 price = product.price;
        
        product.isSold = true;
        
        // Calculate platform fee
        uint256 fee = (price * platformFee) / 1000;
        uint256 sellerAmount = price - fee;
        
        // Transfer payment
        payable(owner()).transfer(fee);
        payable(seller).transfer(sellerAmount);
        
        // Transfer NFT
        _transfer(seller, msg.sender, tokenId);
        
        userPurchases[msg.sender].push(tokenId);
        
        emit ProductSold(tokenId, seller, msg.sender, price);
    }
    
    function authenticateProduct(uint256 tokenId, bool authenticated) external onlyOwner {
        require(products[tokenId].seller != address(0), "Product does not exist");
        products[tokenId].isAuthenticated = authenticated;
        emit ProductAuthenticated(tokenId, authenticated);
    }
    
    function updatePrice(uint256 tokenId, uint256 newPrice) external {
        require(products[tokenId].seller == msg.sender, "Not the seller");
        require(!products[tokenId].isSold, "Product already sold");
        require(newPrice > 0, "Price must be greater than 0");
        
        products[tokenId].price = newPrice;
        emit PriceUpdated(tokenId, newPrice);
    }
    
    function getUserProducts(address user) external view returns (uint256[] memory) {
        return userProducts[user];
    }
    
    function getUserPurchases(address user) external view returns (uint256[] memory) {
        return userPurchases[user];
    }
    
    function updatePlatformFee(uint256 newFee) external onlyOwner {
        require(newFee <= 100, "Fee too high"); // Max 10%
        platformFee = newFee;
    }
    
    function withdraw() external onlyOwner {
        uint256 balance = address(this).balance;
        payable(owner()).transfer(balance);
    }
} 