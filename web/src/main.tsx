import { createRoot } from 'react-dom/client'
import { App } from './App.tsx'
import { defaultTheme, Provider } from '@adobe/react-spectrum'

createRoot(document.getElementById('root')!).render(
  <Provider theme={defaultTheme} colorScheme='light' height={"100vh"}>
    <App />
  </Provider>,
)
