export interface Product {
  id: string;
  name: string;
  description: string;
  price: string;
  image: string;
  category: string;
}

export interface ProductDetails extends Product {
  seller: string;
  images: string[];
  authenticity: string;
  condition: string;
}

export interface UserProfile {
  name: string;
  avatar: string;
  walletAddress: string;
  listings: number;
  sales: number;
  purchases: number;
}

export interface NavigationItem {
  name: string;
  href: string;
} 