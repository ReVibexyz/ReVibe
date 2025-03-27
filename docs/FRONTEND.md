# ReVibe Frontend Development Guide

## Project Structure
```
frontend/
├── public/              # Static files
├── src/
│   ├── components/     # Reusable UI components
│   ├── pages/         # Page components
│   ├── hooks/         # Custom React hooks
│   ├── services/      # API services
│   ├── utils/         # Utility functions
│   ├── types/         # TypeScript types
│   ├── assets/        # Static assets
│   ├── App.tsx        # Root component
│   └── index.tsx      # Entry point
├── package.json       # Dependencies
└── tsconfig.json      # TypeScript config
```

## Technology Stack

### Core
- React 18
- TypeScript
- React Router v6
- Tailwind CSS
- Headless UI
- Heroicons

### State Management
- React Context
- React Query
- Zustand

### Web3 Integration
- ethers.js
- Web3Modal
- wagmi

### Form Handling
- React Hook Form
- Zod validation

### API Client
- Axios
- React Query

## Component Guidelines

### Component Structure
```typescript
import React from 'react';
import { useHook } from '@/hooks/useHook';
import { ComponentProps } from '@/types';

interface Props extends ComponentProps {
  // Component props
}

export const Component: React.FC<Props> = ({ prop1, prop2 }) => {
  const { data, isLoading } = useHook();

  if (isLoading) {
    return <LoadingSpinner />;
  }

  return (
    <div>
      {/* Component JSX */}
    </div>
  );
};
```

### Styling Guidelines
- Use Tailwind CSS classes
- Follow mobile-first approach
- Maintain consistent spacing
- Use design system tokens

### Component Organization
- Atomic design principles
- Feature-based organization
- Shared components library

## State Management

### Global State
```typescript
// store/index.ts
import create from 'zustand';

interface Store {
  user: User | null;
  setUser: (user: User) => void;
}

export const useStore = create<Store>((set) => ({
  user: null,
  setUser: (user) => set({ user }),
}));
```

### API State
```typescript
// hooks/useProducts.ts
import { useQuery } from 'react-query';

export const useProducts = () => {
  return useQuery('products', fetchProducts);
};
```

## Web3 Integration

### Wallet Connection
```typescript
// hooks/useWallet.ts
import { useAccount, useConnect } from 'wagmi';

export const useWallet = () => {
  const { address, isConnected } = useAccount();
  const { connect } = useConnect();

  return {
    address,
    isConnected,
    connect,
  };
};
```

### Contract Interaction
```typescript
// services/contract.ts
import { ethers } from 'ethers';

export const contractService = {
  async buyProduct(tokenId: string, price: string) {
    const contract = getContract();
    const tx = await contract.buyProduct(tokenId, {
      value: ethers.utils.parseEther(price),
    });
    return tx.wait();
  },
};
```

## API Integration

### API Client
```typescript
// services/api.ts
import axios from 'axios';

export const api = axios.create({
  baseURL: process.env.REACT_APP_API_URL,
});

api.interceptors.request.use((config) => {
  const token = localStorage.getItem('token');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});
```

### API Hooks
```typescript
// hooks/useApi.ts
import { useMutation, useQuery } from 'react-query';

export const useProducts = () => {
  return useQuery('products', () => api.get('/products'));
};

export const useCreateProduct = () => {
  return useMutation((data) => api.post('/products', data));
};
```

## Form Handling

### Form Component
```typescript
// components/ProductForm.tsx
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';

export const ProductForm = () => {
  const { register, handleSubmit, errors } = useForm({
    resolver: zodResolver(productSchema),
  });

  const onSubmit = (data) => {
    // Handle form submission
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      {/* Form fields */}
    </form>
  );
};
```

## Testing

### Component Testing
```typescript
// __tests__/ProductCard.test.tsx
import { render, screen } from '@testing-library/react';
import { ProductCard } from '@/components/ProductCard';

describe('ProductCard', () => {
  it('renders product information', () => {
    render(<ProductCard product={mockProduct} />);
    expect(screen.getByText(mockProduct.name)).toBeInTheDocument();
  });
});
```

### Integration Testing
```typescript
// __tests__/ProductFlow.test.tsx
import { render, fireEvent } from '@testing-library/react';
import { ProductFlow } from '@/components/ProductFlow';

describe('ProductFlow', () => {
  it('completes purchase flow', async () => {
    render(<ProductFlow />);
    // Test implementation
  });
});
```

## Performance Optimization

### Code Splitting
```typescript
// App.tsx
import { lazy, Suspense } from 'react';

const ProductDetails = lazy(() => import('@/pages/ProductDetails'));

function App() {
  return (
    <Suspense fallback={<LoadingSpinner />}>
      <ProductDetails />
    </Suspense>
  );
}
```

### Image Optimization
```typescript
// components/ProductImage.tsx
import { Image } from '@/components/Image';

export const ProductImage = ({ src, alt }) => {
  return (
    <Image
      src={src}
      alt={alt}
      loading="lazy"
      sizes="(max-width: 768px) 100vw, (max-width: 1200px) 50vw, 33vw"
    />
  );
};
```

## Error Handling

### Error Boundary
```typescript
// components/ErrorBoundary.tsx
import { ErrorBoundary } from 'react-error-boundary';

export const AppErrorBoundary = ({ children }) => {
  return (
    <ErrorBoundary
      FallbackComponent={ErrorFallback}
      onReset={() => window.location.reload()}
    >
      {children}
    </ErrorBoundary>
  );
};
```

### API Error Handling
```typescript
// hooks/useApi.ts
export const useApi = () => {
  const handleError = (error) => {
    if (error.response?.status === 401) {
      // Handle unauthorized
    }
    // Handle other errors
  };

  return {
    get: async (url) => {
      try {
        const response = await api.get(url);
        return response.data;
      } catch (error) {
        handleError(error);
        throw error;
      }
    },
  };
};
```

## Accessibility

### ARIA Labels
```typescript
// components/Button.tsx
export const Button = ({ children, ...props }) => {
  return (
    <button
      aria-label={props['aria-label']}
      role="button"
      {...props}
    >
      {children}
    </button>
  );
};
```

### Keyboard Navigation
```typescript
// components/ProductGrid.tsx
export const ProductGrid = ({ products }) => {
  return (
    <div role="grid" aria-label="Products">
      {products.map((product) => (
        <ProductCard
          key={product.id}
          product={product}
          tabIndex={0}
        />
      ))}
    </div>
  );
};
```

## Deployment

### Build Process
```bash
# Build for production
npm run build

# Preview production build
npm run preview
```

### Environment Variables
```env
REACT_APP_API_URL=http://localhost:8080
REACT_APP_CHAIN_ID=1
REACT_APP_NETWORK_NAME=Ethereum Mainnet
REACT_APP_INFURA_ID=your-infura-id
REACT_APP_CONTRACT_ADDRESS=your-contract-address
```

## Monitoring

### Error Tracking
```typescript
// utils/errorTracking.ts
import * as Sentry from '@sentry/react';

export const initErrorTracking = () => {
  Sentry.init({
    dsn: process.env.REACT_APP_SENTRY_DSN,
    environment: process.env.NODE_ENV,
  });
};
```

### Analytics
```typescript
// utils/analytics.ts
import { track } from '@/services/analytics';

export const trackEvent = (eventName, properties) => {
  track(eventName, {
    ...properties,
    timestamp: new Date().toISOString(),
  });
};
``` 