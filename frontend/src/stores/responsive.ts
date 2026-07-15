import { useBreakpoints } from '@vueuse/core';
import { defineStore } from 'pinia';

export const useResponsive = defineStore('responsive', () => {
  const breakpoints = useBreakpoints({
    mobile: 0,
    tablet: 640,
    laptop: 1024,
    desktop: 1280,
  });

  const activeBreakpoint = breakpoints.active();
  const isMobile = breakpoints.smaller('tablet');
  const isDesktop = breakpoints.greaterOrEqual('tablet');

  return { activeBreakpoint, isMobile, isDesktop };
});
