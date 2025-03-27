import axios from 'axios';
import { Product, UserProfile } from '../types';

const API_URL = process.env.REACT_APP_API_URL || 'http://localhost:8080';

const api = axios.create({
  baseURL: API_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Request interceptor for adding auth token
api.interceptors.request.use((config) => {
  const token = localStorage.getItem('token');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

// Response interceptor for handling errors
api.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 401) {
      localStorage.removeItem('token');
      window.location.href = '/login';
    }
    return Promise.reject(error);
  }
);

// Auth API
export const authAPI = {
  login: async (walletAddress: string, signature: string) => {
    const response = await api.post('/auth/login', { walletAddress, signature });
    return response.data;
  },
  
  verify: async (token: string) => {
    const response = await api.post('/auth/verify', { token });
    return response.data;
  },
};

// Product API
export const productAPI = {
  getProducts: async (params?: {
    category?: string;
    minPrice?: number;
    maxPrice?: number;
    search?: string;
  }) => {
    const response = await api.get('/products', { params });
    return response.data;
  },
  
  getProduct: async (id: string) => {
    const response = await api.get(`/products/${id}`);
    return response.data;
  },
  
  createProduct: async (product: Omit<Product, 'id'>) => {
    const response = await api.post('/products', product);
    return response.data;
  },
  
  updateProduct: async (id: string, product: Partial<Product>) => {
    const response = await api.put(`/products/${id}`, product);
    return response.data;
  },
  
  deleteProduct: async (id: string) => {
    const response = await api.delete(`/products/${id}`);
    return response.data;
  },
  
  authenticateProduct: async (id: string) => {
    const response = await api.post(`/products/${id}/authenticate`);
    return response.data;
  },
};

// User API
export const userAPI = {
  getUser: async (walletAddress: string) => {
    const response = await api.get(`/users/${walletAddress}`);
    return response.data;
  },
  
  updateUser: async (walletAddress: string, data: Partial<UserProfile>) => {
    const response = await api.put(`/users/${walletAddress}`, data);
    return response.data;
  },
  
  getUserProducts: async (walletAddress: string) => {
    const response = await api.get(`/users/${walletAddress}/products`);
    return response.data;
  },
  
  getUserOrders: async (walletAddress: string) => {
    const response = await api.get(`/users/${walletAddress}/orders`);
    return response.data;
  },
};

// Upload API
export const uploadAPI = {
  uploadImage: async (file: File) => {
    const formData = new FormData();
    formData.append('image', file);
    
    const response = await api.post('/upload', formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    });
    return response.data;
  },
};

export default api; 