export interface NavigationItem {
  label: string
  to: string
  icon: string
}

export const navigationItems: ReadonlyArray<NavigationItem> = [
  { label: 'Home', to: '/home', icon: 'home' },
  { label: 'Goals', to: '/goals', icon: 'target' },
  { label: 'Progress', to: '/progress', icon: 'yeet' },
];

export const bottomNavigationItems: ReadonlyArray<NavigationItem> = [
  { label: 'Settings', to: '/settings', icon: 'yeet' },
];
