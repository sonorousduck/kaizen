export interface NavigationItem {
  label: string
  to: string
  icon: string
}

export const navigationItems: ReadonlyArray<NavigationItem> = [
  { label: 'Home', to: '/home', icon: 'la-home-solid' },
  { label: 'Goals', to: '/goals', icon: 'gi-archery-target' },
  { label: 'Progress', to: '/progress', icon: 'gi-progression' },
];

export const bottomNavigationItems: ReadonlyArray<NavigationItem> = [
  { label: 'Settings', to: '/settings', icon: 'bi-gear' },
];
