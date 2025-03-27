import create from 'zustand';
import { persist } from 'zustand/middleware';
import { Product, UserProfile } from '../types';

interface AppState {
  // User state
  user: UserProfile | null;
  isAuthenticated: boolean;
  setUser: (user: UserProfile | null) => void;
  
  // Product state
  products: Product[];
  selectedProduct: Product | null;
  setProducts: (products: Product[]) => void;
  setSelectedProduct: (product: Product | null) => void;
  
  // UI state
  isLoading: boolean;
  error: string | null;
  setIsLoading: (loading: boolean) => void;
  setError: (error: string | null) => void;
  
  // Cart state
  cart: Product[];
  addToCart: (product: Product) => void;
  removeFromCart: (productId: string) => void;
  clearCart: () => void;
  
  // Web3 state
  walletAddress: string | null;
  chainId: number | null;
  setWalletAddress: (address: string | null) => void;
  setChainId: (chainId: number | null) => void;
}

const useStore = create<AppState>()(
  persist(
    (set) => ({
      // User state
      user: null,
      isAuthenticated: false,
      setUser: (user) => set({ user, isAuthenticated: !!user }),
      
      // Product state
      products: [],
      selectedProduct: null,
      setProducts: (products) => set({ products }),
      setSelectedProduct: (product) => set({ selectedProduct: product }),
      
      // UI state
      isLoading: false,
      error: null,
      setIsLoading: (loading) => set({ isLoading: loading }),
      setError: (error) => set({ error }),
      
      // Cart state
      cart: [],
      addToCart: (product) =>
        set((state) => ({
          cart: [...state.cart, product],
        })),
      removeFromCart: (productId) =>
        set((state) => ({
          cart: state.cart.filter((p) => p.id !== productId),
        })),
      clearCart: () => set({ cart: [] }),
      
      // Web3 state
      walletAddress: null,
      chainId: null,
      setWalletAddress: (address) => set({ walletAddress: address }),
      setChainId: (chainId) => set({ chainId }),
    }),
    {
      name: 'revibe-store',
      partialize: (state) => ({
        user: state.user,
        isAuthenticated: state.isAuthenticated,
        walletAddress: state.walletAddress,
        chainId: state.chainId,
      }),
    }
  )
);

export default useStore; 