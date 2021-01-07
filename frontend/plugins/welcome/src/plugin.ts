import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import LoginPage from './components/Login';

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', LoginPage);
    router.registerRoute('/welcompage', WelcomePage);
  },
});
