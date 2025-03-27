import React, { useState } from 'react';
import { useParams } from 'react-router-dom';

interface ProductDetails {
  id: string;
  name: string;
  description: string;
  price: string;
  seller: string;
  images: string[];
  authenticity: string;
  condition: string;
}

const sampleProduct: ProductDetails = {
  id: '1',
  name: 'Limited Edition Sneaker',
  description: 'Exclusive limited edition sneaker with unique design and premium materials.',
  price: '0.5 ETH',
  seller: '0x1234...5678',
  images: [
    '/assets/products/sneaker-1.jpg',
    '/assets/products/sneaker-2.jpg',
    '/assets/products/sneaker-3.jpg',
  ],
  authenticity: 'Verified by AI Authentication System',
  condition: 'New',
};

const ProductDetails: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const [selectedImage, setSelectedImage] = useState(0);

  return (
    <div className="bg-white">
      <div className="max-w-2xl mx-auto py-16 px-4 sm:py-24 sm:px-6 lg:max-w-7xl lg:px-8">
        <div className="lg:grid lg:grid-cols-2 lg:gap-x-8 lg:items-start">
          {/* Image gallery */}
          <div className="flex flex-col">
            <div className="w-full aspect-w-1 aspect-h-1 bg-gray-200 rounded-lg overflow-hidden">
              <img
                src={sampleProduct.images[selectedImage]}
                alt={sampleProduct.name}
                className="w-full h-full object-center object-cover"
              />
            </div>
            <div className="mt-4 grid grid-cols-4 gap-2">
              {sampleProduct.images.map((image, index) => (
                <button
                  key={index}
                  onClick={() => setSelectedImage(index)}
                  className={`relative rounded-lg overflow-hidden ${
                    selectedImage === index
                      ? 'ring-2 ring-indigo-500'
                      : 'ring-1 ring-gray-200'
                  }`}
                >
                  <img
                    src={image}
                    alt={`View ${index + 1}`}
                    className="w-full h-full object-center object-cover"
                  />
                </button>
              ))}
            </div>
          </div>

          {/* Product info */}
          <div className="mt-10 px-4 sm:px-0 sm:mt-16 lg:mt-0">
            <h1 className="text-3xl font-extrabold tracking-tight text-gray-900">
              {sampleProduct.name}
            </h1>
            <div className="mt-3">
              <h2 className="sr-only">Product information</h2>
              <p className="text-3xl text-gray-900">{sampleProduct.price}</p>
            </div>

            <div className="mt-6">
              <h3 className="sr-only">Description</h3>
              <p className="text-base text-gray-900">{sampleProduct.description}</p>
            </div>

            <div className="mt-6">
              <div className="flex items-center">
                <h4 className="text-sm text-gray-900 font-medium">Seller:</h4>
                <p className="ml-2 text-sm text-gray-500">{sampleProduct.seller}</p>
              </div>
              <div className="mt-2 flex items-center">
                <h4 className="text-sm text-gray-900 font-medium">Authenticity:</h4>
                <p className="ml-2 text-sm text-gray-500">
                  {sampleProduct.authenticity}
                </p>
              </div>
              <div className="mt-2 flex items-center">
                <h4 className="text-sm text-gray-900 font-medium">Condition:</h4>
                <p className="ml-2 text-sm text-gray-500">{sampleProduct.condition}</p>
              </div>
            </div>

            <div className="mt-10">
              <button
                type="button"
                className="w-full bg-indigo-600 border border-transparent rounded-md py-3 px-8 flex items-center justify-center text-base font-medium text-white hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
              >
                Purchase Now
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default ProductDetails; 