export interface NavigationItem {
  label: string
  to: string
  icon: string
}

export const navigationItems: ReadonlyArray<NavigationItem> = [
  { label: 'Home', to: '/home', icon: 'home' },
  { label: 'Goals', to: '/goals', icon: 'target' },
];
