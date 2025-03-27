import React, { useState } from 'react';
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import * as z from 'zod';
import { Product } from '../../types';

const productSchema = z.object({
  name: z.string().min(3, 'Name must be at least 3 characters'),
  description: z.string().min(10, 'Description must be at least 10 characters'),
  price: z.string().min(1, 'Price is required'),
  category: z.string().min(1, 'Category is required'),
  condition: z.string().min(1, 'Condition is required'),
  images: z.array(z.string()).min(1, 'At least one image is required'),
});

type ProductFormData = z.infer<typeof productSchema>;

interface ProductFormProps {
  onSubmit: (data: ProductFormData) => Promise<void>;
  initialData?: Partial<Product>;
}

const ProductForm: React.FC<ProductFormProps> = ({ onSubmit, initialData }) => {
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const {
    register,
    handleSubmit,
    formState: { errors },
    setValue,
  } = useForm<ProductFormData>({
    resolver: zodResolver(productSchema),
    defaultValues: initialData,
  });

  const handleImageUpload = async (e: React.ChangeEvent<HTMLInputElement>) => {
    const files = e.target.files;
    if (!files) return;

    const imageUrls: string[] = [];
    for (let i = 0; i < files.length; i++) {
      const file = files[i];
      // TODO: Implement image upload to IPFS or other storage
      const imageUrl = URL.createObjectURL(file);
      imageUrls.push(imageUrl);
    }

    setValue('images', imageUrls);
  };

  const onSubmitForm = async (data: ProductFormData) => {
    try {
      setIsSubmitting(true);
      setError(null);
      await onSubmit(data);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'An error occurred');
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <form onSubmit={handleSubmit(onSubmitForm)} className="space-y-6">
      <div>
        <label htmlFor="name" className="block text-sm font-medium text-gray-700">
          Product Name
        </label>
        <input
          type="text"
          id="name"
          {...register('name')}
          className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
        />
        {errors.name && (
          <p className="mt-1 text-sm text-red-600">{errors.name.message}</p>
        )}
      </div>

      <div>
        <label htmlFor="description" className="block text-sm font-medium text-gray-700">
          Description
        </label>
        <textarea
          id="description"
          rows={4}
          {...register('description')}
          className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
        />
        {errors.description && (
          <p className="mt-1 text-sm text-red-600">{errors.description.message}</p>
        )}
      </div>

      <div>
        <label htmlFor="price" className="block text-sm font-medium text-gray-700">
          Price (ETH)
        </label>
        <input
          type="number"
          id="price"
          step="0.000001"
          {...register('price')}
          className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
        />
        {errors.price && (
          <p className="mt-1 text-sm text-red-600">{errors.price.message}</p>
        )}
      </div>

      <div>
        <label htmlFor="category" className="block text-sm font-medium text-gray-700">
          Category
        </label>
        <select
          id="category"
          {...register('category')}
          className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
        >
          <option value="">Select a category</option>
          <option value="clothing">Clothing</option>
          <option value="accessories">Accessories</option>
          <option value="art">Art</option>
          <option value="collectibles">Collectibles</option>
        </select>
        {errors.category && (
          <p className="mt-1 text-sm text-red-600">{errors.category.message}</p>
        )}
      </div>

      <div>
        <label htmlFor="condition" className="block text-sm font-medium text-gray-700">
          Condition
        </label>
        <select
          id="condition"
          {...register('condition')}
          className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
        >
          <option value="">Select condition</option>
          <option value="new">New</option>
          <option value="like-new">Like New</option>
          <option value="good">Good</option>
          <option value="fair">Fair</option>
        </select>
        {errors.condition && (
          <p className="mt-1 text-sm text-red-600">{errors.condition.message}</p>
        )}
      </div>

      <div>
        <label htmlFor="images" className="block text-sm font-medium text-gray-700">
          Images
        </label>
        <input
          type="file"
          id="images"
          accept="image/*"
          multiple
          onChange={handleImageUpload}
          className="mt-1 block w-full"
        />
        {errors.images && (
          <p className="mt-1 text-sm text-red-600">{errors.images.message}</p>
        )}
      </div>

      {error && (
        <div className="text-red-600 text-sm">{error}</div>
      )}

      <button
        type="submit"
        disabled={isSubmitting}
        className="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 disabled:opacity-50"
      >
        {isSubmitting ? 'Uploading...' : 'Upload Product'}
      </button>
    </form>
  );
};

export default ProductForm; 