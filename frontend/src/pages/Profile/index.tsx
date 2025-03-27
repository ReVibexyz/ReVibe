import React from 'react';
import { Tab } from '@headlessui/react';

interface UserProfile {
  name: string;
  avatar: string;
  walletAddress: string;
  listings: number;
  sales: number;
  purchases: number;
}

const profile: UserProfile = {
  name: 'John Doe',
  avatar: '/assets/avatar.jpg',
  walletAddress: '0x1234...5678',
  listings: 12,
  sales: 8,
  purchases: 15,
};

function classNames(...classes: string[]) {
  return classes.filter(Boolean).join(' ');
}

const Profile: React.FC = () => {
  return (
    <div className="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
      {/* Profile header */}
      <div className="bg-white shadow">
        <div className="px-4 py-5 sm:px-6">
          <div className="flex items-center">
            <div className="flex-shrink-0 h-20 w-20">
              <img
                className="h-20 w-20 rounded-full"
                src={profile.avatar}
                alt={profile.name}
              />
            </div>
            <div className="ml-4">
              <h1 className="text-2xl font-bold text-gray-900">{profile.name}</h1>
              <p className="text-sm font-medium text-gray-500">
                {profile.walletAddress}
              </p>
            </div>
          </div>
          <div className="mt-6 grid grid-cols-3 gap-5 text-center">
            <div>
              <p className="text-2xl font-semibold text-gray-900">
                {profile.listings}
              </p>
              <p className="text-sm font-medium text-gray-500">Active Listings</p>
            </div>
            <div>
              <p className="text-2xl font-semibold text-gray-900">
                {profile.sales}
              </p>
              <p className="text-sm font-medium text-gray-500">Items Sold</p>
            </div>
            <div>
              <p className="text-2xl font-semibold text-gray-900">
                {profile.purchases}
              </p>
              <p className="text-sm font-medium text-gray-500">Items Purchased</p>
            </div>
          </div>
        </div>
      </div>

      {/* Tabs */}
      <div className="mt-6">
        <Tab.Group>
          <Tab.List className="flex space-x-1 rounded-xl bg-blue-900/20 p-1">
            <Tab
              className={({ selected }) =>
                classNames(
                  'w-full rounded-lg py-2.5 text-sm font-medium leading-5',
                  'ring-white ring-opacity-60 ring-offset-2 ring-offset-blue-400 focus:outline-none focus:ring-2',
                  selected
                    ? 'bg-white text-blue-700 shadow'
                    : 'text-blue-100 hover:bg-white/[0.12] hover:text-white'
                )
              }
            >
              Active Listings
            </Tab>
            <Tab
              className={({ selected }) =>
                classNames(
                  'w-full rounded-lg py-2.5 text-sm font-medium leading-5',
                  'ring-white ring-opacity-60 ring-offset-2 ring-offset-blue-400 focus:outline-none focus:ring-2',
                  selected
                    ? 'bg-white text-blue-700 shadow'
                    : 'text-blue-100 hover:bg-white/[0.12] hover:text-white'
                )
              }
            >
              Transaction History
            </Tab>
          </Tab.List>
          <Tab.Panels className="mt-2">
            <Tab.Panel
              className={classNames(
                'rounded-xl bg-white p-3',
                'ring-white ring-opacity-60 ring-offset-2 ring-offset-blue-400 focus:outline-none focus:ring-2'
              )}
            >
              {/* Active listings content */}
              <div className="space-y-4">
                <p className="text-gray-500">No active listings</p>
              </div>
            </Tab.Panel>
            <Tab.Panel
              className={classNames(
                'rounded-xl bg-white p-3',
                'ring-white ring-opacity-60 ring-offset-2 ring-offset-blue-400 focus:outline-none focus:ring-2'
              )}
            >
              {/* Transaction history content */}
              <div className="space-y-4">
                <p className="text-gray-500">No transaction history</p>
              </div>
            </Tab.Panel>
          </Tab.Panels>
        </Tab.Group>
      </div>
    </div>
  );
};

export default Profile; 