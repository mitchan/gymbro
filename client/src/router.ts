import { createRouter } from '@tanstack/react-router';

import { routeTree } from './routeTree.gen';

export const router = createRouter({
  routeTree,
  defaultPreload: 'intent',
  scrollRestoration: true,
  context: {
    // biome-ignore lint/style/noNonNullAssertion: This will be set after we wrap the app in an AuthProvider
    auth: undefined!,
  },
});

declare module '@tanstack/react-router' {
  interface Register {
    router: typeof router;
  }
}
