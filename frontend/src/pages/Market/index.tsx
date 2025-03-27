import React, { useState } from 'react';
import { Link } from 'react-router-dom';

interface Product {
  id: string;
  name: string;
  price: string;
  image: string;
  category: string;
}

const sampleProducts: Product[] = [
  {
    id: '1',
    name: 'Limited Edition Sneaker',
    price: '0.5 ETH',
    image: '/assets/products/sneaker.jpg',
    category: 'Footwear',
  },
  {
    id: '2',
    name: 'Collectible Art Print',
    price: '0.3 ETH',
    image: '/assets/products/art.jpg',
    category: 'Art',
  },
  {
    id: '3',
    name: 'Designer Watch',
    price: '1.2 ETH',
    image: '/assets/products/watch.jpg',
    category: 'Accessories',
  },
];

const Market: React.FC = () => {
  const [selectedCategory, setSelectedCategory] = useState<string>('all');

  const filteredProducts = selectedCategory === 'all'
    ? sampleProducts
    : sampleProducts.filter(product => product.category === selectedCategory);

  return (
    <div className="bg-white">
      <div className="max-w-2xl mx-auto py-16 px-4 sm:py-24 sm:px-6 lg:max-w-7xl lg:px-8">
        <div className="flex justify-between items-center">
          <h2 className="text-2xl font-extrabold tracking-tight text-gray-900">Market</h2>
          <div className="flex space-x-4">
            <select
              value={selectedCategory}
              onChange={(e) => setSelectedCategory(e.target.value)}
              className="mt-1 block w-full pl-3 pr-10 py-2 text-base border-gray-300 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm rounded-md"
            >
              <option value="all">All Categories</option>
              <option value="Footwear">Footwear</option>
              <option value="Art">Art</option>
              <option value="Accessories">Accessories</option>
            </select>
          </div>
        </div>

        <div className="mt-6 grid grid-cols-1 gap-y-10 gap-x-6 sm:grid-cols-2 lg:grid-cols-4 xl:gap-x-8">
          {filteredProducts.map((product) => (
            <div key={product.id} className="group relative">
              <div className="w-full min-h-80 bg-gray-200 aspect-w-1 aspect-h-1 rounded-md overflow-hidden group-hover:opacity-75 lg:h-80 lg:aspect-none">
                <img
                  src={product.image}
                  alt={product.name}
                  className="w-full h-full object-center object-cover lg:w-full lg:h-full"
                />
              </div>
              <div className="mt-4 flex justify-between">
                <div>
                  <h3 className="text-sm text-gray-700">
                    <Link to={`/product/${product.id}`}>
                      <span aria-hidden="true" className="absolute inset-0" />
                      {product.name}
                    </Link>
                  </h3>
                  <p className="mt-1 text-sm text-gray-500">{product.category}</p>
                </div>
                <p className="text-sm font-medium text-gray-900">{product.price}</p>
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
};

export default Market; 